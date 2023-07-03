package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	SearchPasswordUpdateLogInput struct {
		Limit     int    `json:"limit" query:"limit"`
		Offset    int    `json:"offset" query:"offset"`
		StartTime string `json:"start_time" query:"start_time"`
		EndTime   string `json:"end_time" query:"end_time"`
		Executor  uint   `json:"executor" query:"executor"`
		Target    uint   `json:"target" query:"target"`
	}

	SearchPasswordUpdateLogOutput struct {
		Result []*admin.AdminPasswordLog `json:"result,omitempty"`
	}

	SearchPasswordUpdateLogSuccessResponse struct {
		Code int                           `json:"code,omitempty"`
		Data SearchPasswordUpdateLogOutput `json:"data,omitempty"`
	}
)

type (
	FindAllPasswordUpdateLogInput struct {
		Limit  int `json:"limit" query:"limit"`
		Offset int `json:"offset" query:"offset"`
	}

	FindAllPasswordUpdateLogOutput struct {
		Result []*admin.AdminPasswordLog `json:"result,omitempty"`
	}

	FindAllPasswordUpdateLogSuccessResponse struct {
		Code int                            `json:"code,omitempty"`
		Data FindAllPasswordUpdateLogOutput `json:"data,omitempty"`
	}
)

type (
	FindPasswordUpdateLogOfExecutorInput struct {
		AdminID uint `json:"adminID" query:"admin_id"`
		Limit   int  `json:"limit" query:"limit"`
		Offset  int  `json:"offset" query:"offset"`
	}

	FindPasswordUpdateLogOfExecutorOutput struct {
		Result []*admin.AdminPasswordLog `json:"result,omitempty"`
	}

	FindPasswordUpdateLogOfExecutorSuccessResponse struct {
		Code int                                   `json:"code,omitempty"`
		Data FindPasswordUpdateLogOfExecutorOutput `json:"data,omitempty"`
	}
)

type (
	FindPasswordUpdateLogOfTargetInput struct {
		AdminID uint `json:"adminID" query:"admin_id"`
		Limit   int  `json:"limit" query:"limit"`
		Offset  int  `json:"offset" query:"offset"`
	}

	FindPasswordUpdateLogOfTargetOutput struct {
		Result []*admin.AdminPasswordLog `json:"result,omitempty"`
	}

	FindPasswordUpdateLogOfTargetSuccessResponse struct {
		Code int                                 `json:"code,omitempty"`
		Data FindPasswordUpdateLogOfTargetOutput `json:"data,omitempty"`
	}
)
