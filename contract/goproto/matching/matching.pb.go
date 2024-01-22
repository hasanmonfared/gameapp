// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: contract/protobuf/matching/matching.proto

package matching

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

type MatchUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	UserIds  []uint64 `protobuf:"varint,2,rep,packed,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *MatchUsers) Reset() {
	*x = MatchUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_matching_matching_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchUsers) ProtoMessage() {}

func (x *MatchUsers) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_matching_matching_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchUsers.ProtoReflect.Descriptor instead.
func (*MatchUsers) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_matching_matching_proto_rawDescGZIP(), []int{0}
}

func (x *MatchUsers) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *MatchUsers) GetUserIds() []uint64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

var File_contract_protobuf_matching_matching_proto protoreflect.FileDescriptor

var file_contract_protobuf_matching_matching_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x43, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x42, 0x1b, 0x5a, 0x19, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_protobuf_matching_matching_proto_rawDescOnce sync.Once
	file_contract_protobuf_matching_matching_proto_rawDescData = file_contract_protobuf_matching_matching_proto_rawDesc
)

func file_contract_protobuf_matching_matching_proto_rawDescGZIP() []byte {
	file_contract_protobuf_matching_matching_proto_rawDescOnce.Do(func() {
		file_contract_protobuf_matching_matching_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_protobuf_matching_matching_proto_rawDescData)
	})
	return file_contract_protobuf_matching_matching_proto_rawDescData
}

var file_contract_protobuf_matching_matching_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contract_protobuf_matching_matching_proto_goTypes = []interface{}{
	(*MatchUsers)(nil), // 0: matching.MatchUsers
}
var file_contract_protobuf_matching_matching_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_contract_protobuf_matching_matching_proto_init() }
func file_contract_protobuf_matching_matching_proto_init() {
	if File_contract_protobuf_matching_matching_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_protobuf_matching_matching_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchUsers); i {
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
			RawDescriptor: file_contract_protobuf_matching_matching_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_protobuf_matching_matching_proto_goTypes,
		DependencyIndexes: file_contract_protobuf_matching_matching_proto_depIdxs,
		MessageInfos:      file_contract_protobuf_matching_matching_proto_msgTypes,
	}.Build()
	File_contract_protobuf_matching_matching_proto = out.File
	file_contract_protobuf_matching_matching_proto_rawDesc = nil
	file_contract_protobuf_matching_matching_proto_goTypes = nil
	file_contract_protobuf_matching_matching_proto_depIdxs = nil
}
