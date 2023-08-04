package presenter

import "github.com/gofiber/fiber/v2"

func AssetErrResponse(code int, err error) *fiber.Error {
	return &fiber.Error{
		Code:    code,
		Message: err.Error(),
	}
}

func AssetSuccessResponse(code int, data interface{}) *fiber.Map {
	return &fiber.Map{
		"code":   code,
		"status": true,
		"data":   data,
	}
}
func AssetMsgResponse(code int, message string) *fiber.Map {
	return &fiber.Map{
		"code":    code,
		"status":  true,
		"message": message,
	}
}