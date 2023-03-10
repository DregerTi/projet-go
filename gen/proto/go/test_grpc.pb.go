// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: test.proto

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

// TestApiClient is the client API for TestApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestApiClient interface {
	GetProductById(ctx context.Context, in *ProductByIdRequest, opts ...grpc.CallOption) (*ProductByIdResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error)
	GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error)
	DeleteProductById(ctx context.Context, in *DeleteProductByIdRequest, opts ...grpc.CallOption) (*DeleteProductByIdResponse, error)
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
	UpdatePayment(ctx context.Context, in *UpdatePaymentRequest, opts ...grpc.CallOption) (*UpdatePaymentResponse, error)
	GetPaymentById(ctx context.Context, in *GetPaymentByIdRequest, opts ...grpc.CallOption) (*GetPaymentByIdResponse, error)
	GetAllPayments(ctx context.Context, in *GetAllPaymentsRequest, opts ...grpc.CallOption) (*GetAllPaymentsResponse, error)
	StreamPayments(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (TestApi_StreamPaymentsClient, error)
	GetRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomResponse, error)
	PostToRoom(ctx context.Context, in *Message, opts ...grpc.CallOption) (*RoomResponse, error)
	DeleteRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomResponse, error)
}

type testApiClient struct {
	cc grpc.ClientConnInterface
}

func NewTestApiClient(cc grpc.ClientConnInterface) TestApiClient {
	return &testApiClient{cc}
}

func (c *testApiClient) GetProductById(ctx context.Context, in *ProductByIdRequest, opts ...grpc.CallOption) (*ProductByIdResponse, error) {
	out := new(ProductByIdResponse)
	err := c.cc.Invoke(ctx, "/main.TestApi/GetProductById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/main.TestApi/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error) {
	out := new(UpdateProductResponse)
	err := c.cc.Invoke(ctx, "/main.TestApi/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error) {
	out := new(GetAllProductsResponse)
	err := c.cc.Invoke(ctx, "/main.TestApi/GetAllProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) DeleteProductById(ctx context.Context, in *DeleteProductByIdRequest, opts ...grpc.CallOption) (*DeleteProductByIdResponse, error) {
	out := new(DeleteProductByIdResponse)
	err := c.cc.Invoke(ctx, "/main.TestApi/DeleteProductById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, "/main.TestApi/CreatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) UpdatePayment(ctx context.Context, in *UpdatePaymentRequest, opts ...grpc.CallOption) (*UpdatePaymentResponse, error) {
	out := new(UpdatePaymentResponse)
	err := c.cc.Invoke(ctx, "/main.TestApi/UpdatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) GetPaymentById(ctx context.Context, in *GetPaymentByIdRequest, opts ...grpc.CallOption) (*GetPaymentByIdResponse, error) {
	out := new(GetPaymentByIdResponse)
	err := c.cc.Invoke(ctx, "/main.TestApi/GetPaymentById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) GetAllPayments(ctx context.Context, in *GetAllPaymentsRequest, opts ...grpc.CallOption) (*GetAllPaymentsResponse, error) {
	out := new(GetAllPaymentsResponse)
	err := c.cc.Invoke(ctx, "/main.TestApi/GetAllPayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) StreamPayments(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (TestApi_StreamPaymentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestApi_ServiceDesc.Streams[0], "/main.TestApi/StreamPayments", opts...)
	if err != nil {
		return nil, err
	}
	x := &testApiStreamPaymentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestApi_StreamPaymentsClient interface {
	Recv() (*MessageResponse, error)
	grpc.ClientStream
}

type testApiStreamPaymentsClient struct {
	grpc.ClientStream
}

func (x *testApiStreamPaymentsClient) Recv() (*MessageResponse, error) {
	m := new(MessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testApiClient) GetRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomResponse, error) {
	out := new(RoomResponse)
	err := c.cc.Invoke(ctx, "/main.TestApi/GetRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) PostToRoom(ctx context.Context, in *Message, opts ...grpc.CallOption) (*RoomResponse, error) {
	out := new(RoomResponse)
	err := c.cc.Invoke(ctx, "/main.TestApi/PostToRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) DeleteRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomResponse, error) {
	out := new(RoomResponse)
	err := c.cc.Invoke(ctx, "/main.TestApi/DeleteRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestApiServer is the server API for TestApi service.
// All implementations must embed UnimplementedTestApiServer
// for forward compatibility
type TestApiServer interface {
	GetProductById(context.Context, *ProductByIdRequest) (*ProductByIdResponse, error)
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error)
	GetAllProducts(context.Context, *GetAllProductsRequest) (*GetAllProductsResponse, error)
	DeleteProductById(context.Context, *DeleteProductByIdRequest) (*DeleteProductByIdResponse, error)
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	UpdatePayment(context.Context, *UpdatePaymentRequest) (*UpdatePaymentResponse, error)
	GetPaymentById(context.Context, *GetPaymentByIdRequest) (*GetPaymentByIdResponse, error)
	GetAllPayments(context.Context, *GetAllPaymentsRequest) (*GetAllPaymentsResponse, error)
	StreamPayments(*RoomRequest, TestApi_StreamPaymentsServer) error
	GetRoom(context.Context, *RoomRequest) (*RoomResponse, error)
	PostToRoom(context.Context, *Message) (*RoomResponse, error)
	DeleteRoom(context.Context, *RoomRequest) (*RoomResponse, error)
	mustEmbedUnimplementedTestApiServer()
}

// UnimplementedTestApiServer must be embedded to have forward compatible implementations.
type UnimplementedTestApiServer struct {
}

func (UnimplementedTestApiServer) GetProductById(context.Context, *ProductByIdRequest) (*ProductByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (UnimplementedTestApiServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedTestApiServer) UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedTestApiServer) GetAllProducts(context.Context, *GetAllProductsRequest) (*GetAllProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (UnimplementedTestApiServer) DeleteProductById(context.Context, *DeleteProductByIdRequest) (*DeleteProductByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductById not implemented")
}
func (UnimplementedTestApiServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedTestApiServer) UpdatePayment(context.Context, *UpdatePaymentRequest) (*UpdatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayment not implemented")
}
func (UnimplementedTestApiServer) GetPaymentById(context.Context, *GetPaymentByIdRequest) (*GetPaymentByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentById not implemented")
}
func (UnimplementedTestApiServer) GetAllPayments(context.Context, *GetAllPaymentsRequest) (*GetAllPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPayments not implemented")
}
func (UnimplementedTestApiServer) StreamPayments(*RoomRequest, TestApi_StreamPaymentsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPayments not implemented")
}
func (UnimplementedTestApiServer) GetRoom(context.Context, *RoomRequest) (*RoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
func (UnimplementedTestApiServer) PostToRoom(context.Context, *Message) (*RoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostToRoom not implemented")
}
func (UnimplementedTestApiServer) DeleteRoom(context.Context, *RoomRequest) (*RoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
func (UnimplementedTestApiServer) mustEmbedUnimplementedTestApiServer() {}

// UnsafeTestApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestApiServer will
// result in compilation errors.
type UnsafeTestApiServer interface {
	mustEmbedUnimplementedTestApiServer()
}

func RegisterTestApiServer(s grpc.ServiceRegistrar, srv TestApiServer) {
	s.RegisterService(&TestApi_ServiceDesc, srv)
}

func _TestApi_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestApi/GetProductById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).GetProductById(ctx, req.(*ProductByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestApi/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestApi/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_GetAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).GetAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestApi/GetAllProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).GetAllProducts(ctx, req.(*GetAllProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_DeleteProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).DeleteProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestApi/DeleteProductById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).DeleteProductById(ctx, req.(*DeleteProductByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestApi/CreatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_UpdatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).UpdatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestApi/UpdatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).UpdatePayment(ctx, req.(*UpdatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_GetPaymentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).GetPaymentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestApi/GetPaymentById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).GetPaymentById(ctx, req.(*GetPaymentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_GetAllPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).GetAllPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestApi/GetAllPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).GetAllPayments(ctx, req.(*GetAllPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_StreamPayments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RoomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestApiServer).StreamPayments(m, &testApiStreamPaymentsServer{stream})
}

type TestApi_StreamPaymentsServer interface {
	Send(*MessageResponse) error
	grpc.ServerStream
}

type testApiStreamPaymentsServer struct {
	grpc.ServerStream
}

func (x *testApiStreamPaymentsServer) Send(m *MessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TestApi_GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestApi/GetRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).GetRoom(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_PostToRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).PostToRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestApi/PostToRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).PostToRoom(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestApi/DeleteRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).DeleteRoom(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestApi_ServiceDesc is the grpc.ServiceDesc for TestApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.TestApi",
	HandlerType: (*TestApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductById",
			Handler:    _TestApi_GetProductById_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _TestApi_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _TestApi_UpdateProduct_Handler,
		},
		{
			MethodName: "GetAllProducts",
			Handler:    _TestApi_GetAllProducts_Handler,
		},
		{
			MethodName: "DeleteProductById",
			Handler:    _TestApi_DeleteProductById_Handler,
		},
		{
			MethodName: "CreatePayment",
			Handler:    _TestApi_CreatePayment_Handler,
		},
		{
			MethodName: "UpdatePayment",
			Handler:    _TestApi_UpdatePayment_Handler,
		},
		{
			MethodName: "GetPaymentById",
			Handler:    _TestApi_GetPaymentById_Handler,
		},
		{
			MethodName: "GetAllPayments",
			Handler:    _TestApi_GetAllPayments_Handler,
		},
		{
			MethodName: "GetRoom",
			Handler:    _TestApi_GetRoom_Handler,
		},
		{
			MethodName: "PostToRoom",
			Handler:    _TestApi_PostToRoom_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _TestApi_DeleteRoom_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamPayments",
			Handler:       _TestApi_StreamPayments_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "test.proto",
}
