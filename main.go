package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/lib/pq"
)

// define a struct for a product.
type Product struct {
	Name      string
	Price     float64
	Avaliable bool
}

func main() {
	connectionString := "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable"
	// opening a connection to the postgres db.
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}
	// checking the connection to the database.
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	// defer closing the database connection until the main function exits.
	defer db.Close()
	// creating the product table in the database
	createProductTable(db)

	// creating a sample Product and inserting it into the database.
	product := Product{"Book", 18.99, true}
	pk := insertProduct(db, product)
	fmt.Printf("ID = %d\n", pk)
	// generating and inserting 10 random products into the database.
	generateProduct(db)

}

// function to generate and insert random products into the database.
func generateProduct(db *sql.DB) {
	for i := 0; i < 10; i++ {
		product := Product{generateRandomProduct(), generatePrice(), true}
		pk := insertProduct(db, product)
		fmt.Print(pk)
	}
}

// function to generate a random product name from a predefined list.
func generateRandomProduct() string {
	productNames := []string{"Book", "Beans", "Milk", "Sugar"}

	// Generate a random number between 0 and 3
	randomNumber := rand.Intn(4)
	return productNames[randomNumber]

}
func generatePrice() float64 {
	return rand.Float64()
}

// creating a product table.
func createProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS product(
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC (6,2) NOT NULL,
		avaliable BOOLEAN,
		created timestamp DEFAULT NOW()
	)`

	// database object to create the table.
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}

// function to insert a product into the 'product' table and return the generated ID.
func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, avaliable)
		VALUES ($1,$2,$3) RETURNING id`

	var pk int
	// executing the query and scanning the generated ID into the 'pk' variable.
	err := db.QueryRow(query, product.Name, product.Price, product.Avaliable).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
