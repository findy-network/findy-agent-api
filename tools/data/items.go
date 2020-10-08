package data

import (
	"math/rand"
	"reflect"
	"sort"
	"sync"

	"github.com/findy-network/findy-agent-api/graph/model"
)

type Items struct {
	items   []APIObject
	apiType string
	mutex   sync.RWMutex
}

func (i *Items) Append(object APIObject) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.items = append(i.items, object)
}

func (i *Items) Count() (count int) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	count = len(i.items)
	return
}

func (i *Items) RandomID() (id string) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	max := len(i.items) - 1
	index := rand.Intn(max)
	id = i.items[index].Identifier()
	return
}

func (i *Items) FirstID() (id string) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	id = i.items[0].Identifier()
	return
}

func (i *Items) LastID() (id string) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	id = i.items[len(i.items)-1].Identifier()
	return
}

func (i *Items) CreatedForIndex(index int) (created int64) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	created = i.items[index].Created()
	return
}

func (i *Items) MinCreated() (created int64) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	created = i.items[0].Created()
	return
}

func (i *Items) MaxCreated() (created int64) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	created = i.items[len(i.items)-1].Created()
	return
}

func (i *Items) Sort() {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	s := i.items
	sort.Slice(s, func(i, j int) bool {
		return s[i].Created() < s[j].Created()
	})
}

func (i *Items) PairwiseForID(id string) *model.Pairwise {
	var node *model.Pairwise

	i.mutex.RLock()
	defer i.mutex.RUnlock()

	for _, item := range i.items {
		if item.Identifier() == id {
			node = item.Pairwise().ToNode()
			break
		}
	}

	return node
}

func (i *Items) PairwiseConnection(after, before int) *model.PairwiseConnection {
	i.mutex.RLock()
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
	i.mutex.RUnlock()

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
	i.mutex.RLock()
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
	i.mutex.RUnlock()

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
	Connections *Items
	Events      *Items
}

var State *Data

func init() {
	State = &Data{
		Connections: &Items{items: make([]APIObject, 0), apiType: reflect.TypeOf(model.Pairwise{}).Name()},
		Events:      &Items{items: make([]APIObject, 0), apiType: reflect.TypeOf(model.Event{}).Name()},
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
