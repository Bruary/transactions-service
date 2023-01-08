package service

import (
	"context"
	"fmt"

	"github.com/Bruary/transactions-service/server/pb"
	"github.com/google/uuid"
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
