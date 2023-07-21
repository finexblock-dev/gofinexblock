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

type SetKeyValueContext[param any, result any] struct {
	Key   param
	Value result
	Err   chan error
}

type GetKeyValueContext[param any, result any] struct {
	Key    param
	Result chan result
}