package client_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	client "github.com/devcaldev/devcal-go"
)

var (
	// NOTE: local devcal server must be running
	addr = "localhost:50051"
	// NOTE: obtain test apiKey from local devcal server
	apiKey = "Gtr4D2lXzMy+oVS2Y2rrWyiDQ81tWM/cpD5EwvA9VJJtA7E1Tx1HnT1Moqbt36DYEcivCHDbeJi6GxhnPMuNxw=="
)

func TestNewWithInsecureCredentials(t *testing.T) {
	ctx := context.Background()

	c, err := client.NewWithInsecureCredentials(addr, apiKey)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	dtstart := time.Now()
	dtend := time.Now().Add(time.Hour)
	e, err := c.InsertEvent(ctx, &client.InsertEventParams{
		Dtstart: dtstart,
		Dtend:   dtend,
		Rrule:   "FREQ=WEEKLY;INTERVAL=2;BYDAY=MO,SU",
	})
	if err != nil {
		t.Errorf("could not insert event: %v", err)
	}
	if e.ID == "" {
		t.Error("inserted event has no id")
	}
	if e.AccountID == "" {
		t.Error("inserted event has no account id")
	}
	if !reflect.DeepEqual(e.Dtstart.UTC(), dtstart.UTC()) {
		t.Error("inserted event dtstart not equal", e.Dtstart, dtstart)
	}
	if !reflect.DeepEqual(e.Dtend.UTC(), dtend.UTC()) {
		t.Error("inserted event dtend not equal", e.Dtend, dtend)
	}

	e, err = c.GetEvent(ctx, &client.GetEventParams{ID: e.ID})
	if err != nil {
		t.Errorf("could not get event: %v", err)
	}
	if e.ID == "" {
		t.Error("fetched event has no id")
	}
	if e.AccountID == "" {
		t.Error("fetched event has no account id")
	}
	if !reflect.DeepEqual(e.Dtstart.UTC(), dtstart.UTC()) {
		t.Error("fetched event dtstart not equal", e.Dtstart, dtstart)
	}
	if !reflect.DeepEqual(e.Dtend.UTC(), dtend.UTC()) {
		t.Error("fetched event dtend not equal", e.Dtend, dtend)
	}

	es, err := c.ListEvents(ctx, &client.ListEventsParams{Date: time.Now(), Period: "year"})
	if err != nil {
		t.Errorf("could not list events: %v", err)
	}
	if len(es) == 0 {
		t.Error("listed events for date range returned 0")
	}

	err = c.UpdateEvent(ctx, &client.UpdateEventParams{ID: e.ID, Props: map[string]any{"name": "x"}})
	if err != nil {
		t.Errorf("could not update event: %v", err)
	}

	es, err = c.ListEvents(ctx, &client.ListEventsParams{Props: map[string]any{"name": "x"}})
	if err != nil {
		t.Errorf("could not find events: %v", err)
	}
	if len(es) == 0 {
		t.Error("listed events for props returned 0")
	}

	e, err = c.GetEvent(ctx, &client.GetEventParams{ID: e.ID})
	if err != nil {
		t.Errorf("could not get event: %v", err)
	}
	if e.Props["name"] != "x" {
		t.Error("updated event props does not have name = x")
	}

	err = c.DeleteEvent(ctx, &client.DeleteEventParams{ID: e.ID})
	if err != nil {
		t.Errorf("could not delete event: %v", err)
	}
}
