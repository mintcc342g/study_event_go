package entity

import (
	"math/rand"
	"strings"
	"study-event-go/application/dto"
	"study-event-go/domain/vo"
	"study-event-go/types"

	"github.com/juju/errors"
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
func NewAlarm(alarmDTO *dto.Alarm) (*Alarm, error) {
	if err := validateAlarmDTO(alarmDTO); err != nil {
		return nil, err
	}

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

	if alarmDTO.TotalCount == 0 {
		return errors.BadRequestf("invalid total count")
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

// MakeLegionMemberCount ...
func (a *Alarm) MakeLegionMemberCount() int {
	return rand.Intn(types.TempleLegionMaxNumber-types.TempleLegionMinNumber) + types.TempleLegionMinNumber
}
