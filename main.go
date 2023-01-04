package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Bruary/transactions-service/db"
	pb "github.com/Bruary/transactions-service/server/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type TransactionsServer struct {
	pb.UnimplementedTransactionsServer
}

func (s *TransactionsServer) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	fmt.Println("The request: ", req.GetAmount(), req.GetCurrency())
	user_uid := uuid.New()

	return &pb.CreateTransactionResponse{
		Uid:      user_uid.String(),
		Amount:   req.GetAmount(),
		Currency: req.GetCurrency(),
	}, nil
}

func main() {

	//establish DB connection
	db.EstablishDBConnection()

	// open connection
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	// Create a new grpc server
	s := grpc.NewServer()

	// register the transactions server
	pb.RegisterTransactionsServer(s, &TransactionsServer{})
	log.Printf("Server listening at %v ", lis.Addr())

	// start reading and parsing the input
	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}

}
