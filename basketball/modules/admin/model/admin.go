package model

import "gorm.io/gorm"

type Admin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// TableName 返回 Admin 模型对应的数据库表名
// 返回：字符串形式的表名
func (Admin) TableName() string {
	return "admins"
}

// CreateAdmin 在数据库中创建管理员记录
// 入参：db 为 *gorm.DB，admin 为待创建的管理员实体
// 返回：错误信息，nil 表示成功
func CreateAdmin(db *gorm.DB, admin *Admin) error {
	return db.Create(admin).Error
}

// FindAdminByUsername 根据用户名查询管理员
// 入参：db 为 *gorm.DB，username 为管理员用户名
// 返回：管理员实体指针与错误信息；未找到时返回 nil, gorm.ErrRecordNotFound
func FindAdminByUsername(db *gorm.DB, username string) (*Admin, error) {
	var a Admin
	err := db.Where("username = ?", username).First(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}
