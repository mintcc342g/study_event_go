package worker

import (
	"context"
	"fmt"
	"study-event-go/cmd/container"
	"study-event-go/domain/interfaces"
	"study-event-go/types"
)

// EventWorker ...
type EventWorker struct {
	lilyRepo   interfaces.LilyRepository
	gardenRepo interfaces.GardenRepository
}

// NewEventWorker ...
func NewEventWorker(repoContainer *container.RepositoryContainer) *EventWorker {
	return &EventWorker{
		lilyRepo:   repoContainer.LilyRepo,
		gardenRepo: repoContainer.GardenRepo,
	}
}

// ProcessLegionSortieEvent ...
func (w *EventWorker) ProcessLegionSortieEvent(gardenID uint64, location string, legionMemberCount uint32) error {
	ctx := context.Background()
	// ctx = context.WithValue(ctx, types.RequestIDKey, requestID) 	// TODO

	garden, err := w.gardenRepo.Garden(ctx, types.GardenID(gardenID))
	if err != nil {
		return err
	}

	lilies, err := w.lilyRepo.LiliesByOrderedRank(ctx, garden.ID, legionMemberCount)
	if err != nil {
		return err
	}

	legion, err := garden.NewTopLegion(lilies)
	if err != nil {
		return err
	}

	println("[ANNOUNCEMENT]", "LEGION IS CREATED!!!")
	println("[ANNOUNCEMENT]", "THE MEMBER IS ...")
	for i, l := range legion.Members {
		println("[ANNOUNCEMENT]", fmt.Sprintf("%d)", i), l.Name.FullName())
	}
	return nil
}
