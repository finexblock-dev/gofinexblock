package structs

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
)

type (
	SearchApiLogInput struct {
		Limit     int              `json:"limit,required" query:"limit,required" default:"20"`
		Offset    int              `json:"offset,required" query:"offset,required" default:"0"`
		AdminID   uint             `json:"admin_id,required" query:"admin_id,required"`
		StartTime string           `json:"start_time,required" query:"start_time,required"`
		EndTime   string           `json:"end_time,required" query:"end_time,required"`
		Method    entity.ApiMethod `json:"method,required" query:"method,required" binding:"required,enum"`
		Endpoint  string           `json:"endpoint,required" query:"endpoint,required"`
	}

	SearchApiLogOutput struct {
		Result []*entity.AdminApiLog
	}

	SearchApiLogSuccessResponse struct {
		Code int                `json:"code,required"`
		Data SearchApiLogOutput `json:"data,required"`
	}
)
