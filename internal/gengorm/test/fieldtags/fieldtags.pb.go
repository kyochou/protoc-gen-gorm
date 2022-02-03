// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: fieldtags/fieldtags.proto

package fieldtags

import (
	_ "github.com/complex64/protoc-gen-gorm/gormpb"
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

// Options iterates all common options to assert on the resulting GORM tags.
type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `gorm:"column:my_column"`
	Column string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	// `gorm:"not null"`
	NotNull string `protobuf:"bytes,6,opt,name=not_null,json=notNull,proto3" json:"not_null,omitempty"`
	// `gorm:"default:value"`
	Default string `protobuf:"bytes,7,opt,name=default,proto3" json:"default,omitempty"`
	// `gorm:"unique"`
	Unique string `protobuf:"bytes,8,opt,name=unique,proto3" json:"unique,omitempty"`
	// `gorm:"primaryKey"`
	PrimaryKey string `protobuf:"bytes,9,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
	// `gorm:"index"`
	DefaultIndex string `protobuf:"bytes,10,opt,name=default_index,json=defaultIndex,proto3" json:"default_index,omitempty"`
	// `gorm:"index:my_index"`
	NamedIndex string `protobuf:"bytes,101,opt,name=named_index,json=namedIndex,proto3" json:"named_index,omitempty"`
	// `gorm:"uniqueIndex"`
	UniqueDefaultIndex string `protobuf:"bytes,11,opt,name=unique_default_index,json=uniqueDefaultIndex,proto3" json:"unique_default_index,omitempty"`
	// `gorm:"uniqueIndex:my_unique_index"`
	UniqueNamedIndex string `protobuf:"bytes,111,opt,name=unique_named_index,json=uniqueNamedIndex,proto3" json:"unique_named_index,omitempty"`
	// `gorm:"autoCreateTime"`
	AutoCreateTime string `protobuf:"bytes,14,opt,name=auto_create_time,json=autoCreateTime,proto3" json:"auto_create_time,omitempty"`
	// `gorm:"autoUpdateTime"`
	AutoUpdateTime string `protobuf:"bytes,16,opt,name=auto_update_time,json=autoUpdateTime,proto3" json:"auto_update_time,omitempty"`
	// `gorm:"-"`
	Ignore string `protobuf:"bytes,22,opt,name=ignore,proto3" json:"ignore,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fieldtags_fieldtags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_fieldtags_fieldtags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_fieldtags_fieldtags_proto_rawDescGZIP(), []int{0}
}

func (x *Options) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *Options) GetNotNull() string {
	if x != nil {
		return x.NotNull
	}
	return ""
}

func (x *Options) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *Options) GetUnique() string {
	if x != nil {
		return x.Unique
	}
	return ""
}

func (x *Options) GetPrimaryKey() string {
	if x != nil {
		return x.PrimaryKey
	}
	return ""
}

func (x *Options) GetDefaultIndex() string {
	if x != nil {
		return x.DefaultIndex
	}
	return ""
}

func (x *Options) GetNamedIndex() string {
	if x != nil {
		return x.NamedIndex
	}
	return ""
}

func (x *Options) GetUniqueDefaultIndex() string {
	if x != nil {
		return x.UniqueDefaultIndex
	}
	return ""
}

func (x *Options) GetUniqueNamedIndex() string {
	if x != nil {
		return x.UniqueNamedIndex
	}
	return ""
}

func (x *Options) GetAutoCreateTime() string {
	if x != nil {
		return x.AutoCreateTime
	}
	return ""
}

func (x *Options) GetAutoUpdateTime() string {
	if x != nil {
		return x.AutoUpdateTime
	}
	return ""
}

func (x *Options) GetIgnore() string {
	if x != nil {
		return x.Ignore
	}
	return ""
}

// DenyOptions iterates all permissions to assert on the resulting tags.
type DenyOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `gorm:"-"`
	Ignore string `protobuf:"bytes,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// `gorm:"->:false;<-:update"`
	UpdateOnly string `protobuf:"bytes,2,opt,name=update_only,json=updateOnly,proto3" json:"update_only,omitempty"`
	// `gorm:"->:false;<-:create"`
	CreateOnly string `protobuf:"bytes,3,opt,name=create_only,json=createOnly,proto3" json:"create_only,omitempty"`
	// `gorm:"->:false;<-"`
	WriteOnly string `protobuf:"bytes,4,opt,name=write_only,json=writeOnly,proto3" json:"write_only,omitempty"`
	// `gorm:"->"`
	ReadOnly string `protobuf:"bytes,5,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	// `gorm:"<-:update"`
	ReadUpdate string `protobuf:"bytes,6,opt,name=read_update,json=readUpdate,proto3" json:"read_update,omitempty"`
	// `gorm:"<-:create"`
	ReadCreate string `protobuf:"bytes,7,opt,name=read_create,json=readCreate,proto3" json:"read_create,omitempty"`
}

func (x *DenyOptions) Reset() {
	*x = DenyOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fieldtags_fieldtags_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DenyOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenyOptions) ProtoMessage() {}

func (x *DenyOptions) ProtoReflect() protoreflect.Message {
	mi := &file_fieldtags_fieldtags_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenyOptions.ProtoReflect.Descriptor instead.
func (*DenyOptions) Descriptor() ([]byte, []int) {
	return file_fieldtags_fieldtags_proto_rawDescGZIP(), []int{1}
}

func (x *DenyOptions) GetIgnore() string {
	if x != nil {
		return x.Ignore
	}
	return ""
}

func (x *DenyOptions) GetUpdateOnly() string {
	if x != nil {
		return x.UpdateOnly
	}
	return ""
}

func (x *DenyOptions) GetCreateOnly() string {
	if x != nil {
		return x.CreateOnly
	}
	return ""
}

func (x *DenyOptions) GetWriteOnly() string {
	if x != nil {
		return x.WriteOnly
	}
	return ""
}

func (x *DenyOptions) GetReadOnly() string {
	if x != nil {
		return x.ReadOnly
	}
	return ""
}

func (x *DenyOptions) GetReadUpdate() string {
	if x != nil {
		return x.ReadUpdate
	}
	return ""
}

func (x *DenyOptions) GetReadCreate() string {
	if x != nil {
		return x.ReadCreate
	}
	return ""
}

var File_fieldtags_fieldtags_proto protoreflect.FileDescriptor

var file_fieldtags_fieldtags_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x74, 0x61, 0x67, 0x73, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x12, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x04, 0x0a, 0x07, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xb2, 0xbb, 0x18, 0x0b, 0x0a, 0x09, 0x6d, 0x79,
	0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12,
	0x21, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xb2, 0xbb, 0x18, 0x02, 0x30, 0x01, 0x52, 0x07, 0x6e, 0x6f, 0x74, 0x4e, 0x75,
	0x6c, 0x6c, 0x12, 0x25, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xb2, 0xbb, 0x18, 0x07, 0x3a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xb2, 0xbb, 0x18, 0x02, 0x40,
	0x01, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f,
	0xb2, 0xbb, 0x18, 0x0b, 0x0a, 0x09, 0x6d, 0x79, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52,
	0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x0d, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xb2, 0xbb, 0x18, 0x04, 0x52, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x31, 0x0a, 0x0b, 0x6e, 0x61,
	0x6d, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x10, 0xb2, 0xbb, 0x18, 0x0c, 0x52, 0x0a, 0x12, 0x08, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x3a, 0x0a,
	0x14, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xb2, 0xbb, 0x18,
	0x04, 0x5a, 0x02, 0x08, 0x01, 0x52, 0x12, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x45, 0x0a, 0x12, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x6f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xb2, 0xbb, 0x18, 0x13, 0x5a, 0x11, 0x12, 0x0f, 0x6d,
	0x79, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x10,
	0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x30, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xb2, 0xbb, 0x18, 0x02,
	0x70, 0x01, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x31, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xb2, 0xbb,
	0x18, 0x03, 0x80, 0x01, 0x01, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xb2, 0xbb, 0x18, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06,
	0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x3a, 0x06, 0xaa, 0xbb, 0x18, 0x02, 0x08, 0x01, 0x22, 0xc4,
	0x02, 0x0a, 0x0b, 0x44, 0x65, 0x6e, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25,
	0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d,
	0xb2, 0xbb, 0x18, 0x09, 0xba, 0x01, 0x06, 0x10, 0x01, 0x18, 0x01, 0x20, 0x01, 0x52, 0x06, 0x69,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xb2, 0xbb, 0x18, 0x07,
	0xba, 0x01, 0x04, 0x10, 0x01, 0x20, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x6e, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xb2, 0xbb, 0x18, 0x07, 0xba, 0x01,
	0x04, 0x18, 0x01, 0x20, 0x01, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c,
	0x79, 0x12, 0x28, 0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xb2, 0xbb, 0x18, 0x05, 0xba, 0x01, 0x02, 0x20, 0x01,
	0x52, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x09, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b,
	0xb2, 0xbb, 0x18, 0x07, 0xba, 0x01, 0x04, 0x10, 0x01, 0x18, 0x01, 0x52, 0x08, 0x72, 0x65, 0x61,
	0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xb2, 0xbb, 0x18, 0x05,
	0xba, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xb2, 0xbb, 0x18, 0x05, 0xba, 0x01, 0x02, 0x18,
	0x01, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x06, 0xaa,
	0xbb, 0x18, 0x02, 0x08, 0x01, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x36, 0x34, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6d,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72,
	0x6d, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x74, 0x61, 0x67, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fieldtags_fieldtags_proto_rawDescOnce sync.Once
	file_fieldtags_fieldtags_proto_rawDescData = file_fieldtags_fieldtags_proto_rawDesc
)

func file_fieldtags_fieldtags_proto_rawDescGZIP() []byte {
	file_fieldtags_fieldtags_proto_rawDescOnce.Do(func() {
		file_fieldtags_fieldtags_proto_rawDescData = protoimpl.X.CompressGZIP(file_fieldtags_fieldtags_proto_rawDescData)
	})
	return file_fieldtags_fieldtags_proto_rawDescData
}

var file_fieldtags_fieldtags_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fieldtags_fieldtags_proto_goTypes = []interface{}{
	(*Options)(nil),     // 0: fieldtags.Options
	(*DenyOptions)(nil), // 1: fieldtags.DenyOptions
}
var file_fieldtags_fieldtags_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fieldtags_fieldtags_proto_init() }
func file_fieldtags_fieldtags_proto_init() {
	if File_fieldtags_fieldtags_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fieldtags_fieldtags_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
		file_fieldtags_fieldtags_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DenyOptions); i {
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
			RawDescriptor: file_fieldtags_fieldtags_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fieldtags_fieldtags_proto_goTypes,
		DependencyIndexes: file_fieldtags_fieldtags_proto_depIdxs,
		MessageInfos:      file_fieldtags_fieldtags_proto_msgTypes,
	}.Build()
	File_fieldtags_fieldtags_proto = out.File
	file_fieldtags_fieldtags_proto_rawDesc = nil
	file_fieldtags_fieldtags_proto_goTypes = nil
	file_fieldtags_fieldtags_proto_depIdxs = nil
}
