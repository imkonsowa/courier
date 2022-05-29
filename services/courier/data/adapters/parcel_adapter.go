package adapters

import (
	"time"

	"gorm.io/datatypes"
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

func (p *ParcelAdapter) PaginatedParcels(day *time.Time, offset, limit int) ([]*models.Parcel, error) {
	var parcels []*models.Parcel

	p.DB.
		Where(&models.Parcel{Date: datatypes.Date(*day)}).
		Select("id, parcel_id, weight").
		Offset(offset).
		Limit(limit).
		Order("parcel_id asc").
		Find(&parcels)

	return parcels, nil
}
