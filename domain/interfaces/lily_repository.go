package interfaces

import (
	"context"
	"study_event_go/domain/entity"
)

// LilyRepository ...
type LilyRepository interface {
	TopClassLilies(ctx context.Context, gardenID uint64, memberCnt int) ([]*entity.Lily, error)
}
