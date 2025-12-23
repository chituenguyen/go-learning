package main

import "fmt"

// TypeScript:
// interface User {
//     name: string;
//     age: number;
//     email: string;
// }

// Go Struct (giống class nhưng không có inheritance)
type User struct {
	Name  string // Uppercase = public (exported)
	Age   int
	email string // lowercase = private (unexported)
}

// Khởi tạo struct
func createUsers() {
	// Cách 1: Zero values
	var user1 User
	fmt.Println(user1) // {  0 } - Name="", Age=0, email=""

	// Cách 2: Literal với field names (recommended)
	user2 := User{
		Name:  "Tue",
		Age:   28,
		email: "tue@example.com",
	}

	// Cách 3: Positional (không khuyến khích)
	user3 := User{"John", 30, "john@example.com"}

	// Cách 4: Pointer (quan trọng!)
	user4 := &User{
		Name: "Alice",
		Age:  25,
	}

	// Access fields
	fmt.Println(user2.Name) // Tue
	fmt.Println(user4.Name) // Alice (Go tự dereference)

	// Modify
	user2.Age = 29
}

// Nested structs
type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Employee struct {
	Name    string
	Age     int
	Address Address // Nested
}

func nestedStructs() {
	emp := Employee{
		Name: "Tue",
		Age:  28,
		Address: Address{
			Street:  "123 Main St",
			City:    "HCMC",
			ZipCode: "70000",
		},
	}

	fmt.Println(emp.Address.City) // HCMC
}

// Anonymous structs (giống inline types TS)
func anonymousStruct() {
	// TypeScript: const config: {host: string, port: number} = {...}
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}

	fmt.Println(config.Host)
}

// Struct tags (metadata - dùng cho JSON, DB)
type Product struct {
	ID    int     `json:"id" db:"product_id"`
	Name  string  `json:"name" db:"product_name"`
	Price float64 `json:"price" db:"price"`
}

// Sẽ dùng khi parse JSON:
// {"id": 1, "name": "Phone", "price": 999.99}
