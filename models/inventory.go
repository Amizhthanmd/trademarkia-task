package models

import "time"

type Inventory struct {
	ID        string    `json:"id,omitempty" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
