package handler

import "github.com/gofiber/fiber/v2"

type TradeAPI interface {
	SearchTradeHistory() fiber.Handler
}

type TradeHandler struct {
}
