package structs

import (
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
)

type (
	SearchPasswordUpdateLogInput struct {
		Limit     int    `json:"limit,required" query:"limit,required" default:"20"`
		Offset    int    `json:"offset,required" query:"offset,required" default:"0"`
		StartTime string `json:"startTime,required" query:"startTime,required"`
		EndTime   string `json:"endTime,required" query:"endTime,required"`
		Executor  uint   `json:"executor,required" query:"executor,required"`
		Target    uint   `json:"target,required" query:"target,required"`
	}

	SearchPasswordUpdateLogOutput struct {
		Result []*entity.AdminPasswordLog `json:"result,required"`
	}

	SearchPasswordUpdateLogSuccessResponse struct {
		Code int                           `json:"code,required"`
		Data SearchPasswordUpdateLogOutput `json:"data,required"`
	}
)