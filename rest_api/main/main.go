package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type Categories struct {
	Id   int    `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}

func (Categories) TableName() string { return "schema_thangbq.categories" }

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	fmt.Println("Successfully connected to the database!", db)

	// CRUD: Create, Read, Update, Delete
	// POST /api/v1/categories
	// GET /api/v1/categories
	// GET /api/v1/categories/:id
	// PUT /api/v1/categories/:id
	// DELETE /api/v1/categories/:id

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		item := v1.Group("/categories")
		{
			item.POST("", CreateCategories(db))
			item.GET("")
			item.GET("/:id")
			item.PUT("/:id")
			item.DELETE("")
		}
	}

	r.Run(":3000")

}

func CreateCategories(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var category Categories

		if err := c.ShouldBind(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&category).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   category,
		})
	}
}
