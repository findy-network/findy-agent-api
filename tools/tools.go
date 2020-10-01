package tools

import "github.com/findy-network/findy-agent-api/graph/model"

type InternalPairwise struct {
	ID            string `faker:"uuid_hyphenated"`
	OurDid        string
	TheirDid      string
	TheirEndpoint string `faker:"url"`
	TheirLabel    string `faker:"name"`
	InitiatedByUs bool
	//CreatedMs     int64 `faker:"unix_time"`
	ApprovedMs int64 `faker:"unix_time"`
	CreatedMs  int64 `faker:"unix_time"`
}

type InternalEvent struct {
	ID           string             `faker:"uuid_hyphenated"`
	Description  string             `faker:"sentence"`
	ProtocolType model.ProtocolType `faker:"oneof: model.ProtocolTypeNone, model.ProtocolTypeConnection, model.ProtocolTypeCredential, model.ProtocolTypeProof, model.ProtocolTypeBasicMessage"`
	Type         model.EventType    `faker:"oneof: model.EventTypeNotification, model.EventTypeAction"`
	PairwiseID   string             `faker:"eventPairwiseId"`
	CreatedMs    int64              `faker:"unix_time"`
}
