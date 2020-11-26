package storage

import (
	"time"

	"github.com/ebikode/peaq-challenge/challenge3/exchange/models"
)

// GrowthRecordStorage ...
type GrowthRecordStorage struct {
	*MDatabase
}

// NewGrowthRecordStorage Initialize GrowthRecord Storage
func NewGrowthRecordStorage(db *MDatabase) *GrowthRecordStorage {
	return &GrowthRecordStorage{db}
}

// Fetch rate using user_id and rate_id
func (r *GrowthRecordStorage) Fetch(rateID uint) *models.GrowthRecord {
	rate := models.GrowthRecord{}
	// Select resource from database
	err := r.db.
		Where("id=?", rateID).
		First(&rate).Error

	if rate.ID < 1 || err != nil {
		return nil
	}

	return &rate
}

// FetchByTimestamp Get rates from - to timestamp
func (r *GrowthRecordStorage) FetchByTimestamp(from, to int64) []*models.GrowthRecord {
	var rates []*models.GrowthRecord
	fromTime := time.Unix(from, 0).UTC()
	toTime := time.Unix(to, 0).UTC()

	// Select resource from database
	r.db.
		Where("from >= ?", fromTime).
		Where("to >= ?", toTime).
		Order("created_at desc").
		Find(&rates)

	return rates
}

// Store Add a new rate
func (r *GrowthRecordStorage) Store(p models.GrowthRecord) (*models.GrowthRecord, error) {

	rate := p

	err := r.db.Create(&rate).Error

	if err != nil {
		return nil, err
	}
	return &rate, nil
}

// Update a rate
func (r *GrowthRecordStorage) Update(rate *models.GrowthRecord) (*models.GrowthRecord, error) {

	err := r.db.Save(&rate).Error

	if err != nil {
		return nil, err
	}

	return rate, nil
}
