package auth

import (
	"database/sql"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func (a *authRepository) Conn() *gorm.DB {
	return a.db
}

func newAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db: db}
}

func (a *authRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.db.Begin(&sql.TxOptions{Isolation: level})
}
