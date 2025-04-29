# Student API

A simple RESTful API for managing student records, built with Go and SQLite.

## Features

- Create a new student
- Retrieve a student by ID
- Retrieve a list of all students
- Uses SQLite as the database
- Configuration via YAML or environment variables
- Request Validation
- Graceful server shutdown

---

## Project Structure
Here is a `README.md` file for your project based on the provided code:

<pre> <code>```
Student-API/
├── cmd/
│   └── student-api/
│       └── main.go          # Entry point of the application
├── Internal/
│   ├── config/
│   │   └── config.go        # Configuration loader
│   ├── http/
│   │   └── handler/
│   │       └── student/
│   │           └── student.go # HTTP handlers for student operations
│   ├── sqlite/
│   │   └── sqlite.go        # SQLite database implementation
│   ├── storage/
│   │   └── storage.go       # Storage interface
│   ├── types/
│   │   └── types.go         # Data types (e.g., Student struct)
│   └── utiles/
│       └── response/
│           └── response.go  # Utility functions for HTTP responses
└── go.mod                   # Go module file
```</code> </pre>

## Code Overview

### 1. Configuration (`config/config.go`)
- Loads configuration from a YAML file or environment variables.
- Uses the `cleanenv` package for parsing.

### 2. SQLite Integration (`sqlite/sqlite.go`)
- Implements the Storage interface for database operations.
- Creates the `students` table if it doesn't exist.

### 3. HTTP Handlers (`http/handler/student/student.go`)
- Defines handlers for creating, retrieving, and listing students.
- Validates request payloads using `go-playground/validator`.

### 4. Response Utilities (`utiles/response/response.go`)
- Provides helper functions for writing JSON responses.
- Handles validation and general errors.

---

## Graceful Shutdown

The server listens for termination signals (e.g., `Ctrl+C`) and shuts down gracefully, ensuring all ongoing requests are completed.

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.
```

Let me know if you need any modifications or additional sections!Let me know if you need any modifications or additional sections!
