package interfaces

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/types"
)

// GardenRepository ...
type GardenRepository interface {
	Garden(ctx context.Context, id types.GardenID) (*entity.Garden, error)
}
