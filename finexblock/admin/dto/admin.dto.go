package dto

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/admin/types"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"
)

type (
	FindAllAdminInput struct {
		Limit  int `json:"limit,omitempty" query:"limit,omitempty"`
		Offset int `json:"offset,omitempty" query:"offset,omitempty"`
	}

	FindAllAdminOutput struct {
		Result []*types.PartialAdmin
	}

	FindAllAdminSuccessResponse struct {
		Code int                `json:"code,omitempty"`
		Data FindAllAdminOutput `json:"data,omitempty"`
	}
)

type (
	FindAdminByGradeInput struct {
		Limit  int             `json:"limit" query:"limit"`
		Offset int             `json:"offset" query:"offset"`
		Grade  admin.GradeType `json:"grade" query:"grade"`
	}

	FindAdminByGradeOutput struct {
		Result []*admin.Admin `json:"result,omitempty"`
	}

	FindAdminByGradeSwaggerOutput struct {
		Result []*types.PartialAdmin `json:"result,omitempty"`
	}

	FindAdminByGradeSuccessResponse struct {
		Code int                           `json:"code,omitempty"`
		Data FindAdminByGradeSwaggerOutput `json:"data,omitempty"`
	}
)

type (
	DeleteAdminInput struct {
		AdminID uint `json:"admin_id"`
	}

	DeleteAdminOutput struct {
		Msg string `json:"msg,omitempty"`
	}

	DeleteAdminSuccessResponse struct {
		Code int               `json:"code,omitempty"`
		Data DeleteAdminOutput `json:"data,omitempty"`
	}
)

type (
	UpdatePasswordInput struct {
		AdminID      uint   `json:"admin_id"`
		PrevPassword string `json:"prev_password"`
		NewPassword  string `json:"new_password"`
	}

	UpdatePasswordOutput struct {
		Msg string `json:"msg,omitempty"`
	}

	UpdatePasswordSuccessResponse struct {
		Code int                  `json:"code,omitempty"`
		Data UpdatePasswordOutput `json:"data,omitempty"`
	}
)

type (
	UpdateEmailInput struct {
		AdminID uint   `json:"admin_id"`
		Email   string `json:"email"`
	}

	UpdateEmailOutput struct {
		Msg string `json:"msg,omitempty"`
	}

	UpdateEmailSuccessResponse struct {
		Code int               `json:"code,omitempty"`
		Data UpdateEmailOutput `json:"data,omitempty"`
	}
)

type (
	UpdateGradeInput struct {
		AdminID uint            `json:"admin_id"`
		Grade   admin.GradeType `json:"grade"`
	}

	UpdateGradeOutput struct {
		Msg string `json:"msg,omitempty"`
	}

	UpdateGradeSuccessResponse struct {
		Code int               `json:"code,omitempty"`
		Data UpdateGradeOutput `json:"data,omitempty"`
	}
)

type (
	BlockAdminInput struct {
		AdminID uint `json:"adminID" query:"admin_id" default:"1"`
	}

	BlockAdminOutput struct {
		Msg string `json:"msg,omitempty" default:"Successfully blocked"`
	}

	BlockAdminSuccessResponse struct {
		Code int              `json:"code,omitempty"`
		Data BlockAdminOutput `json:"data,omitempty"`
	}
)

type (
	UnblockAdminInput struct {
		AdminID uint `json:"adminID" query:"admin_id" default:"1"`
	}

	UnblockAdminOutput struct {
		Msg string `json:"msg,omitempty" default:"Successfully unblocked"`
	}

	UnblockAdminSuccessResponse struct {
		Code int                `json:"code,omitempty"`
		Data UnblockAdminOutput `json:"data,omitempty"`
	}
)
