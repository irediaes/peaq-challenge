package main

import (
	"time"

	"github.com/ebikode/peaq-challenge/challenge3/exchange/models"
	pb "github.com/ebikode/peaq-challenge/challenge3/exchange/proto/rate"
	"github.com/ebikode/peaq-challenge/challenge3/exchange/utils"
)

const (
	timeFormat = "2006-01-02 15:04:05"
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
		From:         rate.FromDate,
		To:           rate.ToDate,
		FromRate:     unmarshalRate(rate.FromRate),
		ToRate:       unmarshalRate(rate.ToRate),
	}
}

// unmarshalRate ...
func unmarshalRate(rate models.Rate) *pb.Rate {

	return &pb.Rate{
		MarketName: rate.MarketName,
		Volume:     float32(rate.Volume),
		High:       float32(rate.High),
		Low:        float32(rate.Low),
	}
}

// UnmarshalProtoResponseData ...
func UnmarshalProtoResponseData(records []*models.GrowthRecord, numberOfCrytoRate int) []*pb.ResponseData {

	responseDataList := []*pb.ResponseData{}

	for i := 0; i < len(records); i++ {

		record := records[i]
		responseMarketData := []*pb.ResponseMarketData{}

		marketData := &pb.ResponseMarketData{
			MarketPair: record.FromRate.MarketName,
		}
		growthData := &pb.GrowthData{}

		growthData.VolumeGrowth = float32(utils.Round4Decimal(record.VolumeGrowth))
		growthData.HighGrowth = float32(utils.Round4Decimal(record.HighGrowth))
		growthData.LowGrowth = float32(utils.Round4Decimal(record.LowGrowth))
		marketData.GrowthData = growthData
		responseMarketData = append(responseMarketData, marketData)

		nextRecord := records[i+1]
		if nextRecord.FromDate == record.FromDate && nextRecord.ToDate == record.ToDate {

			nextMarketData := &pb.ResponseMarketData{
				MarketPair: nextRecord.FromRate.MarketName,
			}
			nextGrowthData := &pb.GrowthData{}

			nextGrowthData.VolumeGrowth = float32(utils.Round4Decimal(nextRecord.VolumeGrowth))
			nextGrowthData.HighGrowth = float32(utils.Round4Decimal(nextRecord.HighGrowth))
			nextGrowthData.LowGrowth = float32(utils.Round4Decimal(nextRecord.LowGrowth))
			nextMarketData.GrowthData = nextGrowthData
			responseMarketData = append(responseMarketData, nextMarketData)
			i++
		}

		fromDate := time.Unix(record.FromDate, 0).UTC()
		toDate := time.Unix(record.ToDate, 0).UTC()

		response := &pb.ResponseData{
			From:       fromDate.Format(timeFormat),
			To:         toDate.Format(timeFormat),
			MarketData: responseMarketData,
		}

		responseDataList = append(responseDataList, response)
	}
	return responseDataList
}

func processBatchGrowthData() {

}
