// 定义文件级别的Json标签复制
// @Tag("gorm", "autoUpdateTime:nano")

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.13.0
// source: admin_auth.proto

package admin

import (
	_ "github.com/go-home-admin/go-admin/generate/proto/common/http"
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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" form:"username" gorm:"autoUpdateTime:nano"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" form:"password" gorm:"autoUpdateTime:nano"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_admin_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" form:"token" gorm:"autoUpdateTime:nano"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_admin_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type MyMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MyMenuRequest) Reset() {
	*x = MyMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyMenuRequest) ProtoMessage() {}

func (x *MyMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyMenuRequest.ProtoReflect.Descriptor instead.
func (*MyMenuRequest) Descriptor() ([]byte, []int) {
	return file_admin_auth_proto_rawDescGZIP(), []int{2}
}

type MyMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menus []*MenuInfo `protobuf:"bytes,1,rep,name=menus,proto3" json:"menus,omitempty" form:"menus" gorm:"autoUpdateTime:nano"`
}

func (x *MyMenuResponse) Reset() {
	*x = MyMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyMenuResponse) ProtoMessage() {}

func (x *MyMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyMenuResponse.ProtoReflect.Descriptor instead.
func (*MyMenuResponse) Descriptor() ([]byte, []int) {
	return file_admin_auth_proto_rawDescGZIP(), []int{3}
}

func (x *MyMenuResponse) GetMenus() []*MenuInfo {
	if x != nil {
		return x.Menus
	}
	return nil
}

type MenuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" gorm:"autoUpdateTime:nano"`
	ParentId  uint32      `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty" form:"parent_id" gorm:"autoUpdateTime:nano"`
	Path      string      `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty" form:"path" gorm:"autoUpdateTime:nano"`
	Hidden    bool        `protobuf:"varint,4,opt,name=hidden,proto3" json:"hidden,omitempty" form:"hidden" gorm:"autoUpdateTime:nano"`
	Name      string      `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty" form:"name" gorm:"autoUpdateTime:nano"`
	Redirect  string      `protobuf:"bytes,6,opt,name=redirect,proto3" json:"redirect,omitempty" form:"redirect" gorm:"autoUpdateTime:nano"`
	Component string      `protobuf:"bytes,7,opt,name=component,proto3" json:"component,omitempty" form:"component" gorm:"autoUpdateTime:nano"`
	Meta      *Meta       `protobuf:"bytes,8,opt,name=meta,proto3" json:"meta,omitempty" form:"meta" gorm:"autoUpdateTime:nano"`
	Children  []*MenuInfo `protobuf:"bytes,9,rep,name=children,proto3" json:"children,omitempty" form:"children" gorm:"autoUpdateTime:nano"`
}

func (x *MenuInfo) Reset() {
	*x = MenuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuInfo) ProtoMessage() {}

func (x *MenuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_admin_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuInfo.ProtoReflect.Descriptor instead.
func (*MenuInfo) Descriptor() ([]byte, []int) {
	return file_admin_auth_proto_rawDescGZIP(), []int{4}
}

func (x *MenuInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MenuInfo) GetParentId() uint32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *MenuInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MenuInfo) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *MenuInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuInfo) GetRedirect() string {
	if x != nil {
		return x.Redirect
	}
	return ""
}

func (x *MenuInfo) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *MenuInfo) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *MenuInfo) GetChildren() []*MenuInfo {
	if x != nil {
		return x.Children
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" form:"title" gorm:"autoUpdateTime:nano"`
	Icon  string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty" form:"icon" gorm:"autoUpdateTime:nano"`
	Amis  string `protobuf:"bytes,3,opt,name=amis,proto3" json:"amis,omitempty" form:"amis" gorm:"autoUpdateTime:nano"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_admin_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_admin_auth_proto_rawDescGZIP(), []int{5}
}

func (x *Meta) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Meta) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Meta) GetAmis() string {
	if x != nil {
		return x.Amis
	}
	return ""
}

var File_admin_auth_proto protoreflect.FileDescriptor

var file_admin_auth_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x0f, 0x0a, 0x0d, 0x4d,
	0x79, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x0e,
	0x4d, 0x79, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x6d, 0x65, 0x6e, 0x75, 0x73, 0x22, 0xff, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x44, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x6d, 0x69,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x6d, 0x69, 0x73, 0x32, 0xf1, 0x01,
	0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x12, 0x43, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0f, 0x92, 0xb5, 0x18, 0x0b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x45, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x92, 0xb5, 0x18, 0x0c, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x46, 0x0a, 0x06, 0x4d, 0x79, 0x4d, 0x65, 0x6e,
	0x75, 0x12, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x79, 0x4d, 0x65, 0x6e, 0x75,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x4d, 0x79, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f,
	0x82, 0xb5, 0x18, 0x0b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x1a,
	0x10, 0xc2, 0xf3, 0x18, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2d, 0x68, 0x6f, 0x6d, 0x65, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x6f,
	0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_admin_auth_proto_rawDescOnce sync.Once
	file_admin_auth_proto_rawDescData = file_admin_auth_proto_rawDesc
)

func file_admin_auth_proto_rawDescGZIP() []byte {
	file_admin_auth_proto_rawDescOnce.Do(func() {
		file_admin_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_auth_proto_rawDescData)
	})
	return file_admin_auth_proto_rawDescData
}

var file_admin_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_admin_auth_proto_goTypes = []any{
	(*LoginRequest)(nil),   // 0: admin.LoginRequest
	(*LoginResponse)(nil),  // 1: admin.LoginResponse
	(*MyMenuRequest)(nil),  // 2: admin.MyMenuRequest
	(*MyMenuResponse)(nil), // 3: admin.MyMenuResponse
	(*MenuInfo)(nil),       // 4: admin.MenuInfo
	(*Meta)(nil),           // 5: admin.Meta
}
var file_admin_auth_proto_depIdxs = []int32{
	4, // 0: admin.MyMenuResponse.menus:type_name -> admin.MenuInfo
	5, // 1: admin.MenuInfo.meta:type_name -> admin.Meta
	4, // 2: admin.MenuInfo.children:type_name -> admin.MenuInfo
	0, // 3: admin.AdminAuth.Login:input_type -> admin.LoginRequest
	0, // 4: admin.AdminAuth.Logout:input_type -> admin.LoginRequest
	2, // 5: admin.AdminAuth.MyMenu:input_type -> admin.MyMenuRequest
	1, // 6: admin.AdminAuth.Login:output_type -> admin.LoginResponse
	1, // 7: admin.AdminAuth.Logout:output_type -> admin.LoginResponse
	3, // 8: admin.AdminAuth.MyMenu:output_type -> admin.MyMenuResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_admin_auth_proto_init() }
func file_admin_auth_proto_init() {
	if File_admin_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_auth_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LoginRequest); i {
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
		file_admin_auth_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LoginResponse); i {
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
		file_admin_auth_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MyMenuRequest); i {
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
		file_admin_auth_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MyMenuResponse); i {
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
		file_admin_auth_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MenuInfo); i {
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
		file_admin_auth_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Meta); i {
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
			RawDescriptor: file_admin_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_auth_proto_goTypes,
		DependencyIndexes: file_admin_auth_proto_depIdxs,
		MessageInfos:      file_admin_auth_proto_msgTypes,
	}.Build()
	File_admin_auth_proto = out.File
	file_admin_auth_proto_rawDesc = nil
	file_admin_auth_proto_goTypes = nil
	file_admin_auth_proto_depIdxs = nil
}
