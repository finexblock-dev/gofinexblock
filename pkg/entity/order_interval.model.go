package entity

import (
	"github.com/finexblock-dev/gofinexblock/pkg/types"
	"time"
)

type OrderInterval struct {
	ID        uint           `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'" json:"id"`
	Duration  types.Duration `gorm:"not null;comment:'시간 간격';size:32" json:"duration"`
	StartTime time.Time      `gorm:"not null;type:timestamp;comment:'시작 시간'" json:"startTime"`
	EndTime   time.Time      `gorm:"not null;type:timestamp;comment:'종료 시간'" json:"endTime"`

	OrderMatchingEvent []OrderMatchingEvent `gorm:"foreignKey:OrderIntervalID;constraint:OnUpdate:CASCADE;" json:"orderMatchingEvent"`
	Chart              []Chart              `gorm:"foreignKey:OrderIntervalID;constraint:OnUpdate:CASCADE;" json:"chart"`
}

func (m *OrderInterval) Alias() string {
	return "order_interval oi"
}

func (m *OrderInterval) TableName() string {
	return "order_interval"
}