// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crypto/encoding/proto/types.proto

package proto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PubKey struct {
	// Types that are valid to be assigned to Key:
	//	*PubKey_Ed25519
	//	*PubKey_Secp256K1
	//	*PubKey_Sr25519
	//	*PubKey_Multisig
	Key                  isPubKey_Key `protobuf_oneof:"key"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PubKey) Reset()         { *m = PubKey{} }
func (m *PubKey) String() string { return proto.CompactTextString(m) }
func (*PubKey) ProtoMessage()    {}
func (*PubKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_32876cb9ea95204c, []int{0}
}
func (m *PubKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubKey.Unmarshal(m, b)
}
func (m *PubKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubKey.Marshal(b, m, deterministic)
}
func (m *PubKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKey.Merge(m, src)
}
func (m *PubKey) XXX_Size() int {
	return xxx_messageInfo_PubKey.Size(m)
}
func (m *PubKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKey.DiscardUnknown(m)
}

var xxx_messageInfo_PubKey proto.InternalMessageInfo

type isPubKey_Key interface {
	isPubKey_Key()
}

type PubKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof" json:"ed25519,omitempty"`
}
type PubKey_Secp256K1 struct {
	Secp256K1 []byte `protobuf:"bytes,2,opt,name=secp256k1,proto3,oneof" json:"secp256k1,omitempty"`
}
type PubKey_Sr25519 struct {
	Sr25519 []byte `protobuf:"bytes,3,opt,name=sr25519,proto3,oneof" json:"sr25519,omitempty"`
}
type PubKey_Multisig struct {
	Multisig *PubKeyMultiSigThreshold `protobuf:"bytes,4,opt,name=multisig,proto3,oneof" json:"multisig,omitempty"`
}

func (*PubKey_Ed25519) isPubKey_Key()   {}
func (*PubKey_Secp256K1) isPubKey_Key() {}
func (*PubKey_Sr25519) isPubKey_Key()   {}
func (*PubKey_Multisig) isPubKey_Key()  {}

func (m *PubKey) GetKey() isPubKey_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PubKey) GetEd25519() []byte {
	if x, ok := m.GetKey().(*PubKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PubKey) GetSecp256K1() []byte {
	if x, ok := m.GetKey().(*PubKey_Secp256K1); ok {
		return x.Secp256K1
	}
	return nil
}

func (m *PubKey) GetSr25519() []byte {
	if x, ok := m.GetKey().(*PubKey_Sr25519); ok {
		return x.Sr25519
	}
	return nil
}

func (m *PubKey) GetMultisig() *PubKeyMultiSigThreshold {
	if x, ok := m.GetKey().(*PubKey_Multisig); ok {
		return x.Multisig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PubKey) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PubKey_Ed25519)(nil),
		(*PubKey_Secp256K1)(nil),
		(*PubKey_Sr25519)(nil),
		(*PubKey_Multisig)(nil),
	}
}

type PrivKey struct {
	// Types that are valid to be assigned to Key:
	//	*PrivKey_Ed25519
	//	*PrivKey_Secp256K1
	//	*PrivKey_Sr25519
	Key                  isPrivKey_Key `protobuf_oneof:"key"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PrivKey) Reset()         { *m = PrivKey{} }
func (m *PrivKey) String() string { return proto.CompactTextString(m) }
func (*PrivKey) ProtoMessage()    {}
func (*PrivKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_32876cb9ea95204c, []int{1}
}
func (m *PrivKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivKey.Unmarshal(m, b)
}
func (m *PrivKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivKey.Marshal(b, m, deterministic)
}
func (m *PrivKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivKey.Merge(m, src)
}
func (m *PrivKey) XXX_Size() int {
	return xxx_messageInfo_PrivKey.Size(m)
}
func (m *PrivKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivKey.DiscardUnknown(m)
}

var xxx_messageInfo_PrivKey proto.InternalMessageInfo

type isPrivKey_Key interface {
	isPrivKey_Key()
}

type PrivKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof" json:"ed25519,omitempty"`
}
type PrivKey_Secp256K1 struct {
	Secp256K1 []byte `protobuf:"bytes,2,opt,name=secp256k1,proto3,oneof" json:"secp256k1,omitempty"`
}
type PrivKey_Sr25519 struct {
	Sr25519 []byte `protobuf:"bytes,3,opt,name=sr25519,proto3,oneof" json:"sr25519,omitempty"`
}

func (*PrivKey_Ed25519) isPrivKey_Key()   {}
func (*PrivKey_Secp256K1) isPrivKey_Key() {}
func (*PrivKey_Sr25519) isPrivKey_Key()   {}

func (m *PrivKey) GetKey() isPrivKey_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PrivKey) GetEd25519() []byte {
	if x, ok := m.GetKey().(*PrivKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PrivKey) GetSecp256K1() []byte {
	if x, ok := m.GetKey().(*PrivKey_Secp256K1); ok {
		return x.Secp256K1
	}
	return nil
}

func (m *PrivKey) GetSr25519() []byte {
	if x, ok := m.GetKey().(*PrivKey_Sr25519); ok {
		return x.Sr25519
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PrivKey) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PrivKey_Ed25519)(nil),
		(*PrivKey_Secp256K1)(nil),
		(*PrivKey_Sr25519)(nil),
	}
}

type PubKeyMultiSigThreshold struct {
	K                    uint64   `protobuf:"varint,1,opt,name=k,proto3" json:"k,omitempty"`
	PubKeys              []PubKey `protobuf:"bytes,2,rep,name=pub_keys,json=pubKeys,proto3" json:"pub_keys"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubKeyMultiSigThreshold) Reset()         { *m = PubKeyMultiSigThreshold{} }
func (m *PubKeyMultiSigThreshold) String() string { return proto.CompactTextString(m) }
func (*PubKeyMultiSigThreshold) ProtoMessage()    {}
func (*PubKeyMultiSigThreshold) Descriptor() ([]byte, []int) {
	return fileDescriptor_32876cb9ea95204c, []int{2}
}
func (m *PubKeyMultiSigThreshold) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubKeyMultiSigThreshold.Unmarshal(m, b)
}
func (m *PubKeyMultiSigThreshold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubKeyMultiSigThreshold.Marshal(b, m, deterministic)
}
func (m *PubKeyMultiSigThreshold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKeyMultiSigThreshold.Merge(m, src)
}
func (m *PubKeyMultiSigThreshold) XXX_Size() int {
	return xxx_messageInfo_PubKeyMultiSigThreshold.Size(m)
}
func (m *PubKeyMultiSigThreshold) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKeyMultiSigThreshold.DiscardUnknown(m)
}

var xxx_messageInfo_PubKeyMultiSigThreshold proto.InternalMessageInfo

func (m *PubKeyMultiSigThreshold) GetK() uint64 {
	if m != nil {
		return m.K
	}
	return 0
}

func (m *PubKeyMultiSigThreshold) GetPubKeys() []PubKey {
	if m != nil {
		return m.PubKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*PubKey)(nil), "tendermint.crypto.encoding.proto.PubKey")
	proto.RegisterType((*PrivKey)(nil), "tendermint.crypto.encoding.proto.PrivKey")
	proto.RegisterType((*PubKeyMultiSigThreshold)(nil), "tendermint.crypto.encoding.proto.PubKeyMultiSigThreshold")
}

func init() { proto.RegisterFile("crypto/encoding/proto/types.proto", fileDescriptor_32876cb9ea95204c) }

var fileDescriptor_32876cb9ea95204c = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xbf, 0x4e, 0xf3, 0x30,
	0x1c, 0xac, 0xdb, 0x7e, 0x4d, 0x3f, 0xb7, 0x93, 0x17, 0xa2, 0x0e, 0x10, 0x3a, 0xa0, 0x4c, 0x8e,
	0x1a, 0x94, 0x8a, 0xae, 0x99, 0x8a, 0x10, 0x52, 0x15, 0x90, 0x90, 0x58, 0x2a, 0x92, 0x58, 0x8e,
	0x95, 0x26, 0xb6, 0x6c, 0x07, 0xc9, 0x6f, 0xc7, 0xc8, 0x53, 0xf0, 0x2c, 0xa8, 0xf9, 0x43, 0x18,
	0x40, 0xb0, 0x30, 0xf9, 0x7c, 0xbe, 0xbb, 0xdf, 0xfd, 0x24, 0xc3, 0xf3, 0x44, 0x1a, 0xa1, 0xb9,
	0x47, 0xca, 0x84, 0xa7, 0xac, 0xa4, 0x9e, 0x90, 0x5c, 0x73, 0x4f, 0x1b, 0x41, 0x14, 0xae, 0x31,
	0x72, 0x34, 0x29, 0x53, 0x22, 0x0b, 0x56, 0x6a, 0xdc, 0xa8, 0x71, 0xa7, 0x6e, 0x14, 0x8b, 0x0b,
	0x9d, 0x31, 0x99, 0xee, 0xc5, 0x93, 0xd4, 0xa6, 0x0d, 0xa0, 0x9c, 0xf2, 0x1e, 0x35, 0xba, 0xe5,
	0x0b, 0x80, 0x93, 0x5d, 0x15, 0xdf, 0x10, 0x83, 0x16, 0xd0, 0x22, 0xa9, 0x1f, 0x04, 0xab, 0x8d,
	0x0d, 0x1c, 0xe0, 0xce, 0xb7, 0x83, 0xa8, 0x23, 0xd0, 0x29, 0xfc, 0xaf, 0x48, 0x22, 0xfc, 0x60,
	0x9d, 0xaf, 0xec, 0x61, 0xfb, 0xda, 0x53, 0x47, 0xaf, 0x92, 0x8d, 0x77, 0xd4, 0x79, 0x5b, 0x02,
	0x3d, 0xc0, 0x69, 0x51, 0x1d, 0x34, 0x53, 0x8c, 0xda, 0x63, 0x07, 0xb8, 0x33, 0x7f, 0x83, 0x7f,
	0xea, 0x8f, 0x9b, 0x4e, 0xb7, 0x47, 0xdf, 0x1d, 0xa3, 0xf7, 0x99, 0x24, 0x2a, 0xe3, 0x87, 0x74,
	0x3b, 0x88, 0x3e, 0xc2, 0xc2, 0x7f, 0x70, 0x94, 0x13, 0xb3, 0xcc, 0xa0, 0xb5, 0x93, 0xec, 0xf9,
	0x0f, 0x57, 0xe8, 0x26, 0x49, 0x78, 0xf2, 0x4d, 0x2f, 0x34, 0x87, 0x20, 0xaf, 0x67, 0x8e, 0x23,
	0x90, 0xa3, 0x6b, 0x38, 0x15, 0x55, 0xbc, 0xcf, 0x89, 0x51, 0xf6, 0xd0, 0x19, 0xb9, 0x33, 0xdf,
	0xfd, 0xed, 0xca, 0xe1, 0xf8, 0xf5, 0xed, 0x6c, 0x10, 0x59, 0xa2, 0xbe, 0xa9, 0xf0, 0xea, 0x71,
	0x4d, 0x99, 0xce, 0xaa, 0x18, 0x27, 0xbc, 0xf0, 0xfa, 0x90, 0xcf, 0xf0, 0xcb, 0x0f, 0x13, 0x4f,
	0xea, 0xe3, 0xf2, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x56, 0xfa, 0x97, 0x5b, 0x50, 0x02, 0x00, 0x00,
}