// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: proto/details/book.proto

package details

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BookService_CreateBook_FullMethodName    = "/details.BookService/CreateBook"
	BookService_GetBook_FullMethodName       = "/details.BookService/GetBook"
	BookService_GetBookByISBN_FullMethodName = "/details.BookService/GetBookByISBN"
	BookService_ListBooks_FullMethodName     = "/details.BookService/ListBooks"
	BookService_UpdateBook_FullMethodName    = "/details.BookService/UpdateBook"
	BookService_DeleteBook_FullMethodName    = "/details.BookService/DeleteBook"
	BookService_SearchBooks_FullMethodName   = "/details.BookService/SearchBooks"
)

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	GetBookByISBN(ctx context.Context, in *GetBookByISBNRequest, opts ...grpc.CallOption) (*Book, error)
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error)
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*Empty, error)
	SearchBooks(ctx context.Context, in *SearchBooksRequest, opts ...grpc.CallOption) (*SearchBooksResponse, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_CreateBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_GetBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBookByISBN(ctx context.Context, in *GetBookByISBNRequest, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_GetBookByISBN_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, BookService_ListBooks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_UpdateBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, BookService_DeleteBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) SearchBooks(ctx context.Context, in *SearchBooksRequest, opts ...grpc.CallOption) (*SearchBooksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchBooksResponse)
	err := c.cc.Invoke(ctx, BookService_SearchBooks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility.
type BookServiceServer interface {
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	GetBookByISBN(context.Context, *GetBookByISBNRequest) (*Book, error)
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	DeleteBook(context.Context, *DeleteBookRequest) (*Empty, error)
	SearchBooks(context.Context, *SearchBooksRequest) (*SearchBooksResponse, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookServiceServer struct{}

func (UnimplementedBookServiceServer) CreateBook(context.Context, *CreateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookServiceServer) GetBook(context.Context, *GetBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookServiceServer) GetBookByISBN(context.Context, *GetBookByISBNRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookByISBN not implemented")
}
func (UnimplementedBookServiceServer) ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (UnimplementedBookServiceServer) UpdateBook(context.Context, *UpdateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookServiceServer) DeleteBook(context.Context, *DeleteBookRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookServiceServer) SearchBooks(context.Context, *SearchBooksRequest) (*SearchBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBooks not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}
func (UnimplementedBookServiceServer) testEmbeddedByValue()                     {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	// If the following call pancis, it indicates UnimplementedBookServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_CreateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_GetBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBookByISBN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookByISBNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBookByISBN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_GetBookByISBN_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBookByISBN(ctx, req.(*GetBookByISBNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_ListBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_UpdateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_DeleteBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_SearchBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).SearchBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_SearchBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).SearchBooks(ctx, req.(*SearchBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "details.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _BookService_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
		{
			MethodName: "GetBookByISBN",
			Handler:    _BookService_GetBookByISBN_Handler,
		},
		{
			MethodName: "ListBooks",
			Handler:    _BookService_ListBooks_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookService_DeleteBook_Handler,
		},
		{
			MethodName: "SearchBooks",
			Handler:    _BookService_SearchBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/details/book.proto",
}
