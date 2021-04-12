package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imhofroger/go-skc/models"
)

type CreateEntryInput struct {
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
}

type UpdateEntryInput struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// GET /entries
// Find all entries
func FindEntries(c *gin.Context) {
	var entry []models.entry
	models.DB.Find(&entries)

	c.JSON(http.StatusOK, gin.H{"data": entries})
}

// GET /entries/:id
// Find a entry
func FindEntry(c *gin.Context) {
	// Get model if exist
	var entry models.entry
	if err := models.DB.Where("id = ?", c.Param("id")).First(&entry).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": entry})
}

// POST /entries
// Create new entry
func CreateEntry(c *gin.Context) {
	// Validate input
	var input CreateentryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create entry
	entry := models.entry{Title: input.Title, Artist: input.Artist}
	models.DB.Create(&entry)

	c.JSON(http.StatusOK, gin.H{"data": entry})
}

// PATCH /entries/:id
// Update a entry
func UpdateEntry(c *gin.Context) {
	// Get model if exist
	var entry models.entry
	if err := models.DB.Where("id = ?", c.Param("id")).First(&entry).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateEntryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&entry).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": entry})
}

// DELETE /entries/:id
// Delete a entry
func Deleteentry(c *gin.Context) {
	// Get model if exist
	var entry models.entry
	if err := models.DB.Where("id = ?", c.Param("id")).First(&entry).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&entry)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
