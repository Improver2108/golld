package car

// Package car demonstrates polymorphism through interfaces and method overriding
// This package shows how different implementations can share the same interface

import "fmt"

// car represents the base vehicle with common functionality
// This struct serves as our foundation for polymorphic behavior
type car struct {
	brand        string // Brand name (e.g., "Ford", "Tesla")
	model        string // Model name (e.g., "Mustang", "Model S")
	isEngineOn   bool   // Engine state
	currentSpeed int    // Current speed in km/h
}

// initCar creates a new base car instance with initialization
// This private constructor demonstrates encapsulation within the package
func initCar(brand, model string) *car {
	return &car{brand: brand, model: model}
}

// StartEngine turns on the engine - common to all car types
// This method demonstrates how base functionality can be shared across implementations
func (c *car) StartEngine() {
	c.isEngineOn = true
	fmt.Println(c.brand + " " + c.model + " : Engine starts with a roar!")
}

// StopEngine turns off engine and resets state - common functionality
// This method shows how cleanup operations are shared across different implementations
func (c *car) StopEngine() {
	c.isEngineOn = false
	c.currentSpeed = 0
	fmt.Println(c.brand + " " + c.model + " : Engine turned off.")
}

// GetSpeed provides read access to speed - common interface
// This getter method demonstrates how different implementations can provide the same interface
func (c *car) GetSpeed() int {
	return c.currentSpeed
}
