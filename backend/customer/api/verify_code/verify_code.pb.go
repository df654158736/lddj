// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: api/verify_code/verify_code.proto

package s

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

type Type int32

const (
	Type_DEFAULT Type = 0
	Type_DIGIT   Type = 1
	Type_LETTER  Type = 2
	Type_MIXED   Type = 3
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "DEFAULT",
		1: "DIGIT",
		2: "LETTER",
		3: "MIXED",
	}
	Type_value = map[string]int32{
		"DEFAULT": 0,
		"DIGIT":   1,
		"LETTER":  2,
		"MIXED":   3,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_verify_code_verify_code_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_api_verify_code_verify_code_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_api_verify_code_verify_code_proto_rawDescGZIP(), []int{0}
}

type CreateVerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateVerifyCodeRequest) Reset() {
	*x = CreateVerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_verify_code_verify_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVerifyCodeRequest) ProtoMessage() {}

func (x *CreateVerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_verify_code_verify_code_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*CreateVerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_api_verify_code_verify_code_proto_rawDescGZIP(), []int{0}
}

type CreateVerifyCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateVerifyCodeReply) Reset() {
	*x = CreateVerifyCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_verify_code_verify_code_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVerifyCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVerifyCodeReply) ProtoMessage() {}

func (x *CreateVerifyCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_verify_code_verify_code_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVerifyCodeReply.ProtoReflect.Descriptor instead.
func (*CreateVerifyCodeReply) Descriptor() ([]byte, []int) {
	return file_api_verify_code_verify_code_proto_rawDescGZIP(), []int{1}
}

type UpdateVerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateVerifyCodeRequest) Reset() {
	*x = UpdateVerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_verify_code_verify_code_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVerifyCodeRequest) ProtoMessage() {}

func (x *UpdateVerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_verify_code_verify_code_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*UpdateVerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_api_verify_code_verify_code_proto_rawDescGZIP(), []int{2}
}

type UpdateVerifyCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateVerifyCodeReply) Reset() {
	*x = UpdateVerifyCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_verify_code_verify_code_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVerifyCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVerifyCodeReply) ProtoMessage() {}

func (x *UpdateVerifyCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_verify_code_verify_code_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVerifyCodeReply.ProtoReflect.Descriptor instead.
func (*UpdateVerifyCodeReply) Descriptor() ([]byte, []int) {
	return file_api_verify_code_verify_code_proto_rawDescGZIP(), []int{3}
}

type DeleteVerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteVerifyCodeRequest) Reset() {
	*x = DeleteVerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_verify_code_verify_code_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVerifyCodeRequest) ProtoMessage() {}

func (x *DeleteVerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_verify_code_verify_code_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*DeleteVerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_api_verify_code_verify_code_proto_rawDescGZIP(), []int{4}
}

type DeleteVerifyCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteVerifyCodeReply) Reset() {
	*x = DeleteVerifyCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_verify_code_verify_code_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVerifyCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVerifyCodeReply) ProtoMessage() {}

func (x *DeleteVerifyCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_verify_code_verify_code_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVerifyCodeReply.ProtoReflect.Descriptor instead.
func (*DeleteVerifyCodeReply) Descriptor() ([]byte, []int) {
	return file_api_verify_code_verify_code_proto_rawDescGZIP(), []int{5}
}

type GetVerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Type   Type  `protobuf:"varint,2,opt,name=type,proto3,enum=s.v1.Type" json:"type,omitempty"`
}

func (x *GetVerifyCodeRequest) Reset() {
	*x = GetVerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_verify_code_verify_code_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVerifyCodeRequest) ProtoMessage() {}

func (x *GetVerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_verify_code_verify_code_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*GetVerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_api_verify_code_verify_code_proto_rawDescGZIP(), []int{6}
}

func (x *GetVerifyCodeRequest) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *GetVerifyCodeRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_DEFAULT
}

type GetVerifyCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetVerifyCodeReply) Reset() {
	*x = GetVerifyCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_verify_code_verify_code_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVerifyCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVerifyCodeReply) ProtoMessage() {}

func (x *GetVerifyCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_verify_code_verify_code_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVerifyCodeReply.ProtoReflect.Descriptor instead.
func (*GetVerifyCodeReply) Descriptor() ([]byte, []int) {
	return file_api_verify_code_verify_code_proto_rawDescGZIP(), []int{7}
}

func (x *GetVerifyCodeReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListVerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListVerifyCodeRequest) Reset() {
	*x = ListVerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_verify_code_verify_code_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVerifyCodeRequest) ProtoMessage() {}

func (x *ListVerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_verify_code_verify_code_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*ListVerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_api_verify_code_verify_code_proto_rawDescGZIP(), []int{8}
}

type ListVerifyCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListVerifyCodeReply) Reset() {
	*x = ListVerifyCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_verify_code_verify_code_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVerifyCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVerifyCodeReply) ProtoMessage() {}

func (x *ListVerifyCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_verify_code_verify_code_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVerifyCodeReply.ProtoReflect.Descriptor instead.
func (*ListVerifyCodeReply) Descriptor() ([]byte, []int) {
	return file_api_verify_code_verify_code_proto_rawDescGZIP(), []int{9}
}

var File_api_verify_code_verify_code_proto protoreflect.FileDescriptor

var file_api_verify_code_verify_code_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x4e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x35, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x49,
	0x47, 0x49, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x45, 0x54, 0x54, 0x45, 0x52, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x49, 0x58, 0x45, 0x44, 0x10, 0x03, 0x32, 0x8d, 0x03, 0x0a,
	0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x4e, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x2e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x10, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x2e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x2e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x45, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x2e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x48, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x1f, 0x0a, 0x04,
	0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x15, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_verify_code_verify_code_proto_rawDescOnce sync.Once
	file_api_verify_code_verify_code_proto_rawDescData = file_api_verify_code_verify_code_proto_rawDesc
)

func file_api_verify_code_verify_code_proto_rawDescGZIP() []byte {
	file_api_verify_code_verify_code_proto_rawDescOnce.Do(func() {
		file_api_verify_code_verify_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_verify_code_verify_code_proto_rawDescData)
	})
	return file_api_verify_code_verify_code_proto_rawDescData
}

var file_api_verify_code_verify_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_verify_code_verify_code_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_verify_code_verify_code_proto_goTypes = []any{
	(Type)(0),                       // 0: s.v1.Type
	(*CreateVerifyCodeRequest)(nil), // 1: s.v1.CreateVerifyCodeRequest
	(*CreateVerifyCodeReply)(nil),   // 2: s.v1.CreateVerifyCodeReply
	(*UpdateVerifyCodeRequest)(nil), // 3: s.v1.UpdateVerifyCodeRequest
	(*UpdateVerifyCodeReply)(nil),   // 4: s.v1.UpdateVerifyCodeReply
	(*DeleteVerifyCodeRequest)(nil), // 5: s.v1.DeleteVerifyCodeRequest
	(*DeleteVerifyCodeReply)(nil),   // 6: s.v1.DeleteVerifyCodeReply
	(*GetVerifyCodeRequest)(nil),    // 7: s.v1.GetVerifyCodeRequest
	(*GetVerifyCodeReply)(nil),      // 8: s.v1.GetVerifyCodeReply
	(*ListVerifyCodeRequest)(nil),   // 9: s.v1.ListVerifyCodeRequest
	(*ListVerifyCodeReply)(nil),     // 10: s.v1.ListVerifyCodeReply
}
var file_api_verify_code_verify_code_proto_depIdxs = []int32{
	0,  // 0: s.v1.GetVerifyCodeRequest.type:type_name -> s.v1.Type
	1,  // 1: s.v1.VerifyCode.CreateVerifyCode:input_type -> s.v1.CreateVerifyCodeRequest
	3,  // 2: s.v1.VerifyCode.UpdateVerifyCode:input_type -> s.v1.UpdateVerifyCodeRequest
	5,  // 3: s.v1.VerifyCode.DeleteVerifyCode:input_type -> s.v1.DeleteVerifyCodeRequest
	7,  // 4: s.v1.VerifyCode.GetVerifyCode:input_type -> s.v1.GetVerifyCodeRequest
	9,  // 5: s.v1.VerifyCode.ListVerifyCode:input_type -> s.v1.ListVerifyCodeRequest
	2,  // 6: s.v1.VerifyCode.CreateVerifyCode:output_type -> s.v1.CreateVerifyCodeReply
	4,  // 7: s.v1.VerifyCode.UpdateVerifyCode:output_type -> s.v1.UpdateVerifyCodeReply
	6,  // 8: s.v1.VerifyCode.DeleteVerifyCode:output_type -> s.v1.DeleteVerifyCodeReply
	8,  // 9: s.v1.VerifyCode.GetVerifyCode:output_type -> s.v1.GetVerifyCodeReply
	10, // 10: s.v1.VerifyCode.ListVerifyCode:output_type -> s.v1.ListVerifyCodeReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_verify_code_verify_code_proto_init() }
func file_api_verify_code_verify_code_proto_init() {
	if File_api_verify_code_verify_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_verify_code_verify_code_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateVerifyCodeRequest); i {
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
		file_api_verify_code_verify_code_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateVerifyCodeReply); i {
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
		file_api_verify_code_verify_code_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateVerifyCodeRequest); i {
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
		file_api_verify_code_verify_code_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateVerifyCodeReply); i {
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
		file_api_verify_code_verify_code_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteVerifyCodeRequest); i {
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
		file_api_verify_code_verify_code_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteVerifyCodeReply); i {
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
		file_api_verify_code_verify_code_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetVerifyCodeRequest); i {
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
		file_api_verify_code_verify_code_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetVerifyCodeReply); i {
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
		file_api_verify_code_verify_code_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ListVerifyCodeRequest); i {
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
		file_api_verify_code_verify_code_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ListVerifyCodeReply); i {
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
			RawDescriptor: file_api_verify_code_verify_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_verify_code_verify_code_proto_goTypes,
		DependencyIndexes: file_api_verify_code_verify_code_proto_depIdxs,
		EnumInfos:         file_api_verify_code_verify_code_proto_enumTypes,
		MessageInfos:      file_api_verify_code_verify_code_proto_msgTypes,
	}.Build()
	File_api_verify_code_verify_code_proto = out.File
	file_api_verify_code_verify_code_proto_rawDesc = nil
	file_api_verify_code_verify_code_proto_goTypes = nil
	file_api_verify_code_verify_code_proto_depIdxs = nil
}
