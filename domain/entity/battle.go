package entity

import (
	"math/rand"
	"study_event_go/application/dto"
	"study_event_go/domain/vo"
	"study_event_go/types"
)

// Alarm ...
type Alarm struct {
	ID           uint64
	GardenID     uint64
	CaveLocation string
	Huges        []*vo.Huge
	TotalCount   uint32
	AlertLevel   types.AlertLevel
}

// NewAlarm ...
func NewAlarm(alarmDTO *dto.Alarm) *Alarm {

	alarm := &Alarm{
		GardenID:     alarmDTO.GardenID,
		CaveLocation: alarmDTO.CaveLocation,
		TotalCount:   alarmDTO.TotalCount,
		AlertLevel:   alarmDTO.AlertLevel,
		Huges:        []*vo.Huge{},
	}

	for _, huge := range alarmDTO.Huges {
		alarm.Huges = append(alarm.Huges, &vo.Huge{
			Class: huge.Class,
			Type:  huge.Type,
		})
	}

	return alarm
}

// IsSevere ...
func (a *Alarm) IsSevere() bool {
	return a.AlertLevel == types.LevelThree
}

// MakeLegionMemberCount ...
func (a *Alarm) MakeLegionMemberCount() int {
	return rand.Intn(types.TempleLegionMaxNumber-types.TempleLegionMinNumber) + types.TempleLegionMinNumber
}
