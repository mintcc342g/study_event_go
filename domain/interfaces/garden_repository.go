package interfaces

import (
	"context"
	"study_event_go/domain/entity"
)

// GardenRepository ...
type GardenRepository interface {
	Garden(ctx context.Context, id uint64) (*entity.Garden, error)
}
