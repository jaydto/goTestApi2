package routinggingonic

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)
type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "Pamela my girl", Author: "Mark Rolensco ", Quantity: 20},
	{ID: "2", Title: "Noise", Author: "Kim Li ", Quantity: 20},
	{ID: "3", Title: "The Tribunal", Author: "Jim Maina ", Quantity: 20},
}

func getGooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

}

func createBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, books)
}

func getGookById(id string) (*Book, error) {

	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book not found")

}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getGookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)

}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameters"})
		return
	}
	book, err := getGookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found"})
		return

	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not Available"})
		return
	}
	book.Quantity -= 1

	c.IndentedJSON(http.StatusOK, book)

}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameters"})
		return
	}
	book, err := getGookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found"})
		return

	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not Available"})
		return
	}
	book.Quantity += 1

	c.IndentedJSON(http.StatusOK, book)

}

func Routing(){
	router := gin.Default()
	router.GET("/books", getGooks)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/returnBook", returnBook)
	router.GET("/books/:id", bookById)

	router.Run("localhost:8000")
}