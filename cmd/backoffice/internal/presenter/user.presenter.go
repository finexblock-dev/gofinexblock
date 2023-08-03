package presenter

import (
	"github.com/gofiber/fiber/v2"
)

func UserErrResponse(code int, err error) *fiber.Error {
	return &fiber.Error{
		Code:    code,
		Message: err.Error(),
	}
}

func UserSuccessResponse(code int, data interface{}) *fiber.Map {
	return &fiber.Map{
		"code":   code,
		"status": true,
		"data":   data,
	}
}
func UserMsgResponse(code int, message string) *fiber.Map {
	return &fiber.Map{
		"code":    code,
		"status":  true,
		"message": message,
	}
}
