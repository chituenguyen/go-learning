package main

import "fmt"

// Basic function
// TypeScript: function add(a: number, b: number): number
func add(a int, b int) int {
	return a + b
}

// Short form khi cùng type
func multiply(a, b int) int {
	return a * b
}

// Multiple return values (KHÁC TypeScript!)
// Very common in Go for error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil // nil = no error
}

func useMultipleReturn() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)

	// Ignore return value với _
	result2, _ := divide(10, 2) // Ignore error
	fmt.Println("Result2:", result2)
}

// Named return values
func getUser(id int) (name string, age int, err error) {
	if id <= 0 {
		err = fmt.Errorf("invalid id")
		return // Return zero values + error
	}
	name = "John"
	age = 30
	return // Automatically returns name, age, err
}

// Variadic functions (như ...args trong JS)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func testVariadic() {
	fmt.Println(sum(1, 2, 3))       // 6
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15
}
