package Controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/Models"
)

func FindBooks(c *gin.Context) {
	var books Models.Book
	Models.DB.Find(&books)

	c.JSON(200, gin.H{"data": books})
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(c *gin.Context) {

	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := Models.Book{Title: input.Title, Author: input.Author}
	Models.DB.Create(&book)

	c.JSON(200, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	var book Models.Book

	if err := Models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(200, gin.H{"data": book})
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

//func UpdateBook(c *gin.Context) {
//	var book Models.Book
//	if err := Models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not find !"})
//		return
//	}
//validate input
//	var input UpdateBookInput
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(400, gin.H{"error": err.Error()})
//		return
//	}
//	Models.DB.Model(&book).Updates(input)
//
//	c.JSON(200, gin.H{"data": book})
//}

func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book Models.Book
	if err := Models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Models.DB.Table("books").Where("id = ?", c.Param("id")).Updates(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

//Delete data
func DeleteBook(c *gin.Context) {
	var book Models.Book
	if err := Models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Not found !"})
		return
	}
	Models.DB.Delete(&book)

	c.JSON(200, gin.H{"data": true})
}
