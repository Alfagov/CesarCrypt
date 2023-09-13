// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: protoStructs/store.proto

package protoStructs

import (
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

type KeyRing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uid     string            `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	RootKey string            `protobuf:"bytes,3,opt,name=rootKey,proto3" json:"rootKey,omitempty"`
	Keys    map[string]string `protobuf:"bytes,4,rep,name=keys,proto3" json:"keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *KeyRing) Reset() {
	*x = KeyRing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoStructs_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRing) ProtoMessage() {}

func (x *KeyRing) ProtoReflect() protoreflect.Message {
	mi := &file_protoStructs_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRing.ProtoReflect.Descriptor instead.
func (*KeyRing) Descriptor() ([]byte, []int) {
	return file_protoStructs_store_proto_rawDescGZIP(), []int{0}
}

func (x *KeyRing) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KeyRing) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *KeyRing) GetRootKey() string {
	if x != nil {
		return x.RootKey
	}
	return ""
}

func (x *KeyRing) GetKeys() map[string]string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uid            string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Version        int32  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	KeyringVersion int32  `protobuf:"varint,4,opt,name=keyring_version,json=keyringVersion,proto3" json:"keyring_version,omitempty"`
	Shards         int32  `protobuf:"varint,5,opt,name=shards,proto3" json:"shards,omitempty"`
	Threshold      int32  `protobuf:"varint,6,opt,name=threshold,proto3" json:"threshold,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoStructs_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_protoStructs_store_proto_msgTypes[1]
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
	return file_protoStructs_store_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Config) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Config) GetKeyringVersion() int32 {
	if x != nil {
		return x.KeyringVersion
	}
	return 0
}

func (x *Config) GetShards() int32 {
	if x != nil {
		return x.Shards
	}
	return 0
}

func (x *Config) GetThreshold() int32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

var File_protoStructs_store_proto protoreflect.FileDescriptor

var file_protoStructs_store_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x07, 0x4b,
	0x65, 0x79, 0x52, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x6f, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x6f, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x69, 0x6e, 0x67, 0x2e, 0x4b,
	0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x1a, 0x37,
	0x0a, 0x09, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa7, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6b, 0x65, 0x79,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoStructs_store_proto_rawDescOnce sync.Once
	file_protoStructs_store_proto_rawDescData = file_protoStructs_store_proto_rawDesc
)

func file_protoStructs_store_proto_rawDescGZIP() []byte {
	file_protoStructs_store_proto_rawDescOnce.Do(func() {
		file_protoStructs_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoStructs_store_proto_rawDescData)
	})
	return file_protoStructs_store_proto_rawDescData
}

var file_protoStructs_store_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protoStructs_store_proto_goTypes = []interface{}{
	(*KeyRing)(nil), // 0: KeyRing
	(*Config)(nil),  // 1: Config
	nil,             // 2: KeyRing.KeysEntry
}
var file_protoStructs_store_proto_depIdxs = []int32{
	2, // 0: KeyRing.keys:type_name -> KeyRing.KeysEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protoStructs_store_proto_init() }
func file_protoStructs_store_proto_init() {
	if File_protoStructs_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoStructs_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRing); i {
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
		file_protoStructs_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protoStructs_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protoStructs_store_proto_goTypes,
		DependencyIndexes: file_protoStructs_store_proto_depIdxs,
		MessageInfos:      file_protoStructs_store_proto_msgTypes,
	}.Build()
	File_protoStructs_store_proto = out.File
	file_protoStructs_store_proto_rawDesc = nil
	file_protoStructs_store_proto_goTypes = nil
	file_protoStructs_store_proto_depIdxs = nil
}
