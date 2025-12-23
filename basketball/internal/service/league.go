package service

import (
	"basketball/internal/model"
	"basketball/internal/repository"
	"context"
	"errors"
)

type LeagueService struct {
	Repo repository.LeagueRepoInterface
}

func NewLeagueService(repo repository.LeagueRepoInterface) *LeagueService {
	return &LeagueService{Repo: repo}
}

func (obj *LeagueService) Submit(ctx context.Context, league *model.League) error {

	exist, err := obj.Repo.Exist(ctx, league.Name)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("league name already exist")
	}
	return obj.Repo.Submit(ctx, league)
}
