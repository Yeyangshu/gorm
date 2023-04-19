package model

import (
	"database/sql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// User钩子
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	ctx := tx.Statement.Context
	value := ctx.Value("context")
	log.Printf("hook_context get value: %s", value)

	// 获取gorm设置值
	setKey, ok := tx.Get("setKey")
	if ok {
		log.Printf("hook_context get set value: %t", setKey)
	}

	return nil
}
