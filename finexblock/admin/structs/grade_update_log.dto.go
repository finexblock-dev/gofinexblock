package structs

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
)

type (
	SearchGradeUpdateLogInput struct {
		Executor  uint   `json:"executor,required" query:"executor,required"`
		Target    uint   `json:"target,required" query:"target,required"`
		StartTime string `json:"start_time,required" query:"start_time,required"`
		EndTime   string `json:"end_time,required" query:"end_time,required"`
		Limit     int    `json:"limit,required" query:"limit,required" default:"20"`
		Offset    int    `json:"offset,required" query:"offset,required" default:"0"`
	}

	SearchGradeUpdateLogOutput struct {
		Result []*entity.AdminGradeUpdateLog `json:"result,required"`
	}

	SearchGradeUpdateLogSuccessResponse struct {
		Code int                        `json:"code,required"`
		Data SearchGradeUpdateLogOutput `json:"data,required"`
	}
)
