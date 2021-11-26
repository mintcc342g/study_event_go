package types

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

// HugeType ...
type HugeType uint32

// Huge's Types ...
const (
	NoneType HugeType = iota
	SpecialType
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
