package dto

type (
	ListImageInput struct {
		Limit  int `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)
