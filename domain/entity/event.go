package entity

import "study-event-go/domain/vo"

// SortieEvent ...
type SortieEvent struct {
	Location string
	Legion   *Legion
	Huges    []*vo.Huge
}

// NewSortieEvent ...
func NewSortieEvent(alarm *Alarm, legion *Legion) *SortieEvent {
	return &SortieEvent{
		Location: alarm.CaveLocation,
		Legion:   legion,
		Huges:    alarm.Huges,
	}
}
