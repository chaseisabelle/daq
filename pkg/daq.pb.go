// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: daq.proto

package daqpb

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

type EnqueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *EnqueueRequest) Reset() {
	*x = EnqueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueRequest) ProtoMessage() {}

func (x *EnqueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueRequest.ProtoReflect.Descriptor instead.
func (*EnqueueRequest) Descriptor() ([]byte, []int) {
	return file_daq_proto_rawDescGZIP(), []int{0}
}

func (x *EnqueueRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type EnqueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnqueueResponse) Reset() {
	*x = EnqueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueResponse) ProtoMessage() {}

func (x *EnqueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueResponse.ProtoReflect.Descriptor instead.
func (*EnqueueResponse) Descriptor() ([]byte, []int) {
	return file_daq_proto_rawDescGZIP(), []int{1}
}

type DequeueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DequeueRequest) Reset() {
	*x = DequeueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DequeueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DequeueRequest) ProtoMessage() {}

func (x *DequeueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DequeueRequest.ProtoReflect.Descriptor instead.
func (*DequeueRequest) Descriptor() ([]byte, []int) {
	return file_daq_proto_rawDescGZIP(), []int{2}
}

type DequeueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *DequeueResponse) Reset() {
	*x = DequeueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DequeueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DequeueResponse) ProtoMessage() {}

func (x *DequeueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DequeueResponse.ProtoReflect.Descriptor instead.
func (*DequeueResponse) Descriptor() ([]byte, []int) {
	return file_daq_proto_rawDescGZIP(), []int{3}
}

func (x *DequeueResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *DequeueResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type RequeueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *RequeueRequest) Reset() {
	*x = RequeueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daq_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequeueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequeueRequest) ProtoMessage() {}

func (x *RequeueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daq_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequeueRequest.ProtoReflect.Descriptor instead.
func (*RequeueRequest) Descriptor() ([]byte, []int) {
	return file_daq_proto_rawDescGZIP(), []int{4}
}

func (x *RequeueRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type RequeueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequeueResponse) Reset() {
	*x = RequeueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daq_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequeueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequeueResponse) ProtoMessage() {}

func (x *RequeueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daq_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequeueResponse.ProtoReflect.Descriptor instead.
func (*RequeueResponse) Descriptor() ([]byte, []int) {
	return file_daq_proto_rawDescGZIP(), []int{5}
}

var File_daq_proto protoreflect.FileDescriptor

var file_daq_proto_rawDesc = []byte{
	0x0a, 0x09, 0x64, 0x61, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x44, 0x41, 0x51,
	0x22, 0x24, 0x0a, 0x0e, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x0f, 0x44,
	0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x11, 0x0a, 0x0f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xab, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x45,
	0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x13, 0x2e, 0x44, 0x41, 0x51, 0x2e, 0x45, 0x6e, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x44, 0x41,
	0x51, 0x2e, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x07, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x13, 0x2e, 0x44,
	0x41, 0x51, 0x2e, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x44, 0x41, 0x51, 0x2e, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x12, 0x13, 0x2e, 0x44, 0x41, 0x51, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x44, 0x41, 0x51, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x64, 0x61, 0x71, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_daq_proto_rawDescOnce sync.Once
	file_daq_proto_rawDescData = file_daq_proto_rawDesc
)

func file_daq_proto_rawDescGZIP() []byte {
	file_daq_proto_rawDescOnce.Do(func() {
		file_daq_proto_rawDescData = protoimpl.X.CompressGZIP(file_daq_proto_rawDescData)
	})
	return file_daq_proto_rawDescData
}

var file_daq_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_daq_proto_goTypes = []interface{}{
	(*EnqueueRequest)(nil),  // 0: DAQ.EnqueueRequest
	(*EnqueueResponse)(nil), // 1: DAQ.EnqueueResponse
	(*DequeueRequest)(nil),  // 2: DAQ.DequeueRequest
	(*DequeueResponse)(nil), // 3: DAQ.DequeueResponse
	(*RequeueRequest)(nil),  // 4: DAQ.RequeueRequest
	(*RequeueResponse)(nil), // 5: DAQ.RequeueResponse
}
var file_daq_proto_depIdxs = []int32{
	0, // 0: DAQ.Service.Enqueue:input_type -> DAQ.EnqueueRequest
	2, // 1: DAQ.Service.Dequeue:input_type -> DAQ.DequeueRequest
	4, // 2: DAQ.Service.Requeue:input_type -> DAQ.RequeueRequest
	1, // 3: DAQ.Service.Enqueue:output_type -> DAQ.EnqueueResponse
	3, // 4: DAQ.Service.Dequeue:output_type -> DAQ.DequeueResponse
	5, // 5: DAQ.Service.Requeue:output_type -> DAQ.RequeueResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_daq_proto_init() }
func file_daq_proto_init() {
	if File_daq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_daq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueRequest); i {
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
		file_daq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueResponse); i {
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
		file_daq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DequeueRequest); i {
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
		file_daq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DequeueResponse); i {
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
		file_daq_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequeueRequest); i {
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
		file_daq_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequeueResponse); i {
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
			RawDescriptor: file_daq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_daq_proto_goTypes,
		DependencyIndexes: file_daq_proto_depIdxs,
		MessageInfos:      file_daq_proto_msgTypes,
	}.Build()
	File_daq_proto = out.File
	file_daq_proto_rawDesc = nil
	file_daq_proto_goTypes = nil
	file_daq_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Enqueue(ctx context.Context, in *EnqueueRequest, opts ...grpc.CallOption) (*EnqueueResponse, error)
	Dequeue(ctx context.Context, in *DequeueRequest, opts ...grpc.CallOption) (*DequeueResponse, error)
	Requeue(ctx context.Context, in *RequeueRequest, opts ...grpc.CallOption) (*RequeueResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Enqueue(ctx context.Context, in *EnqueueRequest, opts ...grpc.CallOption) (*EnqueueResponse, error) {
	out := new(EnqueueResponse)
	err := c.cc.Invoke(ctx, "/DAQ.Service/Enqueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Dequeue(ctx context.Context, in *DequeueRequest, opts ...grpc.CallOption) (*DequeueResponse, error) {
	out := new(DequeueResponse)
	err := c.cc.Invoke(ctx, "/DAQ.Service/Dequeue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Requeue(ctx context.Context, in *RequeueRequest, opts ...grpc.CallOption) (*RequeueResponse, error) {
	out := new(RequeueResponse)
	err := c.cc.Invoke(ctx, "/DAQ.Service/Requeue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Enqueue(context.Context, *EnqueueRequest) (*EnqueueResponse, error)
	Dequeue(context.Context, *DequeueRequest) (*DequeueResponse, error)
	Requeue(context.Context, *RequeueRequest) (*RequeueResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Enqueue(context.Context, *EnqueueRequest) (*EnqueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enqueue not implemented")
}
func (*UnimplementedServiceServer) Dequeue(context.Context, *DequeueRequest) (*DequeueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dequeue not implemented")
}
func (*UnimplementedServiceServer) Requeue(context.Context, *RequeueRequest) (*RequeueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Requeue not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DAQ.Service/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Enqueue(ctx, req.(*EnqueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Dequeue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DequeueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Dequeue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DAQ.Service/Dequeue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Dequeue(ctx, req.(*DequeueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Requeue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequeueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Requeue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DAQ.Service/Requeue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Requeue(ctx, req.(*RequeueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DAQ.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enqueue",
			Handler:    _Service_Enqueue_Handler,
		},
		{
			MethodName: "Dequeue",
			Handler:    _Service_Dequeue_Handler,
		},
		{
			MethodName: "Requeue",
			Handler:    _Service_Requeue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daq.proto",
}
