# Simple CRUD API with Go

This is a simple CRUD API built with **Go** using the **Gin** framework and **GORM** for database interactions. The project is designed as a basic demonstration of how to implement CRUD operations in Go.

## Features
- 📌 Create, Read, Update, and Delete (CRUD) operations for a **Hotel** entity.
- 🚀 Uses **Gin** for handling HTTP requests.
- 🗄️ Uses **GORM** as an ORM for SQL database operations.

## Technologies Used
- **Go** (Golang)
- **Gin** (Web framework)
- **GORM** (ORM for Go)
- **SQL Server**

## Installation & Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/ziadhanii/Simple-CRUD-API-with-GO.git
   cd Simple-CRUD-API-with-GO
   ```

2. **Initialize Go modules:**
   ```sh
   go mod tidy
   ```

3. **Run the application:**
   ```sh
   go run main.go
   ```

4. **Access:**
   Open your browser and go to:
   ```
   http://localhost:8080
   ```

## API Endpoints

| Method | Endpoint       | Description          |
|--------|--------------|----------------------|
| GET    | /hotels/      | Get all hotels       |
| GET    | /hotels/:id   | Get a hotel by ID    |
| POST   | /hotels/      | Create a new hotel   |
| PUT    | /hotels/:id   | Update hotel details |
| DELETE | /hotels/:id   | Delete a hotel       |

## Project Structure
```
Simple-CRUD-API-with-GO/
│── handlers/        # Contains API handlers
│── models/          # Database models
│── routes/          # Route definitions
│── main.go          # Entry point of the application
│── go.mod           # Go module file
│── README.md        # Project documentation
```

## Contribution
Feel free to fork this repository, open issues, and submit pull requests! 🚀

## License
This project is open-source and available under the **MIT License**.

