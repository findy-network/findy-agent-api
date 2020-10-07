package resolver

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/findy-network/findy-agent-api/graph/model"
	"github.com/findy-network/findy-agent-api/tools/data"
	"github.com/findy-network/findy-agent-api/tools/faker"
	"github.com/lainio/err2"
)

var eventAddedObserver map[string]chan *model.EventEdge

func initEvents() {
	eventAddedObserver = map[string]chan *model.EventEdge{}

}

func (r *queryResolver) Events(
	_ context.Context,
	after *string, before *string,
	first *int, last *int) (c *model.EventConnection, err error) {
	defer err2.Return(&err)

	afterIndex, beforeIndex, err := pick(data.State.Events, after, before, first, last)
	err2.Check(err)

	return data.State.Events.EventConnection(afterIndex, beforeIndex), nil
}

func (r *subscriptionResolver) EventAdded(ctx context.Context) (<-chan *model.EventEdge, error) {
	id := "tenantId-" + strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println("Add id", id)
	events := make(chan *model.EventEdge, 1)

	go func() {
		<-ctx.Done()
		fmt.Println("Delete id", id)
		delete(eventAddedObserver, id)
	}()

	eventAddedObserver[id] = events

	return events, nil
}

func (r *mutationResolver) AddRandomEvent(ctx context.Context) (bool, error) {
	events, err := faker.FakeEvents(1, data.State.Connections)
	if err == nil {
		data.State.Events.Append(&events[0])
		fmt.Println("Added event", events[0].ID)
		for _, observer := range eventAddedObserver {
			observer <- events[0].ToEdge()
		}
	}
	return true, err
}
