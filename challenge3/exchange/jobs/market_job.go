package jobs

import (
	"fmt"
	"net/http"
	"strings"
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

	for _, data := range respData.Result {

		market := data.MarketName
		// Screen required data
		if _, ok := allowedMarket[market]; ok {
			fmt.Println(data)
			fmt.Println(data.Volume)
			oldRate := rateService.GetByMarketName(market)
			tm := strings.Replace(data.Timestamp, "T", " ", -1)
			parseTimestamp, _ := time.Parse(timeFormat, tm)

			rate := models.Rate{
				MarketName: market,
				High:       data.High,
				Low:        data.Low,
				Volume:     data.Volume,
				Timestamp:  parseTimestamp,
			}
			newRate, err := rateService.CreateRate(rate)

			fmt.Println("newRate", newRate)

			if err != nil {
				fmt.Printf(`an error occurred while creating market rate: "%s"`, err.Error())

			}

			if oldRate != nil {

				volumeGrowth := utils.CalculatePercentageDifference(oldRate.Volume, newRate.Volume)
				highGrowth := utils.CalculatePercentageDifference(oldRate.High, newRate.High)
				lowGrowth := utils.CalculatePercentageDifference(oldRate.Low, newRate.Low)

				record := models.GrowthRecord{
					FromRateID:   oldRate.ID,
					ToRateID:     newRate.ID,
					From:         oldRate.Timestamp,
					To:           oldRate.Timestamp,
					VolumeGrowth: volumeGrowth,
					HighGrowth:   highGrowth,
					LowGrowth:    lowGrowth,
				}

				_, err := growthService.CreateGrowthRecord(record)

				if err != nil {
					fmt.Printf(`an error occurred while creating market growth rate: "%s"`, err.Error())
				}
			}
		}
	}
}
