package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/IrvanWijayaSardam/CashFlow/dto"
	"github.com/IrvanWijayaSardam/CashFlow/helper"
	"github.com/IrvanWijayaSardam/CashFlow/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
	SaveFile(context *gin.Context)
	GetFile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["userid"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["userid"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}

func (c *userController) SaveFile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["userid"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	// File Upload
	file, err := context.FormFile("file")
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	filePath, err := c.userService.SaveFile(file)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	err = c.userService.UpdateUserProfile(id, filePath)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", filePath)
	context.JSON(http.StatusOK, res)
}

func (c *userController) GetFile(context *gin.Context) {
	fileName := context.Param("file_name") // Assuming the filename is passed as a route parameter

	filePath := "./cdn/" + fileName // Construct the file path based on the filename and the directory

	file, err := os.Open(filePath)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to fetch file", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		res := helper.BuildErrorResponse("Failed to fetch file", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	// Get the file's content type
	contentType := http.DetectContentType(nil)

	// Set the appropriate headers for the response
	context.Header("Content-Type", contentType)
	context.Header("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	// Copy the file data to the response body
	_, err = io.Copy(context.Writer, file)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to fetch file", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
}
