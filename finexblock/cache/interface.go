package cache

type KeyValueStore[T any] interface {
	Cache[T]
}

type LeastRecentlyUsedCache[T any] interface {
	Cache[T]
	// TODO: implement this
}

type LeastFrequentlyUsedCache[T any] interface {
	Cache[T]
	// TODO: implement this
}

type Cache[T any] interface {
	Get(key string) (value *T, err error)
	Set(key string, value *T) (err error)
	CurrentSize() (size int)
	Resize(size int) (err error)
	DeleteAll()
}
