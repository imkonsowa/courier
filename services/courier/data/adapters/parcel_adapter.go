package adapters

import (
	"fmt"
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

func (p *ParcelAdapter) Count(day *time.Time) int64 {
	var count int64

	p.DB.
		Where(&models.Parcel{Date: datatypes.Date(*day)}).
		Model(&models.Parcel{}).
		Count(&count)

	return count
}

func (p *ParcelAdapter) PaginatedParcels(day *time.Time, country string, offset, limit int) ([]*models.Parcel, error) {
	if day == nil || day.IsZero() {
		return nil, fmt.Errorf("valid date value is required")
	}

	var parcels []*models.Parcel

	cond := &models.Parcel{
		Date: datatypes.Date(*day),
	}

	if len(country) != 0 {
		cond.Country = country
	}

	p.DB.
		Where(cond).
		Select("id, parcel_id, weight").
		Offset(offset).
		Limit(limit).
		Order("parcel_id asc").
		Find(&parcels)

	return parcels, nil
}
