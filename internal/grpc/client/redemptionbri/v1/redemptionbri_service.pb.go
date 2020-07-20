// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: redemptionbri_service.proto

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

//Request message for service Redemption
type RedemptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*RedemptionData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *RedemptionRequest) Reset() {
	*x = RedemptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redemptionbri_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedemptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedemptionRequest) ProtoMessage() {}

func (x *RedemptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redemptionbri_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedemptionRequest.ProtoReflect.Descriptor instead.
func (*RedemptionRequest) Descriptor() ([]byte, []int) {
	return file_redemptionbri_service_proto_rawDescGZIP(), []int{0}
}

func (x *RedemptionRequest) GetData() []*RedemptionData {
	if x != nil {
		return x.Data
	}
	return nil
}

//Response message for service Redemption
type RedemptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//Status from API Redemption BRI per reference no
	//Format : Success/Failed/Error
	//Success : response from BRI --> http 200 and status true
	//Failed : response from BRI --> http 200 and status false
	//Error : response from BRI --> other than http 200
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	//Response Code from API Redemption BRI per reference no
	MessageCode string `protobuf:"bytes,2,opt,name=message_code,json=messageCode,proto3" json:"message_code,omitempty"`
	//Response Description from API Redemption BRI per reference no
	MessageDescription string `protobuf:"bytes,3,opt,name=message_description,json=messageDescription,proto3" json:"message_description,omitempty"`
	//Data per per reference no
	Data *Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RedemptionResponse) Reset() {
	*x = RedemptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redemptionbri_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedemptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedemptionResponse) ProtoMessage() {}

func (x *RedemptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redemptionbri_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedemptionResponse.ProtoReflect.Descriptor instead.
func (*RedemptionResponse) Descriptor() ([]byte, []int) {
	return file_redemptionbri_service_proto_rawDescGZIP(), []int{1}
}

func (x *RedemptionResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RedemptionResponse) GetMessageCode() string {
	if x != nil {
		return x.MessageCode
	}
	return ""
}

func (x *RedemptionResponse) GetMessageDescription() string {
	if x != nil {
		return x.MessageDescription
	}
	return ""
}

func (x *RedemptionResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_redemptionbri_service_proto protoreflect.FileDescriptor

var file_redemptionbri_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x72, 0x69, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72,
	0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x72, 0x69, 0x1a, 0x1b, 0x72, 0x65,
	0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x72, 0x69, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x11, 0x52, 0x65, 0x64,
	0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72,
	0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x72, 0x69, 0x2e, 0x52, 0x65, 0x64,
	0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0xa9, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x72, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x66, 0x0a,
	0x0d, 0x52, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x72, 0x69, 0x12, 0x55,
	0x0a, 0x0a, 0x52, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x72,
	0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x72, 0x69, 0x2e, 0x52, 0x65, 0x64,
	0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x72, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x72, 0x69, 0x2e, 0x52,
	0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x04, 0x5a, 0x02, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_redemptionbri_service_proto_rawDescOnce sync.Once
	file_redemptionbri_service_proto_rawDescData = file_redemptionbri_service_proto_rawDesc
)

func file_redemptionbri_service_proto_rawDescGZIP() []byte {
	file_redemptionbri_service_proto_rawDescOnce.Do(func() {
		file_redemptionbri_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_redemptionbri_service_proto_rawDescData)
	})
	return file_redemptionbri_service_proto_rawDescData
}

var file_redemptionbri_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_redemptionbri_service_proto_goTypes = []interface{}{
	(*RedemptionRequest)(nil),  // 0: redemptionbri.RedemptionRequest
	(*RedemptionResponse)(nil), // 1: redemptionbri.RedemptionResponse
	(*RedemptionData)(nil),     // 2: redemptionbri.RedemptionData
	(*Data)(nil),               // 3: redemptionbri.Data
}
var file_redemptionbri_service_proto_depIdxs = []int32{
	2, // 0: redemptionbri.RedemptionRequest.data:type_name -> redemptionbri.RedemptionData
	3, // 1: redemptionbri.RedemptionResponse.data:type_name -> redemptionbri.Data
	0, // 2: redemptionbri.RedemptionBri.Redemption:input_type -> redemptionbri.RedemptionRequest
	1, // 3: redemptionbri.RedemptionBri.Redemption:output_type -> redemptionbri.RedemptionResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_redemptionbri_service_proto_init() }
func file_redemptionbri_service_proto_init() {
	if File_redemptionbri_service_proto != nil {
		return
	}
	file_redemptionbri_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_redemptionbri_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedemptionRequest); i {
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
		file_redemptionbri_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedemptionResponse); i {
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
			RawDescriptor: file_redemptionbri_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_redemptionbri_service_proto_goTypes,
		DependencyIndexes: file_redemptionbri_service_proto_depIdxs,
		MessageInfos:      file_redemptionbri_service_proto_msgTypes,
	}.Build()
	File_redemptionbri_service_proto = out.File
	file_redemptionbri_service_proto_rawDesc = nil
	file_redemptionbri_service_proto_goTypes = nil
	file_redemptionbri_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RedemptionBriClient is the client API for RedemptionBri service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RedemptionBriClient interface {
	Redemption(ctx context.Context, in *RedemptionRequest, opts ...grpc.CallOption) (RedemptionBri_RedemptionClient, error)
}

type redemptionBriClient struct {
	cc grpc.ClientConnInterface
}

func NewRedemptionBriClient(cc grpc.ClientConnInterface) RedemptionBriClient {
	return &redemptionBriClient{cc}
}

func (c *redemptionBriClient) Redemption(ctx context.Context, in *RedemptionRequest, opts ...grpc.CallOption) (RedemptionBri_RedemptionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RedemptionBri_serviceDesc.Streams[0], "/redemptionbri.RedemptionBri/Redemption", opts...)
	if err != nil {
		return nil, err
	}
	x := &redemptionBriRedemptionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RedemptionBri_RedemptionClient interface {
	Recv() (*RedemptionResponse, error)
	grpc.ClientStream
}

type redemptionBriRedemptionClient struct {
	grpc.ClientStream
}

func (x *redemptionBriRedemptionClient) Recv() (*RedemptionResponse, error) {
	m := new(RedemptionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RedemptionBriServer is the server API for RedemptionBri service.
type RedemptionBriServer interface {
	Redemption(*RedemptionRequest, RedemptionBri_RedemptionServer) error
}

// UnimplementedRedemptionBriServer can be embedded to have forward compatible implementations.
type UnimplementedRedemptionBriServer struct {
}

func (*UnimplementedRedemptionBriServer) Redemption(*RedemptionRequest, RedemptionBri_RedemptionServer) error {
	return status.Errorf(codes.Unimplemented, "method Redemption not implemented")
}

func RegisterRedemptionBriServer(s *grpc.Server, srv RedemptionBriServer) {
	s.RegisterService(&_RedemptionBri_serviceDesc, srv)
}

func _RedemptionBri_Redemption_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RedemptionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RedemptionBriServer).Redemption(m, &redemptionBriRedemptionServer{stream})
}

type RedemptionBri_RedemptionServer interface {
	Send(*RedemptionResponse) error
	grpc.ServerStream
}

type redemptionBriRedemptionServer struct {
	grpc.ServerStream
}

func (x *redemptionBriRedemptionServer) Send(m *RedemptionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _RedemptionBri_serviceDesc = grpc.ServiceDesc{
	ServiceName: "redemptionbri.RedemptionBri",
	HandlerType: (*RedemptionBriServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Redemption",
			Handler:       _RedemptionBri_Redemption_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "redemptionbri_service.proto",
}
