// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/auth/v1/auth.proto

package authv1pb

import (
	context "context"
	fmt "fmt"
	rpc "github.com/cernbox/cs3apis/gen/proto/go/cs3/rpc"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GenerateAccessTokenRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateAccessTokenRequest) Reset()         { *m = GenerateAccessTokenRequest{} }
func (m *GenerateAccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateAccessTokenRequest) ProtoMessage()    {}
func (*GenerateAccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9840c7a4d8bddec9, []int{0}
}

func (m *GenerateAccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateAccessTokenRequest.Unmarshal(m, b)
}
func (m *GenerateAccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateAccessTokenRequest.Marshal(b, m, deterministic)
}
func (m *GenerateAccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateAccessTokenRequest.Merge(m, src)
}
func (m *GenerateAccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateAccessTokenRequest.Size(m)
}
func (m *GenerateAccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateAccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateAccessTokenRequest proto.InternalMessageInfo

func (m *GenerateAccessTokenRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GenerateAccessTokenRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GenerateAccessTokenResponse struct {
	Status               *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	AccessToken          string      `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GenerateAccessTokenResponse) Reset()         { *m = GenerateAccessTokenResponse{} }
func (m *GenerateAccessTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateAccessTokenResponse) ProtoMessage()    {}
func (*GenerateAccessTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9840c7a4d8bddec9, []int{1}
}

func (m *GenerateAccessTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateAccessTokenResponse.Unmarshal(m, b)
}
func (m *GenerateAccessTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateAccessTokenResponse.Marshal(b, m, deterministic)
}
func (m *GenerateAccessTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateAccessTokenResponse.Merge(m, src)
}
func (m *GenerateAccessTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateAccessTokenResponse.Size(m)
}
func (m *GenerateAccessTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateAccessTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateAccessTokenResponse proto.InternalMessageInfo

func (m *GenerateAccessTokenResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GenerateAccessTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type WhoAmIRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WhoAmIRequest) Reset()         { *m = WhoAmIRequest{} }
func (m *WhoAmIRequest) String() string { return proto.CompactTextString(m) }
func (*WhoAmIRequest) ProtoMessage()    {}
func (*WhoAmIRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9840c7a4d8bddec9, []int{2}
}

func (m *WhoAmIRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoAmIRequest.Unmarshal(m, b)
}
func (m *WhoAmIRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoAmIRequest.Marshal(b, m, deterministic)
}
func (m *WhoAmIRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoAmIRequest.Merge(m, src)
}
func (m *WhoAmIRequest) XXX_Size() int {
	return xxx_messageInfo_WhoAmIRequest.Size(m)
}
func (m *WhoAmIRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoAmIRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WhoAmIRequest proto.InternalMessageInfo

func (m *WhoAmIRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type WhoAmIResponse struct {
	Status               *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	User                 *User       `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WhoAmIResponse) Reset()         { *m = WhoAmIResponse{} }
func (m *WhoAmIResponse) String() string { return proto.CompactTextString(m) }
func (*WhoAmIResponse) ProtoMessage()    {}
func (*WhoAmIResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9840c7a4d8bddec9, []int{3}
}

func (m *WhoAmIResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoAmIResponse.Unmarshal(m, b)
}
func (m *WhoAmIResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoAmIResponse.Marshal(b, m, deterministic)
}
func (m *WhoAmIResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoAmIResponse.Merge(m, src)
}
func (m *WhoAmIResponse) XXX_Size() int {
	return xxx_messageInfo_WhoAmIResponse.Size(m)
}
func (m *WhoAmIResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoAmIResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WhoAmIResponse proto.InternalMessageInfo

func (m *WhoAmIResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *WhoAmIResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*GenerateAccessTokenRequest)(nil), "cs3.authv1.GenerateAccessTokenRequest")
	proto.RegisterType((*GenerateAccessTokenResponse)(nil), "cs3.authv1.GenerateAccessTokenResponse")
	proto.RegisterType((*WhoAmIRequest)(nil), "cs3.authv1.WhoAmIRequest")
	proto.RegisterType((*WhoAmIResponse)(nil), "cs3.authv1.WhoAmIResponse")
}

func init() { proto.RegisterFile("cs3/auth/v1/auth.proto", fileDescriptor_9840c7a4d8bddec9) }

var fileDescriptor_9840c7a4d8bddec9 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0x96, 0xa3, 0xca, 0x4a, 0xc6, 0x6d, 0x5a, 0xb9, 0x55, 0x95, 0x3a, 0x3d, 0xb4, 0x11, 0x22,
	0x9c, 0x6c, 0xc5, 0x7e, 0x00, 0x64, 0xe7, 0x80, 0xe0, 0x42, 0xe4, 0x84, 0x1f, 0x21, 0xa4, 0x68,
	0xb3, 0x0c, 0x89, 0x05, 0xf1, 0x9a, 0xdd, 0x75, 0x78, 0x1f, 0x8e, 0x3c, 0x00, 0x0f, 0xc1, 0x53,
	0xa1, 0x5d, 0xdb, 0xc4, 0x40, 0x10, 0xe2, 0x64, 0x7d, 0xf3, 0xfd, 0x78, 0x66, 0x76, 0xe0, 0x37,
	0x15, 0x81, 0x47, 0x72, 0xb9, 0xf0, 0x56, 0x03, 0xfd, 0x75, 0x33, 0xce, 0x24, 0xb3, 0x81, 0x8a,
	0xc0, 0x55, 0x78, 0x35, 0x70, 0xba, 0x75, 0x0d, 0x47, 0xc1, 0x72, 0x4e, 0x51, 0x14, 0x42, 0xe7,
	0x97, 0x22, 0x79, 0x46, 0x3d, 0x21, 0x89, 0xcc, 0xab, 0xea, 0xdf, 0x39, 0x63, 0xf3, 0x6b, 0xf4,
	0x48, 0x96, 0x78, 0x24, 0x4d, 0x99, 0x24, 0x32, 0x61, 0x69, 0xc9, 0xf6, 0x26, 0xe0, 0xec, 0x61,
	0x8a, 0x9c, 0x48, 0x0c, 0x29, 0x45, 0x21, 0x26, 0xec, 0x0a, 0xd3, 0x18, 0x6f, 0x72, 0x14, 0xd2,
	0x76, 0xa0, 0x99, 0x0b, 0xe4, 0x29, 0x59, 0x62, 0xc7, 0xf8, 0x67, 0xec, 0xb4, 0xe2, 0x67, 0xac,
	0xb8, 0x8c, 0x08, 0x71, 0xcb, 0xf8, 0x45, 0xa7, 0x51, 0x70, 0x15, 0xee, 0x25, 0xd0, 0xdd, 0x98,
	0x2a, 0x32, 0x96, 0x0a, 0xb4, 0xfb, 0x60, 0x16, 0x2d, 0xea, 0x50, 0xcb, 0xff, 0xee, 0xaa, 0x11,
	0x79, 0x46, 0xdd, 0xb1, 0x2e, 0xc7, 0x25, 0x6d, 0xff, 0x87, 0xaf, 0x44, 0xfb, 0xa7, 0x52, 0x05,
	0x94, 0xff, 0xb1, 0xc8, 0x3a, 0xb3, 0xe7, 0xc3, 0xb7, 0x93, 0x05, 0x0b, 0x97, 0xfb, 0x55, 0xcf,
	0xaf, 0x3d, 0xc6, 0x5b, 0xcf, 0x14, 0xda, 0x95, 0xe7, 0xb3, 0x1d, 0x6d, 0xc1, 0x17, 0xb5, 0x01,
	0xdd, 0x89, 0xe5, 0xff, 0x70, 0xd7, 0x6f, 0xe3, 0x1e, 0x09, 0xe4, 0xb1, 0x66, 0xfd, 0x07, 0x03,
	0xac, 0x30, 0x97, 0x8b, 0x31, 0xf2, 0x55, 0x42, 0xd1, 0xbe, 0x84, 0x9f, 0x1b, 0xf6, 0x61, 0x6f,
	0xd7, 0xed, 0xef, 0x3f, 0x83, 0xd3, 0xff, 0x50, 0x57, 0x8e, 0xb1, 0x0b, 0x66, 0x31, 0x98, 0xfd,
	0xa7, 0x6e, 0x79, 0xb1, 0x20, 0xc7, 0xd9, 0x44, 0x15, 0x01, 0xd1, 0x01, 0xb4, 0x29, 0x5b, 0xd6,
	0x04, 0x51, 0x4b, 0xcd, 0x31, 0x52, 0xb7, 0x32, 0x32, 0xce, 0x9a, 0x45, 0x31, 0x9b, 0xdd, 0x35,
	0xcc, 0x61, 0x74, 0x78, 0x1a, 0x46, 0xf7, 0x0d, 0x18, 0x8e, 0x03, 0x57, 0x89, 0x8e, 0x07, 0x8f,
	0x1a, 0x9c, 0x17, 0x60, 0x66, 0xea, 0x0b, 0x0b, 0x9e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xc9,
	0xc4, 0x5e, 0xd8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error)
	WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error) {
	out := new(GenerateAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/cs3.authv1.AuthService/GenerateAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error) {
	out := new(WhoAmIResponse)
	err := c.cc.Invoke(ctx, "/cs3.authv1.AuthService/WhoAmI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	GenerateAccessToken(context.Context, *GenerateAccessTokenRequest) (*GenerateAccessTokenResponse, error)
	WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_GenerateAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GenerateAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.authv1.AuthService/GenerateAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GenerateAccessToken(ctx, req.(*GenerateAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.authv1.AuthService/WhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).WhoAmI(ctx, req.(*WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.authv1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateAccessToken",
			Handler:    _AuthService_GenerateAccessToken_Handler,
		},
		{
			MethodName: "WhoAmI",
			Handler:    _AuthService_WhoAmI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/auth/v1/auth.proto",
}
