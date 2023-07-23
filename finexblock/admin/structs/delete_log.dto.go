package structs

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
)

type (
	SearchDeleteLogInput struct {
		Limit     int    `json:"limit,required" query:"limit,required" default:"20"`
		Offset    int    `json:"offset,required" query:"offset,required" default:"0"`
		StartTime string `json:"start_time,required" query:"start_time,required"`
		EndTime   string `json:"end_time,required" query:"end_time,required"`
		Executor  uint   `json:"executor,required" query:"executor,required"`
		Target    uint   `json:"target,required" query:"target,required"`
	}

	SearchDeleteLogOutput struct {
		Result []*entity.AdminDeleteLog `json:"result,required"`
	}

	SearchDeleteLogSuccessResponse struct {
		Code int                   `json:"code,required"`
		Data SearchDeleteLogOutput `json:"data,required"`
	}
)
