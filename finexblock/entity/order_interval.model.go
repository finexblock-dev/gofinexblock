package entity

import "time"

type OrderInterval struct {
	ID        uint      `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'"`
	Duration  string    `gorm:"not null;comment:'시간 간격';size:32"`
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
