package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	FindLoginHistoryOfAdminInput struct {
		AdminID uint `json:"admin_id" query:"admin_id"`
		Limit   int  `json:"limit" query:"limit" default:"20"`
		Offset  int  `json:"offset" query:"offset" default:"0"`
	}

	FindLoginHistoryOfAdminOutput struct {
		Result []*admin.AdminLoginHistory `json:"result"`
	}

	FindLoginHistoryOfAdminSuccessResponse struct {
		Code int                           `json:"code"`
		Data FindLoginHistoryOfAdminOutput `json:"data"`
	}
)