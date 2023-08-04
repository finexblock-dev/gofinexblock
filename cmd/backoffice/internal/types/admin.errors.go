package types

import "errors"

var (
	ErrFailedToFindAdmin = errors.New("failed to find admin")

	ErrFailedToUpdateEmail    = errors.New("failed to update email")
	ErrFailedToUpdatePassword = errors.New("failed to update password")
	ErrFailedToUpdateGrade    = errors.New("failed to update grade")

	ErrFailedToFindLoginFailedLog = errors.New("failed to find failed log")
	ErrFailedToFindLoginHistory   = errors.New("failed to find login history")

	ErrFailedToSearchApiLog            = errors.New("failed to search api log")
	ErrFailedToSearchDeleteLog         = errors.New("failed to search delete log")
	ErrFailedToSearchGradeUpdateLog    = errors.New("failed to search grade update log")
	ErrFailedToSearchPasswordUpdateLog = errors.New("failed to search password update log")
	ErrFailedToFindAccessToken         = errors.New("failed to find access token")

	ErrFailedToDeleteAdmin = errors.New("failed to delete admin")

	ErrFailedToBlockAdmin   = errors.New("failed to block admin")
	ErrFailedToUnblockAdmin = errors.New("failed to unblock admin")

	ErrCheckPassword = errors.New("check password")
	ErrInvalidGrade  = errors.New("invalid grade")
)