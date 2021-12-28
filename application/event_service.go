package application

import (
	"study-event-go/domain/interfaces"
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
