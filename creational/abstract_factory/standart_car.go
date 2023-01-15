package abstract_factory

type Standartcar struct {
	wheels int
	doors int
	speed int
	hasElectricEngine bool
}

func (standart Standartcar) WheelCount() int {
	return standart.wheels
}

func (standart Standartcar) NumberOfDoors() int {
	return standart.doors
}

func (standart Standartcar) Speed() int {
	return standart.speed
}

func (standart Standartcar) HasElectricEngine() bool {
	return standart.hasElectricEngine
}