package entity

type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	Email     string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password  string `gorm:"->;<-;not null" json:"-"`
	Profile   string `gorm:"type:varchar(255)" json:"profile"`
	Telephone string `gorm:"type:varchar(255)" json:"telp"`
	Pin       string `gorm:"type:varchar(255)" json:"pin"`
	Jk        string `gorm:"type:varchar(255)" json:"jk"`
	Token     string `gorm:"-" json:"token,omitempty"`
}
