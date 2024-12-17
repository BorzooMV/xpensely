package router

import (
	"github.com/BorzooMV/xpensely/internal/handlers"
	"github.com/labstack/echo/v4"
)

func UsersRouter(e *echo.Echo) {
	e.GET(Paths["home"], handlers.HandleUsers)
	// e.POST
	// ...
}
