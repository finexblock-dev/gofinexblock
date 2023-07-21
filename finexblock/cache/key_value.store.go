package cache

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"time"
)

type DefaultKeyValueStore[T any] struct {
	size  int
	store map[string]*T
	exp   map[string]time.Time
	get   chan types.GetKeyValueContext[string, *T]
	set   chan types.SetKeyValueContext[string, *T]
	del   chan string
}

func (k *DefaultKeyValueStore[T]) ConcurrentRead(key string) (*T, error) {
	ctx := types.GetKeyValueContext[string, *T]{
		Key:    key,
		Result: make(chan *T),
	}
	k.get <- ctx
	value := <-ctx.Result

	if value == nil {
		return nil, ErrKeyNotFound
	}

	return value, nil
}

func (k *DefaultKeyValueStore[T]) ConcurrentWrite(key string, value *T) error {
	ctx := types.SetKeyValueContext[string, *T]{
		Key:   key,
		Value: value,
	}

	k.set <- ctx
	return <-ctx.Err
}

func (k *DefaultKeyValueStore[T]) ConcurrentDelete(key string) {
	if k.IsExist(key) {
		k.del <- key
	}
}

func (k *DefaultKeyValueStore[T]) Subscribe() {
	for {
		select {
		case ctx := <-k.get:
			value, _ := k.Get(ctx.Key)
			ctx.Result <- value
		case ctx := <-k.set:
			ctx.Err <- k.Set(ctx.Key, ctx.Value)
		case key := <-k.del:
			k.Del(key)
		}
	}
}

func (k *DefaultKeyValueStore[T]) IsExist(key string) (exist bool) {
	if _, ok := k.store[key]; !ok {
		return false
	}
	return true
}

func (k *DefaultKeyValueStore[T]) SetNX(key string, value *T) (err error) {
	if k.IsExist(key) {
		return ErrKeyAlreadyExist
	}

	return k.Set(key, value)
}

func (k *DefaultKeyValueStore[T]) SetEX(key string, value *T, duration time.Duration) (err error) {
	if err = k.Set(key, value); err != nil {
		return err
	}

	k.exp[key] = time.Now().Add(duration)
	return nil
}

func (k *DefaultKeyValueStore[T]) Del(key string) {
	delete(k.store, key)
	delete(k.exp, key)
}

func (k *DefaultKeyValueStore[T]) DeleteAll() {
	k.store = make(map[string]*T, k.size)
	k.exp = make(map[string]time.Time, k.size)
	k.size = 0
}

func (k *DefaultKeyValueStore[T]) Get(key string) (value *T, err error) {
	if _, ok := k.store[key]; !ok {
		return nil, ErrKeyNotFound
	}

	if time.Now().After(k.exp[key]) {
		k.Del(key)
		return nil, ErrKeyNotFound
	}
	return k.store[key], nil
}

func (k *DefaultKeyValueStore[T]) Set(key string, value *T) (err error) {
	if k.CurrentSize() >= k.size {
		return ErrCacheFull
	}

	k.store[key] = value
	k.exp[key] = time.Now().Add(time.Hour * 24 * 365)
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

func NewDefaultKeyValueStore[T any](size int) *DefaultKeyValueStore[T] {
	return &DefaultKeyValueStore[T]{size: size, store: make(map[string]*T, size)}
}