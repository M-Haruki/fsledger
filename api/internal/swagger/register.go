package swagger

import (
	"github.com/labstack/echo/v5"
)

func RegisterHandlers(router *echo.Echo) {
	router.StaticFS("/swagger", FS)
}
