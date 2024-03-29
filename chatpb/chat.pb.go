// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatpb/chat.proto

package chatpb

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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,json=username,proto3" json:"Username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41cf34dba7f8bdf6, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type LoginResponse struct {
	LoginResult          string   `protobuf:"bytes,1,opt,name=LoginResult,json=loginResult,proto3" json:"LoginResult,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41cf34dba7f8bdf6, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetLoginResult() string {
	if m != nil {
		return m.LoginResult
	}
	return ""
}

type Messages struct {
	Sender               string   `protobuf:"bytes,1,opt,name=Sender,json=sender,proto3" json:"Sender,omitempty"`
	Reciever             string   `protobuf:"bytes,2,opt,name=Reciever,json=reciever,proto3" json:"Reciever,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Messages) Reset()         { *m = Messages{} }
func (m *Messages) String() string { return proto.CompactTextString(m) }
func (*Messages) ProtoMessage()    {}
func (*Messages) Descriptor() ([]byte, []int) {
	return fileDescriptor_41cf34dba7f8bdf6, []int{2}
}

func (m *Messages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Messages.Unmarshal(m, b)
}
func (m *Messages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Messages.Marshal(b, m, deterministic)
}
func (m *Messages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Messages.Merge(m, src)
}
func (m *Messages) XXX_Size() int {
	return xxx_messageInfo_Messages.Size(m)
}
func (m *Messages) XXX_DiscardUnknown() {
	xxx_messageInfo_Messages.DiscardUnknown(m)
}

var xxx_messageInfo_Messages proto.InternalMessageInfo

func (m *Messages) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Messages) GetReciever() string {
	if m != nil {
		return m.Reciever
	}
	return ""
}

func (m *Messages) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SendMessageResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=Status,json=status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageResponse) Reset()         { *m = SendMessageResponse{} }
func (m *SendMessageResponse) String() string { return proto.CompactTextString(m) }
func (*SendMessageResponse) ProtoMessage()    {}
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41cf34dba7f8bdf6, []int{3}
}

func (m *SendMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageResponse.Unmarshal(m, b)
}
func (m *SendMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageResponse.Marshal(b, m, deterministic)
}
func (m *SendMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageResponse.Merge(m, src)
}
func (m *SendMessageResponse) XXX_Size() int {
	return xxx_messageInfo_SendMessageResponse.Size(m)
}
func (m *SendMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageResponse proto.InternalMessageInfo

func (m *SendMessageResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*Messages)(nil), "Messages")
	proto.RegisterType((*SendMessageResponse)(nil), "SendMessageResponse")
}

func init() { proto.RegisterFile("chatpb/chat.proto", fileDescriptor_41cf34dba7f8bdf6) }

var fileDescriptor_41cf34dba7f8bdf6 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x8a, 0x69, 0x3a, 0xb1, 0x05, 0x57, 0x91, 0x90, 0x53, 0xc9, 0x41, 0x44, 0xe8, 0xfa,
	0xf5, 0x0f, 0xf4, 0xaa, 0x97, 0x04, 0x41, 0xbc, 0x6d, 0xe3, 0x90, 0x06, 0xda, 0x4d, 0xdc, 0x99,
	0xe4, 0x97, 0xf8, 0x83, 0x65, 0x37, 0xdb, 0x36, 0x8a, 0xa7, 0xe5, 0xcd, 0x7b, 0x3b, 0xef, 0x3d,
	0x06, 0xce, 0xca, 0xb5, 0xe2, 0x76, 0x75, 0x6b, 0x1f, 0xd9, 0x9a, 0x86, 0x9b, 0xec, 0x06, 0x4e,
	0x5f, 0x9a, 0xaa, 0xd6, 0x39, 0x7e, 0x75, 0x48, 0x2c, 0x52, 0x88, 0xde, 0x08, 0x8d, 0x56, 0x5b,
	0x4c, 0x82, 0x45, 0x70, 0x3d, 0xcd, 0xa3, 0xce, 0xe3, 0xec, 0x1e, 0x66, 0x5e, 0x4b, 0x6d, 0xa3,
	0x09, 0xc5, 0x02, 0xe2, 0xdd, 0xa0, 0xdb, 0xb0, 0xd7, 0xc7, 0x9b, 0xc3, 0x28, 0x7b, 0x87, 0xe8,
	0x15, 0x89, 0x54, 0x85, 0x24, 0x2e, 0x21, 0x2c, 0x50, 0x7f, 0xa2, 0xf1, 0xc2, 0x90, 0x1c, 0xb2,
	0x96, 0x39, 0x96, 0x35, 0xf6, 0x68, 0x92, 0xa3, 0xc1, 0xd2, 0x78, 0x2c, 0x12, 0x98, 0xf8, 0xff,
	0xc9, 0xb1, 0xa3, 0x26, 0xdb, 0x01, 0x66, 0x4b, 0x38, 0xb7, 0xdb, 0x3c, 0xbb, 0x8f, 0x64, 0x4d,
	0x58, 0x71, 0x47, 0x7b, 0x13, 0x87, 0x1e, 0xbe, 0x03, 0x88, 0x9f, 0xd7, 0x8a, 0x0b, 0x34, 0x7d,
	0x5d, 0xa2, 0xb8, 0x82, 0x13, 0x17, 0x5d, 0xcc, 0xe4, 0xb8, 0x7f, 0x3a, 0x97, 0xbf, 0x2b, 0x2e,
	0x61, 0x5e, 0xb0, 0x32, 0x6c, 0x13, 0xf6, 0x58, 0xeb, 0x4a, 0xfc, 0x51, 0xa4, 0x53, 0xb9, 0x6b,
	0x78, 0x17, 0x08, 0x09, 0xf1, 0x28, 0x95, 0x38, 0x70, 0xe9, 0x85, 0xfc, 0x27, 0xee, 0x53, 0xf4,
	0x11, 0x0e, 0x37, 0x59, 0x85, 0xee, 0x1e, 0x8f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x96,
	0xb3, 0xa9, 0xa4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	StartReciveing(ctx context.Context, in *LoginResponse, opts ...grpc.CallOption) (ChatService_StartReciveingClient, error)
	SendMessage(ctx context.Context, in *Messages, opts ...grpc.CallOption) (*SendMessageResponse, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/ChatService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) StartReciveing(ctx context.Context, in *LoginResponse, opts ...grpc.CallOption) (ChatService_StartReciveingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/ChatService/StartReciveing", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceStartReciveingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_StartReciveingClient interface {
	Recv() (*Messages, error)
	grpc.ClientStream
}

type chatServiceStartReciveingClient struct {
	grpc.ClientStream
}

func (x *chatServiceStartReciveingClient) Recv() (*Messages, error) {
	m := new(Messages)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *Messages, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/ChatService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	StartReciveing(*LoginResponse, ChatService_StartReciveingServer) error
	SendMessage(context.Context, *Messages) (*SendMessageResponse, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedChatServiceServer) StartReciveing(req *LoginResponse, srv ChatService_StartReciveingServer) error {
	return status.Errorf(codes.Unimplemented, "method StartReciveing not implemented")
}
func (*UnimplementedChatServiceServer) SendMessage(ctx context.Context, req *Messages) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_StartReciveing_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LoginResponse)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).StartReciveing(m, &chatServiceStartReciveingServer{stream})
}

type ChatService_StartReciveingServer interface {
	Send(*Messages) error
	grpc.ServerStream
}

type chatServiceStartReciveingServer struct {
	grpc.ServerStream
}

func (x *chatServiceStartReciveingServer) Send(m *Messages) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Messages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessage(ctx, req.(*Messages))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ChatService_Login_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartReciveing",
			Handler:       _ChatService_StartReciveing_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chatpb/chat.proto",
}
