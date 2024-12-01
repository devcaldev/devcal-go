package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/devcaldev/devcal-go/rpc"
)

type apiKeyCredentials struct {
	apiKey                   string
	requireTransportSecurity bool
}

func (c *apiKeyCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {

	return map[string]string{
		"authorization": "Bearer " + c.apiKey,
	}, nil
}
func (c *apiKeyCredentials) RequireTransportSecurity() bool {
	return c.requireTransportSecurity
}

type grpcClient struct {
	rpc.EventsServiceClient
	c *grpc.ClientConn
}

func (c *grpcClient) Close() error {
	return c.c.Close()
}

type Client interface {
	rpc.EventsServiceClient
	Close() error
}

func New(addr string, apiKey string) (Client, error) {
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(&apiKeyCredentials{apiKey: apiKey}),
		// for localhost
		// grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
	}

	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, err
	}

	c := &grpcClient{
		rpc.NewEventsServiceClient(conn),
		conn,
	}

	return c, nil
}
