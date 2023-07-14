package cache

import "errors"

var ErrKeyNotFound = errors.New("key not found")

var ErrCacheFull = errors.New("cache is full")

var ErrCacheEmpty = errors.New("cache is empty")

var ErrSizeTooSmall = errors.New("size is too small")

var ErrKeyAlreadyExist = errors.New("key already exist")
