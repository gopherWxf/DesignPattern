package Simple_Factory

import "testing"

func TestFactoryCreateCar(t *testing.T) {
	FactoryCreateCar("bmw").drive()
	FactoryCreateCar("tesla").drive()
	FactoryCreateCar("porsche").drive()
}
