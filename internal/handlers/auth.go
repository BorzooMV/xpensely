package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BorzooMV/xpensely/services"
	"github.com/BorzooMV/xpensely/utils"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Db *sql.DB
}

func (h *AuthHandler) AuthenticateUser(e echo.Context) error {
	getUserByUsernameQs := `
	SELECT username, password FROM users WHERE username = $1;
	`
	var StoredUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var Req struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := json.NewDecoder(e.Request().Body).Decode(&Req); err != nil {
		return utils.SendErrorResponse(e, http.StatusInternalServerError, fmt.Sprintf("couldn't decode the request body: %v", err.Error()))
	}

	if err := e.Validate(&Req); err != nil {
		return utils.SendErrorResponse(e, http.StatusBadRequest, fmt.Sprintf("A required field haven't received: %v", err.Error()))
	}

	row := h.Db.QueryRow(getUserByUsernameQs, Req.Username)
	if err := row.Scan(
		&StoredUser.Username,
		&StoredUser.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return utils.SendErrorResponse(e, http.StatusNotFound, "User not found")
		}

		return utils.SendErrorResponse(e, http.StatusInternalServerError, fmt.Sprintf("Couldn't scan fetched body from database: %v", err.Error()))
	}

	if err := utils.ValidatePassword(StoredUser.Password, Req.Password); err != nil {
		return utils.SendErrorResponse(e, http.StatusBadRequest, "Password is incorrect")
	}

	accessToken, err := utils.CreateAccessToken(StoredUser.Username)
	if err != nil {
		return utils.SendErrorResponse(e, http.StatusInternalServerError, fmt.Sprintf("Couldn't create token: %v", err.Error()))
	}

	refreshToken, err := utils.CreateRefreshToken(StoredUser.Username)
	if err != nil {
		return utils.SendErrorResponse(e, http.StatusInternalServerError, fmt.Sprintf("Couldn't create token: %v", err.Error()))
	}

	return e.JSON(http.StatusOK, map[string]string{
		"access_token":  fmt.Sprintf("Bearer %v", accessToken),
		"refresh_token": fmt.Sprintf("Bearer %v", refreshToken),
	})
}

func (h *AuthHandler) RefreshAccessToken(e echo.Context) error {
	var Req struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}
	ctx := context.Background()
	redisClient := services.ConnectRedis()
	defer redisClient.Close()

	if err := json.NewDecoder(e.Request().Body).Decode(&Req); err != nil {
		return utils.SendErrorResponse(e, http.StatusInternalServerError, fmt.Sprintf("Couldn't decode request body: %v", err.Error()))
	}

	Req.RefreshToken = Req.RefreshToken[len("Bearer "):]

	username, err := redisClient.Get(ctx, Req.RefreshToken).Result()
	if err != nil {
		return utils.SendErrorResponse(e, http.StatusNotFound, "Not found")
	}

	_, err = utils.ValidateToken(Req.RefreshToken)
	if err != nil {
		return utils.SendErrorResponse(e, http.StatusBadRequest, fmt.Sprintf("couldn't validate token: %v", err.Error()))
	}

	accessToken, err := utils.CreateAccessToken(username)
	if err != nil {
		return utils.SendErrorResponse(e, http.StatusInternalServerError, fmt.Sprintf("Couldn't create token: %v", err.Error()))
	}
	newRefreshToken, err := utils.CreateRefreshToken(username)
	if err != nil {
		return utils.SendErrorResponse(e, http.StatusInternalServerError, fmt.Sprintf("Couldn't create token: %v", err.Error()))
	}

	return e.JSON(http.StatusOK, map[string]string{
		"access_token":  fmt.Sprintf("Bearer %v", accessToken),
		"refresh_token": fmt.Sprintf("Bearer %v", newRefreshToken),
	})

}
