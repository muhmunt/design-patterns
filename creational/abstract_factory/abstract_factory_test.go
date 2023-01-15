package abstract_factory

import (
	"fmt"
	"testing"
)

func TestCarFactory(t *testing.T) {
	CarFactory, err := CreateVehicleFactory(CAR)

	if err != nil {
		t.Fatal(err)
	}

	sportCarVehicle, err := CarFactory.NewVehicle(SPORT_CAR)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sportCarVehicle.NumberOfDoors())
	_, ok := sportCarVehicle.(Car)

	if !ok {
		t.Fatal("Struct wrong failed")
	}
}