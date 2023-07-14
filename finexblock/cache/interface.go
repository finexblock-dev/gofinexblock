package cache

import "time"

type LeastRecentlyUsedCache[T any] interface {
	KeyValueStore[T]
	// TODO: implement this
}

type LeastFrequentlyUsedCache[T any] interface {
	KeyValueStore[T]
	// TODO: implement this
}

type KeyValueStore[T any] interface {
	Get(key string) (value *T, err error)
	IsExist(key string) (exist bool)

	Set(key string, value *T) (err error)
	SetNX(key string, value *T) (err error)
	SetEX(key string, value *T, duration time.Duration) (err error)

	CurrentSize() (size int)
	Resize(size int) (err error)

	Del(key string)
	DeleteAll()
}
