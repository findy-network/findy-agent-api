// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package agency

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	// Listen is async function to stream AgentStatus. ClientID must be unique.
	Listen(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (Agent_ListenClient, error)
	// Give is function to give answer to ACTION_NEEDED_xx notifications.
	Give(ctx context.Context, in *Answer, opts ...grpc.CallOption) (*ClientID, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Listen(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (Agent_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Agent_serviceDesc.Streams[0], "/agency.Agent/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Agent_ListenClient interface {
	Recv() (*AgentStatus, error)
	grpc.ClientStream
}

type agentListenClient struct {
	grpc.ClientStream
}

func (x *agentListenClient) Recv() (*AgentStatus, error) {
	m := new(AgentStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) Give(ctx context.Context, in *Answer, opts ...grpc.CallOption) (*ClientID, error) {
	out := new(ClientID)
	err := c.cc.Invoke(ctx, "/agency.Agent/Give", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	// Listen is async function to stream AgentStatus. ClientID must be unique.
	Listen(*ClientID, Agent_ListenServer) error
	// Give is function to give answer to ACTION_NEEDED_xx notifications.
	Give(context.Context, *Answer) (*ClientID, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) Listen(*ClientID, Agent_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedAgentServer) Give(context.Context, *Answer) (*ClientID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Give not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClientID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServer).Listen(m, &agentListenServer{stream})
}

type Agent_ListenServer interface {
	Send(*AgentStatus) error
	grpc.ServerStream
}

type agentListenServer struct {
	grpc.ServerStream
}

func (x *agentListenServer) Send(m *AgentStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Agent_Give_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Answer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Give(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agency.Agent/Give",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Give(ctx, req.(*Answer))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agency.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Give",
			Handler:    _Agent_Give_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _Agent_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "agent.proto",
}