package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	SearchApiLogInput struct {
		Limit     int    `json:"limit" query:"limit" default:"20"`
		Offset    int    `json:"offset" query:"offset" default:"0"`
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
		Code int                `json:"code"`
		Data SearchApiLogOutput `json:"data"`
	}
)

type (
	FindAllApiLogInput struct {
		Limit  int `json:"limit" query:"limit" default:"20"`
		Offset int `json:"offset" query:"offset" default:"0"`
	}

	FindAllApiLogOutput struct {
		Result []*admin.AdminApiLog
	}

	FindAllApiLogSuccessResponse struct {
		Code int                 `json:"code"`
		Data FindAllApiLogOutput `json:"data"`
	}
)

type (
	FindApiLogByAdminInput struct {
		AdminID uint `json:"admin_id" query:"admin_id"`
		Limit   int  `json:"limit" query:"limit" default:"20"`
		Offset  int  `json:"offset" query:"offset" default:"0"`
	}

	FindApiLogByAdminOutput struct {
		Result []*admin.AdminApiLog
	}

	FindApiLogByAdminSuccessResponse struct {
		Code int                     `json:"code"`
		Data FindApiLogByAdminOutput `json:"data"`
	}
)

type (
	FindApiLogByTimeCondInput struct {
		StartTime string `json:"start_time" query:"start_time"`
		EndTime   string `json:"end_time" query:"end_time"`
		Limit     int    `json:"limit" query:"limit" default:"20"`
		Offset    int    `json:"offset" query:"offset" default:"0"`
	}

	FindApiLogByTimeCondOutput struct {
		Result []*admin.AdminApiLog
	}

	FindApiLogByTimeCondSuccessResponse struct {
		Code int                        `json:"code"`
		Data FindApiLogByTimeCondOutput `json:"data"`
	}
)

type (
	FindApiLogByMethodCondInput struct {
		Method string `json:"method" query:"method"`
		Limit  int    `json:"limit" query:"limit" default:"20"`
		Offset int    `json:"offset" query:"offset" default:"0"`
	}

	FindApiLogByMethodCondOutput struct {
		Result []*admin.AdminApiLog
	}

	FindApiLogByMethodCondSuccessResponse struct {
		Code int                          `json:"code"`
		Data FindApiLogByMethodCondOutput `json:"data"`
	}
)

type (
	FindApiLogByEndpointInput struct {
		Endpoint string `json:"endpoint" query:"endpoint"`
		Limit    int    `json:"limit" query:"limit" default:"20"`
		Offset   int    `json:"offset" query:"offset" default:"0"`
	}

	FindApiLogByEndpointOutput struct {
		Result []*admin.AdminApiLog `json:"result"`
	}

	FindApiLogByEndpointSuccessResponse struct {
		Code int                        `json:"code"`
		Data FindApiLogByEndpointOutput `json:"data"`
	}
)