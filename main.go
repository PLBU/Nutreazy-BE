package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"be-nutreazy/models"
	"be-nutreazy/routes"
	"be-nutreazy/utils"
)

func initDB() *gorm.DB {
	// Load environment variables from the .env file
	err := godotenv.Load()
	utils.LogFatal(err)
	
	db, err := gorm.Open(postgres.Open(os.Getenv("URL")), &gorm.Config{})
    utils.LogFatal(err)
	
	db.AutoMigrate(&models.User{})

	return db
}

func main() {
	db := initDB()
    
	router := gin.Default()

	routes.SetUpRoutes(router, db)
	router.Run(":8080")
}