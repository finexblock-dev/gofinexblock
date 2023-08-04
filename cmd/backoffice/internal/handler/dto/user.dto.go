package dto

type (
	FindUserByIDInput struct {
		UserID uint `json:"userId" query:"userId" binding:"required"`
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
		IsBlock         bool   `json:"isBlock" query:"isBlock"`
		IsDormant       bool   `json:"isDormant" query:"isDormant"`
		IsMetaverseUser bool   `json:"isMetaverseUser" query:"isMetaverseUser"`
		IsDropOutUser   bool   `json:"isDropOutUser" query:"isDropOutUser"`
		Description     string `json:"description" query:"description"`
		Limit           int    `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset          int    `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	BlockUserInput struct {
		UserID uint `json:"userId" query:"userId" binding:"required"`
	}

	UnblockUserInput struct {
		UserID uint `json:"userId" query:"userId" binding:"required"`
	}
)

type (
	CreateMemoInput struct {
		UserID      uint   `json:"userId" default:"79" binding:"required"`
		Description string `json:"description" default:"example memo" binding:"required"`
	}
)
