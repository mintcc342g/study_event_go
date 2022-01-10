package types

import (
	"encoding/json"
	"strings"
)

// HugeClass ...
type HugeClass uint32

// Huge's Classes ...
const (
	NoneClass HugeClass = iota
	SmallClass
	MiddleClass
	LargeClass
	GiganticClass
	UltraClass
)

// String ...
func (h HugeClass) String() string {
	return [...]string{"none", "small", "middle", "large", "gigantic", "ultra"}[h]
}

// MarshalJSON ...
func (h *HugeClass) MarshalJSON() ([]byte, error) {
	return json.Marshal((*h).String())
}

// UnmarshalJSON ...
func (h *HugeClass) UnmarshalJSON(data []byte) error {
	strData := strings.Trim(string(data), "\"")
	if strData == "" {
		return nil
	}

	*h = map[string]HugeClass{
		"none":     NoneClass,
		"small":    SmallClass,
		"middle":   MiddleClass,
		"large":    LargeClass,
		"gigantic": GiganticClass,
		"ultra":    UltraClass,
	}[strings.ToLower(strData)]

	return nil
}

// HugeType ...
type HugeType uint32

// Huge's Types ...
const (
	NoneType HugeType = iota
	SpecialType
)

// String ...
func (h HugeType) String() string {
	return [...]string{"none", "special"}[h]
}

// MarshalJSON ...
func (h *HugeType) MarshalJSON() ([]byte, error) {
	return json.Marshal((*h).String())
}

// UnmarshalJSON ...
func (h *HugeType) UnmarshalJSON(data []byte) error {
	strData := strings.Trim(string(data), "\"")
	if strData == "" {
		return nil
	}

	*h = map[string]HugeType{
		"none":    NoneType,
		"special": SpecialType,
	}[strings.ToLower(strData)]

	return nil
}

// AlertLevel ...
type AlertLevel uint32

// Alert Levels ...
const (
	LevelZero AlertLevel = iota
	LevelOne
	LevelTwo
	LevelThree
)

// SkillType ...
type SkillType uint32

// Skill Types ...
const (
	NoneSkillType SkillType = iota
	SkillTypeRare
	SkillTypeSub
	SkillTypeBoosted
)

// String ...
func (s SkillType) String() string {
	return [...]string{"none", "rare", "sub", "boosted"}[s]
}

// MarshalJSON ...
func (s *SkillType) MarshalJSON() ([]byte, error) {
	return json.Marshal((*s).String())
}

// UnmarshalJSON ...
func (s *SkillType) UnmarshalJSON(data []byte) error {
	strData := strings.Trim(string(data), "\"")
	if strData == "" {
		return nil
	}

	*s = map[string]SkillType{
		"none":    NoneSkillType,
		"rare":    SkillTypeRare,
		"sub":     SkillTypeSub,
		"boosted": SkillTypeBoosted,
	}[strings.ToLower(strData)]

	return nil
}

// DeletionReason ...
type DeletionReason uint32

// DeletionReasons ...
const (
	NoneReason DeletionReason = iota
	Retirement
	NonHostileDeath
	KilledInAction
	MissingInAction
)

// String ...
func (d DeletionReason) String() string {
	return [...]string{"NONE", "RETIREMENT", "NON-HOSTILE DEATH", "KILLED IN ACTION", "MISSING IN ACTION"}[d]
}

// MarshalJSON ...
func (d *DeletionReason) MarshalJSON() ([]byte, error) {
	return json.Marshal((*d).String())
}

// UnmarshalJSON ...
func (d *DeletionReason) UnmarshalJSON(data []byte) error {
	strData := strings.Trim(string(data), "\"")
	if strData == "" {
		return nil
	}

	*d = map[string]DeletionReason{
		"NONE":              NoneReason,
		"RETIREMENT":        Retirement,
		"NON-HOSTILE DEATH": NonHostileDeath,
		"KILLED IN ACTION":  KilledInAction,
		"MISSING IN ACTION": MissingInAction,
	}[strings.ToLower(strData)]

	return nil
}
