package parser

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/pkg/order/structs"
)

func SearchOrderMatchingHistoryInput(input *dto.SearchTradeHistoryInput) (*structs.SearchOrderMatchingHistoryInput, error) {
	return &structs.SearchOrderMatchingHistoryInput{
		UserID:        input.UserID,
		OrderSymbolID: input.OrderSymbolID,
		StartTime:     input.StartTime,
		EndTime:       input.EndTime,
		Limit:         input.Limit,
		Offset:        input.Offset,
	}, nil
}