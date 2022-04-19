// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb_chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	pb_ws "suzaku/pkg/proto/pb_ws"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	GetMaxAndMinSeq(ctx context.Context, in *GetMaxAndMinSeqReq, opts ...grpc.CallOption) (*GetMaxAndMinSeqResp, error)
	PullMessageBySeqList(ctx context.Context, in *pb_ws.PullMessageBySeqListReq, opts ...grpc.CallOption) (*pb_ws.PullMessageBySeqListResp, error)
	SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) GetMaxAndMinSeq(ctx context.Context, in *GetMaxAndMinSeqReq, opts ...grpc.CallOption) (*GetMaxAndMinSeqResp, error) {
	out := new(GetMaxAndMinSeqResp)
	err := c.cc.Invoke(ctx, "/pb_auth.Chat/GetMaxAndMinSeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) PullMessageBySeqList(ctx context.Context, in *pb_ws.PullMessageBySeqListReq, opts ...grpc.CallOption) (*pb_ws.PullMessageBySeqListResp, error) {
	out := new(pb_ws.PullMessageBySeqListResp)
	err := c.cc.Invoke(ctx, "/pb_auth.Chat/PullMessageBySeqList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	out := new(SendMsgResp)
	err := c.cc.Invoke(ctx, "/pb_auth.Chat/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	GetMaxAndMinSeq(context.Context, *GetMaxAndMinSeqReq) (*GetMaxAndMinSeqResp, error)
	PullMessageBySeqList(context.Context, *pb_ws.PullMessageBySeqListReq) (*pb_ws.PullMessageBySeqListResp, error)
	SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) GetMaxAndMinSeq(context.Context, *GetMaxAndMinSeqReq) (*GetMaxAndMinSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMaxAndMinSeq not implemented")
}
func (UnimplementedChatServer) PullMessageBySeqList(context.Context, *pb_ws.PullMessageBySeqListReq) (*pb_ws.PullMessageBySeqListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullMessageBySeqList not implemented")
}
func (UnimplementedChatServer) SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_GetMaxAndMinSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMaxAndMinSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetMaxAndMinSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_auth.Chat/GetMaxAndMinSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetMaxAndMinSeq(ctx, req.(*GetMaxAndMinSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_PullMessageBySeqList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb_ws.PullMessageBySeqListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).PullMessageBySeqList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_auth.Chat/PullMessageBySeqList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).PullMessageBySeqList(ctx, req.(*pb_ws.PullMessageBySeqListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_auth.Chat/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMsg(ctx, req.(*SendMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_auth.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMaxAndMinSeq",
			Handler:    _Chat_GetMaxAndMinSeq_Handler,
		},
		{
			MethodName: "PullMessageBySeqList",
			Handler:    _Chat_PullMessageBySeqList_Handler,
		},
		{
			MethodName: "SendMsg",
			Handler:    _Chat_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chart/chat.proto",
}
