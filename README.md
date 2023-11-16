# Go PostgreSQL Database Project

This is a simple Go project that demonstrates how to interact with a PostgreSQL database. The project includes functionality to create a product table, insert sample products, and generate random product data.

## Prerequisites

- Go installed on your machine
- PostgreSQL database server running
- `github.com/lib/pq` Go package for PostgreSQL

## Getting Started

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/your-username/go-postgresql-project.git
   ```

- cd go-postgresql-project
- connectionString := "postgres://your-username:your-password@localhost:5432/your-database?sslmode=disable"
- go run main.go

Project Structure:

main.go: Contains the main functionality of the project, including database connection, table creation, product insertion, and random data generation.

README.md: This file, providing information about the project.

Functions:

createProductTable(db, \*sql.DB): Creates a 'product' table in the PostgreSQL database.

insertProduct(db sql.DB, product Product) int: Inserts a product into the 'product' table and returns the generated ID.

generateProduct(db, \*sql.DB): Generates and inserts random products into the 'product' table.

generateRandomProduct() string: Generates a random product name from a predefined list.

generatePrice()float64: Generates a random price for a product.
