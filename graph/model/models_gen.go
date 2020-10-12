// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Event struct {
	ID          string       `json:"id"`
	Description string       `json:"description"`
	Protocol    ProtocolType `json:"protocol"`
	Type        EventType    `json:"type"`
	CreatedMs   string       `json:"createdMs"`
	Connection  *Pairwise    `json:"connection"`
}

type EventConnection struct {
	Edges      []*EventEdge `json:"edges"`
	Nodes      []*Event     `json:"nodes"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

type EventEdge struct {
	Cursor string `json:"cursor"`
	Node   *Event `json:"node"`
}

type Invitation struct {
	Payload string `json:"payload"`
}

type Offer struct {
	ID     string `json:"id"`
	Accept bool   `json:"accept"`
}

type PageInfo struct {
	EndCursor       *string `json:"endCursor"`
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
}

type Pairwise struct {
	ID            string `json:"id"`
	OurDid        string `json:"ourDid"`
	TheirDid      string `json:"theirDid"`
	TheirEndpoint string `json:"theirEndpoint"`
	TheirLabel    string `json:"theirLabel"`
	CreatedMs     string `json:"createdMs"`
	ApprovedMs    string `json:"approvedMs"`
	InitiatedByUs bool   `json:"initiatedByUs"`
}

type PairwiseConnection struct {
	Edges      []*PairwiseEdge `json:"edges"`
	Nodes      []*Pairwise     `json:"nodes"`
	PageInfo   *PageInfo       `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

type PairwiseEdge struct {
	Cursor string    `json:"cursor"`
	Node   *Pairwise `json:"node"`
}

type Request struct {
	ID     string `json:"id"`
	Accept bool   `json:"accept"`
}

type Response struct {
	Ok bool `json:"ok"`
}

type User struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	LastLoginTimeMs string `json:"lastLoginTimeMs"`
}

type EventType string

const (
	EventTypeNotification EventType = "NOTIFICATION"
	EventTypeAction       EventType = "ACTION"
)

var AllEventType = []EventType{
	EventTypeNotification,
	EventTypeAction,
}

func (e EventType) IsValid() bool {
	switch e {
	case EventTypeNotification, EventTypeAction:
		return true
	}
	return false
}

func (e EventType) String() string {
	return string(e)
}

func (e *EventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventType", str)
	}
	return nil
}

func (e EventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProtocolType string

const (
	ProtocolTypeNone         ProtocolType = "NONE"
	ProtocolTypeConnection   ProtocolType = "CONNECTION"
	ProtocolTypeCredential   ProtocolType = "CREDENTIAL"
	ProtocolTypeProof        ProtocolType = "PROOF"
	ProtocolTypeBasicMessage ProtocolType = "BASIC_MESSAGE"
)

var AllProtocolType = []ProtocolType{
	ProtocolTypeNone,
	ProtocolTypeConnection,
	ProtocolTypeCredential,
	ProtocolTypeProof,
	ProtocolTypeBasicMessage,
}

func (e ProtocolType) IsValid() bool {
	switch e {
	case ProtocolTypeNone, ProtocolTypeConnection, ProtocolTypeCredential, ProtocolTypeProof, ProtocolTypeBasicMessage:
		return true
	}
	return false
}

func (e ProtocolType) String() string {
	return string(e)
}

func (e *ProtocolType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProtocolType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProtocolType", str)
	}
	return nil
}

func (e ProtocolType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
