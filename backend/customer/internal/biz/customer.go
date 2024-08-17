package biz

import (
	"database/sql"

	"gorm.io/gorm"
)

var (
// ErrUserNotFound is user not found.
)

// Greeter is a Greeter model.
type Customer struct {
	CustomerModel
	gorm.Model
}

// Greeter is a Greeter model.
type CustomerModel struct {
	Telephone      string         `gorm:"type:varchar(15);uniqueIndex;" json:"telephone"`
	Token          string         `gorm:"type:varchar(4095);" json:"token"`
	TokenCreatedAt sql.NullTime   `gorm:"" json:"token_created_at"`
	Name           sql.NullString `gorm:"type:varchar(255);uniqueIndex;" json:"name"`
	Email          sql.NullString `gorm:"type:varchar(255);uniqueIndex;" json:"email"`
	Wechat         sql.NullString `gorm:"type:varchar(255);uniqueIndex;" json:"wechat"`
}
