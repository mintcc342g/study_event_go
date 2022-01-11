package application

import (
	"context"
	"study-event-go/application/dto"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
)

// LilyService ...
type LilyService struct {
	// charmRepo  interfaces.CharmRepository
	gardenRepo interfaces.GardenRepository
	lilyRepo   interfaces.LilyRepository
}

// NewLilyService ...
func NewLilyService(
	// charmRepo interfaces.CharmRepository,
	gardenRepo interfaces.GardenRepository,
	lilyRepo interfaces.LilyRepository,
) *LilyService {
	return &LilyService{
		// charmRepo:  charmRepo,
		gardenRepo: gardenRepo,
		lilyRepo:   lilyRepo,
	}
}

// TODO: logger

// New ...
func (l *LilyService) New(ctx context.Context, lilyDTO *dto.Lily) (*dto.Lily, error) {

	// TODO: requestor check with ctx

	_, err := l.gardenRepo.Get(ctx, lilyDTO.GardenID)
	if err != nil {
		return nil, err
	}

	lily, err := entity.NewLily(lilyDTO)
	if err != nil {
		return nil, err
	}

	if lily, err = l.lilyRepo.New(ctx, lily); err != nil {
		return nil, err
	}

	return lily.DTO(), nil
}
