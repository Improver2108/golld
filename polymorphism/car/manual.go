package car

import "fmt"

// ManualCar demonstrates polymorphism through different method implementations
// This struct shows how the same interface can have completely different behavior
type ManualCar struct {
	*car            // Embedded car field - inherits base functionality
	currentGear int // Additional field for manual transmission state
}

// InitManual creates a manual car with polymorphic capabilities
// This constructor demonstrates how different types can be created through the same pattern
func InitManual(brand, model string) *ManualCar {
	return &ManualCar{car: initCar(brand, model)}
}

// ShiftGear adds manual-specific functionality not available in other car types
// This method demonstrates polymorphism by providing functionality only ManualCar has
func (c *ManualCar) ShiftGear(gear int) {
	if !c.isEngineOn {
		fmt.Println(c.brand + " " + c.model + " : Engine is off! Cannot Cannot Shift Gear.")
		return
	}
	c.currentGear = gear
	fmt.Println(c.brand+" "+c.model+" : Shifted to gear ", c.currentGear)
}

// Accelerate implements the Car interface with manual-specific logic
// This method demonstrates polymorphism by providing different behavior for the same signature
func (s *ManualCar) Accelerate() {
	if !s.isEngineOn {
		fmt.Println(s.brand + " " + s.model + " : Cannot accelerate! Engine is off.")
		return
	}
	s.currentSpeed += 20
	fmt.Println(s.brand+" "+s.model+" : Accelerating to ", s.currentSpeed, " km/h")
}

// AccelerateBy implements the Car interface with manual-specific logic
// This method demonstrates polymorphism by providing different behavior for the same signature
func (s *ManualCar) AccelerateBy(speed int) {
	if !s.isEngineOn {
		fmt.Println(s.brand + " " + s.model + " : Cannot accelerate! Engine is off.")
		return
	}
	s.currentSpeed += speed
	fmt.Println(s.brand+" "+s.model+" : Accelerating to ", s.currentSpeed, " km/h")
}

// Brake implements the Car interface with manual-specific braking behavior
// This method demonstrates polymorphism by providing different braking characteristics
func (s *ManualCar) Brake() {
	s.currentSpeed -= 20
	if s.currentSpeed < 0 {
		s.currentSpeed = 0
	}
	fmt.Println(s.brand+" "+s.model+" : Braking! Speed is now ", s.currentSpeed, " km/h")
}
