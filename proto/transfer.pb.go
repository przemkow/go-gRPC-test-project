// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transfer.proto

package todo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Todo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Done                 bool     `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96c3e6bcafb460d3, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Todo) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type NoArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoArgs) Reset()         { *m = NoArgs{} }
func (m *NoArgs) String() string { return proto.CompactTextString(m) }
func (*NoArgs) ProtoMessage()    {}
func (*NoArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_96c3e6bcafb460d3, []int{1}
}

func (m *NoArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoArgs.Unmarshal(m, b)
}
func (m *NoArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoArgs.Marshal(b, m, deterministic)
}
func (m *NoArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoArgs.Merge(m, src)
}
func (m *NoArgs) XXX_Size() int {
	return xxx_messageInfo_NoArgs.Size(m)
}
func (m *NoArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_NoArgs.DiscardUnknown(m)
}

var xxx_messageInfo_NoArgs proto.InternalMessageInfo

type Response struct {
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Todo                 *Todo    `protobuf:"bytes,2,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_96c3e6bcafb460d3, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func init() {
	proto.RegisterType((*Todo)(nil), "todo.Todo")
	proto.RegisterType((*NoArgs)(nil), "todo.NoArgs")
	proto.RegisterType((*Response)(nil), "todo.Response")
}

func init() { proto.RegisterFile("transfer.proto", fileDescriptor_96c3e6bcafb460d3) }

var fileDescriptor_96c3e6bcafb460d3 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x4f, 0xcd, 0x4a, 0xc6, 0x30,
	0x10, 0xfc, 0x22, 0xe5, 0x6b, 0xdc, 0x4a, 0x0f, 0x7b, 0x2a, 0x3d, 0x48, 0xc9, 0xc5, 0x9e, 0x82,
	0xd4, 0x27, 0x28, 0x08, 0xde, 0x3c, 0x44, 0x1f, 0xc0, 0xda, 0xac, 0xe2, 0xc1, 0x6c, 0x49, 0x82,
	0xcf, 0x2f, 0xd9, 0x52, 0xe8, 0x6d, 0x7e, 0x92, 0x99, 0x59, 0x68, 0x73, 0x5c, 0x42, 0xfa, 0xa2,
	0x68, 0xb7, 0xc8, 0x99, 0xb1, 0xca, 0xec, 0xd9, 0x58, 0xa8, 0xde, 0xd9, 0x33, 0x22, 0x54, 0x61,
	0xf9, 0xa5, 0x4e, 0x0d, 0x6a, 0xbc, 0x75, 0x82, 0x8b, 0xe6, 0x39, 0x50, 0x77, 0x33, 0xa8, 0x51,
	0x3b, 0xc1, 0x46, 0xc3, 0xf5, 0x95, 0xe7, 0xf8, 0x9d, 0xcc, 0x33, 0x68, 0x47, 0x69, 0xe3, 0x90,
	0x08, 0x3b, 0xa8, 0xd7, 0x48, 0x4b, 0x26, 0x2f, 0x01, 0xda, 0x1d, 0x14, 0xef, 0x41, 0x7a, 0x24,
	0xa3, 0x99, 0xc0, 0x16, 0x62, 0x4b, 0xa3, 0x13, 0x7d, 0xfa, 0x80, 0xa6, 0xb0, 0x37, 0x8a, 0x7f,
	0x3f, 0x2b, 0xe1, 0x03, 0xd4, 0xb3, 0xf7, 0xb2, 0xe8, 0xf4, 0xb6, 0x6f, 0x77, 0x7c, 0xf4, 0x99,
	0x0b, 0x8e, 0xa0, 0x5f, 0x28, 0x17, 0x33, 0xe1, 0xdd, 0xee, 0xee, 0xbb, 0xfa, 0xd3, 0x3f, 0x73,
	0x79, 0x54, 0x9f, 0x57, 0x39, 0xf7, 0xe9, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x74, 0x6e, 0x1d, 0xad,
	0x00, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	AddTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Response, error)
	GetTodos(ctx context.Context, in *NoArgs, opts ...grpc.CallOption) (TodoService_GetTodosClient, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) AddTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/todo.TodoService/AddTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodos(ctx context.Context, in *NoArgs, opts ...grpc.CallOption) (TodoService_GetTodosClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TodoService_serviceDesc.Streams[0], "/todo.TodoService/GetTodos", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoServiceGetTodosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoService_GetTodosClient interface {
	Recv() (*Todo, error)
	grpc.ClientStream
}

type todoServiceGetTodosClient struct {
	grpc.ClientStream
}

func (x *todoServiceGetTodosClient) Recv() (*Todo, error) {
	m := new(Todo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	AddTodo(context.Context, *Todo) (*Response, error)
	GetTodos(*NoArgs, TodoService_GetTodosServer) error
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) AddTodo(ctx context.Context, req *Todo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodo not implemented")
}
func (*UnimplementedTodoServiceServer) GetTodos(req *NoArgs, srv TodoService_GetTodosServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTodos not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_AddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).AddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/AddTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).AddTodo(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoArgs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServiceServer).GetTodos(m, &todoServiceGetTodosServer{stream})
}

type TodoService_GetTodosServer interface {
	Send(*Todo) error
	grpc.ServerStream
}

type todoServiceGetTodosServer struct {
	grpc.ServerStream
}

func (x *todoServiceGetTodosServer) Send(m *Todo) error {
	return x.ServerStream.SendMsg(m)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTodo",
			Handler:    _TodoService_AddTodo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTodos",
			Handler:       _TodoService_GetTodos_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transfer.proto",
}