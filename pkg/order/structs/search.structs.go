package structs

import "time"

type (
	SearchOrderMatchingHistoryInput struct {
		UserID        uint
		OrderSymbolID uint
		StartTime     time.Time
		EndTime       time.Time
		Limit         int
		Offset        int
	}
)