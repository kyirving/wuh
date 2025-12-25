package controller

import (
	"basketball/internal/dao/admin"
	"basketball/internal/model"
	"basketball/internal/service"
	"basketball/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PlayerController struct {
	PlayerService *service.PlayerService
	Log           *zap.Logger
}

func NewPlayerController(playerService *service.PlayerService, log *zap.Logger) *PlayerController {
	return &PlayerController{PlayerService: playerService, Log: log}
}

func (pc *PlayerController) Submit(c *gin.Context) {
	var reqBody admin.PlayerReq
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		pc.Log.Error("player submit bind json failed", zap.Error(err))
		util.Output(c, http.StatusBadRequest, nil, "请求参数错误")
		return
	}

	player := &model.Player{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		TeamID:      uint(reqBody.TeamID),
		Number:      uint(reqBody.Number),
		Position:    reqBody.Position,
		Height:      uint(reqBody.Height),
		Weight:      uint(reqBody.Weight),
	}

	if err := pc.PlayerService.Create(c, player); err != nil {
		pc.Log.Error("player submit create failed", zap.Error(err))
		util.Output(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	util.Output(c, http.StatusOK, nil, "创建成功")
}

func (pc *PlayerController) List(c *gin.Context) {
	var reqBody admin.PlayerListReq
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		pc.Log.Error("player list bind json failed", zap.Error(err))
		util.Output(c, http.StatusBadRequest, nil, "请求参数错误")
		return
	}

	players, err := pc.PlayerService.List(c, reqBody)
	if err != nil {
		util.Output(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	util.Output(c, http.StatusOK, players, "查询成功")
}
