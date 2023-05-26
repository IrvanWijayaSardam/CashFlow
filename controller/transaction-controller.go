package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IrvanWijayaSardam/CashFlow/dto"
	"github.com/IrvanWijayaSardam/CashFlow/helper"
	"github.com/IrvanWijayaSardam/CashFlow/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TransactionContoller interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
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
	res := helper.BuildResponse(true, "OK!", []interface{}{trx})
	context.JSON(http.StatusOK, res)
}

func (c *transactionController) Insert(context *gin.Context) {
	var transactionCreateDTO dto.TransactionCreateDTO
	errDTO := context.ShouldBind(&transactionCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), []interface{}{helper.EmptyObj{}})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			transactionCreateDTO.UserID = convertedUserID
		}
		result := c.transactionService.InsertTransaction(transactionCreateDTO)
		response := helper.BuildResponse(true, "OK!", []interface{}{result})
		context.JSON(http.StatusCreated, response)
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
