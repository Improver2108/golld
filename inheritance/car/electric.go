package car

import "fmt"

// ElectricCar demonstrates inheritance through composition
// This struct embeds the base car and adds electric-specific functionality
type ElectricCar struct {
	*car             // Embedded car field - this gives ElectricCar all car methods
	batteryLevel int // Additional field specific to electric cars
}

// InitElectric creates an electric car by composing with base car functionality
// This constructor demonstrates how inheritance is achieved through embedding
func InitElectric(brand, model string) *ElectricCar {
	return &ElectricCar{car: Init(brand, model), batteryLevel: 100}
}

// ChargeBattery adds electric-specific functionality
// This method demonstrates how derived classes can add new methods
// while inheriting all base class functionality through embedding
func (c *ElectricCar) ChargeBattery() {
	c.batteryLevel = 100
	fmt.Println(c.brand + " " + c.model + " : Battery fully charged!")
}
