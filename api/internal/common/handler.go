package common

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func RegisterRoutes(g *echo.Group) {
	g.GET("/health", func(c *echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
}
