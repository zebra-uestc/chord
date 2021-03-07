// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: bridge.proto

package bridge

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

type DhtStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DhtStatus) Reset() {
	*x = DhtStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DhtStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DhtStatus) ProtoMessage() {}

func (x *DhtStatus) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DhtStatus.ProtoReflect.Descriptor instead.
func (*DhtStatus) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastBlock   *Block `protobuf:"bytes,1,opt,name=lastBlock,proto3" json:"lastBlock,omitempty"`
	OrdererAddr string `protobuf:"bytes,2,opt,name=ordererAddr,proto3" json:"ordererAddr,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetLastBlock() *Block {
	if x != nil {
		return x.LastBlock
	}
	return nil
}

func (x *Config) GetOrdererAddr() string {
	if x != nil {
		return x.OrdererAddr
	}
	return ""
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigSeq uint64    `protobuf:"varint,1,opt,name=configSeq,proto3" json:"configSeq,omitempty"`
	NormalMsg *Envelope `protobuf:"bytes,2,opt,name=normalMsg,proto3" json:"normalMsg,omitempty"`
	ConfigMsg *Envelope `protobuf:"bytes,3,opt,name=configMsg,proto3" json:"configMsg,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{2}
}

func (x *Msg) GetConfigSeq() uint64 {
	if x != nil {
		return x.ConfigSeq
	}
	return 0
}

func (x *Msg) GetNormalMsg() *Envelope {
	if x != nil {
		return x.NormalMsg
	}
	return nil
}

func (x *Msg) GetConfigMsg() *Envelope {
	if x != nil {
		return x.ConfigMsg
	}
	return nil
}

var File_bridge_proto protoreflect.FileDescriptor

var file_bridge_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x44, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x57, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x22, 0x83, 0x01, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x71, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x71,
	0x12, 0x2e, 0x0a, 0x09, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x09, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x4d, 0x73, 0x67,
	0x12, 0x2e, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x73, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x73, 0x67,
	0x32, 0x40, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x72,
	0x12, 0x30, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0d,
	0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x11, 0x2e,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x44, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x32, 0x6d, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x72,
	0x12, 0x2c, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x0b, 0x2e, 0x62,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x1a, 0x11, 0x2e, 0x62, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x2e, 0x44, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x0a, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x11, 0x2e, 0x62,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x44, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a,
	0x0e, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x7a, 0x65, 0x62, 0x72, 0x61, 0x2d, 0x75, 0x65, 0x73, 0x74, 0x63, 0x2f, 0x63, 0x68, 0x6f, 0x72,
	0x64, 0x2f, 0x64, 0x68, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bridge_proto_rawDescOnce sync.Once
	file_bridge_proto_rawDescData = file_bridge_proto_rawDesc
)

func file_bridge_proto_rawDescGZIP() []byte {
	file_bridge_proto_rawDescOnce.Do(func() {
		file_bridge_proto_rawDescData = protoimpl.X.CompressGZIP(file_bridge_proto_rawDescData)
	})
	return file_bridge_proto_rawDescData
}

var file_bridge_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bridge_proto_goTypes = []interface{}{
	(*DhtStatus)(nil), // 0: bridge.DhtStatus
	(*Config)(nil),    // 1: bridge.Config
	(*Msg)(nil),       // 2: bridge.Msg
	(*Block)(nil),     // 3: bridge.Block
	(*Envelope)(nil),  // 4: bridge.Envelope
}
var file_bridge_proto_depIdxs = []int32{
	3, // 0: bridge.Config.lastBlock:type_name -> bridge.Block
	4, // 1: bridge.Msg.normalMsg:type_name -> bridge.Envelope
	4, // 2: bridge.Msg.configMsg:type_name -> bridge.Envelope
	3, // 3: bridge.BlockTranser.TransBlock:input_type -> bridge.Block
	2, // 4: bridge.MsgTranser.TransMsg:input_type -> bridge.Msg
	0, // 5: bridge.MsgTranser.LoadConfig:input_type -> bridge.DhtStatus
	0, // 6: bridge.BlockTranser.TransBlock:output_type -> bridge.DhtStatus
	0, // 7: bridge.MsgTranser.TransMsg:output_type -> bridge.DhtStatus
	1, // 8: bridge.MsgTranser.LoadConfig:output_type -> bridge.Config
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bridge_proto_init() }
func file_bridge_proto_init() {
	if File_bridge_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bridge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DhtStatus); i {
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
		file_bridge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_bridge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
			RawDescriptor: file_bridge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_bridge_proto_goTypes,
		DependencyIndexes: file_bridge_proto_depIdxs,
		MessageInfos:      file_bridge_proto_msgTypes,
	}.Build()
	File_bridge_proto = out.File
	file_bridge_proto_rawDesc = nil
	file_bridge_proto_goTypes = nil
	file_bridge_proto_depIdxs = nil
}
