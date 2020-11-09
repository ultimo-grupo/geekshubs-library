package main

import (
	"bookstore/api"
	library_pkg "bookstore/library"

	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"

	"github.com/gin-gonic/gin"
)

func main() {

	logger := logrus.New()
	library := library_pkg.NewLibrary(library_pkg.NewMysqlRepository())
	handlers := api.NewAPI(logger)

	router := gin.New()
	router.Use(ginlogrus.Logger(logger.WithField("service", "gin")), gin.Recovery())
	api := router.Group("/api")

	logger.Info("Welcome to the GeeksHubs Library. API is ready.")

	api.GET("/", handlers.Hello)
	api.POST("/books", handlers.AddBook(library))
	api.GET("/books", handlers.GetAllBooks(library))
	api.GET("/books/:id", handlers.GetBook(library))
	api.DELETE("/books/:id", handlers.DeleteBook(library))
	api.PUT("/books/:id", handlers.UpdateBook(library))

	router.Run(":8080")
}
