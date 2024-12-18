package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/BorzooMV/xpensely/internal/models"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Db *sql.DB
}

func (u *UserHandler) GetAllUsers(c echo.Context) error {
	qs := "SELECT * FROM users;"
	rows, err := u.Db.Query(qs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to query database: %v", err.Error()),
		})
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.IsAdmin)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": fmt.Sprintf("Failed to scan data: %v", err.Error()),
			})
		}
		users = append(users, u)
	}

	return c.JSON(http.StatusOK, users)
}
