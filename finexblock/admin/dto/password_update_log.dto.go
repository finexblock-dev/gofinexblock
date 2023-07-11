package dto

import "github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"

type (
	SearchPasswordUpdateLogInput struct {
		Limit     int    `json:"limit,required" query:"limit,required" default:"20"`
		Offset    int    `json:"offset,required" query:"offset,required" default:"0"`
		StartTime string `json:"start_time,required" query:"start_time,required"`
		EndTime   string `json:"end_time,required" query:"end_time,required"`
		Executor  uint   `json:"executor,required" query:"executor,required"`
		Target    uint   `json:"target,required" query:"target,required"`
	}

	SearchPasswordUpdateLogOutput struct {
		Result []*admin.AdminPasswordLog `json:"result,required"`
	}

	SearchPasswordUpdateLogSuccessResponse struct {
		Code int                           `json:"code,required"`
		Data SearchPasswordUpdateLogOutput `json:"data,required"`
	}
)

type (
	FindAllPasswordUpdateLogInput struct {
		Limit  int `json:"limit,required" query:"limit,required" default:"20"`
		Offset int `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindAllPasswordUpdateLogOutput struct {
		Result []*admin.AdminPasswordLog `json:"result,required"`
	}

	FindAllPasswordUpdateLogSuccessResponse struct {
		Code int                            `json:"code,required"`
		Data FindAllPasswordUpdateLogOutput `json:"data,required"`
	}
)

type (
	FindPasswordUpdateLogOfExecutorInput struct {
		AdminID uint `json:"admin_id,required" query:"admin_id,required"`
		Limit   int  `json:"limit,required" query:"limit,required" default:"20"`
		Offset  int  `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindPasswordUpdateLogOfExecutorOutput struct {
		Result []*admin.AdminPasswordLog `json:"result,required"`
	}

	FindPasswordUpdateLogOfExecutorSuccessResponse struct {
		Code int                                   `json:"code,required"`
		Data FindPasswordUpdateLogOfExecutorOutput `json:"data,required"`
	}
)

type (
	FindPasswordUpdateLogOfTargetInput struct {
		AdminID uint `json:"admin_id,required" query:"admin_id,required"`
		Limit   int  `json:"limit,required" query:"limit,required" default:"20"`
		Offset  int  `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindPasswordUpdateLogOfTargetOutput struct {
		Result []*admin.AdminPasswordLog `json:"result,required"`
	}

	FindPasswordUpdateLogOfTargetSuccessResponse struct {
		Code int                                 `json:"code,required"`
		Data FindPasswordUpdateLogOfTargetOutput `json:"data,required"`
	}
)