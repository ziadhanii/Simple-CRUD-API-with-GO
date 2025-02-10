package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"go-crud-api/config"
	"go-crud-api/models"
	"go-crud-api/routes"
)

func main() {
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(&models.Hotel{})
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	router := gin.Default()

	router.SetTrustedProxies(nil)

	routes.RegisterRoutes(router)

	log.Println("🚀 Server is running on http://localhost:8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
