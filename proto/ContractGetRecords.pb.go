// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ContractGetRecords.proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Get all the records for a smart contract instance, for any function call (or the constructor call) during the last 24 hours, for which a Record was requested.
type ContractGetRecordsQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ContractID           *ContractID  `protobuf:"bytes,2,opt,name=contractID,proto3" json:"contractID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContractGetRecordsQuery) Reset()         { *m = ContractGetRecordsQuery{} }
func (m *ContractGetRecordsQuery) String() string { return proto.CompactTextString(m) }
func (*ContractGetRecordsQuery) ProtoMessage()    {}
func (*ContractGetRecordsQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6832342edea4bb5, []int{0}
}

func (m *ContractGetRecordsQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractGetRecordsQuery.Unmarshal(m, b)
}
func (m *ContractGetRecordsQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractGetRecordsQuery.Marshal(b, m, deterministic)
}
func (m *ContractGetRecordsQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractGetRecordsQuery.Merge(m, src)
}
func (m *ContractGetRecordsQuery) XXX_Size() int {
	return xxx_messageInfo_ContractGetRecordsQuery.Size(m)
}
func (m *ContractGetRecordsQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractGetRecordsQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ContractGetRecordsQuery proto.InternalMessageInfo

func (m *ContractGetRecordsQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ContractGetRecordsQuery) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

// Response when the client sends the node ContractGetRecordsQuery
type ContractGetRecordsResponse struct {
	Header               *ResponseHeader      `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ContractID           *ContractID          `protobuf:"bytes,2,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Records              []*TransactionRecord `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ContractGetRecordsResponse) Reset()         { *m = ContractGetRecordsResponse{} }
func (m *ContractGetRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*ContractGetRecordsResponse) ProtoMessage()    {}
func (*ContractGetRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6832342edea4bb5, []int{1}
}

func (m *ContractGetRecordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractGetRecordsResponse.Unmarshal(m, b)
}
func (m *ContractGetRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractGetRecordsResponse.Marshal(b, m, deterministic)
}
func (m *ContractGetRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractGetRecordsResponse.Merge(m, src)
}
func (m *ContractGetRecordsResponse) XXX_Size() int {
	return xxx_messageInfo_ContractGetRecordsResponse.Size(m)
}
func (m *ContractGetRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractGetRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContractGetRecordsResponse proto.InternalMessageInfo

func (m *ContractGetRecordsResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ContractGetRecordsResponse) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

func (m *ContractGetRecordsResponse) GetRecords() []*TransactionRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*ContractGetRecordsQuery)(nil), "proto.ContractGetRecordsQuery")
	proto.RegisterType((*ContractGetRecordsResponse)(nil), "proto.ContractGetRecordsResponse")
}

func init() { proto.RegisterFile("ContractGetRecords.proto", fileDescriptor_b6832342edea4bb5) }

var fileDescriptor_b6832342edea4bb5 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0x8b, 0x15, 0xd2, 0x8b, 0x0d, 0x4a, 0xc3, 0x9e, 0xca, 0x1e, 0xa4, 0x08, 0x2e,
	0xb8, 0xbe, 0x41, 0x2b, 0xa8, 0x37, 0x0d, 0x7d, 0x81, 0x31, 0x3b, 0x98, 0x15, 0x4c, 0xc2, 0x24,
	0x8a, 0x7d, 0x2d, 0x9f, 0x50, 0xc8, 0xce, 0xca, 0xea, 0x1e, 0x3d, 0x85, 0xcc, 0xff, 0xfd, 0xff,
	0x3f, 0x89, 0x50, 0x3b, 0xef, 0x12, 0x81, 0x49, 0x77, 0x98, 0x34, 0x1a, 0x4f, 0x6d, 0xac, 0x03,
	0xf9, 0xe4, 0xe5, 0x71, 0x3e, 0xca, 0xd3, 0x2d, 0xc4, 0xce, 0xec, 0x0f, 0x01, 0x59, 0x28, 0x57,
	0x7b, 0x02, 0x17, 0xc1, 0xa4, 0xce, 0xbb, 0xde, 0xc2, 0xc2, 0xf2, 0xe9, 0x1d, 0xe9, 0x70, 0x8f,
	0xd0, 0x22, 0xf1, 0xe8, 0x4c, 0x63, 0x0c, 0xde, 0x45, 0x1c, 0x4f, 0xab, 0x4f, 0xb1, 0x9a, 0xd6,
	0x66, 0xb3, 0xbc, 0x14, 0x73, 0x9b, 0x51, 0x55, 0xac, 0x8b, 0xcd, 0xa2, 0x91, 0xbd, 0xa5, 0x1e,
	0x45, 0x6b, 0x26, 0xe4, 0xb5, 0x10, 0x86, 0x63, 0x1e, 0x6e, 0xd5, 0x51, 0xe6, 0x97, 0xcc, 0xef,
	0x7e, 0x04, 0x3d, 0x82, 0xaa, 0xaf, 0x42, 0x94, 0xd3, 0xea, 0x61, 0x49, 0x79, 0xf5, 0xa7, 0xfd,
	0x9c, 0xd3, 0x7e, 0xbf, 0xe2, 0x1f, 0x0b, 0xc8, 0x46, 0x9c, 0x50, 0x5f, 0xaa, 0x66, 0xeb, 0xd9,
	0x66, 0xd1, 0x28, 0xe6, 0x27, 0x9f, 0xaa, 0x07, 0x70, 0x7b, 0x21, 0x2a, 0xe3, 0xdf, 0x6a, 0x8b,
	0x2d, 0x12, 0x58, 0x88, 0xf6, 0x85, 0x20, 0xd8, 0x1a, 0x42, 0xc7, 0xde, 0x57, 0xf8, 0x80, 0xc7,
	0xe2, 0x79, 0x9e, 0x6f, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x20, 0x2f, 0xa3, 0xad, 0xd4,
	0x01, 0x00, 0x00,
}
