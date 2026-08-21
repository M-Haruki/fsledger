package server

import (
	"log/slog"

	"github.com/M-Haruki/fsledger/api/internal/common"
	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/M-Haruki/fsledger/api/internal/openapi"
	"github.com/M-Haruki/fsledger/api/internal/stock"
	"github.com/M-Haruki/fsledger/api/internal/tag"
	"github.com/jackc/pgx/v5/pgxpool"
)

type strictServer struct {
	*common.CommonHandler
	*stock.StockHandler
	*tag.TagHandler
}

func newOpenAPIHandler(db *pgxpool.Pool, queries *sqlc.Queries, logger *slog.Logger) openapi.ServerInterface {
	// common
	commonRepo := common.NewRepository(queries)
	commonService := common.NewService(*commonRepo)
	commonHandler := common.NewHandler(commonService, logger)
	// stock
	stockRepo := stock.NewRepository(db, queries)
	stockService := stock.NewService(*stockRepo)
	stockHandler := stock.NewHandler(stockService, logger)
	// tag
	tagRepo := tag.NewRepository(db, queries)
	tagService := tag.NewService(*tagRepo)
	tagHandler := tag.NewHandler(tagService, logger)

	server := strictServer{
		CommonHandler: commonHandler,
		StockHandler:  stockHandler,
		TagHandler:    tagHandler,
	}

	return openapi.NewStrictHandler(server, nil)
}
