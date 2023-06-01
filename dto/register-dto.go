package dto

import "mime/multipart"

type RegisterDTO struct {
	ID        uint64                `json:"id" form:"id"`
	Name      string                `json:"name" form:"name" binding:"required"`
	Email     string                `json:"email" form:"email" binding:"required,email"`
	Profile   *multipart.FileHeader `form:"profile"`
	Jk        string                `json:"jk" form:"jk" binding:"required"`
	Pin       string                `json:"pin" form:"pin" binding:"required"`
	Telephone string                `json:"telp" form:"telp" binding:"required"`
}
