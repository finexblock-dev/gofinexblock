package dto

type (
	FindAllAdminInput struct {
		Limit  int `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindAdminByGradeInput struct {
		Limit  int    `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int    `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
		Grade  string `json:"grade,required" query:"grade,required" example:"SUPERUSER, MAINTAINER, SUPPORT" binding:"required"`
	}
)

type (
	DeleteAdminInput struct {
		AdminID uint `json:"adminId,required" query:"adminId,required" binding:"required"`
	}
)

type (
	UpdatePasswordInput struct {
		AdminID      uint   `json:"adminId,required" binding:"required"`
		PrevPassword string `json:"prevPassword,required" binding:"required"`
		NewPassword  string `json:"newPassword,required" binding:"required"`
	}
)

type (
	UpdateEmailInput struct {
		AdminID uint   `json:"adminId,required" binding:"required"`
		Email   string `json:"email,required" binding:"required"`
	}
)

type (
	UpdateGradeInput struct {
		AdminID uint   `json:"adminId,required" binding:"required" default:"3"`
		Grade   string `json:"grade,required" default:"M" example:"M" binding:"required"`
	}
)

type (
	BlockAdminInput struct {
		AdminID uint `json:"adminId,required" query:"adminId,required" default:"1"`
	}
)

type (
	UnblockAdminInput struct {
		AdminID uint `json:"adminId,required" query:"adminId,required" default:"1"`
	}
)

type (
	SearchPasswordUpdateLogInput struct {
		Limit     int    `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset    int    `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
		StartTime string `json:"startTime" query:"startTime"`
		EndTime   string `json:"endTime" query:"endTime"`
		Executor  uint   `json:"executor" query:"executor"`
		Target    uint   `json:"target" query:"target"`
	}
)

type (
	FindAllPasswordUpdateLogInput struct {
		Limit  int `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindPasswordUpdateLogOfExecutorInput struct {
		AdminID uint `json:"adminId,required" query:"adminId,required" binding:"required"`
		Limit   int  `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset  int  `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindPasswordUpdateLogOfTargetInput struct {
		AdminID uint `json:"adminId,required" query:"adminId,required" binding:"required"`
		Limit   int  `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset  int  `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindLoginHistoryOfAdminInput struct {
		AdminID uint `json:"adminId,required" query:"adminId,required" binding:"required" default:"1"`
		Limit   int  `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset  int  `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindLoginFailedLogInput struct {
		AdminID uint `json:"adminId,required" query:"adminId,required" default:"1" binding:"required"`
		Limit   int  `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset  int  `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindAllGradeUpdateLogInput struct {
		Limit  int `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindGradeUpdateLogOfExecutorInput struct {
		Executor uint `json:"executor,required" query:"executor,required" binding:"required"`
		Limit    int  `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset   int  `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindGradeUpdateLogOfTargetInput struct {
		Target uint `json:"target,required" query:"target,required" binding:"required"`
		Limit  int  `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int  `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	SearchGradeUpdateLogInput struct {
		Executor  uint   `json:"executor" query:"executor"`
		Target    uint   `json:"target" query:"target"`
		StartTime string `json:"startTime" query:"startTime"`
		EndTime   string `json:"endTime" query:"endTime"`
		Limit     int    `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset    int    `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	SearchDeleteLogInput struct {
		Limit     int    `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset    int    `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
		StartTime string `json:"startTime" query:"startTime"`
		EndTime   string `json:"endTime" query:"endTime"`
		Executor  uint   `json:"executor" query:"executor"`
		Target    uint   `json:"target" query:"target"`
	}
)

type (
	FindDeleteLogOfExecutorInput struct {
		AdminID uint `json:"adminId,required" query:"adminId,required" binding:"required"`
		Limit   int  `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset  int  `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindDeleteLogOfTargetInput struct {
		AdminID uint `json:"adminId,required" query:"adminId,required" binding:"required"`
	}
)

type (
	FindAllDeleteLogInput struct {
		Limit  int `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	SearchApiLogInput struct {
		Limit     int    `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset    int    `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
		AdminID   uint   `json:"adminId" query:"adminId"`
		StartTime string `json:"startTime" query:"startTime"`
		EndTime   string `json:"endTime" query:"endTime"`
		Method    string `json:"method" query:"method" example:"GET, POST, PATCH, PUT, DELETE..."`
		Endpoint  string `json:"endpoint" query:"endpoint"`
	}
)

type (
	FindAllApiLogInput struct {
		Limit  int `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindApiLogByAdminInput struct {
		AdminID uint `json:"adminId,required" query:"adminId,required" binding:"required"`
		Limit   int  `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset  int  `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindApiLogByTimeCondInput struct {
		StartTime string `json:"startTime,required" query:"startTime,required" binding:"required"`
		EndTime   string `json:"endTime,required" query:"endTime,required" binding:"required"`
		Limit     int    `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset    int    `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindApiLogByMethodCondInput struct {
		Method string `json:"method,required" query:"method,required" example:"GET, POST, PATCH, PUT, DELETE..." binding:"required"`
		Limit  int    `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int    `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindApiLogByEndpointInput struct {
		Endpoint string `json:"endpoint,required" query:"endpoint,required" binding:"required"`
		Limit    int    `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset   int    `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindOnlineAdminInput struct {
		Limit  int `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)