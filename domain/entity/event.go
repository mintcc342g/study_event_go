package entity

// SortieEvent ...
type SortieEvent struct {
	GardenID          uint64
	Location          string
	LegionMemberCount uint32
}

// NewSortieEvent ...
func NewSortieEvent(alarm *Alarm) *SortieEvent {
	return &SortieEvent{
		GardenID:          uint64(alarm.GardenID),
		Location:          alarm.CaveLocation,
		LegionMemberCount: uint32(alarm.LegionMemberCount),
	}
}
