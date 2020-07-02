// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: read.proto

package proto

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

type GetHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHashRequest) Reset() {
	*x = GetHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHashRequest) ProtoMessage() {}

func (x *GetHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_read_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHashRequest.ProtoReflect.Descriptor instead.
func (*GetHashRequest) Descriptor() ([]byte, []int) {
	return file_read_proto_rawDescGZIP(), []int{0}
}

func (x *GetHashRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetHashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *GetHashResponse) Reset() {
	*x = GetHashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHashResponse) ProtoMessage() {}

func (x *GetHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_read_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHashResponse.ProtoReflect.Descriptor instead.
func (*GetHashResponse) Descriptor() ([]byte, []int) {
	return file_read_proto_rawDescGZIP(), []int{1}
}

func (x *GetHashResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_read_proto protoreflect.FileDescriptor

var file_read_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x32, 0x40, 0x0a, 0x04,
	0x52, 0x65, 0x61, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_read_proto_rawDescOnce sync.Once
	file_read_proto_rawDescData = file_read_proto_rawDesc
)

func file_read_proto_rawDescGZIP() []byte {
	file_read_proto_rawDescOnce.Do(func() {
		file_read_proto_rawDescData = protoimpl.X.CompressGZIP(file_read_proto_rawDescData)
	})
	return file_read_proto_rawDescData
}

var file_read_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_read_proto_goTypes = []interface{}{
	(*GetHashRequest)(nil),  // 0: proto.GetHashRequest
	(*GetHashResponse)(nil), // 1: proto.GetHashResponse
}
var file_read_proto_depIdxs = []int32{
	0, // 0: proto.Read.GetHash:input_type -> proto.GetHashRequest
	1, // 1: proto.Read.GetHash:output_type -> proto.GetHashResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_read_proto_init() }
func file_read_proto_init() {
	if File_read_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_read_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHashRequest); i {
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
		file_read_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHashResponse); i {
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
			RawDescriptor: file_read_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_read_proto_goTypes,
		DependencyIndexes: file_read_proto_depIdxs,
		MessageInfos:      file_read_proto_msgTypes,
	}.Build()
	File_read_proto = out.File
	file_read_proto_rawDesc = nil
	file_read_proto_goTypes = nil
	file_read_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReadClient is the client API for Read service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReadClient interface {
	GetHash(ctx context.Context, in *GetHashRequest, opts ...grpc.CallOption) (*GetHashResponse, error)
}

type readClient struct {
	cc grpc.ClientConnInterface
}

func NewReadClient(cc grpc.ClientConnInterface) ReadClient {
	return &readClient{cc}
}

func (c *readClient) GetHash(ctx context.Context, in *GetHashRequest, opts ...grpc.CallOption) (*GetHashResponse, error) {
	out := new(GetHashResponse)
	err := c.cc.Invoke(ctx, "/proto.Read/GetHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReadServer is the server API for Read service.
type ReadServer interface {
	GetHash(context.Context, *GetHashRequest) (*GetHashResponse, error)
}

// UnimplementedReadServer can be embedded to have forward compatible implementations.
type UnimplementedReadServer struct {
}

func (*UnimplementedReadServer) GetHash(context.Context, *GetHashRequest) (*GetHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHash not implemented")
}

func RegisterReadServer(s *grpc.Server, srv ReadServer) {
	s.RegisterService(&_Read_serviceDesc, srv)
}

func _Read_GetHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadServer).GetHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Read/GetHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadServer).GetHash(ctx, req.(*GetHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Read_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Read",
	HandlerType: (*ReadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHash",
			Handler:    _Read_GetHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "read.proto",
}
