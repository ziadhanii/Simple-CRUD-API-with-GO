package routes

import (
	"github.com/gin-gonic/gin"
	"go-crud-api/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	hotelRoutes := router.Group("/hotels")
	{
		hotelRoutes.GET("/", handlers.GetHotels)
		hotelRoutes.GET("/:id", handlers.GetHotelByID)
		hotelRoutes.POST("/", handlers.CreateHotel)
		hotelRoutes.PUT("/:id", handlers.UpdateHotel)
		hotelRoutes.DELETE("/:id", handlers.DeleteHotel)

	}
}
