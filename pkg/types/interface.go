package types

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type StringEnum interface {
	fmt.Stringer
	Validate() error
}

type Model interface {
	TableName() string
	Alias() string
}

type Repository interface {
	Tx(level sql.IsolationLevel) *gorm.DB
	Conn() *gorm.DB
}

type Service interface {
	Ctx() context.Context
	CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc)
	Tx(level sql.IsolationLevel) *gorm.DB
	Conn() *gorm.DB
}

type Heap[T any] interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Pop() T
	Push(x T)
}

type BaseClaimer interface {
	Claim()
	Claimer(claimer Claimer) Claimer
}

type SingleStreamClaimer interface {
	BaseClaimer
	ReadPendingStream(stream Stream, group Group) (*redis.XPending, error)
	ClaimStream(stream Stream, group Group, claimer Claimer, minIdleTime time.Duration, ids []string) ([]redis.XMessage, error)
}

type BaseConsumer interface {
	Consume()
	Consumer(consumer Consumer) Consumer
}

type SingleStreamConsumer interface {
	BaseConsumer
	ReadStream(stream Stream, group Group, consumer Consumer, count int64, block time.Duration) ([]redis.XStream, error)
}

type MultiStreamConsumer interface {
	BaseConsumer
	ReadStreams(streams []Stream, group Group, consumer Consumer, count int64, block time.Duration) ([]redis.XStream, error)
}

type Metadata map[string]interface{}

func (m Metadata) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func (m *Metadata) Scan(value interface{}) error {
	if value == nil {
		*m = nil
		return nil
	}

	return json.Unmarshal([]byte(fmt.Sprintf("%s", value)), &m)
}