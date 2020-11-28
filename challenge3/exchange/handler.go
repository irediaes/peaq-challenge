package main

import (
	"context"
	"os"
	"strconv"

	pb "github.com/ebikode/peaq-challenge/challenge3/exchange/proto/rate"
)

type handler struct {
	repository
}

// FetchGrowths ...
func (h handler) GetGrowthRecords(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {

	numberOfCrytoMarket := os.Getenv("NUMBER_MARKET")
	numMarket, _ := strconv.ParseInt(numberOfCrytoMarket, 10, 64)

	records, err := h.repository.GetByTimestamp(req.FromTimestamp, req.ToTimestamp)

	response := &pb.Response{}

	response.Data = UnmarshalProtoResponseData(records, int(numMarket))

	return response, err
}
