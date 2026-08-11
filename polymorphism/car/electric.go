package car

import "fmt"

// ElectricCar demonstrates polymorphism through method implementation variation
// This struct shows how different implementations can satisfy the same interface differently
type ElectricCar struct {
	*car             // Embedded car field - inherits base functionality
	batteryLevel int // Additional field specific to electric cars
}

// InitElectric creates an electric car with polymorphic capabilities
// This constructor demonstrates how different types can be created through the same pattern
func InitElectric(brand, model string) *ElectricCar {
	return &ElectricCar{car: initCar(brand, model), batteryLevel: 100}
}

// ChargeBattery adds electric-specific functionality
// This method demonstrates polymorphism by providing functionality only available to ElectricCar
func (c *ElectricCar) ChargeBattery() {
	c.batteryLevel = 100
	fmt.Println(c.car.brand + " " + c.car.model + " : Battery fully charged!")
}

// Accelerate implements the Car interface with electric-specific logic
// This method demonstrates polymorphism by providing different behavior for the same method signature
func (c *ElectricCar) Accelerate() {
	if !c.isEngineOn {
		fmt.Println(c.brand + " " + c.model + " : Cannot accelerate! Engine is off.")
		return
	}
	if c.batteryLevel <= 0 {
		fmt.Println(c.brand + " " + c.model + " : Battery dead! Cannot accelerate.")
		return
	}

	fmt.Println(c.brand+" "+c.model+" : Accelerating to ", c.currentSpeed, " km/h")
}

// AccelerateBy implements the Car interface with electric-specific logic
// This method demonstrates polymorphism by providing different behavior for the same signature
func (c *ElectricCar) AccelerateBy(speed int) {
	if !c.isEngineOn {
		fmt.Println(c.brand + " " + c.model + " : Cannot accelerate! Engine is off.")
		return
	}
	if c.batteryLevel <= 0 {
		fmt.Println(c.brand + " " + c.model + " : Battery dead! Cannot accelerate.")
		return
	}
	c.currentSpeed += speed
	fmt.Println(c.brand+" "+c.model+" : Accelerating to ", c.currentSpeed, " km/h")
}

// Brake implements the Car interface with electric-specific braking behavior
// This method demonstrates polymorphism by providing different braking characteristics
func (c *ElectricCar) Brake() {
	c.currentSpeed -= 15
	if c.currentSpeed < 0 {
		c.currentSpeed = 0
	}
	fmt.Println(c.brand+" "+c.model+" : Braking! Speed is now ", c.currentSpeed, " km/h")
}
