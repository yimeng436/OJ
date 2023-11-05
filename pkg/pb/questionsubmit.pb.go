// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.9
// source: questionsubmit.proto

package pb

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

type QuestionSubmitInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                 // id
	Language    string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`      // 编程语言
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`              // 用户代码
	JudgeInfo   string `protobuf:"bytes,4,opt,name=judgeInfo,proto3" json:"judgeInfo,omitempty"`    // 判题信息（json 对象）
	Status      int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`         // 判题状态（0 - 待判题、1 - 判题中、2 - 成功、3 - 失败）
	QuestionId  int64  `protobuf:"varint,6,opt,name=questionId,proto3" json:"questionId,omitempty"` // 题目 id
	UserId      int64  `protobuf:"varint,7,opt,name=userId,proto3" json:"userId,omitempty"`         // 创建用户 id
	CreateTime  string `protobuf:"bytes,8,opt,name=createTime,proto3" json:"createTime,omitempty"`  // 创建时间
	UpdateTime  string `protobuf:"bytes,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`  // 更新时间
	IsDelete    int32  `protobuf:"varint,10,opt,name=isDelete,proto3" json:"isDelete,omitempty"`    // 是否删除
	JudgeStatus int32  `protobuf:"varint,11,opt,name=judgeStatus,proto3" json:"judgeStatus,omitempty"`
}

func (x *QuestionSubmitInfo) Reset() {
	*x = QuestionSubmitInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionsubmit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionSubmitInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionSubmitInfo) ProtoMessage() {}

func (x *QuestionSubmitInfo) ProtoReflect() protoreflect.Message {
	mi := &file_questionsubmit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionSubmitInfo.ProtoReflect.Descriptor instead.
func (*QuestionSubmitInfo) Descriptor() ([]byte, []int) {
	return file_questionsubmit_proto_rawDescGZIP(), []int{0}
}

func (x *QuestionSubmitInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QuestionSubmitInfo) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *QuestionSubmitInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QuestionSubmitInfo) GetJudgeInfo() string {
	if x != nil {
		return x.JudgeInfo
	}
	return ""
}

func (x *QuestionSubmitInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *QuestionSubmitInfo) GetQuestionId() int64 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *QuestionSubmitInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *QuestionSubmitInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *QuestionSubmitInfo) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *QuestionSubmitInfo) GetIsDelete() int32 {
	if x != nil {
		return x.IsDelete
	}
	return 0
}

func (x *QuestionSubmitInfo) GetJudgeStatus() int32 {
	if x != nil {
		return x.JudgeStatus
	}
	return 0
}

// 提交结果信息
type JudgeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message     string     `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` //提交结果
	Time        int64      `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`      //耗费时间
	Memory      float32    `protobuf:"fixed32,3,opt,name=memory,proto3" json:"memory,omitempty"` //耗费内存
	ErrCase     *JudgeCase `protobuf:"bytes,4,opt,name=errCase,proto3" json:"errCase,omitempty"`
	JudgeStatus int32      `protobuf:"varint,5,opt,name=judgeStatus,proto3" json:"judgeStatus,omitempty"`
}

func (x *JudgeInfo) Reset() {
	*x = JudgeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionsubmit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeInfo) ProtoMessage() {}

func (x *JudgeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_questionsubmit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeInfo.ProtoReflect.Descriptor instead.
func (*JudgeInfo) Descriptor() ([]byte, []int) {
	return file_questionsubmit_proto_rawDescGZIP(), []int{1}
}

func (x *JudgeInfo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *JudgeInfo) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *JudgeInfo) GetMemory() float32 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *JudgeInfo) GetErrCase() *JudgeCase {
	if x != nil {
		return x.ErrCase
	}
	return nil
}

func (x *JudgeInfo) GetJudgeStatus() int32 {
	if x != nil {
		return x.JudgeStatus
	}
	return 0
}

type QuestionSubmitAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language   string   `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Code       string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	QuestionId int64    `protobuf:"varint,3,opt,name=questionId,proto3" json:"questionId,omitempty"`
	Ctx        *Context `protobuf:"bytes,4,opt,name=ctx,proto3" json:"ctx,omitempty"`
}

func (x *QuestionSubmitAddRequest) Reset() {
	*x = QuestionSubmitAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionsubmit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionSubmitAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionSubmitAddRequest) ProtoMessage() {}

func (x *QuestionSubmitAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_questionsubmit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionSubmitAddRequest.ProtoReflect.Descriptor instead.
func (*QuestionSubmitAddRequest) Descriptor() ([]byte, []int) {
	return file_questionsubmit_proto_rawDescGZIP(), []int{2}
}

func (x *QuestionSubmitAddRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *QuestionSubmitAddRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QuestionSubmitAddRequest) GetQuestionId() int64 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *QuestionSubmitAddRequest) GetCtx() *Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

type QuestionSubmitQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language   string   `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Status     int64    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	QuestionId int64    `protobuf:"varint,3,opt,name=questionId,proto3" json:"questionId,omitempty"`
	UserId     int64    `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
	Page       *Page    `protobuf:"bytes,5,opt,name=page,proto3" json:"page,omitempty"`
	Ctx        *Context `protobuf:"bytes,6,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Id         int64    `protobuf:"varint,7,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *QuestionSubmitQueryRequest) Reset() {
	*x = QuestionSubmitQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionsubmit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionSubmitQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionSubmitQueryRequest) ProtoMessage() {}

func (x *QuestionSubmitQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_questionsubmit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionSubmitQueryRequest.ProtoReflect.Descriptor instead.
func (*QuestionSubmitQueryRequest) Descriptor() ([]byte, []int) {
	return file_questionsubmit_proto_rawDescGZIP(), []int{3}
}

func (x *QuestionSubmitQueryRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *QuestionSubmitQueryRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *QuestionSubmitQueryRequest) GetQuestionId() int64 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *QuestionSubmitQueryRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *QuestionSubmitQueryRequest) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QuestionSubmitQueryRequest) GetCtx() *Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *QuestionSubmitQueryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type QuestionSubmitQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionSubmitVO []*QuestionSubmitVo `protobuf:"bytes,1,rep,name=questionSubmitVO,proto3" json:"questionSubmitVO,omitempty"`
	Total            int64               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *QuestionSubmitQueryResponse) Reset() {
	*x = QuestionSubmitQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionsubmit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionSubmitQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionSubmitQueryResponse) ProtoMessage() {}

func (x *QuestionSubmitQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_questionsubmit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionSubmitQueryResponse.ProtoReflect.Descriptor instead.
func (*QuestionSubmitQueryResponse) Descriptor() ([]byte, []int) {
	return file_questionsubmit_proto_rawDescGZIP(), []int{4}
}

func (x *QuestionSubmitQueryResponse) GetQuestionSubmitVO() []*QuestionSubmitVo {
	if x != nil {
		return x.QuestionSubmitVO
	}
	return nil
}

func (x *QuestionSubmitQueryResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type QuestionSubmitVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language     string      `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Code         string      `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	JudgeInfo    *JudgeInfo  `protobuf:"bytes,3,opt,name=judgeInfo,proto3" json:"judgeInfo,omitempty"`
	Status       int32       `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	QuestionId   int64       `protobuf:"varint,5,opt,name=questionId,proto3" json:"questionId,omitempty"`
	UserId       int64       `protobuf:"varint,6,opt,name=userId,proto3" json:"userId,omitempty"`
	CreateTime   string      `protobuf:"bytes,7,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime   string      `protobuf:"bytes,8,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	SubmitUser   *UserVo     `protobuf:"bytes,9,opt,name=submitUser,proto3" json:"submitUser,omitempty"`
	QuestionVo   *QuestionVo `protobuf:"bytes,10,opt,name=questionVo,proto3" json:"questionVo,omitempty"`
	SubmitStatus string      `protobuf:"bytes,11,opt,name=submitStatus,proto3" json:"submitStatus,omitempty"`
}

func (x *QuestionSubmitVo) Reset() {
	*x = QuestionSubmitVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionsubmit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionSubmitVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionSubmitVo) ProtoMessage() {}

func (x *QuestionSubmitVo) ProtoReflect() protoreflect.Message {
	mi := &file_questionsubmit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionSubmitVo.ProtoReflect.Descriptor instead.
func (*QuestionSubmitVo) Descriptor() ([]byte, []int) {
	return file_questionsubmit_proto_rawDescGZIP(), []int{5}
}

func (x *QuestionSubmitVo) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *QuestionSubmitVo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QuestionSubmitVo) GetJudgeInfo() *JudgeInfo {
	if x != nil {
		return x.JudgeInfo
	}
	return nil
}

func (x *QuestionSubmitVo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *QuestionSubmitVo) GetQuestionId() int64 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *QuestionSubmitVo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *QuestionSubmitVo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *QuestionSubmitVo) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *QuestionSubmitVo) GetSubmitUser() *UserVo {
	if x != nil {
		return x.SubmitUser
	}
	return nil
}

func (x *QuestionSubmitVo) GetQuestionVo() *QuestionVo {
	if x != nil {
		return x.QuestionVo
	}
	return nil
}

func (x *QuestionSubmitVo) GetSubmitStatus() string {
	if x != nil {
		return x.SubmitStatus
	}
	return ""
}

var File_questionsubmit_proto protoreflect.FileDescriptor

var file_questionsubmit_proto_rawDesc = []byte{
	0x0a, 0x14, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc0, 0x02, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x99, 0x01, 0x0a, 0x09, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x61,
	0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x86, 0x01, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x03,
	0x63, 0x74, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x22, 0xcf, 0x01, 0x0a, 0x1a, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x72, 0x0a, 0x1b, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x56, 0x4f, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x56, 0x6f, 0x52, 0x10, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x56, 0x4f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xf6,
	0x02, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x56, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x56, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x6f, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x56, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xd6, 0x02, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x10, 0x44, 0x6f, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x42, 0x79, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x0a, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x3e, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x0d, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_questionsubmit_proto_rawDescOnce sync.Once
	file_questionsubmit_proto_rawDescData = file_questionsubmit_proto_rawDesc
)

func file_questionsubmit_proto_rawDescGZIP() []byte {
	file_questionsubmit_proto_rawDescOnce.Do(func() {
		file_questionsubmit_proto_rawDescData = protoimpl.X.CompressGZIP(file_questionsubmit_proto_rawDescData)
	})
	return file_questionsubmit_proto_rawDescData
}

var file_questionsubmit_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_questionsubmit_proto_goTypes = []interface{}{
	(*QuestionSubmitInfo)(nil),          // 0: QuestionSubmitInfo
	(*JudgeInfo)(nil),                   // 1: JudgeInfo
	(*QuestionSubmitAddRequest)(nil),    // 2: QuestionSubmitAddRequest
	(*QuestionSubmitQueryRequest)(nil),  // 3: QuestionSubmitQueryRequest
	(*QuestionSubmitQueryResponse)(nil), // 4: QuestionSubmitQueryResponse
	(*QuestionSubmitVo)(nil),            // 5: QuestionSubmitVo
	(*JudgeCase)(nil),                   // 6: JudgeCase
	(*Context)(nil),                     // 7: Context
	(*Page)(nil),                        // 8: Page
	(*UserVo)(nil),                      // 9: UserVo
	(*QuestionVo)(nil),                  // 10: QuestionVo
	(*Empty)(nil),                       // 11: Empty
	(*IdReQuest)(nil),                   // 12: IdReQuest
	(*IdResponse)(nil),                  // 13: IdResponse
	(*TotalResponse)(nil),               // 14: TotalResponse
	(*BoolResponse)(nil),                // 15: BoolResponse
}
var file_questionsubmit_proto_depIdxs = []int32{
	6,  // 0: JudgeInfo.errCase:type_name -> JudgeCase
	7,  // 1: QuestionSubmitAddRequest.ctx:type_name -> Context
	8,  // 2: QuestionSubmitQueryRequest.page:type_name -> Page
	7,  // 3: QuestionSubmitQueryRequest.ctx:type_name -> Context
	5,  // 4: QuestionSubmitQueryResponse.questionSubmitVO:type_name -> QuestionSubmitVo
	1,  // 5: QuestionSubmitVo.judgeInfo:type_name -> JudgeInfo
	9,  // 6: QuestionSubmitVo.submitUser:type_name -> UserVo
	10, // 7: QuestionSubmitVo.questionVo:type_name -> QuestionVo
	2,  // 8: QuestionSubmitService.DoQuestionSubmit:input_type -> QuestionSubmitAddRequest
	3,  // 9: QuestionSubmitService.ListQuestionSubmitByPage:input_type -> QuestionSubmitQueryRequest
	11, // 10: QuestionSubmitService.GetQuestionSubmitTotal:input_type -> Empty
	12, // 11: QuestionSubmitService.GetQuestionSubmitById:input_type -> IdReQuest
	0,  // 12: QuestionSubmitService.UpdateQuestionStatusById:input_type -> QuestionSubmitInfo
	13, // 13: QuestionSubmitService.DoQuestionSubmit:output_type -> IdResponse
	4,  // 14: QuestionSubmitService.ListQuestionSubmitByPage:output_type -> QuestionSubmitQueryResponse
	14, // 15: QuestionSubmitService.GetQuestionSubmitTotal:output_type -> TotalResponse
	0,  // 16: QuestionSubmitService.GetQuestionSubmitById:output_type -> QuestionSubmitInfo
	15, // 17: QuestionSubmitService.UpdateQuestionStatusById:output_type -> BoolResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_questionsubmit_proto_init() }
func file_questionsubmit_proto_init() {
	if File_questionsubmit_proto != nil {
		return
	}
	file_common_proto_init()
	file_user_proto_init()
	file_question_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_questionsubmit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionSubmitInfo); i {
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
		file_questionsubmit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JudgeInfo); i {
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
		file_questionsubmit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionSubmitAddRequest); i {
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
		file_questionsubmit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionSubmitQueryRequest); i {
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
		file_questionsubmit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionSubmitQueryResponse); i {
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
		file_questionsubmit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionSubmitVo); i {
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
			RawDescriptor: file_questionsubmit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_questionsubmit_proto_goTypes,
		DependencyIndexes: file_questionsubmit_proto_depIdxs,
		MessageInfos:      file_questionsubmit_proto_msgTypes,
	}.Build()
	File_questionsubmit_proto = out.File
	file_questionsubmit_proto_rawDesc = nil
	file_questionsubmit_proto_goTypes = nil
	file_questionsubmit_proto_depIdxs = nil
}
