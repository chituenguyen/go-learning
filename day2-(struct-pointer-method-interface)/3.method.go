package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

// Method với VALUE receiver (gets a copy)
// TypeScript: class User { getName(): string {...} }
func (u User) GetName() string {
	return u.Name
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

// Method với POINTER receiver (modifies original)
func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) IncrementAge() {
	u.Age++
}

// RULE: Value receiver vs Pointer receiver?
// - Pointer receiver: Nếu method cần modify struct
// - Pointer receiver: Nếu struct lớn (avoid copy)
// - Value receiver: Nếu struct nhỏ và không modify

func methodsExample() {
	user := User{Name: "Tue", Age: 28, Email: "tue@example.com"}

	// Call methods
	fmt.Println(user.GetName()) // Tue
	fmt.Println(user.IsAdult()) // true

	// Go tự động convert value ↔ pointer khi call methods
	user.SetName("John") // Go tự động: (&user).SetName("John")
	user.IncrementAge()  // Go tự động: (&user).IncrementAge()

	fmt.Println(user.Name) // John
	fmt.Println(user.Age)  // 29
}

// NestJS DTO pattern trong Go
type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validation method
func (dto *CreateUserDTO) Validate() error {
	if dto.Name == "" {
		return fmt.Errorf("name is required")
	}
	if dto.Email == "" {
		return fmt.Errorf("email is required")
	}
	if len(dto.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	return nil
}

// Transform method
func (dto *CreateUserDTO) ToUser() User {
	return User{
		Name:  dto.Name,
		Email: dto.Email,
	}
}
