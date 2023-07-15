package types

type ErrReceiveContext[T any] struct {
	Tunnel chan error
	Value  T
}

type ResultReceiveContext[param any, result any] struct {
	Tunnel chan result
	Value  param
}
