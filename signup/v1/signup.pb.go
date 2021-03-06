// Code generated by protoc-gen-go. DO NOT EDIT.
// source: signup/v1/signup.proto

package v1 // import "github.com/oswee/stubs/signup/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CreateSignupRequest struct {
	Signup               *Signup  `protobuf:"bytes,1,opt,name=signup,proto3" json:"signup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSignupRequest) Reset()         { *m = CreateSignupRequest{} }
func (m *CreateSignupRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSignupRequest) ProtoMessage()    {}
func (*CreateSignupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_signup_1a6155a5d910dc58, []int{0}
}
func (m *CreateSignupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSignupRequest.Unmarshal(m, b)
}
func (m *CreateSignupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSignupRequest.Marshal(b, m, deterministic)
}
func (dst *CreateSignupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSignupRequest.Merge(dst, src)
}
func (m *CreateSignupRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSignupRequest.Size(m)
}
func (m *CreateSignupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSignupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSignupRequest proto.InternalMessageInfo

func (m *CreateSignupRequest) GetSignup() *Signup {
	if m != nil {
		return m.Signup
	}
	return nil
}

type Signup struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FistName             string   `protobuf:"bytes,2,opt,name=fist_name,json=fistName,proto3" json:"fist_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Username             string   `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Status               int64    `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signup) Reset()         { *m = Signup{} }
func (m *Signup) String() string { return proto.CompactTextString(m) }
func (*Signup) ProtoMessage()    {}
func (*Signup) Descriptor() ([]byte, []int) {
	return fileDescriptor_signup_1a6155a5d910dc58, []int{1}
}
func (m *Signup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signup.Unmarshal(m, b)
}
func (m *Signup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signup.Marshal(b, m, deterministic)
}
func (dst *Signup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signup.Merge(dst, src)
}
func (m *Signup) XXX_Size() int {
	return xxx_messageInfo_Signup.Size(m)
}
func (m *Signup) XXX_DiscardUnknown() {
	xxx_messageInfo_Signup.DiscardUnknown(m)
}

var xxx_messageInfo_Signup proto.InternalMessageInfo

func (m *Signup) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Signup) GetFistName() string {
	if m != nil {
		return m.FistName
	}
	return ""
}

func (m *Signup) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Signup) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Signup) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Signup) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Signup) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateSignupRequest)(nil), "oswee.signup.v1.CreateSignupRequest")
	proto.RegisterType((*Signup)(nil), "oswee.signup.v1.Signup")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SignupServiceClient is the client API for SignupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SignupServiceClient interface {
	CreateSignup(ctx context.Context, in *CreateSignupRequest, opts ...grpc.CallOption) (*Signup, error)
}

type signupServiceClient struct {
	cc *grpc.ClientConn
}

func NewSignupServiceClient(cc *grpc.ClientConn) SignupServiceClient {
	return &signupServiceClient{cc}
}

func (c *signupServiceClient) CreateSignup(ctx context.Context, in *CreateSignupRequest, opts ...grpc.CallOption) (*Signup, error) {
	out := new(Signup)
	err := c.cc.Invoke(ctx, "/oswee.signup.v1.SignupService/CreateSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignupServiceServer is the server API for SignupService service.
type SignupServiceServer interface {
	CreateSignup(context.Context, *CreateSignupRequest) (*Signup, error)
}

func RegisterSignupServiceServer(s *grpc.Server, srv SignupServiceServer) {
	s.RegisterService(&_SignupService_serviceDesc, srv)
}

func _SignupService_CreateSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).CreateSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oswee.signup.v1.SignupService/CreateSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).CreateSignup(ctx, req.(*CreateSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SignupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oswee.signup.v1.SignupService",
	HandlerType: (*SignupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSignup",
			Handler:    _SignupService_CreateSignup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signup/v1/signup.proto",
}

func init() { proto.RegisterFile("signup/v1/signup.proto", fileDescriptor_signup_1a6155a5d910dc58) }

var fileDescriptor_signup_1a6155a5d910dc58 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4e, 0xc3, 0x30,
	0x10, 0x84, 0x95, 0x94, 0x46, 0xe9, 0xf2, 0x27, 0x19, 0x54, 0xa2, 0x72, 0x89, 0x22, 0x0e, 0x3d,
	0xc5, 0x6a, 0x79, 0x03, 0x90, 0xb8, 0xc1, 0x21, 0xbd, 0x71, 0x00, 0x39, 0xcd, 0x52, 0x2c, 0x35,
	0x75, 0xf0, 0xda, 0xe9, 0xab, 0xf1, 0x78, 0x28, 0x76, 0x82, 0x10, 0xd0, 0x9b, 0x67, 0xbe, 0x19,
	0x69, 0xbd, 0x0b, 0x53, 0x92, 0x9b, 0x9d, 0x6d, 0x78, 0xbb, 0xe0, 0xfe, 0x95, 0x37, 0x5a, 0x19,
	0xc5, 0xce, 0x15, 0xed, 0x11, 0xf3, 0xde, 0x6b, 0x17, 0xd9, 0x03, 0x5c, 0xdc, 0x6b, 0x14, 0x06,
	0x57, 0xce, 0x2a, 0xf0, 0xc3, 0x22, 0x19, 0xc6, 0x21, 0xf2, 0x99, 0x24, 0x48, 0x83, 0xf9, 0xf1,
	0xf2, 0x2a, 0xff, 0x55, 0xcc, 0xfb, 0x7c, 0x1f, 0xcb, 0x3e, 0x03, 0x88, 0xbc, 0xc5, 0xce, 0x20,
	0x94, 0x95, 0xeb, 0x8d, 0x8a, 0x50, 0x56, 0xec, 0x1a, 0x26, 0x6f, 0x92, 0xcc, 0xeb, 0x4e, 0xd4,
	0x98, 0x84, 0x69, 0x30, 0x9f, 0x14, 0x71, 0x67, 0x3c, 0x89, 0x1a, 0x3b, 0xb8, 0x15, 0x03, 0x1c,
	0x79, 0xd8, 0x19, 0x0e, 0x5e, 0xc2, 0x18, 0x6b, 0x21, 0xb7, 0xc9, 0x91, 0x03, 0x5e, 0xb0, 0x19,
	0xc4, 0x96, 0x50, 0xbb, 0xc6, 0xd8, 0x37, 0x06, 0xdd, 0xb1, 0x46, 0x10, 0xed, 0x95, 0xae, 0x92,
	0xc8, 0xb3, 0x41, 0xb3, 0x29, 0x44, 0x64, 0x84, 0xb1, 0x94, 0xc4, 0x6e, 0xb6, 0x5e, 0x2d, 0x5f,
	0xe0, 0xd4, 0x4f, 0xbe, 0x42, 0xdd, 0xca, 0x35, 0xb2, 0x47, 0x38, 0xf9, 0xb9, 0x13, 0x76, 0xf3,
	0xe7, 0xf3, 0xff, 0xac, 0x6c, 0x76, 0x68, 0x45, 0x77, 0xd9, 0x73, 0xba, 0x91, 0xe6, 0xdd, 0x96,
	0xf9, 0x5a, 0xd5, 0xdc, 0x85, 0x38, 0x19, 0x5b, 0x12, 0xff, 0x3e, 0x52, 0x19, 0xb9, 0xf3, 0xdc,
	0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x09, 0xe3, 0x95, 0x20, 0xb8, 0x01, 0x00, 0x00,
}
