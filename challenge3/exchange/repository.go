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
func UnmarshalProtoResponseData(records []*models.GrowthRecord, numberOfCrytoMarket int) []*pb.ResponseData {

	responseDataList := []*pb.ResponseData{}

	for i := 0; i < len(records); i++ {

		responseMarketData := []*pb.ResponseMarketData{}

		fromDate := time.Unix(records[i].FromDate, 0).UTC()
		toDate := time.Unix(records[i].ToDate, 0).UTC()

		for n := 0; n < numberOfCrytoMarket; n++ {
			record := records[i+n]
			marketData := processMarketData(record)
			responseMarketData = append(responseMarketData, marketData)
		}
		i += (numberOfCrytoMarket - 1)

		response := &pb.ResponseData{
			From:       fromDate.Format(timeFormat),
			To:         toDate.Format(timeFormat),
			MarketData: responseMarketData,
		}

		responseDataList = append(responseDataList, response)
	}
	return responseDataList
}

func processMarketData(record *models.GrowthRecord) *pb.ResponseMarketData {
	marketData := &pb.ResponseMarketData{
		MarketPair: record.FromRate.MarketName,
	}

	growthData := &pb.GrowthData{}

	growthData.VolumeGrowth = float32(utils.Round4Decimal(record.VolumeGrowth))
	growthData.HighGrowth = float32(utils.Round4Decimal(record.HighGrowth))
	growthData.LowGrowth = float32(utils.Round4Decimal(record.LowGrowth))
	marketData.GrowthData = growthData

	return marketData
}
