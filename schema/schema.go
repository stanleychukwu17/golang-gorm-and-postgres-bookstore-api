package schema

import (
	"time"
)

// --START-- sample objects

// User model
type User struct {
	ID        uint16    `gorm:"primaryKey;type:smallint" json:"id,omitempty"`
	Username  string    `gorm:"uniqueIndex;type:varchar(100);not null" json:"username,omitempty"`
	Email     string    `gorm:"uniqueIndex;type:varchar(100);not null" json:"email,omitempty"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// One-to-Many relationship with Orders
	Orders []Order
}

// Order model
type Order struct {
	ID          uint    `gorm:"primaryKey"`
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

//--EMD--

// Book model
type Book struct {
	ID        uint16    `gorm:"primaryKey;type:smallint" json:"id,omitempty"`
	Author    string    `gorm:"type:varchar(100);not null" json:"author,omitempty"`
	Publisher string    `gorm:"type:varchar(100);not null" json:"publisher,omitempty"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
