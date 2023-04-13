package model

type Category struct {
	Id   uint   `json:"-" gorm:"primaryKey;autoIncrement;"`
	Name string `json:"name"`
}
