package interfaces

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/types"
)

// LilyRepository ...
type LilyRepository interface {
	Save(ctx context.Context, lily *entity.Lily) (*entity.Lily, error)
	Lily(ctx context.Context, id types.LilyID) (*entity.Lily, error)
	LiliesByRank(ctx context.Context, gardenID types.GardenID, rank uint32) ([]*entity.Lily, error)
}
