// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: userApi.proto

package userApi

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type RequestRole int32

const (
	Request_user  RequestRole = 0
	Request_admin RequestRole = 1
)

// Enum value maps for RequestRole.
var (
	RequestRole_name = map[int32]string{
		0: "user",
		1: "admin",
	}
	RequestRole_value = map[string]int32{
		"user":  0,
		"admin": 1,
	}
)

func (x RequestRole) Enum() *RequestRole {
	p := new(RequestRole)
	*p = x
	return p
}

func (x RequestRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestRole) Descriptor() protoreflect.EnumDescriptor {
	return file_userApi_proto_enumTypes[0].Descriptor()
}

func (RequestRole) Type() protoreflect.EnumType {
	return &file_userApi_proto_enumTypes[0]
}

func (x RequestRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestRole.Descriptor instead.
func (RequestRole) EnumDescriptor() ([]byte, []int) {
	return file_userApi_proto_rawDescGZIP(), []int{0, 0}
}

type GetResponseRole int32

const (
	GetResponse_user  GetResponseRole = 0
	GetResponse_admin GetResponseRole = 1
)

// Enum value maps for GetResponseRole.
var (
	GetResponseRole_name = map[int32]string{
		0: "user",
		1: "admin",
	}
	GetResponseRole_value = map[string]int32{
		"user":  0,
		"admin": 1,
	}
)

func (x GetResponseRole) Enum() *GetResponseRole {
	p := new(GetResponseRole)
	*p = x
	return p
}

func (x GetResponseRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetResponseRole) Descriptor() protoreflect.EnumDescriptor {
	return file_userApi_proto_enumTypes[1].Descriptor()
}

func (GetResponseRole) Type() protoreflect.EnumType {
	return &file_userApi_proto_enumTypes[1]
}

func (x GetResponseRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetResponseRole.Descriptor instead.
func (GetResponseRole) EnumDescriptor() ([]byte, []int) {
	return file_userApi_proto_rawDescGZIP(), []int{2, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email           string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password        string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	PasswordConfirm string `protobuf:"bytes,4,opt,name=password_confirm,json=passwordConfirm,proto3" json:"password_confirm,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userApi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_userApi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_userApi_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Request) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Request) GetPasswordConfirm() string {
	if x != nil {
		return x.PasswordConfirm
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userApi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userApi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_userApi_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email     string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userApi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userApi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_userApi_proto_rawDescGZIP(), []int{2}
}

func (x *GetResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetResponse) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetResponse) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email *wrappers.StringValue `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userApi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userApi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_userApi_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *UpdateRequest) GetEmail() *wrappers.StringValue {
	if x != nil {
		return x.Email
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userApi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userApi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_userApi_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_userApi_proto protoreflect.FileDescriptor

var file_userApi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x41, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x22, 0x1b, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x01,
	0x22, 0x20, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xda, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x1b, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x01, 0x22,
	0x85, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xe1, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x70, 0x69, 0x12, 0x33, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x4c, 0x52, 0x41, 0x53,
	0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x70, 0x69, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_userApi_proto_rawDescOnce sync.Once
	file_userApi_proto_rawDescData = file_userApi_proto_rawDesc
)

func file_userApi_proto_rawDescGZIP() []byte {
	file_userApi_proto_rawDescOnce.Do(func() {
		file_userApi_proto_rawDescData = protoimpl.X.CompressGZIP(file_userApi_proto_rawDescData)
	})
	return file_userApi_proto_rawDescData
}

var file_userApi_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_userApi_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_userApi_proto_goTypes = []any{
	(RequestRole)(0),             // 0: userApi.Request.role
	(GetResponseRole)(0),         // 1: userApi.GetResponse.role
	(*Request)(nil),              // 2: userApi.Request
	(*CreateResponse)(nil),       // 3: userApi.CreateResponse
	(*GetResponse)(nil),          // 4: userApi.GetResponse
	(*UpdateRequest)(nil),        // 5: userApi.UpdateRequest
	(*DeleteRequest)(nil),        // 6: userApi.DeleteRequest
	(*timestamp.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*wrappers.StringValue)(nil), // 8: google.protobuf.StringValue
	(*empty.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_userApi_proto_depIdxs = []int32{
	7, // 0: userApi.GetResponse.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: userApi.GetResponse.updated_at:type_name -> google.protobuf.Timestamp
	8, // 2: userApi.UpdateRequest.name:type_name -> google.protobuf.StringValue
	8, // 3: userApi.UpdateRequest.email:type_name -> google.protobuf.StringValue
	2, // 4: userApi.UserApi.Create:input_type -> userApi.Request
	2, // 5: userApi.UserApi.Get:input_type -> userApi.Request
	5, // 6: userApi.UserApi.Update:input_type -> userApi.UpdateRequest
	6, // 7: userApi.UserApi.Delete:input_type -> userApi.DeleteRequest
	3, // 8: userApi.UserApi.Create:output_type -> userApi.CreateResponse
	4, // 9: userApi.UserApi.Get:output_type -> userApi.GetResponse
	9, // 10: userApi.UserApi.Update:output_type -> google.protobuf.Empty
	9, // 11: userApi.UserApi.Delete:output_type -> google.protobuf.Empty
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_userApi_proto_init() }
func file_userApi_proto_init() {
	if File_userApi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userApi_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Request); i {
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
		file_userApi_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateResponse); i {
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
		file_userApi_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetResponse); i {
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
		file_userApi_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRequest); i {
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
		file_userApi_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRequest); i {
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
			RawDescriptor: file_userApi_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userApi_proto_goTypes,
		DependencyIndexes: file_userApi_proto_depIdxs,
		EnumInfos:         file_userApi_proto_enumTypes,
		MessageInfos:      file_userApi_proto_msgTypes,
	}.Build()
	File_userApi_proto = out.File
	file_userApi_proto_rawDesc = nil
	file_userApi_proto_goTypes = nil
	file_userApi_proto_depIdxs = nil
}