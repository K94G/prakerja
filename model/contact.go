package model

type Contact struct {
	Id            uint     `json:"id" gorm:"primaryKey;autoIncrement;"`
	Name          string   `json:"name"`
	Phone         string   `json:"phone"`
	Email         string   `json:"email"`
	Address       string   `json:"address"`
	City          string   `json:"city"`
	Province      string   `json:"province"`
	PostalCode    int      `json:"postal_code"`
	GoogleMapsUrl string   `json:"google_maps_url"`
	CategoryId    uint     `json:"category_id"`
	Category      Category `gorm:"foreignKey:CategoryId"`
	Active        bool     `json:"active"`
}
