package trade

import (
	"encoding/json"
	"errors"
)

const (
	accountLockPrefix = "account-lock:"
)

var (
	lock, _ = json.Marshal(true)

	ErrKeyNotFound     = errors.New("redis: key not found")
	ErrDecimalParse    = errors.New("decimal parse error")
	ErrNegativeBalance = errors.New("negative balance")
)