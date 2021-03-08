// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: chord.proto

package chord

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Node contains a node ID and address.
// ID即为hash后的值，address为通信地址
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Node) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type ER struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ER) Reset() {
	*x = ER{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ER) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ER) ProtoMessage() {}

func (x *ER) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ER.ProtoReflect.Descriptor instead.
func (*ER) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{1}
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{2}
}

func (x *ID) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{4}
}

func (x *GetResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{5}
}

func (x *SetRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *SetRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetResponse) Reset() {
	*x = SetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResponse) ProtoMessage() {}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResponse.ProtoReflect.Descriptor instead.
func (*SetResponse) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{6}
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{8}
}

type MultiDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys [][]byte `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *MultiDeleteRequest) Reset() {
	*x = MultiDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiDeleteRequest) ProtoMessage() {}

func (x *MultiDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiDeleteRequest.ProtoReflect.Descriptor instead.
func (*MultiDeleteRequest) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{9}
}

func (x *MultiDeleteRequest) GetKeys() [][]byte {
	if x != nil {
		return x.Keys
	}
	return nil
}

type RequestKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From []byte `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   []byte `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *RequestKeysRequest) Reset() {
	*x = RequestKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestKeysRequest) ProtoMessage() {}

func (x *RequestKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestKeysRequest.ProtoReflect.Descriptor instead.
func (*RequestKeysRequest) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{10}
}

func (x *RequestKeysRequest) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *RequestKeysRequest) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

type KV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KV) Reset() {
	*x = KV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KV) ProtoMessage() {}

func (x *KV) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KV.ProtoReflect.Descriptor instead.
func (*KV) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{11}
}

func (x *KV) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KV) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type RequestKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*KV `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *RequestKeysResponse) Reset() {
	*x = RequestKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestKeysResponse) ProtoMessage() {}

func (x *RequestKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestKeysResponse.ProtoReflect.Descriptor instead.
func (*RequestKeysResponse) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{12}
}

func (x *RequestKeysResponse) GetValues() []*KV {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_chord_proto protoreflect.FileDescriptor

var file_chord_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x68, 0x6f, 0x72, 0x64, 0x22, 0x2a, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x22, 0x04, 0x0a, 0x02, 0x45, 0x52, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x23, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x34, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x38, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x6f, 0x22,
	0x2c, 0x0a, 0x02, 0x4b, 0x56, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a,
	0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x4b, 0x56, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x32, 0xbf, 0x04, 0x0a, 0x05, 0x43, 0x68, 0x6f, 0x72,
	0x64, 0x12, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x12, 0x09, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x45, 0x52, 0x1a, 0x0b,
	0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x09, 0x2e, 0x63, 0x68,
	0x6f, 0x72, 0x64, 0x2e, 0x45, 0x52, 0x1a, 0x0b, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x0b, 0x2e,
	0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x09, 0x2e, 0x63, 0x68, 0x6f,
	0x72, 0x64, 0x2e, 0x45, 0x52, 0x12, 0x27, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x09, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x49,
	0x44, 0x1a, 0x0b, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x12, 0x09, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x49, 0x44, 0x1a, 0x09, 0x2e,
	0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x45, 0x52, 0x12, 0x28, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x50,
	0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x0b, 0x2e, 0x63, 0x68, 0x6f,
	0x72, 0x64, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x09, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e,
	0x45, 0x52, 0x12, 0x26, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x12, 0x0b, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a,
	0x09, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x45, 0x52, 0x12, 0x2d, 0x0a, 0x04, 0x58, 0x47,
	0x65, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x58, 0x53, 0x65,
	0x74, 0x12, 0x11, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x58, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x68, 0x6f, 0x72,
	0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x40, 0x0a, 0x0c, 0x58, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x19, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x68,
	0x6f, 0x72, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x58, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4b, 0x65,
	0x79, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4b, 0x65, 0x79,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x65, 0x62, 0x72, 0x61, 0x2d, 0x75, 0x65,
	0x73, 0x74, 0x63, 0x2f, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chord_proto_rawDescOnce sync.Once
	file_chord_proto_rawDescData = file_chord_proto_rawDesc
)

func file_chord_proto_rawDescGZIP() []byte {
	file_chord_proto_rawDescOnce.Do(func() {
		file_chord_proto_rawDescData = protoimpl.X.CompressGZIP(file_chord_proto_rawDescData)
	})
	return file_chord_proto_rawDescData
}

var file_chord_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_chord_proto_goTypes = []interface{}{
	(*Node)(nil),                // 0: chord.Node
	(*ER)(nil),                  // 1: chord.ER
	(*ID)(nil),                  // 2: chord.ID
	(*GetRequest)(nil),          // 3: chord.GetRequest
	(*GetResponse)(nil),         // 4: chord.GetResponse
	(*SetRequest)(nil),          // 5: chord.SetRequest
	(*SetResponse)(nil),         // 6: chord.SetResponse
	(*DeleteRequest)(nil),       // 7: chord.DeleteRequest
	(*DeleteResponse)(nil),      // 8: chord.DeleteResponse
	(*MultiDeleteRequest)(nil),  // 9: chord.MultiDeleteRequest
	(*RequestKeysRequest)(nil),  // 10: chord.RequestKeysRequest
	(*KV)(nil),                  // 11: chord.KV
	(*RequestKeysResponse)(nil), // 12: chord.RequestKeysResponse
}
var file_chord_proto_depIdxs = []int32{
	11, // 0: chord.RequestKeysResponse.values:type_name -> chord.KV
	1,  // 1: chord.Chord.GetPredecessor:input_type -> chord.ER
	1,  // 2: chord.Chord.GetSuccessor:input_type -> chord.ER
	0,  // 3: chord.Chord.Notify:input_type -> chord.Node
	2,  // 4: chord.Chord.FindSuccessor:input_type -> chord.ID
	2,  // 5: chord.Chord.CheckPredecessor:input_type -> chord.ID
	0,  // 6: chord.Chord.SetPredecessor:input_type -> chord.Node
	0,  // 7: chord.Chord.SetSuccessor:input_type -> chord.Node
	3,  // 8: chord.Chord.XGet:input_type -> chord.GetRequest
	5,  // 9: chord.Chord.XSet:input_type -> chord.SetRequest
	7,  // 10: chord.Chord.XDelete:input_type -> chord.DeleteRequest
	9,  // 11: chord.Chord.XMultiDelete:input_type -> chord.MultiDeleteRequest
	10, // 12: chord.Chord.XRequestKeys:input_type -> chord.RequestKeysRequest
	0,  // 13: chord.Chord.GetPredecessor:output_type -> chord.Node
	0,  // 14: chord.Chord.GetSuccessor:output_type -> chord.Node
	1,  // 15: chord.Chord.Notify:output_type -> chord.ER
	0,  // 16: chord.Chord.FindSuccessor:output_type -> chord.Node
	1,  // 17: chord.Chord.CheckPredecessor:output_type -> chord.ER
	1,  // 18: chord.Chord.SetPredecessor:output_type -> chord.ER
	1,  // 19: chord.Chord.SetSuccessor:output_type -> chord.ER
	4,  // 20: chord.Chord.XGet:output_type -> chord.GetResponse
	6,  // 21: chord.Chord.XSet:output_type -> chord.SetResponse
	8,  // 22: chord.Chord.XDelete:output_type -> chord.DeleteResponse
	8,  // 23: chord.Chord.XMultiDelete:output_type -> chord.DeleteResponse
	12, // 24: chord.Chord.XRequestKeys:output_type -> chord.RequestKeysResponse
	13, // [13:25] is the sub-list for method output_type
	1,  // [1:13] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_chord_proto_init() }
func file_chord_proto_init() {
	if File_chord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ER); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestKeysRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KV); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestKeysResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chord_proto_goTypes,
		DependencyIndexes: file_chord_proto_depIdxs,
		MessageInfos:      file_chord_proto_msgTypes,
	}.Build()
	File_chord_proto = out.File
	file_chord_proto_rawDesc = nil
	file_chord_proto_goTypes = nil
	file_chord_proto_depIdxs = nil
}