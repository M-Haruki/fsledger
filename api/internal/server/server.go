package server

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/M-Haruki/fsledger/api/internal/common"
	"github.com/M-Haruki/fsledger/api/internal/db"
	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/M-Haruki/fsledger/api/internal/openapi"
	"github.com/M-Haruki/fsledger/api/internal/stock"
	"github.com/M-Haruki/fsledger/api/internal/swagger"
	"github.com/labstack/echo/v5"
	echomiddleware "github.com/oapi-codegen/echo-v5-middleware"
)

type Server struct {
	echo *echo.Echo
}

func New(ctx context.Context, cfg Config) (*Server, error) {
	// logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: cfg.LogLevel}))

	// db
	db, err := db.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	queries := sqlc.New(db)

	// common
	commonRepo := common.NewRepository(queries)
	commonService := common.NewService(*commonRepo)
	commonHandler := common.NewHandler(commonService, logger)
	// stock
	stockRepo := stock.NewRepository(queries)
	stockService := stock.NewService(*stockRepo)
	stockHandler := stock.NewHandler(stockService, logger)

	strictServer := newStrictServer(commonHandler, stockHandler)
	strictHandler := openapi.NewStrictHandler(strictServer, nil)

	// echo
	e := echo.New()

	// openapi
	api := e.Group("/api")
	validateSwagger, err := openapi.GetSpec()
	if err != nil {
		log.Fatal(err)
	}
	api.Use(echomiddleware.OapiRequestValidator(validateSwagger))
	openapi.RegisterHandlers(api, strictHandler)

	// swagger ui
	if cfg.IsDev {
		swagger.RegisterHandlers(e)
	}

	return &Server{
		echo: e,
	}, nil
}

func (s *Server) Start(addr string) error {
	return s.echo.Start(addr)
}
