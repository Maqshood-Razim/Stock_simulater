package main

import (
	"log"
	"stock-market/stocks"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database with GORM
	stocks.InitDB()

	// Setup Router
	router := gin.Default()

	// Routes
	router.GET("/stocks", stocks.GetStocks)         // Get current stock prices
	router.POST("/orders", stocks.PlaceOrder)       // Place a Buy/Sell order
	router.GET("/orders", stocks.GetAllOrders)      // Get all orders
	router.GET("/simulate", stocks.StartSimulation) // Start stock price simulation

	// Start Server
	log.Println("Server started on :8080")
	router.Run(":8080")
}
