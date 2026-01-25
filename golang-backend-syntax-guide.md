# TỔNG HỢP GOLANG SYNTAX CHO BACKEND DEVELOPER

## 📑 MỤC LỤC
1. [Cơ bản](#1-cơ-bản)
2. [Kiểu dữ liệu](#2-kiểu-dữ-liệu)
3. [Control Flow](#3-control-flow)
4. [Functions](#4-functions)
5. [Structs & Methods](#5-structs--methods)
6. [Interfaces](#6-interfaces)
7. [Error Handling](#7-error-handling)
8. [Concurrency](#8-concurrency)
9. [Pointers](#9-pointers)
10. [Packages & Modules](#10-packages--modules)
11. [File I/O](#11-file-io)
12. [HTTP & Web](#12-http--web)
13. [Database](#13-database)
14. [Testing](#14-testing)
15. [Context](#15-context)
16. [JSON](#16-json)
17. [Backend Patterns](#17-backend-patterns)

---

## 1. Cơ BẢN

### Package Declaration
```go
package main  // executable package
package user  // library package
```

### Import
```go
// Single import
import "fmt"

// Multiple imports
import (
    "fmt"
    "net/http"
    "database/sql"
)

// Alias import
import (
    f "fmt"
    _ "github.com/lib/pq"  // import for side effects
)
```

### Variables
```go
// Khai báo với var
var name string = "Tue"
var age int = 25

// Khai báo ngắn gọn (chỉ dùng trong function)
name := "Tue"
age := 25

// Multiple declaration
var (
    host     string = "localhost"
    port     int    = 5432
    database string = "mydb"
)

// Zero values
var s string   // ""
var i int      // 0
var b bool     // false
var p *int     // nil
```

### Constants
```go
const Pi = 3.14

const (
    StatusOK       = 200
    StatusNotFound = 404
    StatusError    = 500
)

// iota - auto increment
const (
    Sunday = iota     // 0
    Monday            // 1
    Tuesday           // 2
    Wednesday         // 3
)

const (
    _ = iota          // skip 0
    KB = 1 << (10 * iota)  // 1024
    MB                     // 1048576
    GB                     // 1073741824
)
```

---

## 2. KIỂU DỮ LIỆU

### Primitive Types
```go
// Integer
var i8 int8       // -128 to 127
var i16 int16     // -32768 to 32767
var i32 int32     // -2147483648 to 2147483647
var i64 int64     // -9223372036854775808 to 9223372036854775807
var ui8 uint8     // 0 to 255
var ui16 uint16   // 0 to 65535
var ui32 uint32   // 0 to 4294967295
var ui64 uint64   // 0 to 18446744073709551615

// Platform dependent
var i int         // 32 or 64 bit
var ui uint       // 32 or 64 bit

// Float
var f32 float32   // IEEE-754 32-bit
var f64 float64   // IEEE-754 64-bit

// String
var s string = "Hello, 世界"
var raw string = `multi
line
string`

// Boolean
var isActive bool = true

// Byte (alias for uint8)
var b byte = 'A'

// Rune (alias for int32, represents Unicode code point)
var r rune = '世'
```

### Arrays
```go
// Fixed size
var arr [5]int
arr[0] = 1

// Initialize
numbers := [5]int{1, 2, 3, 4, 5}
auto := [...]int{1, 2, 3}  // compiler counts

// Multi-dimensional
matrix := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
```

### Slices (Dynamic Arrays)
```go
// Declare
var s []int

// Make slice
s := make([]int, 5)       // length 5
s := make([]int, 5, 10)   // length 5, capacity 10

// Literal
nums := []int{1, 2, 3, 4, 5}

// Slicing
a := nums[1:4]   // [2 3 4]
b := nums[:3]    // [1 2 3]
c := nums[2:]    // [3 4 5]
d := nums[:]     // full copy

// Append
s = append(s, 1)
s = append(s, 2, 3, 4)
s = append(s, []int{5, 6}...)

// Length and capacity
len(s)   // số phần tử hiện tại
cap(s)   // dung lượng tối đa

// Copy
dst := make([]int, len(src))
copy(dst, src)

// Iterate
for i, v := range nums {
    fmt.Printf("Index: %d, Value: %d\n", i, v)
}
```

### Maps
```go
// Declare
var m map[string]int

// Make
m := make(map[string]int)

// Literal
users := map[string]int{
    "alice": 25,
    "bob":   30,
}

// Set value
m["key"] = 42

// Get value
value := m["key"]

// Check if key exists
value, exists := m["key"]
if exists {
    fmt.Println(value)
}

// Delete
delete(m, "key")

// Iterate
for key, value := range users {
    fmt.Printf("%s: %d\n", key, value)
}

// Length
len(m)
```

### Structs
```go
// Define
type User struct {
    ID        int
    Name      string
    Email     string
    CreatedAt time.Time
}

// Initialize
u1 := User{
    ID:    1,
    Name:  "Tue",
    Email: "tue@example.com",
}

u2 := User{1, "Tue", "tue@example.com", time.Now()}

// Anonymous struct
person := struct {
    Name string
    Age  int
}{
    Name: "Tue",
    Age:  25,
}

// Embedded struct
type Admin struct {
    User
    Role string
}

admin := Admin{
    User: User{ID: 1, Name: "Admin"},
    Role: "superadmin",
}

// Struct tags (cho JSON, DB, validation)
type Product struct {
    ID    int    `json:"id" db:"product_id"`
    Name  string `json:"name" db:"product_name" validate:"required"`
    Price float64 `json:"price" db:"price" validate:"gt=0"`
}
```

---

## 3. CONTROL FLOW

### If-Else
```go
// Basic
if x > 0 {
    fmt.Println("Positive")
}

// If-else
if x > 0 {
    fmt.Println("Positive")
} else {
    fmt.Println("Non-positive")
}

// If-else if-else
if x > 0 {
    fmt.Println("Positive")
} else if x < 0 {
    fmt.Println("Negative")
} else {
    fmt.Println("Zero")
}

// If with initialization
if err := doSomething(); err != nil {
    return err
}

// Short variable scope
if val, ok := cache.Get(key); ok {
    return val
}
```

### Switch
```go
// Basic switch
switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("End of work week")
default:
    fmt.Println("Midweek")
}

// Multiple cases
switch day {
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Weekday")
}

// Switch with expression
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
case score >= 70:
    grade = "C"
default:
    grade = "F"
}

// Type switch
switch v := i.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
case bool:
    fmt.Printf("Boolean: %t\n", v)
default:
    fmt.Printf("Unknown type\n")
}

// Fallthrough
switch num {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("One or Two")
}
```

### Loops
```go
// Basic for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style loop
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}

// Infinite loop
for {
    // break để thoát
    if shouldStop {
        break
    }
}

// Range over slice
nums := []int{1, 2, 3, 4, 5}
for i, v := range nums {
    fmt.Printf("Index: %d, Value: %d\n", i, v)
}

// Range over map
for key, value := range myMap {
    fmt.Printf("%s: %v\n", key, value)
}

// Range over string (runes)
for i, ch := range "Hello" {
    fmt.Printf("%d: %c\n", i, ch)
}

// Ignore index
for _, v := range nums {
    fmt.Println(v)
}

// Only index
for i := range nums {
    fmt.Println(i)
}

// Break and continue
for i := 0; i < 10; i++ {
    if i == 5 {
        continue  // skip 5
    }
    if i == 8 {
        break     // stop at 8
    }
    fmt.Println(i)
}

// Labeled break/continue
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer
        }
        fmt.Printf("%d,%d ", i, j)
    }
}
```

### Defer
```go
// Execute sau khi function return
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // sẽ close khi function return
    
    // work with file
    return nil
}

// Multiple defers (LIFO - Last In First Out)
func example() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
    // Output: 3, 2, 1
}

// Common pattern với transaction
func transferMoney(from, to int, amount float64) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()  // rollback nếu có lỗi
    
    // perform operations
    
    return tx.Commit()
}
```

### Panic and Recover
```go
// Panic - stop execution
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}

// Recover - catch panic
func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    
    result = a / b
    return
}

// Middleware pattern với recover
func recoverMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("panic: %v", err)
                http.Error(w, "Internal Server Error", 500)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

---

## 4. FUNCTIONS

### Basic Function
```go
// Simple function
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// With return value
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Named return values
func getUser(id int) (name string, age int, err error) {
    if id <= 0 {
        err = fmt.Errorf("invalid user id")
        return
    }
    name = "Tue"
    age = 25
    return  // naked return
}

// Same type parameters
func sum(a, b, c int) int {
    return a + b + c
}
```

### Variadic Functions
```go
// Accept variable number of arguments
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// Usage
sum(1, 2, 3)
sum(1, 2, 3, 4, 5)

nums := []int{1, 2, 3, 4}
sum(nums...)  // spread operator

// Mix with regular parameters
func log(level string, messages ...string) {
    fmt.Printf("[%s] ", level)
    for _, msg := range messages {
        fmt.Print(msg, " ")
    }
    fmt.Println()
}
```

### Anonymous Functions & Closures
```go
// Anonymous function
func() {
    fmt.Println("Anonymous function")
}()

// Assign to variable
add := func(a, b int) int {
    return a + b
}
result := add(5, 3)

// Closure
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c())  // 1
fmt.Println(c())  // 2
fmt.Println(c())  // 3

// Common backend pattern: middleware
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", 401)
            return
        }
        next(w, r)
    }
}
```

### Function Types
```go
// Define function type
type Operation func(int, int) int

// Use as parameter
func calculate(a, b int, op Operation) int {
    return op(a, b)
}

// Usage
add := func(a, b int) int { return a + b }
result := calculate(5, 3, add)

// Common backend pattern: handler
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Middleware func(HandlerFunc) HandlerFunc
```

### Methods vs Functions
```go
// Method - function with receiver
type Calculator struct {
    value int
}

// Value receiver (không thay đổi original)
func (c Calculator) GetValue() int {
    return c.value
}

// Pointer receiver (có thể thay đổi original)
func (c *Calculator) Add(n int) {
    c.value += n
}

func (c *Calculator) Subtract(n int) {
    c.value -= n
}

// Usage
calc := Calculator{value: 10}
calc.Add(5)
fmt.Println(calc.GetValue())  // 15
```

---

## 5. STRUCTS & METHODS

### Struct Definition
```go
// Basic struct
type User struct {
    ID        int
    Username  string
    Email     string
    Password  string
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Constructor function
func NewUser(username, email, password string) *User {
    return &User{
        Username:  username,
        Email:     email,
        Password:  hashPassword(password),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
}

// Struct with validation
func (u *User) Validate() error {
    if u.Username == "" {
        return fmt.Errorf("username is required")
    }
    if len(u.Password) < 8 {
        return fmt.Errorf("password must be at least 8 characters")
    }
    return nil
}
```

### Methods
```go
type Account struct {
    ID      int
    Balance float64
}

// Value receiver - read only
func (a Account) GetBalance() float64 {
    return a.Balance
}

// Pointer receiver - can modify
func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return fmt.Errorf("amount must be positive")
    }
    a.Balance += amount
    return nil
}

func (a *Account) Withdraw(amount float64) error {
    if amount <= 0 {
        return fmt.Errorf("amount must be positive")
    }
    if amount > a.Balance {
        return fmt.Errorf("insufficient balance")
    }
    a.Balance -= amount
    return nil
}

// Method chaining
func (a *Account) SetBalance(balance float64) *Account {
    a.Balance = balance
    return a
}

// Usage with method chaining
account := &Account{ID: 1}
account.SetBalance(1000).Deposit(500)
```

### Composition (Embedding)
```go
// Base struct
type Model struct {
    ID        int
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Embedded struct
type Product struct {
    Model
    Name  string
    Price float64
    Stock int
}

// Product tự động có các field và method của Model
product := Product{
    Model: Model{
        ID:        1,
        CreatedAt: time.Now(),
    },
    Name:  "Laptop",
    Price: 999.99,
}

// Access embedded fields
fmt.Println(product.ID)
fmt.Println(product.CreatedAt)

// Multiple embedding
type AuditLog struct {
    Model
    UserID int
    Action string
}

type AdminUser struct {
    User
    Permissions []string
}
```

### Private vs Public
```go
// Public (exported) - starts with uppercase
type User struct {
    ID   int     // Public field
    Name string  // Public field
    age  int     // Private field
}

// Public method
func (u *User) GetAge() int {
    return u.age
}

// Private method
func (u *User) validate() error {
    // only accessible within the same package
    return nil
}
```

---

## 6. INTERFACES

### Interface Definition
```go
// Basic interface
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Interface composition
type ReadWriter interface {
    Reader
    Writer
}

// Empty interface (accepts any type)
var i interface{}
i = 42
i = "hello"
i = []int{1, 2, 3}
```

### Common Backend Interfaces
```go
// Repository interface
type UserRepository interface {
    Create(user *User) error
    GetByID(id int) (*User, error)
    GetByEmail(email string) (*User, error)
    Update(user *User) error
    Delete(id int) error
    List(limit, offset int) ([]*User, error)
}

// Implementation
type PostgresUserRepository struct {
    db *sql.DB
}

func (r *PostgresUserRepository) Create(user *User) error {
    query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
    return r.db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
}

func (r *PostgresUserRepository) GetByID(id int) (*User, error) {
    user := &User{}
    query := `SELECT id, username, email FROM users WHERE id = $1`
    err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email)
    if err != nil {
        return nil, err
    }
    return user, nil
}

// Service interface
type UserService interface {
    Register(username, email, password string) (*User, error)
    Login(email, password string) (string, error)
    GetProfile(userID int) (*User, error)
}
```

### Type Assertions
```go
// Type assertion
var i interface{} = "hello"

s := i.(string)  // panic if wrong type
fmt.Println(s)

// Safe type assertion
s, ok := i.(string)
if ok {
    fmt.Println(s)
}

// Type switch
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

### Interface Best Practices
```go
// Small interfaces (1-2 methods)
type Stringer interface {
    String() string
}

type Closer interface {
    Close() error
}

// Accept interfaces, return structs
func ProcessData(r Reader) (*Result, error) {
    // accept interface parameter
    return &Result{}, nil
}

// Dependency injection
type Server struct {
    userRepo    UserRepository
    userService UserService
    logger      Logger
}

func NewServer(repo UserRepository, service UserService, logger Logger) *Server {
    return &Server{
        userRepo:    repo,
        userService: service,
        logger:      logger,
    }
}
```

---

## 7. ERROR HANDLING

### Basic Error Handling
```go
// Return error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Check error
result, err := divide(10, 0)
if err != nil {
    log.Fatal(err)
}

// Multiple return values with error
func readFile(filename string) ([]byte, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to read file: %w", err)
    }
    return data, nil
}
```

### Custom Errors
```go
// Simple custom error
type ValidationError struct {
    Field string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Usage
func validateUser(user *User) error {
    if user.Email == "" {
        return &ValidationError{
            Field:   "email",
            Message: "is required",
        }
    }
    return nil
}

// Complex custom error
type APIError struct {
    StatusCode int
    Message    string
    Err        error
}

func (e *APIError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("API Error %d: %s (caused by: %v)", e.StatusCode, e.Message, e.Err)
    }
    return fmt.Sprintf("API Error %d: %s", e.StatusCode, e.Message)
}

func (e *APIError) Unwrap() error {
    return e.Err
}
```

### Error Wrapping (Go 1.13+)
```go
// Wrap error with context
func processOrder(orderID int) error {
    order, err := getOrder(orderID)
    if err != nil {
        return fmt.Errorf("process order %d: %w", orderID, err)
    }
    // process
    return nil
}

// Check error type
var apiErr *APIError
if errors.As(err, &apiErr) {
    fmt.Printf("API Error with status %d\n", apiErr.StatusCode)
}

// Check if error is specific error
if errors.Is(err, sql.ErrNoRows) {
    // handle not found
}

// Unwrap error
unwrapped := errors.Unwrap(err)
```

### Error Handling Patterns
```go
// Early return pattern
func createUser(user *User) error {
    if err := user.Validate(); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }
    
    if err := checkDuplicate(user.Email); err != nil {
        return fmt.Errorf("duplicate check failed: %w", err)
    }
    
    if err := saveUser(user); err != nil {
        return fmt.Errorf("save user failed: %w", err)
    }
    
    return nil
}

// Sentinel errors
var (
    ErrNotFound      = errors.New("resource not found")
    ErrUnauthorized  = errors.New("unauthorized")
    ErrInvalidInput  = errors.New("invalid input")
)

func GetUser(id int) (*User, error) {
    // ...
    if notFound {
        return nil, ErrNotFound
    }
    return user, nil
}

// Usage
user, err := GetUser(123)
if errors.Is(err, ErrNotFound) {
    // handle not found
}

// Error group pattern (multiple errors)
type MultiError struct {
    Errors []error
}

func (m *MultiError) Error() string {
    var msgs []string
    for _, err := range m.Errors {
        msgs = append(msgs, err.Error())
    }
    return strings.Join(msgs, "; ")
}

func (m *MultiError) Add(err error) {
    if err != nil {
        m.Errors = append(m.Errors, err)
    }
}
```

### Panic vs Error
```go
// Use error for expected errors
func withdraw(amount float64) error {
    if amount > balance {
        return fmt.Errorf("insufficient balance")
    }
    return nil
}

// Use panic for programming errors (should never happen)
func mustConnect(dsn string) *sql.DB {
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        panic(fmt.Sprintf("database connection failed: %v", err))
    }
    return db
}

// Recover from panic in specific cases
func safeHandler(fn func()) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Recovered from panic: %v", r)
        }
    }()
    fn()
}
```

---

## 8. CONCURRENCY

### Goroutines
```go
// Start goroutine
go func() {
    fmt.Println("Running in goroutine")
}()

// With parameters
go processData(data)

// Multiple goroutines
for i := 0; i < 10; i++ {
    go func(id int) {
        fmt.Printf("Worker %d\n", id)
    }(i)  // pass i as parameter
}

// Wait for goroutines
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        // do work
        fmt.Printf("Worker %d done\n", id)
    }(i)
}
wg.Wait()
```

### Channels
```go
// Unbuffered channel
ch := make(chan int)

// Buffered channel
ch := make(chan int, 100)

// Send to channel
ch <- 42

// Receive from channel
value := <-ch

// Close channel
close(ch)

// Check if channel is closed
value, ok := <-ch
if !ok {
    fmt.Println("Channel closed")
}

// Range over channel
for value := range ch {
    fmt.Println(value)
}

// Select statement
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
case ch3 <- 42:
    fmt.Println("Sent to ch3")
default:
    fmt.Println("No communication")
}
```

### Common Concurrency Patterns
```go
// Worker pool pattern
func workerPool(jobs <-chan int, results chan<- int, numWorkers int) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for job := range jobs {
                // process job
                results <- job * 2
            }
        }(i)
    }
    
    wg.Wait()
    close(results)
}

// Usage
jobs := make(chan int, 100)
results := make(chan int, 100)

go workerPool(jobs, results, 5)

for i := 1; i <= 10; i++ {
    jobs <- i
}
close(jobs)

for result := range results {
    fmt.Println(result)
}

// Pipeline pattern
func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Usage
nums := generator(1, 2, 3, 4)
squared := square(nums)
for result := range squared {
    fmt.Println(result)
}

// Fan-out, Fan-in pattern
func fanOut(in <-chan int, numWorkers int) []<-chan int {
    channels := make([]<-chan int, numWorkers)
    for i := 0; i < numWorkers; i++ {
        channels[i] = process(in)
    }
    return channels
}

func fanIn(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)
    
    multiplex := func(ch <-chan int) {
        defer wg.Done()
        for v := range ch {
            out <- v
        }
    }
    
    wg.Add(len(channels))
    for _, ch := range channels {
        go multiplex(ch)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}

// Timeout pattern
func doWithTimeout(timeout time.Duration) error {
    done := make(chan bool)
    
    go func() {
        // long running task
        time.Sleep(2 * time.Second)
        done <- true
    }()
    
    select {
    case <-done:
        return nil
    case <-time.After(timeout):
        return fmt.Errorf("operation timed out")
    }
}
```

### Synchronization Primitives
```go
// Mutex
var (
    mu    sync.Mutex
    count int
)

func increment() {
    mu.Lock()
    defer mu.Unlock()
    count++
}

// RWMutex (multiple readers, single writer)
var (
    rwMu sync.RWMutex
    data = make(map[string]string)
)

func read(key string) string {
    rwMu.RLock()
    defer rwMu.RUnlock()
    return data[key]
}

func write(key, value string) {
    rwMu.Lock()
    defer rwMu.Unlock()
    data[key] = value
}

// sync.Once (execute only once)
var (
    once     sync.Once
    instance *Database
)

func GetInstance() *Database {
    once.Do(func() {
        instance = &Database{}
        instance.Connect()
    })
    return instance
}

// sync.WaitGroup (shown earlier)

// Atomic operations
var counter int64

func incrementAtomic() {
    atomic.AddInt64(&counter, 1)
}

func getCounter() int64 {
    return atomic.LoadInt64(&counter)
}
```

### Context for Cancellation
```go
// With cancel
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    select {
    case <-ctx.Done():
        fmt.Println("Cancelled")
        return
    case <-time.After(5 * time.Second):
        fmt.Println("Completed")
    }
}()

// Cancel after 2 seconds
time.Sleep(2 * time.Second)
cancel()

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

// With deadline
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

// Pass context through functions
func processRequest(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
        // do work
        return nil
    }
}
```

---

## 9. POINTERS

### Pointer Basics
```go
// Declare pointer
var p *int

// Get address of variable
x := 42
p = &x

// Dereference pointer
fmt.Println(*p)  // 42

// Modify through pointer
*p = 100
fmt.Println(x)   // 100

// Nil pointer
var ptr *int
if ptr == nil {
    fmt.Println("Pointer is nil")
}
```

### Pointers with Structs
```go
type User struct {
    Name string
    Age  int
}

// Pointer to struct
user := &User{Name: "Tue", Age: 25}

// Access fields (auto-dereference)
fmt.Println(user.Name)  // same as (*user).Name

// Modify through pointer
func updateUser(u *User) {
    u.Name = "Updated"
    u.Age = 26
}

user := &User{Name: "Tue", Age: 25}
updateUser(user)
```

### When to Use Pointers
```go
// 1. Modify original value
func increment(x *int) {
    *x++
}

num := 10
increment(&num)
fmt.Println(num)  // 11

// 2. Avoid copying large structs
type LargeStruct struct {
    Data [1000000]byte
}

// Good - pass pointer
func processLarge(s *LargeStruct) {
    // work with s
}

// Bad - copies entire struct
func processLargeBad(s LargeStruct) {
    // work with s
}

// 3. Optional values (nil pointer = no value)
func getUser(id int) *User {
    if id <= 0 {
        return nil
    }
    return &User{Name: "Tue"}
}

user := getUser(123)
if user != nil {
    fmt.Println(user.Name)
}

// 4. Methods that modify receiver
func (u *User) UpdateAge(age int) {
    u.Age = age
}
```

### Pointer Pitfalls
```go
// Looping with pointers - WRONG
users := make([]*User, 0)
for i := 0; i < 3; i++ {
    user := User{Name: fmt.Sprintf("User%d", i)}
    users = append(users, &user)  // BUG: all point to same variable
}

// Looping with pointers - CORRECT
users := make([]*User, 0)
for i := 0; i < 3; i++ {
    user := User{Name: fmt.Sprintf("User%d", i)}
    users = append(users, &user)  // OK with proper scoping
}

// Or
for i := 0; i < 3; i++ {
    users = append(users, &User{Name: fmt.Sprintf("User%d", i)})
}
```

---

## 10. PACKAGES & MODULES

### Package Structure
```go
// File: myapp/main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
}

// File: myapp/user/user.go
package user

type User struct {
    ID   int
    Name string
}

func NewUser(name string) *User {
    return &User{Name: name}
}

// File: myapp/user/repository.go
package user

type Repository struct {
    // ...
}
```

### Go Modules
```bash
# Initialize module
go mod init github.com/username/project

# Add dependency
go get github.com/gin-gonic/gin

# Update dependencies
go get -u

# Tidy dependencies
go mod tidy

# Vendor dependencies
go mod vendor
```

### go.mod Example
```go
module github.com/tue/banking-api

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/lib/pq v1.10.9
    github.com/golang-jwt/jwt/v5 v5.0.0
)

require (
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/gabriel-vasile/mimetype v1.4.2 // indirect
)
```

### Import Patterns
```go
// Standard import
import "fmt"
import "net/http"

// Grouped imports
import (
    "fmt"
    "net/http"
    
    "github.com/gin-gonic/gin"
    "github.com/lib/pq"
)

// Aliased import
import (
    f "fmt"
    mypkg "github.com/company/very-long-package-name"
)

// Blank import (for side effects)
import _ "github.com/lib/pq"

// Dot import (not recommended)
import . "fmt"
Println("No prefix needed")
```

### Project Structure
```
myapp/
├── cmd/
│   └── api/
│       └── main.go          # Application entry point
├── internal/                # Private application code
│   ├── domain/
│   │   └── user.go         # Domain models
│   ├── repository/
│   │   └── user_repo.go    # Data access
│   ├── service/
│   │   └── user_service.go # Business logic
│   └── handler/
│       └── user_handler.go # HTTP handlers
├── pkg/                     # Public libraries
│   ├── logger/
│   └── validator/
├── config/                  # Configuration
│   └── config.go
├── migrations/             # Database migrations
├── go.mod
└── go.sum
```

---

## 11. FILE I/O

### Reading Files
```go
// Read entire file
data, err := os.ReadFile("file.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))

// Read line by line
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)
}

if err := scanner.Err(); err != nil {
    log.Fatal(err)
}

// Read with buffer
buf := make([]byte, 1024)
n, err := file.Read(buf)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
```

### Writing Files
```go
// Write entire file
data := []byte("Hello, World!")
err := os.WriteFile("output.txt", data, 0644)
if err != nil {
    log.Fatal(err)
}

// Append to file
file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()

if _, err := file.WriteString("New line\n"); err != nil {
    log.Fatal(err)
}

// Buffered write
writer := bufio.NewWriter(file)
writer.WriteString("Buffered line\n")
writer.Flush()
```

### File Operations
```go
// Check if file exists
if _, err := os.Stat("file.txt"); os.IsNotExist(err) {
    fmt.Println("File does not exist")
}

// Delete file
err := os.Remove("file.txt")

// Rename file
err := os.Rename("old.txt", "new.txt")

// Create directory
err := os.Mkdir("mydir", 0755)
err := os.MkdirAll("path/to/dir", 0755)  // create all parent dirs

// List directory
entries, err := os.ReadDir(".")
for _, entry := range entries {
    fmt.Println(entry.Name(), entry.IsDir())
}
```

### CSV Operations
```go
import "encoding/csv"

// Read CSV
file, _ := os.Open("data.csv")
defer file.Close()

reader := csv.NewReader(file)
records, err := reader.ReadAll()
for _, record := range records {
    fmt.Println(record)
}

// Write CSV
file, _ := os.Create("output.csv")
defer file.Close()

writer := csv.NewWriter(file)
defer writer.Flush()

writer.Write([]string{"Name", "Age", "Email"})
writer.Write([]string{"Tue", "25", "tue@example.com"})
```

---

## 12. HTTP & WEB

### Standard Library HTTP Server
```go
import "net/http"

// Simple handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message":"Hello, World!"}`))
}

// Start server
func main() {
    http.HandleFunc("/hello", helloHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// With ServeMux
mux := http.NewServeMux()
mux.HandleFunc("/users", usersHandler)
mux.HandleFunc("/products", productsHandler)
http.ListenAndServe(":8080", mux)
```

### HTTP Client
```go
// GET request
resp, err := http.Get("https://api.example.com/users")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)
fmt.Println(string(body))

// POST request
data := []byte(`{"name":"Tue"}`)
resp, err := http.Post("https://api.example.com/users", "application/json", bytes.NewBuffer(data))

// Custom request
client := &http.Client{
    Timeout: 10 * time.Second,
}

req, err := http.NewRequest("PUT", "https://api.example.com/users/1", bytes.NewBuffer(data))
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer token")

resp, err := client.Do(req)
```

### Gin Framework
```go
import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    
    // Routes
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    
    r.POST("/users", createUser)
    r.GET("/users/:id", getUser)
    r.PUT("/users/:id", updateUser)
    r.DELETE("/users/:id", deleteUser)
    
    // Route groups
    api := r.Group("/api/v1")
    {
        api.GET("/products", listProducts)
        api.POST("/products", createProduct)
    }
    
    // Middleware
    r.Use(authMiddleware())
    
    r.Run(":8080")
}

// Handler with path parameter
func getUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"user_id": id})
}

// Handler with query parameters
func listUsers(c *gin.Context) {
    page := c.DefaultQuery("page", "1")
    limit := c.Query("limit")
    c.JSON(200, gin.H{
        "page":  page,
        "limit": limit,
    })
}

// Handler with request body
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

func createUser(c *gin.Context) {
    var req CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // Create user
    c.JSON(201, gin.H{"message": "User created"})
}

// Middleware
func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        
        // Validate token
        c.Set("user_id", 123)
        c.Next()
    }
}

// Get value from context
func protectedHandler(c *gin.Context) {
    userID := c.GetInt("user_id")
    c.JSON(200, gin.H{"user_id": userID})
}
```

### REST API Example
```go
type UserHandler struct {
    service UserService
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
    users := r.Group("/api/v1/users")
    {
        users.POST("", h.Create)
        users.GET("/:id", h.GetByID)
        users.GET("", h.List)
        users.PUT("/:id", h.Update)
        users.DELETE("/:id", h.Delete)
    }
}

func (h *UserHandler) Create(c *gin.Context) {
    var req CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
        return
    }
    
    user, err := h.service.Create(&req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
        return
    }
    
    user, err := h.service.GetByID(id)
    if err != nil {
        if errors.Is(err, ErrNotFound) {
            c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, user)
}
```

---

## 13. DATABASE

### Database/SQL
```go
import (
    "database/sql"
    _ "github.com/lib/pq"
)

// Connect
db, err := sql.Open("postgres", "postgres://user:pass@localhost/dbname?sslmode=disable")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Ping to verify connection
if err := db.Ping(); err != nil {
    log.Fatal(err)
}

// Set connection pool
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(5 * time.Minute)
```

### CRUD Operations
```go
// CREATE
func createUser(db *sql.DB, user *User) error {
    query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
    err := db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
    return err
}

// READ - single row
func getUserByID(db *sql.DB, id int) (*User, error) {
    user := &User{}
    query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
    err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("user not found")
    }
    return user, err
}

// READ - multiple rows
func listUsers(db *sql.DB) ([]*User, error) {
    query := `SELECT id, name, email FROM users ORDER BY id`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    users := make([]*User, 0)
    for rows.Next() {
        user := &User{}
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    
    return users, rows.Err()
}

// UPDATE
func updateUser(db *sql.DB, user *User) error {
    query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
    result, err := db.Exec(query, user.Name, user.Email, user.ID)
    if err != nil {
        return err
    }
    
    rows, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rows == 0 {
        return fmt.Errorf("user not found")
    }
    
    return nil
}

// DELETE
func deleteUser(db *sql.DB, id int) error {
    query := `DELETE FROM users WHERE id = $1`
    _, err := db.Exec(query, id)
    return err
}
```

### Transactions
```go
func transferMoney(db *sql.DB, fromID, toID int, amount float64) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()  // rollback if not committed
    
    // Deduct from sender
    _, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, fromID)
    if err != nil {
        return err
    }
    
    // Add to receiver
    _, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, toID)
    if err != nil {
        return err
    }
    
    // Commit transaction
    return tx.Commit()
}
```

### Prepared Statements
```go
// Prepare statement
stmt, err := db.Prepare("SELECT id, name FROM users WHERE email = $1")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

// Execute multiple times
user1 := &User{}
stmt.QueryRow("user1@example.com").Scan(&user1.ID, &user1.Name)

user2 := &User{}
stmt.QueryRow("user2@example.com").Scan(&user2.ID, &user2.Name)
```

### GORM ORM
```go
import "gorm.io/gorm"
import "gorm.io/driver/postgres"

// Connect
dsn := "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// Auto migrate
db.AutoMigrate(&User{}, &Product{})

// CREATE
user := User{Name: "Tue", Email: "tue@example.com"}
result := db.Create(&user)  // user.ID sẽ được fill

// READ
var user User
db.First(&user, 1)  // find by primary key
db.First(&user, "email = ?", "tue@example.com")
db.Where("age > ?", 18).Find(&users)

// UPDATE
db.Model(&user).Update("name", "Updated Name")
db.Model(&user).Updates(User{Name: "New Name", Age: 26})

// DELETE
db.Delete(&user, 1)

// Complex queries
var users []User
db.Where("age > ?", 18).
   Where("name LIKE ?", "%Tue%").
   Order("created_at DESC").
   Limit(10).
   Offset(20).
   Find(&users)

// Joins
db.Joins("Company").Where("companies.name = ?", "ABC").Find(&users)

// Raw SQL
db.Raw("SELECT id, name FROM users WHERE age > ?", 18).Scan(&users)

// Transactions
db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&user1).Error; err != nil {
        return err
    }
    if err := tx.Create(&user2).Error; err != nil {
        return err
    }
    return nil
})
```

---

## 14. TESTING

### Basic Testing
```go
// File: math.go
package math

func Add(a, b int) int {
    return a + b
}

// File: math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

// Table-driven tests
func TestAddTable(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -1, -2},
        {"zero", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

### Test Helper Functions
```go
// Helper function
func assertEqual(t *testing.T, got, want interface{}) {
    t.Helper()  // marks this as helper function
    if got != want {
        t.Errorf("got %v; want %v", got, want)
    }
}

// Usage
func TestSomething(t *testing.T) {
    result := Calculate()
    assertEqual(t, result, 42)
}
```

### Testing with Setup/Teardown
```go
func setupTestDB(t *testing.T) *sql.DB {
    db, err := sql.Open("postgres", "postgres://localhost/test")
    if err != nil {
        t.Fatal(err)
    }
    
    // Setup schema
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (...)")
    if err != nil {
        t.Fatal(err)
    }
    
    return db
}

func teardownTestDB(t *testing.T, db *sql.DB) {
    db.Exec("DROP TABLE users")
    db.Close()
}

func TestUserRepository(t *testing.T) {
    db := setupTestDB(t)
    defer teardownTestDB(t, db)
    
    repo := NewUserRepository(db)
    
    // Run tests
    user, err := repo.Create(&User{Name: "Test"})
    if err != nil {
        t.Fatal(err)
    }
    
    assertEqual(t, user.Name, "Test")
}
```

### Mocking
```go
// Interface to mock
type UserService interface {
    GetUser(id int) (*User, error)
}

// Mock implementation
type MockUserService struct {
    GetUserFunc func(id int) (*User, error)
}

func (m *MockUserService) GetUser(id int) (*User, error) {
    return m.GetUserFunc(id)
}

// Test using mock
func TestHandler(t *testing.T) {
    mockService := &MockUserService{
        GetUserFunc: func(id int) (*User, error) {
            return &User{ID: id, Name: "Test"}, nil
        },
    }
    
    handler := NewHandler(mockService)
    
    // Test handler with mock service
}
```

### Benchmarking
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

func BenchmarkConcatenate(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = "hello" + "world"
    }
}

// Run benchmarks
// go test -bench=.
// go test -bench=. -benchmem  // include memory stats
```

### Test Coverage
```bash
# Run tests with coverage
go test -cover

# Generate coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

## 15. CONTEXT

### Basic Context
```go
import "context"

// Background context
ctx := context.Background()

// TODO context (placeholder)
ctx := context.TODO()

// With cancel
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// With deadline
deadline := time.Now().Add(10 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

### Context Values
```go
// Store value in context
type contextKey string

const userIDKey contextKey = "userID"

ctx := context.WithValue(context.Background(), userIDKey, 123)

// Retrieve value
userID := ctx.Value(userIDKey).(int)

// Type-safe context values
func SetUserID(ctx context.Context, userID int) context.Context {
    return context.WithValue(ctx, userIDKey, userID)
}

func GetUserID(ctx context.Context) (int, bool) {
    userID, ok := ctx.Value(userIDKey).(int)
    return userID, ok
}
```

### Context in HTTP
```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    
    // Add timeout
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()
    
    // Use context in database call
    user, err := getUserFromDB(ctx, userID)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    
    json.NewEncoder(w).Encode(user)
}

func getUserFromDB(ctx context.Context, userID int) (*User, error) {
    query := "SELECT * FROM users WHERE id = $1"
    
    // QueryRowContext respects context cancellation
    user := &User{}
    err := db.QueryRowContext(ctx, query, userID).Scan(&user.ID, &user.Name)
    return user, err
}
```

### Context Patterns
```go
// Cancellation propagation
func processData(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()  // context.Canceled or context.DeadlineExceeded
    default:
        // do work
    }
    
    // Pass context to sub-functions
    return subTask(ctx)
}

// Timeout pattern
func makeRequest(ctx context.Context, url string) (*Response, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    
    return resp, nil
}

// Goroutine with context
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case <-ctx.Done():
            return
        case job := <-jobs:
            // process job
        }
    }
}
```

---

## 16. JSON

### JSON Encoding
```go
import "encoding/json"

type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Age       int       `json:"age,omitempty"`      // omit if zero value
    Password  string    `json:"-"`                  // never marshal
    CreatedAt time.Time `json:"created_at"`
}

// Marshal to JSON
user := User{ID: 1, Name: "Tue", Email: "tue@example.com"}
jsonData, err := json.Marshal(user)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(jsonData))

// Marshal with indentation
jsonData, err := json.MarshalIndent(user, "", "  ")
```

### JSON Decoding
```go
// Unmarshal from JSON
jsonStr := `{"id":1,"name":"Tue","email":"tue@example.com"}`
var user User
err := json.Unmarshal([]byte(jsonStr), &user)
if err != nil {
    log.Fatal(err)
}

// Decode from io.Reader
decoder := json.NewDecoder(r.Body)
var user User
err := decoder.Decode(&user)
```

### JSON with HTTP
```go
// Send JSON response
func sendJSON(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

// Gin framework
func handler(c *gin.Context) {
    user := User{ID: 1, Name: "Tue"}
    c.JSON(200, user)
}

// Parse JSON request
func createUser(c *gin.Context) {
    var req CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // use req
}
```

### Custom JSON Marshaling
```go
type CustomTime time.Time

func (ct CustomTime) MarshalJSON() ([]byte, error) {
    t := time.Time(ct)
    return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02"))), nil
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
    str := string(data[1 : len(data)-1])  // remove quotes
    t, err := time.Parse("2006-01-02", str)
    if err != nil {
        return err
    }
    *ct = CustomTime(t)
    return nil
}
```

---

## 17. BACKEND PATTERNS

### Repository Pattern
```go
type UserRepository interface {
    Create(ctx context.Context, user *User) error
    GetByID(ctx context.Context, id int) (*User, error)
    GetByEmail(ctx context.Context, email string) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id int) error
    List(ctx context.Context, limit, offset int) ([]*User, error)
}

type postgresUserRepository struct {
    db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) UserRepository {
    return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) Create(ctx context.Context, user *User) error {
    query := `
        INSERT INTO users (name, email, password, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
    return r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password, time.Now()).
        Scan(&user.ID)
}

func (r *postgresUserRepository) GetByID(ctx context.Context, id int) (*User, error) {
    user := &User{}
    query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
    err := r.db.QueryRowContext(ctx, query, id).
        Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
    if err == sql.ErrNoRows {
        return nil, ErrUserNotFound
    }
    return user, err
}
```

### Service Layer Pattern
```go
type UserService interface {
    Register(ctx context.Context, req *RegisterRequest) (*User, error)
    Login(ctx context.Context, email, password string) (string, error)
    GetProfile(ctx context.Context, userID int) (*User, error)
    UpdateProfile(ctx context.Context, userID int, req *UpdateProfileRequest) error
}

type userService struct {
    repo   UserRepository
    hasher PasswordHasher
    jwt    JWTService
}

func NewUserService(repo UserRepository, hasher PasswordHasher, jwt JWTService) UserService {
    return &userService{
        repo:   repo,
        hasher: hasher,
        jwt:    jwt,
    }
}

func (s *userService) Register(ctx context.Context, req *RegisterRequest) (*User, error) {
    // Validate
    if err := req.Validate(); err != nil {
        return nil, err
    }
    
    // Check if email exists
    existing, _ := s.repo.GetByEmail(ctx, req.Email)
    if existing != nil {
        return nil, ErrEmailAlreadyExists
    }
    
    // Hash password
    hashedPassword, err := s.hasher.Hash(req.Password)
    if err != nil {
        return nil, err
    }
    
    // Create user
    user := &User{
        Name:     req.Name,
        Email:    req.Email,
        Password: hashedPassword,
    }
    
    if err := s.repo.Create(ctx, user); err != nil {
        return nil, err
    }
    
    return user, nil
}
```

### Middleware Pattern
```go
type Middleware func(http.Handler) http.Handler

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        next.ServeHTTP(w, r)
        
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        
        if token == "" {
            http.Error(w, "Unauthorized", 401)
            return
        }
        
        // Validate token and set user in context
        userID, err := validateToken(token)
        if err != nil {
            http.Error(w, "Invalid token", 401)
            return
        }
        
        ctx := context.WithValue(r.Context(), "userID", userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Chain middlewares
func ChainMiddleware(h http.Handler, middlewares ...Middleware) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        h = middlewares[i](h)
    }
    return h
}

// Usage
handler := ChainMiddleware(
    myHandler,
    LoggingMiddleware,
    AuthMiddleware,
    RateLimitMiddleware,
)
```

### Dependency Injection
```go
type Server struct {
    router      *gin.Engine
    db          *sql.DB
    userRepo    UserRepository
    userService UserService
    userHandler *UserHandler
    config      *Config
}

func NewServer(config *Config) (*Server, error) {
    // Database
    db, err := sql.Open("postgres", config.DatabaseURL)
    if err != nil {
        return nil, err
    }
    
    // Repositories
    userRepo := NewPostgresUserRepository(db)
    
    // Services
    hasher := NewBcryptHasher()
    jwt := NewJWTService(config.JWTSecret)
    userService := NewUserService(userRepo, hasher, jwt)
    
    // Handlers
    userHandler := NewUserHandler(userService)
    
    // Router
    router := gin.Default()
    
    server := &Server{
        router:      router,
        db:          db,
        userRepo:    userRepo,
        userService: userService,
        userHandler: userHandler,
        config:      config,
    }
    
    server.setupRoutes()
    
    return server, nil
}

func (s *Server) setupRoutes() {
    api := s.router.Group("/api/v1")
    {
        users := api.Group("/users")
        {
            users.POST("/register", s.userHandler.Register)
            users.POST("/login", s.userHandler.Login)
            users.GET("/profile", AuthMiddleware(), s.userHandler.GetProfile)
        }
    }
}

func (s *Server) Start() error {
    return s.router.Run(":" + s.config.Port)
}
```

### Configuration Pattern
```go
type Config struct {
    Port        string
    DatabaseURL string
    JWTSecret   string
    RedisURL    string
}

func LoadConfig() (*Config, error) {
    // From environment variables
    config := &Config{
        Port:        getEnv("PORT", "8080"),
        DatabaseURL: getEnv("DATABASE_URL", ""),
        JWTSecret:   getEnv("JWT_SECRET", ""),
        RedisURL:    getEnv("REDIS_URL", ""),
    }
    
    // Validate
    if config.DatabaseURL == "" {
        return nil, errors.New("DATABASE_URL is required")
    }
    
    return config, nil
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}
```

### Error Response Pattern
```go
type ErrorResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Details interface{} `json:"details,omitempty"`
}

func HandleError(c *gin.Context, err error) {
    switch {
    case errors.Is(err, ErrNotFound):
        c.JSON(404, ErrorResponse{
            Code:    "NOT_FOUND",
            Message: "Resource not found",
        })
    case errors.Is(err, ErrUnauthorized):
        c.JSON(401, ErrorResponse{
            Code:    "UNAUTHORIZED",
            Message: "Authentication required",
        })
    case errors.Is(err, ErrValidation):
        c.JSON(400, ErrorResponse{
            Code:    "VALIDATION_ERROR",
            Message: "Invalid input",
            Details: err.Error(),
        })
    default:
        c.JSON(500, ErrorResponse{
            Code:    "INTERNAL_ERROR",
            Message: "Internal server error",
        })
    }
}
```

---

## KẾT LUẬN

Đây là tổng hợp khá đầy đủ các syntax và pattern thường dùng trong Go backend development. Một số tips quan trọng:

1. **Luôn check error** - Go không có exception, phải check error sau mỗi operation
2. **Sử dụng interfaces** cho abstraction và testing
3. **Context** cho timeout và cancellation
4. **Goroutines và channels** cho concurrency
5. **Defer** cho cleanup resources
6. **Pointer vs Value** - hiểu khi nào dùng pointer receiver
7. **Package structure** - tổ chức code theo domain-driven design

Để học sâu hơn, nên:
- Đọc "Effective Go" (official docs)
- Practice với real projects
- Tham khảo source code của các open-source Go projects
- Làm coding challenges trên LeetCode/HackerRank bằng Go

Good luck với journey học Go! 🚀
