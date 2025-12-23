package controller

import (
	"basketball/internal/dao/admin"
	"basketball/internal/model"
	"basketball/internal/service"
	"basketball/pkg/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type UserController struct {
	Service *service.UserService
}

// NewUserController 创建用户控制器实例
func NewUserController(service *service.UserService) *UserController {
	return &UserController{Service: service}
}

// Register 注册用户
func (u *UserController) Register(ctx *gin.Context) {

	var req admin.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.Output(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	if req.Username == "" || req.Password == "" {
		util.Output(ctx, http.StatusBadRequest, nil, "username or password is empty")
		return
	}
	_, exists, err := u.Service.FindByUsername(ctx, req.Username)
	if err != nil {
		util.Output(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	if exists {
		util.Output(ctx, http.StatusBadRequest, nil, "username already exists")
		return
	}

	user := model.NewUser(req.Username, req.Password)
	err = u.Service.Register(ctx, user)
	fmt.Println(err)
	if err != nil {
		util.Output(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	util.Output(ctx, http.StatusOK, map[string]float32{"id": cast.ToFloat32(user.ID)}, "register success")
}

// Login 登录用户
func (u *UserController) Login(ctx *gin.Context) {

	var req admin.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.Output(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	if req.Username == "" || req.Password == "" {
		util.Output(ctx, http.StatusBadRequest, nil, "username or password is empty")
		return
	}

	token, err := u.Service.Login(ctx, req.Username, req.Password)
	if err != nil {
		util.Output(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	util.Output(ctx, http.StatusOK, token, "login success")
}
