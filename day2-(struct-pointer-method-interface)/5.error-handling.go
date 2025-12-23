package main

import (
	"errors"
	"fmt"
)

// Go KHÔNG có try/catch!
// Errors là return values

// TypeScript:
// try {
//     const user = await getUserById(1);
// } catch (error) {
//     console.error(error);
// }

// Go style:
func getUserById(id int) (*User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user id")
	}

	// Simulate DB query
	if id == 999 {
		return nil, errors.New("user not found")
	}

	return &User{Name: "John", Age: 30}, nil
}

func errorHandling() {
	// Pattern: Check error immediately
	user, err := getUserById(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User:", user.Name)
}

// Custom errors
type NotFoundError struct {
	Resource string
	ID       int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with id %d not found", e.Resource, e.ID)
}

func findUser(id int) (*User, error) {
	if id == 999 {
		return nil, &NotFoundError{Resource: "User", ID: id}
	}
	return &User{Name: "John"}, nil
}

// Error wrapping (Go 1.13+)
func processUser(id int) error {
	user, err := getUserById(id)
	if err != nil {
		return fmt.Errorf("failed to process user: %w", err) // %w wraps error
	}

	fmt.Println("Processing:", user.Name)
	return nil
}

func errorWrapping() {
	err := processUser(999)
	if err != nil {
		fmt.Println(err) // "failed to process user: user not found"

		// Unwrap error
		if errors.Is(err, errors.New("user not found")) {
			fmt.Println("User not found!")
		}
	}
}

// Best practices
func bestPractices() {
	// 1. Always check errors
	user, err := getUserById(1)
	if err != nil {
		// Handle error
		return
	}

	// 2. Return early
	if user.Age < 18 {
		fmt.Println("User is minor")
		return
	}

	// 3. Don't ignore errors
	// BAD:
	// getUserById(1)  // Ignoring error

	// GOOD:
	_, _ = getUserById(1) // Explicitly ignore with _
}
