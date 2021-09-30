// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/DB/tokendb.proto

package token

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	justifications "token-strike/tsp2p/justifications"
	lock "token-strike/tsp2p/lock"
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

// Block is struct of block in blockchain
type Block struct {
	// prev_block — hash of previous block
	PrevBlock string `protobuf:"bytes,1,opt,name=prev_block,json=prevBlock,proto3" json:"prev_block,omitempty"`
	// justifications — collection of one of justification structure with
	// payload information.
	Justifications []*Justification `protobuf:"bytes,2,rep,name=justifications,proto3" json:"justifications,omitempty"`
	// creation — date of block creation in unix time format
	Creation int64 `protobuf:"varint,3,opt,name=creation,proto3" json:"creation,omitempty"`
	// state — hash of state structure containing locks, owners and meta token
	// info
	State string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	// pkt_block_hash —  the hash of the most recent PKT block
	PktBlockHash string `protobuf:"bytes,5,opt,name=pkt_block_hash,json=pktBlockHash,proto3" json:"pkt_block_hash,omitempty"`
	// pkt_block_height — the height of the most recent PKT block
	PktBlockHeight int32 `protobuf:"varint,6,opt,name=pkt_block_height,json=pktBlockHeight,proto3" json:"pkt_block_height,omitempty"`
	// height — the current height of this TokenStrike chain
	Height uint64 `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
	// signature — issuer ID, needed for validate. If signature incorrect block
	// is not valid
	Signature            string   `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f335e4a205f88b, []int{0}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetPrevBlock() string {
	if m != nil {
		return m.PrevBlock
	}
	return ""
}

func (m *Block) GetJustifications() []*Justification {
	if m != nil {
		return m.Justifications
	}
	return nil
}

func (m *Block) GetCreation() int64 {
	if m != nil {
		return m.Creation
	}
	return 0
}

func (m *Block) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Block) GetPktBlockHash() string {
	if m != nil {
		return m.PktBlockHash
	}
	return ""
}

func (m *Block) GetPktBlockHeight() int32 {
	if m != nil {
		return m.PktBlockHeight
	}
	return 0
}

func (m *Block) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

// Owner contains information about the holders' wallets and their balances
type Owner struct {
	// holder_wallet — hash of wallet address of holder
	HolderWallet string `protobuf:"bytes,1,opt,name=holder_wallet,json=holderWallet,proto3" json:"holder_wallet,omitempty"`
	// count — number of tokens held on wallet
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Owner) Reset()         { *m = Owner{} }
func (m *Owner) String() string { return proto.CompactTextString(m) }
func (*Owner) ProtoMessage()    {}
func (*Owner) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f335e4a205f88b, []int{1}
}

func (m *Owner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Owner.Unmarshal(m, b)
}
func (m *Owner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Owner.Marshal(b, m, deterministic)
}
func (m *Owner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Owner.Merge(m, src)
}
func (m *Owner) XXX_Size() int {
	return xxx_messageInfo_Owner.Size(m)
}
func (m *Owner) XXX_DiscardUnknown() {
	xxx_messageInfo_Owner.DiscardUnknown(m)
}

var xxx_messageInfo_Owner proto.InternalMessageInfo

func (m *Owner) GetHolderWallet() string {
	if m != nil {
		return m.HolderWallet
	}
	return ""
}

func (m *Owner) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// Token contain information about token
type Token struct {
	// count — number of issued tokens;
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// expiration — number of PKT block after which the token expires
	Expiration int32 `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// creation — date of token creation in unix time format
	Creation int64 `protobuf:"varint,3,opt,name=creation,proto3" json:"creation,omitempty"`
	// issuer_pubkey — public key of issuer used for signing
	IssuerPubkey string `protobuf:"bytes,4,opt,name=issuer_pubkey,json=issuerPubkey,proto3" json:"issuer_pubkey,omitempty"`
	// urls — set of urls for access to blockchain
	Urls                 []string `protobuf:"bytes,5,rep,name=urls,proto3" json:"urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f335e4a205f88b, []int{2}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Token) GetExpiration() int32 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *Token) GetCreation() int64 {
	if m != nil {
		return m.Creation
	}
	return 0
}

func (m *Token) GetIssuerPubkey() string {
	if m != nil {
		return m.IssuerPubkey
	}
	return ""
}

func (m *Token) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

// State is a current state of blockchain
type State struct {
	// token — metadata about token
	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// owners — a set of structures with addresses and their balances
	Owners []*Owner `protobuf:"bytes,2,rep,name=owners,proto3" json:"owners,omitempty"`
	// locks — set of lock structures
	Locks                []*lock.Lock `protobuf:"bytes,3,rep,name=locks,proto3" json:"locks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f335e4a205f88b, []int{3}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *State) GetOwners() []*Owner {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *State) GetLocks() []*lock.Lock {
	if m != nil {
		return m.Locks
	}
	return nil
}

// Justification is a helper to use it in block
type Justification struct {
	// Types that are valid to be assigned to Content:
	//	*Justification_Lock
	//	*Justification_Transfer
	//	*Justification_LockOver
	//	*Justification_Genesis
	Content              isJustification_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Justification) Reset()         { *m = Justification{} }
func (m *Justification) String() string { return proto.CompactTextString(m) }
func (*Justification) ProtoMessage()    {}
func (*Justification) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f335e4a205f88b, []int{4}
}

func (m *Justification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Justification.Unmarshal(m, b)
}
func (m *Justification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Justification.Marshal(b, m, deterministic)
}
func (m *Justification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Justification.Merge(m, src)
}
func (m *Justification) XXX_Size() int {
	return xxx_messageInfo_Justification.Size(m)
}
func (m *Justification) XXX_DiscardUnknown() {
	xxx_messageInfo_Justification.DiscardUnknown(m)
}

var xxx_messageInfo_Justification proto.InternalMessageInfo

type isJustification_Content interface {
	isJustification_Content()
}

type Justification_Lock struct {
	Lock *justifications.LockToken `protobuf:"bytes,1,opt,name=lock,proto3,oneof"`
}

type Justification_Transfer struct {
	Transfer *justifications.TranferToken `protobuf:"bytes,2,opt,name=transfer,proto3,oneof"`
}

type Justification_LockOver struct {
	LockOver *justifications.LockTimeOver `protobuf:"bytes,3,opt,name=lock_over,json=lockOver,proto3,oneof"`
}

type Justification_Genesis struct {
	Genesis *justifications.Genesis `protobuf:"bytes,4,opt,name=genesis,proto3,oneof"`
}

func (*Justification_Lock) isJustification_Content() {}

func (*Justification_Transfer) isJustification_Content() {}

func (*Justification_LockOver) isJustification_Content() {}

func (*Justification_Genesis) isJustification_Content() {}

func (m *Justification) GetContent() isJustification_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Justification) GetLock() *justifications.LockToken {
	if x, ok := m.GetContent().(*Justification_Lock); ok {
		return x.Lock
	}
	return nil
}

func (m *Justification) GetTransfer() *justifications.TranferToken {
	if x, ok := m.GetContent().(*Justification_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (m *Justification) GetLockOver() *justifications.LockTimeOver {
	if x, ok := m.GetContent().(*Justification_LockOver); ok {
		return x.LockOver
	}
	return nil
}

func (m *Justification) GetGenesis() *justifications.Genesis {
	if x, ok := m.GetContent().(*Justification_Genesis); ok {
		return x.Genesis
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Justification) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Justification_Lock)(nil),
		(*Justification_Transfer)(nil),
		(*Justification_LockOver)(nil),
		(*Justification_Genesis)(nil),
	}
}

func init() {
	proto.RegisterType((*Block)(nil), "tokendb.Block")
	proto.RegisterType((*Owner)(nil), "tokendb.Owner")
	proto.RegisterType((*Token)(nil), "tokendb.Token")
	proto.RegisterType((*State)(nil), "tokendb.State")
	proto.RegisterType((*Justification)(nil), "tokendb.Justification")
}

func init() { proto.RegisterFile("protos/DB/tokendb.proto", fileDescriptor_87f335e4a205f88b) }

var fileDescriptor_87f335e4a205f88b = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xde, 0xb4, 0x4d, 0xbb, 0x39, 0xfd, 0x41, 0x06, 0xe9, 0xc6, 0xb2, 0x4a, 0xe8, 0x8a, 0x04,
	0xc1, 0x16, 0xb2, 0x77, 0x0a, 0x5e, 0x14, 0xc1, 0x22, 0xc2, 0xca, 0xb8, 0x20, 0x78, 0x53, 0xd2,
	0xec, 0x69, 0x1b, 0x53, 0x33, 0x61, 0x66, 0xd2, 0xea, 0x6b, 0xf8, 0x00, 0xbe, 0xa7, 0x77, 0x32,
	0x67, 0xd2, 0x6c, 0xb7, 0x88, 0x37, 0x25, 0xdf, 0x5f, 0x73, 0xe6, 0x9b, 0x13, 0xb8, 0x28, 0xa4,
	0xd0, 0x42, 0x4d, 0xdf, 0xcd, 0xa6, 0x5a, 0x64, 0x98, 0xdf, 0x2d, 0x27, 0xc4, 0xb0, 0x4e, 0x05,
	0x47, 0x2f, 0x2b, 0xc7, 0xb7, 0x52, 0xe9, 0x74, 0x95, 0x26, 0xb1, 0x4e, 0x45, 0x7e, 0x0a, 0x6d,
	0x68, 0x34, 0xac, 0xbc, 0x5b, 0x91, 0x64, 0xf4, 0x63, 0xf9, 0xf1, 0xef, 0x06, 0xb8, 0x33, 0x83,
	0xd9, 0x53, 0x80, 0x42, 0xe2, 0x6e, 0xb1, 0x34, 0xc8, 0x77, 0x02, 0x27, 0xf4, 0xb8, 0x67, 0x18,
	0x2b, 0xbf, 0x85, 0xc1, 0xc3, 0x3f, 0xf6, 0x1b, 0x41, 0x33, 0xec, 0x46, 0xc3, 0xc9, 0x61, 0xba,
	0x0f, 0xc7, 0x32, 0x3f, 0x71, 0xb3, 0x11, 0x9c, 0x27, 0x12, 0x09, 0xf8, 0xcd, 0xc0, 0x09, 0x9b,
	0xbc, 0xc6, 0xec, 0x31, 0xb8, 0x4a, 0xc7, 0x1a, 0xfd, 0x16, 0xbd, 0xd5, 0x02, 0xf6, 0x1c, 0x06,
	0x45, 0xa6, 0xed, 0x3c, 0x8b, 0x4d, 0xac, 0x36, 0xbe, 0x4b, 0x72, 0xaf, 0xc8, 0x34, 0xcd, 0x34,
	0x8f, 0xd5, 0x86, 0x85, 0xf0, 0xe8, 0xc8, 0x85, 0xe9, 0x7a, 0xa3, 0xfd, 0x76, 0xe0, 0x84, 0x2e,
	0x1f, 0xd4, 0x3e, 0x62, 0xd9, 0x10, 0xda, 0x95, 0xde, 0x09, 0x9c, 0xb0, 0xc5, 0x2b, 0xc4, 0x2e,
	0xc1, 0x53, 0xe9, 0x3a, 0x8f, 0x75, 0x29, 0xd1, 0x3f, 0xb7, 0xe7, 0xae, 0x89, 0xf1, 0x0c, 0xdc,
	0x9b, 0x7d, 0x8e, 0x92, 0x5d, 0x41, 0x7f, 0x23, 0xb6, 0x77, 0x28, 0x17, 0xfb, 0x78, 0xbb, 0x45,
	0x5d, 0x55, 0xd4, 0xb3, 0xe4, 0x17, 0xe2, 0xcc, 0x49, 0x12, 0x51, 0xe6, 0xda, 0x6f, 0xd0, 0x11,
	0x2d, 0x18, 0xff, 0x72, 0xc0, 0xbd, 0x35, 0x2d, 0xdd, 0xeb, 0xce, 0x91, 0xce, 0x9e, 0x01, 0xe0,
	0x8f, 0x22, 0x95, 0xb6, 0x9d, 0x06, 0x4d, 0x7f, 0xc4, 0xfc, 0xb7, 0xbb, 0x2b, 0xe8, 0xa7, 0x4a,
	0x95, 0x28, 0x17, 0x45, 0xb9, 0xcc, 0xf0, 0x67, 0xd5, 0x61, 0xcf, 0x92, 0x9f, 0x88, 0x63, 0x0c,
	0x5a, 0xa5, 0xdc, 0x2a, 0xdf, 0x0d, 0x9a, 0xa1, 0xc7, 0xe9, 0x79, 0xbc, 0x07, 0xf7, 0x73, 0xd5,
	0xb3, 0x4b, 0x57, 0x48, 0x33, 0x75, 0xa3, 0x41, 0x7d, 0xa1, 0x34, 0x32, 0xb7, 0x22, 0x7b, 0x01,
	0x6d, 0x61, 0x7a, 0x38, 0xdc, 0xfb, 0xbd, 0x8d, 0xea, 0xe1, 0x95, 0xca, 0x02, 0x70, 0x4d, 0xe7,
	0xca, 0x6f, 0x92, 0x0d, 0x26, 0xb4, 0x6c, 0x1f, 0x45, 0x92, 0x71, 0x2b, 0x8c, 0xff, 0x38, 0xd0,
	0x7f, 0xb0, 0x2b, 0x6c, 0x0a, 0xad, 0x7a, 0xe9, 0xba, 0xd1, 0x93, 0xc9, 0xc9, 0x06, 0x9b, 0x30,
	0xcd, 0x32, 0x3f, 0xe3, 0x64, 0x64, 0xaf, 0xe1, 0x5c, 0xcb, 0x38, 0x57, 0x2b, 0x94, 0x54, 0x57,
	0x37, 0xba, 0x3c, 0x0d, 0xdd, 0xca, 0x38, 0x5f, 0xa1, 0x3c, 0xe4, 0x6a, 0x3f, 0x7b, 0x03, 0x1e,
	0xed, 0x8a, 0xd8, 0xa1, 0xa4, 0x36, 0xff, 0x11, 0xa6, 0x37, 0xa6, 0xdf, 0xf1, 0x66, 0x87, 0xd2,
	0x84, 0x4d, 0xc0, 0x3c, 0xb3, 0x6b, 0xe8, 0xac, 0x31, 0x47, 0x95, 0x2a, 0xea, 0xb9, 0x1b, 0x5d,
	0x9c, 0x46, 0xdf, 0x5b, 0x79, 0x7e, 0xc6, 0x0f, 0xce, 0x99, 0x07, 0x9d, 0x44, 0xe4, 0x1a, 0x73,
	0x3d, 0x1b, 0x7d, 0xf5, 0xa9, 0xb6, 0x57, 0x4a, 0xcb, 0x34, 0xc3, 0xa9, 0x56, 0x45, 0x54, 0xd8,
	0xef, 0x7b, 0xd9, 0xa6, 0x2f, 0xf2, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x9d, 0x8a,
	0xcc, 0xf9, 0x03, 0x00, 0x00,
}
