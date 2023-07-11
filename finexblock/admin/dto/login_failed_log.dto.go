package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	FindLoginFailedLogInput struct {
		AdminID uint `json:"admin_id,required" query:"admin_id,required"`
		Limit   int  `json:"limit,required" query:"limit,required" default:"20"`
		Offset  int  `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindLoginFailedLogOutput struct {
		Result []*admin.AdminLoginFailedLog `json:"result,required"`
	}

	FindLoginFailedLogSuccessResponse struct {
		Code int                      `json:"code,required"`
		Data FindLoginFailedLogOutput `json:"data,required"`
	}
)