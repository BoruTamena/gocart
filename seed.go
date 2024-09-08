package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	// Update with your PostgreSQL connection details
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "root"
	dbName     = "cart_db"
)

func Seed_Data() {
	// Create connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")

	// Seed users
	_, err = db.Exec(`INSERT INTO "user" (username, email, password) 
                      VALUES ($1, $2, $3), ($4, $5, $6)`,
		"john_doe", "john@example.com", "password1", // User 1
		"jane_doe", "jane@example.com", "password2") // User 2
	if err != nil {
		log.Fatalf("Error seeding users: %v", err)
	}
	fmt.Println("Users seeded successfully")

	// Seed discounts
	_, err = db.Exec(`INSERT INTO "discount" (name, description, percent)
                      VALUES ($1, $2, $3), ($4, $5, $6)`,
		"Summer Sale", "15% off on all items", 15, // Discount 1
		"Black Friday", "50% off on electronics", 50) // Discount 2
	if err != nil {
		log.Fatalf("Error seeding discounts: %v", err)
	}
	fmt.Println("Discounts seeded successfully")

	// Seed products
	_, err = db.Exec(`INSERT INTO "product" (name, description, sku, category, price, discount_id)
                      VALUES ($1, $2, $3, $4, $5, $6), ($7, $8, $9, $10, $11, $12)`,
		"Laptop", "Gaming Laptop", 1001, "Electronics", 1200.50, 2, // Product 1
		"Phone", "Smartphone", 1002, "Electronics", 800.99, 1) // Product 2
	if err != nil {
		log.Fatalf("Error seeding products: %v", err)
	}
	fmt.Println("Products seeded successfully")

	// Seed shopping session
	_, err = db.Exec(`INSERT INTO "shopping_session" (user_id, total)
                      VALUES ($1, $2), ($3, $4)`,
		1, 1500.50, // Session 1 for User 1
		2, 800.99) // Session 2 for User 2
	if err != nil {
		log.Fatalf("Error seeding shopping sessions: %v", err)
	}
	fmt.Println("Shopping sessions seeded successfully")

	// Seed cart items
	_, err = db.Exec(`INSERT INTO "cart_item" (session_id, product_id, quantity)
                      VALUES ($1, $2, $3), ($4, $5, $6)`,
		1, 1, 1, // Cart item 1
		2, 2, 2) // Cart item 2
	if err != nil {
		log.Fatalf("Error seeding cart items: %v", err)
	}
	fmt.Println("Cart items seeded successfully")

	// Seed order details
	_, err = db.Exec(`INSERT INTO "order_details" (user_id, total, payment_id)
                      VALUES ($1, $2, $3), ($4, $5, $6)`,
		1, 1500.50, 1, // Order 1 for User 1
		2, 800.99, 2) // Order 2 for User 2
	if err != nil {
		log.Fatalf("Error seeding order details: %v", err)
	}
	fmt.Println("Order details seeded successfully")

	// Seed payment
	_, err = db.Exec(`INSERT INTO "payment" (order_id, amount, status)
                      VALUES ($1, $2, $3), ($4, $5, $6)`,
		1, 1500.50, "completed", // Payment 1 for Order 1
		2, 800.99, "pending") // Payment 2 for Order 2
	if err != nil {
		log.Fatalf("Error seeding payments: %v", err)
	}
	fmt.Println("Payments seeded successfully")

	// Seed order items
	_, err = db.Exec(`INSERT INTO "order_items" (order_id, product_id)
                      VALUES ($1, $2), ($3, $4)`,
		1, 1, // Order item 1
		2, 2) // Order item 2
	if err != nil {
		log.Fatalf("Error seeding order items: %v", err)
	}
	fmt.Println("Order items seeded successfully")
}
