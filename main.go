package main

import (
	library_pkg "bookstore/library"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, playground")

	library := library_pkg.NewLibrary(library_pkg.NewMysqlRepository())

	router := gin.Default()
	api := router.Group("/api")

	api.GET("/", hello)
	api.POST("/books", addBook(library))
	api.GET("/books/:id", getBook(library))
	api.DELETE("/books/:id", deleteBook(library))
	api.PUT("/books/:id", updateBook(library))

	router.Run(":8080")
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func addBook(library *library_pkg.Library) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var book library_pkg.Book
		err := c.BindJSON(&book)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		id, err := library.Add(book)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(202, gin.H{
			"id": id,
		})
	}
	return gin.HandlerFunc(fn)
}

func updateBook(library *library_pkg.Library) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		var book library_pkg.Book
		err := c.BindJSON(&book)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		id, err = library.Modify(id, book)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(202, gin.H{
			"id": id,
		})
	}
	return gin.HandlerFunc(fn)
}

func getBook(library *library_pkg.Library) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		book, err := library.Find(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, book)
	}
	return gin.HandlerFunc(fn)
}

func deleteBook(library *library_pkg.Library) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		err := library.Delete(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(202, gin.H{
			"id": id,
		})
	}
	return gin.HandlerFunc(fn)
}
