package dto

type (
	FindUserAssetsInput struct {
		UserID uint `json:"userId" binding:"required" query:"userId"`
	}

	FindUserBalanceUpdateLogInput struct {
		UserID uint `json:"userId" binding:"required" query:"userId"`
		Limit  int  `json:"limit" binding:"required" query:"limit" min:"1" max:"100"`
		Offset int  `json:"offset" binding:"required" query:"offset" min:"0"`
	}
)