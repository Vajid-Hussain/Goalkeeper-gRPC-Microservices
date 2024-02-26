// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/pb/reminder.proto

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

type TodayTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *TodayTaskRequest) Reset() {
	*x = TodayTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_reminder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodayTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodayTaskRequest) ProtoMessage() {}

func (x *TodayTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_reminder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodayTaskRequest.ProtoReflect.Descriptor instead.
func (*TodayTaskRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_reminder_proto_rawDescGZIP(), []int{0}
}

func (x *TodayTaskRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type TodayTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *TodayTaskResponse) Reset() {
	*x = TodayTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_reminder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodayTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodayTaskResponse) ProtoMessage() {}

func (x *TodayTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_reminder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodayTaskResponse.ProtoReflect.Descriptor instead.
func (*TodayTaskResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_reminder_proto_rawDescGZIP(), []int{1}
}

func (x *TodayTaskResponse) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_pb_reminder_proto protoreflect.FileDescriptor

var file_pkg_pb_reminder_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x22, 0x2a, 0x0a, 0x10, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x27,
	0x0a, 0x11, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x5c, 0x0a, 0x10, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x54,
	0x6f, 0x64, 0x61, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x2e, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_reminder_proto_rawDescOnce sync.Once
	file_pkg_pb_reminder_proto_rawDescData = file_pkg_pb_reminder_proto_rawDesc
)

func file_pkg_pb_reminder_proto_rawDescGZIP() []byte {
	file_pkg_pb_reminder_proto_rawDescOnce.Do(func() {
		file_pkg_pb_reminder_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_reminder_proto_rawDescData)
	})
	return file_pkg_pb_reminder_proto_rawDescData
}

var file_pkg_pb_reminder_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_pb_reminder_proto_goTypes = []interface{}{
	(*TodayTaskRequest)(nil),  // 0: remainder.todayTaskRequest
	(*TodayTaskResponse)(nil), // 1: remainder.todayTaskResponse
}
var file_pkg_pb_reminder_proto_depIdxs = []int32{
	0, // 0: remainder.remainderService.TodayTask:input_type -> remainder.todayTaskRequest
	1, // 1: remainder.remainderService.TodayTask:output_type -> remainder.todayTaskResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pb_reminder_proto_init() }
func file_pkg_pb_reminder_proto_init() {
	if File_pkg_pb_reminder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_reminder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodayTaskRequest); i {
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
		file_pkg_pb_reminder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodayTaskResponse); i {
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
			RawDescriptor: file_pkg_pb_reminder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_reminder_proto_goTypes,
		DependencyIndexes: file_pkg_pb_reminder_proto_depIdxs,
		MessageInfos:      file_pkg_pb_reminder_proto_msgTypes,
	}.Build()
	File_pkg_pb_reminder_proto = out.File
	file_pkg_pb_reminder_proto_rawDesc = nil
	file_pkg_pb_reminder_proto_goTypes = nil
	file_pkg_pb_reminder_proto_depIdxs = nil
}
