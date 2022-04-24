package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/zweimach/xendit-test/api"
	"github.com/zweimach/xendit-test/db"
	"github.com/zweimach/xendit-test/utils"
)

func main() {
	driver := os.Getenv("DB_DRIVER")
	source := os.Getenv("DB_SOURCE")
	if driver == "" || source == "" {
		log.Fatal("please provide DB_SOURCE and DB_DRIVER env variable")
	}

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${status} ${uri} ${error}\n",
	}))
	e.Use(middleware.Recover())
	e.Validator = utils.NewCustomValidator(validator.New())

	conn, err := sql.Open(driver, source)
	if err != nil {
		log.Fatal(err.Error())
	}

	db := db.New(conn)

	h := api.NewHandler(context.Background(), db)
	h.Register(e)

	if err := e.Start(":1234"); err != nil {
		e.Logger.Fatal(err)
	}
}
