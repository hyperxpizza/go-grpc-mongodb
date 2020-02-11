package main

import (
	"fmt"
	"context"
	
	pb "git.mip-consult.de/sde/suzuki-framework/microservices/api/proto/v1"
	db "git.mip-consult.de/sde/suzuki-framework/microservices/pkg/database"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)

var client *mongo.Client

type Server struct {
	pb.UnimplementedAcccountServiceServer
}

func (s *Server) InsertAccount (ctx context.Context, req *pb.InsertAccountRequest) (*pb.InsertAccountResponse, error){
	account := req.GetAccount()
	data := db.AccountItem{
		//ID empty, so it gets omitted
		Username: account.GetUsername(),
		Password: account.GetPassword(),
		Email: account.GetEmail(),
	}
	
	result, err := db.InsertAccountToTheDatabase(client, ctx, data)
	if err != nil{
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("[-] Insert Account error: %v"),
		)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	account.Id = oid.Hex()

	return &pb.InsertAccountResponse{Account: account}, nil

}

func (s *Server) GetAccount (ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error){
	
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Sprintf("[-] Could not convert to ObjectID: %v",err)
		)
	}
}

func (s *Server) DeleteAccount (ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error){

}

func (s *Server) UpdateAccount (ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error)
{

}

func (s *Server) StreamAccounts(req *api.StreamAccountsRequest, stream api.AccountService_StreamAccountsServer) error {
	
}

func main() {
	client = db.ConnectToDB()
	fmt.Println("Hello World")
	db.CloseConnectionToDB(client)
}
