package server

import (
	"bytes"
	"fmt"
	auth "github.com/Nevoral/DNZ_app/internal/Authentification"
	"github.com/Nevoral/DNZ_app/internal/database"
	hand "github.com/Nevoral/DNZ_app/internal/handlers"
	zlog "github.com/Nevoral/DNZ_app/internal/logging"
	"github.com/Nevoral/DNZ_app/web/Home"
	"github.com/gofiber/fiber/v3"
	"os"
	"time"
)

func (s *FiberServer) Router() {
	s.App.Use(auth.JWTAuthMiddleware)

	s.App.Get("/js/*", hand.SendJs)
	s.App.Get("/assets", hand.SendAsset)
	s.App.Get("/css/*", hand.SendCss)
	//
	s.App.Get("/", s.HomePage)
	s.App.Get("/createuser", s.createUser)
	s.App.Post("/signup", s.SignUpUser)
	s.App.Get("/signuptab", s.SignUpTab)
	s.App.Post("/login", s.LoginUser)
	s.App.Get("/logintab", s.LogInTab)
	//	s.App.Get("/food", s.FoodProductList)
	//	s.App.Get("/drink", s.DrinkProductList)
	//	s.App.Post("/order", s.OrderLog)
	//	s.App.Post("/clear-order", s.CleanOrder)
	//	s.App.Get("/menu", s.SendForm)
	//	s.App.Post("/create-menu", s.CreateProductMenu)
	//	s.App.Get("/health", s.healthHandler)
	//	s.App.Get("/increment", s.increment)
	//	s.App.Get("/decrement", s.decrement)
	//	s.App.Post("/add-product", s.addProduct)
}

//	func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
//		return c.JSON(s.db.Health())
//	}
func (s *FiberServer) createUser(c fiber.Ctx) error {
	data, err := s.db.CreateUser(c.UserContext(), database.CreateUserParam{
		Username:     "Antak",
		Email:        "NevoralTomas@gmail.com",
		PasswordHash: "nevim",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return err
	}
	return c.JSON(data)
}

func (s *FiberServer) HomePage(c fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)

	var output bytes.Buffer
	for _, value := range Home.Page() {
		if err := value.Render(c.UserContext(), &output); err != nil {
			zlog.ErrorLog(err.Error())
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
	}
	return c.SendStream(&output)
}

func (s *FiberServer) SignUpUser(c fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	if c.FormValue("terms") != "on" {
		err := fmt.Errorf("nebyl dán souhlas s podmínkami")
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusNotAcceptable).SendString(err.Error())
	}

	if stat, err := auth.EmailTest(email); err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(stat).SendString(err.Error())
	}

	if stat, err := auth.UsernameTest(username); err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(stat).SendString(err.Error())
	}

	if stat, err := auth.PasswordTest(password); err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(stat).SendString(err.Error())
	}

	hashPassword, err := auth.GenerateHashPassword(password)
	if err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("problem s bezpečnostním Hashem")
	}

	jwt, err := auth.GenerateToken(email, os.Getenv("JWT_KEY"), time.Now().Add(time.Hour*48))
	if err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("problem s jwt tokenem")
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt-dnz",
		Value:    jwt,
		Expires:  time.Now().Add(48 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict", // Can be Lax or None as per your requirements
	})

	data, err := s.db.CreateUser(c.UserContext(), database.CreateUserParam{
		Username:     username,
		Email:        email,
		PasswordHash: hashPassword,
	})
	if err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("chyba v databázi")
	}
	fmt.Println(data)

	var output bytes.Buffer
	if err = Home.PopupWindowCon(
		Home.WelcomeTab(false),
		Home.EmailConfirm(email),
	).Render(c.UserContext(), &output); err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.SendStream(&output)
}

func (s *FiberServer) SignUpTab(c fiber.Ctx) error {
	var output bytes.Buffer
	if err := Home.PopupWindowCon(
		Home.AuthTab(false),
		Home.WelcomeTab(false),
	).Render(c.UserContext(), &output); err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.SendStream(&output)
}

func (s *FiberServer) LoginUser(c fiber.Ctx) error {
	email := c.FormValue("email")
	//password := c.FormValue("password")

	//hashPassword, err := Authentification.GenerateHashPassword(password)
	//if err != nil {
	//	zlog.ErrorLog(err.Error())
	//	return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	//}

	//data, err := s.db.CreateUser(c.UserContext(), database.CreateUserParam{
	//	Username:     username,
	//	Email:        email,
	//	PasswordHash: hashPassword,
	//})
	//if err != nil {
	//	zlog.ErrorLog(err.Error())
	//	return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	//}
	//fmt.Println(data)

	var output bytes.Buffer
	if err := Home.EmailConfirm(email).Render(c.UserContext(), &output); err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.SendStream(&output)
}

func (s *FiberServer) LogInTab(c fiber.Ctx) error {
	var output bytes.Buffer
	if err := Home.PopupWindowCon(
		Home.WelcomeTab(true),
		Home.AuthTab(true),
	).Render(c.UserContext(), &output); err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.SendStream(&output)
}

//
//func (s *FiberServer) FoodProductList(c *fiber.Ctx) error {
//	products, err := s.db.GetListOfProducts(1, "food")
//	if err != nil {
//		return err
//	}
//	var menuId int64 = 1
//	if len(products) != 0 {
//		menuId = products[0].Productmenuid
//	}
//	p := reg.ItemList(products, menuId, "orange-500", "food")
//	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
//
//	var output bytes.Buffer
//	if err = p.Render(context.Background(), &output); err != nil {
//		zlog.ErrorLog(err.Error())
//		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
//	}
//	return c.SendStream(&output)
//}
//
//func (s *FiberServer) DrinkProductList(c *fiber.Ctx) error {
//	products, err := s.db.GetListOfProducts(1, "drink")
//	if err != nil {
//		return err
//	}
//	var menuId int64 = 1
//	if len(products) != 0 {
//		menuId = products[0].Productmenuid
//	}
//	p := reg.ItemList(products, menuId, "sky-500", "drink")
//	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
//
//	var output bytes.Buffer
//	if err = p.Render(context.Background(), &output); err != nil {
//		zlog.ErrorLog(err.Error())
//		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
//	}
//	return c.SendStream(&output)
//}
//
//func (s *FiberServer) addProduct(c *fiber.Ctx) error {
//	title := c.FormValue("title")
//	category := c.FormValue("category")
//	price, err := strconv.Atoi(c.FormValue("price"))
//	if err != nil {
//		zlog.ErrorLog(err.Error())
//		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
//	}
//	err = s.db.CreateProduct(title, category, int64(price), 1)
//	if err != nil {
//		zlog.ErrorLog(err.Error())
//		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
//	}
//	if category == "food" {
//		return s.FoodProductList(c)
//	}
//	return s.DrinkProductList(c)
//}
//
//func (s *FiberServer) OrderLog(c *fiber.Ctx) error {
//	type response struct {
//		Order string `json:"order"`
//		Total int64  `json:"total"`
//	}
//	var r response
//
//	err := json.Unmarshal(c.Body(), &r)
//	ord := strings.Split(r.Order, ";")
//	for _, o := range ord {
//		row := strings.Split(o, ",")
//		quant, err := strconv.Atoi(row[0])
//		if err != nil {
//			zlog.ErrorLog(err.Error())
//			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
//		}
//		s.db.SetCount(row[1], int64(quant))
//	}
//
//	if err != nil {
//		zlog.ErrorLog(err.Error())
//		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
//	}
//	err = s.db.CreateOrderWithoutUser(1, r.Total, r.Order)
//	if err != nil {
//		zlog.ErrorLog(err.Error())
//		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
//	}
//
//	return c.SendString("OK")
//}
//
//func (s *FiberServer) SendForm(c *fiber.Ctx) error {
//	//fmt.Println("jsem tu")
//	//p := reg.ProductMenuForm()
//	//c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
//	//
//	//var output bytes.Buffer
//	//if err := p.Render(context.Background(), &output); err != nil {
//	//	zlog.ErrorLog(err.Error())
//	//	return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
//	//}
//	//return c.SendStream(&output)
//	return c.SendString("OK")
//}
//
//func (s *FiberServer) CreateProductMenu(c *fiber.Ctx) error {
//	startRegister, err := strconv.Atoi(c.FormValue("startRegister"))
//	if err != nil {
//		zlog.ErrorLog(err.Error())
//		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
//	}
//	menu, err := s.db.CreateProductMenu(int64(startRegister))
//	fmt.Println(menu)
//	if err != nil {
//		zlog.ErrorLog(err.Error())
//		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
//	}
//	return s.FoodProductList(c)
//}
//
//func (s *FiberServer) CleanOrder(c *fiber.Ctx) error {
//	cat := c.Query("cat")
//	msg := bytes.Split(c.Body()[1:len(c.Body())-1], []byte(";"))
//	for _, item := range msg {
//		val := bytes.Split(item, []byte(":"))
//		quant, err := strconv.ParseInt(string(val[1]), 10, 64)
//		if err != nil {
//			zlog.ErrorLog(err.Error())
//			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
//		}
//		id, err := strconv.ParseInt(string(val[0]), 10, 64)
//		if err != nil {
//			zlog.ErrorLog(err.Error())
//			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
//		}
//		for i := int64(0); i < quant; i++ {
//			s.db.GetCount(id, -1)
//		}
//	}
//	if cat == "drink" {
//		return s.DrinkProductList(c)
//	}
//	return s.FoodProductList(c)
//}
//
//func (s *FiberServer) increment(c *fiber.Ctx) error {
//	itemID := c.Query("itemId")
//	if itemID == "" {
//		return c.Status(fiber.StatusBadRequest).SendString("Item ID is required")
//	}
//
//	id, err := strconv.Atoi(itemID)
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).SendString("Invalid Item ID")
//	}
//	// Set the HTTP Content-Type header.
//	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
//
//	// Buffer to store the rendered HTML.
//	var output bytes.Buffer
//
//	p := reg.Count(fmt.Sprintf("%dx", s.db.GetCount(int64(id), 1)))
//
//	// Render the page into the buffer.
//	if err := p.Render(context.Background(), &output); err != nil {
//		// Handle the rendering error, e.g., by logging and returning an HTTP 500 error.
//		log.Printf("Error rendering page: %v", err)
//		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
//	}
//
//	// Send the rendered HTML as the response.
//	return c.SendStream(&output)
//}
//func (s *FiberServer) decrement(c *fiber.Ctx) error {
//	itemID := c.Query("itemId")
//	if itemID == "" {
//		return c.Status(fiber.StatusBadRequest).SendString("Item ID is required")
//	}
//
//	id, err := strconv.Atoi(itemID)
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).SendString("Invalid Item ID")
//	}
//	// Set the HTTP Content-Type header.
//	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
//
//	// Buffer to store the rendered HTML.
//	var output bytes.Buffer
//	p := reg.Count(fmt.Sprintf("%dx", s.db.GetCount(int64(id), -1)))
//
//	// Render the page into the buffer.
//	if err := p.Render(context.Background(), &output); err != nil {
//		// Handle the rendering error, e.g., by logging and returning an HTTP 500 error.
//		log.Printf("Error rendering page: %v", err)
//		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
//	}
//
//	// Send the rendered HTML as the response.
//	return c.SendStream(&output)
//}
