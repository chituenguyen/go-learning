package main

import "fmt"

// TypeScript: Everything is reference (objects) or value (primitives)
// Go: Explicit về references với pointers

func pointersBasics() {
	// Value type
	x := 10
	fmt.Println("x =", x)       // 10
	fmt.Println("Address:", &x) // &x = address of x (e.g., 0xc0000...)

	// Pointer type
	var ptr *int // Pointer to int
	ptr = &x     // ptr holds address of x

	fmt.Println("ptr =", ptr)   // Address (0xc0000...)
	fmt.Println("*ptr =", *ptr) // 10 (dereference - lấy giá trị)

	// Modify through pointer
	*ptr = 20
	fmt.Println("x =", x) // 20 (x changed!)
}

// So sánh Value vs Pointer
func modifyValue(num int) {
	num = 100 // Chỉ thay đổi copy
}

func modifyPointer(num *int) {
	*num = 100 // Thay đổi original value
}

func valueVsPointer() {
	x := 10

	modifyValue(x)
	fmt.Println(x) // 10 (không đổi)

	modifyPointer(&x)
	fmt.Println(x) // 100 (đã đổi!)
}

// Structs với pointers (QUAN TRỌNG!)
type User struct {
	Name string
	Age  int
}

func modifyUserValue(u User) {
	u.Age = 30 // Chỉ thay đổi copy
}

func modifyUserPointer(u *User) {
	u.Age = 30 // Thay đổi original (Go tự dereference)
	// Không cần (*u).Age, Go làm tự động
}

func structPointers() {
	user := User{Name: "Tue", Age: 28}

	modifyUserValue(user)
	fmt.Println(user.Age) // 28 (không đổi)

	modifyUserPointer(&user)
	fmt.Println(user.Age) // 30 (đã đổi!)
}

// Khi nào dùng pointer?
// 1. Muốn modify original value
// 2. Struct lớn (tránh copy overhead)
// 3. Methods cần modify receiver (xem phần tiếp)

// Nil pointer (giống null)
func nilPointers() {
	var ptr *int
	fmt.Println(ptr) // <nil>

	// DANGER! Panic nếu dereference nil
	// fmt.Println(*ptr)  // runtime error

	// Always check nil before use
	if ptr != nil {
		fmt.Println(*ptr)
	}
}

// new() function - allocate memory
func newFunction() {
	// new(T) returns *T with zero value
	ptr := new(int)
	fmt.Println(*ptr) // 0

	*ptr = 42
	fmt.Println(*ptr) // 42

	// With structs
	user := new(User) // *User
	user.Name = "John"
}
