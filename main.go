package main

import (
	"log"
	"net"

	"github.com/Bruary/transactions-service/db"
	pb "github.com/Bruary/transactions-service/server/pb"
	"github.com/Bruary/transactions-service/service"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {

	//establish DB connection
	ConnectToDB()

	// register service
	RegisterService()
}

func ConnectToDB() {
	db.EstablishDBConnection()
}

func RegisterService() {
	// open connection
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	// Create a new grpc server
	s := grpc.NewServer()

	// register the transactions server
	pb.RegisterTransactionsServer(s, &service.TransactionsServer{})
	log.Printf("Server listening at %v ", lis.Addr())

	// start reading and parsing the input
	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
