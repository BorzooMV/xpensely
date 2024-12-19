package utils

import "github.com/labstack/echo/v4"

func SendErrorResponse(c echo.Context, s int, t string) error {
	return c.JSON(s, map[string]string{
		"error": t,
	})
}
