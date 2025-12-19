package route

import (
	"basketball/modules/admin/controller"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由并注入数据库连接
// 入参：db 为已初始化的 *gorm.DB
// 返回：*gin.Engine Gin 引擎实例
func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	adminController := controller.NewAdminController(db)
	r.POST("/user/login", adminController.Login)
	r.POST("/user/register", adminController.Register)

	return r
}
