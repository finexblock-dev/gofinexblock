package structs

import (
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
)

type (
	FindUserByIDInput struct {
		UserID uint `json:"userId,required" query:"userId,required"`
	}

	FindUserByIDOutput struct {
		Result entity.UserMetadata `json:"result,required"`
	}

	FindUserByIDSuccessResponse struct {
		Code int                `json:"code,required"`
		Data FindUserByIDOutput `json:"data,required"`
	}
)

type (
	SearchUserInput struct {
		ID              uint   `json:"id" query:"id"`
		GradeID         uint   `json:"gradeId" query:"gradeId"`
		UUID            string `json:"uuid" query:"uuid"`
		Email           string `json:"email" query:"email"`
		Nickname        string `json:"nickname" query:"nickname"`
		Fullname        string `json:"fullname" query:"fullname"`
		PhoneNumber     string `json:"phoneNumber" query:"phoneNumber"`
		UserType        string `json:"userType" query:"userType"`
		IsAdult         bool   `json:"isAdult" query:"isAdult"`
		IsBlock         bool   `json:"isBlock" query:"isBlock"`
		IsDormant       bool   `json:"isDormant" query:"isDormant"`
		IsMetaverseUser bool   `json:"isMetaverseUser" query:"isMetaverseUser"`
		IsDropOutUser   bool   `json:"isDropOutUser" query:"isDropOutUser"`
		Description     string `json:"description" query:"description"`
		Limit           int    `json:"limit,required" query:"limit,required" default:"20"`
		Offset          int    `json:"offset,required" query:"offset,required" default:"0"`
	}

	SearchUserOutput struct {
		Result []*entity.UserMetadata `json:"result,required"`
	}

	SearchUserSuccessResponse struct {
		Code int              `json:"code,required"`
		Data SearchUserOutput `json:"data,required"`
	}
)

type (
	BlockUserInput struct {
		UserID uint `json:"userId,required" query:"userId,required"`
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
		UserID      uint   `json:"userId,required" default:"79"`
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