package main

import "fmt"

// TypeScript Interface: Structure của data
// Go Interface: Behavior (methods) của type

// Go interface (định nghĩa behaviors)
type Speaker interface {
	Speak() string
}

// Bất kỳ type nào có method Speak() string là Speaker
// KHÔNG CẦN "implements" keyword!

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

// Polymorphism (đa hình)
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func interfaceExample() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	// Cả dog và cat đều implement Speaker
	makeSound(dog) // Woof!
	makeSound(cat) // Meow!
}

// Practical example: Repository pattern (giống NestJS)
type UserRepository interface {
	FindByID(id int) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}

// PostgreSQL implementation
type PostgresUserRepo struct {
	// db connection
}

func (r *PostgresUserRepo) FindByID(id int) (*User, error) {
	// SELECT * FROM users WHERE id = ?
	return &User{Name: "John", Age: 30}, nil
}

func (r *PostgresUserRepo) Create(user *User) error {
	// INSERT INTO users ...
	return nil
}

func (r *PostgresUserRepo) Update(user *User) error {
	// UPDATE users ...
	return nil
}

func (r *PostgresUserRepo) Delete(id int) error {
	// DELETE FROM users ...
	return nil
}

// MongoDB implementation (cùng interface!)
type MongoUserRepo struct {
	// mongo connection
}

func (r *MongoUserRepo) FindByID(id int) (*User, error) {
	// MongoDB query
	return &User{Name: "Jane", Age: 25}, nil
}

func (r *MongoUserRepo) Create(user *User) error {
	return nil
}

func (r *MongoUserRepo) Update(user *User) error {
	return nil
}

func (r *MongoUserRepo) Delete(id int) error {
	return nil
}

// Service sử dụng interface (giống NestJS dependency injection)
type UserService struct {
	repo UserRepository // Interface, not concrete type!
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.FindByID(id)
}

func dependencyInjection() {
	// Có thể swap implementation dễ dàng
	pgRepo := &PostgresUserRepo{}
	mongoRepo := &MongoUserRepo{}

	service1 := &UserService{repo: pgRepo}
	service2 := &UserService{repo: mongoRepo}

	user1, _ := service1.GetUser(1) // Uses Postgres
	user2, _ := service2.GetUser(1) // Uses MongoDB

	fmt.Println(user1, user2)
}

// Empty interface (interface{}) - giống any trong TS
func emptyInterface() {
	var data interface{} // Có thể hold bất kỳ type nào

	data = 42
	fmt.Println(data)

	data = "hello"
	fmt.Println(data)

	data = User{Name: "Tue"}
	fmt.Println(data)

	// Type assertion để lấy lại type cụ thể
	var x interface{} = "hello"

	// Safe type assertion
	str, ok := x.(string)
	if ok {
		fmt.Println("String:", str)
	}

	// Type switch
	switch v := x.(type) {
	case string:
		fmt.Println("String:", v)
	case int:
		fmt.Println("Int:", v)
	case User:
		fmt.Println("User:", v)
	default:
		fmt.Println("Unknown type")
	}
}

// Common interfaces trong Go
// - io.Reader: Read([]byte) (int, error)
// - io.Writer: Write([]byte) (int, error)
// - error: Error() string
// - Stringer (fmt): String() string
