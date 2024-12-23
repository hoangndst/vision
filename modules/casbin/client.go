package casbin

import (
	"context"
	pb "github.com/casbin/casbin-server/proto"
	"google.golang.org/grpc"
)

type Client struct {
	remoteClient pb.CasbinClient
}

func NewClient(_ context.Context, address string, opts ...grpc.DialOption) (*Client, error) {
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{remoteClient: pb.NewCasbinClient(conn)}, nil
}
