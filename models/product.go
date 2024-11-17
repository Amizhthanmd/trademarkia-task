package models

import "time"

type Product struct {
	ID          string  `json:"id,omitempty" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`

	Inventory   Inventory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	InventoryID string    `json:"inventory_id,omitempty"`

	Orders    []Order   `gorm:"many2many:order_products;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
