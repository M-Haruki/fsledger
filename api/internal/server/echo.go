package server

import (
	"github.com/M-Haruki/fsledger/api/internal/openapi"
	"github.com/M-Haruki/fsledger/api/internal/swagger"
	"github.com/labstack/echo/v5"
	echomiddleware "github.com/oapi-codegen/echo-v5-middleware"
)

func newEcho(cfg Config, handler openapi.ServerInterface) (*echo.Echo, error) {
	e := echo.New()

	// openapi
	api := e.Group("/api")
	validateSwagger, err := openapi.GetSpec()
	if err != nil {
		return nil, err
	}
	api.Use(echomiddleware.OapiRequestValidator(validateSwagger))
	openapi.RegisterHandlers(api, handler)

	// swagger ui
	if cfg.IsDev {
		swagger.RegisterHandlers(e)
	}
	return e, nil
}
