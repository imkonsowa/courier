package adapters

import (
	"gorm.io/gorm"

	"courier/services/courier/data/models"
)

type ParcelAdapter struct {
	DB *gorm.DB
}

func NewParcelAdapter(db *gorm.DB) *ParcelAdapter {
	return &ParcelAdapter{
		DB: db,
	}
}

func (p *ParcelAdapter) PatchInsert(parcels []*models.Parcel) ([]*models.Parcel, error) {
	if len(parcels) == 0 {
		return parcels, nil
	}

	err := p.DB.Save(&parcels).Error
	if err != nil {
		return nil, err
	}

	return parcels, nil
}
