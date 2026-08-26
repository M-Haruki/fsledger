package server

import (
	"log/slog"

	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/M-Haruki/fsledger/api/internal/handler"
	"github.com/M-Haruki/fsledger/api/internal/openapi"
	"github.com/M-Haruki/fsledger/api/internal/repository"
	"github.com/M-Haruki/fsledger/api/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

func newOpenAPIHandler(db *pgxpool.Pool, queries *sqlc.Queries, logger *slog.Logger) openapi.ServerInterface {
	repository := repository.NewRepository(db, queries)
	service := service.NewService(*repository)
	handler := handler.NewHandler(service, logger)
	return openapi.NewStrictHandler(handler, nil)
}
