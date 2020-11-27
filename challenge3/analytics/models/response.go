package models

import "time"

// Response ...
type Response struct {
	From time.Time     `json:"from"`
	To   time.Time     `json:"to"`
	Data []*GrowthData `json:"data"`
}

// GrowthData ...
type GrowthData struct {
	MarketPair string `json:"market_pair"`
	Data       struct {
		VolumeGrowth float64 `json:"volume_growth"`
		HighGrowth   float64 `json:"high_growth"`
		LowGrowth    float64 `json:"low_growth"`
	} `json:"data"`
}
