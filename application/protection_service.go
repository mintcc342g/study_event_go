package application

import (
	"context"
	"study_event_go/application/dto"
	"study_event_go/domain/entity"
	"study_event_go/domain/interfaces"
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
func (b *ProtectionService) Alarm(ctx context.Context, alarmDTO *dto.Alarm) (err error) {

	// TODO: requestor check with ctx
	// TODO: logger
	// TODO: db

	garden, err := b.gardenRepo.Garden(ctx, alarmDTO.GardenID)
	if err != nil {
		return err
	}

	alarm := entity.NewAlarm(alarmDTO)

	if alarm.IsSevere() && garden.IsLudovico() {
		return b.sendTempleLegion(ctx, garden, alarm)
	}

	return b.startNormalBattle(ctx, garden, alarm)
}

// sendTempleLegion
func (b *ProtectionService) sendTempleLegion(ctx context.Context, garden *entity.Garden, alarm *entity.Alarm) (err error) {

	lilies, err := b.lilyRepo.TopClassLilies(ctx, alarm.GardenID, alarm.MakeLegionMemberCount())
	if err != nil {
		return err
	}

	legion, err := garden.NewTempleLegion(lilies)
	if err != nil {
		return err
	}

	event := entity.NewSortieEvent(alarm, legion)

	if err = b.eventRepo.SendLegionSortieEvent(ctx, event); err != nil {
		return err
	}

	return nil
}

// startNormalBattle
func (b *ProtectionService) startNormalBattle(ctx context.Context, garden *entity.Garden, alarm *entity.Alarm) (err error) {

	// TODO: send push notification to all lilies in the garden

	return nil
}
