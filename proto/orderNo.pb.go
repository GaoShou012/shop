// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderNo.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type OrderNoGenRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderNoGenRequest) Reset()         { *m = OrderNoGenRequest{} }
func (m *OrderNoGenRequest) String() string { return proto.CompactTextString(m) }
func (*OrderNoGenRequest) ProtoMessage()    {}
func (*OrderNoGenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11a259be2d7b15d, []int{0}
}

func (m *OrderNoGenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderNoGenRequest.Unmarshal(m, b)
}
func (m *OrderNoGenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderNoGenRequest.Marshal(b, m, deterministic)
}
func (m *OrderNoGenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderNoGenRequest.Merge(m, src)
}
func (m *OrderNoGenRequest) XXX_Size() int {
	return xxx_messageInfo_OrderNoGenRequest.Size(m)
}
func (m *OrderNoGenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderNoGenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderNoGenRequest proto.InternalMessageInfo

type OrderNoGenResponse struct {
	Code                 uint64   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	No                   uint64   `protobuf:"varint,3,opt,name=no,proto3" json:"no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderNoGenResponse) Reset()         { *m = OrderNoGenResponse{} }
func (m *OrderNoGenResponse) String() string { return proto.CompactTextString(m) }
func (*OrderNoGenResponse) ProtoMessage()    {}
func (*OrderNoGenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11a259be2d7b15d, []int{1}
}

func (m *OrderNoGenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderNoGenResponse.Unmarshal(m, b)
}
func (m *OrderNoGenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderNoGenResponse.Marshal(b, m, deterministic)
}
func (m *OrderNoGenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderNoGenResponse.Merge(m, src)
}
func (m *OrderNoGenResponse) XXX_Size() int {
	return xxx_messageInfo_OrderNoGenResponse.Size(m)
}
func (m *OrderNoGenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderNoGenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderNoGenResponse proto.InternalMessageInfo

func (m *OrderNoGenResponse) GetCode() uint64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *OrderNoGenResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *OrderNoGenResponse) GetNo() uint64 {
	if m != nil {
		return m.No
	}
	return 0
}

type OrderNoEncryptFormat struct {
	No                   uint64   `protobuf:"varint,1,opt,name=no,proto3" json:"no,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Time                 uint64   `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	Timeout              uint64   `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderNoEncryptFormat) Reset()         { *m = OrderNoEncryptFormat{} }
func (m *OrderNoEncryptFormat) String() string { return proto.CompactTextString(m) }
func (*OrderNoEncryptFormat) ProtoMessage()    {}
func (*OrderNoEncryptFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11a259be2d7b15d, []int{2}
}

func (m *OrderNoEncryptFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderNoEncryptFormat.Unmarshal(m, b)
}
func (m *OrderNoEncryptFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderNoEncryptFormat.Marshal(b, m, deterministic)
}
func (m *OrderNoEncryptFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderNoEncryptFormat.Merge(m, src)
}
func (m *OrderNoEncryptFormat) XXX_Size() int {
	return xxx_messageInfo_OrderNoEncryptFormat.Size(m)
}
func (m *OrderNoEncryptFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderNoEncryptFormat.DiscardUnknown(m)
}

var xxx_messageInfo_OrderNoEncryptFormat proto.InternalMessageInfo

func (m *OrderNoEncryptFormat) GetNo() uint64 {
	if m != nil {
		return m.No
	}
	return 0
}

func (m *OrderNoEncryptFormat) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderNoEncryptFormat) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *OrderNoEncryptFormat) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *OrderNoEncryptFormat) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type OrderNoEncryptReq struct {
	Data                 *OrderNoEncryptFormat `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OrderNoEncryptReq) Reset()         { *m = OrderNoEncryptReq{} }
func (m *OrderNoEncryptReq) String() string { return proto.CompactTextString(m) }
func (*OrderNoEncryptReq) ProtoMessage()    {}
func (*OrderNoEncryptReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11a259be2d7b15d, []int{3}
}

func (m *OrderNoEncryptReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderNoEncryptReq.Unmarshal(m, b)
}
func (m *OrderNoEncryptReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderNoEncryptReq.Marshal(b, m, deterministic)
}
func (m *OrderNoEncryptReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderNoEncryptReq.Merge(m, src)
}
func (m *OrderNoEncryptReq) XXX_Size() int {
	return xxx_messageInfo_OrderNoEncryptReq.Size(m)
}
func (m *OrderNoEncryptReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderNoEncryptReq.DiscardUnknown(m)
}

var xxx_messageInfo_OrderNoEncryptReq proto.InternalMessageInfo

func (m *OrderNoEncryptReq) GetData() *OrderNoEncryptFormat {
	if m != nil {
		return m.Data
	}
	return nil
}

type OrderNoEncryptRsp struct {
	Code                 uint64   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              uint64   `protobuf:"varint,2,opt,name=message,proto3" json:"message,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderNoEncryptRsp) Reset()         { *m = OrderNoEncryptRsp{} }
func (m *OrderNoEncryptRsp) String() string { return proto.CompactTextString(m) }
func (*OrderNoEncryptRsp) ProtoMessage()    {}
func (*OrderNoEncryptRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11a259be2d7b15d, []int{4}
}

func (m *OrderNoEncryptRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderNoEncryptRsp.Unmarshal(m, b)
}
func (m *OrderNoEncryptRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderNoEncryptRsp.Marshal(b, m, deterministic)
}
func (m *OrderNoEncryptRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderNoEncryptRsp.Merge(m, src)
}
func (m *OrderNoEncryptRsp) XXX_Size() int {
	return xxx_messageInfo_OrderNoEncryptRsp.Size(m)
}
func (m *OrderNoEncryptRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderNoEncryptRsp.DiscardUnknown(m)
}

var xxx_messageInfo_OrderNoEncryptRsp proto.InternalMessageInfo

func (m *OrderNoEncryptRsp) GetCode() uint64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *OrderNoEncryptRsp) GetMessage() uint64 {
	if m != nil {
		return m.Message
	}
	return 0
}

func (m *OrderNoEncryptRsp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type OrderNoDecryptReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderNoDecryptReq) Reset()         { *m = OrderNoDecryptReq{} }
func (m *OrderNoDecryptReq) String() string { return proto.CompactTextString(m) }
func (*OrderNoDecryptReq) ProtoMessage()    {}
func (*OrderNoDecryptReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11a259be2d7b15d, []int{5}
}

func (m *OrderNoDecryptReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderNoDecryptReq.Unmarshal(m, b)
}
func (m *OrderNoDecryptReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderNoDecryptReq.Marshal(b, m, deterministic)
}
func (m *OrderNoDecryptReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderNoDecryptReq.Merge(m, src)
}
func (m *OrderNoDecryptReq) XXX_Size() int {
	return xxx_messageInfo_OrderNoDecryptReq.Size(m)
}
func (m *OrderNoDecryptReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderNoDecryptReq.DiscardUnknown(m)
}

var xxx_messageInfo_OrderNoDecryptReq proto.InternalMessageInfo

func (m *OrderNoDecryptReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type OrderNoDecryptRsp struct {
	Code                 uint64                `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              uint64                `protobuf:"varint,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *OrderNoEncryptFormat `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OrderNoDecryptRsp) Reset()         { *m = OrderNoDecryptRsp{} }
func (m *OrderNoDecryptRsp) String() string { return proto.CompactTextString(m) }
func (*OrderNoDecryptRsp) ProtoMessage()    {}
func (*OrderNoDecryptRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11a259be2d7b15d, []int{6}
}

func (m *OrderNoDecryptRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderNoDecryptRsp.Unmarshal(m, b)
}
func (m *OrderNoDecryptRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderNoDecryptRsp.Marshal(b, m, deterministic)
}
func (m *OrderNoDecryptRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderNoDecryptRsp.Merge(m, src)
}
func (m *OrderNoDecryptRsp) XXX_Size() int {
	return xxx_messageInfo_OrderNoDecryptRsp.Size(m)
}
func (m *OrderNoDecryptRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderNoDecryptRsp.DiscardUnknown(m)
}

var xxx_messageInfo_OrderNoDecryptRsp proto.InternalMessageInfo

func (m *OrderNoDecryptRsp) GetCode() uint64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *OrderNoDecryptRsp) GetMessage() uint64 {
	if m != nil {
		return m.Message
	}
	return 0
}

func (m *OrderNoDecryptRsp) GetData() *OrderNoEncryptFormat {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderNoGenRequest)(nil), "proto.OrderNoGenRequest")
	proto.RegisterType((*OrderNoGenResponse)(nil), "proto.OrderNoGenResponse")
	proto.RegisterType((*OrderNoEncryptFormat)(nil), "proto.OrderNoEncryptFormat")
	proto.RegisterType((*OrderNoEncryptReq)(nil), "proto.OrderNoEncryptReq")
	proto.RegisterType((*OrderNoEncryptRsp)(nil), "proto.OrderNoEncryptRsp")
	proto.RegisterType((*OrderNoDecryptReq)(nil), "proto.OrderNoDecryptReq")
	proto.RegisterType((*OrderNoDecryptRsp)(nil), "proto.OrderNoDecryptRsp")
}

func init() {
	proto.RegisterFile("orderNo.proto", fileDescriptor_c11a259be2d7b15d)
}

var fileDescriptor_c11a259be2d7b15d = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x3d, 0x4f, 0xf3, 0x30,
	0x10, 0x7e, 0x9d, 0xb8, 0x6f, 0xc5, 0x21, 0x90, 0x30, 0x1d, 0x4c, 0x58, 0x2a, 0x4f, 0x65, 0x29,
	0x52, 0x59, 0x91, 0x58, 0x0a, 0xdd, 0x40, 0xf2, 0xc2, 0x1c, 0x9a, 0x13, 0x8a, 0x50, 0xec, 0xd4,
	0x76, 0x25, 0xf8, 0x89, 0xfc, 0x2b, 0x14, 0xdb, 0x49, 0x4a, 0x3f, 0xa4, 0x32, 0xd9, 0x77, 0xf7,
	0xe4, 0xf9, 0xb8, 0x18, 0xce, 0xb4, 0x29, 0xd0, 0x3c, 0xeb, 0x69, 0x6d, 0xb4, 0xd3, 0x6c, 0xe0,
	0x0f, 0x71, 0x09, 0x17, 0x2f, 0xa1, 0xbf, 0x40, 0x25, 0x71, 0xb5, 0x46, 0xeb, 0x84, 0x04, 0xb6,
	0xd9, 0xb4, 0xb5, 0x56, 0x16, 0x19, 0x03, 0xba, 0xd4, 0x05, 0x72, 0x32, 0x26, 0x13, 0x2a, 0xfd,
	0x9d, 0x71, 0x18, 0x56, 0x68, 0x6d, 0xfe, 0x8e, 0x3c, 0x19, 0x93, 0xc9, 0x89, 0x6c, 0x4b, 0x76,
	0x0e, 0x89, 0xd2, 0x3c, 0xf5, 0xd8, 0x44, 0x69, 0xf1, 0x09, 0xa3, 0xc8, 0xf9, 0xa8, 0x96, 0xe6,
	0xab, 0x76, 0x4f, 0xda, 0x54, 0xb9, 0x8b, 0x38, 0xd2, 0xe2, 0x9a, 0xba, 0x2c, 0x3c, 0x19, 0x95,
	0x49, 0x59, 0x74, 0xaa, 0xa9, 0xa7, 0x0f, 0xaa, 0x0c, 0xa8, 0x2b, 0x2b, 0xe4, 0x34, 0x38, 0x69,
	0xee, 0x8d, 0x93, 0xe6, 0xd4, 0x6b, 0xc7, 0x07, 0xbe, 0xdd, 0x96, 0x62, 0xde, 0x45, 0x8c, 0xca,
	0x12, 0x57, 0xec, 0x16, 0x68, 0x91, 0xbb, 0xdc, 0x0b, 0x9f, 0xce, 0xae, 0xc3, 0x52, 0xa6, 0xfb,
	0x1c, 0x4a, 0x0f, 0x14, 0xaf, 0x3b, 0x2c, 0xb6, 0x3e, 0x66, 0x25, 0xb4, 0x5f, 0xc9, 0x08, 0x06,
	0x4e, 0x7f, 0xa0, 0x8a, 0x59, 0x42, 0x21, 0x6e, 0x3a, 0xe2, 0x39, 0x76, 0xf6, 0x3a, 0x28, 0xd9,
	0x84, 0x9a, 0x1d, 0xe8, 0x9f, 0x3d, 0xb4, 0xb9, 0xd3, 0x23, 0x73, 0xcf, 0xbe, 0x09, 0x0c, 0xe3,
	0x98, 0xdd, 0x43, 0xba, 0x40, 0xc5, 0xf8, 0xef, 0xaf, 0xfa, 0x87, 0x93, 0x5d, 0xed, 0x99, 0x84,
	0xd7, 0x23, 0xfe, 0xb1, 0x07, 0x18, 0x46, 0x81, 0x6d, 0x86, 0xfe, 0xbf, 0x64, 0x07, 0x26, 0xb6,
	0x0e, 0x04, 0x31, 0xf7, 0x36, 0x41, 0xbf, 0xb9, 0xec, 0xc0, 0xa4, 0x21, 0x78, 0xfb, 0xef, 0x47,
	0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0x75, 0x1f, 0x97, 0x0b, 0x03, 0x00, 0x00,
}