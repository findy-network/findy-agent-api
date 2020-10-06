package tools

import (
	"strconv"

	"github.com/findy-network/findy-agent-api/graph/model"
)

type InternalPairwise struct {
	ID            string `faker:"uuid_hyphenated"`
	OurDid        string
	TheirDid      string
	TheirEndpoint string `faker:"url"`
	TheirLabel    string `faker:"organisationLabel"`
	InitiatedByUs bool
	ApprovedMs    int64 `faker:"unix_time"`
	CreatedMs     int64 `faker:"unix_time"`
}

func (p *InternalPairwise) toNode() *model.Pairwise {
	return &model.Pairwise{
		ID:            p.ID,
		OurDid:        p.OurDid,
		TheirDid:      p.TheirDid,
		TheirEndpoint: p.TheirEndpoint,
		TheirLabel:    p.TheirLabel,
		CreatedMs:     strconv.FormatInt(p.CreatedMs, 10),
		ApprovedMs:    strconv.FormatInt(p.ApprovedMs, 10),
		InitiatedByUs: p.InitiatedByUs,
	}
}

type InternalEvent struct {
	ID           string             `faker:"uuid_hyphenated"`
	Description  string             `faker:"sentence"`
	ProtocolType model.ProtocolType `faker:"oneof: model.ProtocolTypeNone, model.ProtocolTypeConnection, model.ProtocolTypeCredential, model.ProtocolTypeProof, model.ProtocolTypeBasicMessage"`
	Type         model.EventType    `faker:"oneof: model.EventTypeNotification, model.EventTypeAction"`
	PairwiseID   string             `faker:"eventPairwiseId"`
	CreatedMs    int64              `faker:"unix_time"`
}

func (e *InternalEvent) toEdge() *model.EventEdge {
	cursor := CreateCursor(e.CreatedMs, model.Event{})
	return &model.EventEdge{
		Cursor: cursor,
		Node:   e.toNode(),
	}

}

func (e *InternalEvent) toNode() *model.Event {
	createdStr := strconv.FormatInt(e.CreatedMs, 10)
	var conn *InternalPairwise
	for index, c := range Connections {
		if c.ID == e.PairwiseID {
			conn = &Connections[index]
		}
	}
	return &model.Event{
		ID:          e.ID,
		Description: e.Description,
		Protocol:    e.ProtocolType,
		Type:        e.Type,
		CreatedMs:   createdStr,
		Connection:  conn.toNode(),
	}

}
