package application

import (
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
