package routes

import (
	"example/bookstore/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.Engine) {
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.GetBookByID)
	router.POST("/books", handlers.PostBooks)
	router.GET("/", handlers.Homepage)
}
