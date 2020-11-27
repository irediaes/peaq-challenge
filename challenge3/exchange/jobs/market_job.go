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

			volume, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", data.Volume), 64)
			high, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", data.High), 64)
			low, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", data.Low), 64)

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

				vGrowth := utils.CalculatePercentageDifference(oldRate.Volume, newRate.Volume)
				hGrowth := utils.CalculatePercentageDifference(oldRate.High, newRate.High)
				lGrowth := utils.CalculatePercentageDifference(oldRate.Low, newRate.Low)

				volumeGrowth, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", vGrowth), 64)
				highGrowth, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", hGrowth), 64)
				lowGrowth, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", lGrowth), 64)

				record := models.GrowthRecord{
					FromRateID:   oldRate.ID,
					ToRateID:     newRate.ID,
					FromDate:     oldRate.Timestamp,
					ToDate:       now,
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
