package service

import (
	"log"

	"github.com/IrvanWijayaSardam/CashFlow/dto"
	"github.com/IrvanWijayaSardam/CashFlow/entity"
	"github.com/IrvanWijayaSardam/CashFlow/repository"
	"github.com/mashingan/smapping"
)

type TransactionService interface {
	InsertTransaction(b dto.TransactionCreateDTO) entity.Transaction
	All(idUser string) []entity.Transaction
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
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
