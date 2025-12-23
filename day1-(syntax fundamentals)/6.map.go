package main

import "fmt"

func maps() {
	// TypeScript: const user: Record<string, any> = {}
	// Go: map[keyType]valueType

	// Declare & initialize
	var m1 map[string]int // nil map, cannot add items!

	m2 := make(map[string]int) // Empty map, can add
	m2["age"] = 30
	m2["score"] = 95

	// Literal
	user := map[string]string{
		"name":  "Tue",
		"email": "tue@example.com",
	}

	// Access
	name := user["name"]
	fmt.Println(name)

	// Check if key exists (QUAN TRỌNG!)
	email, exists := user["email"]
	if exists {
		fmt.Println("Email:", email)
	}

	// No key returns zero value
	phone := user["phone"] // "" (empty string)

	// Delete
	delete(user, "email")

	// Loop through map
	for key, value := range user {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Get only keys
	for key := range user {
		fmt.Println(key)
	}
}
