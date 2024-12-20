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

type UserHandler struct {
	Db *sql.DB
}

func (u *UserHandler) GetAllUsers(c echo.Context) error {
	getAllUsersQs := `
	SELECT
	id, firstname, lastname, username, email, created_at, updated_at, is_admin
	FROM users;
	`
	rows, err := u.Db.Query(getAllUsersQs)
	if err != nil {
		return utils.SendErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to query database: %v", err.Error()))
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.IsAdmin)
		if err != nil {
			return utils.SendErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to scan data: %v", err.Error()))
		}
		users = append(users, u)
	}

	return c.JSON(http.StatusOK, users)
}

func (u *UserHandler) GetSingleUser(c echo.Context) error {
	getSingleUserQs := `
	SELECT 
	id, firstname, lastname, username, email, created_at, updated_at, is_admin
	FROM users
	WHERE id = $1;
	`
	id := c.Param("id")
	row := u.Db.QueryRow(getSingleUserQs, id)
	user := models.User{}
	err := row.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin)
	if err != nil {
		return utils.SendErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to scan data: %v", err.Error()))
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	var Res struct {
		Id string `json:"id"`
	}
	var Req struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	err := json.NewDecoder(c.Request().Body).Decode(&Req)
	if err != nil {
		return utils.SendErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Can't decode request body:%v", err.Error()))
	}

	createUserQs := `
	INSERT INTO users
	(firstname, lastname, username, email, password)
	VALUES
	($1,$2,$3,$4,$5)
	RETURNING id;
	`
	row := u.Db.QueryRow(createUserQs, Req.Firstname, Req.Lastname, Req.Username, Req.Email, Req.Password)

	err = row.Scan(&Res.Id)
	if err != nil {
		return utils.SendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Can't scan retrieved data from database:%v", err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"user_id": Res.Id,
	})
}
