package models

import (
	"time"

	"gorm.io/gorm"

	"courier/src/courierpb"
)

type Parcel struct {
	gorm.Model
	ParcelID int64
	Email    string
	Phone    string
	Weight   float32
	Date     *time.Time
}

func (Parcel) TableName() string {
	return "parcels"
}

func FromPb(p *courierpb.Parcel) *Parcel {
	return &Parcel{
		ParcelID: p.GetId(),
		Email:    p.GetEmail(),
		Phone:    p.GetPhone(),
		Weight:   p.GetWeight(),
	}
}
