package types

import "errors"

var (
	ErrFailedToParseQuery = errors.New("failed to parse query")

	ErrFailedToParseBody = errors.New("failed to parse body")

	ErrFailedToParseImages = errors.New("failed to parse images")
)