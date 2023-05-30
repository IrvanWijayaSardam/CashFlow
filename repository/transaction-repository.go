package repository

import (
	"github.com/IrvanWijayaSardam/CashFlow/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	InsertTransaction(b *entity.Transaction) entity.Transaction
	UpdateTransaction(b *entity.Transaction) entity.Transaction
	All(idUser string) []entity.Transaction
	SumGroupId(idUser string, idGroup string) int
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

func (db *transactionConnection) SumGroupId(idUser string, groupId string) int {
	var result int
	db.connection.Model(&entity.Transaction{}).
		Select("SUM(transaction_value)").
		Where("user_id = ? AND transaction_group = ?", idUser, groupId).
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
