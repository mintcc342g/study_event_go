package entity

import (
	"strings"
	"study-event-go/application/dto"
	"study-event-go/domain/vo"
	"study-event-go/types"

	"github.com/juju/errors"
)

// Alarm ...
type Alarm struct {
	ID                types.AlarmID
	GardenID          types.GardenID
	CaveLocation      string
	Huges             []*vo.Huge
	TotalHugeCount    uint32
	AlertLevel        types.AlertLevel
	LegionMemberCount uint32
}

// NewAlarm ...
func NewAlarm(garden *Garden, alarmDTO *dto.Alarm) (*Alarm, error) {
	if err := validateAlarmDTO(alarmDTO); err != nil {
		return nil, err
	}

	if garden.IsTopLegionSystem() && alarmDTO.LegionMemberCount == 0 {
		return nil, errors.BadRequestf("invalid legion member count")
	}

	alarm := &Alarm{
		GardenID:          alarmDTO.GardenID,
		CaveLocation:      alarmDTO.CaveLocation,
		Huges:             []*vo.Huge{},
		TotalHugeCount:    alarmDTO.TotalHugeCount,
		AlertLevel:        alarmDTO.AlertLevel,
		LegionMemberCount: alarmDTO.LegionMemberCount,
	}

	for _, huge := range alarmDTO.Huges {
		alarm.Huges = append(alarm.Huges, &vo.Huge{
			Class: huge.Class,
			Type:  huge.Type,
		})
	}

	return alarm, nil
}

func validateAlarmDTO(alarmDTO *dto.Alarm) error {
	if alarmDTO.GardenID == 0 {
		return errors.BadRequestf("invalid garden ID")
	}

	alarmDTO.CaveLocation = strings.TrimSpace(alarmDTO.CaveLocation)

	if alarmDTO.CaveLocation == "" {
		return errors.BadRequestf("invalid cave location")
	}

	if alarmDTO.TotalHugeCount == 0 {
		return errors.BadRequestf("invalid total huge count")
	}

	if len(alarmDTO.Huges) <= 0 {
		return errors.BadRequestf("invalid huges")
	}

	return nil
}

// IsSevere ...
func (a *Alarm) IsSevere() bool {
	return a.AlertLevel == types.LevelThree
}
