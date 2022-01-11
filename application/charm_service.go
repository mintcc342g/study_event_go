package application

import (
	"context"
	"study-event-go/application/dto"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
)

// CharmService ...
type CharmService struct {
	charmRepo interfaces.CharmRepository
}

// NewCharmService ...
func NewCharmService(
	charmRepo interfaces.CharmRepository,
) *CharmService {
	return &CharmService{
		charmRepo: charmRepo,
	}
}

// TODO: logger

// NewCharmCreator ...
func (c *CharmService) NewCharmCreator(ctx context.Context, charmCreatorDTO *dto.CharmCreator) (*dto.CharmCreator, error) {

	// TODO: requestor check with ctx

	charmCreator, err := entity.NewCharmCreator(charmCreatorDTO)
	if err != nil {
		return nil, err
	}

	if charmCreator, err = c.charmRepo.NewCreator(ctx, charmCreator); err != nil {
		return nil, err
	}

	return charmCreator.DTO(), nil
}

// NewCharmModel ...
func (c *CharmService) NewCharmModel(ctx context.Context, charmModelDTO *dto.CharmModel) (*dto.CharmModel, error) {

	// TODO: requestor check with ctx

	if _, err := c.charmRepo.Creator(ctx, charmModelDTO.CreatorID); err != nil {
		return nil, err
	}

	charmModel, err := entity.NewCharmModel(charmModelDTO)
	if err != nil {
		return nil, err
	}

	if charmModel, err = c.charmRepo.NewModel(ctx, charmModel); err != nil {
		return nil, err
	}

	return charmModel.DTO(), nil
}
