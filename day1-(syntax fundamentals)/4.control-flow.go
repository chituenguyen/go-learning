package main

import "fmt"

func calculateScore() int {
	return 85
}

// If statements (NO parentheses!)
func checkAge(age int) {
	// TypeScript: if (age >= 18)
	// Go: NO parentheses
	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 13 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Child")
	}

	// If với initialization statement (unique to Go!)
	if score := calculateScore(); score > 80 {
		fmt.Println("Pass")
	}
}

// Loops (ONLY for, NO while!)
func loops() {
	// C-style for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While-style (just use for)
	count := 0
	for count < 5 { // Giống while
		fmt.Println(count)
		count++
	}

	// Infinite loop
	// for {
	//     fmt.Println("Forever")
	// }

	// Range (giống forEach)
	numbers := []int{1, 2, 3, 4, 5}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Chỉ cần value
	for _, value := range numbers {
		fmt.Println(value)
	}
}

// Switch (cleaner than JS, no break needed!)
func checkDay(day int) {
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 6, 7: // Multiple values
		fmt.Println("Weekend!")
	default:
		fmt.Println("Other day")
	}
	// NO NEED for break!

	// Switch without expression (giống if-else chain)
	score := 85
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}
}
