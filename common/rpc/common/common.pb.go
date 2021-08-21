// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: common.proto

package __

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

type BaseAppReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beid int64 `protobuf:"varint,1,opt,name=beid,proto3" json:"beid,omitempty"`
}

func (x *BaseAppReq) Reset() {
	*x = BaseAppReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseAppReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseAppReq) ProtoMessage() {}

func (x *BaseAppReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseAppReq.ProtoReflect.Descriptor instead.
func (*BaseAppReq) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *BaseAppReq) GetBeid() int64 {
	if x != nil {
		return x.Beid
	}
	return 0
}

type BaseAppResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beid        int64  `protobuf:"varint,1,opt,name=beid,proto3" json:"beid,omitempty"`
	Logo        string `protobuf:"bytes,2,opt,name=logo,proto3" json:"logo,omitempty"`
	Sname       string `protobuf:"bytes,3,opt,name=sname,proto3" json:"sname,omitempty"`
	Isclose     int64  `protobuf:"varint,4,opt,name=isclose,proto3" json:"isclose,omitempty"`
	Fullwebsite string `protobuf:"bytes,5,opt,name=fullwebsite,proto3" json:"fullwebsite,omitempty"`
}

func (x *BaseAppResp) Reset() {
	*x = BaseAppResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseAppResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseAppResp) ProtoMessage() {}

func (x *BaseAppResp) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseAppResp.ProtoReflect.Descriptor instead.
func (*BaseAppResp) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *BaseAppResp) GetBeid() int64 {
	if x != nil {
		return x.Beid
	}
	return 0
}

func (x *BaseAppResp) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *BaseAppResp) GetSname() string {
	if x != nil {
		return x.Sname
	}
	return ""
}

func (x *BaseAppResp) GetIsclose() int64 {
	if x != nil {
		return x.Isclose
	}
	return 0
}

func (x *BaseAppResp) GetFullwebsite() string {
	if x != nil {
		return x.Fullwebsite
	}
	return ""
}

//请求的api
type AppConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beid  int64 `protobuf:"varint,1,opt,name=beid,proto3" json:"beid,omitempty"`
	Ptyid int64 `protobuf:"varint,2,opt,name=ptyid,proto3" json:"ptyid,omitempty"`
}

func (x *AppConfigReq) Reset() {
	*x = AppConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConfigReq) ProtoMessage() {}

func (x *AppConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConfigReq.ProtoReflect.Descriptor instead.
func (*AppConfigReq) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *AppConfigReq) GetBeid() int64 {
	if x != nil {
		return x.Beid
	}
	return 0
}

func (x *AppConfigReq) GetPtyid() int64 {
	if x != nil {
		return x.Ptyid
	}
	return 0
}

//返回的值
type AppConfigResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Beid      int64  `protobuf:"varint,2,opt,name=beid,proto3" json:"beid,omitempty"`
	Ptyid     int64  `protobuf:"varint,3,opt,name=ptyid,proto3" json:"ptyid,omitempty"`
	Appid     string `protobuf:"bytes,4,opt,name=appid,proto3" json:"appid,omitempty"`
	Appsecret string `protobuf:"bytes,5,opt,name=appsecret,proto3" json:"appsecret,omitempty"`
	Title     string `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *AppConfigResp) Reset() {
	*x = AppConfigResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConfigResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConfigResp) ProtoMessage() {}

func (x *AppConfigResp) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConfigResp.ProtoReflect.Descriptor instead.
func (*AppConfigResp) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *AppConfigResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppConfigResp) GetBeid() int64 {
	if x != nil {
		return x.Beid
	}
	return 0
}

func (x *AppConfigResp) GetPtyid() int64 {
	if x != nil {
		return x.Ptyid
	}
	return 0
}

func (x *AppConfigResp) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *AppConfigResp) GetAppsecret() string {
	if x != nil {
		return x.Appsecret
	}
	return ""
}

func (x *AppConfigResp) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x0a, 0x42, 0x61, 0x73, 0x65, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x62, 0x65, 0x69, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x73,
	0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x65, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62, 0x65, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x66, 0x75, 0x6c, 0x6c, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x75, 0x6c, 0x6c, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x62, 0x65, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x74, 0x79, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x74, 0x79, 0x69, 0x64, 0x22, 0x93, 0x01, 0x0a,
	0x0d, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x65, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62, 0x65,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x74, 0x79, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x74, 0x79, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x32, 0x7c, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x73, 0x65, 0x41, 0x70, 0x70, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_proto_goTypes = []interface{}{
	(*BaseAppReq)(nil),    // 0: common.BaseAppReq
	(*BaseAppResp)(nil),   // 1: common.BaseAppResp
	(*AppConfigReq)(nil),  // 2: common.AppConfigReq
	(*AppConfigResp)(nil), // 3: common.AppConfigResp
}
var file_common_proto_depIdxs = []int32{
	2, // 0: common.Common.GetAppConfig:input_type -> common.AppConfigReq
	0, // 1: common.Common.GetBaseApp:input_type -> common.BaseAppReq
	3, // 2: common.Common.GetAppConfig:output_type -> common.AppConfigResp
	1, // 3: common.Common.GetBaseApp:output_type -> common.BaseAppResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseAppReq); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseAppResp); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConfigReq); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConfigResp); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommonClient is the client API for Common service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommonClient interface {
	GetAppConfig(ctx context.Context, in *AppConfigReq, opts ...grpc.CallOption) (*AppConfigResp, error)
	GetBaseApp(ctx context.Context, in *BaseAppReq, opts ...grpc.CallOption) (*BaseAppResp, error)
}

type commonClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonClient(cc grpc.ClientConnInterface) CommonClient {
	return &commonClient{cc}
}

func (c *commonClient) GetAppConfig(ctx context.Context, in *AppConfigReq, opts ...grpc.CallOption) (*AppConfigResp, error) {
	out := new(AppConfigResp)
	err := c.cc.Invoke(ctx, "/common.Common/GetAppConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) GetBaseApp(ctx context.Context, in *BaseAppReq, opts ...grpc.CallOption) (*BaseAppResp, error) {
	out := new(BaseAppResp)
	err := c.cc.Invoke(ctx, "/common.Common/GetBaseApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServer is the server API for Common service.
type CommonServer interface {
	GetAppConfig(context.Context, *AppConfigReq) (*AppConfigResp, error)
	GetBaseApp(context.Context, *BaseAppReq) (*BaseAppResp, error)
}

// UnimplementedCommonServer can be embedded to have forward compatible implementations.
type UnimplementedCommonServer struct {
}

func (*UnimplementedCommonServer) GetAppConfig(context.Context, *AppConfigReq) (*AppConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppConfig not implemented")
}
func (*UnimplementedCommonServer) GetBaseApp(context.Context, *BaseAppReq) (*BaseAppResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBaseApp not implemented")
}

func RegisterCommonServer(s *grpc.Server, srv CommonServer) {
	s.RegisterService(&_Common_serviceDesc, srv)
}

func _Common_GetAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).GetAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.Common/GetAppConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).GetAppConfig(ctx, req.(*AppConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_GetBaseApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).GetBaseApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.Common/GetBaseApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).GetBaseApp(ctx, req.(*BaseAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Common_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.Common",
	HandlerType: (*CommonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppConfig",
			Handler:    _Common_GetAppConfig_Handler,
		},
		{
			MethodName: "GetBaseApp",
			Handler:    _Common_GetBaseApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common.proto",
}
