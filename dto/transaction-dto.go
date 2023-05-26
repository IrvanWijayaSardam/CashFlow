package dto

type TransactionCreateDTO struct {
	UserID            uint64 `json:"userid" form:"userid" binding:"required"`
	TransactionTypeID int    `json:"trxtid" form:"trxtid" binding:"required"`
	Date              string `json:"date" form:"date" binding:"required"`
	TransactionValue  int    `json:"trxvalue" form:"trxvalue" binding:"required"`
}
