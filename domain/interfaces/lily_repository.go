package interfaces

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/types"
)

// LilyRepository ...
type LilyRepository interface {
	New(ctx context.Context, lily *entity.Lily) (*entity.Lily, error)
	TopClassLilies(ctx context.Context, gardenID types.GardenID, memberCnt int) ([]*entity.Lily, error)
}
