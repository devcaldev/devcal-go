// example demonstrates how to insert, get and list events + generating recurrences using rrule-go package
package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	client "github.com/devcaldev/devcal-go"
	"github.com/devcaldev/devcal-go/rpc"
	"github.com/teambition/rrule-go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	addr   = flag.String("addr", "devcal.fly.dev:50051", "grpc server address")
	apiKey = flag.String("apiKey", "", "api key")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	c, err := client.New(*addr, *apiKey)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	var rr string
	rr = "FREQ=WEEKLY;INTERVAL=2;BYDAY=MO,SU"

	e, err := c.InsertEvent(ctx, &rpc.InsertEventParams{
		Dtstart: timestamppb.New(time.Now()),
		Dtend:   timestamppb.New(time.Now().Add(time.Hour)),
		Rrule:   &rr,
	})
	if err != nil {
		log.Fatalf("could not insert event: %v", err)
	}
	log.Println("Event", e)
	log.Println("---")

	e, err = c.GetEvent(ctx, &rpc.GetEventParams{ID: e.GetID()})
	if err != nil {
		log.Fatalf("could not get event: %v", err)
	}
	printEvent(e)
	log.Println("---")

	s, err := c.ListEvents(ctx, &rpc.ListEventsParams{Range: &rpc.ListEventsRange{Date: timestamppb.New(time.Now()), Period: "year"}})
	if err != nil {
		log.Fatalf("could not list events: %v", err)
	}
	for {
		e, err := s.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv failed: %v", err)
		}
		printEvent(e)
	}

	_, err = c.UpdateEvent(ctx, &rpc.UpdateEventParams{ID: e.GetID(), Props: []byte(`{"name":"x"}`)})
	if err != nil {
		log.Fatalf("could not update event: %v", err)
	}

	s, err = c.ListEvents(ctx, &rpc.ListEventsParams{Props: []byte(`{"name":"x"}`)})
	if err != nil {
		log.Fatalf("could not find events: %v", err)
	}
	for {
		e, err := s.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv failed: %v", err)
		}
		printEvent(e)
	}

	_, err = c.DeleteEvent(ctx, &rpc.DeleteEventParams{ID: e.GetID()})
	if err != nil {
		log.Fatalf("could not delete event: %v", err)
	}
}

func printEvent(e *rpc.Event) {
	if e.GetRrule() != "" {
		r, err := rrule.StrToRRule(e.GetRrule() + ";DTSTART=" + e.GetDtstart().AsTime().UTC().Format(rrule.DateTimeFormat))
		if err != nil {
			log.Fatalf("could not str to rrule: %v %s", err, e.GetRrule()+";DTSTART="+e.GetDtstart().AsTime().UTC().Format(rrule.DateTimeFormat))
		}
		q := e.Dtstart.AsTime()

		log.Println("Event:", e)
		log.Println("Event Occurrences:", r.Between(
			q,
			q.Add(time.Hour*24000),
			false,
		))
	} else {
		log.Println("Event:", e)
	}
}
