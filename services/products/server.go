package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "git.mip-consult.de/sde/suzuki-framework/microservices/api/proto/v1"
	db "git.mip-consult.de/sde/suzuki-framework/microservices/pkg/database"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type Server struct {
	pb.UnimplementedProductServiceServer
}

func (s *Server) InsertProduct(ctx context.Context, req *pb.InsertProductRequest) (*pb.InsertProductResponse, error) {
	product := req.GetProduct()
	data := db.ProductItem{
		Name:        product.GetName(),
		Price:       product.GetPrice(),
		Description: product.GetDescription(),
	}

	result, err := db.InsertProductToTheDB(ctx, client, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("[-] Insert Product error: %v"),
		)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	product.Id = oid.Hex()

	return &pb.InsertProductResponse{Product: product}, nil

}

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("[-] Could not convert to ObjectID: %v", err),
		)
	}

	data := db.ProductItem{}

	result := db.GetProductFromDB(ctx, client, oid)
	if err = result.Decode(&data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("[-] Could not find account with ObjectID %s: %v", req.GetId(), err),
		)
	}

	response := &pb.GetProductResponse{
		Product: &pb.Product{
			Id:          oid.Hex(),
			Name:        data.Name,
			Price:       data.Price,
			Description: data.Description,
		},
	}

	return response, nil
}

func (s *Server) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("[-] Could not convert to ObjectID %s: %v", req.GetId(), err),
		)
	}

	err = db.DeleteProductFromDB(ctx, client, oid)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound, fmt.Sprintf("[-] Could not find/delete Account with ObjectID %s: %v", req.GetId(), err),
		)
	}
	fmt.Println("[+] Product has been deleted from the database")

	return &pb.DeleteAccountResponse{
		Success: true,
	}, nil
}

func (s *Server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {

	product := req.GetProduct()
	oid, err := primitive.ObjectIDFromHex(product.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("[-] Could not convert to ObjectID: %v", err),
		)
	}

	update := bson.M{
		"name":        product.GetName(),
		"price":       product.GetPrice(),
		"description": product.GetDescription(),
	}

	result := db.UpdateProductInDB(ctx, client, oid, update)

	decoded := db.ProductItem{}
	err = result.Decode((&decoded))
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("[-] Could not find Account with supplied ID %s: %v", oid, err),
		)
	}

	return &pb.UpdateProductResponse{
		Product: &pb.Product{
			Id:          decoded.ID.Hex(),
			Name:        decoded.Name,
			Price:       decoded.Price,
			Description: decoded.Description,
		},
	}, nil
}

func (s *Server) StreamProducts(req *StreamProductsRequest, stream ProductService_StreamProductsServer) error {

	data := &db.ProductItem{}
	collection := client.Database(db.DBNAME).Collection("products")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		if err != nil {
			return err
		}

		stream.Send(&pb.StreamProductsResponse{
			Product: &pb.Product{
				Id:          data.ID.Hex(),
				Name:        data.Name,
				Price:       data.Price,
				Description: data.Description,
			},
		})

		if err = cursor.Err(); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("[-] Cursor Error: %v", err),
			)
		}

	}

	return nil
}

func ListenGRPC() error {
	client = db.ConnectToDB()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterProductServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)

	return serv.Serve(lis)
}
