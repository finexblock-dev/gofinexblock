package entity

import "time"

type Duration string

const (
	OneMinute     Duration = "ONE_MINUTE"
	FiveMinute    Duration = "FIVE_MINUTE"
	ThreeMinute            = "THREE_MINUTE"
	FifteenMinute          = "FIFTEEN_MINUTE"
	ThirtyMinute           = "THIRTY_MINUTE"
	OneHour                = "ONE_HOUR"
	TwoHour                = "TWO_HOUR"
	FourHour               = "FOUR_HOUR"
	SixHour                = "SIX_HOUR"
	EightHour              = "EIGHT_HOUR"
	TwelveHour             = "TWELVE_HOUR"
	OneDay                 = "ONE_DAY"
	ThreeDay               = "THREE_DAY"
	OneWeek                = "ONE_WEEK"
	OneMonth               = "ONE_MONTH"
)

type OrderInterval struct {
	ID        uint      `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'"`
	Duration  Duration  `gorm:"not null;comment:'시간 간격';size:32"`
	StartTime time.Time `gorm:"not null;type:timestamp;comment:'시작 시간'"`
	EndTime   time.Time `gorm:"not null;type:timestamp;comment:'종료 시간'"`

	OrderMatchingEvent []OrderMatchingEvent `gorm:"foreignKey:OrderIntervalID;constraint:OnUpdate:CASCADE;"`
	Chart              []Chart              `gorm:"foreignKey:OrderIntervalID;constraint:OnUpdate:CASCADE;"`
}

func (m *OrderInterval) Alias() string {
	return "order_interval oi"
}

func (m *OrderInterval) TableName() string {
	return "order_interval"
}
