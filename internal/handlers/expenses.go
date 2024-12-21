package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/BorzooMV/xpensely/internal/models"
	"github.com/BorzooMV/xpensely/utils"
	"github.com/labstack/echo/v4"
)

type ExpensesHandler struct {
	Db *sql.DB
}

func (h *ExpensesHandler) GetAllExpensesOfSingleUser(e echo.Context) error {
	var Response struct {
		Count    int              `json:"count"`
		Expenses []models.Expense `json:"expenses"`
	}

	getExpensesOfSingleUserQs := `
	SELECT
	expenses.id, amount, description, expenses.created_at, expenses.updated_at
	FROM users INNER JOIN expenses
	ON users.id = expenses.user_id
	WHERE user_id = $1;
	`

	rows, err := h.Db.Query(getExpensesOfSingleUserQs, e.Param("id"))
	if err != nil {
		return utils.SendErrorResponse(e, http.StatusInternalServerError, fmt.Sprintf("couldn't query database: %v", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		singleExpense := models.Expense{}
		err := rows.Scan(&singleExpense.Id, &singleExpense.Amount, &singleExpense.Description, &singleExpense.CreatedAt, &singleExpense.UpdatedAt)
		if err != nil {
			return utils.SendErrorResponse(e, http.StatusInternalServerError, fmt.Sprintf("Couldn't scan fetched row from database: %v", err.Error()))
		}
		Response.Expenses = append(Response.Expenses, singleExpense)
	}

	Response.Count = len(Response.Expenses)
	return e.JSON(http.StatusOK, Response)
}
