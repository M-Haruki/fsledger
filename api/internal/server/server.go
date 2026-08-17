package server

import (
	"context"
	"os"

	"github.com/M-Haruki/fsledger/api/internal/common"
	"github.com/M-Haruki/fsledger/api/internal/db"
	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/M-Haruki/fsledger/api/internal/openapi"
	"github.com/labstack/echo/v5"
)

type Server struct {
	echo *echo.Echo
}

func New(ctx context.Context) (*Server, error) {

	db, err := db.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	queries := sqlc.New(db)

	commonRepo := common.NewRepository(queries)
	commonService := common.NewService(*commonRepo)
	commonHandler := common.NewHandler(commonService)

	e := echo.New()

	strictHandler := openapi.NewStrictHandler(commonHandler, nil)

	openapi.RegisterHandlersWithBaseURL(e, strictHandler, "/api")

	return &Server{
		echo: e,
	}, nil
}

func (s *Server) Start(addr string) error {
	return s.echo.Start(addr)
}
