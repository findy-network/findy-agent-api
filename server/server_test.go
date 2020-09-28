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

	Handler(&Resolver{}).ServeHTTP(response, request)

	bytes := response.Body.Bytes()
	fmt.Println(string(bytes))
	json.Unmarshal(bytes, &payload)
	return
}

func connQuery(arguments string) string {
	if len(arguments) > 0 {
		arguments = "(" + arguments + ")"
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
		data := connections
		paginationInvalidError := JSON{Errors: &[]JSONError{
			{
				Message: ErrorFirstLastInvalid,
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
					Message: ErrorFirstLastMissing,
					Path:    errorPath,
				},
			}}},
			{"connections, pagination first too low", args{connQuery("first: 0")}, paginationInvalidError},
			{"connections, pagination first too high", args{connQuery("first: 101")}, paginationInvalidError},
			{"connections, pagination last too low", args{connQuery("last: 0")}, paginationInvalidError},
			{"connections, pagination last too high", args{connQuery("last: 101")}, paginationInvalidError},
			{"first connection ", args{connQuery("first: 1")}, JSON{Data: &JSONData{Connections: model.PairwiseConnection{
				Edges: []*model.PairwiseEdge{
					{Cursor: createCursor(data[0].CreatedMs, reflect.TypeOf(model.Pairwise{})), Node: &model.Pairwise{
						ID:            data[0].ID,
						OurDid:        data[0].OurDid,
						TheirDid:      data[0].TheirDid,
						TheirEndpoint: data[0].TheirEndpoint,
						TheirLabel:    data[0].TheirLabel,
						CreatedMs:     strconv.FormatInt(data[0].CreatedMs, 10),
						ApprovedMs:    strconv.FormatInt(data[0].ApprovedMs, 10),
						InitiatedByUs: data[0].InitiatedByUs,
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
