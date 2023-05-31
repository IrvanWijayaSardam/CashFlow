package repository

import (
	"github.com/IrvanWijayaSardam/CashFlow/entity"
	"github.com/IrvanWijayaSardam/CashFlow/helper"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	InsertTransaction(b *entity.Transaction) entity.Transaction
	UpdateTransaction(b *entity.Transaction) entity.Transaction
	All(idUser string) []entity.Transaction
	SumGroupId(idUser string) []helper.TransactionGroupSum
	TransactionReport(idUser string) helper.TransactionReport
	DeleteTransaction(b *entity.Transaction)
	FindTransactionById(UserID uint64) entity.Transaction
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

func (db *transactionConnection) All(idUser string) []entity.Transaction {
	var transactions []entity.Transaction
	db.connection.Preload("User").Where("user_id = ?", idUser).Find(&transactions)
	return transactions
}

func (db *transactionConnection) SumGroupId(idUser string) []helper.TransactionGroupSum {
	var result []helper.TransactionGroupSum
	db.connection.Model(&entity.Transaction{}).
		Select("transaction_group, SUM(transaction_value) AS total_transaction").
		Where("user_id = ?", idUser).
		Group("transaction_group").
		Scan(&result)
	return result
}

func (db *transactionConnection) TransactionReport(idUser string) helper.TransactionReport {
	var result helper.TransactionReport

	db.connection.Model(&entity.Transaction{}).
		Select("SUM(CASE WHEN transaction_type = '1' THEN transaction_value ELSE 0 END) AS transaction_in, "+
			"SUM(CASE WHEN transaction_type = '2' THEN transaction_value ELSE 0 END) AS transaction_out").
		Where("user_id = ?", idUser).
		Scan(&result)

	return result
}

func (db *transactionConnection) UpdateTransaction(b *entity.Transaction) entity.Transaction {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return *b
}

func (db *transactionConnection) DeleteTransaction(b *entity.Transaction) {
	db.connection.Delete(&b)
}

func (db *transactionConnection) FindTransactionById(UserID uint64) entity.Transaction {
	var transaction entity.Transaction
	db.connection.Preload("User").Find(&transaction, UserID)
	return transaction
}
