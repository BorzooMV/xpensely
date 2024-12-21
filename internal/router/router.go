package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func HandleRoutes(e *echo.Echo, db *sql.DB) {
	UsersRouter(e, db)
	ExpensesRouter(e, db)
}
