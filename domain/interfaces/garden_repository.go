package interfaces

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/types"
)

// GardenRepository ...
type GardenRepository interface {
	New(ctx context.Context, garden *entity.Garden) (*entity.Garden, error)
	Garden(ctx context.Context, id types.GardenID) (*entity.Garden, error)
	GardenByName(ctx context.Context, name string) (*entity.Garden, error)
	Gardens(ctx context.Context, offset uint32) ([]*entity.Garden, error)
	Update(ctx context.Context, garden *entity.Garden) (*entity.Garden, error)
}
