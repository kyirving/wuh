package controller

import (
	"basketball/internal/dao/admin"
	"basketball/internal/model"
	"basketball/internal/service"
	"basketball/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamController struct {
	TeamService *service.TeamService
}

func NewTeamController(teamService *service.TeamService) *TeamController {
	return &TeamController{
		TeamService: teamService,
	}
}

// Submit 创建球队
func (tc *TeamController) Submit(ctx *gin.Context) {
	var reqBody admin.TeamDao
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		util.Output(ctx, http.StatusBadRequest, gin.H{"error": err.Error()}, "invalid request body")
		return
	}

	if reqBody.Name == "" {
		util.Output(ctx, http.StatusBadRequest, gin.H{"error": "name is required"}, "name is required")
		return
	}

	if reqBody.LeagueID == 0 {
		util.Output(ctx, http.StatusBadRequest, gin.H{"error": "league id is required"}, "league id is required")
		return
	}

	team := &model.Team{
		ID:          reqBody.Id,
		Name:        reqBody.Name,
		Description: reqBody.Description,
		LeagueID:    reqBody.LeagueID,
	}

	if err := tc.TeamService.Submit(ctx, team); err != nil {
		util.Output(ctx, http.StatusInternalServerError, nil, "failed to submit team")
		return
	}
	util.Output(ctx, http.StatusOK, gin.H{"id": team.ID}, "team submitted")
}

// List 获取球队列表，返回球队基础信息及所属联赛名称
func (tc *TeamController) List(ctx *gin.Context) {
	list, err := tc.TeamService.List(ctx)
	if err != nil {
		util.Output(ctx, http.StatusInternalServerError, nil, "failed to list teams")
		return
	}
	util.Output(ctx, http.StatusOK, list, "teams listed")
}
