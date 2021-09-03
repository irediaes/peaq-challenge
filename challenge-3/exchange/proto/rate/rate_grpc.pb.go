// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rate

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

// RateServiceClient is the client API for RateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RateServiceClient interface {
	// Fetch Processed Growth Records
	GetGrowthRecords(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
	// Fetch Raw Growth Records
	GetRawGrowthRecords(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*RawResponse, error)
}

type rateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRateServiceClient(cc grpc.ClientConnInterface) RateServiceClient {
	return &rateServiceClient{cc}
}

func (c *rateServiceClient) GetGrowthRecords(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rate.RateService/GetGrowthRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateServiceClient) GetRawGrowthRecords(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*RawResponse, error) {
	out := new(RawResponse)
	err := c.cc.Invoke(ctx, "/rate.RateService/GetRawGrowthRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateServiceServer is the server API for RateService service.
// All implementations must embed UnimplementedRateServiceServer
// for forward compatibility
type RateServiceServer interface {
	// Fetch Processed Growth Records
	GetGrowthRecords(context.Context, *GetRequest) (*Response, error)
	// Fetch Raw Growth Records
	GetRawGrowthRecords(context.Context, *GetRequest) (*RawResponse, error)
	mustEmbedUnimplementedRateServiceServer()
}

// UnimplementedRateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRateServiceServer struct {
}

func (UnimplementedRateServiceServer) GetGrowthRecords(context.Context, *GetRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGrowthRecords not implemented")
}
func (UnimplementedRateServiceServer) GetRawGrowthRecords(context.Context, *GetRequest) (*RawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRawGrowthRecords not implemented")
}
func (UnimplementedRateServiceServer) mustEmbedUnimplementedRateServiceServer() {}

// UnsafeRateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RateServiceServer will
// result in compilation errors.
type UnsafeRateServiceServer interface {
	mustEmbedUnimplementedRateServiceServer()
}

func RegisterRateServiceServer(s grpc.ServiceRegistrar, srv RateServiceServer) {
	s.RegisterService(&RateService_ServiceDesc, srv)
}

func _RateService_GetGrowthRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).GetGrowthRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rate.RateService/GetGrowthRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).GetGrowthRecords(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateService_GetRawGrowthRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).GetRawGrowthRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rate.RateService/GetRawGrowthRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).GetRawGrowthRecords(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RateService_ServiceDesc is the grpc.ServiceDesc for RateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rate.RateService",
	HandlerType: (*RateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGrowthRecords",
			Handler:    _RateService_GetGrowthRecords_Handler,
		},
		{
			MethodName: "GetRawGrowthRecords",
			Handler:    _RateService_GetRawGrowthRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rate.proto",
}
