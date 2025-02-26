# Go REST API - Anıl Yağız

This project is a Book Library Management System REST API developed using the Go programming language and Gin-Gonic framework.

## Features

- Full CRUD operations for authors, books, and reviews
- Relational database integration (PostgreSQL + GORM)
- Validation and error handling
- Pagination support
- Swagger API documentation
- Containerization with Docker and Docker Compose

## Getting Started

### Requirements

- Go 1.18 or higher
- PostgreSQL 
- Docker and Docker Compose (optional)

### Local Installation

1. Clone the repository:

```bash
git clone https://github.com/MentalArts/go-rest-api-anil-yagiz.git
cd go-rest-api-anil-yagiz
```

2. Install dependencies:

```bash
go mod download
```

3. Edit the `.env` file:

```
# Database Configuration
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=bookstore
DB_PORT=5432
DB_SSLMODE=disable

# API Configuration
API_PORT=8000
GIN_MODE=debug
```

4. Run the application:

```bash
go run main.go
```

### Running with Docker

```bash
docker-compose up -d
```

## API Documentation

Swagger documentation can be accessed from the running API at:

```
http://localhost:8000/swagger/index.html
```

## API Endpoints

### Authors

- `GET /api/v1/authors` - List all authors (with pagination)
- `GET /api/v1/authors/:id` - Get author details (with books)
- `POST /api/v1/authors` - Create new author
- `PUT /api/v1/authors/:id` - Update author
- `DELETE /api/v1/authors/:id` - Delete author

### Books

- `GET /api/v1/books` - List all books (with pagination)
- `GET /api/v1/books/:id` - Get book details (with author and reviews)
- `POST /api/v1/books` - Create new book
- `PUT /api/v1/books/:id` - Update book
- `DELETE /api/v1/books/:id` - Delete book

### Reviews

- `GET /api/v1/books/:id/reviews` - Get all reviews for a book
- `POST /api/v1/books/:id/reviews` - Add review to a book
- `PUT /api/v1/reviews/:id` - Update review
- `DELETE /api/v1/reviews/:id` - Delete review

## Project Structure

```
.
├── docker-compose.yaml    # Docker Compose configuration
├── Dockerfile            # Dockerfile for API
├── docs/                 # Swagger documentation
├── dto/                  # Data transfer objects
├── go.mod               # Go module definition
├── go.sum               # Go dependency versions
├── handlers/            # API endpoint handlers
├── main.go              # Main application entry point
├── models/              # Database models
├── README.md            # Project documentation
└── utils/               # Helper functions
```

---


# Project: Building a Simple REST API with Go and Gin-Gonic

# DEADLINE: 09.03.2025

# CONTACT: oguzhan.durgut@btsgrp.com

## Description

The goal of this project is to develop a basic REST API using the Go programming language and the Gin-Gonic framework. The API should be capable of interacting with a database (e.g. PostgreSQL and SQLite). Additionally, the project should utilize containerization to develop the application within a containerized environment.

## Requirements

### Core Technologies

-   **Programming Language**: The project must be implemented using the Go programming language.
-   **Framework**: Utilize the Gin-Gonic framework for building the REST API.

### Functionality

-   **API Functionality**: Build a Book Library Management System with the following entities:

    1. Books (title, author, ISBN, publication_year, description)
    2. Authors (name, biography, birth_date)
    3. Reviews (rating, comment, date_posted)

    Relationships:

    -   One Author can have many Books (1:N)
    -   One Book can have many Reviews (1:N)
    -   Books and Authors have a bidirectional relationship

Required API Endpoints:

-   Books:

    -   GET /api/v1/books (list all books with pagination)
    -   GET /api/v1/books/:id (get book details with author and reviews)
    -   POST /api/v1/books (create new book)
    -   PUT /api/v1/books/:id (update book)
    -   DELETE /api/v1/books/:id (delete book)

-   Authors:

    -   GET /api/v1/authors (list all authors with their books)
    -   GET /api/v1/authors/:id (get author details)
    -   POST /api/v1/authors (create new author)
    -   PUT /api/v1/authors/:id (update author)
    -   DELETE /api/v1/authors/:id (delete author)

-   Reviews:

    -   GET /api/v1/books/:id/reviews (get all reviews for a book)
    -   POST /api/v1/books/:id/reviews (add review to a book)
    -   PUT /api/v1/reviews/:id (update review)
    -   DELETE /api/v1/reviews/:id (delete review)

-   **Database Integration**: Implement the system using PostgreSQL with GORM. The database schema should include proper foreign key relationships and constraints.

### Development & Deployment

-   **Containerization**: Develop both the REST API and the SQL service within a containerized environment using Docker. Provide necessary Dockerfiles and a docker-compose.yaml file for containerization and deployment.
-   **Swagger Documentation**: Generate auto documentation for the API endpoints using Swagger.
-   **Version Control**: Create a new repository under the [MentalArts organization](https://github.com/MentalArts) on GitHub. All development work should be committed to this repository. The repository name should follow the format: `go-rest-api-name-surname`.

### Note

-   **Containerization Challenge**: During the process of containerization using Docker Compose, you might encounter challenges or issues. Encourage them to explore and solve these problems as part of the learning process.

## Extra Features

The following features are optional enhancements that can be implemented to extend the core functionality:

### Authentication & Authorization

-   Implement JWT-based authentication
-   Protected routes requiring authentication
-   Role-based access control (Admin, User)
-   User registration and login endpoints:
    -   POST /api/v1/auth/register
    -   POST /api/v1/auth/login
    -   POST /api/v1/auth/refresh-token

### Additional Enhancements

-   **Rate Limiting**: Implement request rate limiting per user/IP
-   **Caching**: Add Redis caching for frequently accessed data
-   **Logging**: Implement structured logging with levels (info, error, debug)
-   **Input Validation**: Add request payload validation
-   **Error Handling**: Implement consistent error responses
-   **Testing**:
    -   Unit tests for business logic
    -   Integration tests for API endpoints
    -   Load testing with tools like k6

### Monitoring & Observability

-   **Metrics**: Implement Prometheus metrics
-   **Health Checks**: Add health check endpoints
-   **Tracing**: Implement distributed tracing with Jaeger
-   **Monitoring Dashboard**: Set up Grafana for visualization

## Resources

### Go Resources

-   [Tour of Go](https://tour.golang.org/)
-   [Practical Go Lessons Book](https://www.practical-go-lessons.com/)
-   [Go Installation](https://golang.org/doc/install)
-   [Go Installation (Older Versions)](https://golang.org/dl/)

### Docker Resources

-   [Docker Installation](https://docs.docker.com/get-docker/)
-   [OrbStack](https://orbstack.dev/) (Docker Desktop Alternative for MacOS - Optional)

### Framework & Tools

-   [Gin-Gonic Documentation](https://gin-gonic.com/) : REST API Framework
-   [GORM Documentation](https://gorm.io/) : ORM for Golang
-   [Docker Documentation](https://docs.docker.com/) : Containerization
-   [Swagger](https://swagger.io/) : Auto Documentation Tool
-   [Gin-Swagger Github](https://github.com/swaggo/gin-swagger) : Swagger for Gin-Gonic
-   [Swag](https://github.com/swaggo/swag) : Swagger documentation generator for Golang
