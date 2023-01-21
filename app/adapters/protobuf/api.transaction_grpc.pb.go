// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: api.transaction.proto

package protobuf

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

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	Insert(ctx context.Context, in *RequestGetTransaction, opts ...grpc.CallOption) (*ResponseGetTransaction, error)
	Select(ctx context.Context, in *RequestGetTransaction, opts ...grpc.CallOption) (*ResponseGetTransaction, error)
	Get(ctx context.Context, in *RequestGetTransaction, opts ...grpc.CallOption) (*ResponseGetTransaction, error)
	Update(ctx context.Context, in *RequestGetTransaction, opts ...grpc.CallOption) (*ResponseGetTransaction, error)
	Delete(ctx context.Context, in *RequestGetTransaction, opts ...grpc.CallOption) (*ResponseGetTransaction, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) Insert(ctx context.Context, in *RequestGetTransaction, opts ...grpc.CallOption) (*ResponseGetTransaction, error) {
	out := new(ResponseGetTransaction)
	err := c.cc.Invoke(ctx, "/service_transaction.TransactionService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) Select(ctx context.Context, in *RequestGetTransaction, opts ...grpc.CallOption) (*ResponseGetTransaction, error) {
	out := new(ResponseGetTransaction)
	err := c.cc.Invoke(ctx, "/service_transaction.TransactionService/Select", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) Get(ctx context.Context, in *RequestGetTransaction, opts ...grpc.CallOption) (*ResponseGetTransaction, error) {
	out := new(ResponseGetTransaction)
	err := c.cc.Invoke(ctx, "/service_transaction.TransactionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) Update(ctx context.Context, in *RequestGetTransaction, opts ...grpc.CallOption) (*ResponseGetTransaction, error) {
	out := new(ResponseGetTransaction)
	err := c.cc.Invoke(ctx, "/service_transaction.TransactionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) Delete(ctx context.Context, in *RequestGetTransaction, opts ...grpc.CallOption) (*ResponseGetTransaction, error) {
	out := new(ResponseGetTransaction)
	err := c.cc.Invoke(ctx, "/service_transaction.TransactionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations must embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	Insert(context.Context, *RequestGetTransaction) (*ResponseGetTransaction, error)
	Select(context.Context, *RequestGetTransaction) (*ResponseGetTransaction, error)
	Get(context.Context, *RequestGetTransaction) (*ResponseGetTransaction, error)
	Update(context.Context, *RequestGetTransaction) (*ResponseGetTransaction, error)
	Delete(context.Context, *RequestGetTransaction) (*ResponseGetTransaction, error)
	mustEmbedUnimplementedTransactionServiceServer()
}

// UnimplementedTransactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) Insert(context.Context, *RequestGetTransaction) (*ResponseGetTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedTransactionServiceServer) Select(context.Context, *RequestGetTransaction) (*ResponseGetTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Select not implemented")
}
func (UnimplementedTransactionServiceServer) Get(context.Context, *RequestGetTransaction) (*ResponseGetTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTransactionServiceServer) Update(context.Context, *RequestGetTransaction) (*ResponseGetTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTransactionServiceServer) Delete(context.Context, *RequestGetTransaction) (*ResponseGetTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_transaction.TransactionService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Insert(ctx, req.(*RequestGetTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_Select_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Select(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_transaction.TransactionService/Select",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Select(ctx, req.(*RequestGetTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_transaction.TransactionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Get(ctx, req.(*RequestGetTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_transaction.TransactionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Update(ctx, req.(*RequestGetTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_transaction.TransactionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Delete(ctx, req.(*RequestGetTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service_transaction.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _TransactionService_Insert_Handler,
		},
		{
			MethodName: "Select",
			Handler:    _TransactionService_Select_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TransactionService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TransactionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TransactionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.transaction.proto",
}