package database

import (
	"context"
	"fmt"
	zlog "github.com/Nevoral/DNZ_app/internal/logging"
	_ "github.com/joho/godotenv/autoload"
	"github.com/surrealdb/surrealdb.go"
	"os"
)

type DbClient interface {
	CreateUser(ctx context.Context, arg CreateUserParam) (row *CreateUserResult, err error)
}

type surrealClient struct {
	client *surrealdb.DB
}

var (
	dbUrl       = os.Getenv("DB_URL")
	dbPort      = os.Getenv("DB_PORT")
	dbNamespace = os.Getenv("DB_NAMESPACE")
	dbName      = os.Getenv("DB_NAME")
	userName    = os.Getenv("DB_USER_NAME")
	userPass    = os.Getenv("DB_USER_PASS")
)

func New() DbClient {
	db, err := surrealdb.New(fmt.Sprintf("ws://%s%s/rpc", dbUrl, dbPort))
	if err != nil {
		zlog.PanicLog(err.Error())
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": userName,
		"pass": userPass,
	}); err != nil {
		zlog.PanicLog(err.Error())
	}

	if _, err = db.Use(dbNamespace, dbName); err != nil {
		zlog.PanicLog(err.Error())
	}

	return &surrealClient{
		client: db,
	}
}

func (s *surrealClient) DefineTablesDB() error {
	return nil
}

func (s *surrealClient) Query(ctx context.Context, query string, args any, returnVal any) error {
	resultChan := make(chan interface{})
	errChan := make(chan error)

	go func() {
		results, err := s.client.Query(query, args)
		if err != nil {
			errChan <- err
			return
		}
		resultChan <- results
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	case results := <-resultChan:
		ok, err := surrealdb.UnmarshalRaw(results, returnVal)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("Query returned no results")
		}
		return nil
	}
}
