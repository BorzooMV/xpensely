package middlewares

import (
	"net/http"

	"github.com/BorzooMV/xpensely/internal/router"
	"github.com/BorzooMV/xpensely/utils"
	"github.com/labstack/echo/v4"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().RequestURI == router.Paths["auth"] {
			return next(c)
		}

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		}

		reqToken := authHeader[len("Bearer "):]

		if _, err := utils.ValidateToken(reqToken); err != nil {
			return utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}
