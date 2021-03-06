// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/justifications/justifications.proto

package justifications

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	lock "token-strike/tsp2p/server/lock"
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

// LockToken the token locking
type LockToken struct {
	// lock — information about lock
	Lock                 *lock.Lock `protobuf:"bytes,1,opt,name=lock,proto3" json:"lock,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LockToken) Reset()         { *m = LockToken{} }
func (m *LockToken) String() string { return proto.CompactTextString(m) }
func (*LockToken) ProtoMessage()    {}
func (*LockToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee571a5de13873a7, []int{0}
}

func (m *LockToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockToken.Unmarshal(m, b)
}
func (m *LockToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockToken.Marshal(b, m, deterministic)
}
func (m *LockToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockToken.Merge(m, src)
}
func (m *LockToken) XXX_Size() int {
	return xxx_messageInfo_LockToken.Size(m)
}
func (m *LockToken) XXX_DiscardUnknown() {
	xxx_messageInfo_LockToken.DiscardUnknown(m)
}

var xxx_messageInfo_LockToken proto.InternalMessageInfo

func (m *LockToken) GetLock() *lock.Lock {
	if m != nil {
		return m.Lock
	}
	return nil
}

// TranferToken receiving funds for tokens and unlcok them
type TranferToken struct {
	// htlc_secret — htlc genereted issuer
	HtlcSecret string `protobuf:"bytes,1,opt,name=htlc_secret,json=htlcSecret,proto3" json:"htlc_secret,omitempty"`
	// lock — hash information about lock
	Lock                 string   `protobuf:"bytes,2,opt,name=lock,proto3" json:"lock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranferToken) Reset()         { *m = TranferToken{} }
func (m *TranferToken) String() string { return proto.CompactTextString(m) }
func (*TranferToken) ProtoMessage()    {}
func (*TranferToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee571a5de13873a7, []int{1}
}

func (m *TranferToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranferToken.Unmarshal(m, b)
}
func (m *TranferToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranferToken.Marshal(b, m, deterministic)
}
func (m *TranferToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranferToken.Merge(m, src)
}
func (m *TranferToken) XXX_Size() int {
	return xxx_messageInfo_TranferToken.Size(m)
}
func (m *TranferToken) XXX_DiscardUnknown() {
	xxx_messageInfo_TranferToken.DiscardUnknown(m)
}

var xxx_messageInfo_TranferToken proto.InternalMessageInfo

func (m *TranferToken) GetHtlcSecret() string {
	if m != nil {
		return m.HtlcSecret
	}
	return ""
}

func (m *TranferToken) GetLock() string {
	if m != nil {
		return m.Lock
	}
	return ""
}

// LockTimeOver timeout for token locking
type LockTimeOver struct {
	// proof_elapsed — PKT block hash confirming expiration lock
	ProofElapsed string `protobuf:"bytes,1,opt,name=proof_elapsed,json=proofElapsed,proto3" json:"proof_elapsed,omitempty"`
	// lock_id — hash with information about lock justification
	Lock                 string   `protobuf:"bytes,2,opt,name=lock,proto3" json:"lock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockTimeOver) Reset()         { *m = LockTimeOver{} }
func (m *LockTimeOver) String() string { return proto.CompactTextString(m) }
func (*LockTimeOver) ProtoMessage()    {}
func (*LockTimeOver) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee571a5de13873a7, []int{2}
}

func (m *LockTimeOver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockTimeOver.Unmarshal(m, b)
}
func (m *LockTimeOver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockTimeOver.Marshal(b, m, deterministic)
}
func (m *LockTimeOver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockTimeOver.Merge(m, src)
}
func (m *LockTimeOver) XXX_Size() int {
	return xxx_messageInfo_LockTimeOver.Size(m)
}
func (m *LockTimeOver) XXX_DiscardUnknown() {
	xxx_messageInfo_LockTimeOver.DiscardUnknown(m)
}

var xxx_messageInfo_LockTimeOver proto.InternalMessageInfo

func (m *LockTimeOver) GetProofElapsed() string {
	if m != nil {
		return m.ProofElapsed
	}
	return ""
}

func (m *LockTimeOver) GetLock() string {
	if m != nil {
		return m.Lock
	}
	return ""
}

// Genesis initial block justification
type Genesis struct {
	// token — token identification by name
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee571a5de13873a7, []int{3}
}

func (m *Genesis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genesis.Unmarshal(m, b)
}
func (m *Genesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genesis.Marshal(b, m, deterministic)
}
func (m *Genesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genesis.Merge(m, src)
}
func (m *Genesis) XXX_Size() int {
	return xxx_messageInfo_Genesis.Size(m)
}
func (m *Genesis) XXX_DiscardUnknown() {
	xxx_messageInfo_Genesis.DiscardUnknown(m)
}

var xxx_messageInfo_Genesis proto.InternalMessageInfo

func (m *Genesis) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*LockToken)(nil), "justifications.LockToken")
	proto.RegisterType((*TranferToken)(nil), "justifications.TranferToken")
	proto.RegisterType((*LockTimeOver)(nil), "justifications.LockTimeOver")
	proto.RegisterType((*Genesis)(nil), "justifications.Genesis")
}

func init() {
	proto.RegisterFile("protos/justifications/justifications.proto", fileDescriptor_ee571a5de13873a7)
}

var fileDescriptor_ee571a5de13873a7 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc7, 0xa9, 0xf8, 0x41, 0xa7, 0xab, 0x87, 0x20, 0x22, 0x1e, 0xac, 0xd4, 0x4b, 0xa9, 0xd8,
	0x85, 0xfa, 0x06, 0x8a, 0xf4, 0x22, 0x08, 0xb5, 0x27, 0x2f, 0x65, 0x8d, 0xb3, 0x18, 0x77, 0xdd,
	0x09, 0x33, 0x63, 0x9f, 0x5f, 0x76, 0x92, 0x8b, 0xa5, 0x97, 0x90, 0xf9, 0xfd, 0x3f, 0x12, 0x06,
	0x66, 0x91, 0x49, 0x49, 0xca, 0xef, 0x5f, 0xd1, 0x50, 0x07, 0x5f, 0x69, 0xa0, 0x6e, 0x77, 0x9c,
	0x9b, 0xc9, 0x9d, 0xfd, 0xa7, 0x57, 0x17, 0x39, 0xdb, 0x92, 0x6f, 0xec, 0x48, 0xbe, 0xc9, 0x1d,
	0x0c, 0x5f, 0xc8, 0x37, 0x6b, 0x6a, 0xb0, 0x73, 0xd7, 0x70, 0xd8, 0x4b, 0x97, 0x83, 0x9b, 0xc1,
	0x74, 0xb4, 0x80, 0xb9, 0xf9, 0x7a, 0x79, 0x65, 0x7c, 0xf2, 0x04, 0xc5, 0x9a, 0xab, 0xae, 0x46,
	0x4e, 0xfe, 0x31, 0x8c, 0xbe, 0xb4, 0xf5, 0x1b, 0x41, 0xcf, 0xa8, 0x16, 0x1b, 0xae, 0xa0, 0x47,
	0x6f, 0x46, 0x9c, 0xcb, 0x85, 0x07, 0xa6, 0xa4, 0x92, 0x25, 0x14, 0xf6, 0x62, 0xf8, 0xc1, 0xd7,
	0x2d, 0xb2, 0xbb, 0x85, 0xd3, 0xc8, 0x44, 0xf5, 0x06, 0xdb, 0x2a, 0x0a, 0x7e, 0xe6, 0x9a, 0xc2,
	0xe0, 0x73, 0x62, 0x7b, 0x8b, 0xc6, 0x70, 0xb2, 0xc4, 0x0e, 0x25, 0x88, 0x3b, 0x87, 0x23, 0xed,
	0x7f, 0x94, 0xb3, 0x69, 0x78, 0x9c, 0xbd, 0x4f, 0xed, 0x72, 0x2f, 0xca, 0xa1, 0xc1, 0x52, 0x25,
	0x2e, 0x62, 0x29, 0xc8, 0x5b, 0xe4, 0x9d, 0xad, 0x7d, 0x1c, 0xdb, 0x3a, 0x1e, 0xfe, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x37, 0xbf, 0x96, 0xfa, 0x64, 0x01, 0x00, 0x00,
}
