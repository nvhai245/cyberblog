// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: cyber.proto

package cyberProto

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

type User struct {
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

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyber_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_cyber_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_cyber_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetBirthday() int64 {
	if x != nil {
		return x.Birthday
	}
	return 0
}

func (x *User) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *User) GetFacebook() string {
	if x != nil {
		return x.Facebook
	}
	return ""
}

func (x *User) GetInstagram() string {
	if x != nil {
		return x.Instagram
	}
	return ""
}

func (x *User) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *User) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *User) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *User) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type GetUserByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestorId int32 `protobuf:"varint,1,opt,name=requestor_id,json=requestorId,proto3" json:"requestor_id,omitempty"`
	UserId      int32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserByIdRequest) Reset() {
	*x = GetUserByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyber_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdRequest) ProtoMessage() {}

func (x *GetUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cyber_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_cyber_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserByIdRequest) GetRequestorId() int32 {
	if x != nil {
		return x.RequestorId
	}
	return 0
}

func (x *GetUserByIdRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	User    *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByIdResponse) Reset() {
	*x = GetUserByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyber_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdResponse) ProtoMessage() {}

func (x *GetUserByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cyber_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdResponse.ProtoReflect.Descriptor instead.
func (*GetUserByIdResponse) Descriptor() ([]byte, []int) {
	return file_cyber_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByIdResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetUserByIdResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetAllUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId int32 `protobuf:"varint,1,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
}

func (x *GetAllUsersRequest) Reset() {
	*x = GetAllUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyber_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsersRequest) ProtoMessage() {}

func (x *GetAllUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cyber_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsersRequest.ProtoReflect.Descriptor instead.
func (*GetAllUsersRequest) Descriptor() ([]byte, []int) {
	return file_cyber_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllUsersRequest) GetAdminId() int32 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

type GetAllUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Users   []*User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetAllUsersResponse) Reset() {
	*x = GetAllUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyber_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsersResponse) ProtoMessage() {}

func (x *GetAllUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cyber_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsersResponse.ProtoReflect.Descriptor instead.
func (*GetAllUsersResponse) Descriptor() ([]byte, []int) {
	return file_cyber_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllUsersResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetAllUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type EditUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestorEmail   string `protobuf:"bytes,1,opt,name=requestor_email,json=requestorEmail,proto3" json:"requestor_email,omitempty"`
	RequestorIsAdmin bool   `protobuf:"varint,2,opt,name=requestor_is_admin,json=requestorIsAdmin,proto3" json:"requestor_is_admin,omitempty"`
	User             *User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *EditUserRequest) Reset() {
	*x = EditUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyber_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUserRequest) ProtoMessage() {}

func (x *EditUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cyber_proto_msgTypes[5]
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
	return file_cyber_proto_rawDescGZIP(), []int{5}
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

func (x *EditUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type EditUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	User    *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *EditUserResponse) Reset() {
	*x = EditUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyber_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUserResponse) ProtoMessage() {}

func (x *EditUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cyber_proto_msgTypes[6]
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
	return file_cyber_proto_rawDescGZIP(), []int{6}
}

func (x *EditUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *EditUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestorEmail   string `protobuf:"bytes,1,opt,name=requestor_email,json=requestorEmail,proto3" json:"requestor_email,omitempty"`
	RequestorIsAdmin bool   `protobuf:"varint,2,opt,name=requestor_is_admin,json=requestorIsAdmin,proto3" json:"requestor_is_admin,omitempty"`
	UserId           int32  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyber_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cyber_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_cyber_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUserRequest) GetRequestorEmail() string {
	if x != nil {
		return x.RequestorEmail
	}
	return ""
}

func (x *DeleteUserRequest) GetRequestorIsAdmin() bool {
	if x != nil {
		return x.RequestorIsAdmin
	}
	return false
}

func (x *DeleteUserRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	User    *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyber_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cyber_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_cyber_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_cyber_proto protoreflect.FileDescriptor

var file_cyber_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x79, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63,
	0x79, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x02, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f,
	0x6f, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f,
	0x6f, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x50, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x57, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x26,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x79, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x0f, 0x45, 0x64, 0x69, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x5f, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x52, 0x0a, 0x10, 0x45, 0x64, 0x69, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x83, 0x01, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x54, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xbb, 0x02, 0x0a, 0x05, 0x43, 0x79, 0x62, 0x65,
	0x72, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x1e, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x1e, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e,
	0x63, 0x79, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x79, 0x62,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x63, 0x79, 0x62, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cyber_proto_rawDescOnce sync.Once
	file_cyber_proto_rawDescData = file_cyber_proto_rawDesc
)

func file_cyber_proto_rawDescGZIP() []byte {
	file_cyber_proto_rawDescOnce.Do(func() {
		file_cyber_proto_rawDescData = protoimpl.X.CompressGZIP(file_cyber_proto_rawDescData)
	})
	return file_cyber_proto_rawDescData
}

var file_cyber_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cyber_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: cyberProto.User
	(*GetUserByIdRequest)(nil),  // 1: cyberProto.GetUserByIdRequest
	(*GetUserByIdResponse)(nil), // 2: cyberProto.GetUserByIdResponse
	(*GetAllUsersRequest)(nil),  // 3: cyberProto.GetAllUsersRequest
	(*GetAllUsersResponse)(nil), // 4: cyberProto.GetAllUsersResponse
	(*EditUserRequest)(nil),     // 5: cyberProto.EditUserRequest
	(*EditUserResponse)(nil),    // 6: cyberProto.EditUserResponse
	(*DeleteUserRequest)(nil),   // 7: cyberProto.DeleteUserRequest
	(*DeleteUserResponse)(nil),  // 8: cyberProto.DeleteUserResponse
}
var file_cyber_proto_depIdxs = []int32{
	0, // 0: cyberProto.GetUserByIdResponse.user:type_name -> cyberProto.User
	0, // 1: cyberProto.GetAllUsersResponse.users:type_name -> cyberProto.User
	0, // 2: cyberProto.EditUserRequest.user:type_name -> cyberProto.User
	0, // 3: cyberProto.EditUserResponse.user:type_name -> cyberProto.User
	0, // 4: cyberProto.DeleteUserResponse.user:type_name -> cyberProto.User
	1, // 5: cyberProto.Cyber.GetUserById:input_type -> cyberProto.GetUserByIdRequest
	3, // 6: cyberProto.Cyber.GetAllUsers:input_type -> cyberProto.GetAllUsersRequest
	5, // 7: cyberProto.Cyber.EditUser:input_type -> cyberProto.EditUserRequest
	7, // 8: cyberProto.Cyber.DeleteUser:input_type -> cyberProto.DeleteUserRequest
	2, // 9: cyberProto.Cyber.GetUserById:output_type -> cyberProto.GetUserByIdResponse
	4, // 10: cyberProto.Cyber.GetAllUsers:output_type -> cyberProto.GetAllUsersResponse
	6, // 11: cyberProto.Cyber.EditUser:output_type -> cyberProto.EditUserResponse
	8, // 12: cyberProto.Cyber.DeleteUser:output_type -> cyberProto.DeleteUserResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cyber_proto_init() }
func file_cyber_proto_init() {
	if File_cyber_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cyber_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_cyber_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIdRequest); i {
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
		file_cyber_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIdResponse); i {
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
		file_cyber_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUsersRequest); i {
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
		file_cyber_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUsersResponse); i {
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
		file_cyber_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_cyber_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_cyber_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
		file_cyber_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserResponse); i {
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
			RawDescriptor: file_cyber_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cyber_proto_goTypes,
		DependencyIndexes: file_cyber_proto_depIdxs,
		MessageInfos:      file_cyber_proto_msgTypes,
	}.Build()
	File_cyber_proto = out.File
	file_cyber_proto_rawDesc = nil
	file_cyber_proto_goTypes = nil
	file_cyber_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CyberClient is the client API for Cyber service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CyberClient interface {
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error)
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
	EditUser(ctx context.Context, in *EditUserRequest, opts ...grpc.CallOption) (*EditUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type cyberClient struct {
	cc grpc.ClientConnInterface
}

func NewCyberClient(cc grpc.ClientConnInterface) CyberClient {
	return &cyberClient{cc}
}

func (c *cyberClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error) {
	out := new(GetUserByIdResponse)
	err := c.cc.Invoke(ctx, "/cyberProto.Cyber/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cyberClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	out := new(GetAllUsersResponse)
	err := c.cc.Invoke(ctx, "/cyberProto.Cyber/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cyberClient) EditUser(ctx context.Context, in *EditUserRequest, opts ...grpc.CallOption) (*EditUserResponse, error) {
	out := new(EditUserResponse)
	err := c.cc.Invoke(ctx, "/cyberProto.Cyber/EditUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cyberClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/cyberProto.Cyber/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CyberServer is the server API for Cyber service.
type CyberServer interface {
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
	EditUser(context.Context, *EditUserRequest) (*EditUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
}

// UnimplementedCyberServer can be embedded to have forward compatible implementations.
type UnimplementedCyberServer struct {
}

func (*UnimplementedCyberServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (*UnimplementedCyberServer) GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (*UnimplementedCyberServer) EditUser(context.Context, *EditUserRequest) (*EditUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUser not implemented")
}
func (*UnimplementedCyberServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterCyberServer(s *grpc.Server, srv CyberServer) {
	s.RegisterService(&_Cyber_serviceDesc, srv)
}

func _Cyber_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CyberServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyberProto.Cyber/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CyberServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cyber_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CyberServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyberProto.Cyber/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CyberServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cyber_EditUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CyberServer).EditUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyberProto.Cyber/EditUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CyberServer).EditUser(ctx, req.(*EditUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cyber_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CyberServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyberProto.Cyber/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CyberServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cyber_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cyberProto.Cyber",
	HandlerType: (*CyberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _Cyber_GetUserById_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _Cyber_GetAllUsers_Handler,
		},
		{
			MethodName: "EditUser",
			Handler:    _Cyber_EditUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Cyber_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cyber.proto",
}
