package server

import (
	"context"
	"net/http"

	_ "github.com/ebikode/peaq-challenge/challenge-3/exchange/proto/rate"
	pb "github.com/ebikode/peaq-challenge/challenge-3/exchange/proto/rate"
)

// GetGrowthRecords ...
// Fetch Growth Records
//
// Fetch all Growth record data in five minutes interval
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
		// fmt.Println(records, err)

		if format == jsonString {
			if err != nil {
				resp.Message(false, "Error Occurred while fetching Records.")
				resp.ErrorResponse(http.StatusBadRequest, w, r)
				return
			}

			resp.Message(true, "success")
			resp.AddCustomData("results", records.Data)
			// UNCOMMENT THIS IF YOU NEED THE FULL RESULTS DATA
			// resp.AddCustomData("results_records", records.GrowthRecords)
			resp.Respond(w, r)
			return
		}
		resp.Message(false, "Bad Format Supplied")
		resp.ErrorResponse(http.StatusBadRequest, w, r)
	}

}
