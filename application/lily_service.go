package application

import (
	"context"
	"study-event-go/application/dto"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/types"
)

// LilyService ...
type LilyService struct {
	charmRepo  interfaces.CharmRepository
	gardenRepo interfaces.GardenRepository
	lilyRepo   interfaces.LilyRepository
}

// NewLilyService ...
func NewLilyService(
	charmRepo interfaces.CharmRepository,
	gardenRepo interfaces.GardenRepository,
	lilyRepo interfaces.LilyRepository,
) *LilyService {
	return &LilyService{
		charmRepo:  charmRepo,
		gardenRepo: gardenRepo,
		lilyRepo:   lilyRepo,
	}
}

// TODO: logger

// New ...
func (l *LilyService) New(ctx context.Context, lilyDTO *dto.Lily) (*dto.Lily, error) {

	// TODO: requestor check with ctx

	_, err := l.gardenRepo.Garden(ctx, lilyDTO.GardenID)
	if err != nil {
		return nil, err
	}

	lily, err := entity.NewLily(lilyDTO)
	if err != nil {
		return nil, err
	}

	if lily, err = l.lilyRepo.Save(ctx, lily); err != nil {
		return nil, err
	}

	return lily.DTO(), nil
}

// NewCharms ...
func (l *LilyService) NewCharms(ctx context.Context, lilyID types.LilyID, charmModelIDs []types.CharmModelID) (results []*dto.Charm, err error) {

	// TODO: requestor and creator check with ctx

	lily, err := l.lilyRepo.Lily(ctx, lilyID)
	if err != nil {
		return
	}

	charmModels, err := l.charmRepo.ModelsByIDs(ctx, charmModelIDs)
	if err != nil {
		return
	}

	if err = lily.ContractWith(charmModelIDs, charmModels...); err != nil {
		return
	}

	charms, err := l.charmRepo.SaveCharms(ctx, lily.Charms...)
	if err != nil {
		return
	}

	for _, v := range charms {
		results = append(results, &dto.Charm{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
			Name:      v.Name,
			ModelID:   v.ModelID,
			OwnerID:   v.OwnerID,
		})
	}

	return
}
