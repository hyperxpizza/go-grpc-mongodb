package main

import (
	"fmt"
	"context"
	"net"
	"log"
	
	pb "git.mip-consult.de/sde/suzuki-framework/microservices/api/proto/v1"
	db "git.mip-consult.de/sde/suzuki-framework/microservices/pkg/database"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"

	"go.mongodb.org/mongo-driver/bson"
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
			fmt.Sprintf("[-] Could not convert to ObjectID: %v",err),
		)
	}

	data := db.AccountItem{}

	result := db.GetAccountFromDB(client, ctx, oid)
	if err = result.Decode(&data); err != nil{
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("[-] Could not find account with ObjectID %s: %v", req.GetId(), err),
		)
	}

	response := &pb.GetAccountResponse{
		Account: &pb.Account {
			Id: oid.Hex(),
			Username: data.Username,
			Password: data.Password,
			Email: data.Email,
		},
	}

	return response, nil

}

func (s *Server) DeleteAccount (ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error){

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Sprintf("[-] Could not convert to ObjectID: %v",err),
		)
	}

	err = db.DeleteAccountFromDB(client, ctx, oid)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("[-] Could not find/delete Account with ObjectID %s: %v", req.GetId(), err),
		)
	}
	fmt.Println("[+] Account has been deleted from the database")

	return &pb.DeleteAccountResponse{
		Success: true,
	}, nil

}

func (s *Server) UpdateAccount (ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error){
	account := req.GetAccount()
	
	oid, err := primitive.ObjectIDFromHex(account.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("[-] Could not convert to ObjectID: %v",err),
		)
	}

	update := bson.M {
		"username": account.GetUsername(),
		"password": account.GetPassword(),
		"email": account.GetEmail(),
	}
	result := db.UpdateAccountInDB(client, ctx, oid, update)
	
	decoded := db.AccountItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("[-] Could not find Account with supplied ID %s: %v", oid, err),
		)
	}

	return &pb.UpdateAccountResponse{
		Account: &pb.Account{
			Id: decoded.ID.Hex(),
			Username: decoded.Username,
			Password: decoded.Password,
			Email: decoded.Email,
		},
	}, nil

}

func (s *Server) StreamAccounts(req *pb.StreamAccountsRequest, stream pb.AccountService_StreamAccountsServer) error {

	data := &db.AccountItem{}
	collection := client.Database(db.DBNAME).Collection("accounts")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf{
			codes.Internal,
			fmt.Sprintf("[-] Unknown Internal error: %v", err),
		}
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()){
		err := cursor.Decode(&data)
		if err != nil{
			return status.Errorf(
				codes.Unavailable,
				fmt.Sprintf("[-] Could not decode data: %v", err),
			)
		}

		stream.Send(&pb.StreamAccountsRequest{
			Account: &pb.Account{
				Id: data.ID.Hex(),
				Username: data.Username,
				Password: data.Password,
				Email: data.Email,
			},
		})
	}
	if err = cursor.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("[-] Cursor Error: %v", err),
		)
	}

	return nil
}

const port = ":8080"

func main() {
	fmt.Println("[*] Starting Server...")
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("[-] net.Listen error: %v", err)
	}

	fmt.Println("[*] Listening on port :8080")
	
	s := grpc.NewServer()
	pb.RegisterAccountsServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("[-] Serve error: %v", err)
	}

}
