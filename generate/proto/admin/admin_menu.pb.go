// 定义文件级别的Json标签复制
// @Tag("gorm", "autoUpdateTime:nano")

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.13.0
// source: admin_menu.proto

package admin

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

var File_admin_menu_proto protoreflect.FileDescriptor

var file_admin_menu_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x1f, 0x0a, 0x09,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4d, 0x65, 0x6e, 0x75, 0x1a, 0x12, 0xc2, 0xf3, 0x18, 0x05, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0xca, 0xf3, 0x18, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x42, 0x38, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x68,
	0x6f, 0x6d, 0x65, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_admin_menu_proto_goTypes = []any{}
var file_admin_menu_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_menu_proto_init() }
func file_admin_menu_proto_init() {
	if File_admin_menu_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_menu_proto_goTypes,
		DependencyIndexes: file_admin_menu_proto_depIdxs,
	}.Build()
	File_admin_menu_proto = out.File
	file_admin_menu_proto_rawDesc = nil
	file_admin_menu_proto_goTypes = nil
	file_admin_menu_proto_depIdxs = nil
}
