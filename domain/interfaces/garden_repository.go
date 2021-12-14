package interfaces

import (
	"context"
	"study_event_go/domain/entity"
	"study_event_go/types"
)

// GardenRepository ...
type GardenRepository interface {
	Garden(ctx context.Context, id types.GardenID) (*entity.Garden, error)
}
