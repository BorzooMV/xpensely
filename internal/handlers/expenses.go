package handlers

import (
	"database/sql"
	"encoding/json"
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
	var userExists bool

	isUserExistsQs := `
		SELECT EXISTS(
			SELECT 1 FROM users WHERE id = $1
		);
	`
	row := h.Db.QueryRow(isUserExistsQs, e.Param("id"))
	err := row.Scan(&userExists)
	if err != nil {
		return utils.SendErrorResponse(e, http.StatusInternalServerError, fmt.Sprintf("Couldn't query database: %v", err.Error()))
	}
	if !userExists {
		return utils.SendErrorResponse(e, http.StatusNotFound, "User not found")
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

func (h *ExpensesHandler) CreateNewExpense(e echo.Context) error {
	var Request struct {
		UserId      int     `json:"user_id" validate:"required,number"`
		Amount      float64 `json:"amount" validate:"required,number,min=0"`
		Description string  `json:"description" validate:"required,max=255"`
	}

	var Response struct {
		ExpenseId string `json:"id"`
	}

	createNewExpenseQs := `
	INSERT INTO expenses
	(user_id, amount, description)
	VALUES
	($1, $2, $3)
	RETURNING id;
	`

	err := json.NewDecoder(e.Request().Body).Decode(&Request)
	if err != nil {
		return utils.SendErrorResponse(e, http.StatusInternalServerError, fmt.Sprintf("Couldn't decode requested body: %v", err.Error()))
	}

	if err := e.Validate(&Request); err != nil {
		return utils.SendErrorResponse(e, http.StatusBadRequest, err.Error())
	}

	row := h.Db.QueryRow(createNewExpenseQs, Request.UserId, Request.Amount, Request.Description)

	err = row.Scan(&Response.ExpenseId)
	if err != nil {
		return utils.SendErrorResponse(e, http.StatusInternalServerError, fmt.Sprintf("couldn't scan database response: %v", err.Error()))
	}

	return e.JSON(http.StatusOK, map[string]string{
		"expense_id": Response.ExpenseId,
	})

}
