// example demonstrates how to insert, get and list events + generating recurrences using rrule-go package
package main

import (
	"context"
	"flag"
	"log"
	"time"

	client "github.com/devcaldev/devcal-go"
	"github.com/teambition/rrule-go"
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

	e, err := c.InsertEvent(ctx, &client.InsertEventParams{
		Dtstart: time.Now(),
		Dtend:   time.Now().Add(time.Hour),
		Rrule:   "FREQ=WEEKLY;INTERVAL=2;BYDAY=MO,SU",
	})
	if err != nil {
		log.Fatalf("could not insert event: %v", err)
	}
	log.Println("Event", e)
	log.Println("---")

	e, err = c.GetEvent(ctx, &client.GetEventParams{ID: e.ID})
	if err != nil {
		log.Fatalf("could not get event: %v", err)
	}
	printEvent(e)
	log.Println("---")

	es, err := c.ListEvents(ctx, &client.ListEventsParams{Date: time.Now(), Period: "year"})
	if err != nil {
		log.Fatalf("could not list events: %v", err)
	}
	for _, e := range es {
		printEvent(e)
	}

	err = c.UpdateEvent(ctx, &client.UpdateEventParams{ID: e.ID, Props: map[string]any{"name": "x"}})
	if err != nil {
		log.Fatalf("could not update event: %v", err)
	}

	es, err = c.ListEvents(ctx, &client.ListEventsParams{Props: map[string]any{"name": "x"}})
	if err != nil {
		log.Fatalf("could not find events: %v", err)
	}
	for _, e := range es {
		printEvent(e)
	}

	err = c.DeleteEvent(ctx, &client.DeleteEventParams{ID: e.ID})
	if err != nil {
		log.Fatalf("could not delete event: %v", err)
	}
}

func printEvent(e *client.Event) {
	if e.Rrule != "" {
		r, err := rrule.StrToRRule(e.Rrule + ";DTSTART=" + e.Dtstart.UTC().Format(rrule.DateTimeFormat))
		if err != nil {
			log.Fatalf("could not str to rrule: %v %s", err, e.Rrule+";DTSTART="+e.Dtstart.UTC().Format(rrule.DateTimeFormat))
		}
		q := e.Dtstart

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
