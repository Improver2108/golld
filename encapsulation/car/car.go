package car

// Package car demonstrates encapsulation by hiding internal implementation details
// This package shows how we can protect data integrity through controlled access

import "fmt"

// sportsCar represents a sports car with encapsulated internal state
// Note the lowercase field names - this is Go's way of indicating private fields
// In Go, there are no true private fields, but lowercase names indicate "private" by convention
type sportsCar struct {
	brand        string // Private: brand name (e.g., "Ford", "Ferrari")
	model        string // Private: model name (e.g., "Mustang", "F40")
	isEngineOn   bool   // Private: engine state
	currentSpeed int    // Private: current speed in km/h
	currentGear  int    // Private: current gear (0 = neutral/parked)
}

// Init creates a new sportsCar instance with proper initialization
// This is our constructor that sets up the initial state safely
// Note we return *sportsCar (lowercase) which follows Go's convention for unexported types
func Init(brand, model string) *sportsCar {
	return &sportsCar{brand: brand, model: model}
}

// StartEngine turns on the engine and provides feedback
// This method demonstrates encapsulation by:
// 1. Providing controlled access to modify internal state (isEngineOn)
// 2. Including appropriate feedback about the operation
func (s *sportsCar) StartEngine() {
	s.isEngineOn = true
	fmt.Println(s.brand + " " + s.model + " : Engine starts with a roar!")
}

// ShiftGear changes the transmission with validation logic
// This method demonstrates encapsulation by:
// 1. Protecting internal state (currentGear) through controlled access
// 2. Enforcing business rules (can't shift if engine is off)
// 3. Hiding the complexity of gear shifting from external code
func (s *sportsCar) ShiftGear(gear int) {
	if !s.isEngineOn {
		fmt.Println(s.brand + " " + s.model + " : Engine is off! Cannot Cannot Shift Gear.")
		return
	}
	s.currentGear = gear
	fmt.Println(s.brand+" "+s.model+" : Shifted to gear ", s.currentGear)
}

// Accelerate increases speed with safety checks
// This method demonstrates encapsulation by:
// 1. Providing controlled access to modify internal state (currentSpeed)
// 2. Including validation logic before modifying state
// 3. Returning appropriate feedback about the operation
func (s *sportsCar) Accelerate() {
	if !s.isEngineOn {
		fmt.Println(s.brand + " " + s.model + " : Engine is off! Cannot accelerate.")
	}
	s.currentSpeed += 20
	fmt.Println(s.brand+" "+s.model+" : Accelerating to ", s.currentSpeed, " km/h")
}

// Brake reduces speed with boundary protection
// This method demonstrates encapsulation by:
// 1. Providing controlled access to modify internal state (currentSpeed)
// 2. Including edge case handling (speed can't go below 0)
// 3. Returning feedback about the new state
func (s *sportsCar) Brake() {
	s.currentSpeed -= 20
	if s.currentSpeed < 0 {
		s.currentSpeed = 0
	}
	fmt.Println(s.brand+" "+s.model+" : Braking! Speed is now ", s.currentSpeed, " km/h")
}

// StopEngine safely turns off engine and resets state
// This method demonstrates encapsulation by:
// 1. Providing controlled access to modify multiple internal states at once
// 2. Ensuring proper cleanup (gear to neutral, speed to 0)
// 3. Returning feedback about the complete operation
func (s *sportsCar) StopEngine() {
	s.isEngineOn = false
	s.currentGear = 0
	s.currentSpeed = 0
	fmt.Println(s.brand + " " + s.model + " : Engine turned off.")
}

// GetSpeed provides read-only access to internal state
// This method demonstrates encapsulation by:
// 1. Providing controlled access to read private field (currentSpeed)
// 2. Allowing external code to get necessary information without exposing internals
// 3. Following the pattern of getter methods for encapsulated data
func (s *sportsCar) GetSpeed() int {
	return s.currentSpeed
}
