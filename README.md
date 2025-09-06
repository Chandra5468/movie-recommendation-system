# Go Movie API

## Overview
The Go Movie API is a simple RESTful API built with Go that allows users to fetch and manage movie records stored in a PostgreSQL database. This project demonstrates the use of a clean architecture pattern, separating concerns into different layers: repository, service, and handler.

## Project Structure
```
go-movie-api
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handler
│   │   └── handler.go   # HTTP request handlers
│   ├── service
│   │   └── service.go   # Business logic for movie operations
│   └── repository
│       └── repository.go # Database interaction layer
├── pkg
│   └── db
│       └── postgres.go   # Database connection logic
├── go.mod                # Module definition and dependencies
├── go.sum                # Dependency checksums
└── README.md             # Project documentation
```

## Setup Instructions

1. **Clone the Repository**
   ```bash
   git clone https://github.com/yourusername/go-movie-api.git
   cd go-movie-api
   ```

2. **Install Dependencies**
   Ensure you have Go installed, then run:
   ```bash
   go mod tidy
   ```

3. **Set Up PostgreSQL**
   - Install PostgreSQL and create a database for the API.
   - Update the database connection settings in `pkg/db/postgres.go`.

4. **Run the Application**
   ```bash
   go run cmd/main.go
   ```

## API Usage
- **GET /movies**: Fetch all movies from the database.
- **POST /movies**: Add a new movie to the database. Requires a JSON body with movie details.

## Contributing
Feel free to fork the repository and submit pull requests for any improvements or features you would like to add.

## License
This project is licensed under the MIT License. See the LICENSE file for details.