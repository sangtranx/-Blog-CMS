package postlikestorage

import (
	"gorm.io/gorm"
)

type sqlstorage struct {
	db *gorm.DB
}

func NewSQLStorage(db *gorm.DB) *sqlstorage {
	return &sqlstorage{db: db}
}
