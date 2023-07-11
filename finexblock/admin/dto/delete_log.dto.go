package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	FindAllDeleteLogInput struct {
		Limit  int `json:"limit" query:"limit" default:"20"`
		Offset int `json:"offset" query:"offset" default:"0"`
	}

	FindAllDeleteLogOutput struct {
		Result []*admin.AdminDeleteLog `json:"result"`
	}

	FindAllDeleteLogSuccessResponse struct {
		Code int                    `json:"code"`
		Data FindAllDeleteLogOutput `json:"data"`
	}
)

type (
	SearchDeleteLogInput struct {
		Limit     int    `json:"limit" query:"limit" default:"20"`
		Offset    int    `json:"offset" query:"offset" default:"0"`
		StartTime string `json:"start_time" query:"start_time"`
		EndTime   string `json:"end_time" query:"end_time"`
		Executor  uint   `json:"executor" query:"executor"`
		Target    uint   `json:"target" query:"target"`
	}

	SearchDeleteLogOutput struct {
		Result []*admin.AdminDeleteLog `json:"result"`
	}

	SearchDeleteLogSuccessResponse struct {
		Code int                   `json:"code"`
		Data SearchDeleteLogOutput `json:"data"`
	}
)

type (
	FindDeleteLogOfExecutorInput struct {
		AdminID uint `json:"admin_id" query:"admin_id"`
		Limit   int  `json:"limit" query:"limit" default:"20"`
		Offset  int  `json:"offset" query:"offset" default:"0"`
	}

	FindDeleteLogOfExecutorOutput struct {
		Result []*admin.AdminDeleteLog `json:"result"`
	}

	FindDeleteLogOfExecutorSuccessResponse struct {
		Code int                           `json:"code"`
		Data FindDeleteLogOfExecutorOutput `json:"data"`
	}
)

type (
	FindDeleteLogOfTargetInput struct {
		AdminID uint `json:"adminID" query:"admin_id"`
	}

	FindDeleteLogOfTargetOutput struct {
		Result *admin.AdminDeleteLog `json:"result"`
	}

	FindDeleteLogOfTargetSuccessResponse struct {
		Code int                         `json:"code"`
		Data FindDeleteLogOfTargetOutput `json:"data"`
	}
)