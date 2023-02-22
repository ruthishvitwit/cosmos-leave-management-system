// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/cosmos/std/v1beta1/tx.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LeaveStatus int32

const (
	LeaveStatus_STATUS_UNDEFINED LeaveStatus = 0
	LeaveStatus_STATUS_ACCEPTED  LeaveStatus = 1
	LeaveStatus_STATUS_REJECTED  LeaveStatus = 2
)

// Enum value maps for LeaveStatus.
var (
	LeaveStatus_name = map[int32]string{
		0: "STATUS_UNDEFINED",
		1: "STATUS_ACCEPTED",
		2: "STATUS_REJECTED",
	}
	LeaveStatus_value = map[string]int32{
		"STATUS_UNDEFINED": 0,
		"STATUS_ACCEPTED":  1,
		"STATUS_REJECTED":  2,
	}
)

func (x LeaveStatus) Enum() *LeaveStatus {
	p := new(LeaveStatus)
	*p = x
	return p
}

func (x LeaveStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LeaveStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_cosmos_std_v1beta1_tx_proto_enumTypes[0].Descriptor()
}

func (LeaveStatus) Type() protoreflect.EnumType {
	return &file_proto_cosmos_std_v1beta1_tx_proto_enumTypes[0]
}

func (x LeaveStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LeaveStatus.Descriptor instead.
func (LeaveStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_cosmos_std_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

type RegisterAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the account address of the admin.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// name is the admin name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RegisterAdminRequest) Reset() {
	*x = RegisterAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAdminRequest) ProtoMessage() {}

func (x *RegisterAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAdminRequest.ProtoReflect.Descriptor instead.
func (*RegisterAdminRequest) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_std_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterAdminRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegisterAdminRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegisterAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterAdminResponse) Reset() {
	*x = RegisterAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAdminResponse) ProtoMessage() {}

func (x *RegisterAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAdminResponse.ProtoReflect.Descriptor instead.
func (*RegisterAdminResponse) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_std_v1beta1_tx_proto_rawDescGZIP(), []int{1}
}

type AddStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the admin
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// students is the list of Student.
	Students []*Student `protobuf:"bytes,2,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *AddStudentRequest) Reset() {
	*x = AddStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStudentRequest) ProtoMessage() {}

func (x *AddStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStudentRequest.ProtoReflect.Descriptor instead.
func (*AddStudentRequest) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_std_v1beta1_tx_proto_rawDescGZIP(), []int{2}
}

func (x *AddStudentRequest) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *AddStudentRequest) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

type AddStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddStudentResponse) Reset() {
	*x = AddStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStudentResponse) ProtoMessage() {}

func (x *AddStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStudentResponse.ProtoReflect.Descriptor instead.
func (*AddStudentResponse) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_std_v1beta1_tx_proto_rawDescGZIP(), []int{3}
}

type ApplyLeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Reason  string                 `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	From    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *ApplyLeaveRequest) Reset() {
	*x = ApplyLeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyLeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyLeaveRequest) ProtoMessage() {}

func (x *ApplyLeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyLeaveRequest.ProtoReflect.Descriptor instead.
func (*ApplyLeaveRequest) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_std_v1beta1_tx_proto_rawDescGZIP(), []int{4}
}

func (x *ApplyLeaveRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ApplyLeaveRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ApplyLeaveRequest) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *ApplyLeaveRequest) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

type ApplyLeaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApplyLeaveResponse) Reset() {
	*x = ApplyLeaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyLeaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyLeaveResponse) ProtoMessage() {}

func (x *ApplyLeaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyLeaveResponse.ProtoReflect.Descriptor instead.
func (*ApplyLeaveResponse) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_std_v1beta1_tx_proto_rawDescGZIP(), []int{5}
}

type AcceptLeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin   string      `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	LeaveId uint64      `protobuf:"varint,2,opt,name=leave_id,json=leaveId,proto3" json:"leave_id,omitempty"`
	Status  LeaveStatus `protobuf:"varint,3,opt,name=status,proto3,enum=std.v1beta1.LeaveStatus" json:"status,omitempty"`
}

func (x *AcceptLeaveRequest) Reset() {
	*x = AcceptLeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptLeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptLeaveRequest) ProtoMessage() {}

func (x *AcceptLeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptLeaveRequest.ProtoReflect.Descriptor instead.
func (*AcceptLeaveRequest) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_std_v1beta1_tx_proto_rawDescGZIP(), []int{6}
}

func (x *AcceptLeaveRequest) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *AcceptLeaveRequest) GetLeaveId() uint64 {
	if x != nil {
		return x.LeaveId
	}
	return 0
}

func (x *AcceptLeaveRequest) GetStatus() LeaveStatus {
	if x != nil {
		return x.Status
	}
	return LeaveStatus_STATUS_UNDEFINED
}

type AcceptLeaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AcceptLeaveResponse) Reset() {
	*x = AcceptLeaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptLeaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptLeaveResponse) ProtoMessage() {}

func (x *AcceptLeaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptLeaveResponse.ProtoReflect.Descriptor instead.
func (*AcceptLeaveResponse) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_std_v1beta1_tx_proto_rawDescGZIP(), []int{7}
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the account address of the student.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// name is the student name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// id is a unique student id.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_std_v1beta1_tx_proto_rawDescGZIP(), []int{8}
}

func (x *Student) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_cosmos_std_v1beta1_tx_proto protoreflect.FileDescriptor

var file_proto_cosmos_std_v1beta1_tx_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73,
	0x74, 0x64, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x44, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x5b, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x30, 0x0a, 0x08, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x14, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x77, 0x0a,
	0x12, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x61,
	0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6c, 0x65, 0x61,
	0x76, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x0a,
	0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x4d, 0x0a, 0x0b, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0xcd, 0x02, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x56, 0x0a,
	0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x21,
	0x2e, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x78, 0x2f, 0x73, 0x74, 0x64, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cosmos_std_v1beta1_tx_proto_rawDescOnce sync.Once
	file_proto_cosmos_std_v1beta1_tx_proto_rawDescData = file_proto_cosmos_std_v1beta1_tx_proto_rawDesc
)

func file_proto_cosmos_std_v1beta1_tx_proto_rawDescGZIP() []byte {
	file_proto_cosmos_std_v1beta1_tx_proto_rawDescOnce.Do(func() {
		file_proto_cosmos_std_v1beta1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cosmos_std_v1beta1_tx_proto_rawDescData)
	})
	return file_proto_cosmos_std_v1beta1_tx_proto_rawDescData
}

var file_proto_cosmos_std_v1beta1_tx_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_cosmos_std_v1beta1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_cosmos_std_v1beta1_tx_proto_goTypes = []interface{}{
	(LeaveStatus)(0),              // 0: std.v1beta1.LeaveStatus
	(*RegisterAdminRequest)(nil),  // 1: std.v1beta1.RegisterAdminRequest
	(*RegisterAdminResponse)(nil), // 2: std.v1beta1.RegisterAdminResponse
	(*AddStudentRequest)(nil),     // 3: std.v1beta1.AddStudentRequest
	(*AddStudentResponse)(nil),    // 4: std.v1beta1.AddStudentResponse
	(*ApplyLeaveRequest)(nil),     // 5: std.v1beta1.ApplyLeaveRequest
	(*ApplyLeaveResponse)(nil),    // 6: std.v1beta1.ApplyLeaveResponse
	(*AcceptLeaveRequest)(nil),    // 7: std.v1beta1.AcceptLeaveRequest
	(*AcceptLeaveResponse)(nil),   // 8: std.v1beta1.AcceptLeaveResponse
	(*Student)(nil),               // 9: std.v1beta1.Student
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_proto_cosmos_std_v1beta1_tx_proto_depIdxs = []int32{
	9,  // 0: std.v1beta1.AddStudentRequest.students:type_name -> std.v1beta1.Student
	10, // 1: std.v1beta1.ApplyLeaveRequest.from:type_name -> google.protobuf.Timestamp
	10, // 2: std.v1beta1.ApplyLeaveRequest.to:type_name -> google.protobuf.Timestamp
	0,  // 3: std.v1beta1.AcceptLeaveRequest.status:type_name -> std.v1beta1.LeaveStatus
	1,  // 4: std.v1beta1.Msg.RegisterAdmin:input_type -> std.v1beta1.RegisterAdminRequest
	3,  // 5: std.v1beta1.Msg.AddStudent:input_type -> std.v1beta1.AddStudentRequest
	5,  // 6: std.v1beta1.Msg.ApplyLeave:input_type -> std.v1beta1.ApplyLeaveRequest
	7,  // 7: std.v1beta1.Msg.AcceptLeave:input_type -> std.v1beta1.AcceptLeaveRequest
	2,  // 8: std.v1beta1.Msg.RegisterAdmin:output_type -> std.v1beta1.RegisterAdminResponse
	4,  // 9: std.v1beta1.Msg.AddStudent:output_type -> std.v1beta1.AddStudentResponse
	6,  // 10: std.v1beta1.Msg.ApplyLeave:output_type -> std.v1beta1.ApplyLeaveResponse
	8,  // 11: std.v1beta1.Msg.AcceptLeave:output_type -> std.v1beta1.AcceptLeaveResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_cosmos_std_v1beta1_tx_proto_init() }
func file_proto_cosmos_std_v1beta1_tx_proto_init() {
	if File_proto_cosmos_std_v1beta1_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAdminRequest); i {
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
		file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAdminResponse); i {
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
		file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStudentRequest); i {
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
		file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStudentResponse); i {
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
		file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyLeaveRequest); i {
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
		file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyLeaveResponse); i {
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
		file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptLeaveRequest); i {
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
		file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptLeaveResponse); i {
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
		file_proto_cosmos_std_v1beta1_tx_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
			RawDescriptor: file_proto_cosmos_std_v1beta1_tx_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cosmos_std_v1beta1_tx_proto_goTypes,
		DependencyIndexes: file_proto_cosmos_std_v1beta1_tx_proto_depIdxs,
		EnumInfos:         file_proto_cosmos_std_v1beta1_tx_proto_enumTypes,
		MessageInfos:      file_proto_cosmos_std_v1beta1_tx_proto_msgTypes,
	}.Build()
	File_proto_cosmos_std_v1beta1_tx_proto = out.File
	file_proto_cosmos_std_v1beta1_tx_proto_rawDesc = nil
	file_proto_cosmos_std_v1beta1_tx_proto_goTypes = nil
	file_proto_cosmos_std_v1beta1_tx_proto_depIdxs = nil
}
