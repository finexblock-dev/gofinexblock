package dto

type (
	ScanWithdrawalRequestByStatusInput struct {
		Status string `json:"status" binding:"required" query:"status" example:"SUBMITTED, APPROVED, CANCELED, REJECTED, PENDING, COMPLETED, FAILED" default:"SUBMITTED"`
		Limit  int    `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int    `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}

	FindWithdrawalRequestsByUserIDInput struct {
		UserID uint `json:"userId" binding:"required" query:"userId"`
		Limit  int  `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int  `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}

	RejectWithdrawalRequestsInput struct {
		ID uint `json:"id" binding:"required" query:"id" validate:"min=1"`
	}

	ApproveWithdrawalRequestsInput struct {
		ID uint `json:"id" binding:"required" query:"id" validate:"min=1"`
	}
)