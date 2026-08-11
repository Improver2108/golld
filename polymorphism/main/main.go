package main

// Package main demonstrates polymorphism through interfaces in Go
// This file shows how different implementations can be used interchangeably

import (
	"fmt"

	"github.com/improver2108/golld/polymorphism/car"
)

// Car defines the polymorphic contract that all car types must satisfy
// This interface demonstrates how different implementations can share the same behavior signature
type Car interface {
	StartEngine()
	Accelerate()
	AccelerateBy(speed int)
	Brake()
	StopEngine()
	GetSpeed() int
}

// useCar demonstrates polymorphic behavior with type-specific operations
// This function shows how to work with different implementations through type assertions
func useCar(c Car) {
	c.StartEngine()
	if g, ok := c.(*car.ManualCar); ok {
		g.ShiftGear(3) // Manual-specific functionality
	}
	if e, ok := c.(*car.ElectricCar); ok {
		e.ChargeBattery() // Electric-specific functionality
	}
}

// testDrive demonstrates polymorphic behavior through the common interface
// This function works with any Car implementation without knowing the specific type
// It shows how polymorphism enables writing generic code that works with multiple implementations
func testDrive(c Car) {
	c.StartEngine()
	c.Accelerate()
	c.Accelerate()
	useCar(c)
	c.Brake()
	fmt.Println(c.GetSpeed())
	c.StopEngine()
}

// main function demonstrates the complete polymorphism workflow
// 1. Create different car types using their specific constructors
// 2. Use them through the common Car interface
// 3. Show how each type provides its own implementation of the same methods
// 4. Demonstrate that the same code works with completely different behaviors
func main() {
	manual := car.InitManual("Ford", "Mustang")
	testDrive(manual)
	fmt.Println("-------")
	electric := car.InitElectric("Tesla", "Model S")
	testDrive(electric)
}
