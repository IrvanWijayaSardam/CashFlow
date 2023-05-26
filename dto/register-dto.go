package dto

type RegisterDTO struct {
	Name      string `json:"name" form:"name" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password" form:"password" binding:"required"`
	Profile   string `json:"profile" form:"profile" binding:"required"`
	Telephone string `json:"telp" form:"telp" binding:"required"`
	Pin       string `json:"pin" form:"pin" binding:"required"`
	Jk        string `json:"jk" form:"jk" binding:"required"`
}
