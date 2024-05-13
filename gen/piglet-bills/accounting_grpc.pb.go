// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: api/proto/accounting.proto

package billsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PigletBillsClient is the client API for PigletBills service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PigletBillsClient interface {
	CreateBill(ctx context.Context, in *CreateBillRequest, opts ...grpc.CallOption) (*BillResponse, error)
	GetAllAccounts(ctx context.Context, in *GetSomeBillsRequest, opts ...grpc.CallOption) (*GetSomeBillsResponse, error)
	GetAllGoals(ctx context.Context, in *GetSomeBillsRequest, opts ...grpc.CallOption) (*GetSomeBillsResponse, error)
	GetBill(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*BillResponse, error)
	UpdateBill(ctx context.Context, in *UpdateBillRequest, opts ...grpc.CallOption) (*BillResponse, error)
	DeleteBill(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	VerifyBill(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	FixBillSum(ctx context.Context, in *FixBillSumRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type pigletBillsClient struct {
	cc grpc.ClientConnInterface
}

func NewPigletBillsClient(cc grpc.ClientConnInterface) PigletBillsClient {
	return &pigletBillsClient{cc}
}

func (c *pigletBillsClient) CreateBill(ctx context.Context, in *CreateBillRequest, opts ...grpc.CallOption) (*BillResponse, error) {
	out := new(BillResponse)
	err := c.cc.Invoke(ctx, "/accounting.pigletBills/CreateBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletBillsClient) GetAllAccounts(ctx context.Context, in *GetSomeBillsRequest, opts ...grpc.CallOption) (*GetSomeBillsResponse, error) {
	out := new(GetSomeBillsResponse)
	err := c.cc.Invoke(ctx, "/accounting.pigletBills/GetAllAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletBillsClient) GetAllGoals(ctx context.Context, in *GetSomeBillsRequest, opts ...grpc.CallOption) (*GetSomeBillsResponse, error) {
	out := new(GetSomeBillsResponse)
	err := c.cc.Invoke(ctx, "/accounting.pigletBills/GetAllGoals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletBillsClient) GetBill(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*BillResponse, error) {
	out := new(BillResponse)
	err := c.cc.Invoke(ctx, "/accounting.pigletBills/GetBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletBillsClient) UpdateBill(ctx context.Context, in *UpdateBillRequest, opts ...grpc.CallOption) (*BillResponse, error) {
	out := new(BillResponse)
	err := c.cc.Invoke(ctx, "/accounting.pigletBills/UpdateBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletBillsClient) DeleteBill(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/accounting.pigletBills/DeleteBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletBillsClient) VerifyBill(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/accounting.pigletBills/VerifyBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletBillsClient) FixBillSum(ctx context.Context, in *FixBillSumRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/accounting.pigletBills/FixBillSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PigletBillsServer is the server API for PigletBills service.
// All implementations must embed UnimplementedPigletBillsServer
// for forward compatibility
type PigletBillsServer interface {
	CreateBill(context.Context, *CreateBillRequest) (*BillResponse, error)
	GetAllAccounts(context.Context, *GetSomeBillsRequest) (*GetSomeBillsResponse, error)
	GetAllGoals(context.Context, *GetSomeBillsRequest) (*GetSomeBillsResponse, error)
	GetBill(context.Context, *IdRequest) (*BillResponse, error)
	UpdateBill(context.Context, *UpdateBillRequest) (*BillResponse, error)
	DeleteBill(context.Context, *IdRequest) (*SuccessResponse, error)
	VerifyBill(context.Context, *IdRequest) (*SuccessResponse, error)
	FixBillSum(context.Context, *FixBillSumRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPigletBillsServer()
}

// UnimplementedPigletBillsServer must be embedded to have forward compatible implementations.
type UnimplementedPigletBillsServer struct {
}

func (UnimplementedPigletBillsServer) CreateBill(context.Context, *CreateBillRequest) (*BillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBill not implemented")
}
func (UnimplementedPigletBillsServer) GetAllAccounts(context.Context, *GetSomeBillsRequest) (*GetSomeBillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAccounts not implemented")
}
func (UnimplementedPigletBillsServer) GetAllGoals(context.Context, *GetSomeBillsRequest) (*GetSomeBillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGoals not implemented")
}
func (UnimplementedPigletBillsServer) GetBill(context.Context, *IdRequest) (*BillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBill not implemented")
}
func (UnimplementedPigletBillsServer) UpdateBill(context.Context, *UpdateBillRequest) (*BillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBill not implemented")
}
func (UnimplementedPigletBillsServer) DeleteBill(context.Context, *IdRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBill not implemented")
}
func (UnimplementedPigletBillsServer) VerifyBill(context.Context, *IdRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyBill not implemented")
}
func (UnimplementedPigletBillsServer) FixBillSum(context.Context, *FixBillSumRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FixBillSum not implemented")
}
func (UnimplementedPigletBillsServer) mustEmbedUnimplementedPigletBillsServer() {}

// UnsafePigletBillsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PigletBillsServer will
// result in compilation errors.
type UnsafePigletBillsServer interface {
	mustEmbedUnimplementedPigletBillsServer()
}

func RegisterPigletBillsServer(s grpc.ServiceRegistrar, srv PigletBillsServer) {
	s.RegisterService(&PigletBills_ServiceDesc, srv)
}

func _PigletBills_CreateBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletBillsServer).CreateBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounting.pigletBills/CreateBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletBillsServer).CreateBill(ctx, req.(*CreateBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletBills_GetAllAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSomeBillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletBillsServer).GetAllAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounting.pigletBills/GetAllAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletBillsServer).GetAllAccounts(ctx, req.(*GetSomeBillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletBills_GetAllGoals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSomeBillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletBillsServer).GetAllGoals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounting.pigletBills/GetAllGoals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletBillsServer).GetAllGoals(ctx, req.(*GetSomeBillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletBills_GetBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletBillsServer).GetBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounting.pigletBills/GetBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletBillsServer).GetBill(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletBills_UpdateBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletBillsServer).UpdateBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounting.pigletBills/UpdateBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletBillsServer).UpdateBill(ctx, req.(*UpdateBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletBills_DeleteBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletBillsServer).DeleteBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounting.pigletBills/DeleteBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletBillsServer).DeleteBill(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletBills_VerifyBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletBillsServer).VerifyBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounting.pigletBills/VerifyBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletBillsServer).VerifyBill(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletBills_FixBillSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FixBillSumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletBillsServer).FixBillSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounting.pigletBills/FixBillSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletBillsServer).FixBillSum(ctx, req.(*FixBillSumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PigletBills_ServiceDesc is the grpc.ServiceDesc for PigletBills service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PigletBills_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accounting.pigletBills",
	HandlerType: (*PigletBillsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBill",
			Handler:    _PigletBills_CreateBill_Handler,
		},
		{
			MethodName: "GetAllAccounts",
			Handler:    _PigletBills_GetAllAccounts_Handler,
		},
		{
			MethodName: "GetAllGoals",
			Handler:    _PigletBills_GetAllGoals_Handler,
		},
		{
			MethodName: "GetBill",
			Handler:    _PigletBills_GetBill_Handler,
		},
		{
			MethodName: "UpdateBill",
			Handler:    _PigletBills_UpdateBill_Handler,
		},
		{
			MethodName: "DeleteBill",
			Handler:    _PigletBills_DeleteBill_Handler,
		},
		{
			MethodName: "VerifyBill",
			Handler:    _PigletBills_VerifyBill_Handler,
		},
		{
			MethodName: "FixBillSum",
			Handler:    _PigletBills_FixBillSum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/accounting.proto",
}
