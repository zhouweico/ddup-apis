package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID          string     `gorm:"type:char(36);uniqueIndex;not null" json:"uuid"`
	Username      string     `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password      string     `gorm:"size:100;not null" json:"-"` // json:"-" 确保密码不会被序列化
	Email         string     `gorm:"size:100;null" json:"email"`
	Nickname      string     `gorm:"size:50" json:"nickname"`                 // 用户昵称
	Bio           string     `gorm:"size:500" json:"bio"`                     // 用户简介
	Gender        string     `gorm:"size:10;default:'unknown'" json:"gender"` // 性别: male/female/unknown
	Birthday      *time.Time `json:"birthday"`                                // 出生日期
	Avatar        string     `gorm:"size:255" json:"avatar"`                  // 头像URL
	Status        int        `gorm:"default:1;not null" json:"status"`        // 1:正常 2:未验证 0:禁用
	LastLogin     *time.Time `json:"last_login"`
	LoginAttempts int        `gorm:"default:0" json:"-"`
	LockedUntil   *time.Time `json:"-"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// BeforeCreate 创建前的钩子
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UUID = uuid.New().String()
	if u.Nickname == "" {
		u.Nickname = u.Username
	}
	if u.Gender == "" {
		u.Gender = "unknown"
	}
	return nil
}
