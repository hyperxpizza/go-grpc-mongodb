package products

import (
	"context"

	pb "git.mip-consult.de/sde/suzuki-framework/microservices/api/proto/v1"
	"github.com/golang/protobuf/protoc-gen-go/grpc"
)

type Client struct {
	connection *grpc.ClientConn
	service    pb.ProductServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	s := pb.NewProductServiceClient(conn)
	return &Client{connection: conn, service: s}, nil
}

func (c *Client) CloseConnection() {
	c.connection.Close()
}

func (c *Client) PostProduct(ctx context.Context, name string, price float32, description string) (*Product, error){

	r, err := c.service.InsertProduct(
		ctx,
		&pb.InsertProductRequest{
			Product: &&pb.Product{
				Name: name,
				Price: price,
				Description: description,
			}
		}
	)
	if err != nil{
		return nil, err
	}

	return &Product{
		ID: r.Product.Id,
		Name: r.Product.Name,
		Price: r.Product.Price
		Description: r.Product.Description,
	}, nil

}

func (c *Client) GetProduct(ctx context.Context, id string) (*Product, error){
	
	r, err := c.service.GetProduct{
		&pb.GetProductRequest{
			Id: id,
		},
	}
	if err != nil{
		return nil, err
	}

	return &Product{
		ID: r.Product.Id,
		Name: r.Product.Name,
		Price: r.Product.Price,
		DescDescription: r.Product.Description,
	}, nil
}

//var success - boolean
func (c *Client) DeleteProduct(ctx context.Context, id string) (bool, error){

	r, err := c.service.DeleteProduct(
		ctx,
		&pb.DeleteAccountRequest{
			Id: id,
		},
	)
	if err != nil {
		return nil, err
	}

	success := r.Success
	return success, nil
}

func (c *Client) UpdateProduct(ctx context.Context, id, name string, price float32, description string) (*Product, error){

	r, err := c.service.UpdateProduct(
		ctx,
		&pb.UpdateProductRequest{
			Product: &pb.Product{
				Id: id,
				Name: name,
				Price: price,
				Description: description,
			},
		},
	)
	if err != nil{
		return nil, err
	}

	return &Product{
		ID: r.Product.Id,
		Name: r.Prodcut.Name,
		Price: r.Product.Price,
		Description: r.Product.Description,
	}, nil
}

/* Stream function - to be developed
func (c *Client) ListProducts() ([]Products, error)

*/