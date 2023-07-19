package types

type ErrReceiveContext[param any] struct {
	Tunnel chan error
	Value  param
}

type ResultReceiveContext[param any, result any] struct {
	Result chan result
	Err    chan error
	Value  param
}