package dto

type (
	FindUserAssetsInput struct {
		UserID uint `json:"userId" binding:"required" query:"userId"`
	}

	FindUserBalanceUpdateLogInput struct {
		UserID uint `json:"userId" binding:"required" query:"userId"`
		CoinID uint `json:"coinId" binding:"required" query:"coinId"`
		Limit  int  `json:"limit" binding:"required" query:"limit" min:"1" max:"100" default:"20"`
		Offset int  `json:"offset" binding:"required" query:"offset" min:"0" default:"0"`
	}

	FindUserAssetsByCondInput struct {
		UserID uint `json:"userId" binding:"required" query:"userId"`
		CoinID uint `json:"coinId" binding:"required" query:"coinId"`
	}
)
