package cache

type DefaultKeyValueStore[T any] struct {
	size  int
	store map[string]*T
}

func NewDefaultKeyValueStore[T any](size int) *DefaultKeyValueStore[T] {
	return &DefaultKeyValueStore[T]{size: 0, store: make(map[string]*T, size)}
}

func (k *DefaultKeyValueStore[T]) DeleteAll() {
	k.store = make(map[string]*T, k.size)
	k.size = 0
}

func (k *DefaultKeyValueStore[T]) Get(key string) (value *T, err error) {
	if _, ok := k.store[key]; !ok {
		return nil, ErrKeyNotFound
	}
	return k.store[key], nil
}

func (k *DefaultKeyValueStore[T]) Set(key string, value *T) (err error) {
	if k.CurrentSize() >= k.size {
		return ErrCacheFull
	}

	k.store[key] = value
	return nil
}

func (k *DefaultKeyValueStore[T]) CurrentSize() (size int) {
	return len(k.store)
}

func (k *DefaultKeyValueStore[T]) Resize(size int) (err error) {
	if size < k.CurrentSize() {
		return ErrSizeTooSmall
	}
	k.size = size
	return nil
}
