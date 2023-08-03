package presenter

import "github.com/gofiber/fiber/v2"

func AuthErrResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status":  false,
		"message": err.Error(),
	}
}

func AuthSuccessResponse(code int, data interface{}) *fiber.Map {
	return &fiber.Map{
		"code":   code,
		"status": true,
		"data":   data,
	}
}

func AuthMsgResponse(code int, message string) *fiber.Map {
	return &fiber.Map{
		"code":    code,
		"status":  true,
		"message": message,
	}
}