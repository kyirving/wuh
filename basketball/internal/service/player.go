package service

import (
	"basketball/internal/dao/admin"
	"basketball/internal/model"
	"basketball/internal/repository"
	"context"
)

type PlayerService struct {
	PlayerRepo repository.PlayerRepositoryInterface
}

func NewPlayerService(playerRepository repository.PlayerRepositoryInterface) *PlayerService {
	return &PlayerService{PlayerRepo: playerRepository}
}

func (obj *PlayerService) Create(ctx context.Context, player *model.Player) error {
	return obj.PlayerRepo.Create(ctx, player)
}

func (obj *PlayerService) List(ctx context.Context, req admin.PlayerListReq) (admin.PlayerListResp, error) {
	players, total, err := obj.PlayerRepo.List(ctx, req)
	if err != nil {
		return admin.PlayerListResp{}, err
	}
	res := make([]admin.PlayerListItem, 0, len(players))
	for _, player := range players {
		res = append(res, admin.PlayerListItem{
			Id:        int(player.BaseModel.ID),
			Name:      player.Name,
			Number:    int(player.Number),
			Position:  player.Position,
			Height:    int(player.Height),
			Weight:    int(player.Weight),
			TeamID:    int(player.TeamID),
			TeamName:  player.Team.Name,
			CreatedAt: player.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return admin.PlayerListResp{
		Players: res,
		Total:   total,
	}, nil
}
