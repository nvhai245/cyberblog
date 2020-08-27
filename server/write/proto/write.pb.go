// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: write.proto

package writeProto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NewUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	FirstName string `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Avatar    string `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Birthday  int64  `protobuf:"varint,7,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Bio       string `protobuf:"bytes,8,opt,name=bio,proto3" json:"bio,omitempty"`
	Facebook  string `protobuf:"bytes,9,opt,name=facebook,proto3" json:"facebook,omitempty"`
	Instagram string `protobuf:"bytes,10,opt,name=instagram,proto3" json:"instagram,omitempty"`
	Twitter   string `protobuf:"bytes,11,opt,name=twitter,proto3" json:"twitter,omitempty"`
	IsAdmin   bool   `protobuf:"varint,12,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	CreatedAt int64  `protobuf:"varint,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64  `protobuf:"varint,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *NewUser) Reset() {
	*x = NewUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_write_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUser) ProtoMessage() {}

func (x *NewUser) ProtoReflect() protoreflect.Message {
	mi := &file_write_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUser.ProtoReflect.Descriptor instead.
func (*NewUser) Descriptor() ([]byte, []int) {
	return file_write_proto_rawDescGZIP(), []int{0}
}

func (x *NewUser) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *NewUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NewUser) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *NewUser) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *NewUser) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *NewUser) GetBirthday() int64 {
	if x != nil {
		return x.Birthday
	}
	return 0
}

func (x *NewUser) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *NewUser) GetFacebook() string {
	if x != nil {
		return x.Facebook
	}
	return ""
}

func (x *NewUser) GetInstagram() string {
	if x != nil {
		return x.Instagram
	}
	return ""
}

func (x *NewUser) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *NewUser) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *NewUser) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *NewUser) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type SaveUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *NewUser `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Hash string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *SaveUserRequest) Reset() {
	*x = SaveUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_write_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveUserRequest) ProtoMessage() {}

func (x *SaveUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_write_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveUserRequest.ProtoReflect.Descriptor instead.
func (*SaveUserRequest) Descriptor() ([]byte, []int) {
	return file_write_proto_rawDescGZIP(), []int{1}
}

func (x *SaveUserRequest) GetUser() *NewUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SaveUserRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type SaveUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *NewUser `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Success bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SaveUserResponse) Reset() {
	*x = SaveUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_write_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveUserResponse) ProtoMessage() {}

func (x *SaveUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_write_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveUserResponse.ProtoReflect.Descriptor instead.
func (*SaveUserResponse) Descriptor() ([]byte, []int) {
	return file_write_proto_rawDescGZIP(), []int{2}
}

func (x *SaveUserResponse) GetUser() *NewUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SaveUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type EditUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestorEmail   string   `protobuf:"bytes,1,opt,name=requestor_email,json=requestorEmail,proto3" json:"requestor_email,omitempty"`
	RequestorIsAdmin bool     `protobuf:"varint,2,opt,name=requestor_is_admin,json=requestorIsAdmin,proto3" json:"requestor_is_admin,omitempty"`
	User             *NewUser `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *EditUserRequest) Reset() {
	*x = EditUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_write_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUserRequest) ProtoMessage() {}

func (x *EditUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_write_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditUserRequest.ProtoReflect.Descriptor instead.
func (*EditUserRequest) Descriptor() ([]byte, []int) {
	return file_write_proto_rawDescGZIP(), []int{3}
}

func (x *EditUserRequest) GetRequestorEmail() string {
	if x != nil {
		return x.RequestorEmail
	}
	return ""
}

func (x *EditUserRequest) GetRequestorIsAdmin() bool {
	if x != nil {
		return x.RequestorIsAdmin
	}
	return false
}

func (x *EditUserRequest) GetUser() *NewUser {
	if x != nil {
		return x.User
	}
	return nil
}

type EditUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	User    *NewUser `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *EditUserResponse) Reset() {
	*x = EditUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_write_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUserResponse) ProtoMessage() {}

func (x *EditUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_write_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditUserResponse.ProtoReflect.Descriptor instead.
func (*EditUserResponse) Descriptor() ([]byte, []int) {
	return file_write_proto_rawDescGZIP(), []int{4}
}

func (x *EditUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *EditUserResponse) GetUser() *NewUser {
	if x != nil {
		return x.User
	}
	return nil
}

var File_write_proto protoreflect.FileDescriptor

var file_write_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x02, 0x0a, 0x07, 0x4e, 0x65,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63,
	0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x63,
	0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72,
	0x61, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67,
	0x72, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4e, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x55, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x91, 0x01,
	0x0a, 0x0f, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x55, 0x0a, 0x10, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x27, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0x95, 0x01, 0x0a, 0x05, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b,
	0x2e, 0x77, 0x72, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x45, 0x64, 0x69,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x77, 0x72, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_write_proto_rawDescOnce sync.Once
	file_write_proto_rawDescData = file_write_proto_rawDesc
)

func file_write_proto_rawDescGZIP() []byte {
	file_write_proto_rawDescOnce.Do(func() {
		file_write_proto_rawDescData = protoimpl.X.CompressGZIP(file_write_proto_rawDescData)
	})
	return file_write_proto_rawDescData
}

var file_write_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_write_proto_goTypes = []interface{}{
	(*NewUser)(nil),          // 0: writeProto.NewUser
	(*SaveUserRequest)(nil),  // 1: writeProto.SaveUserRequest
	(*SaveUserResponse)(nil), // 2: writeProto.SaveUserResponse
	(*EditUserRequest)(nil),  // 3: writeProto.EditUserRequest
	(*EditUserResponse)(nil), // 4: writeProto.EditUserResponse
}
var file_write_proto_depIdxs = []int32{
	0, // 0: writeProto.SaveUserRequest.user:type_name -> writeProto.NewUser
	0, // 1: writeProto.SaveUserResponse.user:type_name -> writeProto.NewUser
	0, // 2: writeProto.EditUserRequest.user:type_name -> writeProto.NewUser
	0, // 3: writeProto.EditUserResponse.user:type_name -> writeProto.NewUser
	1, // 4: writeProto.Write.SaveUser:input_type -> writeProto.SaveUserRequest
	3, // 5: writeProto.Write.EditUser:input_type -> writeProto.EditUserRequest
	2, // 6: writeProto.Write.SaveUser:output_type -> writeProto.SaveUserResponse
	4, // 7: writeProto.Write.EditUser:output_type -> writeProto.EditUserResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_write_proto_init() }
func file_write_proto_init() {
	if File_write_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_write_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewUser); i {
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
		file_write_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveUserRequest); i {
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
		file_write_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveUserResponse); i {
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
		file_write_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditUserRequest); i {
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
		file_write_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditUserResponse); i {
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
			RawDescriptor: file_write_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_write_proto_goTypes,
		DependencyIndexes: file_write_proto_depIdxs,
		MessageInfos:      file_write_proto_msgTypes,
	}.Build()
	File_write_proto = out.File
	file_write_proto_rawDesc = nil
	file_write_proto_goTypes = nil
	file_write_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WriteClient is the client API for Write service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WriteClient interface {
	SaveUser(ctx context.Context, in *SaveUserRequest, opts ...grpc.CallOption) (*SaveUserResponse, error)
	EditUser(ctx context.Context, in *EditUserRequest, opts ...grpc.CallOption) (*EditUserResponse, error)
}

type writeClient struct {
	cc grpc.ClientConnInterface
}

func NewWriteClient(cc grpc.ClientConnInterface) WriteClient {
	return &writeClient{cc}
}

func (c *writeClient) SaveUser(ctx context.Context, in *SaveUserRequest, opts ...grpc.CallOption) (*SaveUserResponse, error) {
	out := new(SaveUserResponse)
	err := c.cc.Invoke(ctx, "/writeProto.Write/SaveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writeClient) EditUser(ctx context.Context, in *EditUserRequest, opts ...grpc.CallOption) (*EditUserResponse, error) {
	out := new(EditUserResponse)
	err := c.cc.Invoke(ctx, "/writeProto.Write/EditUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriteServer is the server API for Write service.
type WriteServer interface {
	SaveUser(context.Context, *SaveUserRequest) (*SaveUserResponse, error)
	EditUser(context.Context, *EditUserRequest) (*EditUserResponse, error)
}

// UnimplementedWriteServer can be embedded to have forward compatible implementations.
type UnimplementedWriteServer struct {
}

func (*UnimplementedWriteServer) SaveUser(context.Context, *SaveUserRequest) (*SaveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUser not implemented")
}
func (*UnimplementedWriteServer) EditUser(context.Context, *EditUserRequest) (*EditUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUser not implemented")
}

func RegisterWriteServer(s *grpc.Server, srv WriteServer) {
	s.RegisterService(&_Write_serviceDesc, srv)
}

func _Write_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriteServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writeProto.Write/SaveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriteServer).SaveUser(ctx, req.(*SaveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Write_EditUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriteServer).EditUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writeProto.Write/EditUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriteServer).EditUser(ctx, req.(*EditUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Write_serviceDesc = grpc.ServiceDesc{
	ServiceName: "writeProto.Write",
	HandlerType: (*WriteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUser",
			Handler:    _Write_SaveUser_Handler,
		},
		{
			MethodName: "EditUser",
			Handler:    _Write_EditUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "write.proto",
}
