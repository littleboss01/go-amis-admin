// 定义文件级别的Json标签复制
// @Tag("gorm", "autoUpdateTime:nano")

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.13.0
// source: mailRules.proto

package mail

import (
	_ "github.com/go-home-admin/go-admin/generate/proto/common/http"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_mailRules_proto protoreflect.FileDescriptor

var file_mailRules_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x1f, 0x0a, 0x09, 0x4d, 0x61,
	0x69, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x12, 0xc2, 0xf3, 0x18, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0xca, 0xf3, 0x18, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x68, 0x6f, 0x6d,
	0x65, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6d, 0x61, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mailRules_proto_goTypes = []any{}
var file_mailRules_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mailRules_proto_init() }
func file_mailRules_proto_init() {
	if File_mailRules_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mailRules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mailRules_proto_goTypes,
		DependencyIndexes: file_mailRules_proto_depIdxs,
	}.Build()
	File_mailRules_proto = out.File
	file_mailRules_proto_rawDesc = nil
	file_mailRules_proto_goTypes = nil
	file_mailRules_proto_depIdxs = nil
}
