package dto

type RegisterDTO struct {
	ID        uint64 `json:"id" form:"id"`
	Name      string `json:"name" form:"name" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password" form:"password" binding:"required"`
	Profile   string `json:"profile" form:"profile"` // Changed type to string
	Jk        string `json:"jk" form:"jk" binding:"required"`
	Pin       string `json:"pin" form:"pin" binding:"required"`
	Telephone string `json:"telp" form:"telp" binding:"required"`
}
