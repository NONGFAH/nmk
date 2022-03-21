// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: nmk.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	Type_EvMouseMove      Type = 0
	Type_EvMouseWheelMove Type = 1
	Type_EvMouseKey       Type = 2
	Type_EvKeyboardKey    Type = 3
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "EvMouseMove",
		1: "EvMouseWheelMove",
		2: "EvMouseKey",
		3: "EvKeyboardKey",
	}
	Type_value = map[string]int32{
		"EvMouseMove":      0,
		"EvMouseWheelMove": 1,
		"EvMouseKey":       2,
		"EvKeyboardKey":    3,
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
	return file_nmk_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_nmk_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_nmk_proto_rawDescGZIP(), []int{0}
}

type EventStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  Type  `protobuf:"varint,1,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	Code  int32 `protobuf:"zigzag32,2,opt,name=code,proto3" json:"code,omitempty"`
	Value int32 `protobuf:"zigzag32,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EventStreamRequest) Reset() {
	*x = EventStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nmk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStreamRequest) ProtoMessage() {}

func (x *EventStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nmk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStreamRequest.ProtoReflect.Descriptor instead.
func (*EventStreamRequest) Descriptor() ([]byte, []int) {
	return file_nmk_proto_rawDescGZIP(), []int{0}
}

func (x *EventStreamRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_EvMouseMove
}

func (x *EventStreamRequest) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *EventStreamRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_nmk_proto protoreflect.FileDescriptor

var file_nmk_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x6d, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2a, 0x50, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x45,
	0x76, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x45, 0x76, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x57, 0x68, 0x65, 0x65, 0x6c, 0x4d, 0x6f, 0x76, 0x65,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x76, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x4b, 0x65, 0x79,
	0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x76, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x4b, 0x65, 0x79, 0x10, 0x03, 0x32, 0x45, 0x0a, 0x07, 0x4e, 0x4d, 0x6b, 0x47, 0x72, 0x70, 0x63,
	0x12, 0x3a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x13, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x42, 0x1f, 0x50, 0x01,
	0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6e,
	0x67, 0x66, 0x61, 0x68, 0x2f, 0x6e, 0x6d, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nmk_proto_rawDescOnce sync.Once
	file_nmk_proto_rawDescData = file_nmk_proto_rawDesc
)

func file_nmk_proto_rawDescGZIP() []byte {
	file_nmk_proto_rawDescOnce.Do(func() {
		file_nmk_proto_rawDescData = protoimpl.X.CompressGZIP(file_nmk_proto_rawDescData)
	})
	return file_nmk_proto_rawDescData
}

var file_nmk_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nmk_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nmk_proto_goTypes = []interface{}{
	(Type)(0),                  // 0: Type
	(*EventStreamRequest)(nil), // 1: EventStreamRequest
	(*emptypb.Empty)(nil),      // 2: google.protobuf.Empty
}
var file_nmk_proto_depIdxs = []int32{
	0, // 0: EventStreamRequest.type:type_name -> Type
	1, // 1: NMkGrpc.control:input_type -> EventStreamRequest
	2, // 2: NMkGrpc.control:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nmk_proto_init() }
func file_nmk_proto_init() {
	if File_nmk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nmk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventStreamRequest); i {
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
			RawDescriptor: file_nmk_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nmk_proto_goTypes,
		DependencyIndexes: file_nmk_proto_depIdxs,
		EnumInfos:         file_nmk_proto_enumTypes,
		MessageInfos:      file_nmk_proto_msgTypes,
	}.Build()
	File_nmk_proto = out.File
	file_nmk_proto_rawDesc = nil
	file_nmk_proto_goTypes = nil
	file_nmk_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NMkGrpcClient is the client API for NMkGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NMkGrpcClient interface {
	Control(ctx context.Context, opts ...grpc.CallOption) (NMkGrpc_ControlClient, error)
}

type nMkGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewNMkGrpcClient(cc grpc.ClientConnInterface) NMkGrpcClient {
	return &nMkGrpcClient{cc}
}

func (c *nMkGrpcClient) Control(ctx context.Context, opts ...grpc.CallOption) (NMkGrpc_ControlClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NMkGrpc_serviceDesc.Streams[0], "/NMkGrpc/control", opts...)
	if err != nil {
		return nil, err
	}
	x := &nMkGrpcControlClient{stream}
	return x, nil
}

type NMkGrpc_ControlClient interface {
	Send(*EventStreamRequest) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type nMkGrpcControlClient struct {
	grpc.ClientStream
}

func (x *nMkGrpcControlClient) Send(m *EventStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nMkGrpcControlClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NMkGrpcServer is the server API for NMkGrpc service.
type NMkGrpcServer interface {
	Control(NMkGrpc_ControlServer) error
}

// UnimplementedNMkGrpcServer can be embedded to have forward compatible implementations.
type UnimplementedNMkGrpcServer struct {
}

func (*UnimplementedNMkGrpcServer) Control(NMkGrpc_ControlServer) error {
	return status.Errorf(codes.Unimplemented, "method Control not implemented")
}

func RegisterNMkGrpcServer(s *grpc.Server, srv NMkGrpcServer) {
	s.RegisterService(&_NMkGrpc_serviceDesc, srv)
}

func _NMkGrpc_Control_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NMkGrpcServer).Control(&nMkGrpcControlServer{stream})
}

type NMkGrpc_ControlServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*EventStreamRequest, error)
	grpc.ServerStream
}

type nMkGrpcControlServer struct {
	grpc.ServerStream
}

func (x *nMkGrpcControlServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nMkGrpcControlServer) Recv() (*EventStreamRequest, error) {
	m := new(EventStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _NMkGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NMkGrpc",
	HandlerType: (*NMkGrpcServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "control",
			Handler:       _NMkGrpc_Control_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "nmk.proto",
}
