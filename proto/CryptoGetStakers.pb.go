// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CryptoGetStakers.proto

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

// Get all the accounts that are proxy staking to this account. For each of them, give the amount currently staked. This is not yet implemented, but will be in a future version of the API.
type CryptoGetStakersQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	AccountID            *AccountID   `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CryptoGetStakersQuery) Reset()         { *m = CryptoGetStakersQuery{} }
func (m *CryptoGetStakersQuery) String() string { return proto.CompactTextString(m) }
func (*CryptoGetStakersQuery) ProtoMessage()    {}
func (*CryptoGetStakersQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af165a1b4c34d1f, []int{0}
}

func (m *CryptoGetStakersQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetStakersQuery.Unmarshal(m, b)
}
func (m *CryptoGetStakersQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetStakersQuery.Marshal(b, m, deterministic)
}
func (m *CryptoGetStakersQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetStakersQuery.Merge(m, src)
}
func (m *CryptoGetStakersQuery) XXX_Size() int {
	return xxx_messageInfo_CryptoGetStakersQuery.Size(m)
}
func (m *CryptoGetStakersQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetStakersQuery.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetStakersQuery proto.InternalMessageInfo

func (m *CryptoGetStakersQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CryptoGetStakersQuery) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

// information about a single account that is proxy staking
type ProxyStaker struct {
	AccountID            *AccountID `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Amount               int64      `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProxyStaker) Reset()         { *m = ProxyStaker{} }
func (m *ProxyStaker) String() string { return proto.CompactTextString(m) }
func (*ProxyStaker) ProtoMessage()    {}
func (*ProxyStaker) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af165a1b4c34d1f, []int{1}
}

func (m *ProxyStaker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyStaker.Unmarshal(m, b)
}
func (m *ProxyStaker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyStaker.Marshal(b, m, deterministic)
}
func (m *ProxyStaker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyStaker.Merge(m, src)
}
func (m *ProxyStaker) XXX_Size() int {
	return xxx_messageInfo_ProxyStaker.Size(m)
}
func (m *ProxyStaker) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyStaker.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyStaker proto.InternalMessageInfo

func (m *ProxyStaker) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *ProxyStaker) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// all of the accounts proxy staking to a given account, and the amounts proxy staked
type AllProxyStakers struct {
	AccountID            *AccountID     `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	ProxyStaker          []*ProxyStaker `protobuf:"bytes,2,rep,name=proxyStaker,proto3" json:"proxyStaker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AllProxyStakers) Reset()         { *m = AllProxyStakers{} }
func (m *AllProxyStakers) String() string { return proto.CompactTextString(m) }
func (*AllProxyStakers) ProtoMessage()    {}
func (*AllProxyStakers) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af165a1b4c34d1f, []int{2}
}

func (m *AllProxyStakers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllProxyStakers.Unmarshal(m, b)
}
func (m *AllProxyStakers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllProxyStakers.Marshal(b, m, deterministic)
}
func (m *AllProxyStakers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllProxyStakers.Merge(m, src)
}
func (m *AllProxyStakers) XXX_Size() int {
	return xxx_messageInfo_AllProxyStakers.Size(m)
}
func (m *AllProxyStakers) XXX_DiscardUnknown() {
	xxx_messageInfo_AllProxyStakers.DiscardUnknown(m)
}

var xxx_messageInfo_AllProxyStakers proto.InternalMessageInfo

func (m *AllProxyStakers) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *AllProxyStakers) GetProxyStaker() []*ProxyStaker {
	if m != nil {
		return m.ProxyStaker
	}
	return nil
}

// Response when the client sends the node CryptoGetStakersQuery
type CryptoGetStakersResponse struct {
	Header               *ResponseHeader  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Stakers              *AllProxyStakers `protobuf:"bytes,3,opt,name=stakers,proto3" json:"stakers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CryptoGetStakersResponse) Reset()         { *m = CryptoGetStakersResponse{} }
func (m *CryptoGetStakersResponse) String() string { return proto.CompactTextString(m) }
func (*CryptoGetStakersResponse) ProtoMessage()    {}
func (*CryptoGetStakersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af165a1b4c34d1f, []int{3}
}

func (m *CryptoGetStakersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetStakersResponse.Unmarshal(m, b)
}
func (m *CryptoGetStakersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetStakersResponse.Marshal(b, m, deterministic)
}
func (m *CryptoGetStakersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetStakersResponse.Merge(m, src)
}
func (m *CryptoGetStakersResponse) XXX_Size() int {
	return xxx_messageInfo_CryptoGetStakersResponse.Size(m)
}
func (m *CryptoGetStakersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetStakersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetStakersResponse proto.InternalMessageInfo

func (m *CryptoGetStakersResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CryptoGetStakersResponse) GetStakers() *AllProxyStakers {
	if m != nil {
		return m.Stakers
	}
	return nil
}

func init() {
	proto.RegisterType((*CryptoGetStakersQuery)(nil), "proto.CryptoGetStakersQuery")
	proto.RegisterType((*ProxyStaker)(nil), "proto.ProxyStaker")
	proto.RegisterType((*AllProxyStakers)(nil), "proto.AllProxyStakers")
	proto.RegisterType((*CryptoGetStakersResponse)(nil), "proto.CryptoGetStakersResponse")
}

func init() { proto.RegisterFile("CryptoGetStakers.proto", fileDescriptor_4af165a1b4c34d1f) }

var fileDescriptor_4af165a1b4c34d1f = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0xe9, 0x86, 0x15, 0xaf, 0x0f, 0xce, 0xe0, 0x4a, 0xd9, 0xd3, 0xe8, 0x83, 0x88, 0x60,
	0x11, 0xf5, 0x0b, 0x6c, 0x0a, 0xea, 0xdb, 0x8c, 0xfa, 0x01, 0xce, 0xee, 0x30, 0xd3, 0x6d, 0x09,
	0x49, 0xa6, 0x16, 0xbf, 0xbc, 0x98, 0xa6, 0x2e, 0xed, 0x8b, 0xf8, 0xd4, 0xde, 0xe5, 0xf7, 0xff,
	0x5f, 0xee, 0x1f, 0x48, 0xaf, 0x74, 0xa5, 0xac, 0xbc, 0x21, 0xfb, 0x60, 0xf1, 0x8d, 0xb4, 0x29,
	0x94, 0x96, 0x56, 0xb2, 0x1d, 0xf7, 0x19, 0x0d, 0xa6, 0x68, 0x16, 0xe5, 0x63, 0xa5, 0xc8, 0x1f,
	0x8c, 0x0e, 0xee, 0x37, 0xa4, 0xab, 0x5b, 0xc2, 0x39, 0x69, 0xdf, 0x3a, 0xe4, 0x64, 0x94, 0x5c,
	0x1b, 0x0a, 0xbb, 0xb9, 0x81, 0x61, 0xd7, 0xdb, 0x49, 0xd9, 0x09, 0xc4, 0xc2, 0x81, 0x59, 0x34,
	0x8e, 0x8e, 0x93, 0x73, 0x56, 0x0b, 0x8a, 0xc0, 0x98, 0x7b, 0x82, 0x15, 0xb0, 0x87, 0x65, 0x29,
	0x37, 0x6b, 0x7b, 0x77, 0x9d, 0xf5, 0x1c, 0x3e, 0xf0, 0xf8, 0xa4, 0xe9, 0xf3, 0x2d, 0x92, 0x3f,
	0x41, 0x32, 0xd3, 0xf2, 0xb3, 0xaa, 0x07, 0xb6, 0xe5, 0xd1, 0x9f, 0x72, 0x96, 0x42, 0x8c, 0xab,
	0x9f, 0x7f, 0x37, 0xab, 0xcf, 0x7d, 0x95, 0x7f, 0xc0, 0xfe, 0x64, 0xb9, 0x0c, 0x9c, 0xcd, 0xbf,
	0xad, 0x2f, 0x21, 0x51, 0x5b, 0x7d, 0xd6, 0x1b, 0xf7, 0x83, 0xd5, 0x03, 0x67, 0x1e, 0x62, 0xf9,
	0x17, 0x64, 0xdd, 0x10, 0x9b, 0xb0, 0xd9, 0x69, 0x27, 0xc7, 0xa1, 0x37, 0x6b, 0xbf, 0xc6, 0x6f,
	0x94, 0x67, 0xb0, 0x6b, 0x6a, 0x87, 0xac, 0xef, 0xf8, 0xb4, 0xb9, 0x6e, 0x7b, 0x33, 0xde, 0x60,
	0xd3, 0x23, 0xc8, 0x4b, 0xb9, 0x2a, 0x04, 0xcd, 0x49, 0xa3, 0x40, 0x23, 0x5e, 0x34, 0x2a, 0x51,
	0xa0, 0x5a, 0x78, 0xe5, 0x2b, 0xbe, 0xe3, 0x2c, 0x7a, 0x8e, 0x5d, 0x75, 0xf1, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0xd9, 0x28, 0xcc, 0x52, 0x4c, 0x02, 0x00, 0x00,
}