package model

import "gorm.io/gorm"

// Category struct
type Category struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	ParentId    int    `gorm:"not null" json:"parent_id"`

	Product []*Product `json:"product,omitempty" gorm:"foreignKey:CategoryId;references:ID"`
}
