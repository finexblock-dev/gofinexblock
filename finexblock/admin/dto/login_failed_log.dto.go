package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	FindLoginFailedLogInput struct {
		AdminID uint `json:"admin_id" query:"admin_id"`
		Limit   int  `json:"limit" query:"limit"`
		Offset  int  `json:"offset" query:"offset"`
	}

	FindLoginFailedLogOutput struct {
		Result []*admin.AdminLoginFailedLog `json:"result,omitempty"`
	}

	FindLoginFailedLogSuccessResponse struct {
		Code int                      `json:"code,omitempty"`
		Data FindLoginFailedLogOutput `json:"data,omitempty"`
	}
)
