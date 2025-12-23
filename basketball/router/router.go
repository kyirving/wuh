package router

import (
	"basketball/internal/conf"
	"basketball/internal/controller"
	"basketball/internal/repository"
	"basketball/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, config *conf.Config) *gin.Engine {
	r := gin.Default()

	// 初始化仓储
	userRepository := repository.NewUserRepository(db)

	// 初始化服务
	userService := service.NewUserService(userRepository, config)

	// 初始化控制器
	userController := controller.NewUserController(userService)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
	}

	return r
}
