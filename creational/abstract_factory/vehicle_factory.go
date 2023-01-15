package abstract_factory

type VehicleFactory interface {
	NewVehicle(int) (Vehicle, error)
}