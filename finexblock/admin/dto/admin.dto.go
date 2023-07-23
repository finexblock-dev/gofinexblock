package dto

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
)

type (
	FindAllAdminInput struct {
		Limit  int `json:"limit,required" query:"limit,required" default:"20"`
		Offset int `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindAllAdminOutput struct {
		Result []*entity.PartialAdmin
	}

	FindAllAdminSuccessResponse struct {
		Code int                `json:"code,required"`
		Data FindAllAdminOutput `json:"data,required"`
	}
)

type (
	FindAdminByGradeInput struct {
		Limit  int              `json:"limit,required" query:"limit,required" default:"20"`
		Offset int              `json:"offset,required" query:"offset,required" default:"0"`
		Grade  entity.GradeType `json:"grade,required" query:"grade,required" binding:"required,enum"`
	}

	FindAdminByGradeOutput struct {
		Result []*entity.Admin `json:"result,required"`
	}

	FindAdminByGradeSwaggerOutput struct {
		Result []*entity.PartialAdmin `json:"result,required"`
	}

	FindAdminByGradeSuccessResponse struct {
		Code int                           `json:"code,required"`
		Data FindAdminByGradeSwaggerOutput `json:"data,required"`
	}
)

type (
	DeleteAdminInput struct {
		AdminID uint `json:"admin_id,required"`
	}

	DeleteAdminOutput struct {
		Msg string `json:"msg,required"`
	}

	DeleteAdminSuccessResponse struct {
		Code int               `json:"code,required"`
		Data DeleteAdminOutput `json:"data,required"`
	}
)

type (
	UpdatePasswordInput struct {
		AdminID      uint   `json:"admin_id,required"`
		PrevPassword string `json:"prev_password,required"`
		NewPassword  string `json:"new_password,required"`
	}

	UpdatePasswordOutput struct {
		Msg string `json:"msg,required"`
	}

	UpdatePasswordSuccessResponse struct {
		Code int                  `json:"code,required"`
		Data UpdatePasswordOutput `json:"data,required"`
	}
)

type (
	UpdateEmailInput struct {
		AdminID uint   `json:"admin_id,required"`
		Email   string `json:"email,required"`
	}

	UpdateEmailOutput struct {
		Msg string `json:"msg,required"`
	}

	UpdateEmailSuccessResponse struct {
		Code int               `json:"code,required"`
		Data UpdateEmailOutput `json:"data,required"`
	}
)

type (
	UpdateGradeInput struct {
		AdminID uint             `json:"admin_id,required" query:"adminId,required" binding:"required"`
		Grade   entity.GradeType `json:"grade,required" query:"grade,required" binding:"required,enum"`
	}

	UpdateGradeOutput struct {
		Msg string `json:"msg,required"`
	}

	UpdateGradeSuccessResponse struct {
		Code int               `json:"code,required"`
		Data UpdateGradeOutput `json:"data,required"`
	}
)

type (
	BlockAdminInput struct {
		AdminID uint `json:"admin_id,required" query:"admin_id,required" default:"1"`
	}

	BlockAdminOutput struct {
		Msg string `json:"msg,required" default:"Successfully blocked"`
	}

	BlockAdminSuccessResponse struct {
		Code int              `json:"code,required"`
		Data BlockAdminOutput `json:"data,required"`
	}
)

type (
	UnblockAdminInput struct {
		AdminID uint `json:"admin_id,required" query:"admin_id,required" default:"1"`
	}

	UnblockAdminOutput struct {
		Msg string `json:"msg,required" default:"Successfully unblocked"`
	}

	UnblockAdminSuccessResponse struct {
		Code int                `json:"code,required"`
		Data UnblockAdminOutput `json:"data,required"`
	}
)
