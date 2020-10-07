package data

import (
	"math/rand"
	"reflect"
	"sort"

	"github.com/findy-network/findy-agent-api/graph/model"
)

type Items struct {
	items   []APIObject
	apiType string
}

func (i *Items) Append(object APIObject) {
	i.items = append(i.items, object)
}

func (i *Items) Count() int {
	return len(i.items)
}

func (i *Items) RandomID() string {
	max := len(i.items) - 1
	index := rand.Intn(max)
	return i.items[index].Identifier()
}

func (i *Items) FirstID() string {
	return i.items[0].Identifier()
}

func (i *Items) LastID() string {
	return i.items[len(i.items)-1].Identifier()
}

func (i *Items) CreatedForIndex(index int) int64 {
	return i.items[index].Created()
}

func (i *Items) MinCreated() int64 {
	return i.items[0].Created()
}

func (i *Items) MaxCreated() int64 {
	return i.items[len(i.items)-1].Created()
}

func (i *Items) Sort() {
	s := i.items
	sort.Slice(s, func(i, j int) bool {
		return s[i].Created() < s[i].Created()
	})
}

func (i *Items) PairwiseConnection(after, before int) *model.PairwiseConnection {
	result := i.items[after:before]
	totalCount := len(result)

	edges := make([]*model.PairwiseEdge, totalCount)
	nodes := make([]*model.Pairwise, totalCount)
	for index, pairwise := range result {
		node := pairwise.Pairwise().ToNode()
		edges[index] = &model.PairwiseEdge{
			Cursor: CreateCursor(pairwise.Pairwise().CreatedMs, model.Pairwise{}),
			Node:   node,
		}
		nodes[index] = node
	}

	var startCursor, endCursor *string
	if totalCount > 0 {
		startCursor = &edges[0].Cursor
		endCursor = &edges[totalCount-1].Cursor
	}
	p := &model.PairwiseConnection{
		Edges: edges,
		Nodes: nodes,
		PageInfo: &model.PageInfo{
			EndCursor:       endCursor,
			HasNextPage:     edges[len(edges)-1].Node.ID != i.LastID(),
			HasPreviousPage: edges[0].Node.ID != i.FirstID(),
			StartCursor:     startCursor,
		},
		TotalCount: totalCount,
	}
	return p
}

func (i *Items) EventConnection(after, before int) *model.EventConnection {
	result := i.items[after:before]
	totalCount := len(result)

	edges := make([]*model.EventEdge, totalCount)
	nodes := make([]*model.Event, totalCount)
	for index, event := range result {
		node := event.Event().ToNode()
		edges[index] = &model.EventEdge{
			Cursor: CreateCursor(event.Event().CreatedMs, model.Event{}),
			Node:   node,
		}
		nodes[index] = node
	}

	var startCursor, endCursor *string
	if totalCount > 0 {
		startCursor = &edges[0].Cursor
		endCursor = &edges[totalCount-1].Cursor
	}
	c := &model.EventConnection{
		Edges: edges,
		Nodes: nodes,
		PageInfo: &model.PageInfo{
			EndCursor:       endCursor,
			HasNextPage:     edges[len(edges)-1].Node.ID != i.LastID(),
			HasPreviousPage: edges[0].Node.ID != i.FirstID(),
			StartCursor:     startCursor,
		},
		TotalCount: totalCount,
	}
	return c
}

type Data struct {
	Connections Items
	Events      Items
}

var State *Data

func init() {
	State = &Data{
		Connections: Items{items: make([]APIObject, 0), apiType: reflect.TypeOf(model.Pairwise{}).Name()},
		Events:      Items{items: make([]APIObject, 0), apiType: reflect.TypeOf(model.Event{}).Name()},
	}
	for index := range Connections {
		State.Connections.items = append(State.Connections.items, &Connections[index])
	}
	State.Connections.Sort()

	for index := range Events {
		State.Events.items = append(State.Events.items, &Events[index])
	}
	State.Events.Sort()
}
