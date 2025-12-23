package service

import (
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
