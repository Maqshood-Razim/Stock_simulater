package stocks


type Order struct {
	ID        uint    `gorm:"primaryKey"`
	StockName string  `json:"stock_name" gorm:"not null"`                 
	OrderType string  `json:"order_type" gorm:"type:enum('BUY', 'SELL')"` 
	Price     float64 `json:"price" gorm:"not null"`                      
	Quantity  int     `json:"quantity" gorm:"not null"`                   
}
