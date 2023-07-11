package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	FindLoginFailedLogInput struct {
		AdminID uint `json:"admin_id" query:"admin_id"`
		Limit   int  `json:"limit" query:"limit" default:"20"`
		Offset  int  `json:"offset" query:"offset" default:"0"`
	}

	FindLoginFailedLogOutput struct {
		Result []*admin.AdminLoginFailedLog `json:"result"`
	}

	FindLoginFailedLogSuccessResponse struct {
		Code int                      `json:"code"`
		Data FindLoginFailedLogOutput `json:"data"`
	}
)