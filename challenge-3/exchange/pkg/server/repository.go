package server

import (
	"time"

	"github.com/ebikode/peaq-challenge/challenge-3/exchange/models"
	pb "github.com/ebikode/peaq-challenge/challenge-3/exchange/proto/rate"
	"github.com/ebikode/peaq-challenge/challenge-3/exchange/utils"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type Repository interface {
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

	growthData.VolumeGrowth = utils.Round4Decimal(record.VolumeGrowth)
	growthData.HighGrowth = utils.Round4Decimal(record.HighGrowth)
	growthData.LowGrowth = utils.Round4Decimal(record.LowGrowth)
	marketData.GrowthData = growthData

	return marketData
}
