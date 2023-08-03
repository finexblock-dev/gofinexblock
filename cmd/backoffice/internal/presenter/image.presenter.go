package presenter

import "github.com/gofiber/fiber/v2"

func ImageErrResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status":  false,
		"message": err.Error(),
	}
}

func ImageSuccessResponse(code int, data interface{}) *fiber.Map {
	return &fiber.Map{
		"code":   code,
		"status": true,
		"data":   data,
	}
}

func ImageMsgResponse(code int, message string) *fiber.Map {
	return &fiber.Map{
		"code":    code,
		"status":  true,
		"message": message,
	}
}
