// Code generated by protoc-gen-go. DO NOT EDIT.
// source: AdminUndelete.proto

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

// Undelete a file or smart contract that was deleted by AdminDelete - can only be done with a Hedera admin multisig. When it is deleted, it immediately disappears from the system as seen by the user, but is still stored internally until the expiration time, at which time it is truly and permanently deleted. Until that time, it can be undeleted by the Hedera admin multisig. When a smart contract is deleted, the cryptocurrency account within it continues to exist, and is not affected by the expiration time here.
type AdminUndeleteTransactionBody struct {
	// Types that are valid to be assigned to Id:
	//	*AdminUndeleteTransactionBody_FileID
	//	*AdminUndeleteTransactionBody_ContractID
	Id                   isAdminUndeleteTransactionBody_Id `protobuf_oneof:"id"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *AdminUndeleteTransactionBody) Reset()         { *m = AdminUndeleteTransactionBody{} }
func (m *AdminUndeleteTransactionBody) String() string { return proto.CompactTextString(m) }
func (*AdminUndeleteTransactionBody) ProtoMessage()    {}
func (*AdminUndeleteTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c4361be43a376a1, []int{0}
}

func (m *AdminUndeleteTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminUndeleteTransactionBody.Unmarshal(m, b)
}
func (m *AdminUndeleteTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminUndeleteTransactionBody.Marshal(b, m, deterministic)
}
func (m *AdminUndeleteTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminUndeleteTransactionBody.Merge(m, src)
}
func (m *AdminUndeleteTransactionBody) XXX_Size() int {
	return xxx_messageInfo_AdminUndeleteTransactionBody.Size(m)
}
func (m *AdminUndeleteTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminUndeleteTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_AdminUndeleteTransactionBody proto.InternalMessageInfo

type isAdminUndeleteTransactionBody_Id interface {
	isAdminUndeleteTransactionBody_Id()
}

type AdminUndeleteTransactionBody_FileID struct {
	FileID *FileID `protobuf:"bytes,1,opt,name=fileID,proto3,oneof"`
}

type AdminUndeleteTransactionBody_ContractID struct {
	ContractID *ContractID `protobuf:"bytes,2,opt,name=contractID,proto3,oneof"`
}

func (*AdminUndeleteTransactionBody_FileID) isAdminUndeleteTransactionBody_Id() {}

func (*AdminUndeleteTransactionBody_ContractID) isAdminUndeleteTransactionBody_Id() {}

func (m *AdminUndeleteTransactionBody) GetId() isAdminUndeleteTransactionBody_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *AdminUndeleteTransactionBody) GetFileID() *FileID {
	if x, ok := m.GetId().(*AdminUndeleteTransactionBody_FileID); ok {
		return x.FileID
	}
	return nil
}

func (m *AdminUndeleteTransactionBody) GetContractID() *ContractID {
	if x, ok := m.GetId().(*AdminUndeleteTransactionBody_ContractID); ok {
		return x.ContractID
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AdminUndeleteTransactionBody) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AdminUndeleteTransactionBody_OneofMarshaler, _AdminUndeleteTransactionBody_OneofUnmarshaler, _AdminUndeleteTransactionBody_OneofSizer, []interface{}{
		(*AdminUndeleteTransactionBody_FileID)(nil),
		(*AdminUndeleteTransactionBody_ContractID)(nil),
	}
}

func _AdminUndeleteTransactionBody_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AdminUndeleteTransactionBody)
	// id
	switch x := m.Id.(type) {
	case *AdminUndeleteTransactionBody_FileID:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FileID); err != nil {
			return err
		}
	case *AdminUndeleteTransactionBody_ContractID:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContractID); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AdminUndeleteTransactionBody.Id has unexpected type %T", x)
	}
	return nil
}

func _AdminUndeleteTransactionBody_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AdminUndeleteTransactionBody)
	switch tag {
	case 1: // id.fileID
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FileID)
		err := b.DecodeMessage(msg)
		m.Id = &AdminUndeleteTransactionBody_FileID{msg}
		return true, err
	case 2: // id.contractID
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContractID)
		err := b.DecodeMessage(msg)
		m.Id = &AdminUndeleteTransactionBody_ContractID{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AdminUndeleteTransactionBody_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AdminUndeleteTransactionBody)
	// id
	switch x := m.Id.(type) {
	case *AdminUndeleteTransactionBody_FileID:
		s := proto.Size(x.FileID)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdminUndeleteTransactionBody_ContractID:
		s := proto.Size(x.ContractID)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*AdminUndeleteTransactionBody)(nil), "proto.AdminUndeleteTransactionBody")
}

func init() { proto.RegisterFile("AdminUndelete.proto", fileDescriptor_2c4361be43a376a1) }

var fileDescriptor_2c4361be43a376a1 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8c, 0xbf, 0x4a, 0xc6, 0x30,
	0x14, 0xc5, 0xbf, 0x14, 0xed, 0x10, 0x11, 0xb5, 0x2e, 0xa5, 0x38, 0x48, 0x07, 0x75, 0xca, 0x60,
	0x9f, 0xc0, 0x58, 0xa4, 0xdd, 0xa4, 0xd4, 0x07, 0xb8, 0x26, 0x57, 0x13, 0x69, 0xfe, 0x90, 0x04,
	0xa1, 0xab, 0x4f, 0x2e, 0xa4, 0x45, 0xfc, 0xa6, 0xcb, 0x3d, 0xbf, 0xdf, 0x39, 0xf4, 0xfa, 0x49,
	0x1a, 0x6d, 0xdf, 0xac, 0xc4, 0x05, 0x13, 0x32, 0x1f, 0x5c, 0x72, 0xd5, 0x69, 0x3e, 0xcd, 0x25,
	0x87, 0xa8, 0xc5, 0xbc, 0x7a, 0x8c, 0x1b, 0x68, 0x2e, 0x66, 0x6d, 0x30, 0x26, 0x30, 0x7e, 0x0b,
	0xda, 0x1f, 0x42, 0x6f, 0x8e, 0x16, 0xe6, 0x00, 0x36, 0x82, 0x48, 0xda, 0x59, 0xee, 0xe4, 0x5a,
	0xdd, 0xd3, 0xf2, 0x43, 0x2f, 0x38, 0xf6, 0x35, 0xb9, 0x25, 0x0f, 0x67, 0x8f, 0xe7, 0x5b, 0x91,
	0xbd, 0xe4, 0x70, 0x38, 0x4c, 0x3b, 0xae, 0x3a, 0x4a, 0x85, 0xb3, 0x29, 0x80, 0x48, 0x63, 0x5f,
	0x17, 0x59, 0xbe, 0xda, 0xe5, 0xe7, 0x3f, 0x30, 0x1c, 0xa6, 0x7f, 0x1a, 0x3f, 0xa1, 0x85, 0x96,
	0xfc, 0x8e, 0xb6, 0xc2, 0x19, 0xa6, 0x50, 0x62, 0x00, 0x05, 0x51, 0x7d, 0x06, 0xf0, 0x8a, 0x81,
	0xd7, 0x7b, 0xff, 0x0b, 0xbe, 0xe1, 0x95, 0xbc, 0x97, 0xf9, 0xeb, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xeb, 0x2a, 0xe6, 0x43, 0xf4, 0x00, 0x00, 0x00,
}
