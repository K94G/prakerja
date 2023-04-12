package model

type Category struct {
	Id   uint   `json:"id" gorm:"primaryKey;autoIncrement;"`
	Name string `json:"name"`
}
