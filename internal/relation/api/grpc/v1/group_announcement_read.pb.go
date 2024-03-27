// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: api/grpc/v1/group_announcement_read.proto

package v1

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

type MarkAnnouncementAsReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,1,opt,name=GroupId,proto3" json:"group_id"`
	// @inject_tag: json:"announcement_id"
	AnnouncementId uint32 `protobuf:"varint,2,opt,name=AnnouncementId,proto3" json:"announcement_id"`
	// @inject_tag: json:"user_ids"
	UserIds []string `protobuf:"bytes,3,rep,name=UserIds,proto3" json:"user_ids"`
}

func (x *MarkAnnouncementAsReadRequest) Reset() {
	*x = MarkAnnouncementAsReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkAnnouncementAsReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkAnnouncementAsReadRequest) ProtoMessage() {}

func (x *MarkAnnouncementAsReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkAnnouncementAsReadRequest.ProtoReflect.Descriptor instead.
func (*MarkAnnouncementAsReadRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_announcement_read_proto_rawDescGZIP(), []int{0}
}

func (x *MarkAnnouncementAsReadRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MarkAnnouncementAsReadRequest) GetAnnouncementId() uint32 {
	if x != nil {
		return x.AnnouncementId
	}
	return 0
}

func (x *MarkAnnouncementAsReadRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type MarkAnnouncementAsReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
}

func (x *MarkAnnouncementAsReadResponse) Reset() {
	*x = MarkAnnouncementAsReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkAnnouncementAsReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkAnnouncementAsReadResponse) ProtoMessage() {}

func (x *MarkAnnouncementAsReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkAnnouncementAsReadResponse.ProtoReflect.Descriptor instead.
func (*MarkAnnouncementAsReadResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_announcement_read_proto_rawDescGZIP(), []int{1}
}

func (x *MarkAnnouncementAsReadResponse) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetReadUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,1,opt,name=GroupId,proto3" json:"group_id"`
	// @inject_tag: json:"announcement_id"
	AnnouncementId uint32 `protobuf:"varint,2,opt,name=AnnouncementId,proto3" json:"announcement_id"`
}

func (x *GetReadUsersRequest) Reset() {
	*x = GetReadUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadUsersRequest) ProtoMessage() {}

func (x *GetReadUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadUsersRequest.ProtoReflect.Descriptor instead.
func (*GetReadUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_announcement_read_proto_rawDescGZIP(), []int{2}
}

func (x *GetReadUsersRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GetReadUsersRequest) GetAnnouncementId() uint32 {
	if x != nil {
		return x.AnnouncementId
	}
	return 0
}

type GetReadUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"announcement_read_users"
	AnnouncementReadUsers []*AnnouncementRead `protobuf:"bytes,1,rep,name=AnnouncementReadUsers,proto3" json:"announcement_read_users"`
}

func (x *GetReadUsersResponse) Reset() {
	*x = GetReadUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadUsersResponse) ProtoMessage() {}

func (x *GetReadUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadUsersResponse.ProtoReflect.Descriptor instead.
func (*GetReadUsersResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_announcement_read_proto_rawDescGZIP(), []int{3}
}

func (x *GetReadUsersResponse) GetAnnouncementReadUsers() []*AnnouncementRead {
	if x != nil {
		return x.AnnouncementReadUsers
	}
	return nil
}

type AnnouncementRead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,1,opt,name=GroupId,proto3" json:"group_id"`
	// @inject_tag: json:"announcement_id"
	AnnouncementId uint32 `protobuf:"varint,2,opt,name=AnnouncementId,proto3" json:"announcement_id"`
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,3,opt,name=UserId,proto3" json:"user_id"`
	// @inject_tag: json:"read_at"
	ReadAt uint64 `protobuf:"varint,4,opt,name=ReadAt,proto3" json:"read_at"`
	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,5,opt,name=ID,proto3" json:"id"`
}

func (x *AnnouncementRead) Reset() {
	*x = AnnouncementRead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnouncementRead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnouncementRead) ProtoMessage() {}

func (x *AnnouncementRead) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnouncementRead.ProtoReflect.Descriptor instead.
func (*AnnouncementRead) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_announcement_read_proto_rawDescGZIP(), []int{4}
}

func (x *AnnouncementRead) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *AnnouncementRead) GetAnnouncementId() uint32 {
	if x != nil {
		return x.AnnouncementId
	}
	return 0
}

func (x *AnnouncementRead) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AnnouncementRead) GetReadAt() uint64 {
	if x != nil {
		return x.ReadAt
	}
	return 0
}

func (x *AnnouncementRead) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetAnnouncementReadByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,1,opt,name=GroupId,proto3" json:"group_id"`
	// @inject_tag: json:"announcement_id"
	AnnouncementId uint32 `protobuf:"varint,2,opt,name=AnnouncementId,proto3" json:"announcement_id"`
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,3,opt,name=UserId,proto3" json:"user_id"`
}

func (x *GetAnnouncementReadByUserIdRequest) Reset() {
	*x = GetAnnouncementReadByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnnouncementReadByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnnouncementReadByUserIdRequest) ProtoMessage() {}

func (x *GetAnnouncementReadByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnnouncementReadByUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetAnnouncementReadByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_announcement_read_proto_rawDescGZIP(), []int{5}
}

func (x *GetAnnouncementReadByUserIdRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GetAnnouncementReadByUserIdRequest) GetAnnouncementId() uint32 {
	if x != nil {
		return x.AnnouncementId
	}
	return 0
}

func (x *GetAnnouncementReadByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetAnnouncementReadByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,1,opt,name=GroupId,proto3" json:"group_id"`
	// @inject_tag: json:"announcement_id"
	AnnouncementId uint32 `protobuf:"varint,2,opt,name=AnnouncementId,proto3" json:"announcement_id"`
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,3,opt,name=UserId,proto3" json:"user_id"`
	// @inject_tag: json:"read_at"
	ReadAt uint64 `protobuf:"varint,4,opt,name=ReadAt,proto3" json:"read_at"`
	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,5,opt,name=ID,proto3" json:"id"`
}

func (x *GetAnnouncementReadByUserIdResponse) Reset() {
	*x = GetAnnouncementReadByUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnnouncementReadByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnnouncementReadByUserIdResponse) ProtoMessage() {}

func (x *GetAnnouncementReadByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_announcement_read_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnnouncementReadByUserIdResponse.ProtoReflect.Descriptor instead.
func (*GetAnnouncementReadByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_announcement_read_proto_rawDescGZIP(), []int{6}
}

func (x *GetAnnouncementReadByUserIdResponse) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GetAnnouncementReadByUserIdResponse) GetAnnouncementId() uint32 {
	if x != nil {
		return x.AnnouncementId
	}
	return 0
}

func (x *GetAnnouncementReadByUserIdResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetAnnouncementReadByUserIdResponse) GetReadAt() uint64 {
	if x != nil {
		return x.ReadAt
	}
	return 0
}

func (x *GetAnnouncementReadByUserIdResponse) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_api_grpc_v1_group_announcement_read_proto protoreflect.FileDescriptor

var file_api_grpc_v1_group_announcement_read_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22,
	0x7b, 0x0a, 0x1d, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x30, 0x0a, 0x1e,
	0x4d, 0x61, 0x72, 0x6b, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x22, 0x57,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4a, 0x0a, 0x15, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x15, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x10,
	0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65,
	0x61, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x52, 0x65, 0x61, 0x64,
	0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x7e, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x32, 0xb2, 0x02, 0x0a,
	0x1c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a,
	0x16, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x72, 0x6b, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x17,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6e, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x26, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61,
	0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x73, 0x73, 0x69, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_v1_group_announcement_read_proto_rawDescOnce sync.Once
	file_api_grpc_v1_group_announcement_read_proto_rawDescData = file_api_grpc_v1_group_announcement_read_proto_rawDesc
)

func file_api_grpc_v1_group_announcement_read_proto_rawDescGZIP() []byte {
	file_api_grpc_v1_group_announcement_read_proto_rawDescOnce.Do(func() {
		file_api_grpc_v1_group_announcement_read_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_v1_group_announcement_read_proto_rawDescData)
	})
	return file_api_grpc_v1_group_announcement_read_proto_rawDescData
}

var file_api_grpc_v1_group_announcement_read_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_grpc_v1_group_announcement_read_proto_goTypes = []interface{}{
	(*MarkAnnouncementAsReadRequest)(nil),       // 0: v1.MarkAnnouncementAsReadRequest
	(*MarkAnnouncementAsReadResponse)(nil),      // 1: v1.MarkAnnouncementAsReadResponse
	(*GetReadUsersRequest)(nil),                 // 2: v1.GetReadUsersRequest
	(*GetReadUsersResponse)(nil),                // 3: v1.GetReadUsersResponse
	(*AnnouncementRead)(nil),                    // 4: v1.AnnouncementRead
	(*GetAnnouncementReadByUserIdRequest)(nil),  // 5: v1.GetAnnouncementReadByUserIdRequest
	(*GetAnnouncementReadByUserIdResponse)(nil), // 6: v1.GetAnnouncementReadByUserIdResponse
}
var file_api_grpc_v1_group_announcement_read_proto_depIdxs = []int32{
	4, // 0: v1.GetReadUsersResponse.AnnouncementReadUsers:type_name -> v1.AnnouncementRead
	0, // 1: v1.GroupAnnouncementReadService.MarkAnnouncementAsRead:input_type -> v1.MarkAnnouncementAsReadRequest
	2, // 2: v1.GroupAnnouncementReadService.GetReadUsers:input_type -> v1.GetReadUsersRequest
	5, // 3: v1.GroupAnnouncementReadService.GetAnnouncementReadByUserId:input_type -> v1.GetAnnouncementReadByUserIdRequest
	1, // 4: v1.GroupAnnouncementReadService.MarkAnnouncementAsRead:output_type -> v1.MarkAnnouncementAsReadResponse
	3, // 5: v1.GroupAnnouncementReadService.GetReadUsers:output_type -> v1.GetReadUsersResponse
	6, // 6: v1.GroupAnnouncementReadService.GetAnnouncementReadByUserId:output_type -> v1.GetAnnouncementReadByUserIdResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_grpc_v1_group_announcement_read_proto_init() }
func file_api_grpc_v1_group_announcement_read_proto_init() {
	if File_api_grpc_v1_group_announcement_read_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_v1_group_announcement_read_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkAnnouncementAsReadRequest); i {
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
		file_api_grpc_v1_group_announcement_read_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkAnnouncementAsReadResponse); i {
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
		file_api_grpc_v1_group_announcement_read_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReadUsersRequest); i {
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
		file_api_grpc_v1_group_announcement_read_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReadUsersResponse); i {
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
		file_api_grpc_v1_group_announcement_read_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnouncementRead); i {
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
		file_api_grpc_v1_group_announcement_read_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnnouncementReadByUserIdRequest); i {
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
		file_api_grpc_v1_group_announcement_read_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnnouncementReadByUserIdResponse); i {
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
			RawDescriptor: file_api_grpc_v1_group_announcement_read_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_v1_group_announcement_read_proto_goTypes,
		DependencyIndexes: file_api_grpc_v1_group_announcement_read_proto_depIdxs,
		MessageInfos:      file_api_grpc_v1_group_announcement_read_proto_msgTypes,
	}.Build()
	File_api_grpc_v1_group_announcement_read_proto = out.File
	file_api_grpc_v1_group_announcement_read_proto_rawDesc = nil
	file_api_grpc_v1_group_announcement_read_proto_goTypes = nil
	file_api_grpc_v1_group_announcement_read_proto_depIdxs = nil
}
