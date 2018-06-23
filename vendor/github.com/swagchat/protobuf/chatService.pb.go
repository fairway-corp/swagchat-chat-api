// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatService.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChatIncoming service

type ChatIncomingClient interface {
	// Incoming
	PostMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type chatIncomingClient struct {
	cc *grpc.ClientConn
}

func NewChatIncomingClient(cc *grpc.ClientConn) ChatIncomingClient {
	return &chatIncomingClient{cc}
}

func (c *chatIncomingClient) PostMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.ChatIncoming/PostMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChatIncoming service

type ChatIncomingServer interface {
	// Incoming
	PostMessage(context.Context, *Message) (*google_protobuf1.Empty, error)
}

func RegisterChatIncomingServer(s *grpc.Server, srv ChatIncomingServer) {
	s.RegisterService(&_ChatIncoming_serviceDesc, srv)
}

func _ChatIncoming_PostMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatIncomingServer).PostMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.ChatIncoming/PostMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatIncomingServer).PostMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatIncoming_serviceDesc = grpc.ServiceDesc{
	ServiceName: "swagchat.protobuf.ChatIncoming",
	HandlerType: (*ChatIncomingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostMessage",
			Handler:    _ChatIncoming_PostMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatService.proto",
}

// Client API for ChatOutgoing service

type ChatOutgoingClient interface {
	// Outgoing
	PostWebhookRoom(ctx context.Context, in *Room, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	PostWebhookMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type chatOutgoingClient struct {
	cc *grpc.ClientConn
}

func NewChatOutgoingClient(cc *grpc.ClientConn) ChatOutgoingClient {
	return &chatOutgoingClient{cc}
}

func (c *chatOutgoingClient) PostWebhookRoom(ctx context.Context, in *Room, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.ChatOutgoing/PostWebhookRoom", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatOutgoingClient) PostWebhookMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.ChatOutgoing/PostWebhookMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChatOutgoing service

type ChatOutgoingServer interface {
	// Outgoing
	PostWebhookRoom(context.Context, *Room) (*google_protobuf1.Empty, error)
	PostWebhookMessage(context.Context, *Message) (*google_protobuf1.Empty, error)
}

func RegisterChatOutgoingServer(s *grpc.Server, srv ChatOutgoingServer) {
	s.RegisterService(&_ChatOutgoing_serviceDesc, srv)
}

func _ChatOutgoing_PostWebhookRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Room)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatOutgoingServer).PostWebhookRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.ChatOutgoing/PostWebhookRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatOutgoingServer).PostWebhookRoom(ctx, req.(*Room))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatOutgoing_PostWebhookMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatOutgoingServer).PostWebhookMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.ChatOutgoing/PostWebhookMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatOutgoingServer).PostWebhookMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatOutgoing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "swagchat.protobuf.ChatOutgoing",
	HandlerType: (*ChatOutgoingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostWebhookRoom",
			Handler:    _ChatOutgoing_PostWebhookRoom_Handler,
		},
		{
			MethodName: "PostWebhookMessage",
			Handler:    _ChatOutgoing_PostWebhookMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatService.proto",
}

func init() { proto.RegisterFile("chatService.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xce, 0x48, 0x2c,
	0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2c,
	0x2e, 0x4f, 0x4c, 0x07, 0x09, 0x43, 0xf8, 0x49, 0xa5, 0x69, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9,
	0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9,
	0x79, 0xc5, 0x10, 0x05, 0x52, 0xd2, 0x50, 0x59, 0x98, 0x72, 0xfd, 0xd4, 0xdc, 0x82, 0x92, 0x4a,
	0xa8, 0x24, 0xd8, 0x02, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0xa8, 0x05, 0x46, 0xc1, 0x5c, 0x3c,
	0xce, 0x19, 0x89, 0x25, 0x9e, 0x79, 0xc9, 0xf9, 0xb9, 0x99, 0x79, 0xe9, 0x42, 0xce, 0x5c, 0xdc,
	0x01, 0xf9, 0xc5, 0x30, 0x45, 0x42, 0x52, 0x7a, 0x18, 0x0e, 0xd0, 0x83, 0xca, 0x49, 0x89, 0xe9,
	0x41, 0xec, 0x42, 0xc8, 0xb8, 0x82, 0xec, 0x52, 0x62, 0x30, 0xba, 0xca, 0x08, 0x31, 0xd5, 0xbf,
	0xb4, 0x24, 0x3d, 0x1f, 0x64, 0x6a, 0x2c, 0x17, 0x3f, 0xc8, 0xd4, 0xf0, 0xd4, 0xa4, 0x8c, 0xfc,
	0xfc, 0xec, 0xa0, 0xfc, 0xfc, 0x5c, 0x21, 0x71, 0x2c, 0x26, 0x83, 0x24, 0x70, 0x1a, 0x2b, 0xd9,
	0x74, 0xf9, 0xc9, 0x64, 0x26, 0x61, 0x25, 0x3e, 0xfd, 0x72, 0x88, 0x31, 0xc5, 0xfa, 0x45, 0xf9,
	0xf9, 0xb9, 0x56, 0x8c, 0x5a, 0x42, 0x69, 0x5c, 0x42, 0x48, 0xc6, 0x53, 0xe2, 0x76, 0x19, 0xb0,
	0x25, 0x62, 0x4a, 0x82, 0x08, 0x4b, 0x72, 0x21, 0x5a, 0xac, 0x18, 0xb5, 0x9c, 0xe4, 0xa2, 0x64,
	0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x61, 0xa6, 0xc3, 0xc3, 0x3a,
	0x89, 0x0d, 0xcc, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xb3, 0x3f, 0xc6, 0xc9, 0x01,
	0x00, 0x00,
}
