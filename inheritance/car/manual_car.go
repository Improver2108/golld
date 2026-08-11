package car

import "fmt"

// ManualCar demonstrates inheritance and method overriding through composition
// This struct embeds the base car and adds manual transmission functionality
type ManualCar struct {
	*car            // Embedded car field - inherits all base car methods
	currentGear int // Additional field for manual transmission state
}

// InitManual creates a manual car by composing with base car functionality
// This constructor demonstrates inheritance through embedding
func InitManual(brand, model string) *ManualCar {
	return &ManualCar{car: Init(brand, model)}
}

// ShiftGear adds manual transmission functionality specific to ManualCar
// This method demonstrates how derived classes can add new methods
func (c *ManualCar) ShiftGear(gear int) {
	if !c.isEngineOn {
		fmt.Println(c.brand + " " + c.model + " : Engine is off! Cannot Cannot Shift Gear.")
		return
	}
	c.currentGear = gear
	fmt.Println(c.brand+" "+c.model+" : Shifted to gear ", c.currentGear)
}

// Accelerate overrides the base class method with custom behavior
// This demonstrates method overriding in Go's composition model
// Note: This is a simplified override - in real code you'd typically call the parent method
func (s *ManualCar) Accelerate() {
	fmt.Println("this is manual car")
}
