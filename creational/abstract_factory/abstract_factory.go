package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	CAR = 1
)

func CreateVehicleFactory(vfType int) (VehicleFactory, error) {
	switch vfType {
	case CAR:
		return new(CarFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Unrecognized factory type: %d", vfType))
	}
}