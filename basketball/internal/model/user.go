package model

import (
	"basketball/pkg/util"

	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	BaseModel
}

// NewUser 创建用户实例
func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

// 钩子函数 在创建用户前调用
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// 对密码进行加密
	u.Password = util.EncryptPassword(u.Password)
	return
}
