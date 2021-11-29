package entity

import (
	"study_event_go/application/dto"
	"study_event_go/domain/vo"
	"study_event_go/types"
)

// Alarm ...
type Alarm struct {
	id           uint64
	caveLocation string
	huges        []*vo.Huge
	totalCount   uint32
	alertLevel   types.AlertLevel
}

// NewAlarm ...
func NewAlarm(alarmDTO *dto.Alarm) *Alarm {
	alarm := &Alarm{
		caveLocation: alarmDTO.CaveLocation,
		totalCount:   alarmDTO.TotalCount,
		alertLevel:   alarmDTO.AlertLevel,
		huges:        []*vo.Huge{},
	}

	for _, huge := range alarmDTO.Huges {
		alarm.huges = append(alarm.huges, &vo.Huge{
			Class: huge.Class,
			Type:  huge.Type,
		})
	}

	return alarm
}

// IsSevere ...
func (w *Alarm) IsSevere() bool {
	return w.alertLevel == types.LevelThree
}
