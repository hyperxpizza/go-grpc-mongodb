package accounts

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

	s := pb.NewAccountServiceClient(conn)
	return &Client{connection: conn, service: s}, nil
}

func (c *Client) CloseConnection() {
	c.connection.Close()
}

func (c *Client) PostAccount(ctx context.Context, username, password, email string) (*Account, error) {
	
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

func (c *Client) GetAccount(ctx context, id string) (*Account, error){
	
	r, err := c.service.GetAccount{
		ctx,
		&pb.GetAccountRequest{
			Id: id,
		}
	}
	if err != nil{
		return nil, err
	}

	return Account{
		ID: r.Account.Id,
		Username: r.Account.Username,
		Password: r.Account.Password,
		Email: r.Account.Email,
	}, nil
}

func (c *Client) DeleteAccount(ctx context.Context, id string) (bool, error){
	
	r, err := c.service.DeleteAccount(
		ctx,
		&pb.DeleteAccountRequest{
			Id: id,
		}
	)
	if err != nil{
		return nil, err
	}
	
	success := r.Success

	return success, nil
}

func (c *Client) UpdateAccount(ctx context.Context, id, username, password, email string) (*Account, error){

	r, err := c.service.UpdateAccount(
		ctx,
		&pb.UpdateAccountRequest{
			Account: &pb.Account{
				Id: id,
				Username: username,
				Password: password,
				Email: email,
			},
		}
	)
	if err != nil{
		return nil, err
	}

	account := &Account{
		ID: r.Account.Id,
		Username: r.Account.Username,
		Password: r.Account.Password,
		Email: r.Account.Email,
	}

	return account, nil
}
/*
func (c *Client) ListAccounts() ([Accounts, error){

	return accounts, nil
}
*/