// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: sampleManager.proto

package proto

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

// SampleManagerClient is the client API for SampleManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleManagerClient interface {
	CreateMapping(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	GetSampleId(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type sampleManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleManagerClient(cc grpc.ClientConnInterface) SampleManagerClient {
	return &sampleManagerClient{cc}
}

func (c *sampleManagerClient) CreateMapping(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/proto.SampleManager/CreateMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleManagerClient) GetSampleId(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/proto.SampleManager/GetSampleId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleManagerServer is the server API for SampleManager service.
// All implementations must embed UnimplementedSampleManagerServer
// for forward compatibility
type SampleManagerServer interface {
	CreateMapping(context.Context, *CreateRequest) (*CreateResponse, error)
	GetSampleId(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedSampleManagerServer()
}

// UnimplementedSampleManagerServer must be embedded to have forward compatible implementations.
type UnimplementedSampleManagerServer struct {
}

func (UnimplementedSampleManagerServer) CreateMapping(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMapping not implemented")
}
func (UnimplementedSampleManagerServer) GetSampleId(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSampleId not implemented")
}
func (UnimplementedSampleManagerServer) mustEmbedUnimplementedSampleManagerServer() {}

// UnsafeSampleManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleManagerServer will
// result in compilation errors.
type UnsafeSampleManagerServer interface {
	mustEmbedUnimplementedSampleManagerServer()
}

func RegisterSampleManagerServer(s grpc.ServiceRegistrar, srv SampleManagerServer) {
	s.RegisterService(&SampleManager_ServiceDesc, srv)
}

func _SampleManager_CreateMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleManagerServer).CreateMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SampleManager/CreateMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleManagerServer).CreateMapping(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleManager_GetSampleId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleManagerServer).GetSampleId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SampleManager/GetSampleId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleManagerServer).GetSampleId(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SampleManager_ServiceDesc is the grpc.ServiceDesc for SampleManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SampleManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SampleManager",
	HandlerType: (*SampleManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMapping",
			Handler:    _SampleManager_CreateMapping_Handler,
		},
		{
			MethodName: "GetSampleId",
			Handler:    _SampleManager_GetSampleId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sampleManager.proto",
}
