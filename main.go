package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Pages string `json:"pages"`
}

var bookshelf = []Book{
	// init data
	{
		Id:    "1",
		Name:  "Blue Bird",
		Pages: "500",
	},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bookshelf)
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	for i := range bookshelf {
		if bookshelf[i].id == id {
			c.IndentedJSON(http.StatusOK, bookshelf[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func addBook(c *gin.Context) {
	var book Book
	c.BindJSON(&book)
	bookshelf = append(bookshelf, book)
	c.IndentedJSON(http.StatusOK, bookshelf[len(bookshelf)-1])
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i := range bookshelf {
		if bookshelf[i].id == id {
			c.IndentedJSON(http.StatusOK, bookshelf[i])
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func updateBook(c *gin.Context) {
	var book Book
	id := c.Param("id")
	c.BindJSON(&book)
	for i := range bookshelf {
		if bookshelf[i].id == id {
			bookshelf[i] = book
			c.IndentedJSON(http.StatusOK, bookshelf[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func main() {
	r := gin.Default()
	r.RedirectFixedPath = true
	r.GET("/bookshelf", getBooks)
	r.GET("/bookshelf/:id", getBook)
	r.POST("/bookshelf", addBook)
	r.DELETE("/bookshelf/:id", deleteBook)
	r.PUT("/bookshelf/:id", updateBook)

	port := "8080"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	r.Run(":" + port)
}
