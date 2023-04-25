package model

// Account defines struct of account data.
type Account struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Name        string     `json:"name"`
	Password    string     `json:"-"`
	AuthorityID uint       `json:"authority_id"`
	Authority   *Authority `json:"authority"`
}
