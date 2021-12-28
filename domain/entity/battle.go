package entity

import (
	"math/rand"
	"study-event-go/application/dto"
	"study-event-go/domain/vo"
	"study-event-go/types"
)

// Alarm ...
type Alarm struct {
	ID           types.AlarmID
	GardenID     types.GardenID
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
