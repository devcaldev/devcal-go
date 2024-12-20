# devcal-go

devcal go client is a thin wrapper for devcal grpc go client which handles type conversions between grpc and go internally.

## quick start

import devcal-go package

```go
import (
  "log"
  devcal "github.com/devcaldev/devcal-go"
)
```

initialize client

```go
client, err := devcal.New("devcal.dev:50051", apiKey)
if err != nil {
	log.Fatal(err)
}
defer client.Close()
```

insert non recurring event

```go
event, err := client.InsertEvent(ctx, &devcal.InsertEventParams{
	Dtstart: time.Now(),
	Dtend:   time.Now().Add(time.Hour),
})
if err != nil {
	log.Fatal(err)
}
log.Println("Event", event)
```

insert recurring event

```go
event, err := client.InsertEvent(ctx, &devcal.InsertEventParams{
	Dtstart: time.Now(),
	Dtend:   time.Now().Add(time.Hour),
	Rrule:   "FREQ=WEEKLY;INTERVAL=2;BYDAY=MO,SU",
})
if err != nil {
	log.Fatal(err)
}
log.Println("Event", event)
```

insert event with additional properties

```go
event, err := client.InsertEvent(ctx, &devcal.InsertEventParams{
	Dtstart: time.Now(),
	Dtend:   time.Now().Add(time.Hour),
	Rrule:   "FREQ=WEEKLY;INTERVAL=2;BYDAY=MO,SU",
  Props:   map[string]any{"calendar_id": "c1"},
})
if err != nil {
	log.Fatal(err)
}
log.Println("Event", event)
```

see `client_test.go` file and `example` folder for detailed usage.

goto [devcal.dev](https://devcal.dev) to get an `api key`
