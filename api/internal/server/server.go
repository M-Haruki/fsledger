package server

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
)

type Server struct {
	echo *echo.Echo
	db   *pgxpool.Pool
}

func New(ctx context.Context, cfg Config) (*Server, error) {
	// logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: cfg.LogLevel}))

	// db
	db, queries, err := newDB(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// handler
	handler := newOpenAPIHandler(db, queries, logger)

	// echo
	e, err := newEcho(cfg, handler)
	if err != nil {
		db.Close()
		return nil, err
	}

	return &Server{
		echo: e,
		db:   db,
	}, nil
}

func (s *Server) Start(ctx context.Context, addr string) error {
	sc := echo.StartConfig{
		Address:         addr,
		GracefulTimeout: 5 * time.Second,
	}
	return sc.Start(ctx, s.echo)
}

func (s *Server) Close() {
	s.db.Close()
}
