package model

// Book defines struct of book data.
type Book struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	Title      string    `json:"title"`
	Isbn       string    `json:"isbn"`
	CategoryID uint      `json:"categoryId"`
	Category   *Category `json:"category"`
	FormatID   uint      `json:"formatId"`
	Format     *Format   `json:"format"`
}
