// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crm/v1/crm.proto

package v1 // import "github.com/oswee/stubs/crm/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListCustomersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCustomersRequest) Reset()         { *m = ListCustomersRequest{} }
func (m *ListCustomersRequest) String() string { return proto.CompactTextString(m) }
func (*ListCustomersRequest) ProtoMessage()    {}
func (*ListCustomersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_crm_672403bd48afa943, []int{0}
}
func (m *ListCustomersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCustomersRequest.Unmarshal(m, b)
}
func (m *ListCustomersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCustomersRequest.Marshal(b, m, deterministic)
}
func (dst *ListCustomersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCustomersRequest.Merge(dst, src)
}
func (m *ListCustomersRequest) XXX_Size() int {
	return xxx_messageInfo_ListCustomersRequest.Size(m)
}
func (m *ListCustomersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCustomersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCustomersRequest proto.InternalMessageInfo

type ListCustomersResponse struct {
	Customers            []*Customer `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListCustomersResponse) Reset()         { *m = ListCustomersResponse{} }
func (m *ListCustomersResponse) String() string { return proto.CompactTextString(m) }
func (*ListCustomersResponse) ProtoMessage()    {}
func (*ListCustomersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_crm_672403bd48afa943, []int{1}
}
func (m *ListCustomersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCustomersResponse.Unmarshal(m, b)
}
func (m *ListCustomersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCustomersResponse.Marshal(b, m, deterministic)
}
func (dst *ListCustomersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCustomersResponse.Merge(dst, src)
}
func (m *ListCustomersResponse) XXX_Size() int {
	return xxx_messageInfo_ListCustomersResponse.Size(m)
}
func (m *ListCustomersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCustomersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCustomersResponse proto.InternalMessageInfo

func (m *ListCustomersResponse) GetCustomers() []*Customer {
	if m != nil {
		return m.Customers
	}
	return nil
}

type Customer struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_crm_672403bd48afa943, []int{2}
}
func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (dst *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(dst, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ListCustomersRequest)(nil), "oswee.crm.v1.ListCustomersRequest")
	proto.RegisterType((*ListCustomersResponse)(nil), "oswee.crm.v1.ListCustomersResponse")
	proto.RegisterType((*Customer)(nil), "oswee.crm.v1.Customer")
}

func init() { proto.RegisterFile("crm/v1/crm.proto", fileDescriptor_crm_672403bd48afa943) }

var fileDescriptor_crm_672403bd48afa943 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2e, 0xca, 0xd5,
	0x2f, 0x33, 0xd4, 0x4f, 0x2e, 0xca, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc9, 0x2f,
	0x2e, 0x4f, 0x4d, 0xd5, 0x03, 0x09, 0x94, 0x19, 0x2a, 0x89, 0x71, 0x89, 0xf8, 0x64, 0x16, 0x97,
	0x38, 0x97, 0x16, 0x97, 0xe4, 0xe7, 0xa6, 0x16, 0x15, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x28, 0xf9, 0x72, 0x89, 0xa2, 0x89, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x99, 0x70, 0x71,
	0x26, 0xc3, 0x04, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xc4, 0xf4, 0x90, 0x8d, 0xd4, 0x83,
	0xe9, 0x09, 0x42, 0x28, 0x54, 0xd2, 0xe3, 0xe2, 0x80, 0x09, 0x0b, 0xf1, 0x71, 0x31, 0x65, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25,
	0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x46, 0xa9, 0x5c, 0x9c, 0x70,
	0xab, 0x85, 0x22, 0xb8, 0x78, 0x51, 0xdc, 0x22, 0xa4, 0x84, 0x6a, 0x21, 0x36, 0x0f, 0x48, 0x29,
	0xe3, 0x55, 0x03, 0xf1, 0x8c, 0x93, 0x7c, 0x94, 0x6c, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0x3e, 0x58, 0x83, 0x7e, 0x71, 0x49, 0x69, 0x52, 0xb1, 0x3e, 0x24, 0xd8, 0x92,
	0xd8, 0xc0, 0x61, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x93, 0x66, 0xf4, 0x24, 0x47, 0x01,
	0x00, 0x00,
}