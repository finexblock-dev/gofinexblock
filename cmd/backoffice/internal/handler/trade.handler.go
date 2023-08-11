package handler

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/parser"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/presenter"
	"github.com/finexblock-dev/gofinexblock/pkg/order"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet"
	"github.com/gofiber/fiber/v2"
)

type TradeAPI interface {
	SearchTradeHistory() fiber.Handler
}

type TradeHandler struct {
	walletService wallet.Service
	orderService  order.Service
}

func NewTradeHandler(walletService wallet.Service, orderService order.Service) *TradeHandler {
	return &TradeHandler{walletService: walletService, orderService: orderService}
}

// SearchTradeHistory @SearchTradeHistory
// @description	Search trade history
// @security		BearerAuth
// @tags			Trade
// @accept			json
// @produce		json
// @param			SearchTradeHistoryInput	query		dto.SearchTradeHistoryInput	true	"SearchTradeHistoryInput"
// @success		200							{object}	[]entity.OrderMatchingHistory			"Success"
// @failure		400							{object}	presenter.MsgResponse	"Failed"
// @router			/trade/search [get]
func (t *TradeHandler) SearchTradeHistory() fiber.Handler {
	return func(c *fiber.Ctx) error {
		
		query := new(dto.SearchTradeHistoryInput)
		
		if err := c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.NewErrResponse(err))
		}
		
		input, err := parser.SearchOrderMatchingHistoryInput(query)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.NewErrResponse(err))
		}
		
		result, err := t.orderService.SearchOrderMatchingHistory(input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.NewErrResponse(err))
		}
		
		return c.JSON(result)
	}
}