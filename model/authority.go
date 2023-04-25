package model

// Authority defines struct of authority data.
type Authority struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
