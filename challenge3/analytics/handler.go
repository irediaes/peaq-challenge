package main

import (
	"context"
	"net/http"
	"time"

	"github.com/ebikode/peaq-challenge/challenge3/analytics/models"
	pb "github.com/ebikode/peaq-challenge/challenge3/exchange/proto/rate"
)

// GetGrowthRecords ...
func GetGrowthRecords(rateService pb.RateServiceClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()

		from, to, format := PaginationParams(r)

		println(from, to)

		req := &pb.GetRequest{
			FromTimestamp: from,
			ToTimestamp:   to,
		}

		records, _ := rateService.GetGrowthRecords(ctx, req)
		// fmt.Println(records, err)

		if format == jsonString {
			parsedRecords := unmarshalProtoGrowthRecordList(records.GrowthRecords)

			responseDataList := []*models.Response{}

			for i := 0; i < len(parsedRecords); i++ {

				record := parsedRecords[i]
				dataList := []*models.GrowthData{}

				data := &models.GrowthData{
					MarketPair: record.FromRate.MarketName,
				}
				data.Data.VolumeGrowth = record.VolumeGrowth
				data.Data.HighGrowth = record.HighGrowth
				data.Data.LowGrowth = record.LowGrowth
				dataList = append(dataList, data)

				nextRecord := parsedRecords[i+1]
				if nextRecord.FromDate == record.FromDate && nextRecord.ToDate == record.ToDate {
					nextData := &models.GrowthData{
						MarketPair: nextRecord.FromRate.MarketName,
					}
					nextData.Data.VolumeGrowth = nextRecord.VolumeGrowth
					nextData.Data.HighGrowth = nextRecord.HighGrowth
					nextData.Data.LowGrowth = nextRecord.LowGrowth
					dataList = append(dataList, nextData)
					i++
				}

				response := &models.Response{
					From: record.FromDate,
					To:   record.ToDate,
					Data: dataList,
				}

				responseDataList = append(responseDataList, response)
			}

			Respond(w, r, responseDataList)
			return
		}
		resp := Message(false, "Bad Format Supplied")
		ErrorResponse(http.StatusBadRequest, w, r, resp)
		return
	}

}

// unmarshalGrowthRecordList ...
func unmarshalProtoGrowthRecordList(records []*pb.GrowthRecord) []*models.GrowthRecord {
	recordList := []*models.GrowthRecord{}
	for _, record := range records {
		parsedRecord := unmarshalProtoGrowthRecord(record)
		recordList = append(recordList, parsedRecord)
	}
	return recordList
}

// unmarshalGrowthRecord ...
func unmarshalProtoGrowthRecord(rate *pb.GrowthRecord) *models.GrowthRecord {

	fromDate, _ := time.Parse(timeFormat, rate.From)
	toDate, _ := time.Parse(timeFormat, rate.To)

	return &models.GrowthRecord{
		FromRateID:   uint(rate.FromRateId),
		ToRateID:     uint(rate.ToRateId),
		VolumeGrowth: float64(rate.VolumeGrowth),
		HighGrowth:   float64(rate.HighGrowth),
		LowGrowth:    float64(rate.LowGrowth),
		FromDate:     fromDate,
		ToDate:       toDate,
		FromRate:     unmarshalProtoRate(rate.FromRate),
		ToRate:       unmarshalProtoRate(rate.ToRate),
	}
}

// unmarshalGrowthRecord ...
func unmarshalProtoRate(rate *pb.Rate) models.Rate {

	return models.Rate{
		MarketName: rate.MarketName,
		Volume:     float64(rate.Volume),
		High:       float64(rate.High),
		Low:        float64(rate.Low),
	}
}
