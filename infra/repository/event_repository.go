package repository

import (
	"context"
	"study_event_go/domain/entity"
	"study_event_go/domain/interfaces"
)

type machineryRepository struct {
}

// NewMachineryRepository ...
func NewMachineryRepository() interfaces.EventRepository {
	return &machineryRepository{}
}

func (m *machineryRepository) SendLegionSortieEvent(ctx context.Context, sortieEvent *entity.SortieEvent) error {
	return nil
}
