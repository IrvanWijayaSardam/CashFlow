package service

import (
	"fmt"
	"log"

	"github.com/IrvanWijayaSardam/CashFlow/dto"
	"github.com/IrvanWijayaSardam/CashFlow/entity"
	"github.com/IrvanWijayaSardam/CashFlow/helper"
	"github.com/IrvanWijayaSardam/CashFlow/repository"
	"github.com/mashingan/smapping"
)

type TransactionService interface {
	InsertTransaction(b dto.TransactionCreateDTO) entity.Transaction
	UpdateTransaction(b dto.TransactionUpdateDTO) entity.Transaction
	Delete(b entity.Transaction)
	IsAllowedToEdit(userID string, transactionID uint64) bool
	All(idUser string) []entity.Transaction
	SumGroupId(idUser string) []helper.TransactionGroupSum
	TransactionReport(idUser string) helper.TransactionReport
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

// TransactionReport implements TransactionService
func (service *transactionService) TransactionReport(idUser string) helper.TransactionReport {
	return service.transactionRepository.TransactionReport(idUser)
}

// SumGroupId implements TransactionService
func (service *transactionService) SumGroupId(idUser string) []helper.TransactionGroupSum {
	return service.transactionRepository.SumGroupId(idUser)
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepo,
	}
}

func (service *transactionService) InsertTransaction(b dto.TransactionCreateDTO) entity.Transaction {
	trx := entity.Transaction{}
	err := smapping.FillStruct(&trx, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.transactionRepository.InsertTransaction(&trx)
	return res
}

func (service *transactionService) All(idUser string) []entity.Transaction {
	return service.transactionRepository.All(idUser)
}

func (service *transactionService) UpdateTransaction(b dto.TransactionUpdateDTO) entity.Transaction {
	transaction := entity.Transaction{}
	err := smapping.FillStruct(&transaction, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v : ", err)
	}
	res := service.transactionRepository.UpdateTransaction(&transaction)
	return res
}

func (service *transactionService) Delete(b entity.Transaction) {
	service.transactionRepository.DeleteTransaction(&b)
}

func (service *transactionService) IsAllowedToEdit(userID string, transactionID uint64) bool {
	b := service.transactionRepository.FindTransactionById(transactionID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
