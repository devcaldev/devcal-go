package client

import (
	"context"
	"encoding/json"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

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

type Client struct {
	events rpc.EventsServiceClient
	c      *grpc.ClientConn
}

type Event struct {
	ID        string
	AccountID string
	Dtstart   time.Time
	Dtend     time.Time
	Rrule     string
	Props     map[string]any
}

type InsertEventParams struct {
	Dtstart time.Time
	Dtend   time.Time
	Rrule   string
	Props   map[string]any
}

func (c *Client) InsertEvent(p *InsertEventParams) (*Event, error) {
	rp := &rpc.InsertEventParams{
		Dtstart: timestamppb.New(p.Dtstart),
		Dtend:   timestamppb.New(p.Dtend),
		Rrule:   &p.Rrule,
	}
	rp.Props, _ = json.Marshal(p.Props)

	re, err := c.events.InsertEvent(
		context.Background(),
		rp,
	)
	if err != nil {
		return nil, err
	}
	e := &Event{
		ID:        re.ID,
		AccountID: re.AccountID,
		Dtstart:   re.Dtstart.AsTime(),
		Dtend:     re.Dtend.AsTime(),
		Rrule:     re.Rrule,
	}
	json.Unmarshal(re.Props, &e.Props)
	return e, nil
}

func (c *Client) Close() error {
	return c.c.Close()
}

func New(addr, apiKey string) (*Client, error) {
	return NewWithOptions(
		addr,
		apiKey,
		grpc.WithPerRPCCredentials(&apiKeyCredentials{apiKey: apiKey}),
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
	)
}

func NewWithInsecureCredentials(addr, apiKey string) (*Client, error) {
	return NewWithOptions(
		addr,
		apiKey,
		grpc.WithPerRPCCredentials(&apiKeyCredentials{apiKey: apiKey}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}

func NewWithOptions(addr string, apiKey string, opts ...grpc.DialOption) (*Client, error) {
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, err
	}

	c := &Client{
		events: rpc.NewEventsServiceClient(conn),
		c:      conn,
	}

	return c, nil
}
