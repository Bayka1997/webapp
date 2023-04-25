package model

// Category defines struct of category data.
type Category struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `validate:"required" json:"name"`
}
