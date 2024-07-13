package controllers

import (
	"net/http"
	"strconv"

	"simple-user-rest-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes, kullanıcı route'larını kayıt eder
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/users", func(c *gin.Context) {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	router.GET("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		ctx.JSON(http.StatusOK, user)
	})

	router.POST("/users", func(ctx *gin.Context) {
		var newUser models.User
		if err := ctx.ShouldBindJSON(&newUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&newUser).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "id": newUser.ID})
	})

	router.PUT("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var updateUser models.User
		if err := ctx.ShouldBindJSON(&updateUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := db.Model(&models.User{}).Where("id = ?", id).Updates(updateUser)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	})

	router.DELETE("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		if err := db.Delete(&models.User{}, id).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	})
}
