package router

import (
	"basketball/internal/conf"
	"basketball/internal/controller"
	"basketball/internal/middleware"
	"basketball/internal/repository"
	"basketball/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, config *conf.Config) *gin.Engine {
	r := gin.Default()

	conf := cors.DefaultConfig()
	conf.AllowOrigins = []string{"http://localhost:9527", "http://127.0.0.1:9527"} // 允许的源
	conf.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	conf.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	conf.AllowCredentials = true
	conf.MaxAge = 12 * 60 * 60 // 缓存时间
	r.Use(cors.New(conf))

	// 初始化仓储
	userRepository := repository.NewUserRepository(db)
	leagueRepository := repository.NewLeagueRepo(db)
	teamRepository := repository.NewTeamRepository(db)
	playerRepository := repository.NewPlayerRepository(db)

	// 初始化服务
	userService := service.NewUserService(userRepository, config)
	leagueService := service.NewLeagueService(leagueRepository)
	teamService := service.NewTeamService(teamRepository, leagueRepository)
	playerService := service.NewPlayerService(playerRepository)

	// 初始化控制器
	userController := controller.NewUserController(userService)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
	}

	// 对所有其他路由应用JWT中间件
	r.Use(middleware.JWT(config))
	// 初始化控制器
	leagueController := controller.NewLeagueController(leagueService)
	playerController := controller.NewPlayerController(playerService)
	teamController := controller.NewTeamController(teamService)

	leagueGroup := r.Group("/league")
	{
		leagueGroup.POST("/submit", leagueController.Submit)
		leagueGroup.POST("/list", leagueController.List)
	}
	teamGroup := r.Group("/team")
	{
		teamGroup.POST("/submit", teamController.Submit)
	}
	playerGroup := r.Group("/player")
	{
		playerGroup.POST("/submit", playerController.Submit)
	}
	return r
}
