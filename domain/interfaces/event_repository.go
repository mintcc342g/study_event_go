package interfaces

import (
	"context"
	"study_event_go/domain/entity"
)

// EventRepository ...
type EventRepository interface {
	SendLegionSortieEvent(ctx context.Context, sortieEvent *entity.SortieEvent) error
}
