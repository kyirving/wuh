package controller

import (
	"basketball/internal/dao/admin"
	"basketball/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GameController struct {
	gameService *service.GameService
	Log         *zap.Logger
}

func NewGameController(gameService *service.GameService, log *zap.Logger) *GameController {
	return &GameController{
		gameService: gameService,
		Log:         log,
	}
}

func (c *GameController) Submit(ctx *gin.Context) {
	var req admin.GameSubmitReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 提交游戏
	if err := c.gameService.Submit(ctx, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (c *GameController) List(ctx *gin.Context) {
	var req admin.GameListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// games, err := c.gameService.List(ctx, req)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, gin.H{"games": games})
}
