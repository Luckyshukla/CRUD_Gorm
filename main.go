package main

import (
	"github.com/gin-gonic/gin"
	"main.go/Controller"
	"main.go/Models"
)

func main() {
	r := gin.Default()

	Models.ConnectDataBase()
	r.GET("/books", Controller.FindBooks)
	r.POST("/books", Controller.CreateBook)
	r.POST("/books/:id", Controller.FindBook)
	r.PATCH("/books/:id", Controller.UpdateBook)
	r.DELETE("/books/:id", Controller.DeleteBook)
	r.Run()
}
