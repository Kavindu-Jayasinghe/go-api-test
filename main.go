package main

import (
	"errors"
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
func BookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)

}
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil

		}

	}
	return nil, errors.New("book not found")

}
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id"})
		return
	}
	book, err := getBookById(id)

	if err != nil {

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return

	}

	if book.Quantity <= "0" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not found"})
		return
	}
	book.Quantity -= "1"

	c.IndentedJSON(http.StatusOK, book)

}
func returnBook(c *gin.Context) {

	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id"})
		return
	}
	book, err := getBookById(id)

	if err != nil {

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return

	}

	if book.Quantity <= "0" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not found"})
		return
	}
	book.Quantity += "1"

	c.IndentedJSON(http.StatusOK, book)

}
func main() {

	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", BookById)
	router.POST("books", createBooks)
	router.PATCH("/checkout", checkoutBook)
	router.Run("localhost:8090")

}
