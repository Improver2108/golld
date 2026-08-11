package car

// Package car demonstrates inheritance through composition in Go
// Since Go doesn't have traditional class inheritance, we use embedding to achieve similar results

import "fmt"

// car represents the base vehicle with common functionality
// This struct serves as our "base class" that other car types will inherit from
type car struct {
	brand        string // Brand name (e.g., "Ford", "Tesla")
	model        string // Model name (e.g., "Mustang", "Model S")
	isEngineOn   bool   // Engine state
	currentSpeed int    // Current speed in km/h
}

// Init creates a new base car instance with initialization
// This serves as our constructor for the base "class"
func Init(brand, model string) *car {
	return &car{brand: brand, model: model}
}

// StartEngine turns on the engine - inherited by all car types
// This method will be available to all cars that embed this struct
func (s *car) StartEngine() {
	s.isEngineOn = true
	fmt.Println(s.brand + " " + s.model + " : Engine starts with a roar!")
}

// Accelerate increases speed - can be overridden by subclasses
// This method demonstrates how base functionality can be inherited and optionally overridden
func (s *car) Accelerate() {
	if !s.isEngineOn {
		fmt.Println(s.brand + " " + s.model + " : Engine is off! Cannot accelerate.")
	}
	s.currentSpeed += 20
	fmt.Println(s.brand+" "+s.model+" : Accelerating to ", s.currentSpeed, " km/h")
}

// Brake reduces speed - inherited by all car types as-is
// This method shows how common functionality is shared across different car implementations
func (s *car) Brake() {
	s.currentSpeed -= 20
	if s.currentSpeed < 0 {
		s.currentSpeed = 0
	}
	fmt.Println(s.brand+" "+s.model+" : Braking! Speed is now ", s.currentSpeed, " km/h")
}

// StopEngine turns off engine and resets state - inherited functionality
// This method demonstrates how cleanup operations are shared across car types
func (s *car) StopEngine() {
	s.isEngineOn = false
	s.currentSpeed = 0
	fmt.Println(s.brand + " " + s.model + " : Engine turned off.")
}

// GetSpeed provides read access to speed - inherited by all car types
// This getter method is available to any car that embeds this struct
func (s *car) GetSpeed() int {
	return s.currentSpeed
}
