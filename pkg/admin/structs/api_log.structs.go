package structs

import (
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
)

type (
	SearchApiLogInput struct {
		Limit     int              `json:"limit,required" query:"limit,required" default:"20"`
		Offset    int              `json:"offset,required" query:"offset,required" default:"0"`
		AdminID   uint             `json:"adminId,required" query:"adminId,required"`
		StartTime string           `json:"startTime,required" query:"startTime,required"`
		EndTime   string           `json:"endTime,required" query:"endTime,required"`
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