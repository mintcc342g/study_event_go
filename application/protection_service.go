package application

import (
	"context"
	"study_event_go/application/dto"
	"study_event_go/domain/entity"
	"study_event_go/domain/interfaces"
)

// ProtectionService ...
type ProtectionService struct {
	studentRepo interfaces.StudentRepository
}

// NewBattleService ...
func NewBattleService(studentRepo interfaces.StudentRepository) *ProtectionService {
	return &ProtectionService{
		studentRepo: studentRepo,
	}
}

// Alarm ...
func (b *ProtectionService) Alarm(ctx context.Context, alarmDTO *dto.Alarm) (err error) {

	// TODO: requestor check with ctx

	alarm := entity.NewAlarm(alarmDTO)

	if alarm.IsSevere() {
		return b.sendTempleRegion(ctx, alarm)
	}

	return b.startNormalBattle(ctx, alarm)
}

// Sortie

// sendTempleRegion
func (b *ProtectionService) sendTempleRegion(ctx context.Context, alarm *entity.Alarm) (err error) {
	return nil
}

// startNormalBattle
func (b *ProtectionService) startNormalBattle(ctx context.Context, alarm *entity.Alarm) (err error) {
	return nil
}
