package repository

import (
	"context"
	"study_event_go/domain/entity"
	"study_event_go/domain/interfaces"
	"study_event_go/types"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
)

type machineryRepository struct {
	server *machinery.Server
}

// NewMachineryRepository ...
func NewMachineryRepository(server *machinery.Server) interfaces.EventRepository {
	return &machineryRepository{
		server: server,
	}
}

func (m *machineryRepository) SendLegionSortieEvent(ctx context.Context, sortieEvent *entity.SortieEvent) error {

	signature := &tasks.Signature{
		Name: types.LegionSortieEvent,
		Args: []tasks.Arg{
			{
				Name:  "location",
				Type:  "string",
				Value: sortieEvent.Location,
			},
			{
				Name:  "legion",
				Type:  "interface{}",
				Value: sortieEvent.Legion,
			},
			{
				Name:  "huges",
				Type:  "[]interface{}",
				Value: sortieEvent.Huges,
			},
		},
	}

	if _, err := m.server.SendTask(signature); err != nil {
		return err
	}

	return nil
}
