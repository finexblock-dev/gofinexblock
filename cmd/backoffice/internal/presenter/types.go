package presenter

type ErrResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type MsgResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func NewErrResponse(err error) *ErrResponse {
	return &ErrResponse{
		Status:  false,
		Message: err.Error(),
	}
}