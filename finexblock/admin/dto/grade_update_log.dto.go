package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	FindAllGradeUpdateLogInput struct {
		Limit  int `json:"limit" query:"limit"`
		Offset int `json:"offset" query:"offset"`
	}

	FindAllGradeUpdateLogOutput struct {
		Result []*admin.AdminGradeUpdateLog `json:"result,omitempty"`
	}

	FindAllGradeUpdateLogSuccessResponse struct {
		Code int                         `json:"code,omitempty"`
		Data FindAllGradeUpdateLogOutput `json:"data,omitempty"`
	}
)

type (
	FindGradeUpdateLogOfExecutorInput struct {
		Executor uint `json:"executor" query:"executor"`
		Limit    int  `json:"limit" query:"limit"`
		Offset   int  `json:"offset" query:"offset"`
	}

	FindGradeUpdateLogOfExecutorOutput struct {
		Result []*admin.AdminGradeUpdateLog `json:"result,omitempty"`
	}

	FindGradeUpdateLogOfExecutorSuccessResponse struct {
		Code int                                `json:"code,omitempty"`
		Data FindGradeUpdateLogOfExecutorOutput `json:"data,omitempty"`
	}
)

type (
	FindGradeUpdateLogOfTargetInput struct {
		Target uint `json:"target" query:"target"`
		Limit  int  `json:"limit" query:"limit"`
		Offset int  `json:"offset" query:"offset"`
	}

	FindGradeUpdateLogOfTargetOutput struct {
		Result []*admin.AdminGradeUpdateLog `json:"result,omitempty"`
	}

	FindGradeUpdateLogOfTargetSuccessResponse struct {
		Code int                              `json:"code,omitempty"`
		Data FindGradeUpdateLogOfTargetOutput `json:"data,omitempty"`
	}
)

type (
	SearchGradeUpdateLogInput struct {
		Executor  uint   `json:"executor" query:"executor"`
		Target    uint   `json:"target" query:"target"`
		StartTime string `json:"start_time" query:"start_time"`
		EndTime   string `json:"end_time" query:"end_time"`
		Limit     int    `json:"limit" query:"limit"`
		Offset    int    `json:"offset" query:"offset"`
	}

	SearchGradeUpdateLogOutput struct {
		Result []*admin.AdminGradeUpdateLog `json:"result,omitempty"`
	}

	SearchGradeUpdateLogSuccessResponse struct {
		Code int                        `json:"code,omitempty"`
		Data SearchGradeUpdateLogOutput `json:"data,omitempty"`
	}
)
