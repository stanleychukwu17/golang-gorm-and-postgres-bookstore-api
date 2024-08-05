package schema

import (
	"time"
)

// --START-- sample objects
// User model
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// One-to-Many relationship with Orders
	Orders []Order
}

// Order model
type Order struct {
	ID          uint    `gorm:"primaryKey"`
	OrderNumber string  `gorm:"uniqueIndex;not null"`
	TotalAmount float64 `gorm:"not null"`
	UserID      uint    // Foreign key to User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Product model
type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"uniqueIndex;not null"`
	Description string
	Price       float64 `gorm:"not null"`
	Orders      []Order `gorm:"many2many:order_products;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// OrderProduct model (Join table for many-to-many relationship)
type OrderProduct struct {
	OrderID   uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Quantity  int  `gorm:"not null"`
}

//--EMD--

type Book struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Author    string `gorm:"type:varchar(100);not null" json:"author"`
	Publisher string `gorm:"type:varchar(100);not null" json:"publisher"`
	Title     string `gorm:"type:varchar(100);not null" json:"title"`
}
