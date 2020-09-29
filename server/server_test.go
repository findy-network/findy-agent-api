package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/findy-network/findy-agent-api/graph/model"
	"github.com/findy-network/findy-agent-api/resolver"
	"github.com/findy-network/findy-agent-api/tools"
)

type JSONError struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}

type JSONData struct {
	Connections model.PairwiseConnection `json:"connections"`
}

type JSON struct {
	Data   *JSONData    `json:"data"`
	Errors *[]JSONError `json:"errors"`
}

func queryJSON(content string) string {
	content = strings.Replace(content, "\t", "", -1)
	content = strings.Replace(content, "\n", " ", -1)
	return `{
		"query": "` + content + `"
		}`
}

func doQuery(query string) (payload JSON) {
	request, _ := http.NewRequest(http.MethodPost, "/query", strings.NewReader(queryJSON(query)))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	Handler(&tools.Resolver{}).ServeHTTP(response, request)

	bytes := response.Body.Bytes()
	fmt.Println(string(bytes))
	json.Unmarshal(bytes, &payload)
	return
}

func connQuery(arguments string) string {
	if len(arguments) > 0 {
		arguments = "(" + strings.Replace(arguments, "\"", "\\\"", -1) + ")"
	}
	return `{
		connections` + arguments + ` {
			edges {
				cursor
				node {
					id
					ourDid
					theirDid
					theirEndpoint
					theirLabel
					createdMs
					approvedMs
					initiatedByUs
				}
			}
		}
	}`
}

func TestGetConnections(t *testing.T) {
	t.Run("get connections", func(t *testing.T) {
		type args struct {
			query string
		}
		errorPath := []string{"connections"}
		first := tools.Connections[0]
		firstCursor := tools.CreateCursor(first.CreatedMs, reflect.TypeOf(model.Pairwise{}))
		second := tools.Connections[1]
		secondCursor := tools.CreateCursor(second.CreatedMs, reflect.TypeOf(model.Pairwise{}))
		last := tools.Connections[len(tools.Connections)-1]
		paginationInvalidError := JSON{Errors: &[]JSONError{
			{
				Message: resolver.ErrorFirstLastInvalid,
				Path:    errorPath,
			},
		}}
		tests := []struct {
			name   string
			args   args
			result JSON
		}{
			{"connections, pagination missing", args{connQuery("")}, JSON{Errors: &[]JSONError{
				{
					Message: resolver.ErrorFirstLastMissing,
					Path:    errorPath,
				},
			}}},
			{"connections, pagination first too low", args{connQuery("first: 0")}, paginationInvalidError},
			{"connections, pagination first too high", args{connQuery("first: 101")}, paginationInvalidError},
			{"connections, pagination last too low", args{connQuery("last: 0")}, paginationInvalidError},
			{"connections, pagination last too high", args{connQuery("last: 101")}, paginationInvalidError},
			{"connections, after cursor invalid", args{connQuery("first: 1, after: \"1\"")}, JSON{Errors: &[]JSONError{
				{
					Message: resolver.ErrorCursorInvalid,
					Path:    errorPath,
				},
			}}},
			{"first connection ", args{connQuery("first: 1")}, JSON{Data: &JSONData{Connections: model.PairwiseConnection{
				Edges: []*model.PairwiseEdge{
					{Cursor: firstCursor, Node: &model.Pairwise{
						ID:            first.ID,
						OurDid:        first.OurDid,
						TheirDid:      first.TheirDid,
						TheirEndpoint: first.TheirEndpoint,
						TheirLabel:    first.TheirLabel,
						CreatedMs:     strconv.FormatInt(first.CreatedMs, 10),
						ApprovedMs:    strconv.FormatInt(first.ApprovedMs, 10),
						InitiatedByUs: first.InitiatedByUs,
					}},
				},
			}}}},
			{"last connection ", args{connQuery("last: 1")}, JSON{Data: &JSONData{Connections: model.PairwiseConnection{
				Edges: []*model.PairwiseEdge{
					{Cursor: tools.CreateCursor(last.CreatedMs, reflect.TypeOf(model.Pairwise{})), Node: &model.Pairwise{
						ID:            last.ID,
						OurDid:        last.OurDid,
						TheirDid:      last.TheirDid,
						TheirEndpoint: last.TheirEndpoint,
						TheirLabel:    last.TheirLabel,
						CreatedMs:     strconv.FormatInt(last.CreatedMs, 10),
						ApprovedMs:    strconv.FormatInt(last.ApprovedMs, 10),
						InitiatedByUs: last.InitiatedByUs,
					}},
				},
			}}}},
			{"second connection ", args{connQuery("first: 1, after: \"" + firstCursor + "\"")}, JSON{Data: &JSONData{Connections: model.PairwiseConnection{
				Edges: []*model.PairwiseEdge{
					{Cursor: secondCursor, Node: &model.Pairwise{
						ID:            second.ID,
						OurDid:        second.OurDid,
						TheirDid:      second.TheirDid,
						TheirEndpoint: second.TheirEndpoint,
						TheirLabel:    second.TheirLabel,
						CreatedMs:     strconv.FormatInt(second.CreatedMs, 10),
						ApprovedMs:    strconv.FormatInt(second.ApprovedMs, 10),
						InitiatedByUs: second.InitiatedByUs,
					}},
				},
			}}}},
			{"previous to second connection ", args{connQuery("first: 1, before: \"" + secondCursor + "\"")}, JSON{Data: &JSONData{Connections: model.PairwiseConnection{
				Edges: []*model.PairwiseEdge{
					{Cursor: firstCursor, Node: &model.Pairwise{
						ID:            first.ID,
						OurDid:        first.OurDid,
						TheirDid:      first.TheirDid,
						TheirEndpoint: first.TheirEndpoint,
						TheirLabel:    first.TheirLabel,
						CreatedMs:     strconv.FormatInt(first.CreatedMs, 10),
						ApprovedMs:    strconv.FormatInt(first.ApprovedMs, 10),
						InitiatedByUs: first.InitiatedByUs,
					}},
				},
			}}}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := doQuery(tt.args.query); !reflect.DeepEqual(got, tt.result) {
					t.Errorf("%s = %v, want %v", tt.name, got.Data, tt.result.Data)
				}
			})
		}

	})
}
