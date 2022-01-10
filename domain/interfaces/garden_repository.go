package interfaces

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/types"
)

// GardenRepository ...
type GardenRepository interface {
	New(ctx context.Context, garden *entity.Garden) (*entity.Garden, error)
	Get(ctx context.Context, id types.GardenID) (*entity.Garden, error)
	GetByName(ctx context.Context, name string) (*entity.Garden, error)
	List(ctx context.Context, offset uint32) ([]*entity.Garden, error)
}
