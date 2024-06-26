// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: api/grpc/v1/group.proto

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

type GroupStatus int32

const (
	GroupStatus_GROUP_STATUS_NORMAL  GroupStatus = 0 // 正常状态
	GroupStatus_GROUP_STATUS_LOCKED  GroupStatus = 1 // 锁定状态
	GroupStatus_GROUP_STATUS_DELETED GroupStatus = 2 // 删除状态
)

// Enum value maps for GroupStatus.
var (
	GroupStatus_name = map[int32]string{
		0: "GROUP_STATUS_NORMAL",
		1: "GROUP_STATUS_LOCKED",
		2: "GROUP_STATUS_DELETED",
	}
	GroupStatus_value = map[string]int32{
		"GROUP_STATUS_NORMAL":  0,
		"GROUP_STATUS_LOCKED":  1,
		"GROUP_STATUS_DELETED": 2,
	}
)

func (x GroupStatus) Enum() *GroupStatus {
	p := new(GroupStatus)
	*p = x
	return p
}

func (x GroupStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_grpc_v1_group_proto_enumTypes[0].Descriptor()
}

func (GroupStatus) Type() protoreflect.EnumType {
	return &file_api_grpc_v1_group_proto_enumTypes[0]
}

func (x GroupStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupStatus.Descriptor instead.
func (GroupStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_proto_rawDescGZIP(), []int{0}
}

type GroupType int32

const (
	GroupType_Private GroupType = 0 // 私有群
	GroupType_Public  GroupType = 1 // 公开群
)

// Enum value maps for GroupType.
var (
	GroupType_name = map[int32]string{
		0: "Private",
		1: "Public",
	}
	GroupType_value = map[string]int32{
		"Private": 0,
		"Public":  1,
	}
)

func (x GroupType) Enum() *GroupType {
	p := new(GroupType)
	*p = x
	return p
}

func (x GroupType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_grpc_v1_group_proto_enumTypes[1].Descriptor()
}

func (GroupType) Type() protoreflect.EnumType {
	return &file_api_grpc_v1_group_proto_enumTypes[1]
}

func (x GroupType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupType.Descriptor instead.
func (GroupType) EnumDescriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_proto_rawDescGZIP(), []int{1}
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	Id uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"id"`
	// @inject_tag: json:"type"
	Type GroupType `protobuf:"varint,2,opt,name=Type,proto3,enum=group_v1.GroupType" json:"type"`
	// @inject_tag: json:"status"
	Status GroupStatus `protobuf:"varint,3,opt,name=Status,proto3,enum=group_v1.GroupStatus" json:"status"`
	// @inject_tag: json:"max_members_limit"
	MaxMembersLimit int32 `protobuf:"varint,4,opt,name=Max_members_limit,json=MaxMembersLimit,proto3" json:"max_members_limit"`
	// @inject_tag: json:"creator_id"
	CreatorId string `protobuf:"bytes,5,opt,name=Creator_id,json=CreatorId,proto3" json:"creator_id"`
	// @inject_tag: json:"name"
	Name string `protobuf:"bytes,6,opt,name=Name,proto3" json:"name"`
	// @inject_tag: json:"avatar"
	Avatar string `protobuf:"bytes,7,opt,name=Avatar,proto3" json:"avatar"`
	// @inject_tag: json:"member"
	Member []string `protobuf:"bytes,8,rep,name=Member,proto3" json:"member"` // 创建群聊的时候邀请的成员
	// @inject_tag: json:"silence_time"
	// 禁言时间，如果为0表示不开启禁言
	SilenceTime int64 `protobuf:"varint,9,opt,name=SilenceTime,proto3" json:"silence_time"`
	// @inject_tag: json:"join_approve"
	// 入群申请是否需要审核
	JoinApprove bool `protobuf:"varint,10,opt,name=JoinApprove,proto3" json:"join_approve"`
	// @inject_tag: json:"encrypt"
	// 是否开启群聊加密，只有当群聊类型为私密时有效
	Encrypt bool `protobuf:"varint,11,opt,name=Encrypt,proto3" json:"encrypt"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Group) GetType() GroupType {
	if x != nil {
		return x.Type
	}
	return GroupType_Private
}

func (x *Group) GetStatus() GroupStatus {
	if x != nil {
		return x.Status
	}
	return GroupStatus_GROUP_STATUS_NORMAL
}

func (x *Group) GetMaxMembersLimit() int32 {
	if x != nil {
		return x.MaxMembersLimit
	}
	return 0
}

func (x *Group) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Group) GetMember() []string {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *Group) GetSilenceTime() int64 {
	if x != nil {
		return x.SilenceTime
	}
	return 0
}

func (x *Group) GetJoinApprove() bool {
	if x != nil {
		return x.JoinApprove
	}
	return false
}

func (x *Group) GetEncrypt() bool {
	if x != nil {
		return x.Encrypt
	}
	return false
}

type GetGroupInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"gid"
	Gid uint32 `protobuf:"varint,1,opt,name=Gid,proto3" json:"gid"`
}

func (x *GetGroupInfoRequest) Reset() {
	*x = GetGroupInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupInfoRequest) ProtoMessage() {}

func (x *GetGroupInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupInfoRequest.ProtoReflect.Descriptor instead.
func (*GetGroupInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_proto_rawDescGZIP(), []int{1}
}

func (x *GetGroupInfoRequest) GetGid() uint32 {
	if x != nil {
		return x.Gid
	}
	return 0
}

type GetBatchGroupInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"group_ids"
	GroupIds []uint32 `protobuf:"varint,1,rep,packed,name=Group_ids,json=GroupIds,proto3" json:"group_ids"`
}

func (x *GetBatchGroupInfoRequest) Reset() {
	*x = GetBatchGroupInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchGroupInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchGroupInfoRequest) ProtoMessage() {}

func (x *GetBatchGroupInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchGroupInfoRequest.ProtoReflect.Descriptor instead.
func (*GetBatchGroupInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_proto_rawDescGZIP(), []int{2}
}

func (x *GetBatchGroupInfoRequest) GetGroupIds() []uint32 {
	if x != nil {
		return x.GroupIds
	}
	return nil
}

type GetBatchGroupInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"groups"
	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups"`
}

func (x *GetBatchGroupInfoResponse) Reset() {
	*x = GetBatchGroupInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchGroupInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchGroupInfoResponse) ProtoMessage() {}

func (x *GetBatchGroupInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchGroupInfoResponse.ProtoReflect.Descriptor instead.
func (*GetBatchGroupInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_proto_rawDescGZIP(), []int{3}
}

func (x *GetBatchGroupInfoResponse) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_api_grpc_v1_group_proto protoreflect.FileDescriptor

var file_api_grpc_v1_group_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x76, 0x31, 0x22, 0xdc, 0x02, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76,
	0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x4d, 0x61, 0x78, 0x5f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x4d, 0x61, 0x78, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x53, 0x69, 0x6c, 0x65, 0x6e,
	0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4a, 0x6f, 0x69,
	0x6e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x22, 0x27, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x47, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x73, 0x22, 0x44, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2a, 0x59, 0x0a, 0x0b, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x24, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x10, 0x01, 0x32, 0xb6, 0x01, 0x0a, 0x0c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x47, 0x69,
	0x64, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x61, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x73, 0x12, 0x22, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x73, 0x69, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x73, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_v1_group_proto_rawDescOnce sync.Once
	file_api_grpc_v1_group_proto_rawDescData = file_api_grpc_v1_group_proto_rawDesc
)

func file_api_grpc_v1_group_proto_rawDescGZIP() []byte {
	file_api_grpc_v1_group_proto_rawDescOnce.Do(func() {
		file_api_grpc_v1_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_v1_group_proto_rawDescData)
	})
	return file_api_grpc_v1_group_proto_rawDescData
}

var file_api_grpc_v1_group_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_grpc_v1_group_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_grpc_v1_group_proto_goTypes = []interface{}{
	(GroupStatus)(0),                  // 0: group_v1.GroupStatus
	(GroupType)(0),                    // 1: group_v1.GroupType
	(*Group)(nil),                     // 2: group_v1.Group
	(*GetGroupInfoRequest)(nil),       // 3: group_v1.GetGroupInfoRequest
	(*GetBatchGroupInfoRequest)(nil),  // 4: group_v1.GetBatchGroupInfoRequest
	(*GetBatchGroupInfoResponse)(nil), // 5: group_v1.GetBatchGroupInfoResponse
}
var file_api_grpc_v1_group_proto_depIdxs = []int32{
	1, // 0: group_v1.Group.Type:type_name -> group_v1.GroupType
	0, // 1: group_v1.Group.Status:type_name -> group_v1.GroupStatus
	2, // 2: group_v1.GetBatchGroupInfoResponse.groups:type_name -> group_v1.Group
	3, // 3: group_v1.GroupService.GetGroupInfoByGid:input_type -> group_v1.GetGroupInfoRequest
	4, // 4: group_v1.GroupService.GetBatchGroupInfoByIDs:input_type -> group_v1.GetBatchGroupInfoRequest
	2, // 5: group_v1.GroupService.GetGroupInfoByGid:output_type -> group_v1.Group
	5, // 6: group_v1.GroupService.GetBatchGroupInfoByIDs:output_type -> group_v1.GetBatchGroupInfoResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_grpc_v1_group_proto_init() }
func file_api_grpc_v1_group_proto_init() {
	if File_api_grpc_v1_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_v1_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_api_grpc_v1_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupInfoRequest); i {
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
		file_api_grpc_v1_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchGroupInfoRequest); i {
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
		file_api_grpc_v1_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchGroupInfoResponse); i {
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
			RawDescriptor: file_api_grpc_v1_group_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_v1_group_proto_goTypes,
		DependencyIndexes: file_api_grpc_v1_group_proto_depIdxs,
		EnumInfos:         file_api_grpc_v1_group_proto_enumTypes,
		MessageInfos:      file_api_grpc_v1_group_proto_msgTypes,
	}.Build()
	File_api_grpc_v1_group_proto = out.File
	file_api_grpc_v1_group_proto_rawDesc = nil
	file_api_grpc_v1_group_proto_goTypes = nil
	file_api_grpc_v1_group_proto_depIdxs = nil
}
