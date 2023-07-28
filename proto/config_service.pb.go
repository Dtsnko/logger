// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: proto/config_service.proto

package proto

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

type Configs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configs []*Config `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *Configs) Reset() {
	*x = Configs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configs) ProtoMessage() {}

func (x *Configs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configs.ProtoReflect.Descriptor instead.
func (*Configs) Descriptor() ([]byte, []int) {
	return file_proto_config_service_proto_rawDescGZIP(), []int{0}
}

func (x *Configs) GetConfigs() []*Config {
	if x != nil {
		return x.Configs
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ObjectId    int32  `protobuf:"varint,2,opt,name=objectId,proto3" json:"objectId,omitempty"`
	Enabled     bool   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DaysToStore int32  `protobuf:"varint,4,opt,name=daysToStore,proto3" json:"daysToStore,omitempty"`
	Period      string `protobuf:"bytes,5,opt,name=period,proto3" json:"period,omitempty"`
	StorageId   int32  `protobuf:"varint,6,opt,name=storageId,proto3" json:"storageId,omitempty"`
	DomainId    int32  `protobuf:"varint,7,opt,name=domainId,proto3" json:"domainId,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_service_proto_msgTypes[1]
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
	return file_proto_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Config) GetObjectId() int32 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *Config) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Config) GetDaysToStore() int32 {
	if x != nil {
		return x.DaysToStore
	}
	return 0
}

func (x *Config) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *Config) GetStorageId() int32 {
	if x != nil {
		return x.StorageId
	}
	return 0
}

func (x *Config) GetDomainId() int32 {
	if x != nil {
		return x.DomainId
	}
	return 0
}

type GetConfigByObjectIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId int32 `protobuf:"varint,1,opt,name=objectId,proto3" json:"objectId,omitempty"`
	DomainId int32 `protobuf:"varint,2,opt,name=domainId,proto3" json:"domainId,omitempty"`
}

func (x *GetConfigByObjectIdRequest) Reset() {
	*x = GetConfigByObjectIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigByObjectIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigByObjectIdRequest) ProtoMessage() {}

func (x *GetConfigByObjectIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigByObjectIdRequest.ProtoReflect.Descriptor instead.
func (*GetConfigByObjectIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_config_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetConfigByObjectIdRequest) GetObjectId() int32 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *GetConfigByObjectIdRequest) GetDomainId() int32 {
	if x != nil {
		return x.DomainId
	}
	return 0
}

type GetConfigByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //  int32 domainId = 8;
}

func (x *GetConfigByIdRequest) Reset() {
	*x = GetConfigByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigByIdRequest) ProtoMessage() {}

func (x *GetConfigByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigByIdRequest.ProtoReflect.Descriptor instead.
func (*GetConfigByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_config_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetConfigByIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAllConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size   int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q      string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort   string   `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields []string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *GetAllConfigsRequest) Reset() {
	*x = GetAllConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllConfigsRequest) ProtoMessage() {}

func (x *GetAllConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllConfigsRequest.ProtoReflect.Descriptor instead.
func (*GetAllConfigsRequest) Descriptor() ([]byte, []int) {
	return file_proto_config_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllConfigsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllConfigsRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetAllConfigsRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *GetAllConfigsRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *GetAllConfigsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type UpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId    int32  `protobuf:"varint,1,opt,name=objectId,proto3" json:"objectId,omitempty"`
	Enabled     bool   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DaysToStore int32  `protobuf:"varint,3,opt,name=daysToStore,proto3" json:"daysToStore,omitempty"`
	Period      string `protobuf:"bytes,4,opt,name=period,proto3" json:"period,omitempty"`
	StorageId   int32  `protobuf:"varint,5,opt,name=storageId,proto3" json:"storageId,omitempty"`
}

func (x *UpdateConfigRequest) Reset() {
	*x = UpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigRequest) ProtoMessage() {}

func (x *UpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_config_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateConfigRequest) GetObjectId() int32 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *UpdateConfigRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *UpdateConfigRequest) GetDaysToStore() int32 {
	if x != nil {
		return x.DaysToStore
	}
	return 0
}

func (x *UpdateConfigRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *UpdateConfigRequest) GetStorageId() int32 {
	if x != nil {
		return x.StorageId
	}
	return 0
}

type InsertConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId    int32  `protobuf:"varint,1,opt,name=objectId,proto3" json:"objectId,omitempty"`
	Enabled     bool   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DaysToStore int32  `protobuf:"varint,3,opt,name=daysToStore,proto3" json:"daysToStore,omitempty"`
	Period      string `protobuf:"bytes,4,opt,name=period,proto3" json:"period,omitempty"`
	StorageId   int32  `protobuf:"varint,5,opt,name=storageId,proto3" json:"storageId,omitempty"`
}

func (x *InsertConfigRequest) Reset() {
	*x = InsertConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertConfigRequest) ProtoMessage() {}

func (x *InsertConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertConfigRequest.ProtoReflect.Descriptor instead.
func (*InsertConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_config_service_proto_rawDescGZIP(), []int{6}
}

func (x *InsertConfigRequest) GetObjectId() int32 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *InsertConfigRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *InsertConfigRequest) GetDaysToStore() int32 {
	if x != nil {
		return x.DaysToStore
	}
	return 0
}

func (x *InsertConfigRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *InsertConfigRequest) GetStorageId() int32 {
	if x != nil {
		return x.StorageId
	}
	return 0
}

var File_proto_config_service_proto protoreflect.FileDescriptor

var file_proto_config_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x65,
	0x62, 0x69, 0x74, 0x65, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74,
	0x65, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0xc2, 0x01, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61,
	0x79, 0x73, 0x54, 0x6f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x64, 0x61, 0x79, 0x73, 0x54, 0x6f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x54,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x78, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x0c, 0x0a, 0x01,
	0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x54, 0x6f, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x54,
	0x6f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0xa3, 0x01, 0x0a,
	0x13, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61,
	0x79, 0x73, 0x54, 0x6f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x64, 0x61, 0x79, 0x73, 0x54, 0x6f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x32, 0xad, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x5f, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x77, 0x65, 0x62, 0x69,
	0x74, 0x65, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x23, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74,
	0x65, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x00, 0x12, 0x5b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x2e, 0x77, 0x65, 0x62, 0x69,
	0x74, 0x65, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12,
	0x4f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x24, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c,
	0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x12, 0x24, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65,
	0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_config_service_proto_rawDescOnce sync.Once
	file_proto_config_service_proto_rawDescData = file_proto_config_service_proto_rawDesc
)

func file_proto_config_service_proto_rawDescGZIP() []byte {
	file_proto_config_service_proto_rawDescOnce.Do(func() {
		file_proto_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_config_service_proto_rawDescData)
	})
	return file_proto_config_service_proto_rawDescData
}

var file_proto_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_config_service_proto_goTypes = []interface{}{
	(*Configs)(nil),                    // 0: webitel_logger.Configs
	(*Config)(nil),                     // 1: webitel_logger.Config
	(*GetConfigByObjectIdRequest)(nil), // 2: webitel_logger.GetConfigByObjectIdRequest
	(*GetConfigByIdRequest)(nil),       // 3: webitel_logger.GetConfigByIdRequest
	(*GetAllConfigsRequest)(nil),       // 4: webitel_logger.GetAllConfigsRequest
	(*UpdateConfigRequest)(nil),        // 5: webitel_logger.UpdateConfigRequest
	(*InsertConfigRequest)(nil),        // 6: webitel_logger.InsertConfigRequest
}
var file_proto_config_service_proto_depIdxs = []int32{
	1, // 0: webitel_logger.Configs.configs:type_name -> webitel_logger.Config
	5, // 1: webitel_logger.ConfigService.UpdateConfig:input_type -> webitel_logger.UpdateConfigRequest
	6, // 2: webitel_logger.ConfigService.InsertConfig:input_type -> webitel_logger.InsertConfigRequest
	2, // 3: webitel_logger.ConfigService.GetConfigByObjectId:input_type -> webitel_logger.GetConfigByObjectIdRequest
	3, // 4: webitel_logger.ConfigService.GetConfigById:input_type -> webitel_logger.GetConfigByIdRequest
	4, // 5: webitel_logger.ConfigService.GetAllConfigs:input_type -> webitel_logger.GetAllConfigsRequest
	1, // 6: webitel_logger.ConfigService.UpdateConfig:output_type -> webitel_logger.Config
	1, // 7: webitel_logger.ConfigService.InsertConfig:output_type -> webitel_logger.Config
	1, // 8: webitel_logger.ConfigService.GetConfigByObjectId:output_type -> webitel_logger.Config
	1, // 9: webitel_logger.ConfigService.GetConfigById:output_type -> webitel_logger.Config
	0, // 10: webitel_logger.ConfigService.GetAllConfigs:output_type -> webitel_logger.Configs
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_config_service_proto_init() }
func file_proto_config_service_proto_init() {
	if File_proto_config_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_config_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configs); i {
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
		file_proto_config_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_config_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigByObjectIdRequest); i {
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
		file_proto_config_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigByIdRequest); i {
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
		file_proto_config_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllConfigsRequest); i {
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
		file_proto_config_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigRequest); i {
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
		file_proto_config_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertConfigRequest); i {
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
			RawDescriptor: file_proto_config_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_config_service_proto_goTypes,
		DependencyIndexes: file_proto_config_service_proto_depIdxs,
		MessageInfos:      file_proto_config_service_proto_msgTypes,
	}.Build()
	File_proto_config_service_proto = out.File
	file_proto_config_service_proto_rawDesc = nil
	file_proto_config_service_proto_goTypes = nil
	file_proto_config_service_proto_depIdxs = nil
}