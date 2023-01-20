package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/vandyahmad24/moonlay-test/app/route"
	"github.com/vandyahmad24/moonlay-test/config"
	"github.com/vandyahmad24/moonlay-test/config/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	configData := config.GetEnvVariable("DB_DRIVER")
	database.InitDbMysql()

	if configData == "PSQL" {
		database.InitDbPsql()
	}
	database.InitMigration()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Moonlay Test Code By Vandy Ahmad")
	})
	e.Validator = &CustomValidator{validator: validator.New()}

	route.RouteTodo(e)

	portServer := config.GetEnvVariable("PORT_SERVER")
	e.Logger.Fatal(e.Start(":" + portServer))
}
