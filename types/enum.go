package types

// HugeClass ...
type HugeClass uint32

// Huge's Classes ...
const (
	NoneClass HugeClass = iota
	SmallClass
	MiddleClass
	LargeClass
	UltraClass
	SpecialClass
)

// AlertLevel ...
type AlertLevel uint32

// Alert Levels ...
const (
	LevelZero AlertLevel = iota
	LevelOne
	LevelTwo
	LevelThree
)
