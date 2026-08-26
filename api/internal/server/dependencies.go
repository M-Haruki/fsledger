package server

import (
	"log/slog"

	"github.com/M-Haruki/fsledger/api/internal/common"
	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/M-Haruki/fsledger/api/internal/flow"
	"github.com/M-Haruki/fsledger/api/internal/openapi"
	"github.com/M-Haruki/fsledger/api/internal/stock"
	"github.com/M-Haruki/fsledger/api/internal/transaction"
	"github.com/jackc/pgx/v5/pgxpool"
)

type strictServer struct {
	*common.CommonHandler
	*stock.StockHandler
	*transaction.TransactionHandler
	*flow.FlowHandler
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
	// transaction
	transactionRepo := transaction.NewRepository(db, queries)
	transactionService := transaction.NewService(*transactionRepo)
	transactionHandler := transaction.NewHandler(transactionService, logger)
	// flow
	flowRepo := flow.NewRepository(db, queries)
	flowService := flow.NewService(*flowRepo)
	flowHandler := flow.NewHandler(flowService, logger)

	server := strictServer{
		CommonHandler:      commonHandler,
		StockHandler:       stockHandler,
		TransactionHandler: transactionHandler,
		FlowHandler:        flowHandler,
	}

	return openapi.NewStrictHandler(server, nil)
}
