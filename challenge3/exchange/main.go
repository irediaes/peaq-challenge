package main

// import (
// 	"fmt"
// 	"math/big"
// )

import (
	"log"
	"net"

	"github.com/ebikode/peaq-challenge/challenge3/exchange/jobs"
	"github.com/ebikode/peaq-challenge/challenge3/exchange/pkg/growth"
	pb "github.com/ebikode/peaq-challenge/challenge3/exchange/proto/rate"
	storage "github.com/ebikode/peaq-challenge/challenge3/exchange/storage/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	// initialize DB
	dbConfig := storage.New()
	mdb, err := dbConfig.InitDB()

	// if an error occurred while initialising db
	if err != nil {
		log.Println(err)
	}

	growthStorage := storage.NewGrowthRecordStorage(mdb)

	repo := growth.NewService(growthStorage)

	h := &handler{repo}

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterRateServiceServer(s, h)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Init Jobs
	jobs.Init(mdb)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// func main() {
// 	oldNum := float64(8.0497183772403904E+17)
// 	newNum := big.NewRat(1, 1)
// 	newNum.SetFloat64(oldNum)
// 	fmt.Println(newNum.FloatString(0))
// }
