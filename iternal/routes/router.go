package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"service/iternal/handlers"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.Static("/img", "./img")
		api.GET("/getProducts", handlers.GetProducts(db))
		api.POST("/createProducts", handlers.CreateProducts(db))
		api.PUT("/updateProducts", handlers.UpdateProducts(db))
		api.DELETE("/deleteProducts/:id", handlers.DeleteProducts(db))
		api.GET("/getStocks", handlers.GetStocks(db))
		api.POST("/createStocks", handlers.CreateStocks(db))
		api.PUT("/updateStocks", handlers.UpdateStocks(db))
		api.DELETE("/deleteStocks/:id", handlers.DeleteStocks(db))
		api.GET("/getStores", handlers.GetStores(db))
		api.POST("/createStores", handlers.CreateStores(db))
		api.PUT("/updateStores", handlers.UpdateStores(db))
		api.DELETE("/deleteStores/:id", handlers.DeleteStores(db))
	}

	return router
}
