package main

import (
	"log"
	"net/http"

	pb "github.com/ebikode/peaq-challenge/challenge3/exchange/proto/rate"
	"google.golang.org/grpc"
)

func main() {

	// serverAddr := os.Getenv("EXCHANGE_HOST")
	serverAddr := ":50051"
	// var opts []grpc.DialOption

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	client := pb.NewRateServiceClient(conn)

	http.HandleFunc("/export/analytics", GetGrowthRecords(client))
	http.ListenAndServe(":50052", nil)
}
