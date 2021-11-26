package application

import (
	"study_event_go/domain/interfaces"
)

// EventService ...
type EventService struct {
	reidsRepo interfaces.RedisRepository
}

// NewEventService ...
func NewEventService(reidsRepo interfaces.RedisRepository) *EventService {
	return &EventService{
		reidsRepo: reidsRepo,
	}
}
