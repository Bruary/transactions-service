package server

import (
	"math/rand"

	"github.com/Bruary/api-service/server/pb"
	uuid "github.com/google/uuid"
)

type TransactionsServer struct {
}

func (t *TransactionsServer) CreateTransaction(req *pb.CreateTransactionRequest) (pb.CreateTransactionResponse, error) {

	return pb.CreateTransactionResponse{
		Uid:      uuid.NewString(),
		Amount:   float32(rand.Float64() * 5),
		Currency: "AED",
	}, nil
}

func (t *TransactionsServer) GetTransaction() {

}

func (t *TransactionsServer) GetTransactions() {

}
