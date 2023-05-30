package entity

type Transaction struct {
	ID               uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID           uint64 `gorm:"index" json:"user_id"`
	TransactionType  string `json:"transaction_type"`
	Description      string `json:"description"`
	Date             string `json:"date"`
	TransactionValue int    `json:"transaction_value"`
	TransactionGroup string `json:"transaction_group"`
}
