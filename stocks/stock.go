package stocks

// Stock represents a company's stock in the database
type Stock struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `json:"name" gorm:"unique;not null"` // Stock name (e.g., AAPL, MSFT)
	Price float64 `json:"price" gorm:"not null"`       // Current price of the stock
}
