package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"courier/courierpb"
)

type Parcel struct {
	gorm.Model
	ParcelID int64 `gorm:"index:idx_parcel_id"`
	Email    string
	Phone    string
	Weight   float64
	Country  string         `gorm:"index:idx_country"`
	Date     datatypes.Date `gorm:"index:idx_date"`
}

func (Parcel) TableName() string {
	return "parcels"
}

func FromPb(p *courierpb.Parcel, d *time.Time) *Parcel {
	return &Parcel{
		ParcelID: p.GetId(),
		Email:    p.GetEmail(),
		Phone:    p.GetPhone(),
		Weight:   float64(p.GetWeight()),
		Country:  p.Country,
		Date:     datatypes.Date(*d),
	}
}
