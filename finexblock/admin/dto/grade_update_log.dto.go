package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	FindAllGradeUpdateLogInput struct {
		Limit  int `json:"limit" query:"limit" default:"20"`
		Offset int `json:"offset" query:"offset" default:"0"`
	}

	FindAllGradeUpdateLogOutput struct {
		Result []*admin.AdminGradeUpdateLog `json:"result"`
	}

	FindAllGradeUpdateLogSuccessResponse struct {
		Code int                         `json:"code"`
		Data FindAllGradeUpdateLogOutput `json:"data"`
	}
)

type (
	FindGradeUpdateLogOfExecutorInput struct {
		Executor uint `json:"executor" query:"executor"`
		Limit    int  `json:"limit" query:"limit" default:"20"`
		Offset   int  `json:"offset" query:"offset" default:"0"`
	}

	FindGradeUpdateLogOfExecutorOutput struct {
		Result []*admin.AdminGradeUpdateLog `json:"result"`
	}

	FindGradeUpdateLogOfExecutorSuccessResponse struct {
		Code int                                `json:"code"`
		Data FindGradeUpdateLogOfExecutorOutput `json:"data"`
	}
)

type (
	FindGradeUpdateLogOfTargetInput struct {
		Target uint `json:"target" query:"target"`
		Limit  int  `json:"limit" query:"limit" default:"20"`
		Offset int  `json:"offset" query:"offset" default:"0"`
	}

	FindGradeUpdateLogOfTargetOutput struct {
		Result []*admin.AdminGradeUpdateLog `json:"result"`
	}

	FindGradeUpdateLogOfTargetSuccessResponse struct {
		Code int                              `json:"code"`
		Data FindGradeUpdateLogOfTargetOutput `json:"data"`
	}
)

type (
	SearchGradeUpdateLogInput struct {
		Executor  uint   `json:"executor" query:"executor"`
		Target    uint   `json:"target" query:"target"`
		StartTime string `json:"start_time" query:"start_time"`
		EndTime   string `json:"end_time" query:"end_time"`
		Limit     int    `json:"limit" query:"limit" default:"20"`
		Offset    int    `json:"offset" query:"offset" default:"0"`
	}

	SearchGradeUpdateLogOutput struct {
		Result []*admin.AdminGradeUpdateLog `json:"result"`
	}

	SearchGradeUpdateLogSuccessResponse struct {
		Code int                        `json:"code"`
		Data SearchGradeUpdateLogOutput `json:"data"`
	}
)