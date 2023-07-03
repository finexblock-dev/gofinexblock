package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	SearchApiLogInput struct {
		Limit     int    `json:"limit" query:"limit"`
		Offset    int    `json:"offset" query:"offset"`
		AdminID   uint   `json:"admin_id" query:"admin_id"`
		StartTime string `json:"start_time" query:"start_time"`
		EndTime   string `json:"end_time" query:"end_time"`
		Method    string `json:"method" query:"method"`
		Endpoint  string `json:"endpoint" query:"endpoint"`
	}

	SearchApiLogOutput struct {
		Result []*admin.AdminApiLog
	}

	SearchApiLogSuccessResponse struct {
		Code int                `json:"code,omitempty"`
		Data SearchApiLogOutput `json:"data,omitempty"`
	}
)

type (
	FindAllApiLogInput struct {
		Limit  int `json:"limit" query:"limit"`
		Offset int `json:"offset" query:"offset"`
	}

	FindAllApiLogOutput struct {
		Result []*admin.AdminApiLog
	}

	FindAllApiLogSuccessResponse struct {
		Code int                 `json:"code,omitempty"`
		Data FindAllApiLogOutput `json:"data,omitempty"`
	}
)

type (
	FindApiLogByAdminInput struct {
		AdminID uint `json:"admin_id" query:"admin_id"`
		Limit   int  `json:"limit" query:"limit"`
		Offset  int  `json:"offset" query:"offset"`
	}

	FindApiLogByAdminOutput struct {
		Result []*admin.AdminApiLog
	}

	FindApiLogByAdminSuccessResponse struct {
		Code int                     `json:"code,omitempty"`
		Data FindApiLogByAdminOutput `json:"data,omitempty"`
	}
)

type (
	FindApiLogByTimeCondInput struct {
		StartTime string `json:"start_time" query:"start_time"`
		EndTime   string `json:"end_time" query:"end_time"`
		Limit     int    `json:"limit" query:"limit"`
		Offset    int    `json:"offset" query:"offset"`
	}

	FindApiLogByTimeCondOutput struct {
		Result []*admin.AdminApiLog
	}

	FindApiLogByTimeCondSuccessResponse struct {
		Code int                        `json:"code,omitempty"`
		Data FindApiLogByTimeCondOutput `json:"data,omitempty"`
	}
)

type (
	FindApiLogByMethodCondInput struct {
		Method string `json:"method" query:"method"`
		Limit  int    `json:"limit" query:"limit"`
		Offset int    `json:"offset" query:"offset"`
	}

	FindApiLogByMethodCondOutput struct {
		Result []*admin.AdminApiLog
	}

	FindApiLogByMethodCondSuccessResponse struct {
		Code int                          `json:"code,omitempty"`
		Data FindApiLogByMethodCondOutput `json:"data,omitempty"`
	}
)

type (
	FindApiLogByEndpointInput struct {
		Endpoint string `json:"endpoint" query:"endpoint"`
		Limit    int    `json:"limit" query:"limit"`
		Offset   int    `json:"offset" query:"offset"`
	}

	FindApiLogByEndpointOutput struct {
		Result []*admin.AdminApiLog `json:"result,omitempty"`
	}

	FindApiLogByEndpointSuccessResponse struct {
		Code int                        `json:"code,omitempty"`
		Data FindApiLogByEndpointOutput `json:"data,omitempty"`
	}
)
