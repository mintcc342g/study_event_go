package interfaces

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/types"
)

// LilyRepository ...
type LilyRepository interface {
	TopClassLilies(ctx context.Context, gardenID types.GardenID, memberCnt int) ([]*entity.Lily, error)
}
