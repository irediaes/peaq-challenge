package main

import (
	"context"
	"net/http"

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
		// fmt.Println(records.Data, err)

		if format == jsonString {

			Respond(w, r, records.Data)
			return
		}
		resp := Message(false, "Bad Format Supplied")
		ErrorResponse(http.StatusBadRequest, w, r, resp)
		return
	}

}
