// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: service/server/handlers/shipper/proto/shipper.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddShipperReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ImageURL    string `protobuf:"bytes,2,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	MaxWeight   int32  `protobuf:"varint,4,opt,name=MaxWeight,proto3" json:"MaxWeight,omitempty"`
	CreatedBy   int32  `protobuf:"varint,5,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
}

func (x *AddShipperReq) Reset() {
	*x = AddShipperReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddShipperReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddShipperReq) ProtoMessage() {}

func (x *AddShipperReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddShipperReq.ProtoReflect.Descriptor instead.
func (*AddShipperReq) Descriptor() ([]byte, []int) {
	return file_service_server_handlers_shipper_proto_shipper_proto_rawDescGZIP(), []int{0}
}

func (x *AddShipperReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddShipperReq) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *AddShipperReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddShipperReq) GetMaxWeight() int32 {
	if x != nil {
		return x.MaxWeight
	}
	return 0
}

func (x *AddShipperReq) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

type AddShipperResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *AddShipperResp) Reset() {
	*x = AddShipperResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddShipperResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddShipperResp) ProtoMessage() {}

func (x *AddShipperResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddShipperResp.ProtoReflect.Descriptor instead.
func (*AddShipperResp) Descriptor() ([]byte, []int) {
	return file_service_server_handlers_shipper_proto_shipper_proto_rawDescGZIP(), []int{1}
}

func (x *AddShipperResp) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AddShipperResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetShipperReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetShipperReq) Reset() {
	*x = GetShipperReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShipperReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShipperReq) ProtoMessage() {}

func (x *GetShipperReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShipperReq.ProtoReflect.Descriptor instead.
func (*GetShipperReq) Descriptor() ([]byte, []int) {
	return file_service_server_handlers_shipper_proto_shipper_proto_rawDescGZIP(), []int{2}
}

func (x *GetShipperReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetShipperResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []*ShipperData `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	Message string         `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *GetShipperResp) Reset() {
	*x = GetShipperResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShipperResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShipperResp) ProtoMessage() {}

func (x *GetShipperResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShipperResp.ProtoReflect.Descriptor instead.
func (*GetShipperResp) Descriptor() ([]byte, []int) {
	return file_service_server_handlers_shipper_proto_shipper_proto_rawDescGZIP(), []int{3}
}

func (x *GetShipperResp) GetData() []*ShipperData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetShipperResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ShipperData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ImageURL    string `protobuf:"bytes,3,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	MaxWeight   int32  `protobuf:"varint,5,opt,name=MaxWeight,proto3" json:"MaxWeight,omitempty"`
	CreatedAt   int32  `protobuf:"varint,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	CreatedBy   int32  `protobuf:"varint,7,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
	UpdatedAt   int32  `protobuf:"varint,8,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	UpdatedBy   int32  `protobuf:"varint,9,opt,name=UpdatedBy,proto3" json:"UpdatedBy,omitempty"`
}

func (x *ShipperData) Reset() {
	*x = ShipperData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipperData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipperData) ProtoMessage() {}

func (x *ShipperData) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipperData.ProtoReflect.Descriptor instead.
func (*ShipperData) Descriptor() ([]byte, []int) {
	return file_service_server_handlers_shipper_proto_shipper_proto_rawDescGZIP(), []int{4}
}

func (x *ShipperData) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ShipperData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShipperData) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *ShipperData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ShipperData) GetMaxWeight() int32 {
	if x != nil {
		return x.MaxWeight
	}
	return 0
}

func (x *ShipperData) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ShipperData) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *ShipperData) GetUpdatedAt() int32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *ShipperData) GetUpdatedBy() int32 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

type DeleteShipperReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteShipperReq) Reset() {
	*x = DeleteShipperReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShipperReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShipperReq) ProtoMessage() {}

func (x *DeleteShipperReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShipperReq.ProtoReflect.Descriptor instead.
func (*DeleteShipperReq) Descriptor() ([]byte, []int) {
	return file_service_server_handlers_shipper_proto_shipper_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteShipperReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteShipperResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *DeleteShipperResp) Reset() {
	*x = DeleteShipperResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShipperResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShipperResp) ProtoMessage() {}

func (x *DeleteShipperResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShipperResp.ProtoReflect.Descriptor instead.
func (*DeleteShipperResp) Descriptor() ([]byte, []int) {
	return file_service_server_handlers_shipper_proto_shipper_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteShipperResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_service_server_handlers_shipper_proto_shipper_proto protoreflect.FileDescriptor

var file_service_server_handlers_shipper_proto_shipper_proto_rawDesc = []byte{
	0x0a, 0x33, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x9d,
	0x01, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x61, 0x78, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4d, 0x61, 0x78, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x3a,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x54, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x85, 0x02, 0x0a, 0x0b, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52,
	0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52,
	0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x61, 0x78, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4d, 0x61, 0x78, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x2d, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd5, 0x01, 0x0a,
	0x07, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x53,
	0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x68, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69,
	0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x73, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_server_handlers_shipper_proto_shipper_proto_rawDescOnce sync.Once
	file_service_server_handlers_shipper_proto_shipper_proto_rawDescData = file_service_server_handlers_shipper_proto_shipper_proto_rawDesc
)

func file_service_server_handlers_shipper_proto_shipper_proto_rawDescGZIP() []byte {
	file_service_server_handlers_shipper_proto_shipper_proto_rawDescOnce.Do(func() {
		file_service_server_handlers_shipper_proto_shipper_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_server_handlers_shipper_proto_shipper_proto_rawDescData)
	})
	return file_service_server_handlers_shipper_proto_shipper_proto_rawDescData
}

var file_service_server_handlers_shipper_proto_shipper_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_server_handlers_shipper_proto_shipper_proto_goTypes = []interface{}{
	(*AddShipperReq)(nil),     // 0: product.AddShipperReq
	(*AddShipperResp)(nil),    // 1: product.AddShipperResp
	(*GetShipperReq)(nil),     // 2: product.GetShipperReq
	(*GetShipperResp)(nil),    // 3: product.GetShipperResp
	(*ShipperData)(nil),       // 4: product.ShipperData
	(*DeleteShipperReq)(nil),  // 5: product.DeleteShipperReq
	(*DeleteShipperResp)(nil), // 6: product.DeleteShipperResp
}
var file_service_server_handlers_shipper_proto_shipper_proto_depIdxs = []int32{
	4, // 0: product.GetShipperResp.Data:type_name -> product.ShipperData
	0, // 1: product.Shipper.AddShipper:input_type -> product.AddShipperReq
	2, // 2: product.Shipper.GetShipper:input_type -> product.GetShipperReq
	5, // 3: product.Shipper.DeleteShipper:input_type -> product.DeleteShipperReq
	1, // 4: product.Shipper.AddShipper:output_type -> product.AddShipperResp
	3, // 5: product.Shipper.GetShipper:output_type -> product.GetShipperResp
	6, // 6: product.Shipper.DeleteShipper:output_type -> product.DeleteShipperResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_server_handlers_shipper_proto_shipper_proto_init() }
func file_service_server_handlers_shipper_proto_shipper_proto_init() {
	if File_service_server_handlers_shipper_proto_shipper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddShipperReq); i {
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
		file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddShipperResp); i {
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
		file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShipperReq); i {
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
		file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShipperResp); i {
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
		file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipperData); i {
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
		file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShipperReq); i {
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
		file_service_server_handlers_shipper_proto_shipper_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShipperResp); i {
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
			RawDescriptor: file_service_server_handlers_shipper_proto_shipper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_server_handlers_shipper_proto_shipper_proto_goTypes,
		DependencyIndexes: file_service_server_handlers_shipper_proto_shipper_proto_depIdxs,
		MessageInfos:      file_service_server_handlers_shipper_proto_shipper_proto_msgTypes,
	}.Build()
	File_service_server_handlers_shipper_proto_shipper_proto = out.File
	file_service_server_handlers_shipper_proto_shipper_proto_rawDesc = nil
	file_service_server_handlers_shipper_proto_shipper_proto_goTypes = nil
	file_service_server_handlers_shipper_proto_shipper_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShipperClient is the client API for Shipper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShipperClient interface {
	AddShipper(ctx context.Context, in *AddShipperReq, opts ...grpc.CallOption) (*AddShipperResp, error)
	GetShipper(ctx context.Context, in *GetShipperReq, opts ...grpc.CallOption) (*GetShipperResp, error)
	DeleteShipper(ctx context.Context, in *DeleteShipperReq, opts ...grpc.CallOption) (*DeleteShipperResp, error)
}

type shipperClient struct {
	cc grpc.ClientConnInterface
}

func NewShipperClient(cc grpc.ClientConnInterface) ShipperClient {
	return &shipperClient{cc}
}

func (c *shipperClient) AddShipper(ctx context.Context, in *AddShipperReq, opts ...grpc.CallOption) (*AddShipperResp, error) {
	out := new(AddShipperResp)
	err := c.cc.Invoke(ctx, "/product.Shipper/AddShipper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipperClient) GetShipper(ctx context.Context, in *GetShipperReq, opts ...grpc.CallOption) (*GetShipperResp, error) {
	out := new(GetShipperResp)
	err := c.cc.Invoke(ctx, "/product.Shipper/GetShipper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipperClient) DeleteShipper(ctx context.Context, in *DeleteShipperReq, opts ...grpc.CallOption) (*DeleteShipperResp, error) {
	out := new(DeleteShipperResp)
	err := c.cc.Invoke(ctx, "/product.Shipper/DeleteShipper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShipperServer is the server API for Shipper service.
type ShipperServer interface {
	AddShipper(context.Context, *AddShipperReq) (*AddShipperResp, error)
	GetShipper(context.Context, *GetShipperReq) (*GetShipperResp, error)
	DeleteShipper(context.Context, *DeleteShipperReq) (*DeleteShipperResp, error)
}

// UnimplementedShipperServer can be embedded to have forward compatible implementations.
type UnimplementedShipperServer struct {
}

func (*UnimplementedShipperServer) AddShipper(context.Context, *AddShipperReq) (*AddShipperResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddShipper not implemented")
}
func (*UnimplementedShipperServer) GetShipper(context.Context, *GetShipperReq) (*GetShipperResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShipper not implemented")
}
func (*UnimplementedShipperServer) DeleteShipper(context.Context, *DeleteShipperReq) (*DeleteShipperResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShipper not implemented")
}

func RegisterShipperServer(s *grpc.Server, srv ShipperServer) {
	s.RegisterService(&_Shipper_serviceDesc, srv)
}

func _Shipper_AddShipper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddShipperReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipperServer).AddShipper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Shipper/AddShipper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipperServer).AddShipper(ctx, req.(*AddShipperReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shipper_GetShipper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShipperReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipperServer).GetShipper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Shipper/GetShipper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipperServer).GetShipper(ctx, req.(*GetShipperReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shipper_DeleteShipper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShipperReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipperServer).DeleteShipper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Shipper/DeleteShipper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipperServer).DeleteShipper(ctx, req.(*DeleteShipperReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shipper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.Shipper",
	HandlerType: (*ShipperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddShipper",
			Handler:    _Shipper_AddShipper_Handler,
		},
		{
			MethodName: "GetShipper",
			Handler:    _Shipper_GetShipper_Handler,
		},
		{
			MethodName: "DeleteShipper",
			Handler:    _Shipper_DeleteShipper_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/server/handlers/shipper/proto/shipper.proto",
}
