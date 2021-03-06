// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: pesertabkn_service.proto

package v1

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

type PesertaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nip string `protobuf:"bytes,1,opt,name=nip,proto3" json:"nip,omitempty"`
}

func (x *PesertaRequest) Reset() {
	*x = PesertaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pesertabkn_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PesertaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PesertaRequest) ProtoMessage() {}

func (x *PesertaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pesertabkn_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PesertaRequest.ProtoReflect.Descriptor instead.
func (*PesertaRequest) Descriptor() ([]byte, []int) {
	return file_pesertabkn_service_proto_rawDescGZIP(), []int{0}
}

func (x *PesertaRequest) GetNip() string {
	if x != nil {
		return x.Nip
	}
	return ""
}

type PesertaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status             string       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	MessageCode        string       `protobuf:"bytes,2,opt,name=message_code,json=messageCode,proto3" json:"message_code,omitempty"`
	MessageDescription string       `protobuf:"bytes,3,opt,name=message_description,json=messageDescription,proto3" json:"message_description,omitempty"`
	Data               *PesertaData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PesertaResponse) Reset() {
	*x = PesertaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pesertabkn_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PesertaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PesertaResponse) ProtoMessage() {}

func (x *PesertaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pesertabkn_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PesertaResponse.ProtoReflect.Descriptor instead.
func (*PesertaResponse) Descriptor() ([]byte, []int) {
	return file_pesertabkn_service_proto_rawDescGZIP(), []int{1}
}

func (x *PesertaResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PesertaResponse) GetMessageCode() string {
	if x != nil {
		return x.MessageCode
	}
	return ""
}

func (x *PesertaResponse) GetMessageDescription() string {
	if x != nil {
		return x.MessageDescription
	}
	return ""
}

func (x *PesertaResponse) GetData() *PesertaData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pesertabkn_service_proto protoreflect.FileDescriptor

var file_pesertabkn_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x65, 0x73, 0x65, 0x72, 0x74, 0x61, 0x62, 0x6b, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x65, 0x73, 0x65,
	0x72, 0x74, 0x61, 0x62, 0x6b, 0x6e, 0x1a, 0x18, 0x70, 0x65, 0x73, 0x65, 0x72, 0x74, 0x61, 0x62,
	0x6b, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x22, 0x0a, 0x0e, 0x50, 0x65, 0x73, 0x65, 0x72, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6e, 0x69, 0x70, 0x22, 0xaa, 0x01, 0x0a, 0x0f, 0x50, 0x65, 0x73, 0x65, 0x72, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x65, 0x73, 0x65, 0x72, 0x74, 0x61, 0x62, 0x6b, 0x6e, 0x2e,
	0x50, 0x65, 0x73, 0x65, 0x72, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0x55, 0x0a, 0x0a, 0x50, 0x65, 0x73, 0x65, 0x72, 0x74, 0x61, 0x42, 0x6b, 0x6e, 0x12,
	0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x65, 0x73, 0x65, 0x72, 0x74, 0x61, 0x12, 0x1a, 0x2e,
	0x70, 0x65, 0x73, 0x65, 0x72, 0x74, 0x61, 0x62, 0x6b, 0x6e, 0x2e, 0x50, 0x65, 0x73, 0x65, 0x72,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x65, 0x73, 0x65,
	0x72, 0x74, 0x61, 0x62, 0x6b, 0x6e, 0x2e, 0x50, 0x65, 0x73, 0x65, 0x72, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pesertabkn_service_proto_rawDescOnce sync.Once
	file_pesertabkn_service_proto_rawDescData = file_pesertabkn_service_proto_rawDesc
)

func file_pesertabkn_service_proto_rawDescGZIP() []byte {
	file_pesertabkn_service_proto_rawDescOnce.Do(func() {
		file_pesertabkn_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pesertabkn_service_proto_rawDescData)
	})
	return file_pesertabkn_service_proto_rawDescData
}

var file_pesertabkn_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pesertabkn_service_proto_goTypes = []interface{}{
	(*PesertaRequest)(nil),  // 0: pesertabkn.PesertaRequest
	(*PesertaResponse)(nil), // 1: pesertabkn.PesertaResponse
	(*PesertaData)(nil),     // 2: pesertabkn.PesertaData
}
var file_pesertabkn_service_proto_depIdxs = []int32{
	2, // 0: pesertabkn.PesertaResponse.data:type_name -> pesertabkn.PesertaData
	0, // 1: pesertabkn.PesertaBkn.GetPeserta:input_type -> pesertabkn.PesertaRequest
	1, // 2: pesertabkn.PesertaBkn.GetPeserta:output_type -> pesertabkn.PesertaResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pesertabkn_service_proto_init() }
func file_pesertabkn_service_proto_init() {
	if File_pesertabkn_service_proto != nil {
		return
	}
	file_pesertabkn_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pesertabkn_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PesertaRequest); i {
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
		file_pesertabkn_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PesertaResponse); i {
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
			RawDescriptor: file_pesertabkn_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pesertabkn_service_proto_goTypes,
		DependencyIndexes: file_pesertabkn_service_proto_depIdxs,
		MessageInfos:      file_pesertabkn_service_proto_msgTypes,
	}.Build()
	File_pesertabkn_service_proto = out.File
	file_pesertabkn_service_proto_rawDesc = nil
	file_pesertabkn_service_proto_goTypes = nil
	file_pesertabkn_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PesertaBknClient is the client API for PesertaBkn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PesertaBknClient interface {
	GetPeserta(ctx context.Context, in *PesertaRequest, opts ...grpc.CallOption) (*PesertaResponse, error)
}

type pesertaBknClient struct {
	cc grpc.ClientConnInterface
}

func NewPesertaBknClient(cc grpc.ClientConnInterface) PesertaBknClient {
	return &pesertaBknClient{cc}
}

func (c *pesertaBknClient) GetPeserta(ctx context.Context, in *PesertaRequest, opts ...grpc.CallOption) (*PesertaResponse, error) {
	out := new(PesertaResponse)
	err := c.cc.Invoke(ctx, "/pesertabkn.PesertaBkn/GetPeserta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PesertaBknServer is the server API for PesertaBkn service.
type PesertaBknServer interface {
	GetPeserta(context.Context, *PesertaRequest) (*PesertaResponse, error)
}

// UnimplementedPesertaBknServer can be embedded to have forward compatible implementations.
type UnimplementedPesertaBknServer struct {
}

func (*UnimplementedPesertaBknServer) GetPeserta(context.Context, *PesertaRequest) (*PesertaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeserta not implemented")
}

func RegisterPesertaBknServer(s *grpc.Server, srv PesertaBknServer) {
	s.RegisterService(&_PesertaBkn_serviceDesc, srv)
}

func _PesertaBkn_GetPeserta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PesertaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PesertaBknServer).GetPeserta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pesertabkn.PesertaBkn/GetPeserta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PesertaBknServer).GetPeserta(ctx, req.(*PesertaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PesertaBkn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pesertabkn.PesertaBkn",
	HandlerType: (*PesertaBknServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPeserta",
			Handler:    _PesertaBkn_GetPeserta_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pesertabkn_service.proto",
}
