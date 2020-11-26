package handler

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/ebikode/peaq-challenge/challenge3/exchange/proto/rate"
)

func GetGrowthRecords(rateService pb.RateServiceClient) http.HandlerFunc {

	return func(w http.WriterResponse, r *http.Request) {

		ctx := context.Background()

		from, to, format := PaginationParams(r)

		req := &pb.GetRequest{from, to}

		records, err := rateService.GetGrowthRecords(ctx, req)
		fmt.Println(records, err)

		if format == jsonString {
			resp := records

			Respond(w, r, resp)
			return
		}
		resp := Message(false, "Bad Format Supplied")
		ErrorResponse(http.BadGatewayStatus, w, r, resp)
		return
	}

}
