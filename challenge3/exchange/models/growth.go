package models

import (
	"time"
)

// GrowthRecord ...
type GrowthRecord struct {
	BaseModel
	FromRateID   uint      `json:"from_rate_id" gorm:"not null;type:int(11)"`
	ToRateID     uint      `json:"to_rate_id" gorm:"not null;type:int(11)"`
	VolumeGrowth float64   `json:"volume_growth" gorm:"type:float(16,8);default:0.0"`
	HighGrowth   float64   `json:"high_growth" gorm:"type:float(16,8);default:0.0"`
	LowGrowth    float64   `json:"low_growth" gorm:"type:float(16,8);default:0.0"`
	From         time.Time `json:"from"`
	To           time.Time `json:"to"`
	FromRate     Rate      `json:"from_rate"`
	ToRate       Rate      `json:"to_rate"`
}
