// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: piglet-transactions/transactions.proto

package transv1

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

// PigletTransactionsClient is the client API for PigletTransactions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PigletTransactionsClient interface {
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	UpdateTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	DeleteTransaction(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	GetTransaction(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	GetAllTransactions(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetAllTransactionsResponse, error)
	UpdateBill(ctx context.Context, in *Bill, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	UpdateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*CategoryResponse, error)
	GetCategory(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	GetAllCategories(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error)
	DeleteCategory(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type pigletTransactionsClient struct {
	cc grpc.ClientConnInterface
}

func NewPigletTransactionsClient(cc grpc.ClientConnInterface) PigletTransactionsClient {
	return &pigletTransactionsClient{cc}
}

func (c *pigletTransactionsClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/transactions.PigletTransactions/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletTransactionsClient) UpdateTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/transactions.PigletTransactions/UpdateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletTransactionsClient) DeleteTransaction(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/transactions.PigletTransactions/DeleteTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletTransactionsClient) GetTransaction(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/transactions.PigletTransactions/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletTransactionsClient) GetAllTransactions(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetAllTransactionsResponse, error) {
	out := new(GetAllTransactionsResponse)
	err := c.cc.Invoke(ctx, "/transactions.PigletTransactions/GetAllTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletTransactionsClient) UpdateBill(ctx context.Context, in *Bill, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/transactions.PigletTransactions/UpdateBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletTransactionsClient) AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/transactions.PigletTransactions/AddCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletTransactionsClient) UpdateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/transactions.PigletTransactions/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletTransactionsClient) GetCategory(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/transactions.PigletTransactions/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletTransactionsClient) GetAllCategories(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error) {
	out := new(GetAllCategoriesResponse)
	err := c.cc.Invoke(ctx, "/transactions.PigletTransactions/GetAllCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pigletTransactionsClient) DeleteCategory(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/transactions.PigletTransactions/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PigletTransactionsServer is the server API for PigletTransactions service.
// All implementations must embed UnimplementedPigletTransactionsServer
// for forward compatibility
type PigletTransactionsServer interface {
	CreateTransaction(context.Context, *CreateTransactionRequest) (*TransactionResponse, error)
	UpdateTransaction(context.Context, *Transaction) (*TransactionResponse, error)
	DeleteTransaction(context.Context, *IdRequest) (*SuccessResponse, error)
	GetTransaction(context.Context, *IdRequest) (*TransactionResponse, error)
	GetAllTransactions(context.Context, *EmptyRequest) (*GetAllTransactionsResponse, error)
	UpdateBill(context.Context, *Bill) (*emptypb.Empty, error)
	AddCategory(context.Context, *AddCategoryRequest) (*CategoryResponse, error)
	UpdateCategory(context.Context, *Category) (*CategoryResponse, error)
	GetCategory(context.Context, *IdRequest) (*CategoryResponse, error)
	GetAllCategories(context.Context, *EmptyRequest) (*GetAllCategoriesResponse, error)
	DeleteCategory(context.Context, *IdRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedPigletTransactionsServer()
}

// UnimplementedPigletTransactionsServer must be embedded to have forward compatible implementations.
type UnimplementedPigletTransactionsServer struct {
}

func (UnimplementedPigletTransactionsServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedPigletTransactionsServer) UpdateTransaction(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransaction not implemented")
}
func (UnimplementedPigletTransactionsServer) DeleteTransaction(context.Context, *IdRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransaction not implemented")
}
func (UnimplementedPigletTransactionsServer) GetTransaction(context.Context, *IdRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedPigletTransactionsServer) GetAllTransactions(context.Context, *EmptyRequest) (*GetAllTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTransactions not implemented")
}
func (UnimplementedPigletTransactionsServer) UpdateBill(context.Context, *Bill) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBill not implemented")
}
func (UnimplementedPigletTransactionsServer) AddCategory(context.Context, *AddCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedPigletTransactionsServer) UpdateCategory(context.Context, *Category) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedPigletTransactionsServer) GetCategory(context.Context, *IdRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedPigletTransactionsServer) GetAllCategories(context.Context, *EmptyRequest) (*GetAllCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategories not implemented")
}
func (UnimplementedPigletTransactionsServer) DeleteCategory(context.Context, *IdRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedPigletTransactionsServer) mustEmbedUnimplementedPigletTransactionsServer() {}

// UnsafePigletTransactionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PigletTransactionsServer will
// result in compilation errors.
type UnsafePigletTransactionsServer interface {
	mustEmbedUnimplementedPigletTransactionsServer()
}

func RegisterPigletTransactionsServer(s grpc.ServiceRegistrar, srv PigletTransactionsServer) {
	s.RegisterService(&PigletTransactions_ServiceDesc, srv)
}

func _PigletTransactions_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletTransactionsServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.PigletTransactions/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletTransactionsServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletTransactions_UpdateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletTransactionsServer).UpdateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.PigletTransactions/UpdateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletTransactionsServer).UpdateTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletTransactions_DeleteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletTransactionsServer).DeleteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.PigletTransactions/DeleteTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletTransactionsServer).DeleteTransaction(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletTransactions_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletTransactionsServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.PigletTransactions/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletTransactionsServer).GetTransaction(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletTransactions_GetAllTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletTransactionsServer).GetAllTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.PigletTransactions/GetAllTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletTransactionsServer).GetAllTransactions(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletTransactions_UpdateBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bill)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletTransactionsServer).UpdateBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.PigletTransactions/UpdateBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletTransactionsServer).UpdateBill(ctx, req.(*Bill))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletTransactions_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletTransactionsServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.PigletTransactions/AddCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletTransactionsServer).AddCategory(ctx, req.(*AddCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletTransactions_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletTransactionsServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.PigletTransactions/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletTransactionsServer).UpdateCategory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletTransactions_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletTransactionsServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.PigletTransactions/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletTransactionsServer).GetCategory(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletTransactions_GetAllCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletTransactionsServer).GetAllCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.PigletTransactions/GetAllCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletTransactionsServer).GetAllCategories(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PigletTransactions_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PigletTransactionsServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.PigletTransactions/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PigletTransactionsServer).DeleteCategory(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PigletTransactions_ServiceDesc is the grpc.ServiceDesc for PigletTransactions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PigletTransactions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transactions.PigletTransactions",
	HandlerType: (*PigletTransactionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransaction",
			Handler:    _PigletTransactions_CreateTransaction_Handler,
		},
		{
			MethodName: "UpdateTransaction",
			Handler:    _PigletTransactions_UpdateTransaction_Handler,
		},
		{
			MethodName: "DeleteTransaction",
			Handler:    _PigletTransactions_DeleteTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _PigletTransactions_GetTransaction_Handler,
		},
		{
			MethodName: "GetAllTransactions",
			Handler:    _PigletTransactions_GetAllTransactions_Handler,
		},
		{
			MethodName: "UpdateBill",
			Handler:    _PigletTransactions_UpdateBill_Handler,
		},
		{
			MethodName: "AddCategory",
			Handler:    _PigletTransactions_AddCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _PigletTransactions_UpdateCategory_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _PigletTransactions_GetCategory_Handler,
		},
		{
			MethodName: "GetAllCategories",
			Handler:    _PigletTransactions_GetAllCategories_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _PigletTransactions_DeleteCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "piglet-transactions/transactions.proto",
}
