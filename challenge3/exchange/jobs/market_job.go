package jobs

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ebikode/peaq-challenge/challenge3/exchange/models"
	"github.com/ebikode/peaq-challenge/challenge3/exchange/pkg/growth"
	"github.com/ebikode/peaq-challenge/challenge3/exchange/pkg/rate"
	"github.com/ebikode/peaq-challenge/challenge3/exchange/utils"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	url        = "https://bittrex.com/api/v1.1/public/getmarketsummaries"
	btcAda     = "BTC-ADA"
	ethAda     = "ETH-ADA"
)

var allowedMarket = map[string]string{
	btcAda: btcAda,
	ethAda: ethAda,
}

func getMarketData(rateService rate.Service, growthService growth.Service) {

	respData := &models.MarketResponse{}

	err := utils.SendRequestAndParseResponse(http.MethodGet, url, nil, respData)

	if err != nil {
		fmt.Printf(`an error occurred while fetching market data: "%s"`, err.Error())
		return
	}

	now := time.Now()
	for _, data := range respData.Result {
		market := data.MarketName
		// Screen required data
		if _, ok := allowedMarket[market]; ok {

			oldRate := rateService.GetByMarketName(market)

			volume := data.Volume
			high := data.High
			low := data.Low

			rate := models.Rate{
				MarketName: market,
				High:       high,
				Low:        low,
				Volume:     volume,
				Timestamp:  now,
			}
			newRate, err := rateService.CreateRate(rate)

			if err != nil {
				fmt.Printf(`an error occurred while creating market rate: "%s"`, err.Error())

			}

			if oldRate != nil {

				volumeGrowth := utils.CalculatePercentageDifference(newRate.Volume, oldRate.Volume)
				highGrowth := utils.CalculatePercentageDifference(newRate.High, oldRate.High)
				lowGrowth := utils.CalculatePercentageDifference(newRate.Low, oldRate.Low)

				record := models.GrowthRecord{
					FromRateID:   oldRate.ID,
					ToRateID:     newRate.ID,
					FromDate:     oldRate.Timestamp.Unix(),
					ToDate:       now.Unix(),
					VolumeGrowth: volumeGrowth,
					LowGrowth:    lowGrowth,
					HighGrowth:   highGrowth,
				}

				// b, _ := json.Marshal(record)

				// fmt.Println("======", string(b), "======")
				// fmt.Println(record.VolumeGrowth)
				// fmt.Println(record.HighGrowth)
				// fmt.Println(record.LowGrowth)

				_, err := growthService.CreateGrowthRecord(record)

				if err != nil {
					fmt.Printf(`an error occurred while creating market growth rate: "%s"`, err.Error())
				}
			}
		}
	}
}

func formatFloat(value float64) float64 {
	str := strconv.FormatFloat(value, 'f', -1, 64)
	result, _ := strconv.ParseFloat(str, 64)

	return result
}
