package server

import (
	hand "BuffetRegister/internal/handlers"
	zlog "BuffetRegister/internal/logging"
	page "BuffetRegister/web/pages"
	reg "BuffetRegister/web/pages/Register"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"strings"
)

func (s *FiberServer) Router() {
	s.App.Get("/js/*", hand.SendJs)
	s.App.Get("/assets/*", hand.SendAsset)
	s.App.Get("/css/*", hand.SendCss)

	s.App.Get("/", hand.SendHTML(page.Layout()...))
	s.App.Get("/food", s.FoodProductList)
	s.App.Get("/drink", s.DrinkProductList)
	s.App.Post("/order", s.OrderLog)
	s.App.Post("/clear-order", s.CleanOrder)
	s.App.Get("/menu", s.SendForm)
	s.App.Post("/create-menu", s.CreateProductMenu)
	s.App.Get("/health", s.healthHandler)
	s.App.Get("/increment", s.increment)
	s.App.Get("/decrement", s.decrement)
	s.App.Post("/add-product", s.addProduct)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

func (s *FiberServer) FoodProductList(c *fiber.Ctx) error {
	products, err := s.db.GetListOfProducts(1, "food")
	if err != nil {
		return err
	}
	var menuId int64 = 1
	if len(products) != 0 {
		menuId = products[0].Productmenuid
	}
	p := reg.ItemList(products, menuId, "orange-500", "food")
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)

	var output bytes.Buffer
	if err = p.Render(context.Background(), &output); err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.SendStream(&output)
}

func (s *FiberServer) DrinkProductList(c *fiber.Ctx) error {
	products, err := s.db.GetListOfProducts(1, "drink")
	if err != nil {
		return err
	}
	var menuId int64 = 1
	if len(products) != 0 {
		menuId = products[0].Productmenuid
	}
	p := reg.ItemList(products, menuId, "sky-500", "drink")
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)

	var output bytes.Buffer
	if err = p.Render(context.Background(), &output); err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.SendStream(&output)
}

func (s *FiberServer) addProduct(c *fiber.Ctx) error {
	title := c.FormValue("title")
	category := c.FormValue("category")
	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	err = s.db.CreateProduct(title, category, int64(price), 1)
	if err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if category == "food" {
		return s.FoodProductList(c)
	}
	return s.DrinkProductList(c)
}

func (s *FiberServer) OrderLog(c *fiber.Ctx) error {
	type response struct {
		Order string `json:"order"`
		Total int64  `json:"total"`
	}
	var r response

	err := json.Unmarshal(c.Body(), &r)
	ord := strings.Split(r.Order, ";")
	for _, o := range ord {
		row := strings.Split(o, ",")
		quant, err := strconv.Atoi(row[0])
		if err != nil {
			zlog.ErrorLog(err.Error())
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		s.db.SetCount(row[1], int64(quant))
	}

	if err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	err = s.db.CreateOrderWithoutUser(1, r.Total, r.Order)
	if err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("OK")
}

func (s *FiberServer) SendForm(c *fiber.Ctx) error {
	//fmt.Println("jsem tu")
	//p := reg.ProductMenuForm()
	//c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	//
	//var output bytes.Buffer
	//if err := p.Render(context.Background(), &output); err != nil {
	//	zlog.ErrorLog(err.Error())
	//	return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	//}
	//return c.SendStream(&output)
	return c.SendString("OK")
}

func (s *FiberServer) CreateProductMenu(c *fiber.Ctx) error {
	startRegister, err := strconv.Atoi(c.FormValue("startRegister"))
	if err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	menu, err := s.db.CreateProductMenu(int64(startRegister))
	fmt.Println(menu)
	if err != nil {
		zlog.ErrorLog(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return s.FoodProductList(c)
}

func (s *FiberServer) CleanOrder(c *fiber.Ctx) error {
	cat := c.Query("cat")
	msg := bytes.Split(c.Body()[1:len(c.Body())-1], []byte(";"))
	for _, item := range msg {
		val := bytes.Split(item, []byte(":"))
		quant, err := strconv.ParseInt(string(val[1]), 10, 64)
		if err != nil {
			zlog.ErrorLog(err.Error())
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		id, err := strconv.ParseInt(string(val[0]), 10, 64)
		if err != nil {
			zlog.ErrorLog(err.Error())
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		for i := int64(0); i < quant; i++ {
			s.db.GetCount(id, -1)
		}
	}
	if cat == "drink" {
		return s.DrinkProductList(c)
	}
	return s.FoodProductList(c)
}

func (s *FiberServer) increment(c *fiber.Ctx) error {
	itemID := c.Query("itemId")
	if itemID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Item ID is required")
	}

	id, err := strconv.Atoi(itemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Item ID")
	}
	// Set the HTTP Content-Type header.
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)

	// Buffer to store the rendered HTML.
	var output bytes.Buffer

	p := reg.Count(fmt.Sprintf("%dx", s.db.GetCount(int64(id), 1)))

	// Render the page into the buffer.
	if err := p.Render(context.Background(), &output); err != nil {
		// Handle the rendering error, e.g., by logging and returning an HTTP 500 error.
		log.Printf("Error rendering page: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Send the rendered HTML as the response.
	return c.SendStream(&output)
}
func (s *FiberServer) decrement(c *fiber.Ctx) error {
	itemID := c.Query("itemId")
	if itemID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Item ID is required")
	}

	id, err := strconv.Atoi(itemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Item ID")
	}
	// Set the HTTP Content-Type header.
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)

	// Buffer to store the rendered HTML.
	var output bytes.Buffer
	p := reg.Count(fmt.Sprintf("%dx", s.db.GetCount(int64(id), -1)))

	// Render the page into the buffer.
	if err := p.Render(context.Background(), &output); err != nil {
		// Handle the rendering error, e.g., by logging and returning an HTTP 500 error.
		log.Printf("Error rendering page: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Send the rendered HTML as the response.
	return c.SendStream(&output)
}