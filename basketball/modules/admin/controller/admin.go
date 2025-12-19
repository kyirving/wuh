package controller

import (
	"basketball/dao/admin"
	adminModel "basketball/modules/admin/model"
	"basketball/pkg/util"
	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	DB *gorm.DB
}

// NewAdminController 创建并返回带数据库依赖的 AdminController
// 入参：db 为已初始化的 *gorm.DB
// 返回：*AdminController 控制器实例
func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{DB: db}
}

func (a *AdminController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
	})
}

func (a *AdminController) Register(c *gin.Context) {

	var reqBody admin.RegisterReq
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		util.Output(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	// 检查用户名是否已存在
	user, err := adminModel.FindAdminByUsername(a.DB, reqBody.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		util.Output(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	if user != nil {
		util.Output(c, http.StatusBadRequest, nil, "用户名已存在")
		return
	}

	// 创建新管理员
	toCreate := &adminModel.Admin{Username: reqBody.Username, Password: reqBody.Password}
	if err := adminModel.CreateAdmin(a.DB, toCreate); err != nil {
		util.Output(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	util.Output(c, http.StatusOK, gin.H{"id": toCreate.ID}, "register success")
}
