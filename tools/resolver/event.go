package resolver

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"time"

	"github.com/findy-network/findy-agent-api/tools/faker"

	"github.com/findy-network/findy-agent-api/graph/model"
	"github.com/findy-network/findy-agent-api/resolver"
	"github.com/findy-network/findy-agent-api/tools/data"
)

var eventAddedObserver map[string]chan *model.EventEdge

func initEvents() {
	eventAddedObserver = map[string]chan *model.EventEdge{}

}

func (r *queryResolver) Events(_ context.Context, after *string, before *string, first *int, last *int) (*model.EventConnection, error) {
	fmt.Println("after", after)
	fmt.Println("before", before)
	if first == nil && last == nil {
		return nil, errors.New(resolver.ErrorFirstLastMissing)
	}
	if (first != nil && (*first < 1 || *first > 100)) || (last != nil && (*last < 1 || *last > 100)) {
		return nil, errors.New(resolver.ErrorFirstLastInvalid)
	}

	sort.Slice(data.Events, func(i, j int) bool {
		return data.Events[i].CreatedMs < data.Events[j].CreatedMs
	})

	afterIndex := 0
	beforeIndex := len(data.Events) - 1
	if after != nil || before != nil {
		var afterVal int64
		var beforeVal int64
		var err error
		if after != nil {
			afterVal, err = parseCursor(*after, reflect.TypeOf(model.Event{}))
			if err != nil {
				return nil, err
			}
		}
		if before != nil {
			beforeVal, err = parseCursor(*before, reflect.TypeOf(model.Event{}))
			if err != nil {
				return nil, err
			}
		}
		for index, value := range data.Events {
			if afterVal > 0 && value.CreatedMs <= afterVal {
				afterIndex = index + 1
			}
			if beforeVal > 0 && value.CreatedMs < beforeVal {
				beforeIndex = index
			}
			if (beforeVal > 0 && value.CreatedMs > beforeVal) || (beforeVal == 0 && value.CreatedMs > afterVal) {
				break
			}
		}
	}

	if first != nil {
		afterPlusFirst := afterIndex + (*first - 1)
		if beforeIndex > afterPlusFirst {
			beforeIndex = afterPlusFirst
		}
	} else if last != nil {
		beforeMinusLast := beforeIndex - (*last - 1)
		if afterIndex < beforeMinusLast {
			afterIndex = beforeMinusLast
		}
	}
	result := data.Events[afterIndex:(beforeIndex + 1)]
	totalCount := len(result)
	nodes := make([]*model.Event, totalCount)

	edges := make([]*model.EventEdge, totalCount)
	for index, event := range result {
		edges[index] = event.ToEdge()
	}

	for index, edge := range edges {
		nodes[index] = edge.Node
	}

	var startCursor *string
	var endCursor *string
	if totalCount > 0 {
		startCursor = &edges[0].Cursor
		endCursor = &edges[totalCount-1].Cursor
	}
	p := &model.EventConnection{
		Edges: edges,
		Nodes: nodes,
		PageInfo: &model.PageInfo{
			EndCursor:       endCursor,
			HasNextPage:     edges[len(edges)-1].Node.ID != data.Events[len(data.Events)-1].ID,
			HasPreviousPage: edges[0].Node.ID != data.Events[0].ID,
			StartCursor:     startCursor,
		},
		TotalCount: totalCount,
	}
	return p, nil
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
	events, err := faker.FakeEvents(1, data.Connections)
	if err == nil {
		data.Events = append(data.Events, events...)
		fmt.Println("Added event", events[0].ID)
		for _, observer := range eventAddedObserver {
			observer <- data.Events[len(data.Events)-1].ToEdge()
		}
	}
	return true, err
}
