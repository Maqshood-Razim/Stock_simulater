package main

import (
	"log"
	"stock-market/stocks"

	"github.com/gin-gonic/gin"
)

func main() {

	stocks.InitDB()

	router := gin.Default()

	router.GET("/stocks", stocks.GetStocks)
	router.POST("/orders", stocks.PlaceOrder)
	router.GET("/orders", stocks.GetAllOrders)
	router.GET("/simulate", stocks.StartSimulation)

	log.Println("Server started on :8080")
	router.Run(":8080")
}
