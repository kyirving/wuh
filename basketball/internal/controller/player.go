package controller

import (
	"basketball/internal/dao/admin"
	"basketball/internal/model"
	"basketball/internal/service"
	"basketball/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	PlayerService *service.PlayerService
}

func NewPlayerController(playerService *service.PlayerService) *PlayerController {
	return &PlayerController{PlayerService: playerService}
}

func (pc *PlayerController) Submit(c *gin.Context) {
	var reqBody admin.PlayerReq
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		util.Output(c, http.StatusBadRequest, nil, "请求参数错误")
		return
	}

	player := &model.Player{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		TeamID:      reqBody.TeamID,
		Number:      reqBody.Number,
		Position:    reqBody.Position,
		Height:      reqBody.Height,
		Weight:      reqBody.Weight,
	}

	if err := pc.PlayerService.Create(c, player); err != nil {
		util.Output(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	util.Output(c, http.StatusOK, nil, "创建成功")
}
