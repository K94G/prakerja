package model

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement;"`
	Name     string `json:"name"`
	Username string `json:"username" gorm:"index:idx_username,unique"`
	Password string `json:"password"`
}
