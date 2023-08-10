package handler

import (
	"github.com/finexblock-dev/gofinexblock/pkg/wallet"
	"github.com/gofiber/fiber/v2"
)

type TradeAPI interface {
	SearchTradeHistory() fiber.Handler
}

type TradeHandler struct {
	walletService wallet.Service
}
