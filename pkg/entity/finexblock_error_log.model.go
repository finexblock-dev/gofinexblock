package entity

import (
	"errors"
	"github.com/finexblock-dev/gofinexblock/pkg/types"
)

type Priority string

func (p Priority) String() string {
	return string(p)
}

func (p Priority) Validate() error {
	switch p {
	case HIGH, MEDIUM, LOW:
		return nil
	}
	return errors.New("invalid priority")
}

const (
	HIGH   Priority = "HIGH"
	MEDIUM Priority = "MEDIUM"
	LOW    Priority = "LOW"
)

type FinexblockErrorLog struct {
	ID          uint           `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'" json:"id"`
	ServerID    uint           `gorm:"not null;comment:서버 id" json:"serverId"`
	Process     string         `gorm:"type:LONGTEXT;not null;comment:프로세스" json:"process"`
	Priority    Priority       `gorm:"type:enum('HIGH', 'MEDIUM', 'LOW');not null;comment:중요도" json:"priority"`
	Description string         `gorm:"type:LONGTEXT;not null;comment:부가 설명" json:"description"`
	Err         string         `gorm:"type:LONGTEXT;not null;comment:에러 메세지" json:"err"`
	Metadata    types.Metadata `gorm:"type:json;comment:첨부 metadata" json:"metadata"`
}

func (e *FinexblockErrorLog) TableName() string {
	return "finexblock_error_log"
}

func (e *FinexblockErrorLog) Alias() string {
	return "finexblock_error_log fel"
}