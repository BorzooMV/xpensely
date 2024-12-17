package router

import (
	"github.com/labstack/echo/v4"
)

func HandleRoutes(e *echo.Echo) {
	UsersRouter(e)
	// ExpensesRouter(e)
}
