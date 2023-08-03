package dto

type (
	// XRangeInput @XRangeInput
	XRangeInput struct {
		Stream string `json:"stream" validate:"required"`
		Start  string `json:"start" validate:"required"`
		End    string `json:"end" validate:"required"`
	}
)

type (
	// XInfoStreamInput @XInfoStreamInput
	XInfoStreamInput struct {
		Stream string `json:"stream" validate:"required"`
	}
)

type (
	// GetInput @GetInput
	GetInput struct {
		Key string `json:"key" validate:"required"`
	}
)

type (
	// SetInput @SetInput
	SetInput struct {
		Key   string `json:"key" validate:"required"`
		Value string `json:"value" validate:"required"`
	}
)

type (
	// DelInput @DelInput
	DelInput struct {
		Key string `json:"key" validate:"required"`
	}
)