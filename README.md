# Project: Building a Simple REST API with Go and Gin-Gonic

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