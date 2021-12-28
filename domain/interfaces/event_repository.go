package interfaces

import (
	"context"
	"study-event-go/domain/entity"
)

// EventRepository ...
type EventRepository interface {
	SendLegionSortieEvent(ctx context.Context, sortieEvent *entity.SortieEvent) error
}
