package stocks

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// StartSimulation simulates stock price changes over time
func StartSimulation(c *gin.Context) {
	go func() {
		stocks := []string{"Apple", "Google", "Amazon", "Microsoft"}

		// Create a new random generator with a custom seed
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))

		for {
			for _, stock := range stocks {
				price := 100 + rng.Float64()*50 // Use the local random generator
				var stockRecord Stock
				if err := DB.FirstOrCreate(&stockRecord, Stock{Name: stock}).Error; err == nil {
					stockRecord.Price = price
					DB.Save(&stockRecord)
				}
			}
			time.Sleep(2 * time.Second)
		}
	}()
	c.JSON(http.StatusOK, gin.H{"message": "Stock price simulation started"})
}

// GetStocks retrieves all current stock prices
func GetStocks(c *gin.Context) {
	var stocks []Stock
	if err := DB.Find(&stocks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stocks)
}

// PlaceOrder places a new buy/sell order
func PlaceOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order": order})
}

// GetAllOrders retrieves all orders from the database
func GetAllOrders(c *gin.Context) {
	var orders []Order
	if err := DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
