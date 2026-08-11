package main

// Package main demonstrates inheritance through composition in Go
// This file shows how embedding achieves similar results to traditional class inheritance

import (
	"fmt"

	"github.com/improver2108/golld/inheritance/car"
)

// Car defines the common interface for all car types
// This interface demonstrates how different implementations can share a contract
type Car interface {
	StartEngine()
	Accelerate()
	Brake()
	StopEngine()
	GetSpeed() int
}

// useCar demonstrates type-specific functionality while using the common interface
// This function shows how to work with different car types through type assertion
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
func testDrive(c Car) {
	c.StartEngine()
	c.Accelerate()
	c.Accelerate()
	useCar(c)
	c.Brake()
	fmt.Println(c.GetSpeed())
	c.StopEngine()
}

// main function demonstrates the complete inheritance workflow
// 1. Create different car types using their specific constructors
// 2. Use them through the common Car interface
// 3. Show how each type inherits base functionality while adding its own features
func main() {
	manual := car.InitManual("Ford", "Mustang")
	testDrive(manual)
	electric := car.InitElectric("Tesla", "Model S")
	testDrive(electric)
}
