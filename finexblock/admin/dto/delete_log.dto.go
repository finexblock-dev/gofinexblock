package dto

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
)

type (
	FindAllDeleteLogInput struct {
		Limit  int `json:"limit,required" query:"limit,required" default:"20"`
		Offset int `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindAllDeleteLogOutput struct {
		Result []*entity.AdminDeleteLog `json:"result,required"`
	}

	FindAllDeleteLogSuccessResponse struct {
		Code int                    `json:"code,required"`
		Data FindAllDeleteLogOutput `json:"data,required"`
	}
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

type (
	FindDeleteLogOfExecutorInput struct {
		AdminID uint `json:"admin_id,required" query:"admin_id,required"`
		Limit   int  `json:"limit,required" query:"limit,required" default:"20"`
		Offset  int  `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindDeleteLogOfExecutorOutput struct {
		Result []*entity.AdminDeleteLog `json:"result,required"`
	}

	FindDeleteLogOfExecutorSuccessResponse struct {
		Code int                           `json:"code,required"`
		Data FindDeleteLogOfExecutorOutput `json:"data,required"`
	}
)

type (
	FindDeleteLogOfTargetInput struct {
		AdminID uint `json:"adminID,required" query:"admin_id,required"`
	}

	FindDeleteLogOfTargetOutput struct {
		Result *entity.AdminDeleteLog `json:"result,required"`
	}

	FindDeleteLogOfTargetSuccessResponse struct {
		Code int                         `json:"code,required"`
		Data FindDeleteLogOfTargetOutput `json:"data,required"`
	}
)
