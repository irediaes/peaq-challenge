package models

import (
	"time"
)

// GrowthRecord ...
type GrowthRecord struct {
	FromRateID   uint      `json:"from_rate_id"`
	ToRateID     uint      `json:"to_rate_id"`
	VolumeGrowth float64   `json:"volume_growth"`
	HighGrowth   float64   `json:"high_growth"`
	LowGrowth    float64   `json:"low_growth"`
	FromDate     time.Time `json:"from"`
	ToDate       time.Time `json:"to"`
	FromRate     Rate      `json:"from_rate"`
	ToRate       Rate      `json:"to_rate"`
}
