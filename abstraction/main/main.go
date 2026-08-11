package main

// Package main demonstrates the abstraction concept in Go
// This file shows how to use interfaces to create abstract, reusable code

import "github.com/improver2108/golld/abstraction/car"

// Vehicle is an interface that defines the abstract contract for any vehicle
// This demonstrates abstraction by:
// 1. Hiding implementation details behind a clean interface
// 2. Allowing different vehicle types to be used interchangeably
// 3. Enforcing consistent behavior across all vehicle implementations
type Vehicle interface {
	StartEngine()
	ShiftGear(gear int)
	Accelerate()
	Brake()
	StopEngine()
}

// testDrive demonstrates polymorphism through abstraction
// This function works with ANY Vehicle implementation without knowing the specific type
// It shows how abstraction allows us to write generic code that works with multiple implementations
func testDrive(v Vehicle) {
	v.StartEngine()
	v.ShiftGear(1)
	v.Accelerate()
	v.ShiftGear(2)
	v.Accelerate()
	v.Brake()
	v.StopEngine()
}

// main function demonstrates the complete abstraction workflow
// 1. Create a concrete implementation (SportsCar)
// 2. Use it through an abstract interface (Vehicle)
// 3. Show how the same code works regardless of specific implementation
func main() {
	myCar := car.Init("Ford", "Mustang")

	// Here we pass a concrete SportsCar but treat it as the abstract Vehicle type
	// This is the essence of abstraction in object-oriented programming
	testDrive(myCar)
}
