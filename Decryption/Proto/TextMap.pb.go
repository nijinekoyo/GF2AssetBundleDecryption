// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.1
// source: Decryption/Proto/pb/TextMap.proto

package Proto

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

type TextMapTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*TextMap `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *TextMapTable) Reset() {
	*x = TextMapTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Decryption_Proto_pb_TextMap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextMapTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextMapTable) ProtoMessage() {}

func (x *TextMapTable) ProtoReflect() protoreflect.Message {
	mi := &file_Decryption_Proto_pb_TextMap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextMapTable.ProtoReflect.Descriptor instead.
func (*TextMapTable) Descriptor() ([]byte, []int) {
	return file_Decryption_Proto_pb_TextMap_proto_rawDescGZIP(), []int{0}
}

func (x *TextMapTable) GetData() []*TextMap {
	if x != nil {
		return x.Data
	}
	return nil
}

type TextMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *TextMap) Reset() {
	*x = TextMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Decryption_Proto_pb_TextMap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextMap) ProtoMessage() {}

func (x *TextMap) ProtoReflect() protoreflect.Message {
	mi := &file_Decryption_Proto_pb_TextMap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextMap.ProtoReflect.Descriptor instead.
func (*TextMap) Descriptor() ([]byte, []int) {
	return file_Decryption_Proto_pb_TextMap_proto_rawDescGZIP(), []int{1}
}

func (x *TextMap) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TextMap) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_Decryption_Proto_pb_TextMap_proto protoreflect.FileDescriptor

var file_Decryption_Proto_pb_TextMap_proto_rawDesc = []byte{
	0x0a, 0x21, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0c, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x70, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x33, 0x0a, 0x07, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x2f, 0x44, 0x65, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Decryption_Proto_pb_TextMap_proto_rawDescOnce sync.Once
	file_Decryption_Proto_pb_TextMap_proto_rawDescData = file_Decryption_Proto_pb_TextMap_proto_rawDesc
)

func file_Decryption_Proto_pb_TextMap_proto_rawDescGZIP() []byte {
	file_Decryption_Proto_pb_TextMap_proto_rawDescOnce.Do(func() {
		file_Decryption_Proto_pb_TextMap_proto_rawDescData = protoimpl.X.CompressGZIP(file_Decryption_Proto_pb_TextMap_proto_rawDescData)
	})
	return file_Decryption_Proto_pb_TextMap_proto_rawDescData
}

var file_Decryption_Proto_pb_TextMap_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Decryption_Proto_pb_TextMap_proto_goTypes = []interface{}{
	(*TextMapTable)(nil), // 0: TextMapTable
	(*TextMap)(nil),      // 1: TextMap
}
var file_Decryption_Proto_pb_TextMap_proto_depIdxs = []int32{
	1, // 0: TextMapTable.Data:type_name -> TextMap
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Decryption_Proto_pb_TextMap_proto_init() }
func file_Decryption_Proto_pb_TextMap_proto_init() {
	if File_Decryption_Proto_pb_TextMap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Decryption_Proto_pb_TextMap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextMapTable); i {
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
		file_Decryption_Proto_pb_TextMap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextMap); i {
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
			RawDescriptor: file_Decryption_Proto_pb_TextMap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Decryption_Proto_pb_TextMap_proto_goTypes,
		DependencyIndexes: file_Decryption_Proto_pb_TextMap_proto_depIdxs,
		MessageInfos:      file_Decryption_Proto_pb_TextMap_proto_msgTypes,
	}.Build()
	File_Decryption_Proto_pb_TextMap_proto = out.File
	file_Decryption_Proto_pb_TextMap_proto_rawDesc = nil
	file_Decryption_Proto_pb_TextMap_proto_goTypes = nil
	file_Decryption_Proto_pb_TextMap_proto_depIdxs = nil
}
