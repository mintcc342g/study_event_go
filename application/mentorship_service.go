package application

import (
	"context"
	"study_event_go/application/dto"
	"study_event_go/domain/entity"
	"study_event_go/domain/interfaces"
	"study_event_go/types"
)

// MentorshipService ...
type MentorshipService struct {
	mentorshipRepo interfaces.MentorshipRepository
}

// NewMentorshipService ...
func NewMentorshipService(
	mentorshipRepo interfaces.MentorshipRepository,
) *MentorshipService {
	return &MentorshipService{
		mentorshipRepo: mentorshipRepo,
	}
}

// TODO: logger

// New ...
func (m *MentorshipService) New(ctx context.Context, mentorshipDTO *dto.Mentorship) (*dto.Mentorship, error) {

	mentorship, err := m.mentorshipRepo.New(ctx, entity.NewMentorship(mentorshipDTO))
	if err != nil {
		return nil, err
	}

	return mentorship.DTO(), nil
}

// Get ...
func (m *MentorshipService) Get(ctx context.Context, id types.MentorshipSystemID) (*dto.Mentorship, error) {

	mentorship, err := m.mentorshipRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return mentorship.DTO(), nil
}
