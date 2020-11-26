package main

import (
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
)

func main() {

	serverAddr := os.Getenv("EXCHANGE_HOST")
	var opts []grpc.DialOption

	conn, err := grpc.Dial(*serverAddr, opts...)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	client := &pb.NewRateServiceClient(conn)

	http.HandleFunc("/export/analytics", GetGrowthRecords(client))
	http.ListenAndServe(":50052", nil)
}
