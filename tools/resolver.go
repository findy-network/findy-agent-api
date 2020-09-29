package tools

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/findy-network/findy-agent-api/graph/generated"
	"github.com/findy-network/findy-agent-api/graph/model"
	"github.com/findy-network/findy-agent-api/resolver"
)

var Connections = []InternalPairwise{
	{"4b7565eb-062b-4286-9115-c0584fa486bf", "wHELJGmdZnWZKSttXfrTlNadR", "gvZZqQTEsyijwXEBaLyHyKKfi", "http://www.BvYYcKn.com/", "Ms. Vivian Dibbert", true, 190080541, 274877981},
	{"2de0c34e-3d97-4cba-95a6-99d2f675e2b7", "iGlpntctWWocPKVMNkeYNsRbZ", "DWXRRNVCDcSEDVrcNtDvTJsje", "http://nUrkwXD.com/QiOfbAg.html", "Miss Angie Volkman", true, 972509823, 1225650817},
	{"01fbf139-9ef6-44b5-a8ed-355d737442d7", "GCdvuVODIqrmnjLwtpYZueqnp", "qirbjDmmNwuHVYebuEswnGItS", "https://www.JCawACK.com/VALjSMm", "Prof. Name Satterfield", false, 1370099585, 895722201},
	{"96458265-5e4d-462a-a107-04ac606a8c79", "eHWscqOMorqcRitXTooXxkglU", "SjlQjYwLbkflyLljskJfGjnZR", "https://www.mhAyHBc.biz/", "Miss Breana Goodwin", false, 1403296283, 1508512828},
	{"27f18a75-5ca2-42c6-b509-08c5fe07a65d", "vCrqgVmpbyltcKcFAeJGnpIhh", "euDSUETPKeZQmTunecrAuiyWU", "https://www.BdFpLDu.org/qhnQhcS", "Princess Patricia Gleason", true, 1545036772, 328612424},
}

type ByCreated []InternalPairwise

func (a ByCreated) Len() int           { return len(a) }
func (a ByCreated) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCreated) Less(i, j int) bool { return a[i].CreatedMs < a[j].CreatedMs }

func CreateCursor(created int64, t reflect.Type) string {
	return base64.StdEncoding.EncodeToString([]byte(t.Name() + ":" + strconv.FormatInt(created, 10)))
}

func parseCursor(cursor string, t reflect.Type) (int64, error) {
	plain, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, errors.New(resolver.ErrorCursorInvalid)
	}

	parts := strings.Split(string(plain), ":")
	if len(parts) != 2 {
		return 0, errors.New(resolver.ErrorCursorInvalid)
	}

	value, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, errors.New(resolver.ErrorCursorInvalid)
	}

	return value, nil
}

type Resolver struct{}

func (r *mutationResolver) Invite(ctx context.Context) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Connect(ctx context.Context, input model.Invitation) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendMessage(ctx context.Context) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AcceptOffer(ctx context.Context, input model.Offer) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AcceptRequest(ctx context.Context, input model.Request) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

/*

if first, last missing, return error

Start from the greedy query: SELECT * FROM table ORDER BY created
If the after argument is provided, add id > parsed_cursor to the WHERE clause
If the before argument is provided, add id < parsed_cursor to the WHERE clause
If the first argument is provided, add ORDER BY id DESC LIMIT first+1 to the query
If the last argument is provided, add ORDER BY id ASC LIMIT last+1 to the query
If the last argument is provided, I reverse the order of the results
If the first argument is provided then I set hasPreviousPage: false (see spec for a description of this behavior).
If no less than first+1 results are returned, I set hasNextPage: true, otherwise I set it to false.
If the last argument is provided then I set hasNextPage: false (see spec for a description of this behavior).
If no less last+1 results are returned, I set hasPreviousPage: true, otherwise I set it to false.
*/

func (r *queryResolver) Connections(ctx context.Context, after *string, before *string, first *int, last *int) (*model.PairwiseConnection, error) {
	if first == nil && last == nil {
		return nil, errors.New(resolver.ErrorFirstLastMissing)
	}
	if (first != nil && (*first < 1 || *first > 100)) || (last != nil && (*last < 1 || *last > 100)) {
		return nil, errors.New(resolver.ErrorFirstLastInvalid)
	}

	sort.Sort(ByCreated(Connections))
	afterIndex := 0
	beforeIndex := len(Connections) - 1
	if after != nil || before != nil {
		var afterVal int64
		var beforeVal int64
		var err error
		if after != nil {
			afterVal, err = parseCursor(*after, reflect.TypeOf(model.Pairwise{}))
			if err != nil {
				return nil, err
			}
		}
		if before != nil {
			beforeVal, err = parseCursor(*before, reflect.TypeOf(model.Pairwise{}))
			if err != nil {
				return nil, err
			}
		}
		for index, value := range Connections {
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

	//hasPreviousPage := first == nil
	//hasNextPage := last == nil
	if first != nil {
		afterPlusFirst := afterIndex + (*first - 1)
		if beforeIndex > afterPlusFirst {
			beforeIndex = afterPlusFirst
			//hasNextPage = true
		}
	} else if last != nil {
		beforeMinusLast := beforeIndex - (*last - 1)
		if afterIndex < beforeMinusLast {
			afterIndex = beforeMinusLast
			//hasPreviousPage = true
		}
	}
	result := Connections[afterIndex:(beforeIndex + 1)]
	totalCount := len(result)
	nodes := make([]*model.Pairwise, totalCount)
	for index, pairwise := range result {
		nodes[index] = &model.Pairwise{
			ID:            pairwise.ID,
			OurDid:        pairwise.OurDid,
			TheirDid:      pairwise.TheirDid,
			TheirEndpoint: pairwise.TheirEndpoint,
			TheirLabel:    pairwise.TheirLabel,
			CreatedMs:     strconv.FormatInt(pairwise.CreatedMs, 10),
			ApprovedMs:    strconv.FormatInt(pairwise.ApprovedMs, 10),
			InitiatedByUs: pairwise.InitiatedByUs,
		}
	}

	edges := make([]*model.PairwiseEdge, totalCount)
	for index, pairwise := range nodes {
		edges[index] = &model.PairwiseEdge{
			Cursor: CreateCursor(result[index].CreatedMs, reflect.TypeOf(model.Pairwise{})),
			Node:   pairwise,
		}
	}

	var startCursor *string
	var endCursor *string
	if totalCount > 0 {
		startCursor = &edges[0].Cursor
		endCursor = &edges[totalCount-1].Cursor
	}
	p := &model.PairwiseConnection{
		Edges: edges,
		Nodes: nodes,
		PageInfo: &model.PageInfo{
			EndCursor:       endCursor,
			HasNextPage:     edges[len(edges)-1].Node.ID != Connections[len(Connections)-1].ID,
			HasPreviousPage: edges[0].Node.ID != Connections[0].ID,
			StartCursor:     startCursor,
		},
		TotalCount: totalCount,
	}
	return p, nil
}

func (r *queryResolver) Connection(ctx context.Context, id string) (*model.Pairwise, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Events(ctx context.Context, after *string, before *string, first *int, last *int) (*model.EventConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Event(ctx context.Context, id string) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) EventAdded(ctx context.Context, eventID string) (<-chan *model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
