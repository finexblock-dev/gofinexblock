package dto

import "time"

type (
	SearchTradeHistoryInput struct {
		UserID        uint      `json:"userId" query:"userId" default:"" validate:"min=1"`
		OrderSymbolID uint      `json:"orderSymbolId" query:"orderSymbolId" default:"15" validate:"min=1"`
		StartTime     time.Time `json:"startTime" query:"startTime" default:"2022-01-01T00:00:00Z"`
		EndTime       time.Time `json:"endTime" query:"endTime" default:"2022-01-01T00:00:00Z"`
		Limit         int       `json:"limit" query:"limit" default:"20" validate:"min=1,max=100"`
		Offset        int       `json:"offset" query:"offset" default:"0" validate:"min=0"`
	}
)