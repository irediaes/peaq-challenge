package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "github.com/ebikode/peaq-challenge/challenge3/exchange/proto/rate"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := os.Getenv("EXCHANGE_HOST")
	host := fmt.Sprintf(":%s", os.Getenv("PORT"))
	// host := ":50052"
	// serverAddr := ":50051"
	// var opts []grpc.DialOption

	fmt.Printf("gPRC Server started on %s \n\n", serverAddr)
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()

	client := pb.NewRateServiceClient(conn)

	fmt.Printf("Server started on %s \n\n", host)
	http.HandleFunc("/export/analytics", GetGrowthRecords(ctx, client))
	http.ListenAndServe(host, nil)

}
