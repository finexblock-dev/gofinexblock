package cache

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
	Set(key string, value *T) (err error)
	CurrentSize() (size int)
	Resize(size int) (err error)
	DeleteAll()
}
