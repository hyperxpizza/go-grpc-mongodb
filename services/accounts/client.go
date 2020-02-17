package account

import (
	"context"

	pb "git.mip-consult.de/sde/suzuki-framework/microservices/api/proto/v1"
	"git.mip-consult.de/sde/suzuki-framework/microservices/pkg/database"
	"google.golang.org/grpc"
)

type Client struct {
	connection *grpc.ClientConn
	service    pb.AcccountServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	s := pb.NewAccountServiceClient(connection)
	return &Client{connection: conn, service: s}, nil
}

func (c *Client) CloseConnection() {
	c.connection.Close()
}

func (c *Client) PostAccount(ctx context.Context, username, password, email string) (Account, error) {
	r, err := c.service.InsertAccount(
		ctx,
		&pb.InsertAccountRequest{
			Account: &pb.Account{
				Username: username,
				Password: password,
				Email: email,
			},
		},
	)
	if err != nil{
		return nil, err
	}

	return &
}

func (c *Client) GetAccount(ctx context, id string) (Account, error){

	return Account{}, nil
}

func (c *Client) DeleteAccount(ctx context.Context, id string) (bool, error){
	
	return success, nil
}

func (c *Client) UpdateAccount(ctx context.Context,username, password, email string) (Account, error){

	return Account{}, nil
}

func (c *Client) StreamAccounts() ([Accounts, error){

	return accounts, nil
}
