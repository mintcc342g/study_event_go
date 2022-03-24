package application

import (
	"context"
	"study-event-go/application/dto"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
)

// ProtectionService ...
type ProtectionService struct {
	lilyRepo   interfaces.LilyRepository
	gardenRepo interfaces.GardenRepository
	eventRepo  interfaces.EventRepository
}

// NewProtectionService ...
func NewProtectionService(
	lilyRepo interfaces.LilyRepository,
	gardenRepo interfaces.GardenRepository,
	eventRepo interfaces.EventRepository,
) *ProtectionService {
	return &ProtectionService{
		lilyRepo:   lilyRepo,
		gardenRepo: gardenRepo,
		eventRepo:  eventRepo,
	}
}

// Alarm ...
func (p *ProtectionService) Alarm(ctx context.Context, alarmDTO *dto.Alarm) (err error) {

	// TODO: requestor check with ctx
	// TODO: logger

	garden, err := p.gardenRepo.Garden(ctx, alarmDTO.GardenID)
	if err != nil {
		return err
	}

	alarm, err := entity.NewAlarm(garden, alarmDTO)
	if err != nil {
		return err
	}

	if alarm.IsSevere() && garden.IsTopLegionSystem() {
		return p.makeTopLegion(ctx, garden, alarm)
	}

	return p.startNormalBattle(ctx, garden, alarm)
}

func (p *ProtectionService) makeTopLegion(ctx context.Context, garden *entity.Garden, alarm *entity.Alarm) (err error) {

	lilies, err := p.lilyRepo.LiliesByRank(ctx, alarm.GardenID, alarm.LegionMemberCount)
	if err != nil {
		return err
	}

	legion, err := garden.NewTopLegion(lilies)
	if err != nil {
		return err
	}

	event := entity.NewSortieEvent(alarm, legion)

	if err = p.eventRepo.SendLegionSortieEvent(ctx, event); err != nil {
		return err
	}

	return nil
}

// startNormalBattle
func (p *ProtectionService) startNormalBattle(ctx context.Context, garden *entity.Garden, alarm *entity.Alarm) (err error) {

	// TODO: send push notification to all lilies in the garden

	return nil
}
