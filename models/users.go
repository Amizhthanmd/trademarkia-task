package models

import "time"

type User struct {
	ID        string `json:"id,omitempty" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	Role      string `json:"role" gorm:"default:user"`

	// One-to-Many (orders)
	Orders    []Order   `json:"orders,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Order struct {
	ID          string  `json:"id,omitempty" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	OrderNumber string  `json:"order_number"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`

	UserID string `json:"user_id,omitempty"`
	User   User   `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// Many-to-Many (Products)
	Products  []Product `gorm:"many2many:order_products;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
