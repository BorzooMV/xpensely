package internal

import (
	"database/sql"

	"github.com/BorzooMV/xpensely/internal/router"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.Validator.Struct(i)
}

func CreateEchoServer(db *sql.DB) *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	validate := validator.New(validator.WithRequiredStructEnabled())
	e.Validator = &CustomValidator{Validator: validate}
	router.HandleRoutes(e, db)
	return e
}
