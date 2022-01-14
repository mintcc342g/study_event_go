package application

import (
	"context"
	"study-event-go/application/dto"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/types"

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

	_, err := g.gardenRepo.GardenByName(ctx, gardenDTO.Name)
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	if _, err = g.mentorshipRepo.Mentorship(ctx, gardenDTO.MentorshipID); err != nil {
		return nil, err
	}

	garden, err := entity.NewGarden(gardenDTO)
	if err != nil {
		return nil, err
	}

	if garden, err = g.gardenRepo.Save(ctx, garden); err != nil {
		return nil, err
	}

	return garden.DTO(), nil
}

// Get ...
func (g *GardenService) Get(ctx context.Context, id types.GardenID) (*dto.Garden, error) {

	garden, err := g.gardenRepo.Garden(ctx, id)
	if err != nil {
		return nil, err
	}

	return garden.DTO(), nil
}

// List ...
func (g *GardenService) List(ctx context.Context, offset uint32) ([]*dto.Garden, error) {

	// TODO: cursor pagination

	gardens, err := g.gardenRepo.Gardens(ctx, offset)
	if err != nil {
		return nil, err
	}

	results := make([]*dto.Garden, len(gardens))
	for i, v := range gardens {
		results[i] = &dto.Garden{
			ID:           v.ID,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			DeletedAt:    v.DeletedAt,
			Name:         v.Name,
			Location:     v.Location,
			MentorshipID: v.MentorshipID,
		}
	}

	return results, nil
}

// Update ...
func (g *GardenService) Update(ctx context.Context, gardenDTO *dto.Garden) (*dto.Garden, error) {

	comparable, err := g.gardenRepo.GardenByName(ctx, gardenDTO.Name)
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	if gardenDTO.MentorshipID != 0 {
		if _, err = g.mentorshipRepo.Mentorship(ctx, gardenDTO.MentorshipID); err != nil {
			return nil, err
		}
	}

	garden, err := g.gardenRepo.Garden(ctx, gardenDTO.ID)
	if err != nil {
		return nil, err
	}

	if err = garden.Update(gardenDTO, comparable); err != nil {
		return nil, err
	}

	if garden, err = g.gardenRepo.Update(ctx, garden); err != nil {
		return nil, err
	}

	return garden.DTO(), nil
}

// SoftDelete ...
func (g *GardenService) SoftDelete(ctx context.Context, id types.GardenID) (*dto.Garden, error) {

	garden, err := g.gardenRepo.Garden(ctx, id)
	if err != nil {
		return nil, err
	}

	garden.Delete()

	if garden, err = g.gardenRepo.Update(ctx, garden); err != nil {
		return nil, err
	}

	return garden.DTO(), nil
}
