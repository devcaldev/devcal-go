package client

import (
	"context"
	"io"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"
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
	r rpc.EventsServiceClient
	c *grpc.ClientConn
}

type InsertEventParams struct {
	Dtstart time.Time
	Dtend   time.Time
	Rrule   string
	Props   map[string]any
}

type GetEventParams struct {
	ID string
}

type ListEventsParams struct {
	Date   time.Time
	Period string
	Props  map[string]any
}

type UpdateEventParams struct {
	ID      string
	Dtstart time.Time
	Dtend   time.Time
	Rrule   string
	Props   map[string]any
}

type DeleteEventParams struct {
	ID string
}

type Event struct {
	ID        string
	AccountID string
	Dtstart   time.Time
	Dtend     time.Time
	Rrule     string
	Props     map[string]any
}

func eventFromRpcEvent(re *rpc.Event) *Event {
	return &Event{
		ID:        re.GetID(),
		AccountID: re.GetAccountID(),
		Dtstart:   re.GetDtstart().AsTime(),
		Dtend:     re.GetDtend().AsTime(),
		Rrule:     re.GetRrule(),
		Props:     re.GetProps().AsMap(),
	}
}

func (c *Client) InsertEvent(ctx context.Context, p *InsertEventParams) (e *Event, err error) {
	rp := &rpc.InsertEventParams{
		Dtstart: timestamppb.New(p.Dtstart),
		Dtend:   timestamppb.New(p.Dtend),
	}
	if p.Rrule != "" {
		rp.Rrule = &p.Rrule
	}
	if p.Props != nil {
		rp.Props, err = structpb.NewStruct(p.Props)
		if err != nil {
			return nil, err
		}
	}

	re, err := c.r.InsertEvent(ctx, rp)
	if err != nil {
		return nil, err
	}

	return eventFromRpcEvent(re), nil
}

func (c *Client) GetEvent(ctx context.Context, p *GetEventParams) (e *Event, err error) {
	rp := &rpc.GetEventParams{
		ID: p.ID,
	}

	re, err := c.r.GetEvent(ctx, rp)
	if err != nil {
		return nil, err
	}

	return eventFromRpcEvent(re), nil
}

func (c *Client) ListEvents(ctx context.Context, p *ListEventsParams) (es []*Event, err error) {
	rp := &rpc.ListEventsParams{}
	if !p.Date.IsZero() && p.Period != "" {
		rp.Range = &rpc.ListEventsRange{Date: timestamppb.New(p.Date), Period: p.Period}
	}
	if p.Props != nil {
		rp.Props, err = structpb.NewStruct(p.Props)
		if err != nil {
			return nil, err
		}
	}

	res, err := c.r.ListEvents(ctx, rp)
	if err != nil {
		return nil, err
	}
	es = make([]*Event, 0)
	for {
		e, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		es = append(es, eventFromRpcEvent(e))
	}

	return es, nil
}

func (c *Client) UpdateEvent(ctx context.Context, p *UpdateEventParams) (err error) {
	rp := &rpc.UpdateEventParams{
		ID: p.ID,
	}
	if !p.Dtstart.IsZero() {
		rp.Dtstart = timestamppb.New(p.Dtstart)
	}
	if !p.Dtend.IsZero() {
		rp.Dtend = timestamppb.New(p.Dtend)
	}
	if p.Rrule != "" {
		rp.Rrule = &p.Rrule
	}
	if p.Props != nil {
		rp.Props, err = structpb.NewStruct(p.Props)
		if err != nil {
			return err
		}
	}

	_, err = c.r.UpdateEvent(ctx, rp)

	return
}

func (c *Client) DeleteEvent(ctx context.Context, p *DeleteEventParams) (err error) {
	rp := &rpc.DeleteEventParams{
		ID: p.ID,
	}

	_, err = c.r.DeleteEvent(ctx, rp)

	return
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
		r: rpc.NewEventsServiceClient(conn),
		c: conn,
	}

	return c, nil
}
