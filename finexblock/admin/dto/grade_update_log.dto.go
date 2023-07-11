package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	FindAllGradeUpdateLogInput struct {
		Limit  int `json:"limit,required" query:"limit,required" default:"20"`
		Offset int `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindAllGradeUpdateLogOutput struct {
		Result []*admin.AdminGradeUpdateLog `json:"result,required"`
	}

	FindAllGradeUpdateLogSuccessResponse struct {
		Code int                         `json:"code,required"`
		Data FindAllGradeUpdateLogOutput `json:"data,required"`
	}
)

type (
	FindGradeUpdateLogOfExecutorInput struct {
		Executor uint `json:"executor,required" query:"executor,required"`
		Limit    int  `json:"limit,required" query:"limit,required" default:"20"`
		Offset   int  `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindGradeUpdateLogOfExecutorOutput struct {
		Result []*admin.AdminGradeUpdateLog `json:"result,required"`
	}

	FindGradeUpdateLogOfExecutorSuccessResponse struct {
		Code int                                `json:"code,required"`
		Data FindGradeUpdateLogOfExecutorOutput `json:"data,required"`
	}
)

type (
	FindGradeUpdateLogOfTargetInput struct {
		Target uint `json:"target,required" query:"target,required"`
		Limit  int  `json:"limit,required" query:"limit,required" default:"20"`
		Offset int  `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindGradeUpdateLogOfTargetOutput struct {
		Result []*admin.AdminGradeUpdateLog `json:"result,required"`
	}

	FindGradeUpdateLogOfTargetSuccessResponse struct {
		Code int                              `json:"code,required"`
		Data FindGradeUpdateLogOfTargetOutput `json:"data,required"`
	}
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
		Result []*admin.AdminGradeUpdateLog `json:"result,required"`
	}

	SearchGradeUpdateLogSuccessResponse struct {
		Code int                        `json:"code,required"`
		Data SearchGradeUpdateLogOutput `json:"data,required"`
	}
)