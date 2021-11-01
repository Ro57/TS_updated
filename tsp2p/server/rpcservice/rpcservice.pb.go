// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/PRCSevice/rpcservice.proto

package rpcservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	DB "token-strike/tsp2p/server/DB"
	tokenstrike "token-strike/tsp2p/server/tokenstrike"
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

// TransferTokensRequest contain information about sending transaction
type TransferTokensRequest struct {
	// token — token name
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// lock — hash of locked tokens
	Lock []byte `protobuf:"bytes,2,opt,name=lock,proto3" json:"lock,omitempty"`
	// htlc — funds transfer contract generated in lightning network
	Htlc                 []byte   `protobuf:"bytes,3,opt,name=htlc,proto3" json:"htlc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferTokensRequest) Reset()         { *m = TransferTokensRequest{} }
func (m *TransferTokensRequest) String() string { return proto.CompactTextString(m) }
func (*TransferTokensRequest) ProtoMessage()    {}
func (*TransferTokensRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c2ffff63f68be3, []int{0}
}

func (m *TransferTokensRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferTokensRequest.Unmarshal(m, b)
}
func (m *TransferTokensRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferTokensRequest.Marshal(b, m, deterministic)
}
func (m *TransferTokensRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferTokensRequest.Merge(m, src)
}
func (m *TransferTokensRequest) XXX_Size() int {
	return xxx_messageInfo_TransferTokensRequest.Size(m)
}
func (m *TransferTokensRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferTokensRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferTokensRequest proto.InternalMessageInfo

func (m *TransferTokensRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *TransferTokensRequest) GetLock() []byte {
	if m != nil {
		return m.Lock
	}
	return nil
}

func (m *TransferTokensRequest) GetHtlc() []byte {
	if m != nil {
		return m.Htlc
	}
	return nil
}

// TransferTokensResponse — contain transaction id
type TransferTokensResponse struct {
	// txid — hash of transaction information
	Txid                 []byte   `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferTokensResponse) Reset()         { *m = TransferTokensResponse{} }
func (m *TransferTokensResponse) String() string { return proto.CompactTextString(m) }
func (*TransferTokensResponse) ProtoMessage()    {}
func (*TransferTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c2ffff63f68be3, []int{1}
}

func (m *TransferTokensResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferTokensResponse.Unmarshal(m, b)
}
func (m *TransferTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferTokensResponse.Marshal(b, m, deterministic)
}
func (m *TransferTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferTokensResponse.Merge(m, src)
}
func (m *TransferTokensResponse) XXX_Size() int {
	return xxx_messageInfo_TransferTokensResponse.Size(m)
}
func (m *TransferTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferTokensResponse proto.InternalMessageInfo

func (m *TransferTokensResponse) GetTxid() []byte {
	if m != nil {
		return m.Txid
	}
	return nil
}

// IssueTokenRequest contains a list of owners and the number of pkt blocks
// before the token expires
type IssueTokenRequest struct {
	// owners - contains information about all token holders,
	// the number of tokens is the sum of all tokens of the owners
	Owners []*DB.Owner `protobuf:"bytes,1,rep,name=owners,proto3" json:"owners,omitempty"`
	// expiration — number of PKT block after which the token expires
	Expiration           int32    `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueTokenRequest) Reset()         { *m = IssueTokenRequest{} }
func (m *IssueTokenRequest) String() string { return proto.CompactTextString(m) }
func (*IssueTokenRequest) ProtoMessage()    {}
func (*IssueTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c2ffff63f68be3, []int{2}
}

func (m *IssueTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueTokenRequest.Unmarshal(m, b)
}
func (m *IssueTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueTokenRequest.Marshal(b, m, deterministic)
}
func (m *IssueTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueTokenRequest.Merge(m, src)
}
func (m *IssueTokenRequest) XXX_Size() int {
	return xxx_messageInfo_IssueTokenRequest.Size(m)
}
func (m *IssueTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueTokenRequest proto.InternalMessageInfo

func (m *IssueTokenRequest) GetOwners() []*DB.Owner {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *IssueTokenRequest) GetExpiration() int32 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

// IssueTokenResponse contain token id for access to it
type IssueTokenResponse struct {
	// token_id — hash of token struct
	TokenId              string   `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueTokenResponse) Reset()         { *m = IssueTokenResponse{} }
func (m *IssueTokenResponse) String() string { return proto.CompactTextString(m) }
func (*IssueTokenResponse) ProtoMessage()    {}
func (*IssueTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c2ffff63f68be3, []int{3}
}

func (m *IssueTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueTokenResponse.Unmarshal(m, b)
}
func (m *IssueTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueTokenResponse.Marshal(b, m, deterministic)
}
func (m *IssueTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueTokenResponse.Merge(m, src)
}
func (m *IssueTokenResponse) XXX_Size() int {
	return xxx_messageInfo_IssueTokenResponse.Size(m)
}
func (m *IssueTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IssueTokenResponse proto.InternalMessageInfo

func (m *IssueTokenResponse) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

// LockTokenRequest send information about token
type LockTokenRequest struct {
	// token_id — hash of token struct
	TokenId string `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// amount of locked token
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// recipient — token holder address
	Recipient string `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// secret_hash — hash of htlc
	SecretHash           string   `protobuf:"bytes,4,opt,name=secret_hash,json=secretHash,proto3" json:"secret_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockTokenRequest) Reset()         { *m = LockTokenRequest{} }
func (m *LockTokenRequest) String() string { return proto.CompactTextString(m) }
func (*LockTokenRequest) ProtoMessage()    {}
func (*LockTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c2ffff63f68be3, []int{4}
}

func (m *LockTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockTokenRequest.Unmarshal(m, b)
}
func (m *LockTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockTokenRequest.Marshal(b, m, deterministic)
}
func (m *LockTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockTokenRequest.Merge(m, src)
}
func (m *LockTokenRequest) XXX_Size() int {
	return xxx_messageInfo_LockTokenRequest.Size(m)
}
func (m *LockTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LockTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LockTokenRequest proto.InternalMessageInfo

func (m *LockTokenRequest) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *LockTokenRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *LockTokenRequest) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *LockTokenRequest) GetSecretHash() string {
	if m != nil {
		return m.SecretHash
	}
	return ""
}

// LockTokenResponse response with hash of lock
type LockTokenResponse struct {
	// lock_id — hash of lock
	LockId               []byte   `protobuf:"bytes,1,opt,name=lock_id,json=lockId,proto3" json:"lock_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockTokenResponse) Reset()         { *m = LockTokenResponse{} }
func (m *LockTokenResponse) String() string { return proto.CompactTextString(m) }
func (*LockTokenResponse) ProtoMessage()    {}
func (*LockTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c2ffff63f68be3, []int{5}
}

func (m *LockTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockTokenResponse.Unmarshal(m, b)
}
func (m *LockTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockTokenResponse.Marshal(b, m, deterministic)
}
func (m *LockTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockTokenResponse.Merge(m, src)
}
func (m *LockTokenResponse) XXX_Size() int {
	return xxx_messageInfo_LockTokenResponse.Size(m)
}
func (m *LockTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LockTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LockTokenResponse proto.InternalMessageInfo

func (m *LockTokenResponse) GetLockId() []byte {
	if m != nil {
		return m.LockId
	}
	return nil
}

// PeerRequest contain url with host
type PeerRequest struct {
	// url for rpc connection
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerRequest) Reset()         { *m = PeerRequest{} }
func (m *PeerRequest) String() string { return proto.CompactTextString(m) }
func (*PeerRequest) ProtoMessage()    {}
func (*PeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c2ffff63f68be3, []int{6}
}

func (m *PeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerRequest.Unmarshal(m, b)
}
func (m *PeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerRequest.Marshal(b, m, deterministic)
}
func (m *PeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerRequest.Merge(m, src)
}
func (m *PeerRequest) XXX_Size() int {
	return xxx_messageInfo_PeerRequest.Size(m)
}
func (m *PeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeerRequest proto.InternalMessageInfo

func (m *PeerRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*TransferTokensRequest)(nil), "rpcservice.TransferTokensRequest")
	proto.RegisterType((*TransferTokensResponse)(nil), "rpcservice.TransferTokensResponse")
	proto.RegisterType((*IssueTokenRequest)(nil), "rpcservice.IssueTokenRequest")
	proto.RegisterType((*IssueTokenResponse)(nil), "rpcservice.IssueTokenResponse")
	proto.RegisterType((*LockTokenRequest)(nil), "rpcservice.LockTokenRequest")
	proto.RegisterType((*LockTokenResponse)(nil), "rpcservice.LockTokenResponse")
	proto.RegisterType((*PeerRequest)(nil), "rpcservice.PeerRequest")
}

func init() { proto.RegisterFile("protos/PRCSevice/rpcservice.proto", fileDescriptor_30c2ffff63f68be3) }

var fileDescriptor_30c2ffff63f68be3 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4f, 0x4f, 0xdb, 0x4e,
	0x10, 0x55, 0x7e, 0x81, 0x80, 0x07, 0x84, 0x60, 0x7f, 0x94, 0x18, 0xf3, 0x2f, 0x58, 0x08, 0xe5,
	0x40, 0xed, 0x2a, 0xbd, 0x54, 0xbd, 0x15, 0xa8, 0x20, 0x6a, 0xa5, 0x46, 0x1b, 0x7a, 0x69, 0x0f,
	0x91, 0x63, 0x0f, 0xc4, 0x4a, 0xf0, 0x9a, 0xdd, 0x75, 0x4a, 0xef, 0xfd, 0xb0, 0xfd, 0x18, 0xd5,
	0xae, 0xd7, 0x64, 0x03, 0x4d, 0x6f, 0xb3, 0xef, 0x3d, 0x3f, 0xbf, 0x19, 0xef, 0x18, 0x8e, 0x73,
	0xce, 0x24, 0x13, 0x61, 0x8f, 0x5e, 0xf4, 0x71, 0x9a, 0xc6, 0x18, 0xf2, 0x3c, 0x16, 0xc8, 0x55,
	0x19, 0x68, 0x8e, 0xc0, 0x0c, 0xf1, 0xf6, 0xee, 0x18, 0xbb, 0x9b, 0x60, 0xa8, 0x99, 0x61, 0x71,
	0x1b, 0xe2, 0x7d, 0x2e, 0x7f, 0x96, 0x42, 0xaf, 0x69, 0xbc, 0x2e, 0xcf, 0x43, 0xc9, 0xc6, 0x98,
	0x25, 0x43, 0x43, 0x9c, 0x18, 0x42, 0xa3, 0x42, 0xf2, 0x74, 0x8c, 0x76, 0x5d, 0xaa, 0xfc, 0xaf,
	0xf0, 0xea, 0x86, 0x47, 0x99, 0xb8, 0x45, 0x7e, 0xa3, 0x49, 0x8a, 0x0f, 0x05, 0x0a, 0x49, 0xb6,
	0x61, 0x59, 0xab, 0xdd, 0x5a, 0xab, 0xd6, 0x76, 0x68, 0x79, 0x20, 0x04, 0x96, 0x26, 0x2c, 0x1e,
	0xbb, 0xff, 0xb5, 0x6a, 0xed, 0x75, 0xaa, 0x6b, 0x85, 0x8d, 0xe4, 0x24, 0x76, 0xeb, 0x25, 0xa6,
	0x6a, 0xff, 0x0c, 0x76, 0x9e, 0xdb, 0x8a, 0x9c, 0x65, 0x02, 0x95, 0x5a, 0x3e, 0xa6, 0x89, 0xb6,
	0x5d, 0xa7, 0xba, 0xf6, 0xbf, 0xc3, 0x56, 0x57, 0x88, 0x02, 0xb5, 0xb4, 0x0a, 0x70, 0x0a, 0x0d,
	0xf6, 0x23, 0x43, 0x2e, 0xdc, 0x5a, 0xab, 0xde, 0x5e, 0xeb, 0x6c, 0x04, 0x55, 0x7f, 0x5f, 0x14,
	0x4c, 0x0d, 0x4b, 0x0e, 0x01, 0xf0, 0x31, 0x4f, 0x79, 0x24, 0x53, 0x96, 0xe9, 0x60, 0xcb, 0xd4,
	0x42, 0xfc, 0x10, 0x88, 0x6d, 0x6e, 0x62, 0xec, 0xc2, 0xaa, 0xb6, 0x1b, 0x98, 0x28, 0x0e, 0x5d,
	0xd1, 0xe7, 0x6e, 0xe2, 0xff, 0xaa, 0xc1, 0xe6, 0x67, 0x16, 0x8f, 0xe7, 0xd2, 0x2c, 0xd6, 0x93,
	0x1d, 0x68, 0x44, 0xf7, 0xac, 0xc8, 0xa4, 0x7e, 0xf9, 0x12, 0x35, 0x27, 0xb2, 0x0f, 0x0e, 0xc7,
	0x38, 0xcd, 0x53, 0xcc, 0xa4, 0x1e, 0x8e, 0x43, 0x67, 0x00, 0x39, 0x82, 0x35, 0x81, 0x31, 0x47,
	0x39, 0x18, 0x45, 0x62, 0xe4, 0x2e, 0x69, 0x1e, 0x4a, 0xe8, 0x3a, 0x12, 0x23, 0xff, 0x0c, 0xb6,
	0xac, 0x14, 0x26, 0x76, 0x13, 0x56, 0xd4, 0xcc, 0x07, 0x4f, 0x03, 0x6c, 0xa8, 0x63, 0x37, 0xf1,
	0x8f, 0x60, 0xad, 0x87, 0xc8, 0xab, 0xb8, 0x9b, 0x50, 0x2f, 0xf8, 0xc4, 0x24, 0x55, 0x65, 0xe7,
	0x77, 0x1d, 0x80, 0xf6, 0x2e, 0xfa, 0xe5, 0x9d, 0x22, 0x14, 0x9c, 0x3e, 0x66, 0x89, 0x76, 0x27,
	0xc7, 0x81, 0x75, 0xff, 0xfe, 0x7a, 0x1d, 0x3c, 0xff, 0x5f, 0x12, 0x13, 0xee, 0x13, 0xc0, 0x6c,
	0xd2, 0xe4, 0xc0, 0x7e, 0xe2, 0xc5, 0xe7, 0xf5, 0x0e, 0x17, 0xd1, 0xc6, 0xec, 0x1a, 0x9c, 0xa7,
	0xf6, 0xc9, 0xbe, 0x2d, 0x7e, 0xfe, 0x6d, 0xbc, 0x83, 0x05, 0xac, 0x71, 0x7a, 0x0f, 0x2b, 0x1f,
	0x92, 0x44, 0x4d, 0x87, 0x34, 0x6d, 0xa5, 0x35, 0x2f, 0x6f, 0x27, 0x28, 0x77, 0x2c, 0xa8, 0x76,
	0x2c, 0xf8, 0xa8, 0x76, 0x8c, 0xbc, 0x81, 0x7a, 0x37, 0x9b, 0x92, 0xff, 0x03, 0x7b, 0x73, 0xba,
	0xd9, 0x94, 0xe2, 0x83, 0xb7, 0xfd, 0x12, 0x14, 0x39, 0x79, 0x07, 0xab, 0x3d, 0x26, 0xe4, 0x65,
	0x24, 0x23, 0xb2, 0x35, 0xa7, 0x50, 0x90, 0xb7, 0x3b, 0x07, 0x55, 0x4a, 0xfd, 0xe4, 0x15, 0x6c,
	0x5c, 0xa1, 0xd4, 0xd9, 0xfb, 0x32, 0x92, 0x85, 0x20, 0x7b, 0x73, 0x62, 0x8b, 0x51, 0xaf, 0x77,
	0x17, 0x91, 0xe7, 0xa7, 0xdf, 0x4e, 0x34, 0xf5, 0xba, 0xda, 0x7a, 0x91, 0x77, 0xf2, 0x50, 0xb5,
	0x8d, 0xdc, 0xfa, 0xd3, 0x0c, 0x1b, 0xba, 0xd9, 0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6c,
	0xdb, 0x0c, 0x60, 0x8f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCServiceClient is the client API for RPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCServiceClient interface {
	// SendToken — send token to other holder
	SendToken(ctx context.Context, in *TransferTokensRequest, opts ...grpc.CallOption) (*TransferTokensResponse, error)
	// IssueToken — Issue new token with given expiration data
	// sand return tokenID.
	IssueToken(ctx context.Context, in *IssueTokenRequest, opts ...grpc.CallOption) (*IssueTokenResponse, error)
	// LockToken — Return hash of lock token for verify htlc and information
	// about transaction
	LockToken(ctx context.Context, in *LockTokenRequest, opts ...grpc.CallOption) (*LockTokenResponse, error)
	// AddPeer append new peer to peer slice
	AddPeer(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Inv — send meta information to token strike
	Inv(ctx context.Context, in *tokenstrike.InvReq, opts ...grpc.CallOption) (*tokenstrike.InvResp, error)
	// PostData — send full data to replication
	PostData(ctx context.Context, in *tokenstrike.Data, opts ...grpc.CallOption) (*tokenstrike.PostDataResp, error)
	// GetTokenStatus — response with information about token
	GetTokenStatus(ctx context.Context, in *tokenstrike.TokenStatusReq, opts ...grpc.CallOption) (*tokenstrike.TokenStatus, error)
}

type rPCServiceClient struct {
	cc *grpc.ClientConn
}

func NewRPCServiceClient(cc *grpc.ClientConn) RPCServiceClient {
	return &rPCServiceClient{cc}
}

func (c *rPCServiceClient) SendToken(ctx context.Context, in *TransferTokensRequest, opts ...grpc.CallOption) (*TransferTokensResponse, error) {
	out := new(TransferTokensResponse)
	err := c.cc.Invoke(ctx, "/rpcservice.RPCService/SendToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) IssueToken(ctx context.Context, in *IssueTokenRequest, opts ...grpc.CallOption) (*IssueTokenResponse, error) {
	out := new(IssueTokenResponse)
	err := c.cc.Invoke(ctx, "/rpcservice.RPCService/IssueToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) LockToken(ctx context.Context, in *LockTokenRequest, opts ...grpc.CallOption) (*LockTokenResponse, error) {
	out := new(LockTokenResponse)
	err := c.cc.Invoke(ctx, "/rpcservice.RPCService/LockToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) AddPeer(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rpcservice.RPCService/AddPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) Inv(ctx context.Context, in *tokenstrike.InvReq, opts ...grpc.CallOption) (*tokenstrike.InvResp, error) {
	out := new(tokenstrike.InvResp)
	err := c.cc.Invoke(ctx, "/rpcservice.RPCService/Inv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) PostData(ctx context.Context, in *tokenstrike.Data, opts ...grpc.CallOption) (*tokenstrike.PostDataResp, error) {
	out := new(tokenstrike.PostDataResp)
	err := c.cc.Invoke(ctx, "/rpcservice.RPCService/PostData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) GetTokenStatus(ctx context.Context, in *tokenstrike.TokenStatusReq, opts ...grpc.CallOption) (*tokenstrike.TokenStatus, error) {
	out := new(tokenstrike.TokenStatus)
	err := c.cc.Invoke(ctx, "/rpcservice.RPCService/GetTokenStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServiceServer is the server API for RPCService service.
type RPCServiceServer interface {
	// SendToken — send token to other holder
	SendToken(context.Context, *TransferTokensRequest) (*TransferTokensResponse, error)
	// IssueToken — Issue new token with given expiration data
	// sand return tokenID.
	IssueToken(context.Context, *IssueTokenRequest) (*IssueTokenResponse, error)
	// LockToken — Return hash of lock token for verify htlc and information
	// about transaction
	LockToken(context.Context, *LockTokenRequest) (*LockTokenResponse, error)
	// AddPeer append new peer to peer slice
	AddPeer(context.Context, *PeerRequest) (*empty.Empty, error)
	// Inv — send meta information to token strike
	Inv(context.Context, *tokenstrike.InvReq) (*tokenstrike.InvResp, error)
	// PostData — send full data to replication
	PostData(context.Context, *tokenstrike.Data) (*tokenstrike.PostDataResp, error)
	// GetTokenStatus — response with information about token
	GetTokenStatus(context.Context, *tokenstrike.TokenStatusReq) (*tokenstrike.TokenStatus, error)
}

// UnimplementedRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRPCServiceServer struct {
}

func (*UnimplementedRPCServiceServer) SendToken(ctx context.Context, req *TransferTokensRequest) (*TransferTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToken not implemented")
}
func (*UnimplementedRPCServiceServer) IssueToken(ctx context.Context, req *IssueTokenRequest) (*IssueTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueToken not implemented")
}
func (*UnimplementedRPCServiceServer) LockToken(ctx context.Context, req *LockTokenRequest) (*LockTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockToken not implemented")
}
func (*UnimplementedRPCServiceServer) AddPeer(ctx context.Context, req *PeerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPeer not implemented")
}
func (*UnimplementedRPCServiceServer) Inv(ctx context.Context, req *tokenstrike.InvReq) (*tokenstrike.InvResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inv not implemented")
}
func (*UnimplementedRPCServiceServer) PostData(ctx context.Context, req *tokenstrike.Data) (*tokenstrike.PostDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostData not implemented")
}
func (*UnimplementedRPCServiceServer) GetTokenStatus(ctx context.Context, req *tokenstrike.TokenStatusReq) (*tokenstrike.TokenStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenStatus not implemented")
}

func RegisterRPCServiceServer(s *grpc.Server, srv RPCServiceServer) {
	s.RegisterService(&_RPCService_serviceDesc, srv)
}

func _RPCService_SendToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).SendToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcservice.RPCService/SendToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).SendToken(ctx, req.(*TransferTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_IssueToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).IssueToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcservice.RPCService/IssueToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).IssueToken(ctx, req.(*IssueTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_LockToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).LockToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcservice.RPCService/LockToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).LockToken(ctx, req.(*LockTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcservice.RPCService/AddPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).AddPeer(ctx, req.(*PeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_Inv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(tokenstrike.InvReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).Inv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcservice.RPCService/Inv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).Inv(ctx, req.(*tokenstrike.InvReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_PostData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(tokenstrike.Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).PostData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcservice.RPCService/PostData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).PostData(ctx, req.(*tokenstrike.Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_GetTokenStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(tokenstrike.TokenStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).GetTokenStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcservice.RPCService/GetTokenStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).GetTokenStatus(ctx, req.(*tokenstrike.TokenStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcservice.RPCService",
	HandlerType: (*RPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendToken",
			Handler:    _RPCService_SendToken_Handler,
		},
		{
			MethodName: "IssueToken",
			Handler:    _RPCService_IssueToken_Handler,
		},
		{
			MethodName: "LockToken",
			Handler:    _RPCService_LockToken_Handler,
		},
		{
			MethodName: "AddPeer",
			Handler:    _RPCService_AddPeer_Handler,
		},
		{
			MethodName: "Inv",
			Handler:    _RPCService_Inv_Handler,
		},
		{
			MethodName: "PostData",
			Handler:    _RPCService_PostData_Handler,
		},
		{
			MethodName: "GetTokenStatus",
			Handler:    _RPCService_GetTokenStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/PRCSevice/rpcservice.proto",
}
