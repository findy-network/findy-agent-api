// Copyright 2020 Harri @ OP Techlab.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: agent.proto

package agency

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type is enum type to tell what happening
type Notification_Type int32

const (
	Notification_STATUS_UPDATE               Notification_Type = 0 // General status update where no action needed
	Notification_ACTION_NEEDED               Notification_Type = 1 // General action needed update notification
	Notification_ANSWER_NEEDED_PING          Notification_Type = 2 // Your CA controller has been pinged
	Notification_ANSWER_NEEDED_ISSUE_PROPOSE Notification_Type = 3 // Issuing is proposed
	Notification_ANSWER_NEEDED_PROOF_PROPOSE Notification_Type = 4 // Proof is proposed
	Notification_ANSWER_NEEDED_PROOF_VERIFY  Notification_Type = 5 // During proof values need to be verified
)

// Enum value maps for Notification_Type.
var (
	Notification_Type_name = map[int32]string{
		0: "STATUS_UPDATE",
		1: "ACTION_NEEDED",
		2: "ANSWER_NEEDED_PING",
		3: "ANSWER_NEEDED_ISSUE_PROPOSE",
		4: "ANSWER_NEEDED_PROOF_PROPOSE",
		5: "ANSWER_NEEDED_PROOF_VERIFY",
	}
	Notification_Type_value = map[string]int32{
		"STATUS_UPDATE":               0,
		"ACTION_NEEDED":               1,
		"ANSWER_NEEDED_PING":          2,
		"ANSWER_NEEDED_ISSUE_PROPOSE": 3,
		"ANSWER_NEEDED_PROOF_PROPOSE": 4,
		"ANSWER_NEEDED_PROOF_VERIFY":  5,
	}
)

func (x Notification_Type) Enum() *Notification_Type {
	p := new(Notification_Type)
	*p = x
	return p
}

func (x Notification_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Notification_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_agent_proto_enumTypes[0].Descriptor()
}

func (Notification_Type) Type() protoreflect.EnumType {
	return &file_agent_proto_enumTypes[0]
}

func (x Notification_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Notification_Type.Descriptor instead.
func (Notification_Type) EnumDescriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{3, 0}
}

//
//Answer is a message send by Give function of Agent service.
type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                             // Same as Notification ID (UUID)
	ClientId *ClientID `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // Same as your ClientID when Listening was started
	Ack      bool      `protobuf:"varint,3,opt,name=ack,proto3" json:"ack,omitempty"`                          // Your response to the protocol question
	Info     string    `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`                         // General info, mostly used for debugging
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{0}
}

func (x *Answer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Answer) GetClientId() *ClientID {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *Answer) GetAck() bool {
	if x != nil {
		return x.Ack
	}
	return false
}

func (x *Answer) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

// ClientID is UUID. If user has many different client device connected to
// cloud agent it must identify who is talking to.
type ClientID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // UUID of the client
}

func (x *ClientID) Reset() {
	*x = ClientID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientID) ProtoMessage() {}

func (x *ClientID) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientID.ProtoReflect.Descriptor instead.
func (*ClientID) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{1}
}

func (x *ClientID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

//
//AgentStatus is a message identifying current agent events returned as
//notifications.
type AgentStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     *ClientID     `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // UUID of the client listening
	Notification *Notification `protobuf:"bytes,3,opt,name=notification,proto3" json:"notification,omitempty"`         // The actual Notification message
}

func (x *AgentStatus) Reset() {
	*x = AgentStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentStatus) ProtoMessage() {}

func (x *AgentStatus) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentStatus.ProtoReflect.Descriptor instead.
func (*AgentStatus) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{2}
}

func (x *AgentStatus) GetClientId() *ClientID {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *AgentStatus) GetNotification() *Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

//
//Notification is a message used to tell meaningful events outside from cloud
//agent.
type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId         Notification_Type `protobuf:"varint,1,opt,name=type_id,json=typeId,proto3,enum=agency.Notification_Type" json:"type_id,omitempty"` // Notification type, see Type
	Id             string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`                                                      // Notification's unique ID
	ConnectionId   string            `protobuf:"bytes,3,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`              // Current pairwise ID between agents
	ProtocolId     string            `protobuf:"bytes,4,opt,name=protocol_id,json=protocolId,proto3" json:"protocol_id,omitempty"`                    // Current protocol ID, see Aries Thread ID
	ProtocolFamily string            `protobuf:"bytes,5,opt,name=protocol_family,json=protocolFamily,proto3" json:"protocol_family,omitempty"`        // Text version of the protocol family/namespace
	Timestamp      uint64            `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                       // timestamp in ms
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{3}
}

func (x *Notification) GetTypeId() Notification_Type {
	if x != nil {
		return x.TypeId
	}
	return Notification_STATUS_UPDATE
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *Notification) GetProtocolId() string {
	if x != nil {
		return x.ProtocolId
	}
	return ""
}

func (x *Notification) GetProtocolFamily() string {
	if x != nil {
		return x.ProtocolFamily
	}
	return ""
}

func (x *Notification) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x6d, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2d, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x63, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x76, 0x0a, 0x0b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2d, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x38,
	0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x88, 0x03, 0x0a, 0x0c, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x07, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xa6, 0x01, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4e, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x4e, 0x53,
	0x57, 0x45, 0x52, 0x5f, 0x4e, 0x45, 0x45, 0x44, 0x45, 0x44, 0x5f, 0x50, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x5f, 0x4e, 0x45, 0x45, 0x44,
	0x45, 0x44, 0x5f, 0x49, 0x53, 0x53, 0x55, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x45,
	0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x5f, 0x4e, 0x45, 0x45,
	0x44, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x4f, 0x46, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53,
	0x45, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x5f, 0x4e, 0x45,
	0x45, 0x44, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x4f, 0x46, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46,
	0x59, 0x10, 0x05, 0x32, 0x68, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x06,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x10, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x13, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x2a, 0x0a, 0x04, 0x47, 0x69, 0x76, 0x65, 0x12, 0x0e, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x00, 0x42, 0x36, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6e, 0x64,
	0x79, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x79, 0x2d,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_proto_rawDescOnce sync.Once
	file_agent_proto_rawDescData = file_agent_proto_rawDesc
)

func file_agent_proto_rawDescGZIP() []byte {
	file_agent_proto_rawDescOnce.Do(func() {
		file_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_proto_rawDescData)
	})
	return file_agent_proto_rawDescData
}

var file_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_agent_proto_goTypes = []interface{}{
	(Notification_Type)(0), // 0: agency.Notification.Type
	(*Answer)(nil),         // 1: agency.Answer
	(*ClientID)(nil),       // 2: agency.ClientID
	(*AgentStatus)(nil),    // 3: agency.AgentStatus
	(*Notification)(nil),   // 4: agency.Notification
}
var file_agent_proto_depIdxs = []int32{
	2, // 0: agency.Answer.client_id:type_name -> agency.ClientID
	2, // 1: agency.AgentStatus.client_id:type_name -> agency.ClientID
	4, // 2: agency.AgentStatus.notification:type_name -> agency.Notification
	0, // 3: agency.Notification.type_id:type_name -> agency.Notification.Type
	2, // 4: agency.Agent.Listen:input_type -> agency.ClientID
	1, // 5: agency.Agent.Give:input_type -> agency.Answer
	3, // 6: agency.Agent.Listen:output_type -> agency.AgentStatus
	2, // 7: agency.Agent.Give:output_type -> agency.ClientID
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_agent_proto_init() }
func file_agent_proto_init() {
	if File_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agent_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_proto_goTypes,
		DependencyIndexes: file_agent_proto_depIdxs,
		EnumInfos:         file_agent_proto_enumTypes,
		MessageInfos:      file_agent_proto_msgTypes,
	}.Build()
	File_agent_proto = out.File
	file_agent_proto_rawDesc = nil
	file_agent_proto_goTypes = nil
	file_agent_proto_depIdxs = nil
}
