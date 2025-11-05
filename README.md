# Go Simple API CRUD

This is a simple Go application that demonstrates how to create a RESTful API with CRUD operations using the Gorilla Mux router
and a PostgresSQL database.

## Technologies and Libraries Used

- **Go** - Programming language
- **Gorilla Mux** - HTTP router and URL matcher for building Go web servers
- **PGX** - PostgreSQL driver and toolkit for Go
- **Squirrel** - SQL query builder for Go
- **Testify** - Testing toolkit for Go
- **golang-migrate/migrate** - Database migration tool for Go

## Prerequisites

- Go 1.25 or higher
- PostgreSQL database
- Git

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd go-simple-api-crud
```

2. Install dependencies:

```bash
go mod tidy
```

3. Set up your PostgreSQL database and update the connection string in your configuration.

4. Run database migrations:

```bash
migrate -path db/migrations -database "postgresql://username:password@localhost/dbname?sslmode=disable" up
```

Migration files are located in `db/migrations/` directory.

## Running the Application

1. Start the server:

```bash
go run cmd/api/main.go
```

2. The API will be available at `http://localhost:8080`

## API Documentation

The API documentation is available in OpenAPI format at `api/api_user_openapi.json`. You can import this file into tools like Postman or Swagger UI for
interactive API documentation.