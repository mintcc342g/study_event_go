package types

import "strings"

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

// UnmarshalJSON ...
func (h *HugeClass) UnmarshalJSON(data []byte) (err error) {
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

// UnmarshalJSON ...
func (h *HugeType) UnmarshalJSON(data []byte) (err error) {
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
