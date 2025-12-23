package main

import (
	"encoding/json"
	"fmt"
)

// Struct tags cho JSON
type User struct {
	ID       int    `json:"id"`            // Map to "id"
	Name     string `json:"name"`          // Map to "name"
	Email    string `json:"email"`         // Map to "email"
	Password string `json:"-"`             // Ignore field (never in JSON)
	Age      int    `json:"age,omitempty"` // Omit if zero value
	IsActive bool   `json:"is_active,omitempty"`
}

// JSON Marshaling (struct → JSON)
// NestJS: JSON.stringify()
func marshaling() {
	user := User{
		ID:       1,
		Name:     "Tue",
		Email:    "tue@example.com",
		Password: "secret123",
		Age:      28,
		IsActive: true,
	}

	// Marshal to JSON bytes
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert to string
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
	// Output: {"id":1,"name":"Tue","email":"tue@example.com","age":28,"is_active":true}
	// Note: Password không có (json:"-")

	// Marshal với indent (pretty print)
	prettyJSON, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(prettyJSON))
}

// JSON Unmarshaling (JSON → struct)
// NestJS: JSON.parse()
func unmarshaling() {
	jsonStr := `{
        "id": 1,
        "name": "Tue",
        "email": "tue@example.com",
        "age": 28
    }`

	var user User
	err := json.Unmarshal([]byte(jsonStr), &user) // Pass pointer!
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("User: %+v\n", user)
	// Output: User: {ID:1 Name:Tue Email:tue@example.com Password: Age:28 IsActive:false}
}

// Working với arrays
func jsonArrays() {
	users := []User{
		{ID: 1, Name: "Tue", Email: "tue@example.com"},
		{ID: 2, Name: "John", Email: "john@example.com"},
	}

	// Marshal array
	jsonBytes, _ := json.Marshal(users)
	fmt.Println(string(jsonBytes))

	// Unmarshal array
	jsonStr := `[
        {"id": 1, "name": "Tue"},
        {"id": 2, "name": "John"}
    ]`

	var parsedUsers []User
	json.Unmarshal([]byte(jsonStr), &parsedUsers)
	fmt.Printf("%+v\n", parsedUsers)
}

// Generic JSON (unknown structure)
func genericJSON() {
	jsonStr := `{"name": "Tue", "age": 28, "hobbies": ["coding", "gaming"]}`

	// Use map[string]interface{} (giống any trong TS)
	var data map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &data)

	fmt.Println(data["name"]) // Tue
	fmt.Println(data["age"])  // 28 (as float64!)

	// Type assertion
	if name, ok := data["name"].(string); ok {
		fmt.Println("Name:", name)
	}

	if age, ok := data["age"].(float64); ok { // JSON numbers → float64
		fmt.Println("Age:", int(age))
	}
}
