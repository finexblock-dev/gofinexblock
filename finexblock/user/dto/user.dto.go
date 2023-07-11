package dto

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
)

type (
	FindUserByIDInput struct {
		UserID uint `json:"user_id,required" query:"user_id,required"`
	}

	FindUserByIDOutput struct {
		Result types.Metadata `json:"result,required"`
	}

	FindUserByIDSuccessResponse struct {
		Code int                `json:"code,required"`
		Data FindUserByIDOutput `json:"data,required"`
	}
)

type (
	SearchUserInput struct {
		ID              uint   `json:"id,required" query:"id,required"`
		GradeID         uint   `json:"grade_id,required" query:"grade_id,required"`
		UUID            string `json:"uuid,required" query:"uuid,required"`
		Email           string `json:"email,required" query:"email,required"`
		Nickname        string `json:"nickname,required" query:"nickname,required"`
		Fullname        string `json:"fullname,required" query:"fullname,required"`
		PhoneNumber     string `json:"phone_number,required" query:"phone_number,required"`
		UserType        string `json:"user_type,required" query:"user_type,required"`
		IsBlock         bool   `json:"is_block,required" query:"is_block,required"`
		IsDormant       bool   `json:"is_dormant,required" query:"is_dormant,required"`
		IsMetaverseUser bool   `json:"is_metaverse_user,required" query:"is_metaverse_user,required"`
		IsDropOutUser   bool   `json:"is_drop_out_user,required" query:"is_drop_out_user,required"`
		Description     string `json:"description,required" query:"description,required"`
		Limit           int    `json:"limit,required,required" query:"limit,required,required" default:"20"`
		Offset          int    `json:"offset,required,required" query:"offset,required,required" default:"0"`
	}

	SearchUserOutput struct {
		Result []*types.Metadata `json:"result,required"`
	}

	SearchUserSuccessResponse struct {
		Code int              `json:"code,required"`
		Data SearchUserOutput `json:"data,required"`
	}
)

type (
	BlockUserInput struct {
		UserID uint `json:"user_id,required" query:"user_id,required"`
	}

	BlockUserOutput struct {
		Msg string `json:"msg,required"`
	}

	BlockUserSuccessResponse struct {
		Code int             `json:"code,required"`
		Data BlockUserOutput `json:"data,required"`
	}
)

type (
	CreateMemoInput struct {
		UserID      uint   `json:"user_id,required" default:"79"`
		Description string `json:"description,required" default:"example memo"`
	}

	CreateMemoOutput struct {
		Msg string `json:"msg,required"`
	}

	CreateMemoSuccessResponse struct {
		Code int              `json:"code,required"`
		Data CreateMemoOutput `json:"data,required"`
	}
)