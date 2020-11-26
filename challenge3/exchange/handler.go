package main

import (
	"context"

	pb "github.com/ebikode/peaq-challenge/challenge3/exchange/proto/rate"
)

type handler struct {
	repository
}

// FetchGrowths ...
func (h handler) GetGrowthRecords(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {

	records, err := h.repository.GetByTimestamp(req.FromTimestamp, req.ToTimestamp)

	recordList := UnmarshalGrowthRecordList(records)

	return &pb.Response{GrowthRecords: recordList}, err
}
