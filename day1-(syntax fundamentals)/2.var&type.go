package main

import "fmt"

func variables() {
	// TypeScript: const name: string = "Tue"
	// Go có 3 cách khai báo:

	// Cách 1: var (giống let/const)
	var name string = "Tue"
	var age int = 28
	var isActive bool = true

	// Cách 2: := (short declaration - dùng nhiều nhất)
	// Go tự infer type
	name2 := "John"   // string
	age2 := 30        // int
	salary := 5000.50 // float64

	// Cách 3: Declare nhiều
	var (
		username = "admin"
		password = "secret"
	)

	// Zero values (khác TypeScript undefined)
	var count int    // 0
	var text string  // ""
	var enabled bool // false
	var ptr *int     // nil

	fmt.Println(name, age, isActive)
}

func types() {
	// Basic types
	var i int = 42 // int32 or int64 (depends on system)
	var f float64 = 3.14
	var b bool = true
	var s string = "hello"

	// Type conversion (NO implicit conversion!)
	// TypeScript: let x: number = 10; let y: string = x; // OK
	// Go: PHẢI explicit
	var x int = 10
	var y float64 = float64(x) // Must convert

	// Constants
	const Pi = 3.14
	const AppName = "MyApp"
}
