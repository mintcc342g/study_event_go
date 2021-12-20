package repository

import (
	"context"
	"study_event_go/domain/entity"
	"study_event_go/domain/interfaces"
	"study_event_go/ent"
	"study_event_go/types"
)

type gardenRepository struct {
	conn *ent.Client
}

// NewGardenRepository ...
func NewGardenRepository(conn *ent.Client) interfaces.GardenRepository {
	return &gardenRepository{
		conn: conn,
	}
}

func (g *gardenRepository) Garden(ctx context.Context, id types.GardenID) (*entity.Garden, error) {

	// TODO: RDB

	return nil, nil
}
