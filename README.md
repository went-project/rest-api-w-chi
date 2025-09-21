# Went Chi API Template

A comprehensive REST API template built with Go programming language and Chi router, designed for rapid development of scalable web services.

## 🚀 Features

- � **Chi Router**: Lightweight, fast, and composable HTTP router
- 🛡️ **CORS Support**: Cross-origin resource sharing middleware
- � **Swagger Integration**: Auto-generated API documentation
- 🗄️ **GORM ORM**: Powerful object-relational mapping
- 🔐 **JWT Authentication**: JSON Web Token-based authentication
- 🐳 **Docker Ready**: Container support for deployment
- 🔧 **Multi-Environment**: Development, test, and production configurations
- 📊 **Database Support**: SQLite, PostgreSQL, and MySQL
- 🧪 **Graceful Shutdown**: Proper server shutdown handling
- 📝 **Structured Logging**: Request logging middleware

## 🛠️ Technology Stack

| Component | Technology | Version |
|-----------|------------|---------|
| **HTTP Router** | Chi | v5.2.3 |
| **ORM** | GORM | v1.31.0 |
| **Database Drivers** | SQLite, PostgreSQL, MySQL | Latest |
| **Documentation** | Swagger/OpenAPI | v3 |
| **Authentication** | JWT | - |
| **Configuration** | Environment Variables | - |
| **Validation** | Go Playground Validator | v10 |

## 📋 Prerequisites

- Go 1.24.2 or higher
- Database server (PostgreSQL, MySQL, or SQLite)
- Git

## 🚀 Quick Start

### 1. Clone the Repository
```bash
git clone <repository-url>
cd rest-api-w-chi
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Setup Environment
```bash
cp .env.example .env.local
# Edit .env.local with your configuration
```

### 4. Run the Application
```bash
go run main.go
```

The server will start at `http://localhost:8080`

## ⚙️ Configuration

This project supports multiple environments through separate `.env` files. Use the `APP_ENV` environment variable to specify which configuration file to load.

### Supported Environments

| Environment | ENV File | Command |
|-------------|----------|---------|
| **Local** | `.env.local` | `go run main.go` |
| **Development** | `.env.development` | `APP_ENV=development go run main.go` |
| **Test** | `.env.test` | `APP_ENV=test go run main.go` |
| **Production** | `.env` | `APP_ENV=production go run main.go` |

### Environment Variables

Create your environment file with the following variables:

```env
# Server Configuration
PORT=8080
APP_ENV=local

# Database Configuration
DB_DRIVER=sqlite          # sqlite, postgres, mysql
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASS=your_password
DB_NAME=your_database

# JWT Configuration
JWT_SECRET=your_super_secret_key

# CORS Configuration
CORS_ALLOWED_ORIGINS=*
CORS_ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
CORS_ALLOWED_HEADERS=Origin,Content-Type,Authorization
```

### Database Configuration Examples

#### SQLite (Default)
```env
DB_DRIVER=sqlite
DB_NAME=./db.sqlite
```

#### PostgreSQL
```env
DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=password
DB_NAME=myapp
```

#### MySQL
```env
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=password
DB_NAME=myapp
```

### Usage Examples

```bash
# Run in local environment (.env.local)
go run main.go

# Run in development environment
APP_ENV=development go run main.go

# Run in test environment
APP_ENV=test go run main.go

# Run in production environment
APP_ENV=production go run main.go
```

## 🐳 Docker Usage

### Build Docker Image
```bash
docker build -t went-chi-api .
```

### Run with Docker
```bash
# Development environment
docker run -e APP_ENV=development -p 8080:8080 went-chi-api

# Production environment
docker run -e APP_ENV=production -p 8080:8080 went-chi-api
```

### Docker Compose
```yaml
version: '3.8'
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=development
    volumes:
      - .:/app
```

## 📡 API Documentation

### Base URL
```
http://localhost:8080
```

### Health Check
```http
GET /ping
```
**Response:**
```json
{
  "status": "pong"
}
```

### User Management API

#### Get All Users
```http
GET /users
```
**Response:**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2025-09-21T10:00:00Z",
    "updated_at": "2025-09-21T10:00:00Z"
  }
]
```

#### Get User by ID
```http
GET /users/{id}
```
**Parameters:**
- `id` (path, required): User ID

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2025-09-21T10:00:00Z",
  "updated_at": "2025-09-21T10:00:00Z"
}
```

#### Create User
```http
POST /users
Content-Type: application/json
```
**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com"
}
```
**Response:** `201 Created`

#### Update User
```http
PUT /users/{id}
Content-Type: application/json
```
**Request Body:**
```json
{
  "name": "John Updated",
  "email": "john.updated@example.com"
}
```
**Response:** `200 OK`

#### Delete User
```http
DELETE /users/{id}
```
**Response:**
```json
{
  "message": "User deleted"
}
```

### Interactive API Documentation
Visit `http://localhost:8080/swagger/` for interactive Swagger UI documentation.

## 🏗️ Project Structure

```
.
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
├── README.md              # Project documentation
├── LICENSE                # License file
├── db.sqlite              # SQLite database file
├── app/
│   ├── controllers/       # HTTP request handlers
│   │   └── user_controller.go
│   └── models/           # Database models
│       └── user_model.go
├── config/
│   └── config.go         # Configuration management
├── docs/                 # Swagger documentation
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── config/          # Internal configuration
│   │   └── config.go
│   ├── middlewares/     # HTTP middlewares
│   │   ├── Auth.go
│   │   └── Cors.go
│   ├── providers/       # Service providers
│   │   └── database_provider.go
│   ├── responses/       # Response structures
│   │   └── responses.go
│   └── utils/          # Utility functions
│       └── Hash.go
├── pkg/                # Public packages
└── routes/             # Route definitions
    ├── main_router.go
    └── user_router.go
```

## 🧪 Testing

### Run Tests
```bash
go test ./...
```

### Run Tests with Coverage
```bash
go test -cover ./...
```

### Run Specific Test
```bash
go test ./app/controllers -v
```

## 🚀 Development

### Build the Application
```bash
go build -o bin/went-chi-api .
```

### Run the Application
```bash
go run main.go
```

### Generate Swagger Documentation
```bash
swag init
```

### Format Code
```bash
go fmt ./...
```

### Lint Code
```bash
golangci-lint run
```

## 🌟 Why Chi Router?

### Advantages over Gin

| Feature | Chi | Gin |
|---------|-----|-----|
| **Performance** | Faster, lightweight | Good performance |
| **Standard Compliance** | Uses standard `http.Handler` | Custom context |
| **Middleware Ecosystem** | Rich, standard-compatible | Limited to Gin ecosystem |
| **URL Parameters** | Clean `{param}` syntax | `:param` syntax |
| **Dependencies** | Minimal | More dependencies |
| **Learning Curve** | Familiar to Go developers | Framework-specific |

### Migration from Gin to Chi

This template was migrated from Gin to Chi router. Key changes include:

| Component | Gin | Chi |
|-----------|-----|-----|
| **Router** | `gin.Engine` | `chi.Mux` |
| **Context** | `gin.Context` | `http.ResponseWriter` + `http.Request` |
| **Middleware** | `gin.HandlerFunc` | `func(http.Handler) http.Handler` |
| **URL Parameters** | `c.Param("id")` | `chi.URLParam(r, "id")` |
| **JSON Response** | `c.JSON(200, data)` | `render.JSON(w, r, data)` |
| **Request Binding** | `c.ShouldBindJSON(&user)` | `json.NewDecoder(r.Body).Decode(&user)` |

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go conventions and best practices
- Write tests for new features
- Update documentation for API changes
- Use meaningful commit messages
- Ensure code passes all tests and linting

## 📝 API Response Format

### Success Response
```json
{
  "data": {},
  "status": "success"
}
```

### Error Response
```json
{
  "error": "Error message",
  "status": "error"
}
```

## 🔐 Authentication

This template includes JWT authentication middleware. To use protected endpoints:

1. Implement login endpoint to generate JWT token
2. Include token in Authorization header: `Bearer <token>`
3. Middleware will validate token automatically

Example:
```http
GET /protected-endpoint
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

## 📚 Additional Resources

- [Chi Router Documentation](https://github.com/go-chi/chi)
- [GORM Documentation](https://gorm.io/docs/)
- [Swagger Documentation](https://swagger.io/docs/)
- [Go Best Practices](https://golang.org/doc/effective_go.html)

## 📄 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Chi Router](https://github.com/go-chi/chi) for the excellent HTTP router
- [GORM](https://gorm.io/) for the powerful ORM
- [Swagger](https://swagger.io/) for API documentation
- Go community for amazing tools and libraries