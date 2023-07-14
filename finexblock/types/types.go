package types

import "time"

type PartialAdmin struct {
	ID           uint      `json:"id"`
	Email        string    `json:"email"`
	Grade        string    `json:"grade"`
	IsBlocked    bool      `json:"is_blocked"`
	InitialLogin bool      `json:"initial_login"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	PwdUpdatedAt time.Time `json:"pwd_updated_at"`
}
