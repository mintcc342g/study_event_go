package interfaces

import (
	"context"
	"study_event_go/domain/entity"
	"study_event_go/types"
)

// LilyRepository ...
type LilyRepository interface {
	TopClassLilies(ctx context.Context, gardenID types.GardenID, memberCnt int) ([]*entity.Lily, error)
}
