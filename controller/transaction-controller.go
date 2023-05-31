package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IrvanWijayaSardam/CashFlow/dto"
	"github.com/IrvanWijayaSardam/CashFlow/entity"
	"github.com/IrvanWijayaSardam/CashFlow/helper"
	"github.com/IrvanWijayaSardam/CashFlow/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TransactionContoller interface {
	All(context *gin.Context)
	SumGroupId(context *gin.Context)
	TransactionReport(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type transactionController struct {
	transactionService service.TransactionService
	jwtService         service.JWTService
}

func NewTransactionController(trxServ service.TransactionService, jwtServ service.JWTService) TransactionContoller {
	return &transactionController{
		transactionService: trxServ,
		jwtService:         jwtServ,
	}
}

func (c *transactionController) All(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	trx := c.transactionService.All(userID)
	res := helper.BuildResponse(true, "OK!", trx)
	context.JSON(http.StatusOK, res)
}

func (c *transactionController) SumGroupId(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	trx := c.transactionService.SumGroupId(userID)
	res := helper.BuildResponse(true, "OK!", trx)
	context.JSON(http.StatusOK, res)
}

func (c *transactionController) TransactionReport(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	trx := c.transactionService.TransactionReport(userID)
	res := helper.BuildResponse(true, "OK!", trx)
	context.JSON(http.StatusOK, res)
}

func (c *transactionController) Insert(context *gin.Context) {
	var transactionCreateDTO dto.TransactionCreateDTO
	errDTO := context.ShouldBind(&transactionCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			transactionCreateDTO.UserID = convertedUserID
		}
		result := c.transactionService.InsertTransaction(transactionCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *transactionController) Update(context *gin.Context) {
	var transactionUpdateDTO dto.TransactionUpdateDTO
	errDTO := context.ShouldBind(&transactionUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["userid"])
	if c.transactionService.IsAllowedToEdit(userID, transactionUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			transactionUpdateDTO.UserID = id
		}
		result := c.transactionService.UpdateTransaction(transactionUpdateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *transactionController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["userid"])
	return id
}

func (c *transactionController) Delete(context *gin.Context) {
	var transaction entity.Transaction
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	transaction.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["userid"])
	if c.transactionService.IsAllowedToEdit(userID, transaction.ID) {
		c.transactionService.Delete(transaction)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}
