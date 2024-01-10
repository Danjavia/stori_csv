package main

import (
    "database/sql"
    "log"

    "github.com/danjavia/stori_csv/cmd/api"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

func main() {
    log.Printf("Starting application")

    // Connect to Postgres database
    db, err := sql.Open("postgres", "postgres://danjavia:djvx2024@5432:5432/stori")
    if err != nil {
        log.Fatal(err) // Handle error gracefully
    }

    r := gin.Default()

    r.Use(cors.Default())

    // Available services (using Postgres adapter functions)
    r.POST("/transactions", api.CheckTransactions(db))
    r.POST("/send-email", api.SendEmail(db))

    // Start the Gin server
    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}