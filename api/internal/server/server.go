package server

import (
	"context"
	"log/slog"
	"os"

	"github.com/M-Haruki/fsledger/api/internal/common"
	"github.com/M-Haruki/fsledger/api/internal/db"
	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/M-Haruki/fsledger/api/internal/openapi"
	"github.com/M-Haruki/fsledger/api/internal/stock"
	"github.com/M-Haruki/fsledger/api/internal/swagger"
	"github.com/labstack/echo/v5"
)

type Server struct {
	echo *echo.Echo
}

func New(ctx context.Context, cfg Config) (*Server, error) {
	// logger
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: cfg.LogLevel}))

	// db
	db, err := db.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	queries := sqlc.New(db)

	// common
	commonRepo := common.NewRepository(queries)
	commonService := common.NewService(*commonRepo)
	commonHandler := common.NewHandler(commonService, log)
	// stock
	stockRepo := stock.NewRepository(queries)
	stockService := stock.NewService(*stockRepo)
	stockHandler := stock.NewHandler(stockService, log)

	strictServer := newStrictServer(commonHandler, stockHandler)
	strictHandler := openapi.NewStrictHandler(strictServer, nil)

	e := echo.New()
	openapi.RegisterHandlersWithBaseURL(e, strictHandler, "/api")

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
