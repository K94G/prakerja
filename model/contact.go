package model

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	Id            uint      `json:"id" gorm:"primaryKey;autoIncrement;"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Address       string    `json:"address"`
	City          string    `json:"city"`
	Province      string    `json:"province"`
	PostalCode    int       `json:"postal_code"`
	GoogleMapsUrl string    `json:"google_maps_url"`
	CategoryId    uint      `json:"category_id"`
	Category      Category  `json:"category"`
	Active        bool      `json:"active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"-"`
}

func TimeLocalJakarta() time.Time {
	now := time.Now().UTC()
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		now = time.Now()
	} else {
		now = now.In(loc)
	}

	return now
}

func (m *Contact) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = TimeLocalJakarta()
	m.UpdatedAt = TimeLocalJakarta()
	return nil
}
