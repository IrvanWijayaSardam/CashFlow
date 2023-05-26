package entity

type Transaction struct {
	ID               int    `gorm:"primary_key:auto_increment" json:"id"`
	UserID           uint64 `gorm:"index" json:"user_id"`
	TransactionType  string `json:"transaction_type"`
	Date             string `json:"date"`
	TransactionValue int    `json:"transaction_value"`
}
