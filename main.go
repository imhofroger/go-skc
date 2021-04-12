package main

import (
	"github.com/imhofroger/go-skc/controllers"
	"github.com/imhofroger/go-skc/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/entries", controllers.FindBooks)
	r.GET("/entries/:id", controllers.FindBook)
	r.POST("/entries", controllers.CreateBook)
	r.PATCH("/entries/:id", controllers.UpdateBook)
	r.DELETE("/entries/:id", controllers.DeleteBook)

	// Run the server
	r.Run()
}
