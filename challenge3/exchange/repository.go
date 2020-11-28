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
