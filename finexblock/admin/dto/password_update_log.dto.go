package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	SearchPasswordUpdateLogInput struct {
		Limit     int    `json:"limit" query:"limit" default:"20"`
		Offset    int    `json:"offset" query:"offset" default:"0"`
		StartTime string `json:"start_time" query:"start_time"`
		EndTime   string `json:"end_time" query:"end_time"`
		Executor  uint   `json:"executor" query:"executor"`
		Target    uint   `json:"target" query:"target"`
	}

	SearchPasswordUpdateLogOutput struct {
		Result []*admin.AdminPasswordLog `json:"result"`
	}

	SearchPasswordUpdateLogSuccessResponse struct {
		Code int                           `json:"code"`
		Data SearchPasswordUpdateLogOutput `json:"data"`
	}
)

type (
	FindAllPasswordUpdateLogInput struct {
		Limit  int `json:"limit" query:"limit" default:"20"`
		Offset int `json:"offset" query:"offset" default:"0"`
	}

	FindAllPasswordUpdateLogOutput struct {
		Result []*admin.AdminPasswordLog `json:"result"`
	}

	FindAllPasswordUpdateLogSuccessResponse struct {
		Code int                            `json:"code"`
		Data FindAllPasswordUpdateLogOutput `json:"data"`
	}
)

type (
	FindPasswordUpdateLogOfExecutorInput struct {
		AdminID uint `json:"admin_id" query:"admin_id"`
		Limit   int  `json:"limit" query:"limit" default:"20"`
		Offset  int  `json:"offset" query:"offset" default:"0"`
	}

	FindPasswordUpdateLogOfExecutorOutput struct {
		Result []*admin.AdminPasswordLog `json:"result"`
	}

	FindPasswordUpdateLogOfExecutorSuccessResponse struct {
		Code int                                   `json:"code"`
		Data FindPasswordUpdateLogOfExecutorOutput `json:"data"`
	}
)

type (
	FindPasswordUpdateLogOfTargetInput struct {
		AdminID uint `json:"admin_id" query:"admin_id"`
		Limit   int  `json:"limit" query:"limit" default:"20"`
		Offset  int  `json:"offset" query:"offset" default:"0"`
	}

	FindPasswordUpdateLogOfTargetOutput struct {
		Result []*admin.AdminPasswordLog `json:"result"`
	}

	FindPasswordUpdateLogOfTargetSuccessResponse struct {
		Code int                                 `json:"code"`
		Data FindPasswordUpdateLogOfTargetOutput `json:"data"`
	}
)