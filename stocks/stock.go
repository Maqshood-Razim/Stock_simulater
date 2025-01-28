package stocks


type Stock struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `json:"name" gorm:"unique;not null"` 
	Price float64 `json:"price" gorm:"not null"`      
}
