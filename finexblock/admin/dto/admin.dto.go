package dto

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/admin/types"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"
)

type (
	FindAllAdminInput struct {
		Limit  int `json:"limit" query:"limit" default:"20"`
		Offset int `json:"offset" query:"offset" default:"0"`
	}

	FindAllAdminOutput struct {
		Result []*types.PartialAdmin
	}

	FindAllAdminSuccessResponse struct {
		Code int                `json:"code"`
		Data FindAllAdminOutput `json:"data"`
	}
)

type (
	FindAdminByGradeInput struct {
		Limit  int             `json:"limit" query:"limit" default:"20"`
		Offset int             `json:"offset" query:"offset" default:"0"`
		Grade  admin.GradeType `json:"grade" query:"grade"`
	}

	FindAdminByGradeOutput struct {
		Result []*admin.Admin `json:"result"`
	}

	FindAdminByGradeSwaggerOutput struct {
		Result []*types.PartialAdmin `json:"result"`
	}

	FindAdminByGradeSuccessResponse struct {
		Code int                           `json:"code"`
		Data FindAdminByGradeSwaggerOutput `json:"data"`
	}
)

type (
	DeleteAdminInput struct {
		AdminID uint `json:"admin_id"`
	}

	DeleteAdminOutput struct {
		Msg string `json:"msg"`
	}

	DeleteAdminSuccessResponse struct {
		Code int               `json:"code"`
		Data DeleteAdminOutput `json:"data"`
	}
)

type (
	UpdatePasswordInput struct {
		AdminID      uint   `json:"admin_id"`
		PrevPassword string `json:"prev_password"`
		NewPassword  string `json:"new_password"`
	}

	UpdatePasswordOutput struct {
		Msg string `json:"msg"`
	}

	UpdatePasswordSuccessResponse struct {
		Code int                  `json:"code"`
		Data UpdatePasswordOutput `json:"data"`
	}
)

type (
	UpdateEmailInput struct {
		AdminID uint   `json:"admin_id"`
		Email   string `json:"email"`
	}

	UpdateEmailOutput struct {
		Msg string `json:"msg"`
	}

	UpdateEmailSuccessResponse struct {
		Code int               `json:"code"`
		Data UpdateEmailOutput `json:"data"`
	}
)

type (
	UpdateGradeInput struct {
		AdminID uint            `json:"admin_id"`
		Grade   admin.GradeType `json:"grade"`
	}

	UpdateGradeOutput struct {
		Msg string `json:"msg"`
	}

	UpdateGradeSuccessResponse struct {
		Code int               `json:"code"`
		Data UpdateGradeOutput `json:"data"`
	}
)

type (
	BlockAdminInput struct {
		AdminID uint `json:"admin_id" query:"admin_id" default:"1"`
	}

	BlockAdminOutput struct {
		Msg string `json:"msg" default:"Successfully blocked"`
	}

	BlockAdminSuccessResponse struct {
		Code int              `json:"code"`
		Data BlockAdminOutput `json:"data"`
	}
)

type (
	UnblockAdminInput struct {
		AdminID uint `json:"admin_id" query:"admin_id" default:"1"`
	}

	UnblockAdminOutput struct {
		Msg string `json:"msg" default:"Successfully unblocked"`
	}

	UnblockAdminSuccessResponse struct {
		Code int                `json:"code"`
		Data UnblockAdminOutput `json:"data"`
	}
)