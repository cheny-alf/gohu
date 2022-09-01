// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: info.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InfoClient is the client API for Info service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfoClient interface {
	GetNotification(ctx context.Context, in *GetNotificationReq, opts ...grpc.CallOption) (*GetNotificationRes, error)
}

type infoClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoClient(cc grpc.ClientConnInterface) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) GetNotification(ctx context.Context, in *GetNotificationReq, opts ...grpc.CallOption) (*GetNotificationRes, error) {
	out := new(GetNotificationRes)
	err := c.cc.Invoke(ctx, "/info.Info/GetNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServer is the server API for Info service.
// All implementations must embed UnimplementedInfoServer
// for forward compatibility
type InfoServer interface {
	GetNotification(context.Context, *GetNotificationReq) (*GetNotificationRes, error)
	mustEmbedUnimplementedInfoServer()
}

// UnimplementedInfoServer must be embedded to have forward compatible implementations.
type UnimplementedInfoServer struct {
}

func (UnimplementedInfoServer) GetNotification(context.Context, *GetNotificationReq) (*GetNotificationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotification not implemented")
}
func (UnimplementedInfoServer) mustEmbedUnimplementedInfoServer() {}

// UnsafeInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoServer will
// result in compilation errors.
type UnsafeInfoServer interface {
	mustEmbedUnimplementedInfoServer()
}

func RegisterInfoServer(s grpc.ServiceRegistrar, srv InfoServer) {
	s.RegisterService(&Info_ServiceDesc, srv)
}

func _Info_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/GetNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).GetNotification(ctx, req.(*GetNotificationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Info_ServiceDesc is the grpc.ServiceDesc for Info service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Info_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "info.Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotification",
			Handler:    _Info_GetNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info.proto",
}
