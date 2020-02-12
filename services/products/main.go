package main

import (
	"context"

	pb "git.mip-consult.de/sde/suzuki-framework/microservices/api/proto/v1"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type Server struct {
	pb.UnimplementedProductsServiceServer
}

func (s *Server) InsertProduct(ctx context.Context, req *pb.InsertProductRequest) (*pb.InsertProductResponse, error) {

}

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {

}

func (s *Server) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {

}

func (s *Server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {

}

func (s *Server) StreamProducts(req *pb.StreamProductsRequest, stream pb.StreamProductsService_StremProductsServer) error {

}

func main() {

}
