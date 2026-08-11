package main

// Package main demonstrates encapsulation in Go
// This file shows how to use controlled access patterns and interfaces

import (
	"fmt"

	"github.com/improver2108/golld/encapsulation/car"
)

// Vehicle defines the contract for any vehicle implementation
// This interface demonstrates encapsulation by:
// 1. Specifying exactly what operations are allowed
// 2. Hiding implementation details behind method signatures
// 3. Providing a clean, consistent API for all vehicle types
type Vehicle interface {
	StartEngine()
	ShiftGear(gear int)
	Accelerate()
	Brake()
	StopEngine()
	GetSpeed() int
}

// testDrive demonstrates how encapsulation enables reusable code
// This function works with any Vehicle implementation through the interface
// It shows how encapsulated methods provide all necessary operations without exposing internals
func testDrive(v Vehicle) {
	v.StartEngine()
	v.ShiftGear(1)
	v.Accelerate()
	v.ShiftGear(2)
	v.Accelerate()
	v.Brake()
	fmt.Println(v.GetSpeed()) // Using the encapsulated getter method
	v.StopEngine()
}

// main function demonstrates the complete encapsulation workflow
// 1. Create an instance using the encapsulated constructor (Init)
// 2. Use it through the Vehicle interface
// 3. Show how all operations are performed through controlled methods
func main() {
	myCar := car.Init("Ford", "Mustang")

	// All interactions go through the encapsulated interface methods
	// External code never directly accesses internal fields or state
	testDrive(myCar)
}
