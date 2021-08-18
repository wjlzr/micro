// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oauth.proto

package oauth

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

//user start---------------------
type LoginReq struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce0b12f599e9f07, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResp struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	CurrentAuthority     string   `protobuf:"bytes,2,opt,name=currentAuthority,proto3" json:"currentAuthority,omitempty"`
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,4,opt,name=userName,proto3" json:"userName,omitempty"`
	AccessToken          string   `protobuf:"bytes,5,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	AccessExpire         int64    `protobuf:"varint,6,opt,name=AccessExpire,proto3" json:"AccessExpire,omitempty"`
	RefreshAfter         int64    `protobuf:"varint,7,opt,name=RefreshAfter,proto3" json:"RefreshAfter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce0b12f599e9f07, []int{1}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *LoginResp) GetCurrentAuthority() string {
	if m != nil {
		return m.CurrentAuthority
	}
	return ""
}

func (m *LoginResp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LoginResp) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginResp) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *LoginResp) GetAccessExpire() int64 {
	if m != nil {
		return m.AccessExpire
	}
	return 0
}

func (m *LoginResp) GetRefreshAfter() int64 {
	if m != nil {
		return m.RefreshAfter
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "oauth.LoginReq")
	proto.RegisterType((*LoginResp)(nil), "oauth.loginResp")
}

func init() { proto.RegisterFile("oauth.proto", fileDescriptor_7ce0b12f599e9f07) }

var fileDescriptor_7ce0b12f599e9f07 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x6a, 0x62, 0x32, 0x15, 0x2d, 0x7b, 0x90, 0xa5, 0x5e, 0x42, 0x4e, 0xa5, 0x87,
	0x1e, 0xec, 0x13, 0x04, 0xf1, 0x26, 0x0a, 0xc1, 0xbb, 0xc4, 0x64, 0x6a, 0x16, 0x6d, 0x36, 0xce,
	0xec, 0xa2, 0x3e, 0xb0, 0xef, 0x21, 0xd9, 0x4d, 0xa4, 0xda, 0xe3, 0xff, 0xcd, 0xf0, 0xcf, 0xf0,
	0xc1, 0x5c, 0x57, 0xd6, 0xb4, 0x9b, 0x9e, 0xb4, 0xd1, 0x22, 0x72, 0x21, 0xbf, 0x81, 0xe4, 0x4e,
	0xbf, 0xa8, 0xae, 0xc4, 0x77, 0x71, 0x05, 0xa9, 0x65, 0xa4, 0xa7, 0xae, 0xda, 0xa3, 0x0c, 0xb2,
	0x60, 0x95, 0x96, 0xc9, 0x00, 0xee, 0xab, 0x3d, 0x8a, 0x25, 0x24, 0x7d, 0xc5, 0xfc, 0xa1, 0xa9,
	0x91, 0xa1, 0x9f, 0x4d, 0x39, 0xff, 0x0e, 0x20, 0x7d, 0xf3, 0x2d, 0xdc, 0x8b, 0x4b, 0x88, 0xd9,
	0x54, 0xc6, 0xf2, 0xd8, 0x31, 0x26, 0xb1, 0x86, 0x45, 0x6d, 0x89, 0xb0, 0x33, 0x85, 0x35, 0xad,
	0x26, 0x65, 0xbe, 0xc6, 0xa6, 0x23, 0x2e, 0xce, 0x21, 0x54, 0x8d, 0x9c, 0x65, 0xc1, 0x6a, 0x56,
	0x86, 0xaa, 0x19, 0xae, 0x4f, 0x9f, 0xc8, 0x93, 0x7f, 0x9f, 0x65, 0x30, 0x2f, 0xea, 0x1a, 0x99,
	0x1f, 0xf5, 0x2b, 0x76, 0x32, 0x72, 0xe3, 0x43, 0x24, 0x72, 0x38, 0xf3, 0xf1, 0xf6, 0xb3, 0x57,
	0x84, 0x32, 0x76, 0xbd, 0x7f, 0xd8, 0xb0, 0x53, 0xe2, 0x8e, 0x90, 0xdb, 0x62, 0x67, 0x90, 0xe4,
	0xa9, 0xdf, 0x39, 0x64, 0xd7, 0x5b, 0x88, 0x1e, 0x06, 0x6b, 0x62, 0x0d, 0x91, 0xb3, 0x26, 0x2e,
	0x36, 0xde, 0xe9, 0xe4, 0x70, 0xb9, 0x18, 0xc1, 0xaf, 0x8e, 0xe7, 0xd8, 0xf9, 0xde, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x4e, 0x1e, 0x8e, 0x7e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OauthClient is the client API for Oauth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OauthClient interface {
	// Login 登录
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
}

type oauthClient struct {
	cc *grpc.ClientConn
}

func NewOauthClient(cc *grpc.ClientConn) OauthClient {
	return &oauthClient{cc}
}

func (c *oauthClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/oauth.Oauth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OauthServer is the server API for Oauth service.
type OauthServer interface {
	// Login 登录
	Login(context.Context, *LoginReq) (*LoginResp, error)
}

// UnimplementedOauthServer can be embedded to have forward compatible implementations.
type UnimplementedOauthServer struct {
}

func (*UnimplementedOauthServer) Login(ctx context.Context, req *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterOauthServer(s *grpc.Server, srv OauthServer) {
	s.RegisterService(&_Oauth_serviceDesc, srv)
}

func _Oauth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.Oauth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Oauth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oauth.Oauth",
	HandlerType: (*OauthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Oauth_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauth.proto",
}