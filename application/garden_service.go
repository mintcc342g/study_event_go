package application

import (
	"context"
	"study-event-go/application/dto"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"

	"github.com/juju/errors"
)

// GardenService ...
type GardenService struct {
	gardenRepo     interfaces.GardenRepository
	mentorshipRepo interfaces.MentorshipRepository
}

// NewGardenService ...
func NewGardenService(
	gardenRepo interfaces.GardenRepository,
	mentorshipRepo interfaces.MentorshipRepository,
) *GardenService {
	return &GardenService{
		gardenRepo:     gardenRepo,
		mentorshipRepo: mentorshipRepo,
	}
}

// TODO: logger

// New ...
func (g *GardenService) New(ctx context.Context, gardenDTO *dto.Garden) (*dto.Garden, error) {

	_, err := g.gardenRepo.GetByName(ctx, gardenDTO.Name)
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	if gardenDTO.MentorshipID != 0 {
		if _, err = g.mentorshipRepo.Get(ctx, gardenDTO.MentorshipID); err != nil {
			return nil, err
		}
	}

	garden, err := entity.NewGarden(gardenDTO)
	if err != nil {
		return nil, err
	}

	if garden, err = g.gardenRepo.New(ctx, garden); err != nil {
		return nil, err
	}

	return garden.DTO(), nil
}
