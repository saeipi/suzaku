// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: chart/chat.proto

package pb_chat

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	pb_com "suzaku/pkg/proto/pb_com"
	pb_ws "suzaku/pkg/proto/pb_ws"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgDataToMQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string         `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	OperationId string         `protobuf:"bytes,2,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	MsgData     *pb_ws.MsgData `protobuf:"bytes,3,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`
}

func (x *MsgDataToMQ) Reset() {
	*x = MsgDataToMQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDataToMQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDataToMQ) ProtoMessage() {}

func (x *MsgDataToMQ) ProtoReflect() protoreflect.Message {
	mi := &file_chart_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDataToMQ.ProtoReflect.Descriptor instead.
func (*MsgDataToMQ) Descriptor() ([]byte, []int) {
	return file_chart_chat_proto_rawDescGZIP(), []int{0}
}

func (x *MsgDataToMQ) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *MsgDataToMQ) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *MsgDataToMQ) GetMsgData() *pb_ws.MsgData {
	if x != nil {
		return x.MsgData
	}
	return nil
}

type MsgDataToDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgData     *pb_ws.MsgData `protobuf:"bytes,1,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`
	OperationId string         `protobuf:"bytes,2,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
}

func (x *MsgDataToDB) Reset() {
	*x = MsgDataToDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDataToDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDataToDB) ProtoMessage() {}

func (x *MsgDataToDB) ProtoReflect() protoreflect.Message {
	mi := &file_chart_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDataToDB.ProtoReflect.Descriptor instead.
func (*MsgDataToDB) Descriptor() ([]byte, []int) {
	return file_chart_chat_proto_rawDescGZIP(), []int{1}
}

func (x *MsgDataToDB) GetMsgData() *pb_ws.MsgData {
	if x != nil {
		return x.MsgData
	}
	return nil
}

func (x *MsgDataToDB) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

type PushMsgDataToMQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationId  string         `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	MsgData      *pb_ws.MsgData `protobuf:"bytes,2,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`
	PushToUserId string         `protobuf:"bytes,3,opt,name=push_to_user_id,json=pushToUserId,proto3" json:"push_to_user_id,omitempty"`
}

func (x *PushMsgDataToMQ) Reset() {
	*x = PushMsgDataToMQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgDataToMQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgDataToMQ) ProtoMessage() {}

func (x *PushMsgDataToMQ) ProtoReflect() protoreflect.Message {
	mi := &file_chart_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgDataToMQ.ProtoReflect.Descriptor instead.
func (*PushMsgDataToMQ) Descriptor() ([]byte, []int) {
	return file_chart_chat_proto_rawDescGZIP(), []int{2}
}

func (x *PushMsgDataToMQ) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *PushMsgDataToMQ) GetMsgData() *pb_ws.MsgData {
	if x != nil {
		return x.MsgData
	}
	return nil
}

func (x *PushMsgDataToMQ) GetPushToUserId() string {
	if x != nil {
		return x.PushToUserId
	}
	return ""
}

type GetMinMaxSeqReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OperationId string `protobuf:"bytes,2,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
}

func (x *GetMinMaxSeqReq) Reset() {
	*x = GetMinMaxSeqReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMinMaxSeqReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMinMaxSeqReq) ProtoMessage() {}

func (x *GetMinMaxSeqReq) ProtoReflect() protoreflect.Message {
	mi := &file_chart_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMinMaxSeqReq.ProtoReflect.Descriptor instead.
func (*GetMinMaxSeqReq) Descriptor() ([]byte, []int) {
	return file_chart_chat_proto_rawDescGZIP(), []int{3}
}

func (x *GetMinMaxSeqReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetMinMaxSeqReq) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

type GetMinMaxSeqResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32  `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	MaxSeq  uint32 `protobuf:"varint,3,opt,name=max_seq,json=maxSeq,proto3" json:"max_seq,omitempty"`
	MinSeq  uint32 `protobuf:"varint,4,opt,name=min_seq,json=minSeq,proto3" json:"min_seq,omitempty"`
}

func (x *GetMinMaxSeqResp) Reset() {
	*x = GetMinMaxSeqResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMinMaxSeqResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMinMaxSeqResp) ProtoMessage() {}

func (x *GetMinMaxSeqResp) ProtoReflect() protoreflect.Message {
	mi := &file_chart_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMinMaxSeqResp.ProtoReflect.Descriptor instead.
func (*GetMinMaxSeqResp) Descriptor() ([]byte, []int) {
	return file_chart_chat_proto_rawDescGZIP(), []int{4}
}

func (x *GetMinMaxSeqResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *GetMinMaxSeqResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *GetMinMaxSeqResp) GetMaxSeq() uint32 {
	if x != nil {
		return x.MaxSeq
	}
	return 0
}

func (x *GetMinMaxSeqResp) GetMinSeq() uint32 {
	if x != nil {
		return x.MinSeq
	}
	return 0
}

type SendMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string         `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	OperationId string         `protobuf:"bytes,2,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	MsgData     *pb_ws.MsgData `protobuf:"bytes,3,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`
}

func (x *SendMsgReq) Reset() {
	*x = SendMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgReq) ProtoMessage() {}

func (x *SendMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_chart_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgReq.ProtoReflect.Descriptor instead.
func (*SendMsgReq) Descriptor() ([]byte, []int) {
	return file_chart_chat_proto_rawDescGZIP(), []int{5}
}

func (x *SendMsgReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SendMsgReq) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *SendMsgReq) GetMsgData() *pb_ws.MsgData {
	if x != nil {
		return x.MsgData
	}
	return nil
}

type SendMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode     int32  `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	ErrMsg      string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	ServerMsgId string `protobuf:"bytes,3,opt,name=server_msg_id,json=serverMsgId,proto3" json:"server_msg_id,omitempty"`
	ClientMsgId string `protobuf:"bytes,4,opt,name=client_msg_id,json=clientMsgId,proto3" json:"client_msg_id,omitempty"`
	SendTs      int64  `protobuf:"varint,5,opt,name=send_ts,json=sendTs,proto3" json:"send_ts,omitempty"`
}

func (x *SendMsgResp) Reset() {
	*x = SendMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgResp) ProtoMessage() {}

func (x *SendMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_chart_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgResp.ProtoReflect.Descriptor instead.
func (*SendMsgResp) Descriptor() ([]byte, []int) {
	return file_chart_chat_proto_rawDescGZIP(), []int{6}
}

func (x *SendMsgResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *SendMsgResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *SendMsgResp) GetServerMsgId() string {
	if x != nil {
		return x.ServerMsgId
	}
	return ""
}

func (x *SendMsgResp) GetClientMsgId() string {
	if x != nil {
		return x.ClientMsgId
	}
	return ""
}

func (x *SendMsgResp) GetSendTs() int64 {
	if x != nil {
		return x.SendTs
	}
	return 0
}

type GetHistoryMessagesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize    int32  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Seq         int64  `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	SessionId   string `protobuf:"bytes,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	SessionType int32  `protobuf:"varint,5,opt,name=session_type,json=sessionType,proto3" json:"session_type,omitempty"`
	Back        bool   `protobuf:"varint,6,opt,name=back,proto3" json:"back,omitempty"`
}

func (x *GetHistoryMessagesReq) Reset() {
	*x = GetHistoryMessagesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoryMessagesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryMessagesReq) ProtoMessage() {}

func (x *GetHistoryMessagesReq) ProtoReflect() protoreflect.Message {
	mi := &file_chart_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryMessagesReq.ProtoReflect.Descriptor instead.
func (*GetHistoryMessagesReq) Descriptor() ([]byte, []int) {
	return file_chart_chat_proto_rawDescGZIP(), []int{7}
}

func (x *GetHistoryMessagesReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetHistoryMessagesReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetHistoryMessagesReq) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *GetHistoryMessagesReq) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *GetHistoryMessagesReq) GetSessionType() int32 {
	if x != nil {
		return x.SessionType
	}
	return 0
}

func (x *GetHistoryMessagesReq) GetBack() bool {
	if x != nil {
		return x.Back
	}
	return false
}

type GetHistoryMessagesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common  *pb_com.CommonResp `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	MsgList []*pb_ws.MsgData   `protobuf:"bytes,2,rep,name=msg_list,json=msgList,proto3" json:"msg_list,omitempty"`
}

func (x *GetHistoryMessagesResp) Reset() {
	*x = GetHistoryMessagesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chart_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoryMessagesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryMessagesResp) ProtoMessage() {}

func (x *GetHistoryMessagesResp) ProtoReflect() protoreflect.Message {
	mi := &file_chart_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryMessagesResp.ProtoReflect.Descriptor instead.
func (*GetHistoryMessagesResp) Descriptor() ([]byte, []int) {
	return file_chart_chat_proto_rawDescGZIP(), []int{8}
}

func (x *GetHistoryMessagesResp) GetCommon() *pb_com.CommonResp {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *GetHistoryMessagesResp) GetMsgList() []*pb_ws.MsgData {
	if x != nil {
		return x.MsgList
	}
	return nil
}

var File_chart_chat_proto protoreflect.FileDescriptor

var file_chart_chat_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x0e, 0x70, 0x62, 0x5f,
	0x77, 0x73, 0x2f, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x62, 0x5f,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x71, 0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x4d, 0x51, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f,
	0x77, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x5b, 0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x54, 0x6f,
	0x44, 0x42, 0x12, 0x29, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x77, 0x73, 0x2e, 0x4d, 0x73, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x86, 0x01, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x54, 0x6f, 0x4d, 0x51, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x77,
	0x73, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x25, 0x0a, 0x0f, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x75, 0x73,
	0x68, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4d, 0x69, 0x6e, 0x4d, 0x61, 0x78, 0x53, 0x65, 0x71, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d,
	0x69, 0x6e, 0x4d, 0x61, 0x78, 0x53, 0x65, 0x71, 0x52, 0x65, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67,
	0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x69, 0x6e,
	0x5f, 0x73, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x53,
	0x65, 0x71, 0x22, 0x70, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x6d, 0x73, 0x67,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62,
	0x5f, 0x77, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6d, 0x73, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x22, 0xa2, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x61, 0x63,
	0x6b, 0x22, 0x6f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62,
	0x5f, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52,
	0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x77,
	0x73, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x32, 0xb1, 0x02, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x43, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x4d, 0x69, 0x6e, 0x4d, 0x61, 0x78, 0x53, 0x65, 0x71, 0x12, 0x18, 0x2e, 0x70, 0x62,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x4d, 0x61, 0x78, 0x53,
	0x65, 0x71, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x4d, 0x61, 0x78, 0x53, 0x65, 0x71, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x57, 0x0a, 0x14, 0x50, 0x75, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x79, 0x53, 0x65, 0x71, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x5f, 0x77, 0x73,
	0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x53, 0x65,
	0x71, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x5f, 0x77, 0x73,
	0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x53, 0x65,
	0x71, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x07, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x73, 0x67, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x55, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x3b, 0x70, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chart_chat_proto_rawDescOnce sync.Once
	file_chart_chat_proto_rawDescData = file_chart_chat_proto_rawDesc
)

func file_chart_chat_proto_rawDescGZIP() []byte {
	file_chart_chat_proto_rawDescOnce.Do(func() {
		file_chart_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chart_chat_proto_rawDescData)
	})
	return file_chart_chat_proto_rawDescData
}

var file_chart_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_chart_chat_proto_goTypes = []interface{}{
	(*MsgDataToMQ)(nil),                    // 0: pb_auth.MsgDataToMQ
	(*MsgDataToDB)(nil),                    // 1: pb_auth.MsgDataToDB
	(*PushMsgDataToMQ)(nil),                // 2: pb_auth.PushMsgDataToMQ
	(*GetMinMaxSeqReq)(nil),                // 3: pb_auth.GetMinMaxSeqReq
	(*GetMinMaxSeqResp)(nil),               // 4: pb_auth.GetMinMaxSeqResp
	(*SendMsgReq)(nil),                     // 5: pb_auth.SendMsgReq
	(*SendMsgResp)(nil),                    // 6: pb_auth.SendMsgResp
	(*GetHistoryMessagesReq)(nil),          // 7: pb_auth.GetHistoryMessagesReq
	(*GetHistoryMessagesResp)(nil),         // 8: pb_auth.GetHistoryMessagesResp
	(*pb_ws.MsgData)(nil),                  // 9: pb_ws.MsgData
	(*pb_com.CommonResp)(nil),              // 10: pb_com.CommonResp
	(*pb_ws.PullMessageBySeqListReq)(nil),  // 11: pb_ws.PullMessageBySeqListReq
	(*pb_ws.PullMessageBySeqListResp)(nil), // 12: pb_ws.PullMessageBySeqListResp
}
var file_chart_chat_proto_depIdxs = []int32{
	9,  // 0: pb_auth.MsgDataToMQ.msg_data:type_name -> pb_ws.MsgData
	9,  // 1: pb_auth.MsgDataToDB.msg_data:type_name -> pb_ws.MsgData
	9,  // 2: pb_auth.PushMsgDataToMQ.msg_data:type_name -> pb_ws.MsgData
	9,  // 3: pb_auth.SendMsgReq.msg_data:type_name -> pb_ws.MsgData
	10, // 4: pb_auth.GetHistoryMessagesResp.common:type_name -> pb_com.CommonResp
	9,  // 5: pb_auth.GetHistoryMessagesResp.msg_list:type_name -> pb_ws.MsgData
	3,  // 6: pb_auth.Chat.GetMinMaxSeq:input_type -> pb_auth.GetMinMaxSeqReq
	11, // 7: pb_auth.Chat.PullMessageBySeqList:input_type -> pb_ws.PullMessageBySeqListReq
	5,  // 8: pb_auth.Chat.SendMsg:input_type -> pb_auth.SendMsgReq
	7,  // 9: pb_auth.Chat.GetHistoryMessages:input_type -> pb_auth.GetHistoryMessagesReq
	4,  // 10: pb_auth.Chat.GetMinMaxSeq:output_type -> pb_auth.GetMinMaxSeqResp
	12, // 11: pb_auth.Chat.PullMessageBySeqList:output_type -> pb_ws.PullMessageBySeqListResp
	6,  // 12: pb_auth.Chat.SendMsg:output_type -> pb_auth.SendMsgResp
	8,  // 13: pb_auth.Chat.GetHistoryMessages:output_type -> pb_auth.GetHistoryMessagesResp
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_chart_chat_proto_init() }
func file_chart_chat_proto_init() {
	if File_chart_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chart_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDataToMQ); i {
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
		file_chart_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDataToDB); i {
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
		file_chart_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgDataToMQ); i {
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
		file_chart_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMinMaxSeqReq); i {
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
		file_chart_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMinMaxSeqResp); i {
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
		file_chart_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgReq); i {
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
		file_chart_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgResp); i {
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
		file_chart_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoryMessagesReq); i {
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
		file_chart_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoryMessagesResp); i {
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
			RawDescriptor: file_chart_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chart_chat_proto_goTypes,
		DependencyIndexes: file_chart_chat_proto_depIdxs,
		MessageInfos:      file_chart_chat_proto_msgTypes,
	}.Build()
	File_chart_chat_proto = out.File
	file_chart_chat_proto_rawDesc = nil
	file_chart_chat_proto_goTypes = nil
	file_chart_chat_proto_depIdxs = nil
}
