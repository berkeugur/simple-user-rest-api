package main

import (
	"log"
	"simple-user-rest-api/controllers"
	"simple-user-rest-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// GORM ile veritabanı bağlantısını yap
	db, err := gorm.Open(sqlite.Open("example.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Veritabanı tablolarını migrate et
	db.AutoMigrate(&models.User{})

	// Gin router oluştur
	router := gin.Default()

	// Controller'dan route'ları ayarla
	controllers.RegisterRoutes(router, db)

	// Gin server'ı 8080 portunda başlat
	router.Run(":8080")
}
