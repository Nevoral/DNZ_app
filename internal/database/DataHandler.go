package database

import (
	"context"
	"fmt"
)

func (s *surrealClient) CreateUser(ctx context.Context, arg CreateUserParam) (row *CreateUserResult, err error) {
	var result []*CreateUserResult
	err = s.Query(ctx, createUser, arg, &result)
	if err != nil {
		return nil, err
	}
	if len(result) != 1 {
		return nil, fmt.Errorf("expected 1 row, got %d", len(result))
	}
	return result[0], nil
}

//	func (s *service) Health() map[string]string {
//		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
//		defer cancel()
//
//		err := s.db.PingContext(ctx)
//		if err != nil {
//			log.Fatalf(fmt.Sprintf("db down: %v", err))
//		}
//
//		return map[string]string{
//			"message": "It's healthy",
//		}
//	}
//
//	func (s *service) InitDBTables() error {
//		statements := strings.Split(ddl, ";")
//		for _, statement := range statements {
//			statement = strings.TrimSpace(statement)
//			if statement == "" {
//				continue
//			}
//
//			if _, err := s.db.ExecContext(s.ctx, statement); err != nil {
//				if !strings.Contains(statement, "CREATE INDEX") {
//					zlog.ErrorLog(fmt.Sprintf("executing statement %q: %s", statement, err))
//					return fmt.Errorf("executing statement %q: %s", statement, err)
//				}
//			}
//		}
//		return nil
//	}
//func (s *service) InitDBQueries() {
//	s.queries = buf.New(s.db)
//}

//
//func (s *service) CreateProductMenu(start int64) (*buf.ProductMenu, error) {
//	menu, err := s.queries.CreateProductMenuAndReturnIt(s.ctx, buf.CreateProductMenuAndReturnItParams{
//		Date:          time.Now(),
//		Startregister: start,
//	})
//	if err != nil {
//		zlog.ErrorLog(err.Error())
//		return nil, err
//
//	}
//	return &menu, nil
//}
//
//func (s *service) GetListOfProductMenu() ([]buf.ProductMenu, error) {
//	menu, err := s.queries.GetListProductMenu(s.ctx)
//	if err != nil {
//		return nil, err
//	}
//	return menu, nil
//}
//
//func (s *service) GetCount(id, quant int64) int64 {
//	product, err := s.queries.GetProductByID(s.ctx, id)
//	if err != nil {
//
//	}
//	count := product.Served
//	if count+quant > -1 {
//		count += quant
//	}
//	err = s.queries.SetProductCounterByID(s.ctx, buf.SetProductCounterByIDParams{
//		Served: count,
//		ID:     product.ID,
//	})
//	if err != nil {
//		fmt.Println(err)
//		return 0
//	}
//	return count
//}
//
//func (s *service) SetCount(title string, quant int64) {
//	product, err := s.queries.GetProductByTitle(s.ctx, title)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	count := product.Served
//	if count+quant > -1 {
//		count += quant
//	}
//	err = s.queries.SetProductCounterByID(s.ctx, buf.SetProductCounterByIDParams{
//		Served: count,
//		ID:     product.ID,
//	})
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	return
//}
//
//func (s *service) CreateProduct(title, category string, price, idMenu int64) error {
//	_, err := s.queries.CreateProductAndReturnIt(s.ctx, buf.CreateProductAndReturnItParams{
//		Productmenuid: idMenu,
//		Title:         title,
//		Price:         price,
//		Served:        0,
//		Category:      category,
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (s *service) GetListOfProducts(id int64, category string) ([]buf.Product, error) {
//	products, err := s.queries.GetProductListByProductMenuIDAndCategory(s.ctx, buf.GetProductListByProductMenuIDAndCategoryParams{
//		Productmenuid: id,
//		Category:      category,
//	})
//	if err != nil {
//		return nil, err
//	}
//	return products, nil
//}
//
//func (s *service) CreateOrderWithoutUser(productMenuID, total int64, items string) error {
//	_, err := s.queries.CreateOpenOrderAndReturnIt(s.ctx, buf.CreateOpenOrderAndReturnItParams{
//		Productmenuid: productMenuID,
//		Customerid: sql.NullInt64{
//			Int64: 0,
//			Valid: false,
//		},
//		Date:         time.Now(),
//		Summary:      total,
//		Status:       "close",
//		Itemsordered: items,
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (s *service) CreateOrderWithUser(productMenuID, customerID, total int64, items string) error {
//	_, err := s.queries.CreateOpenOrderAndReturnIt(s.ctx, buf.CreateOpenOrderAndReturnItParams{
//		Productmenuid: productMenuID,
//		Customerid: sql.NullInt64{
//			Int64: customerID,
//			Valid: true,
//		},
//		Date:         time.Now(),
//		Summary:      total,
//		Status:       "open",
//		Itemsordered: items,
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//
//}
//
//func (s *service) CreateUser(name, phoneNumber, role string) (*buf.Customer, error) {
//	customer, err := s.queries.CreateCustomerAndReturnIt(s.ctx, buf.CreateCustomerAndReturnItParams{
//		Name: name,
//		Phonenumber: sql.NullString{
//			String: phoneNumber,
//			Valid:  true,
//		},
//		Role: role,
//	})
//	if err != nil {
//		return nil, err
//	}
//	return &customer, nil
//}
//
//func (s *service) UpdateOrderWithUser() {
//
//}
