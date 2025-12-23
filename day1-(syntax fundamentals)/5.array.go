package main

import "fmt"

func arraysAndSlices() {
	// ARRAY: Fixed size (ít dùng)
	// TypeScript: const arr: number[] = [1, 2, 3]
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr) // [1 2 3]

	// SLICE: Dynamic size (dùng nhiều, giống TS array)
	// Declare
	var slice1 []int = []int{1, 2, 3}
	slice2 := []string{"a", "b", "c"}

	// Make slice với capacity
	slice3 := make([]int, 5)     // len=5, cap=5, values=[0,0,0,0,0]
	slice4 := make([]int, 0, 10) // len=0, cap=10

	// Append (slice grows automatically)
	slice4 = append(slice4, 1)
	slice4 = append(slice4, 2, 3, 4)
	fmt.Println(slice4) // [1 2 3 4]

	// Slicing (giống JS slice)
	numbers := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(numbers[1:4]) // [1 2 3] - from index 1 to 3
	fmt.Println(numbers[:3])  // [0 1 2] - from start to 2
	fmt.Println(numbers[2:])  // [2 3 4 5] - from 2 to end

	// Length & Capacity
	s := make([]int, 3, 5)
	fmt.Println(len(s)) // 3 - current length
	fmt.Println(cap(s)) // 5 - capacity
}

// Common slice operations
func sliceOperations() {
	// Like TypeScript array methods

	// 1. forEach (dùng range)
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num * 2)
	}

	// 2. map (phải tự implement)
	doubled := make([]int, len(nums))
	for i, num := range nums {
		doubled[i] = num * 2
	}

	// 3. filter (phải tự implement)
	var evens []int
	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	fmt.Println(evens) // [2 4]

	// 4. Copy slice
	original := []int{1, 2, 3}
	copied := make([]int, len(original))
	copy(copied, original)
}
