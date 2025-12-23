package service

import (
	"basketball/internal/model"
	"basketball/internal/repository"
	"context"
	"errors"
)

type TeamService struct {
	TeamRepo   repository.TeamRepoInterface
	LeagueRepo repository.LeagueRepoInterface
}

func NewTeamService(teamRepo repository.TeamRepoInterface, leagueRepo repository.LeagueRepoInterface) *TeamService {
	return &TeamService{
		TeamRepo:   teamRepo,
		LeagueRepo: leagueRepo,
	}
}

func (obj *TeamService) Submit(ctx context.Context, team *model.Team) error {

	leagueExist, err := obj.LeagueRepo.FindByID(ctx, uint(team.LeagueID))
	if err != nil {
		return err
	}
	if leagueExist == nil {
		return errors.New("league not exist")
	}

	exist, err := obj.TeamRepo.Exist(ctx, team)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("team already exist")
	}

	return obj.TeamRepo.Submit(ctx, team)
}
