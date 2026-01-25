# Gin Framework Task Management API

A simple REST API built with Gin framework using a JSON file as the database for data isolation.

## Project Structure

```
gin-framework/
├── database/
│   └── db.go           # JSON database operations
├── handlers/
│   └── task_handler.go # HTTP handlers for CRUD operations
├── models/
│   └── task.go         # Task data models
├── db.json             # JSON file database
├── main.go             # Application entry point
├── go.mod              # Go module definition
└── README.md           # This file
```

## Features

- RESTful API design
- CRUD operations for tasks
- JSON file as database (isolated and portable)
- Thread-safe database operations with mutex
- Input validation
- Proper error handling
- Clean architecture with separation of concerns

## Installation & Setup

1. Install dependencies:
```bash
cd gin-framework
go mod tidy
```

2. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Get All Tasks
```bash
GET /api/tasks
```

Example:
```bash
curl http://localhost:8080/api/tasks
```

Response:
```json
{
  "count": 3,
  "data": [
    {
      "id": 1,
      "title": "Learn Gin Framework",
      "description": "Study the basics of Gin web framework for Go",
      "completed": false,
      "created_at": "2026-01-25T10:00:00Z",
      "updated_at": "2026-01-25T10:00:00Z"
    }
  ]
}
```

### Get Task by ID
```bash
GET /api/tasks/:id
```

Example:
```bash
curl http://localhost:8080/api/tasks/1
```

### Create New Task
```bash
POST /api/tasks
```

Example:
```bash
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Task",
    "description": "Task description"
  }'
```

Response:
```json
{
  "message": "Task created successfully",
  "data": {
    "id": 4,
    "title": "New Task",
    "description": "Task description",
    "completed": false,
    "created_at": "2026-01-25T14:30:00Z",
    "updated_at": "2026-01-25T14:30:00Z"
  }
}
```

### Update Task
```bash
PUT /api/tasks/:id
```

Example (update specific fields):
```bash
curl -X PUT http://localhost:8080/api/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "completed": true
  }'
```

You can update any combination of:
- `title` (string)
- `description` (string)
- `completed` (boolean)

### Delete Task
```bash
DELETE /api/tasks/:id
```

Example:
```bash
curl -X DELETE http://localhost:8080/api/tasks/1
```

Response:
```json
{
  "message": "Task deleted successfully"
}
```

## Code Explanation

### Database Layer (`database/db.go`)
- Handles reading/writing JSON files
- Thread-safe operations using `sync.RWMutex`
- Automatically creates file if it doesn't exist

### Models (`models/task.go`)
- `Task`: Main task structure
- `CreateTaskInput`: Validation for creating tasks
- `UpdateTaskInput`: Partial update support using pointers

### Handlers (`handlers/task_handler.go`)
- `GetAllTasks`: Retrieve all tasks
- `GetTaskByID`: Get single task
- `CreateTask`: Create new task with auto-generated ID
- `UpdateTask`: Partial update support
- `DeleteTask`: Remove task by ID

### Main Application (`main.go`)
- Initialize database and handlers
- Setup routes with grouping
- Start HTTP server

## Testing

Test the API using curl, Postman, or any HTTP client.

Example test flow:
```bash
# 1. Get all tasks
curl http://localhost:8080/api/tasks

# 2. Create a new task
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Task","description":"Testing API"}'

# 3. Update the task
curl -X PUT http://localhost:8080/api/tasks/4 \
  -H "Content-Type: application/json" \
  -d '{"completed":true}'

# 4. Get specific task
curl http://localhost:8080/api/tasks/4

# 5. Delete the task
curl -X DELETE http://localhost:8080/api/tasks/4
```

## Key Concepts Demonstrated

1. **Gin Framework Basics**
   - Routing and route groups
   - Request binding and validation
   - JSON responses
   - Middleware (logger, recovery)

2. **Go Best Practices**
   - Package organization
   - Dependency injection
   - Error handling
   - Concurrency safety (mutex)

3. **REST API Design**
   - Proper HTTP methods
   - Status codes
   - Request/response structure
   - CRUD operations

4. **Data Persistence**
   - JSON file as database
   - Read/write operations
   - Auto-incrementing IDs
   - Partial updates

## Notes

- The `db.json` file is created automatically if it doesn't exist
- All database operations are thread-safe
- The API uses JSON for both request and response bodies
- Input validation is handled by Gin's binding tags
