package main

import (
	"github.com/M-Haruki/fsledger/api/internal/common"
	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()
	api := e.Group("/api")
	common.RegisterRoutes(api)
	e.Start(":1323")
}
