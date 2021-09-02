package main

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/ebikode/peaq-challenge/challenge3/exchange/proto/rate"
)

// GetGrowthRecords ...
func GetGrowthRecords(ctx context.Context, rateService pb.RateServiceClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		from, to, format := PaginationParams(r)

		println(from, to)

		req := &pb.GetRequest{
			FromTimestamp: from,
			ToTimestamp:   to,
		}

		records, err := rateService.GetGrowthRecords(ctx, req)
		fmt.Println(records, err)

		if format == jsonString {
			if err != nil {
				resp := Message(false, "Error Occurred while fetching Records.")
				ErrorResponse(http.StatusBadRequest, w, r, resp)
				return
			}

			Respond(w, r, records.Data)
			return
		}
		resp := Message(false, "Bad Format Supplied")
		ErrorResponse(http.StatusBadRequest, w, r, resp)
		return
	}

}
