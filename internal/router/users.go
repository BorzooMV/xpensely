package router

import (
	"database/sql"

	"github.com/BorzooMV/xpensely/internal/handlers"
	"github.com/labstack/echo/v4"
)

func UsersRouter(e *echo.Echo, db *sql.DB) {
	handler := handlers.UserHandler{Db: db}
	e.GET(Paths["users"], handler.GetAllUsers)
	e.POST(Paths["users"], handler.CreateUser)
	e.GET(Paths["userWithId"], handler.GetSingleUser)
	e.DELETE(Paths["userWithId"], handler.DeleteUser)
}
