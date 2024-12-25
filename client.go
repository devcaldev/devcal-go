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

// Client represents a client for interacting with the EventsService.
type Client struct {
	r rpc.EventsServiceClient // gRPC client for the EventsService.
	c *grpc.ClientConn        // gRPC client connection.
}

// InsertEventParams holds parameters for inserting an event.
// Rrule and Props fields are optional
type InsertEventParams struct {
	Dtstart time.Time      `json:"dtstart"`         // Start time of the event.
	Dtend   time.Time      `json:"dtend"`           // End time of the event.
	Rrule   string         `json:"rrule,omitempty"` // Recurrence rule for the event.
	Props   map[string]any `json:"props,omitempty"` // Additional properties of the event.
}

// GetEventParams holds parameters for retrieving an event.
type GetEventParams struct {
	ID string `json:"id"` // Unique identifier of the event.
}

// ListEventsParams holds parameters for listing events.
// Either Date and Period together, or Props, or all of them must be provided
type ListEventsParams struct {
	Date   time.Time      `json:"date,omitempty"`   // Date to filter events.
	Period string         `json:"period,omitempty"` // Period to filter events.
	Props  map[string]any `json:"props,omitempty"`  // Additional properties to filter events.
}

// UpdateEventParams holds parameters for updating an event.
// Except ID, all other fields are optional
type UpdateEventParams struct {
	ID      string         `json:"id"`                // Unique identifier of the event.
	Dtstart time.Time      `json:"dtstart,omitempty"` // New start time of the event.
	Dtend   time.Time      `json:"dtend,omitempty"`   // New end time of the event.
	Rrule   string         `json:"rrule,omitempty"`   // New recurrence rule for the event.
	Props   map[string]any `json:"props,omitempty"`   // New additional properties of the event.
}

// DeleteEventParams holds parameters for deleting an event.
type DeleteEventParams struct {
	ID string `json:"id"` // Unique identifier of the event.
}

// Event represents an event with its details.
type Event struct {
	ID        string         `json:"id"`         // Unique identifier of the event.
	AccountID string         `json:"account_id"` // Account ID associated with the event.
	Dtstart   time.Time      `json:"dtstart"`    // Start time of the event.
	Dtend     time.Time      `json:"dtend"`      // End time of the event.
	Rrule     string         `json:"rrule"`      // Recurrence rule for the event.
	Props     map[string]any `json:"props"`      // Additional properties of the event.
}

func eventFromRpcEvent(re *rpc.Event) *Event {
	return &Event{
		ID:        re.GetId(),
		AccountID: re.GetAccountId(),
		Dtstart:   re.GetDtstart().AsTime(),
		Dtend:     re.GetDtend().AsTime(),
		Rrule:     re.GetRrule(),
		Props:     re.GetProps().AsMap(),
	}
}

// InsertEvent inserts a new event and returns the created event or an error.
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

// GetEvent retrieves an event by ID and returns the event or an error.
func (c *Client) GetEvent(ctx context.Context, p *GetEventParams) (e *Event, err error) {
	rp := &rpc.GetEventParams{
		Id: p.ID,
	}

	re, err := c.r.GetEvent(ctx, rp)
	if err != nil {
		return nil, err
	}

	return eventFromRpcEvent(re), nil
}

// ListEvents lists events based on the provided parameters and returns a slice of events or an error.
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

// UpdateEvent updates an existing event and returns an error if any.
func (c *Client) UpdateEvent(ctx context.Context, p *UpdateEventParams) (err error) {
	rp := &rpc.UpdateEventParams{
		Id: p.ID,
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

// DeleteEvent deletes an event by ID and returns an error if any.
func (c *Client) DeleteEvent(ctx context.Context, p *DeleteEventParams) (err error) {
	rp := &rpc.DeleteEventParams{
		Id: p.ID,
	}

	_, err = c.r.DeleteEvent(ctx, rp)

	return
}

// Close closes the client connection.
func (c *Client) Close() error {
	return c.c.Close()
}

// New creates a new Client with secure credentials.
func New(addr, apiKey string) (*Client, error) {
	return NewWithOptions(
		addr,
		apiKey,
		grpc.WithPerRPCCredentials(&apiKeyCredentials{apiKey: apiKey}),
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
	)
}

// NewWithInsecureCredentials creates a new Client with insecure credentials.
func NewWithInsecureCredentials(addr, apiKey string) (*Client, error) {
	return NewWithOptions(
		addr,
		apiKey,
		grpc.WithPerRPCCredentials(&apiKeyCredentials{apiKey: apiKey}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}

// NewWithOptions creates a new Client with the provided gRPC dial options.
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
