package dto

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
)

type (
	SearchApiLogInput struct {
		Limit     int    `json:"limit,required" query:"limit,required" default:"20"`
		Offset    int    `json:"offset,required" query:"offset,required" default:"0"`
		AdminID   uint   `json:"admin_id,required" query:"admin_id,required"`
		StartTime string `json:"start_time,required" query:"start_time,required"`
		EndTime   string `json:"end_time,required" query:"end_time,required"`
		Method    string `json:"method,required" query:"method,required"`
		Endpoint  string `json:"endpoint,required" query:"endpoint,required"`
	}

	SearchApiLogOutput struct {
		Result []*entity.AdminApiLog
	}

	SearchApiLogSuccessResponse struct {
		Code int                `json:"code,required"`
		Data SearchApiLogOutput `json:"data,required"`
	}
)

type (
	FindAllApiLogInput struct {
		Limit  int `json:"limit,required" query:"limit,required" default:"20"`
		Offset int `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindAllApiLogOutput struct {
		Result []*entity.AdminApiLog
	}

	FindAllApiLogSuccessResponse struct {
		Code int                 `json:"code,required"`
		Data FindAllApiLogOutput `json:"data,required"`
	}
)

type (
	FindApiLogByAdminInput struct {
		AdminID uint `json:"admin_id,required" query:"admin_id,required"`
		Limit   int  `json:"limit,required" query:"limit,required" default:"20"`
		Offset  int  `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindApiLogByAdminOutput struct {
		Result []*entity.AdminApiLog
	}

	FindApiLogByAdminSuccessResponse struct {
		Code int                     `json:"code,required"`
		Data FindApiLogByAdminOutput `json:"data,required"`
	}
)

type (
	FindApiLogByTimeCondInput struct {
		StartTime string `json:"start_time,required" query:"start_time,required"`
		EndTime   string `json:"end_time,required" query:"end_time,required"`
		Limit     int    `json:"limit,required" query:"limit,required" default:"20"`
		Offset    int    `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindApiLogByTimeCondOutput struct {
		Result []*entity.AdminApiLog
	}

	FindApiLogByTimeCondSuccessResponse struct {
		Code int                        `json:"code,required"`
		Data FindApiLogByTimeCondOutput `json:"data,required"`
	}
)

type (
	FindApiLogByMethodCondInput struct {
		Method string `json:"method,required" query:"method,required"`
		Limit  int    `json:"limit,required" query:"limit,required" default:"20"`
		Offset int    `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindApiLogByMethodCondOutput struct {
		Result []*entity.AdminApiLog
	}

	FindApiLogByMethodCondSuccessResponse struct {
		Code int                          `json:"code,required"`
		Data FindApiLogByMethodCondOutput `json:"data,required"`
	}
)

type (
	FindApiLogByEndpointInput struct {
		Endpoint string `json:"endpoint,required" query:"endpoint,required"`
		Limit    int    `json:"limit,required" query:"limit,required" default:"20"`
		Offset   int    `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindApiLogByEndpointOutput struct {
		Result []*entity.AdminApiLog `json:"result,required"`
	}

	FindApiLogByEndpointSuccessResponse struct {
		Code int                        `json:"code,required"`
		Data FindApiLogByEndpointOutput `json:"data,required"`
	}
)
