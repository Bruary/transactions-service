package main

import (
	"log"
	"net"

	"github.com/Bruary/transactions-service/server/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type TransactionsServer struct {
	pb.UnimplementedTransactionsServer
}

func main() {

	// open connection
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}
	// Create a new grpc server
	s := grpc.NewServer()

	// register the transactions server
	pb.RegisterTransactionsServer(s, &TransactionsServer{})
	log.Printf("Server listening at %v ", lis.Addr())

	// start reading and parsing the input
	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}

}
