package repository

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/ent"
	"study-event-go/types"
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
