package types

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")

	ErrFailedToFindUser = errors.New("failed to find user")

	ErrFailedToSearchUser = errors.New("failed to search user")

	ErrFailedToBlockUser = errors.New("failed to block user")

	ErrFailedToCreateMemo = errors.New("failed to create memo")
)