package router

import (
	"database/sql"

	"github.com/BorzooMV/xpensely/internal/handlers"
	"github.com/labstack/echo/v4"
)

func Auth(e *echo.Echo, db *sql.DB) {
	handler := handlers.AuthHandler{
		Db: db,
	}

	e.POST(Paths["auth"], handler.AuthenticateUser)
}
