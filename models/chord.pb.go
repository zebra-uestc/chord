// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chord.proto

package models // import "github.com/zebra-uestc/chord/models"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Node contains a node ID and address.
// ID即为hash后的值，address为通信地址
type Node struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Node) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type ER struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ER) Reset()         { *m = ER{} }
func (m *ER) String() string { return proto.CompactTextString(m) }
func (*ER) ProtoMessage()    {}
func (*ER) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{1}
}
func (m *ER) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ER.Unmarshal(m, b)
}
func (m *ER) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ER.Marshal(b, m, deterministic)
}
func (dst *ER) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ER.Merge(dst, src)
}
func (m *ER) XXX_Size() int {
	return xxx_messageInfo_ER.Size(m)
}
func (m *ER) XXX_DiscardUnknown() {
	xxx_messageInfo_ER.DiscardUnknown(m)
}

var xxx_messageInfo_ER proto.InternalMessageInfo

type ID struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{2}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (dst *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(dst, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{3}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResponse struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{4}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{5}
}
func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (dst *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(dst, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{6}
}
func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (dst *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(dst, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

type DeleteRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{7}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{8}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(dst, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type MultiDeleteRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiDeleteRequest) Reset()         { *m = MultiDeleteRequest{} }
func (m *MultiDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*MultiDeleteRequest) ProtoMessage()    {}
func (*MultiDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{9}
}
func (m *MultiDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiDeleteRequest.Unmarshal(m, b)
}
func (m *MultiDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiDeleteRequest.Marshal(b, m, deterministic)
}
func (dst *MultiDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiDeleteRequest.Merge(dst, src)
}
func (m *MultiDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_MultiDeleteRequest.Size(m)
}
func (m *MultiDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiDeleteRequest proto.InternalMessageInfo

func (m *MultiDeleteRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type RequestKeysRequest struct {
	From                 []byte   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   []byte   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestKeysRequest) Reset()         { *m = RequestKeysRequest{} }
func (m *RequestKeysRequest) String() string { return proto.CompactTextString(m) }
func (*RequestKeysRequest) ProtoMessage()    {}
func (*RequestKeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{10}
}
func (m *RequestKeysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestKeysRequest.Unmarshal(m, b)
}
func (m *RequestKeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestKeysRequest.Marshal(b, m, deterministic)
}
func (dst *RequestKeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestKeysRequest.Merge(dst, src)
}
func (m *RequestKeysRequest) XXX_Size() int {
	return xxx_messageInfo_RequestKeysRequest.Size(m)
}
func (m *RequestKeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestKeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RequestKeysRequest proto.InternalMessageInfo

func (m *RequestKeysRequest) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *RequestKeysRequest) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

type KV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{11}
}
func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (dst *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(dst, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RequestKeysResponse struct {
	Values               []*KV    `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestKeysResponse) Reset()         { *m = RequestKeysResponse{} }
func (m *RequestKeysResponse) String() string { return proto.CompactTextString(m) }
func (*RequestKeysResponse) ProtoMessage()    {}
func (*RequestKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chord_fd1a63cdbe16875e, []int{12}
}
func (m *RequestKeysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestKeysResponse.Unmarshal(m, b)
}
func (m *RequestKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestKeysResponse.Marshal(b, m, deterministic)
}
func (dst *RequestKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestKeysResponse.Merge(dst, src)
}
func (m *RequestKeysResponse) XXX_Size() int {
	return xxx_messageInfo_RequestKeysResponse.Size(m)
}
func (m *RequestKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestKeysResponse proto.InternalMessageInfo

func (m *RequestKeysResponse) GetValues() []*KV {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "models.Node")
	proto.RegisterType((*ER)(nil), "models.ER")
	proto.RegisterType((*ID)(nil), "models.ID")
	proto.RegisterType((*GetRequest)(nil), "models.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "models.GetResponse")
	proto.RegisterType((*SetRequest)(nil), "models.SetRequest")
	proto.RegisterType((*SetResponse)(nil), "models.SetResponse")
	proto.RegisterType((*DeleteRequest)(nil), "models.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "models.DeleteResponse")
	proto.RegisterType((*MultiDeleteRequest)(nil), "models.MultiDeleteRequest")
	proto.RegisterType((*RequestKeysRequest)(nil), "models.RequestKeysRequest")
	proto.RegisterType((*KV)(nil), "models.KV")
	proto.RegisterType((*RequestKeysResponse)(nil), "models.RequestKeysResponse")
}

func init() { proto.RegisterFile("chord.proto", fileDescriptor_chord_fd1a63cdbe16875e) }

var fileDescriptor_chord_fd1a63cdbe16875e = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x6d, 0x6b, 0xd3, 0x50,
	0x14, 0xa6, 0x59, 0x56, 0xd9, 0xd3, 0xb4, 0x94, 0xb3, 0x29, 0x25, 0x82, 0xd4, 0x3b, 0x84, 0x58,
	0xb4, 0x85, 0xe9, 0x87, 0xf9, 0x75, 0xeb, 0x0c, 0xa5, 0x38, 0x24, 0x81, 0x51, 0xfc, 0xd6, 0x26,
	0x67, 0x36, 0xb4, 0xdd, 0x9d, 0xc9, 0x8d, 0x50, 0xff, 0xa8, 0x7f, 0x47, 0xf2, 0x66, 0x93, 0xbe,
	0xe8, 0xbe, 0x9d, 0x9c, 0xe7, 0x25, 0xe7, 0xdc, 0xfb, 0x70, 0xd1, 0xf0, 0xe6, 0x32, 0xf4, 0xfb,
	0x8f, 0xa1, 0x54, 0x92, 0xea, 0x2b, 0xe9, 0xf3, 0x32, 0x12, 0x3d, 0xe8, 0xb7, 0xd2, 0x67, 0x6a,
	0x41, 0x0b, 0xfc, 0x4e, 0xad, 0x5b, 0xb3, 0x0c, 0x47, 0x0b, 0x7c, 0x22, 0xe8, 0x53, 0xdf, 0x0f,
	0x3b, 0x5a, 0xb7, 0x66, 0x9d, 0x38, 0x69, 0x2d, 0x74, 0x68, 0x37, 0x8e, 0x38, 0x83, 0x36, 0x1a,
	0x6e, 0xf3, 0xc5, 0x2b, 0xc0, 0x66, 0xe5, 0xf0, 0x8f, 0x98, 0x23, 0x45, 0x6d, 0x1c, 0x2d, 0x78,
	0x9d, 0xc2, 0x27, 0x4e, 0x52, 0x8a, 0x73, 0x34, 0x52, 0x3c, 0x7a, 0x94, 0x0f, 0x11, 0xd3, 0x19,
	0x8e, 0x7f, 0x4e, 0x97, 0x31, 0xe7, 0x0e, 0xd9, 0x87, 0xf8, 0x08, 0xb8, 0xff, 0x30, 0xd9, 0xa8,
	0xb2, 0xa9, 0x72, 0x55, 0x13, 0x0d, 0x77, 0x63, 0x2d, 0x5e, 0xa3, 0x39, 0xe4, 0x25, 0x2b, 0x3e,
	0x3c, 0x4c, 0x1b, 0xad, 0x82, 0x92, 0x8b, 0x2c, 0xd0, 0x97, 0x78, 0xa9, 0x82, 0xaa, 0x92, 0xa0,
	0x2f, 0x78, 0x1d, 0x75, 0x6a, 0xdd, 0xa3, 0xe4, 0x10, 0x92, 0x5a, 0x5c, 0x82, 0x72, 0x78, 0xcc,
	0xeb, 0xa8, 0xc4, 0xbc, 0x0f, 0xe5, 0x2a, 0x5f, 0x27, 0xad, 0x93, 0x23, 0x52, 0x32, 0x1d, 0xd5,
	0x70, 0x34, 0x25, 0xc5, 0x3b, 0x68, 0xe3, 0xbb, 0x27, 0x6f, 0xf5, 0x09, 0xa7, 0x95, 0xff, 0xe4,
	0x07, 0x27, 0x50, 0x4f, 0xf1, 0x6c, 0xa8, 0xc6, 0x05, 0xfa, 0xd9, 0x45, 0xf6, 0xc7, 0x77, 0x4e,
	0x8e, 0x5c, 0xfc, 0xd6, 0x71, 0x7c, 0x9d, 0xdc, 0x35, 0xf5, 0xd0, 0xb2, 0x59, 0x7d, 0x0d, 0xd9,
	0x67, 0x8f, 0xa3, 0x48, 0x86, 0xf4, 0x97, 0x7f, 0xe3, 0x98, 0x46, 0x51, 0xa7, 0x09, 0xb0, 0x60,
	0xd8, 0xac, 0xdc, 0xd8, 0xfb, 0x2f, 0x53, 0xa0, 0x7e, 0x2b, 0x55, 0x70, 0xbf, 0xa6, 0x4a, 0xdf,
	0x2c, 0x29, 0xe8, 0x2d, 0x9a, 0x9f, 0x83, 0x07, 0x7f, 0x8f, 0xdd, 0x68, 0xb8, 0x65, 0xd7, 0x43,
	0xfb, 0x7a, 0xce, 0xde, 0x62, 0xef, 0x98, 0xa3, 0x61, 0xc5, 0xb6, 0x87, 0x96, 0x5b, 0x5d, 0xe8,
	0xf0, 0x08, 0x16, 0x0c, 0xb7, 0xbc, 0xd0, 0x61, 0xe6, 0x00, 0xfa, 0xc4, 0x66, 0x45, 0x54, 0xf4,
	0x36, 0x51, 0x36, 0x4f, 0x2b, 0xbd, 0xfc, 0x16, 0x12, 0x81, 0x5b, 0x16, 0xb8, 0x7b, 0x04, 0xa5,
	0x50, 0xd2, 0x25, 0x9e, 0x4d, 0xb2, 0x6c, 0xd1, 0xf3, 0x02, 0xaf, 0x64, 0xcd, 0x7c, 0xb1, 0xdd,
	0xce, 0x95, 0x57, 0x30, 0x26, 0xa5, 0x68, 0x92, 0x59, 0xf0, 0x76, 0xf3, 0x7a, 0xd0, 0xc3, 0x86,
	0x31, 0x29, 0x85, 0x69, 0xe3, 0xb1, 0x9b, 0x64, 0xf3, 0xe5, 0x5e, 0x2c, 0x33, 0xba, 0x7a, 0xf3,
	0xed, 0xfc, 0x7b, 0xa0, 0xe6, 0xf1, 0xac, 0xef, 0xc9, 0xd5, 0xe0, 0x17, 0xcf, 0xc2, 0xe9, 0xfb,
	0x84, 0xe4, 0x0d, 0xd2, 0xb7, 0x65, 0x90, 0x49, 0x67, 0xf5, 0xf4, 0x8d, 0xf9, 0xf0, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x3f, 0x01, 0xc7, 0x46, 0x72, 0x04, 0x00, 0x00,
}