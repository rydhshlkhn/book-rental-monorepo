// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.14.0
// source: token/token.proto

package token

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_token_token_proto_msgTypes[0]
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
	return file_token_token_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_token_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_token_token_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Validate token payload
type PayloadValidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *PayloadValidate) Reset() {
	*x = PayloadValidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadValidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadValidate) ProtoMessage() {}

func (x *PayloadValidate) ProtoReflect() protoreflect.Message {
	mi := &file_token_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadValidate.ProtoReflect.Descriptor instead.
func (*PayloadValidate) Descriptor() ([]byte, []int) {
	return file_token_token_proto_rawDescGZIP(), []int{2}
}

func (x *PayloadValidate) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ResponseValidation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool                          `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Claim   *ResponseValidation_ClaimData `protobuf:"bytes,2,opt,name=Claim,proto3" json:"Claim,omitempty"`
}

func (x *ResponseValidation) Reset() {
	*x = ResponseValidation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseValidation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseValidation) ProtoMessage() {}

func (x *ResponseValidation) ProtoReflect() protoreflect.Message {
	mi := &file_token_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseValidation.ProtoReflect.Descriptor instead.
func (*ResponseValidation) Descriptor() ([]byte, []int) {
	return file_token_token_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseValidation) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ResponseValidation) GetClaim() *ResponseValidation_ClaimData {
	if x != nil {
		return x.Claim
	}
	return nil
}

type ResponseGenerate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool                    `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Data    *ResponseGenerate_Token `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ResponseGenerate) Reset() {
	*x = ResponseGenerate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGenerate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGenerate) ProtoMessage() {}

func (x *ResponseGenerate) ProtoReflect() protoreflect.Message {
	mi := &file_token_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGenerate.ProtoReflect.Descriptor instead.
func (*ResponseGenerate) Descriptor() ([]byte, []int) {
	return file_token_token_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseGenerate) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ResponseGenerate) GetData() *ResponseGenerate_Token {
	if x != nil {
		return x.Data
	}
	return nil
}

type ClaimData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Audience  string    `protobuf:"bytes,1,opt,name=Audience,proto3" json:"Audience,omitempty"`
	ExpiresAt int64     `protobuf:"varint,2,opt,name=ExpiresAt,proto3" json:"ExpiresAt,omitempty"`
	IssuedAt  int64     `protobuf:"varint,3,opt,name=IssuedAt,proto3" json:"IssuedAt,omitempty"`
	Issuer    string    `protobuf:"bytes,4,opt,name=Issuer,proto3" json:"Issuer,omitempty"`
	NotBefore int64     `protobuf:"varint,5,opt,name=NotBefore,proto3" json:"NotBefore,omitempty"`
	Subject   string    `protobuf:"bytes,6,opt,name=Subject,proto3" json:"Subject,omitempty"`
	DeviceID  string    `protobuf:"bytes,7,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	User      *UserData `protobuf:"bytes,8,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *ClaimData) Reset() {
	*x = ClaimData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimData) ProtoMessage() {}

func (x *ClaimData) ProtoReflect() protoreflect.Message {
	mi := &file_token_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimData.ProtoReflect.Descriptor instead.
func (*ClaimData) Descriptor() ([]byte, []int) {
	return file_token_token_proto_rawDescGZIP(), []int{5}
}

func (x *ClaimData) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

func (x *ClaimData) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *ClaimData) GetIssuedAt() int64 {
	if x != nil {
		return x.IssuedAt
	}
	return 0
}

func (x *ClaimData) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *ClaimData) GetNotBefore() int64 {
	if x != nil {
		return x.NotBefore
	}
	return 0
}

func (x *ClaimData) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *ClaimData) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *ClaimData) GetUser() *UserData {
	if x != nil {
		return x.User
	}
	return nil
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	RoleID   string `protobuf:"bytes,3,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	DeviceID string `protobuf:"bytes,4,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_token_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_token_token_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_token_token_proto_rawDescGZIP(), []int{6}
}

func (x *UserData) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UserData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserData) GetRoleID() string {
	if x != nil {
		return x.RoleID
	}
	return ""
}

func (x *UserData) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

type ResponseValidation_ClaimData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Audience  string    `protobuf:"bytes,1,opt,name=Audience,proto3" json:"Audience,omitempty"`
	ExpiresAt int64     `protobuf:"varint,2,opt,name=ExpiresAt,proto3" json:"ExpiresAt,omitempty"`
	IssuedAt  int64     `protobuf:"varint,3,opt,name=IssuedAt,proto3" json:"IssuedAt,omitempty"`
	Issuer    string    `protobuf:"bytes,4,opt,name=Issuer,proto3" json:"Issuer,omitempty"`
	NotBefore int64     `protobuf:"varint,5,opt,name=NotBefore,proto3" json:"NotBefore,omitempty"`
	Subject   string    `protobuf:"bytes,6,opt,name=Subject,proto3" json:"Subject,omitempty"`
	DeviceID  string    `protobuf:"bytes,7,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	User      *UserData `protobuf:"bytes,8,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *ResponseValidation_ClaimData) Reset() {
	*x = ResponseValidation_ClaimData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_token_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseValidation_ClaimData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseValidation_ClaimData) ProtoMessage() {}

func (x *ResponseValidation_ClaimData) ProtoReflect() protoreflect.Message {
	mi := &file_token_token_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseValidation_ClaimData.ProtoReflect.Descriptor instead.
func (*ResponseValidation_ClaimData) Descriptor() ([]byte, []int) {
	return file_token_token_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ResponseValidation_ClaimData) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

func (x *ResponseValidation_ClaimData) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *ResponseValidation_ClaimData) GetIssuedAt() int64 {
	if x != nil {
		return x.IssuedAt
	}
	return 0
}

func (x *ResponseValidation_ClaimData) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *ResponseValidation_ClaimData) GetNotBefore() int64 {
	if x != nil {
		return x.NotBefore
	}
	return 0
}

func (x *ResponseValidation_ClaimData) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *ResponseValidation_ClaimData) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *ResponseValidation_ClaimData) GetUser() *UserData {
	if x != nil {
		return x.User
	}
	return nil
}

type ResponseGenerate_Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string     `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	RefreshToken string     `protobuf:"bytes,2,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	Claim        *ClaimData `protobuf:"bytes,3,opt,name=Claim,proto3" json:"Claim,omitempty"`
}

func (x *ResponseGenerate_Token) Reset() {
	*x = ResponseGenerate_Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_token_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGenerate_Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGenerate_Token) ProtoMessage() {}

func (x *ResponseGenerate_Token) ProtoReflect() protoreflect.Message {
	mi := &file_token_token_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGenerate_Token.ProtoReflect.Descriptor instead.
func (*ResponseGenerate_Token) Descriptor() ([]byte, []int) {
	return file_token_token_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ResponseGenerate_Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ResponseGenerate_Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *ResponseGenerate_Token) GetClaim() *ClaimData {
	if x != nil {
		return x.Claim
	}
	return nil
}

var File_token_token_proto protoreflect.FileDescriptor

var file_token_token_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x23, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xde,
	0x02, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x39, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x1a, 0xf2, 0x01, 0x0a, 0x09, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x42, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4e, 0x6f, 0x74, 0x42, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22,
	0xca, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x31,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x69, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x22, 0xf2, 0x01, 0x0a,
	0x09, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x42,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4e, 0x6f, 0x74,
	0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x6a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6c,
	0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x32, 0x8d, 0x01,
	0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x42,
	0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x16, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x19, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x0f, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x17, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x42, 0x30, 0x5a,
	0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_token_proto_rawDescOnce sync.Once
	file_token_token_proto_rawDescData = file_token_token_proto_rawDesc
)

func file_token_token_proto_rawDescGZIP() []byte {
	file_token_token_proto_rawDescOnce.Do(func() {
		file_token_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_token_proto_rawDescData)
	})
	return file_token_token_proto_rawDescData
}

var file_token_token_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_token_token_proto_goTypes = []interface{}{
	(*Request)(nil),                      // 0: token.Request
	(*Response)(nil),                     // 1: token.Response
	(*PayloadValidate)(nil),              // 2: token.PayloadValidate
	(*ResponseValidation)(nil),           // 3: token.ResponseValidation
	(*ResponseGenerate)(nil),             // 4: token.ResponseGenerate
	(*ClaimData)(nil),                    // 5: token.ClaimData
	(*UserData)(nil),                     // 6: token.UserData
	(*ResponseValidation_ClaimData)(nil), // 7: token.ResponseValidation.ClaimData
	(*ResponseGenerate_Token)(nil),       // 8: token.ResponseGenerate.Token
}
var file_token_token_proto_depIdxs = []int32{
	7, // 0: token.ResponseValidation.Claim:type_name -> token.ResponseValidation.ClaimData
	8, // 1: token.ResponseGenerate.Data:type_name -> token.ResponseGenerate.Token
	6, // 2: token.ClaimData.User:type_name -> token.UserData
	6, // 3: token.ResponseValidation.ClaimData.User:type_name -> token.UserData
	5, // 4: token.ResponseGenerate.Token.Claim:type_name -> token.ClaimData
	2, // 5: token.TokenHandler.ValidateToken:input_type -> token.PayloadValidate
	6, // 6: token.TokenHandler.GenerateToken:input_type -> token.UserData
	3, // 7: token.TokenHandler.ValidateToken:output_type -> token.ResponseValidation
	4, // 8: token.TokenHandler.GenerateToken:output_type -> token.ResponseGenerate
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_token_token_proto_init() }
func file_token_token_proto_init() {
	if File_token_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_token_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_token_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadValidate); i {
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
		file_token_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseValidation); i {
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
		file_token_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGenerate); i {
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
		file_token_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimData); i {
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
		file_token_token_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_token_token_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseValidation_ClaimData); i {
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
		file_token_token_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGenerate_Token); i {
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
			RawDescriptor: file_token_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_token_proto_goTypes,
		DependencyIndexes: file_token_token_proto_depIdxs,
		MessageInfos:      file_token_token_proto_msgTypes,
	}.Build()
	File_token_token_proto = out.File
	file_token_token_proto_rawDesc = nil
	file_token_token_proto_goTypes = nil
	file_token_token_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TokenHandlerClient is the client API for TokenHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenHandlerClient interface {
	ValidateToken(ctx context.Context, in *PayloadValidate, opts ...grpc.CallOption) (*ResponseValidation, error)
	GenerateToken(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*ResponseGenerate, error)
}

type tokenHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenHandlerClient(cc grpc.ClientConnInterface) TokenHandlerClient {
	return &tokenHandlerClient{cc}
}

func (c *tokenHandlerClient) ValidateToken(ctx context.Context, in *PayloadValidate, opts ...grpc.CallOption) (*ResponseValidation, error) {
	out := new(ResponseValidation)
	err := c.cc.Invoke(ctx, "/token.TokenHandler/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenHandlerClient) GenerateToken(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*ResponseGenerate, error) {
	out := new(ResponseGenerate)
	err := c.cc.Invoke(ctx, "/token.TokenHandler/GenerateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenHandlerServer is the server API for TokenHandler service.
type TokenHandlerServer interface {
	ValidateToken(context.Context, *PayloadValidate) (*ResponseValidation, error)
	GenerateToken(context.Context, *UserData) (*ResponseGenerate, error)
}

// UnimplementedTokenHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedTokenHandlerServer struct {
}

func (*UnimplementedTokenHandlerServer) ValidateToken(context.Context, *PayloadValidate) (*ResponseValidation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (*UnimplementedTokenHandlerServer) GenerateToken(context.Context, *UserData) (*ResponseGenerate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}

func RegisterTokenHandlerServer(s *grpc.Server, srv TokenHandlerServer) {
	s.RegisterService(&_TokenHandler_serviceDesc, srv)
}

func _TokenHandler_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadValidate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenHandlerServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.TokenHandler/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenHandlerServer).ValidateToken(ctx, req.(*PayloadValidate))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenHandler_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenHandlerServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.TokenHandler/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenHandlerServer).GenerateToken(ctx, req.(*UserData))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "token.TokenHandler",
	HandlerType: (*TokenHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateToken",
			Handler:    _TokenHandler_ValidateToken_Handler,
		},
		{
			MethodName: "GenerateToken",
			Handler:    _TokenHandler_GenerateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token/token.proto",
}
