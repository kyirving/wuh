package service

import (
	"basketball/internal/conf"
	"basketball/internal/model"
	"basketball/internal/repository"
	"basketball/pkg/auth"
	"basketball/pkg/util"
	"context"
	"errors"
	"time"

	"github.com/spf13/cast"
)

type UserService struct {
	Repo   repository.UserRepository
	config *conf.Config
}

// NewUserService 创建用户服务实例
func NewUserService(repo repository.UserRepository, config *conf.Config) *UserService {
	return &UserService{Repo: repo, config: config}
}

func (u *UserService) Register(ctx context.Context, user *model.User) error {
	return u.Repo.Register(ctx, user)
}

func (u *UserService) Login(ctx context.Context, username, password string) (*auth.JWTToken, error) {
	user, exists, err := u.Repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("用户不存在")
	}

	if user.Password != util.EncryptPassword(password) {
		return nil, errors.New("password not match")
	}

	// 生成 JWT 令牌
	jwt := auth.NewJWTProvider([]byte(u.config.Manager.JWTSecret), user.Username, time.Duration(u.config.Manager.AccessExpire), time.Duration(u.config.Manager.RefreshExpire))
	token, err := jwt.GenerateToken(cast.ToInt64(user.ID), []string{"admin"})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (u *UserService) FindByUsername(ctx context.Context, username string) (*model.User, bool, error) {
	return u.Repo.FindByUsername(ctx, username)
}
