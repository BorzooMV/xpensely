# Xpensely

## Summary

Xpensely is a backend application that serves APIs for an expense tracker application. It is built using Go and provides CRUD operations for managing users and expenses.

## Technologies

- **Programming Language**: Go
- **Framework**: Echo
- **Database**: PostgreSQL
- **Caching**: Redis
- **Containerization**: Docker
- **Third-party Libraries**:
  - `godotenv`: For loading environment variables from a `.env` file.
  - `jwt-go`: For JWT authentication.

## API Resources

Xpensely exposes the following API endpoints for managing users and expenses:

- **Users**: CRUD operations for user management.
- **Expenses**: CRUD operations for expense management.
- **Auth**: Authentication endpoints for user login and token management.

### Database Schema

The application manages two tables:

1. **Users**: Contains user information.
2. **Expenses**: Contains expense records, with a `user_id` column that relates to the Users table.

### Authentication

JWT authentication is implemented. Users can obtain:

- `access_token` and `refresh_token` from the `/auth` endpoint by providing valid credentials.
- A new `access_token` using their `refresh_token` from the `/auth/refresh` endpoint.

## Getting Started

### Prerequisites

To run this application, you need to have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [Docker](https://www.docker.com/get-started) (and Docker Compose)

### Environment Variables

Create a `.env` file in the root of the project with the following variables:

```env
# Server Configuration
SERVER_LISTENING_PORT=

# PostgreSQL Configuration
POSTGRES_HOST=
POSTGRES_PASSWORD=
POSTGRES_PORT=
POSTGRES_USER=
POSTGRES_DB_NAME=

# Redis Configuration
REDIS_PASS=
REDIS_ADDRESS=
REDIS_PORT=

# JWT Configuration
JWT_SECRET=
JWT_REFRESH_SECRET=
```

### Redis Configuration

You can configure Redis by creating a `redis.conf` file at the root of the project. Set the following properties:

```conf
# Enable protected mode for security
protected-mode <is protected>

# Enable password authentication
requirepass "<password>"

# Set the port Redis will listen on
port <port>
```

### Docker Setup

To set up the necessary containers (Redis and PostgreSQL), run the following command in the root of the project:

```bash
docker compose up
```

### Makefile

The `Makefile` includes scripts for repetitive operations, which can be found in the `cmd/scripts/` directory.

#### Targets

- init-db: The most important one, you should run this before starting the server for the first time to create needed tables in the database.
- clean-db: This one clears all tables.
- seed-db-example: This one seeds database with sample data for testing purposes.

### Running the Application

The main entry point of the application is located at:

```
cmd/xpensely/main.go
```

You can run the application using:

```bash
go run cmd/xpensely/main.go
```

### Seeding the Database

Sample entries can be seeded into the database using JSON files located in the `assets/data` directory.
