// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/tokenstrike/tokenstrike.proto

package tokenstrike

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	DB "token-strike/tsp2p/server/DB"
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

type InvReq struct {
	Invs                 []*Inv   `protobuf:"bytes,1,rep,name=invs,proto3" json:"invs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvReq) Reset()         { *m = InvReq{} }
func (m *InvReq) String() string { return proto.CompactTextString(m) }
func (*InvReq) ProtoMessage()    {}
func (*InvReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da023b1b27b350a, []int{0}
}

func (m *InvReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvReq.Unmarshal(m, b)
}
func (m *InvReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvReq.Marshal(b, m, deterministic)
}
func (m *InvReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvReq.Merge(m, src)
}
func (m *InvReq) XXX_Size() int {
	return xxx_messageInfo_InvReq.Size(m)
}
func (m *InvReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InvReq.DiscardUnknown(m)
}

var xxx_messageInfo_InvReq proto.InternalMessageInfo

func (m *InvReq) GetInvs() []*Inv {
	if m != nil {
		return m.Invs
	}
	return nil
}

// Inv contains info about data that need replicates to other replicators
type Inv struct {
	/// For block or lock, the "parent" is the token id (hash block 0)
	/// For new token notifications, the parent is the issuer
	Parent               []byte   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Type                 uint32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	EntityHash           []byte   `protobuf:"bytes,3,opt,name=entity_hash,json=entityHash,proto3" json:"entity_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Inv) Reset()         { *m = Inv{} }
func (m *Inv) String() string { return proto.CompactTextString(m) }
func (*Inv) ProtoMessage()    {}
func (*Inv) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da023b1b27b350a, []int{1}
}

func (m *Inv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inv.Unmarshal(m, b)
}
func (m *Inv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inv.Marshal(b, m, deterministic)
}
func (m *Inv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inv.Merge(m, src)
}
func (m *Inv) XXX_Size() int {
	return xxx_messageInfo_Inv.Size(m)
}
func (m *Inv) XXX_DiscardUnknown() {
	xxx_messageInfo_Inv.DiscardUnknown(m)
}

var xxx_messageInfo_Inv proto.InternalMessageInfo

func (m *Inv) GetParent() []byte {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *Inv) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Inv) GetEntityHash() []byte {
	if m != nil {
		return m.EntityHash
	}
	return nil
}

type InvResp struct {
	Needed               []bool   `protobuf:"varint,1,rep,packed,name=needed,proto3" json:"needed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvResp) Reset()         { *m = InvResp{} }
func (m *InvResp) String() string { return proto.CompactTextString(m) }
func (*InvResp) ProtoMessage()    {}
func (*InvResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da023b1b27b350a, []int{2}
}

func (m *InvResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvResp.Unmarshal(m, b)
}
func (m *InvResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvResp.Marshal(b, m, deterministic)
}
func (m *InvResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvResp.Merge(m, src)
}
func (m *InvResp) XXX_Size() int {
	return xxx_messageInfo_InvResp.Size(m)
}
func (m *InvResp) XXX_DiscardUnknown() {
	xxx_messageInfo_InvResp.DiscardUnknown(m)
}

var xxx_messageInfo_InvResp proto.InternalMessageInfo

func (m *InvResp) GetNeeded() []bool {
	if m != nil {
		return m.Needed
	}
	return nil
}

// PostData
type PostDataResp struct {
	Warning              []string `protobuf:"bytes,2,rep,name=warning,proto3" json:"warning,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostDataResp) Reset()         { *m = PostDataResp{} }
func (m *PostDataResp) String() string { return proto.CompactTextString(m) }
func (*PostDataResp) ProtoMessage()    {}
func (*PostDataResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da023b1b27b350a, []int{3}
}

func (m *PostDataResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostDataResp.Unmarshal(m, b)
}
func (m *PostDataResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostDataResp.Marshal(b, m, deterministic)
}
func (m *PostDataResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostDataResp.Merge(m, src)
}
func (m *PostDataResp) XXX_Size() int {
	return xxx_messageInfo_PostDataResp.Size(m)
}
func (m *PostDataResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PostDataResp.DiscardUnknown(m)
}

var xxx_messageInfo_PostDataResp proto.InternalMessageInfo

func (m *PostDataResp) GetWarning() []string {
	if m != nil {
		return m.Warning
	}
	return nil
}

type Data struct {
	// Types that are valid to be assigned to Data:
	//	*Data_Lock
	//	*Data_Block
	//	*Data_Transfer
	Data                 isData_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da023b1b27b350a, []int{4}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

type isData_Data interface {
	isData_Data()
}

type Data_Lock struct {
	Lock *lock.Lock `protobuf:"bytes,1,opt,name=lock,proto3,oneof"`
}

type Data_Block struct {
	Block *DB.Block `protobuf:"bytes,2,opt,name=block,proto3,oneof"`
}

type Data_Transfer struct {
	Transfer *TransferTokens `protobuf:"bytes,3,opt,name=transfer,proto3,oneof"`
}

func (*Data_Lock) isData_Data() {}

func (*Data_Block) isData_Data() {}

func (*Data_Transfer) isData_Data() {}

func (m *Data) GetData() isData_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Data) GetLock() *lock.Lock {
	if x, ok := m.GetData().(*Data_Lock); ok {
		return x.Lock
	}
	return nil
}

func (m *Data) GetBlock() *DB.Block {
	if x, ok := m.GetData().(*Data_Block); ok {
		return x.Block
	}
	return nil
}

func (m *Data) GetTransfer() *TransferTokens {
	if x, ok := m.GetData().(*Data_Transfer); ok {
		return x.Transfer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Data) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Data_Lock)(nil),
		(*Data_Block)(nil),
		(*Data_Transfer)(nil),
	}
}

type TransferTokens struct {
	Htlc                 []byte   `protobuf:"bytes,1,opt,name=htlc,proto3" json:"htlc,omitempty"`
	Lock                 []byte   `protobuf:"bytes,2,opt,name=lock,proto3" json:"lock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferTokens) Reset()         { *m = TransferTokens{} }
func (m *TransferTokens) String() string { return proto.CompactTextString(m) }
func (*TransferTokens) ProtoMessage()    {}
func (*TransferTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da023b1b27b350a, []int{5}
}

func (m *TransferTokens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferTokens.Unmarshal(m, b)
}
func (m *TransferTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferTokens.Marshal(b, m, deterministic)
}
func (m *TransferTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferTokens.Merge(m, src)
}
func (m *TransferTokens) XXX_Size() int {
	return xxx_messageInfo_TransferTokens.Size(m)
}
func (m *TransferTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferTokens.DiscardUnknown(m)
}

var xxx_messageInfo_TransferTokens proto.InternalMessageInfo

func (m *TransferTokens) GetHtlc() []byte {
	if m != nil {
		return m.Htlc
	}
	return nil
}

func (m *TransferTokens) GetLock() []byte {
	if m != nil {
		return m.Lock
	}
	return nil
}

func init() {
	proto.RegisterType((*InvReq)(nil), "tokenstrike.InvReq")
	proto.RegisterType((*Inv)(nil), "tokenstrike.Inv")
	proto.RegisterType((*InvResp)(nil), "tokenstrike.InvResp")
	proto.RegisterType((*PostDataResp)(nil), "tokenstrike.PostDataResp")
	proto.RegisterType((*Data)(nil), "tokenstrike.Data")
	proto.RegisterType((*TransferTokens)(nil), "tokenstrike.TransferTokens")
}

func init() {
	proto.RegisterFile("protos/tokenstrike/tokenstrike.proto", fileDescriptor_7da023b1b27b350a)
}

var fileDescriptor_7da023b1b27b350a = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x41, 0xaf, 0x93, 0x40,
	0x10, 0x7e, 0x14, 0xe4, 0xd5, 0xa1, 0xbe, 0xe8, 0x6a, 0x9e, 0x88, 0x07, 0x91, 0x3c, 0x95, 0x8b,
	0xd4, 0xe0, 0xa5, 0x5e, 0x49, 0x0f, 0x6d, 0xe2, 0xc1, 0xac, 0x3d, 0x79, 0x31, 0xdb, 0xb2, 0x0a,
	0xa1, 0x59, 0xe8, 0xee, 0x06, 0xd3, 0xdf, 0xe1, 0x1f, 0x7e, 0x61, 0x16, 0x1a, 0x48, 0x2f, 0x64,
	0xe6, 0xfb, 0xbe, 0xd9, 0xf9, 0xf2, 0x31, 0xf0, 0xd0, 0xc8, 0x5a, 0xd7, 0x6a, 0xa9, 0xeb, 0x8a,
	0x0b, 0xa5, 0x65, 0x59, 0xf1, 0x71, 0x9d, 0x20, 0x4d, 0xbc, 0x11, 0x14, 0xdc, 0xf7, 0x23, 0xc7,
	0xfa, 0x50, 0xe1, 0xc7, 0x88, 0x82, 0xd7, 0x3d, 0xbe, 0xce, 0xcc, 0x0b, 0xf9, 0xde, 0x10, 0x51,
	0x02, 0xee, 0x56, 0xb4, 0x94, 0x9f, 0xc8, 0x03, 0x38, 0xa5, 0x68, 0x95, 0x6f, 0x85, 0x76, 0xec,
	0xa5, 0xcf, 0x93, 0xf1, 0xa6, 0x4e, 0x82, 0x6c, 0x44, 0xc1, 0xde, 0x8a, 0x96, 0xdc, 0x83, 0xdb,
	0x30, 0xc9, 0x85, 0xf6, 0xad, 0xd0, 0x8a, 0x17, 0xb4, 0xef, 0x08, 0x01, 0x47, 0x9f, 0x1b, 0xee,
	0xcf, 0x42, 0x2b, 0x7e, 0x46, 0xb1, 0x26, 0xef, 0xc0, 0xe3, 0x42, 0x97, 0xfa, 0xfc, 0xbb, 0x60,
	0xaa, 0xf0, 0x6d, 0x1c, 0x00, 0x03, 0x6d, 0x98, 0x2a, 0xa2, 0xf7, 0x70, 0x8b, 0x1e, 0x54, 0xd3,
	0xbd, 0x2b, 0x38, 0xcf, 0x79, 0x8e, 0x36, 0xe6, 0xb4, 0xef, 0xa2, 0x18, 0x16, 0x3f, 0x6a, 0xa5,
	0xd7, 0x4c, 0x33, 0xd4, 0xf9, 0x70, 0xfb, 0x8f, 0x49, 0x51, 0x8a, 0xbf, 0xfe, 0x2c, 0xb4, 0xe3,
	0xa7, 0x74, 0x68, 0xa3, 0xff, 0x16, 0x38, 0x9d, 0x8c, 0x84, 0xe0, 0x74, 0x01, 0xa0, 0x41, 0x2f,
	0x85, 0x04, 0xd3, 0xf8, 0x5e, 0x1f, 0xaa, 0xcd, 0x0d, 0x45, 0x86, 0x7c, 0x84, 0x27, 0x7b, 0x94,
	0xcc, 0x50, 0x72, 0x97, 0x0c, 0xd1, 0x64, 0x47, 0x23, 0x33, 0x34, 0xf9, 0x06, 0x73, 0x2d, 0x99,
	0x50, 0x7f, 0xb8, 0x44, 0xf7, 0x5e, 0xfa, 0x76, 0x92, 0xce, 0xae, 0x27, 0x77, 0x88, 0x6d, 0x6e,
	0xe8, 0x45, 0x9e, 0xb9, 0xe0, 0xe4, 0x4c, 0xb3, 0x68, 0x05, 0x77, 0x53, 0x55, 0x97, 0x54, 0xa1,
	0x8f, 0x87, 0x3e, 0x3f, 0xac, 0x3b, 0xec, 0xe2, 0x67, 0x61, 0x4c, 0xa6, 0x67, 0xf0, 0x70, 0xe2,
	0x27, 0xee, 0x22, 0x5f, 0x4c, 0xfe, 0x2f, 0xaf, 0x7e, 0x0f, 0x3f, 0x05, 0xaf, 0xae, 0x41, 0xd5,
	0x90, 0x15, 0xcc, 0x87, 0xe8, 0xc8, 0x8b, 0x89, 0xa2, 0x83, 0x82, 0x37, 0x13, 0x68, 0x1c, 0x72,
	0xf6, 0xe9, 0xd7, 0x07, 0xe4, 0x3e, 0x0f, 0xb7, 0xa7, 0x9a, 0xb4, 0x59, 0x2a, 0x2e, 0x5b, 0x2e,
	0xc7, 0x87, 0xb8, 0x77, 0xf1, 0x96, 0xbe, 0x3e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x5f, 0x8e,
	0x19, 0xb1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TokenStrikeClient is the client API for TokenStrike service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenStrikeClient interface {
	Inv(ctx context.Context, in *InvReq, opts ...grpc.CallOption) (*InvResp, error)
	PostData(ctx context.Context, in *Data, opts ...grpc.CallOption) (*PostDataResp, error)
}

type tokenStrikeClient struct {
	cc *grpc.ClientConn
}

func NewTokenStrikeClient(cc *grpc.ClientConn) TokenStrikeClient {
	return &tokenStrikeClient{cc}
}

func (c *tokenStrikeClient) Inv(ctx context.Context, in *InvReq, opts ...grpc.CallOption) (*InvResp, error) {
	out := new(InvResp)
	err := c.cc.Invoke(ctx, "/tokenstrike.TokenStrike/Inv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenStrikeClient) PostData(ctx context.Context, in *Data, opts ...grpc.CallOption) (*PostDataResp, error) {
	out := new(PostDataResp)
	err := c.cc.Invoke(ctx, "/tokenstrike.TokenStrike/PostData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenStrikeServer is the server API for TokenStrike service.
type TokenStrikeServer interface {
	Inv(context.Context, *InvReq) (*InvResp, error)
	PostData(context.Context, *Data) (*PostDataResp, error)
}

// UnimplementedTokenStrikeServer can be embedded to have forward compatible implementations.
type UnimplementedTokenStrikeServer struct {
}

func (*UnimplementedTokenStrikeServer) Inv(ctx context.Context, req *InvReq) (*InvResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inv not implemented")
}
func (*UnimplementedTokenStrikeServer) PostData(ctx context.Context, req *Data) (*PostDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostData not implemented")
}

func RegisterTokenStrikeServer(s *grpc.Server, srv TokenStrikeServer) {
	s.RegisterService(&_TokenStrike_serviceDesc, srv)
}

func _TokenStrike_Inv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenStrikeServer).Inv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenstrike.TokenStrike/Inv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenStrikeServer).Inv(ctx, req.(*InvReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenStrike_PostData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenStrikeServer).PostData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenstrike.TokenStrike/PostData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenStrikeServer).PostData(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenStrike_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tokenstrike.TokenStrike",
	HandlerType: (*TokenStrikeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Inv",
			Handler:    _TokenStrike_Inv_Handler,
		},
		{
			MethodName: "PostData",
			Handler:    _TokenStrike_PostData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/tokenstrike/tokenstrike.proto",
}
