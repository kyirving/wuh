package repository

import (
	"basketball/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

// 用户仓储接口 定义了用户相关的数据库操作方法
type UserRepository interface {
	// Register 注册用户
	Register(ctx context.Context, user *model.User) error
	// Create 创建用户
	Create(ctx context.Context, user *model.User) error
	// FindByUsername 根据用户名查询用户
	FindByUsername(ctx context.Context, username string) (*model.User, bool, error)
}

// 实现用户仓储接口
type UserRepo struct {
	Db *gorm.DB
}

// NewUserRepository 创建用户仓储实例
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepo{Db: db}
}

func (u *UserRepo) Register(ctx context.Context, user *model.User) error {
	return u.Db.WithContext(ctx).Create(user).Error
}

func (u *UserRepo) Create(ctx context.Context, user *model.User) error {
	return u.Db.WithContext(ctx).Create(user).Error
}

func (u *UserRepo) FindByUsername(ctx context.Context, username string) (*model.User, bool, error) {
	var user model.User
	result := u.Db.WithContext(ctx).Where("username = ?", username).First(&user)
	// 用户名不存在
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}

	if result.Error != nil {
		return nil, false, result.Error
	}

	return &user, true, nil
}
