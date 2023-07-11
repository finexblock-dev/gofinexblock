package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	FindLoginHistoryOfAdminInput struct {
		AdminID uint `json:"admin_id,required" query:"admin_id,required"`
		Limit   int  `json:"limit,required" query:"limit,required" default:"20"`
		Offset  int  `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindLoginHistoryOfAdminOutput struct {
		Result []*admin.AdminLoginHistory `json:"result,required"`
	}

	FindLoginHistoryOfAdminSuccessResponse struct {
		Code int                           `json:"code,required"`
		Data FindLoginHistoryOfAdminOutput `json:"data,required"`
	}
)