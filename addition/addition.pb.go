// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: addition.proto

package addition

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

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number        int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	AnotherNumber int32 `protobuf:"varint,2,opt,name=anotherNumber,proto3" json:"anotherNumber,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_addition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_addition_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *AddRequest) GetAnotherNumber() int32 {
	if x != nil {
		return x.AnotherNumber
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_addition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_addition_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_addition_proto protoreflect.FileDescriptor

var file_addition_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x32, 0x40, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x14, 0x2e, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_addition_proto_rawDescOnce sync.Once
	file_addition_proto_rawDescData = file_addition_proto_rawDesc
)

func file_addition_proto_rawDescGZIP() []byte {
	file_addition_proto_rawDescOnce.Do(func() {
		file_addition_proto_rawDescData = protoimpl.X.CompressGZIP(file_addition_proto_rawDescData)
	})
	return file_addition_proto_rawDescData
}

var file_addition_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_addition_proto_goTypes = []interface{}{
	(*AddRequest)(nil),  // 0: addition.AddRequest
	(*AddResponse)(nil), // 1: addition.AddResponse
}
var file_addition_proto_depIdxs = []int32{
	0, // 0: addition.Addition.Add:input_type -> addition.AddRequest
	1, // 1: addition.Addition.Add:output_type -> addition.AddResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_addition_proto_init() }
func file_addition_proto_init() {
	if File_addition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_addition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_addition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
			RawDescriptor: file_addition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_addition_proto_goTypes,
		DependencyIndexes: file_addition_proto_depIdxs,
		MessageInfos:      file_addition_proto_msgTypes,
	}.Build()
	File_addition_proto = out.File
	file_addition_proto_rawDesc = nil
	file_addition_proto_goTypes = nil
	file_addition_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdditionClient is the client API for Addition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdditionClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type additionClient struct {
	cc grpc.ClientConnInterface
}

func NewAdditionClient(cc grpc.ClientConnInterface) AdditionClient {
	return &additionClient{cc}
}

func (c *additionClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/addition.Addition/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdditionServer is the server API for Addition service.
type AdditionServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
}

// UnimplementedAdditionServer can be embedded to have forward compatible implementations.
type UnimplementedAdditionServer struct {
}

func (*UnimplementedAdditionServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterAdditionServer(s *grpc.Server, srv AdditionServer) {
	s.RegisterService(&_Addition_serviceDesc, srv)
}

func _Addition_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdditionServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addition.Addition/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdditionServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Addition_serviceDesc = grpc.ServiceDesc{
	ServiceName: "addition.Addition",
	HandlerType: (*AdditionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Addition_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "addition.proto",
}
