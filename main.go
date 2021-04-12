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
	r.GET("/entries", controllers.FindEntries)
	r.GET("/entries/:id", controllers.FindEntries)
	r.POST("/entries", controllers.CreateEntry)
	r.PATCH("/entries/:id", controllers.UpdateEntry)
	r.DELETE("/entries/:id", controllers.DeleteEntry)

	// Run the server
	r.Run()
}
