package server

import (
	"github.com/M-Haruki/fsledger/api/internal/common"
	"github.com/M-Haruki/fsledger/api/internal/stock"
)

type strictServer struct {
	*common.CommonHandler
	*stock.StockHandler
}

func newStrictServer(
	commonH *common.CommonHandler,
	stockH *stock.StockHandler,
) *strictServer {
	return &strictServer{
		CommonHandler: commonH,
		StockHandler:  stockH,
	}
}
