// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/lock/lock.proto

package lock

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

// Lock contain information about tokens and contract for transferring
type Lock struct {
	// count — number of sending tokens
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// recipient — wallet addres of new owner of tokens
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// sender — owner of the wallet address to which tokens will be returned
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	// htlc_secret_hash — hash of contract
	HtlcSecretHash string `protobuf:"bytes,4,opt,name=htlc_secret_hash,json=htlcSecretHash,proto3" json:"htlc_secret_hash,omitempty"`
	// proof_count — lock expiration time in PKT blocks
	ProofCount int32 `protobuf:"varint,5,opt,name=proof_count,json=proofCount,proto3" json:"proof_count,omitempty"`
	// creation_height — creation height in token blockchain
	CreationHeight uint64 `protobuf:"varint,6,opt,name=creation_height,json=creationHeight,proto3" json:"creation_height,omitempty"`
	// signature generated with old owner private key
	Signature            string   `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lock) Reset()         { *m = Lock{} }
func (m *Lock) String() string { return proto.CompactTextString(m) }
func (*Lock) ProtoMessage()    {}
func (*Lock) Descriptor() ([]byte, []int) {
	return fileDescriptor_d340bab4d79d59c9, []int{0}
}

func (m *Lock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lock.Unmarshal(m, b)
}
func (m *Lock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lock.Marshal(b, m, deterministic)
}
func (m *Lock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lock.Merge(m, src)
}
func (m *Lock) XXX_Size() int {
	return xxx_messageInfo_Lock.Size(m)
}
func (m *Lock) XXX_DiscardUnknown() {
	xxx_messageInfo_Lock.DiscardUnknown(m)
}

var xxx_messageInfo_Lock proto.InternalMessageInfo

func (m *Lock) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Lock) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *Lock) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Lock) GetHtlcSecretHash() string {
	if m != nil {
		return m.HtlcSecretHash
	}
	return ""
}

func (m *Lock) GetProofCount() int32 {
	if m != nil {
		return m.ProofCount
	}
	return 0
}

func (m *Lock) GetCreationHeight() uint64 {
	if m != nil {
		return m.CreationHeight
	}
	return 0
}

func (m *Lock) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*Lock)(nil), "lock.Lock")
}

func init() { proto.RegisterFile("protos/lock/lock.proto", fileDescriptor_d340bab4d79d59c9) }

var fileDescriptor_d340bab4d79d59c9 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd0, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0x06, 0x70, 0xe2, 0xb6, 0x95, 0x8d, 0xb0, 0x4a, 0x90, 0x25, 0x07, 0xd1, 0xe0, 0xc5, 0x5c,
	0xb4, 0xa0, 0x6f, 0xa0, 0x97, 0x3d, 0x78, 0xaa, 0x37, 0x2f, 0xa5, 0xc6, 0x71, 0x13, 0xba, 0x24,
	0x61, 0x32, 0xf5, 0x89, 0x7d, 0x10, 0xe9, 0xd4, 0x3f, 0x97, 0x90, 0xf9, 0x25, 0x30, 0x1f, 0x9f,
	0xdc, 0x66, 0x4c, 0x94, 0x4a, 0x7b, 0x48, 0x6e, 0xe4, 0xe3, 0x8e, 0x41, 0x55, 0xf3, 0xfd, 0xfa,
	0x4b, 0xc8, 0xea, 0x39, 0xb9, 0x51, 0x9d, 0xcb, 0xda, 0xa5, 0x29, 0x92, 0x16, 0x46, 0xd8, 0x55,
	0xb7, 0x0c, 0xea, 0x42, 0xae, 0x11, 0x5c, 0xc8, 0x01, 0x22, 0xe9, 0x23, 0x23, 0xec, 0xba, 0xfb,
	0x07, 0xb5, 0x95, 0x4d, 0x81, 0xf8, 0x0e, 0xa8, 0x57, 0xfc, 0xf4, 0x33, 0x29, 0x2b, 0xcf, 0x3c,
	0x1d, 0x5c, 0x5f, 0xc0, 0x21, 0x50, 0xef, 0x87, 0xe2, 0x75, 0xc5, 0x3f, 0x36, 0xb3, 0xbf, 0x30,
	0xef, 0x86, 0xe2, 0xd5, 0x95, 0x3c, 0xc9, 0x98, 0xd2, 0x47, 0xbf, 0xec, 0xae, 0x8d, 0xb0, 0x75,
	0x27, 0x99, 0x9e, 0x38, 0xc0, 0x8d, 0x3c, 0x75, 0x08, 0x03, 0x85, 0x14, 0x7b, 0x0f, 0x61, 0xef,
	0x49, 0x37, 0x46, 0xd8, 0xaa, 0xdb, 0xfc, 0xf2, 0x8e, 0x75, 0x4e, 0x5a, 0xc2, 0x3e, 0x0e, 0x34,
	0x21, 0xe8, 0xe3, 0x25, 0xe9, 0x1f, 0x3c, 0x9a, 0xd7, 0x4b, 0x4a, 0x23, 0xc4, 0xdb, 0x42, 0x18,
	0x46, 0x68, 0xa9, 0xe4, 0xfb, 0xdc, 0x16, 0xc0, 0x4f, 0x40, 0x2e, 0xe5, 0xad, 0xe1, 0x56, 0x1e,
	0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x41, 0x86, 0x92, 0x9c, 0x2f, 0x01, 0x00, 0x00,
}
