package handlers

import (
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHotels godoc
// @Summary Get all hotels
// @Description Retrieve all hotels from the database
// @Tags Hotels
// @Accept json
// @Produce json
// @Success 200 {array} models.Hotel
// @Router /hotels [get]
func GetHotels(c *gin.Context) {
	var hotels []models.Hotel
	config.DB.Find(&hotels)
	c.JSON(http.StatusOK, hotels)
}

// GetHotelByID godoc
// @Summary Get a hotel by ID
// @Description Retrieve a specific hotel by its ID
// @Tags Hotels
// @Accept json
// @Produce json
// @Param id path int true "Hotel ID"
// @Success 200 {object} models.Hotel
// @Failure 404 {object} map[string]string
// @Router /hotels/{id} [get]
func GetHotelByID(c *gin.Context) {
	id := c.Param("id")
	var hotel models.Hotel
	if err := config.DB.First(&hotel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}
	c.JSON(http.StatusOK, hotel)
}

// CreateHotel godoc
// @Summary Create a new hotel
// @Description Add a new hotel to the database
// @Tags Hotels
// @Accept json
// @Produce json
// @Param hotel body models.Hotel true "Hotel Data"
// @Success 201 {object} models.Hotel
// @Failure 400 {object} map[string]string
// @Router /hotels [post]
func CreateHotel(c *gin.Context) {
	var hotel models.Hotel
	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&hotel)
	c.JSON(http.StatusCreated, hotel)
}

// UpdateHotel godoc
// @Summary Update an existing hotel
// @Description Update the details of a specific hotel
// @Tags Hotels
// @Accept json
// @Produce json
// @Param id path int true "Hotel ID"
// @Param hotel body models.Hotel true "Updated Hotel Data"
// @Success 200 {object} models.Hotel
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /hotels/{id} [put]
func UpdateHotel(c *gin.Context) {
	id := c.Param("id")
	var hotel models.Hotel
	if err := config.DB.First(&hotel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}
	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&hotel)
	c.JSON(http.StatusOK, hotel)
}

// DeleteHotel godoc
// @Summary Delete a hotel
// @Description Remove a hotel from the database
// @Tags Hotels
// @Accept json
// @Produce json
// @Param id path int true "Hotel ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /hotels/{id} [delete]
func DeleteHotel(c *gin.Context) {
	id := c.Param("id")
	var hotel models.Hotel

	if err := config.DB.First(&hotel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	if err := config.DB.Delete(&hotel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete hotel"})
		return
	}

	c.Status(http.StatusNoContent)
}
