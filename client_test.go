package client_test

import (
	"context"
	"io"
	"log"
	"testing"
	"time"

	client "github.com/devcaldev/devcal-go"
	"github.com/devcaldev/devcal-go/rpc"
)

// NOTE: local devcal server must be running
// FIX ME: apiKey only available on local development machine and needs to be replaced whenever repo is cloned
var (
	addr   = "localhost:50051"
	apiKey = "YWQ5I+5PbcUmIUNdb4oZ0EsqsjZ/9heZGMcQp0jLXNDLcqZj4AwIzCJ0T6HeqQqLsDj7v4AyDt3CC33J89wZFA=="
)

func TestNewWithInsecureCredentials(t *testing.T) {
	ctx := context.Background()

	c, err := client.NewWithInsecureCredentials(addr, apiKey)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	var rr string
	rr = "FREQ=WEEKLY;INTERVAL=2;BYDAY=MO,SU"

	e, err := c.InsertEvent(ctx, &rpc.InsertEventParams{
		Dtstart: time.Now().Format(time.RFC3339),
		Dtend:   time.Now().Add(time.Hour).Format(time.RFC3339),
		Rrule:   &rr,
	})
	if err != nil {
		t.Fatalf("could not insert event: %v", err)
	}
	log.Println("Event", e)
	log.Println("---")

	e, err = c.GetEvent(ctx, &rpc.GetEventParams{ID: e.GetID()})
	if err != nil {
		t.Fatalf("could not get event: %v", err)
	}
	log.Println("Event", e)
	log.Println("---")

	s, err := c.ListEvents(ctx, &rpc.ListEventsParams{Date: time.Now().Format(time.RFC3339), Period: "year"})
	if err != nil {
		t.Fatalf("could not list events: %v", err)
	}
	for {
		e, err := s.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("stream.Recv failed: %v", err)
		}
		log.Println("Event", e)
		log.Println("---")
	}

	_, err = c.UpdateEvent(ctx, &rpc.UpdateEventParams{ID: e.GetID(), Props: []byte(`{"name":"x"}`)})
	if err != nil {
		t.Fatalf("could not update event: %v", err)
	}

	s, err = c.FindEvents(ctx, &rpc.FindEventsParams{Props: []byte(`{"name":"x"}`)})
	if err != nil {
		t.Fatalf("could not find events: %v", err)
	}
	for {
		e, err := s.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("stream.Recv failed: %v", err)
		}
		log.Println("Event", e)
		log.Println("---")
	}

	e, err = c.GetEvent(ctx, &rpc.GetEventParams{ID: e.GetID()})
	if err != nil {
		t.Fatalf("could not get event: %v", err)
	}
	log.Println("Updated Event", e)
	log.Println("---")

	_, err = c.DeleteEvent(ctx, &rpc.DeleteEventParams{ID: e.GetID()})
	if err != nil {
		t.Fatalf("could not delete event: %v", err)
	}
}
