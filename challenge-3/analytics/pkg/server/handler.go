package server

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/ebikode/peaq-challenge/challenge3/exchange/proto/rate"
)

// GetGrowthRecords ...
func GetGrowthRecords(ctx context.Context, rateService pb.RateServiceClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		resp := NewResponse()

		from, to, format := PaginationParams(r)

		println(from, to)

		req := &pb.GetRequest{
			FromTimestamp: from,
			ToTimestamp:   to,
		}

		records, err := rateService.GetGrowthRecords(ctx, req)
		fmt.Println(records, err)
		fmt.Println(records.Data[0].MarketData, "records.Data[0].MarketData")

		if format == jsonString {
			if err != nil {
				resp.Message(false, "Error Occurred while fetching Records.")
				resp.ErrorResponse(http.StatusBadRequest, w, r)
				return
			}

			resp.Message(true, "success")
			resp.AddCustomData("results", records.Data)
			resp.Respond(w, r)
			return
		}
		resp.Message(false, "Bad Format Supplied")
		resp.ErrorResponse(http.StatusBadRequest, w, r)
	}

}
