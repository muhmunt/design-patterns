package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	STANDART_CAR = 1
	SPORT_CAR = 2
)

type CarFactory struct{}

func (f *CarFactory) NewVehicle(carType int) (Vehicle, error) {
	switch carType {
	case STANDART_CAR:
		return new(Standartcar), nil
	case SPORT_CAR:
		return new(Sportcar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Unsupported var vehicle type : %d", carType))
	}
}