package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	FindAllDeleteLogInput struct {
		Limit  int `json:"limit" query:"limit"`
		Offset int `json:"offset" query:"offset"`
	}

	FindAllDeleteLogOutput struct {
		Result []*admin.AdminDeleteLog `json:"result,omitempty"`
	}

	FindAllDeleteLogSuccessResponse struct {
		Code int                    `json:"code,omitempty"`
		Data FindAllDeleteLogOutput `json:"data,omitempty"`
	}
)

type (
	SearchDeleteLogInput struct {
		Limit     int    `json:"limit" query:"limit"`
		Offset    int    `json:"offset" query:"offset"`
		StartTime string `json:"start_time" query:"start_time"`
		EndTime   string `json:"end_time" query:"end_time"`
		Executor  uint   `json:"executor" query:"executor"`
		Target    uint   `json:"target" query:"target"`
	}

	SearchDeleteLogOutput struct {
		Result []*admin.AdminDeleteLog `json:"result,omitempty"`
	}

	SearchDeleteLogSuccessResponse struct {
		Code int                   `json:"code,omitempty"`
		Data SearchDeleteLogOutput `json:"data,omitempty"`
	}
)

type (
	FindDeleteLogOfExecutorInput struct {
		AdminID uint `json:"adminID" query:"admin_id"`
		Limit   int  `json:"limit" query:"limit"`
		Offset  int  `json:"offset" query:"offset"`
	}

	FindDeleteLogOfExecutorOutput struct {
		Result []*admin.AdminDeleteLog `json:"result,omitempty"`
	}

	FindDeleteLogOfExecutorSuccessResponse struct {
		Code int                           `json:"code,omitempty"`
		Data FindDeleteLogOfExecutorOutput `json:"data,omitempty"`
	}
)

type (
	FindDeleteLogOfTargetInput struct {
		AdminID uint `json:"adminID" query:"admin_id"`
	}

	FindDeleteLogOfTargetOutput struct {
		Result *admin.AdminDeleteLog `json:"result,omitempty"`
	}

	FindDeleteLogOfTargetSuccessResponse struct {
		Code int                         `json:"code,omitempty"`
		Data FindDeleteLogOfTargetOutput `json:"data,omitempty"`
	}
)
