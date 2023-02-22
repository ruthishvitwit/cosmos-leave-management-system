// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/cosmos/std/v1beta1/query.proto

package types

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

type GetStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStudentRequest) Reset() {
	*x = GetStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_std_v1beta1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentRequest) ProtoMessage() {}

func (x *GetStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_std_v1beta1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentRequest.ProtoReflect.Descriptor instead.
func (*GetStudentRequest) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_std_v1beta1_query_proto_rawDescGZIP(), []int{0}
}

type GetStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// students is the list of Student.
	Students []*Student `protobuf:"bytes,2,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *GetStudentResponse) Reset() {
	*x = GetStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_std_v1beta1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentResponse) ProtoMessage() {}

func (x *GetStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_std_v1beta1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentResponse.ProtoReflect.Descriptor instead.
func (*GetStudentResponse) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_std_v1beta1_query_proto_rawDescGZIP(), []int{1}
}

func (x *GetStudentResponse) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *GetStudentResponse) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

var File_proto_cosmos_std_v1beta1_query_proto protoreflect.FileDescriptor

var file_proto_cosmos_std_v1beta1_query_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73,
	0x74, 0x64, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x73, 0x74, 0x64, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5c, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x64, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x56, 0x0a, 0x05, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x4d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x12, 0x1e, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x78, 0x2f, 0x73, 0x74, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cosmos_std_v1beta1_query_proto_rawDescOnce sync.Once
	file_proto_cosmos_std_v1beta1_query_proto_rawDescData = file_proto_cosmos_std_v1beta1_query_proto_rawDesc
)

func file_proto_cosmos_std_v1beta1_query_proto_rawDescGZIP() []byte {
	file_proto_cosmos_std_v1beta1_query_proto_rawDescOnce.Do(func() {
		file_proto_cosmos_std_v1beta1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cosmos_std_v1beta1_query_proto_rawDescData)
	})
	return file_proto_cosmos_std_v1beta1_query_proto_rawDescData
}

var file_proto_cosmos_std_v1beta1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_cosmos_std_v1beta1_query_proto_goTypes = []interface{}{
	(*GetStudentRequest)(nil),  // 0: std.v1beta1.GetStudentRequest
	(*GetStudentResponse)(nil), // 1: std.v1beta1.GetStudentResponse
	(*Student)(nil),            // 2: std.v1beta1.Student
}
var file_proto_cosmos_std_v1beta1_query_proto_depIdxs = []int32{
	2, // 0: std.v1beta1.GetStudentResponse.students:type_name -> std.v1beta1.Student
	0, // 1: std.v1beta1.Query.GetStudent:input_type -> std.v1beta1.GetStudentRequest
	1, // 2: std.v1beta1.Query.GetStudent:output_type -> std.v1beta1.GetStudentResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_cosmos_std_v1beta1_query_proto_init() }
func file_proto_cosmos_std_v1beta1_query_proto_init() {
	if File_proto_cosmos_std_v1beta1_query_proto != nil {
		return
	}
	file_proto_cosmos_std_v1beta1_tx_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_cosmos_std_v1beta1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentRequest); i {
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
		file_proto_cosmos_std_v1beta1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentResponse); i {
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
			RawDescriptor: file_proto_cosmos_std_v1beta1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cosmos_std_v1beta1_query_proto_goTypes,
		DependencyIndexes: file_proto_cosmos_std_v1beta1_query_proto_depIdxs,
		MessageInfos:      file_proto_cosmos_std_v1beta1_query_proto_msgTypes,
	}.Build()
	File_proto_cosmos_std_v1beta1_query_proto = out.File
	file_proto_cosmos_std_v1beta1_query_proto_rawDesc = nil
	file_proto_cosmos_std_v1beta1_query_proto_goTypes = nil
	file_proto_cosmos_std_v1beta1_query_proto_depIdxs = nil
}
