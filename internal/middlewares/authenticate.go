package middlewares

import (
	"net/http"

	"github.com/BorzooMV/xpensely/internal/router"
	"github.com/BorzooMV/xpensely/utils"
	"github.com/labstack/echo/v4"
)

var unguardedRoutes = map[string]string{
	router.Paths["users"]: http.MethodPost,
	router.Paths["auth"]:  http.MethodPost,
}

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if method, ok := unguardedRoutes[c.Request().RequestURI]; ok && c.Request().Method == method {
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
