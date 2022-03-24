package repository

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/types"

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
				Name:  "gardenID",
				Type:  "uint64",
				Value: sortieEvent.GardenID,
			},
			{
				Name:  "location",
				Type:  "string",
				Value: sortieEvent.Location,
			},
			{
				Name:  "legionMemberCount",
				Type:  "uint32",
				Value: sortieEvent.LegionMemberCount,
			},
			// TODO
			// {
			// 	Name:  "requestID",
			// 	Type:  "string",
			// 	Value: util.GetRequestID(ctx),
			// },
		},
	}

	if _, err := m.server.SendTask(signature); err != nil {
		return err
	}

	return nil
}
