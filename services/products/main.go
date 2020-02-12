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
	pb.UnimplementedProductsServiceServer
}

func (s *Server)

func main(){

}

func (s *Server) InsertProduct (ctx context.Context, req *pb.InsertProductRequest) (*pb.InsertProductResponse, error){

}

func (s *Server) GetProduct (ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {

}

func (s *Server) DeleteProduct (ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error){

}

func (s *Server) UpdateProduct (ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {

}

func (s *Server) StreamProducts (req *pb.StreamProductsRequest, stream pb.StreamProductsService_StremProductsServer) error {
	
}