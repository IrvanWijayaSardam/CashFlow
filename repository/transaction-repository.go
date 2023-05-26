package repository

import (
	"github.com/IrvanWijayaSardam/CashFlow/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	InsertTransaction(b *entity.Transaction) entity.Transaction
	All() []entity.Transaction
}

type transactionConnection struct {
	connection *gorm.DB
}

func NewTransactionRepository(dbConn *gorm.DB) TransactionRepository {
	return &transactionConnection{
		connection: dbConn,
	}
}

func (db *transactionConnection) InsertTransaction(b *entity.Transaction) entity.Transaction {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return *b
}

func (db *transactionConnection) All() []entity.Transaction {
	var transactions []entity.Transaction
	db.connection.Preload("UserID").Find(&transactions)
	return transactions
}
