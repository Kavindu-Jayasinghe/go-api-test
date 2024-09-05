package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity string `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "in search of the lost", Author: "Meracale prost", Quantity: "2"},
	{ID: "2", Title: "new 2", Author: "kavindu", Quantity: "5"},
	{ID: "3", Title: "Book 3", Author: "Author 3", Quantity: "6"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

}

func createBooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
func main() {

	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("books", createBooks)
	router.Run("localhost:8090")

}
