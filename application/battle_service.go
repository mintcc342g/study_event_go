package application

import (
	"context"
	"study_event_go/application/dto"
	"study_event_go/domain/interfaces"
)

// BattleService ...
type BattleService struct {
	studentRepo interfaces.StudentRepository
}

// NewBattleService ...
func NewBattleService(studentRepo interfaces.StudentRepository) *BattleService {
	return &BattleService{
		studentRepo: studentRepo,
	}
}

// Warning ...
func (b *BattleService) Warning(ctx context.Context, warningDTO *dto.Warning) (err error) {
	return nil
}

// // Sortie ...
// func (b *BattleService) Sortie(ctx context.Context, warningDTO *dto.Warning) (err error) {
// 	return nil
// }
