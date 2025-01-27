package stocks

// Order represents a Buy/Sell order for a stock
type Order struct {
	ID        uint    `gorm:"primaryKey"`
	StockName string  `json:"stock_name" gorm:"not null"`                 // The stock name for the order
	OrderType string  `json:"order_type" gorm:"type:enum('BUY', 'SELL')"` // Type of the order: BUY or SELL
	Price     float64 `json:"price" gorm:"not null"`                      // Order price
	Quantity  int     `json:"quantity" gorm:"not null"`                   // Quantity of stocks in the order
}
