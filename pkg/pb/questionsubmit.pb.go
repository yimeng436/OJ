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

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                 // id
	Language   string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`      // 编程语言
	Code       string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`              // 用户代码
	JudgeInfo  string `protobuf:"bytes,4,opt,name=judgeInfo,proto3" json:"judgeInfo,omitempty"`    // 判题信息（json 对象）
	Status     int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`         // 判题状态（0 - 待判题、1 - 判题中、2 - 成功、3 - 失败）
	QuestionId int64  `protobuf:"varint,6,opt,name=questionId,proto3" json:"questionId,omitempty"` // 题目 id
	UserId     int64  `protobuf:"varint,7,opt,name=userId,proto3" json:"userId,omitempty"`         // 创建用户 id
	CreateTime string `protobuf:"bytes,8,opt,name=createTime,proto3" json:"createTime,omitempty"`  // 创建时间
	UpdateTime string `protobuf:"bytes,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`  // 更新时间
	IsDelete   int32  `protobuf:"varint,10,opt,name=isDelete,proto3" json:"isDelete,omitempty"`    // 是否删除
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

// 提交结果信息
type JudgeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` //提交结果
	Time    int64  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`      //耗费时间
	Memory  int64  `protobuf:"varint,3,opt,name=memory,proto3" json:"memory,omitempty"`  //耗费内存
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

func (x *JudgeInfo) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

var File_questionsubmit_proto protoreflect.FileDescriptor

var file_questionsubmit_proto_rawDesc = []byte{
	0x0a, 0x14, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69,
	0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x51, 0x0a, 0x09, 0x4a, 0x75, 0x64, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e,
	0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_questionsubmit_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_questionsubmit_proto_goTypes = []interface{}{
	(*QuestionSubmitInfo)(nil), // 0: QuestionSubmitInfo
	(*JudgeInfo)(nil),          // 1: JudgeInfo
}
var file_questionsubmit_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_questionsubmit_proto_init() }
func file_questionsubmit_proto_init() {
	if File_questionsubmit_proto != nil {
		return
	}
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_questionsubmit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
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
