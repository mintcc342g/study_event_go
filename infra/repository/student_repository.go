package repository

import (
	"study_event_go/domain/interfaces"
)

type studentORMRepository struct {
}

// NewStudentRepository ...
func NewStudentRepository() interfaces.StudentRepository {
	return &studentORMRepository{}
}
