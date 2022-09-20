package service

import (
	"fmt"

	pb "github.com/Bruary/transactions-service/models/proto"
)

func CreateTransaction(req pb.CreateTransactionRequest) (pb.CreateTransactionResponse, error) {

	fmt.Println("Ok here we go ladies and gentlemen: ", req)
	return pb.CreateTransactionResponse{}, nil
}
