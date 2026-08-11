package car

// Package car demonstrates abstraction by encapsulating sports car functionality
// This package shows how we can hide implementation details while exposing only necessary operations

import "fmt"

// SportsCar represents a sports car with encapsulated state and behavior
// The struct contains private fields that represent the internal state
type SportsCar struct {
	Brand        string // Brand of the sports car (e.g., "Ford", "Ferrari")
	Model        string // Model name of the sports car (e.g., "Mustang", "F40")
	IsEngineOn   bool   // Internal state: whether engine is running
	CurrentSpeed int    // Internal state: current speed in km/h
	CurrentGear  int    // Internal state: current gear (0 = neutral/parked)
}

// Init creates a new SportsCar instance with the given brand and model
// This is our constructor function that initializes the car's basic properties
// Note that engine starts off, speed is 0, and gear is neutral (0)
func Init(brand, model string) *SportsCar {
	return &SportsCar{Brand: brand, Model: model}
}

// StartEngine turns on the car's engine and sets initial state
// This method demonstrates abstraction by hiding how the engine actually starts
// The user only needs to call this method, not worry about internal implementation
func (s *SportsCar) StartEngine() {
	s.IsEngineOn = true
	fmt.Println(s.Brand + " " + s.Model + " : Engine starts with a roar!")
}

// ShiftGear changes the transmission gear with safety checks
// This method demonstrates abstraction by:
// 1. Hiding the internal logic for gear shifting
// 2. Enforcing business rules (can't shift if engine is off)
// The user only interacts through this simple interface
func (s *SportsCar) ShiftGear(gear int) {
	if !s.IsEngineOn {
		fmt.Println(s.Brand + " " + s.Model + " : Engine is off! Cannot Cannot Shift Gear.")
		return
	}
	s.CurrentGear = gear
	fmt.Println(s.Brand+" "+s.Model+" : Shifted to gear ", s.CurrentGear)
}

// Accelerate increases the car's speed with safety validation
// This method demonstrates abstraction by:
// 1. Hiding how acceleration is implemented internally
// 2. Providing appropriate error handling for invalid states
// The user gets a clean interface without seeing internal complexity
func (s *SportsCar) Accelerate() {
	if !s.IsEngineOn {
		fmt.Println(s.Brand + " " + s.Model + " : Engine is off! Cannot accelerate.")
	}
	s.CurrentSpeed += 20
	fmt.Println(s.Brand+" "+s.Model+" : Accelerating to ", s.CurrentSpeed, " km/h")
}

// Brake reduces the car's speed with boundary protection
// This method demonstrates abstraction by:
// 1. Encapsulating the braking logic
// 2. Handling edge cases (speed can't go below 0)
// The user gets predictable behavior without understanding implementation details
func (s *SportsCar) Brake() {
	s.CurrentSpeed -= 20
	if s.CurrentSpeed < 0 {
		s.CurrentSpeed = 0
	}
	fmt.Println(s.Brand+" "+s.Model+" : Braking! Speed is now ", s.CurrentSpeed, " km/h")
}

// StopEngine safely turns off the engine and resets state
// This method demonstrates abstraction by:
// 1. Handling all necessary cleanup in one operation
// 2. Ensuring proper state transitions (gear to neutral, speed to 0)
// The user gets a complete operation without managing individual steps
func (s *SportsCar) StopEngine() {
	s.IsEngineOn = false
	s.CurrentGear = 0
	s.CurrentSpeed = 0
	fmt.Println(s.Brand + " " + s.Model + " : Engine turned off.")
}
