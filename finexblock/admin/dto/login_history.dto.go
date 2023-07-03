package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	FindLoginHistoryOfAdminInput struct {
		AdminID uint `json:"adminID" query:"admin_id"`
		Limit   int  `json:"limit" query:"limit"`
		Offset  int  `json:"offset" query:"offset"`
	}

	FindLoginHistoryOfAdminOutput struct {
		Result []*admin.AdminLoginHistory `json:"result,omitempty"`
	}

	FindLoginHistoryOfAdminSuccessResponse struct {
		Code int                           `json:"code,omitempty"`
		Data FindLoginHistoryOfAdminOutput `json:"data,omitempty"`
	}
)
