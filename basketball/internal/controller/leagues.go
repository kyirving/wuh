package controller

import (
	"basketball/internal/dao/admin"
	"basketball/internal/model"
	"basketball/internal/service"
	"basketball/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LeagueController struct {
	leagueService *service.LeagueService
}

func NewLeagueController(service *service.LeagueService) *LeagueController {
	return &LeagueController{
		leagueService: service,
	}
}

func (obj *LeagueController) Submit(c *gin.Context) {
	var reqBody admin.LeagueSubmit

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		util.Output(c, http.StatusBadRequest, nil, "invalid request body")
		return
	}

	league := model.NewLeague(reqBody.Name, reqBody.Description)
	if err := obj.leagueService.Submit(c, league); err != nil {
		util.Output(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	util.Output(c, http.StatusOK, nil, "success")
}

func (obj *LeagueController) List(c *gin.Context) {
	leagues, err := obj.leagueService.List(c)
	if err != nil {
		util.Output(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	util.Output(c, http.StatusOK, leagues, "success")
}
