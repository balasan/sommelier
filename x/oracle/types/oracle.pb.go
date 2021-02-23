// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/v1/oracle.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// OracleFeed represents an array of oracle data that is
type OracleFeed struct {
	Data []*types.Any `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *OracleFeed) Reset()         { *m = OracleFeed{} }
func (m *OracleFeed) String() string { return proto.CompactTextString(m) }
func (*OracleFeed) ProtoMessage()    {}
func (*OracleFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_652b57db11528d07, []int{0}
}
func (m *OracleFeed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleFeed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleFeed.Merge(m, src)
}
func (m *OracleFeed) XXX_Size() int {
	return m.Size()
}
func (m *OracleFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleFeed.DiscardUnknown(m)
}

var xxx_messageInfo_OracleFeed proto.InternalMessageInfo

func (m *OracleFeed) GetData() []*types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

// UniswapPair represents an SDK compatible uniswap pair info fetched from The Graph.
type UniswapPair struct {
	ID          string                                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Reserve0    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=reserve0,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reserve0" yaml:"reserve0"`
	Reserve1    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=reserve1,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reserve1" yaml:"reserve1"`
	ReserveUSD  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=reserve_usd,json=reserveUsd,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reserveUSD" yaml:"reserveUSD"`
	Token0      UniswapToken                           `protobuf:"bytes,5,opt,name=token0,proto3" json:"token0"`
	Token1      UniswapToken                           `protobuf:"bytes,6,opt,name=token1,proto3" json:"token1"`
	Token0Price github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=token0_price,json=token0Price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"token0Price" yaml:"token0Price"`
	Token1Price github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=token1_price,json=token1Price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"token1Price" yaml:"token1Price"`
	TotalSupply github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=total_supply,json=totalSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"totalSupply" yaml:"totalSupply"`
}

func (m *UniswapPair) Reset()         { *m = UniswapPair{} }
func (m *UniswapPair) String() string { return proto.CompactTextString(m) }
func (*UniswapPair) ProtoMessage()    {}
func (*UniswapPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_652b57db11528d07, []int{1}
}
func (m *UniswapPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UniswapPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UniswapPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UniswapPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniswapPair.Merge(m, src)
}
func (m *UniswapPair) XXX_Size() int {
	return m.Size()
}
func (m *UniswapPair) XXX_DiscardUnknown() {
	xxx_messageInfo_UniswapPair.DiscardUnknown(m)
}

var xxx_messageInfo_UniswapPair proto.InternalMessageInfo

func (m *UniswapPair) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UniswapPair) GetToken0() UniswapToken {
	if m != nil {
		return m.Token0
	}
	return UniswapToken{}
}

func (m *UniswapPair) GetToken1() UniswapToken {
	if m != nil {
		return m.Token1
	}
	return UniswapToken{}
}

// UniswapToken is the returned uniswap token representation
type UniswapToken struct {
	// token address
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// number of decimal positions of the pair token
	Decimals uint64 `protobuf:"varint,2,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (m *UniswapToken) Reset()         { *m = UniswapToken{} }
func (m *UniswapToken) String() string { return proto.CompactTextString(m) }
func (*UniswapToken) ProtoMessage()    {}
func (*UniswapToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_652b57db11528d07, []int{2}
}
func (m *UniswapToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UniswapToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UniswapToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UniswapToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniswapToken.Merge(m, src)
}
func (m *UniswapToken) XXX_Size() int {
	return m.Size()
}
func (m *UniswapToken) XXX_DiscardUnknown() {
	xxx_messageInfo_UniswapToken.DiscardUnknown(m)
}

var xxx_messageInfo_UniswapToken proto.InternalMessageInfo

func (m *UniswapToken) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UniswapToken) GetDecimals() uint64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

// UniswapToken is the returned uniswap token representation
type OracleVote struct {
	Salt []string    `protobuf:"bytes,1,rep,name=salt,proto3" json:"salt,omitempty"`
	Feed *OracleFeed `protobuf:"bytes,2,opt,name=feed,proto3" json:"feed,omitempty"`
}

func (m *OracleVote) Reset()         { *m = OracleVote{} }
func (m *OracleVote) String() string { return proto.CompactTextString(m) }
func (*OracleVote) ProtoMessage()    {}
func (*OracleVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_652b57db11528d07, []int{3}
}
func (m *OracleVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleVote.Merge(m, src)
}
func (m *OracleVote) XXX_Size() int {
	return m.Size()
}
func (m *OracleVote) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleVote.DiscardUnknown(m)
}

var xxx_messageInfo_OracleVote proto.InternalMessageInfo

func (m *OracleVote) GetSalt() []string {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *OracleVote) GetFeed() *OracleFeed {
	if m != nil {
		return m.Feed
	}
	return nil
}

// OraclePrevote defines an array of hashed from oracle data that are used
// for the prevote phase of the oracle data feeding.
type OraclePrevote struct {
	// hex formated hashes of each oracle feed
	Hashes DataHashes `protobuf:"bytes,1,rep,name=hashes,proto3,castrepeated=DataHashes" json:"hashes,omitempty"`
}

func (m *OraclePrevote) Reset()         { *m = OraclePrevote{} }
func (m *OraclePrevote) String() string { return proto.CompactTextString(m) }
func (*OraclePrevote) ProtoMessage()    {}
func (*OraclePrevote) Descriptor() ([]byte, []int) {
	return fileDescriptor_652b57db11528d07, []int{4}
}
func (m *OraclePrevote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OraclePrevote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OraclePrevote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OraclePrevote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OraclePrevote.Merge(m, src)
}
func (m *OraclePrevote) XXX_Size() int {
	return m.Size()
}
func (m *OraclePrevote) XXX_DiscardUnknown() {
	xxx_messageInfo_OraclePrevote.DiscardUnknown(m)
}

var xxx_messageInfo_OraclePrevote proto.InternalMessageInfo

func (m *OraclePrevote) GetHashes() DataHashes {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func init() {
	proto.RegisterType((*OracleFeed)(nil), "oracle.v1.OracleFeed")
	proto.RegisterType((*UniswapPair)(nil), "oracle.v1.UniswapPair")
	proto.RegisterType((*UniswapToken)(nil), "oracle.v1.UniswapToken")
	proto.RegisterType((*OracleVote)(nil), "oracle.v1.OracleVote")
	proto.RegisterType((*OraclePrevote)(nil), "oracle.v1.OraclePrevote")
}

func init() { proto.RegisterFile("oracle/v1/oracle.proto", fileDescriptor_652b57db11528d07) }

var fileDescriptor_652b57db11528d07 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4f, 0x6f, 0x94, 0x4e,
	0x18, 0xc7, 0x97, 0xed, 0xfe, 0xf8, 0x75, 0x87, 0x5a, 0x23, 0xa9, 0x15, 0x7b, 0x80, 0x86, 0x43,
	0x53, 0x63, 0x0a, 0x9d, 0x7a, 0xd0, 0x78, 0x31, 0x92, 0xb5, 0xd1, 0x18, 0x63, 0xc3, 0x5a, 0x0f,
	0x5e, 0x36, 0x53, 0x98, 0x52, 0x2c, 0xec, 0x10, 0x86, 0x45, 0xb9, 0xf9, 0x12, 0x7c, 0x1d, 0x9e,
	0x7d, 0x11, 0x8d, 0xa7, 0x1e, 0x8d, 0x07, 0x34, 0xdb, 0x8b, 0xd9, 0xe3, 0xbe, 0x02, 0xc3, 0xcc,
	0x2c, 0x4b, 0xa3, 0x1e, 0xd6, 0x9e, 0x98, 0xe7, 0xdf, 0xf7, 0x03, 0x0f, 0xf3, 0x3c, 0x60, 0x9d,
	0xa4, 0xc8, 0x8b, 0xb0, 0x9d, 0x43, 0x9b, 0x9f, 0xac, 0x24, 0x25, 0x19, 0x51, 0xbb, 0xc2, 0xca,
	0xe1, 0xc6, 0x5a, 0x40, 0x02, 0xc2, 0xbc, 0x76, 0x75, 0xe2, 0x09, 0x1b, 0xb7, 0x03, 0x42, 0x82,
	0x08, 0xdb, 0xcc, 0x3a, 0x1a, 0x1d, 0xdb, 0x68, 0x58, 0xcc, 0x42, 0x1e, 0xa1, 0x31, 0xa1, 0x03,
	0x5e, 0xc3, 0x0d, 0x1e, 0x32, 0xf7, 0x01, 0x78, 0xc9, 0x84, 0xf7, 0x31, 0xf6, 0xd5, 0x07, 0xa0,
	0xe3, 0xa3, 0x0c, 0x69, 0xd2, 0xe6, 0xd2, 0xb6, 0xb2, 0xb7, 0x66, 0x71, 0x49, 0x6b, 0x26, 0x69,
	0x3d, 0x1e, 0x16, 0xce, 0xea, 0x97, 0xcf, 0x3b, 0xa2, 0xa6, 0x87, 0x32, 0xe4, 0xb2, 0x0a, 0xf3,
	0xa7, 0x0c, 0x94, 0xc3, 0x61, 0x48, 0xdf, 0xa1, 0xe4, 0x00, 0x85, 0xa9, 0xba, 0x0e, 0xda, 0xa1,
	0xaf, 0x49, 0x9b, 0xd2, 0x76, 0xd7, 0x91, 0xc7, 0xa5, 0xd1, 0x7e, 0xd6, 0x73, 0xdb, 0xa1, 0xaf,
	0x86, 0x60, 0x39, 0xc5, 0x14, 0xa7, 0x39, 0xde, 0xd5, 0xda, 0x2c, 0xfa, 0xe2, 0xac, 0x34, 0x5a,
	0xdf, 0x4a, 0x63, 0x2b, 0x08, 0xb3, 0x93, 0xd1, 0x91, 0xe5, 0x91, 0x58, 0xbc, 0xa2, 0x78, 0xec,
	0x50, 0xff, 0xd4, 0xce, 0x8a, 0x04, 0x53, 0xab, 0x87, 0xbd, 0x49, 0x69, 0xd4, 0x0a, 0xd3, 0xd2,
	0xb8, 0x5e, 0xa0, 0x38, 0x7a, 0x68, 0xce, 0x3c, 0xa6, 0x5b, 0x07, 0x1b, 0x28, 0xa8, 0x2d, 0x5d,
	0x11, 0x05, 0x7f, 0x43, 0xc1, 0x39, 0x0a, 0xaa, 0x1f, 0x24, 0xa0, 0x08, 0x63, 0x30, 0xa2, 0xbe,
	0xd6, 0x61, 0xb8, 0xc1, 0x62, 0xb8, 0x71, 0x69, 0x00, 0x97, 0x8b, 0x1c, 0xf6, 0x7b, 0x93, 0xd2,
	0x00, 0x69, 0x6d, 0x4d, 0x4b, 0xe3, 0xc6, 0x25, 0xfc, 0x61, 0xbf, 0x67, 0xba, 0x75, 0x02, 0xf5,
	0xd5, 0x47, 0x40, 0xce, 0xc8, 0x29, 0x1e, 0xee, 0x6a, 0xff, 0x6d, 0x4a, 0xdb, 0xca, 0xde, 0x2d,
	0xab, 0xbe, 0x30, 0x96, 0xf8, 0x31, 0xaf, 0xaa, 0xb8, 0xb3, 0x5a, 0xbd, 0xd5, 0xa4, 0x34, 0x44,
	0xba, 0x2b, 0x9e, 0xb5, 0x00, 0xd4, 0xe4, 0x45, 0x04, 0xa0, 0x10, 0x80, 0x6a, 0x0e, 0x56, 0xb8,
	0xd4, 0x20, 0x49, 0x43, 0x0f, 0x6b, 0xff, 0xb3, 0x26, 0xf4, 0x17, 0xee, 0xb9, 0xc2, 0x55, 0x0e,
	0x2a, 0x91, 0x69, 0x69, 0xa8, 0xfc, 0xbb, 0x1b, 0x4e, 0xd3, 0x6d, 0xa6, 0xd4, 0x5c, 0x28, 0xb8,
	0xcb, 0x57, 0xe2, 0xc2, 0x3f, 0x71, 0xe1, 0x25, 0x2e, 0x6c, 0x70, 0x33, 0x14, 0x0d, 0xe8, 0x28,
	0x49, 0xa2, 0x42, 0xeb, 0xfe, 0x3b, 0x37, 0x43, 0x51, 0x9f, 0x89, 0x34, 0xb9, 0xb5, 0x93, 0x71,
	0xe7, 0x96, 0x03, 0x56, 0x9a, 0xff, 0xe3, 0xaf, 0xa3, 0xb6, 0x01, 0x96, 0x7d, 0xec, 0x85, 0x31,
	0x8a, 0x28, 0x1b, 0xb5, 0x8e, 0x5b, 0xdb, 0xe6, 0xf3, 0xd9, 0xd8, 0xbf, 0x26, 0x19, 0x56, 0x55,
	0xd0, 0xa1, 0x28, 0xca, 0xd8, 0xd8, 0x77, 0x5d, 0x76, 0x56, 0xef, 0x80, 0xce, 0x31, 0xc6, 0x3e,
	0xab, 0x54, 0xf6, 0x6e, 0x36, 0x2e, 0xc3, 0x7c, 0x5f, 0xb8, 0x2c, 0xc5, 0xbc, 0x0f, 0xae, 0x71,
	0xdf, 0x41, 0x8a, 0xf3, 0x4a, 0x6f, 0x0b, 0xc8, 0x27, 0x88, 0x9e, 0x60, 0xca, 0x14, 0x57, 0x9c,
	0xd5, 0x4f, 0xdf, 0x0d, 0x50, 0x2d, 0x8b, 0xa7, 0xcc, 0xeb, 0x8a, 0xa8, 0xf3, 0xe4, 0x6c, 0xac,
	0x4b, 0xe7, 0x63, 0x5d, 0xfa, 0x31, 0xd6, 0xa5, 0x8f, 0x17, 0x7a, 0xeb, 0xfc, 0x42, 0x6f, 0x7d,
	0xbd, 0xd0, 0x5b, 0x6f, 0xee, 0x36, 0xba, 0x97, 0xe0, 0x20, 0x28, 0xde, 0xe6, 0x36, 0x25, 0x71,
	0x8c, 0xa3, 0x10, 0xa7, 0xf6, 0x7b, 0xb1, 0x1a, 0x79, 0x1b, 0x8f, 0x64, 0xb6, 0x9f, 0xee, 0xfd,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x6b, 0xc8, 0xf6, 0x3b, 0x05, 0x00, 0x00,
}

func (m *OracleFeed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleFeed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleFeed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOracle(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UniswapPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UniswapPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UniswapPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalSupply.Size()
		i -= size
		if _, err := m.TotalSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.Token1Price.Size()
		i -= size
		if _, err := m.Token1Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.Token0Price.Size()
		i -= size
		if _, err := m.Token0Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.Token1.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Token0.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.ReserveUSD.Size()
		i -= size
		if _, err := m.ReserveUSD.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Reserve1.Size()
		i -= size
		if _, err := m.Reserve1.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Reserve0.Size()
		i -= size
		if _, err := m.Reserve0.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UniswapToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UniswapToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UniswapToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimals != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OracleVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Feed != nil {
		{
			size, err := m.Feed.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOracle(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Salt) > 0 {
		for iNdEx := len(m.Salt) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Salt[iNdEx])
			copy(dAtA[i:], m.Salt[iNdEx])
			i = encodeVarintOracle(dAtA, i, uint64(len(m.Salt[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *OraclePrevote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OraclePrevote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OraclePrevote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for iNdEx := len(m.Hashes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Hashes[iNdEx])
			copy(dAtA[i:], m.Hashes[iNdEx])
			i = encodeVarintOracle(dAtA, i, uint64(len(m.Hashes[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OracleFeed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovOracle(uint64(l))
		}
	}
	return n
}

func (m *UniswapPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = m.Reserve0.Size()
	n += 1 + l + sovOracle(uint64(l))
	l = m.Reserve1.Size()
	n += 1 + l + sovOracle(uint64(l))
	l = m.ReserveUSD.Size()
	n += 1 + l + sovOracle(uint64(l))
	l = m.Token0.Size()
	n += 1 + l + sovOracle(uint64(l))
	l = m.Token1.Size()
	n += 1 + l + sovOracle(uint64(l))
	l = m.Token0Price.Size()
	n += 1 + l + sovOracle(uint64(l))
	l = m.Token1Price.Size()
	n += 1 + l + sovOracle(uint64(l))
	l = m.TotalSupply.Size()
	n += 1 + l + sovOracle(uint64(l))
	return n
}

func (m *UniswapToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovOracle(uint64(m.Decimals))
	}
	return n
}

func (m *OracleVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Salt) > 0 {
		for _, s := range m.Salt {
			l = len(s)
			n += 1 + l + sovOracle(uint64(l))
		}
	}
	if m.Feed != nil {
		l = m.Feed.Size()
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}

func (m *OraclePrevote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for _, b := range m.Hashes {
			l = len(b)
			n += 1 + l + sovOracle(uint64(l))
		}
	}
	return n
}

func sovOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracle(x uint64) (n int) {
	return sovOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OracleFeed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OracleFeed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleFeed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &types.Any{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UniswapPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UniswapPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UniswapPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserve0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reserve0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserve1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reserve1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveUSD", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReserveUSD.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token0", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token0Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token0Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token1Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token1Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UniswapToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UniswapToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UniswapToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OracleVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OracleVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt = append(m.Salt, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Feed == nil {
				m.Feed = &OracleFeed{}
			}
			if err := m.Feed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OraclePrevote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OraclePrevote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OraclePrevote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashes = append(m.Hashes, make([]byte, postIndex-iNdEx))
			copy(m.Hashes[len(m.Hashes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthOracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracle = fmt.Errorf("proto: unexpected end of group")
)
