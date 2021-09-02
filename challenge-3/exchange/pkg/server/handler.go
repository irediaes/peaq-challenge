package server

import (
	"context"
	"os"
	"strconv"

	pb "github.com/ebikode/peaq-challenge/challenge3/exchange/proto/rate"
)

type exchangeServiceServer struct {
	repository Repository
	pb.UnimplementedRateServiceServer
}

func NewExchangeServiceServer(repo Repository) *exchangeServiceServer {
	return &exchangeServiceServer{repository: repo}
}

// FetchGrowths ...
func (h *exchangeServiceServer) GetGrowthRecords(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {

	numberOfCrytoMarket := os.Getenv("NUMBER_MARKET")
	numMarket, _ := strconv.ParseInt(numberOfCrytoMarket, 10, 64)

	records, err := h.repository.GetByTimestamp(req.FromTimestamp, req.ToTimestamp)

	response := &pb.Response{}

	response.Data = UnmarshalProtoResponseData(records, int(numMarket))

	return response, err
}
