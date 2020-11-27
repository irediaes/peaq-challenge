package models

// Rate ...
type Rate struct {
	MarketName string  `json:"market_name"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Volume     float64 `json:"volume"`
}
