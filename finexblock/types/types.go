package types

type ErrReceiveContext[param any] struct {
	Tunnel chan error
	Value  param
}

type ResultReceiveContext[param any, result any] struct {
	Tunnel chan result
	Value  param
}
