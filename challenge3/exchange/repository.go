package main

import (
	"github.com/ebikode/peaq-challenge/challenge3/exchange/models"
	pb "github.com/ebikode/peaq-challenge/challenge3/exchange/proto/rate"
)

type repository interface {
	GetByTimestamp(int64, int64) ([]*models.GrowthRecord, error)
}

// UnmarshalGrowthRecordList ...
func UnmarshalGrowthRecordList(records []*models.GrowthRecord) []*pb.GrowthRecord {
	recordList := make([]*pb.GrowthRecord, 0)
	for _, record := range records {
		recordList = append(recordList, UnmarshalGrowthRecord(record))
	}
	return recordList
}

// UnmarshalGrowthRecord ...
func UnmarshalGrowthRecord(rate *models.GrowthRecord) *pb.GrowthRecord {

	return &pb.GrowthRecord{
		Id:           int32(rate.ID),
		FromRateId:   int32(rate.FromRateID),
		ToRateId:     int32(rate.ToRateID),
		VolumeGrowth: float32(rate.VolumeGrowth),
		HighGrowth:   float32(rate.HighGrowth),
		LowGrowth:    float32(rate.LowGrowth),
		From:         rate.From.Format("2006-01-02 15:04:05"),
		To:           rate.To.Format("2006-01-02 15:04:05"),
	}
}

// formatTime() {
// 	t.Format("2006-01-02 15:04:05")
// }
