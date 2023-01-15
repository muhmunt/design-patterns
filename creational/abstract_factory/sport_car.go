package abstract_factory

type Sportcar struct {
	wheels int
	doors int
	speed int
	hasElectricEngine bool
}

func (standart Sportcar) WheelCount() int {
	return standart.wheels
}

func (standart Sportcar) NumberOfDoors() int {
	return standart.doors
}

func (standart Sportcar) Speed() int {
	return standart.speed
}

func (standart Sportcar) HasElectricEngine() bool {
	return standart.hasElectricEngine
}