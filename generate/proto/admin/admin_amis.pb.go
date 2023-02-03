// 定义文件级别的Json标签复制
// @Tag("gorm", "autoUpdateTime:nano")

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: admin_amis.proto

package admin

import (
	_ "github.com/go-home-admin/go-admin/generate/proto/common/http"
	protobuf "github.com/go-home-admin/home/protobuf"
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

type DemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DemoRequest) Reset() {
	*x = DemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_amis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoRequest) ProtoMessage() {}

func (x *DemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_amis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoRequest.ProtoReflect.Descriptor instead.
func (*DemoRequest) Descriptor() ([]byte, []int) {
	return file_admin_amis_proto_rawDescGZIP(), []int{0}
}

type DemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" form:"title"`
	Body  *protobuf.Any `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty" form:"body"`
}

func (x *DemoResponse) Reset() {
	*x = DemoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_amis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoResponse) ProtoMessage() {}

func (x *DemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_amis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoResponse.ProtoReflect.Descriptor instead.
func (*DemoResponse) Descriptor() ([]byte, []int) {
	return file_admin_amis_proto_rawDescGZIP(), []int{1}
}

func (x *DemoResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DemoResponse) GetBody() *protobuf.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_admin_amis_proto protoreflect.FileDescriptor

var file_admin_amis_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x6d, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x44, 0x65, 0x6d, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x0c, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x6f, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x32, 0x52, 0x0a, 0x04, 0x41, 0x6d, 0x69, 0x73, 0x12, 0x3f, 0x0a, 0x04,
	0x44, 0x65, 0x6d, 0x6f, 0x12, 0x12, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6d,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x82,
	0xb5, 0x18, 0x0a, 0x2f, 0x61, 0x6d, 0x69, 0x73, 0x2f, 0x74, 0x61, 0x62, 0x73, 0x1a, 0x09, 0xc2,
	0xf3, 0x18, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x68, 0x6f, 0x6d, 0x65, 0x2d, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_amis_proto_rawDescOnce sync.Once
	file_admin_amis_proto_rawDescData = file_admin_amis_proto_rawDesc
)

func file_admin_amis_proto_rawDescGZIP() []byte {
	file_admin_amis_proto_rawDescOnce.Do(func() {
		file_admin_amis_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_amis_proto_rawDescData)
	})
	return file_admin_amis_proto_rawDescData
}

var file_admin_amis_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_admin_amis_proto_goTypes = []interface{}{
	(*DemoRequest)(nil),  // 0: admin.DemoRequest
	(*DemoResponse)(nil), // 1: admin.DemoResponse
	(*protobuf.Any)(nil), // 2: home.protobuf.Any
}
var file_admin_amis_proto_depIdxs = []int32{
	2, // 0: admin.DemoResponse.body:type_name -> home.protobuf.Any
	0, // 1: admin.Amis.Demo:input_type -> admin.DemoRequest
	1, // 2: admin.Amis.Demo:output_type -> admin.DemoResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_admin_amis_proto_init() }
func file_admin_amis_proto_init() {
	if File_admin_amis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_amis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoRequest); i {
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
		file_admin_amis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoResponse); i {
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
			RawDescriptor: file_admin_amis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_amis_proto_goTypes,
		DependencyIndexes: file_admin_amis_proto_depIdxs,
		MessageInfos:      file_admin_amis_proto_msgTypes,
	}.Build()
	File_admin_amis_proto = out.File
	file_admin_amis_proto_rawDesc = nil
	file_admin_amis_proto_goTypes = nil
	file_admin_amis_proto_depIdxs = nil
}
