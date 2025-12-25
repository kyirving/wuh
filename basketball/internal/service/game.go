package service

import (
	"basketball/internal/dao/admin"
	"basketball/internal/model"
	"basketball/internal/repository"
	"context"
	"errors"
	"fmt"
	"time"
)

type GameService struct {
	gameRepo repository.GameRepositoryInterface
}

func NewGameService(gameRepo repository.GameRepositoryInterface) *GameService {
	return &GameService{gameRepo: gameRepo}
}

func (obj *GameService) List(ctx context.Context, req admin.GameListReq) ([]model.Game, int64, error) {
	return obj.gameRepo.List(ctx, req)
}

func (obj *GameService) Submit(ctx context.Context, game admin.GameSubmitReq) error {
	if game.HomeTeamID == game.AwayTeamID {
		return errors.New("home team and away team cannot be the same")
	}

	fmt.Println(game)

	var gameModel = &model.Game{}
	gameModel.Id = int64(game.Id)
	gameModel.LeagueID = game.LeagueID
	gameModel.Name = game.Name
	gameModel.HomeTeamID = game.HomeTeamID
	gameModel.AwayTeamID = game.AwayTeamID
	gameModel.DateTime, _ = time.Parse(time.DateTime, game.DateTime)
	gameModel.Status = int64(game.Status)
	gameModel.Description = game.Description

	if game.Id == 0 {
		return obj.gameRepo.Create(ctx, gameModel)
	}
	return obj.gameRepo.Update(ctx, gameModel)
}
