package resolver

import (
	"context"
	"errors"
	"reflect"
	"sort"
	"strconv"

	"github.com/findy-network/findy-agent-api/graph/model"
	"github.com/findy-network/findy-agent-api/resolver"
	"github.com/findy-network/findy-agent-api/tools/data"
)

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

func (r *queryResolver) Connections(_ context.Context, after *string, before *string, first *int, last *int) (*model.PairwiseConnection, error) {
	if first == nil && last == nil {
		return nil, errors.New(resolver.ErrorFirstLastMissing)
	}
	if (first != nil && (*first < 1 || *first > 100)) || (last != nil && (*last < 1 || *last > 100)) {
		return nil, errors.New(resolver.ErrorFirstLastInvalid)
	}

	sort.Slice(data.Connections, func(i, j int) bool {
		return data.Connections[i].CreatedMs < data.Connections[j].CreatedMs
	})

	afterIndex := 0
	beforeIndex := len(data.Connections) - 1
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
		for index, value := range data.Connections {
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
	result := data.Connections[afterIndex:(beforeIndex + 1)]
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
			Cursor: data.CreateCursor(result[index].CreatedMs, model.Pairwise{}),
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
			HasNextPage:     edges[len(edges)-1].Node.ID != data.Connections[len(data.Connections)-1].ID,
			HasPreviousPage: edges[0].Node.ID != data.Connections[0].ID,
			StartCursor:     startCursor,
		},
		TotalCount: totalCount,
	}
	return p, nil
}
