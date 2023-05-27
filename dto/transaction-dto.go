package dto

type TransactionCreateDTO struct {
	UserID           uint64 `json:"userid" form:"userid" binding:"required"`
	TransactionType  string `json:"trxtype" form:"trxtype" binding:"required"`
	Date             string `json:"date" form:"date" binding:"required"`
	TransactionValue int    `json:"trxvalue" form:"trxvalue" binding:"required"`
}

type TransactionUpdateDTO struct {
	ID               uint64 `json:"id" form:"id" binding:"required"`
	UserID           uint64 `json:"userid" form:"userid" binding:"required"`
	TransactionType  string `json:"trxtype" form:"trxtype" binding:"required"`
	Date             string `json:"date" form:"date" binding:"required"`
	TransactionValue int    `json:"trxvalue" form:"trxvalue" binding:"required"`
}
