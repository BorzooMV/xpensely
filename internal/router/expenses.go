package router

import (
	"database/sql"

	"github.com/BorzooMV/xpensely/internal/handlers"
	"github.com/labstack/echo/v4"
)

func ExpensesRouter(e *echo.Echo, db *sql.DB) {
	handler := handlers.ExpensesHandler{
		Db: db,
	}

	e.GET(Paths["expensesOfSingleUser"], handler.GetAllExpensesOfSingleUser)
}
