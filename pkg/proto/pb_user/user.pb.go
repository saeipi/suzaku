// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: pb_user/user.proto

package pb_user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	pb_com "suzaku/pkg/proto/pb_com"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of google/protobuf/timestamp.proto.

type Timestamp = timestamppb.Timestamp

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户ID
	SzkId    string `protobuf:"bytes,2,opt,name=szk_id,json=szkId,proto3" json:"szk_id,omitempty"`    // 账户ID
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`           // 昵称
	Gender   int32  `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`              // 性别
	//google.protobuf.Timestamp birth = 6; // 生日
	BirthTs   int64  `protobuf:"varint,5,opt,name=birth_ts,json=birthTs,proto3" json:"birth_ts,omitempty"`      // 生日
	Email     string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`                          // Email
	Mobile    string `protobuf:"bytes,7,opt,name=mobile,proto3" json:"mobile,omitempty"`                        // 手机号
	AvatarUrl string `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"` // 头像
	CityId    int64  `protobuf:"varint,9,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`         // 城市ID
	Country   string `protobuf:"bytes,10,opt,name=country,proto3" json:"country,omitempty"`                     // 国家
	City      string `protobuf:"bytes,11,opt,name=city,proto3" json:"city,omitempty"`                           // 城市
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_pb_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserInfo) GetSzkId() string {
	if x != nil {
		return x.SzkId
	}
	return ""
}

func (x *UserInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserInfo) GetBirthTs() int64 {
	if x != nil {
		return x.BirthTs
	}
	return 0
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserInfo) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *UserInfo) GetCityId() int64 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *UserInfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UserInfo) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type UserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户ID
}

func (x *UserInfoReq) Reset() {
	*x = UserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReq) ProtoMessage() {}

func (x *UserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReq.ProtoReflect.Descriptor instead.
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return file_pb_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfoReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common   *pb_com.CommonResp `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	UserInfo *UserInfo          `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *UserInfoResp) Reset() {
	*x = UserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResp) ProtoMessage() {}

func (x *UserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResp.ProtoReflect.Descriptor instead.
func (*UserInfoResp) Descriptor() ([]byte, []int) {
	return file_pb_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfoResp) GetCommon() *pb_com.CommonResp {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *UserInfoResp) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type EditUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // 用户ID
	SzkId     string `protobuf:"bytes,2,opt,name=szk_id,json=szkId,proto3" json:"szk_id,omitempty"`             // 账户ID (一年只允许修改一次)
	Nickname  string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`                    // 昵称
	Gender    int32  `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`                       // 性别
	BirthTs   int64  `protobuf:"varint,5,opt,name=birth_ts,json=birthTs,proto3" json:"birth_ts,omitempty"`      // 生日
	Email     string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`                          // Email
	Mobile    string `protobuf:"bytes,7,opt,name=mobile,proto3" json:"mobile,omitempty"`                        // 手机号
	AvatarUrl string `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"` // 头像
	CityId    int64  `protobuf:"varint,10,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`        // 城市ID
}

func (x *EditUserInfoReq) Reset() {
	*x = EditUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUserInfoReq) ProtoMessage() {}

func (x *EditUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditUserInfoReq.ProtoReflect.Descriptor instead.
func (*EditUserInfoReq) Descriptor() ([]byte, []int) {
	return file_pb_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *EditUserInfoReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EditUserInfoReq) GetSzkId() string {
	if x != nil {
		return x.SzkId
	}
	return ""
}

func (x *EditUserInfoReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *EditUserInfoReq) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *EditUserInfoReq) GetBirthTs() int64 {
	if x != nil {
		return x.BirthTs
	}
	return 0
}

func (x *EditUserInfoReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EditUserInfoReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *EditUserInfoReq) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *EditUserInfoReq) GetCityId() int64 {
	if x != nil {
		return x.CityId
	}
	return 0
}

var File_pb_user_user_proto protoreflect.FileDescriptor

var file_pb_user_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x13, 0x70,
	0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x7a, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x7a, 0x6b, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x74, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x69, 0x72, 0x74, 0x68, 0x54, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x22, 0x26, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62,
	0x5f, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52,
	0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x0f, 0x45, 0x64, 0x69, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x7a, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x7a, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x62, 0x69, 0x72, 0x74, 0x68, 0x54, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x32, 0x7d, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3c, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x64, 0x69, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62,
	0x5f, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x22, 0x5a, 0x20, 0x73, 0x75, 0x7a, 0x61, 0x6b, 0x75, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x70, 0x62, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_user_user_proto_rawDescOnce sync.Once
	file_pb_user_user_proto_rawDescData = file_pb_user_user_proto_rawDesc
)

func file_pb_user_user_proto_rawDescGZIP() []byte {
	file_pb_user_user_proto_rawDescOnce.Do(func() {
		file_pb_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_user_user_proto_rawDescData)
	})
	return file_pb_user_user_proto_rawDescData
}

var file_pb_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_user_user_proto_goTypes = []interface{}{
	(*UserInfo)(nil),          // 0: pb_user.UserInfo
	(*UserInfoReq)(nil),       // 1: pb_user.UserInfoReq
	(*UserInfoResp)(nil),      // 2: pb_user.UserInfoResp
	(*EditUserInfoReq)(nil),   // 3: pb_user.EditUserInfoReq
	(*pb_com.CommonResp)(nil), // 4: pb_com.CommonResp
}
var file_pb_user_user_proto_depIdxs = []int32{
	4, // 0: pb_user.UserInfoResp.common:type_name -> pb_com.CommonResp
	0, // 1: pb_user.UserInfoResp.user_info:type_name -> pb_user.UserInfo
	1, // 2: pb_user.User.UserInfo:input_type -> pb_user.UserInfoReq
	3, // 3: pb_user.User.EditUserInfo:input_type -> pb_user.EditUserInfoReq
	2, // 4: pb_user.User.UserInfo:output_type -> pb_user.UserInfoResp
	4, // 5: pb_user.User.EditUserInfo:output_type -> pb_com.CommonResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_user_user_proto_init() }
func file_pb_user_user_proto_init() {
	if File_pb_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_pb_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoReq); i {
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
		file_pb_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResp); i {
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
		file_pb_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditUserInfoReq); i {
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
			RawDescriptor: file_pb_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_user_user_proto_goTypes,
		DependencyIndexes: file_pb_user_user_proto_depIdxs,
		MessageInfos:      file_pb_user_user_proto_msgTypes,
	}.Build()
	File_pb_user_user_proto = out.File
	file_pb_user_user_proto_rawDesc = nil
	file_pb_user_user_proto_goTypes = nil
	file_pb_user_user_proto_depIdxs = nil
}
