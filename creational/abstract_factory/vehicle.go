package abstract_factory

type Vehicle interface {
	WheelCount() int
	NumberOfDoors() int
	Speed() int
}