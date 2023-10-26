// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: http.proto

package rpc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InfoRefsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base specifies the base read parameters
	Base *ReadRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// Service can be: upload-pack or receive-pack
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// Parameters to use with git -c (key=value pairs)
	GitConfigOptions []string `protobuf:"bytes,3,rep,name=git_config_options,json=gitConfigOptions,proto3" json:"git_config_options,omitempty"`
	// Git protocol version
	GitProtocol string `protobuf:"bytes,4,opt,name=git_protocol,json=gitProtocol,proto3" json:"git_protocol,omitempty"`
}

func (x *InfoRefsRequest) Reset() {
	*x = InfoRefsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRefsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRefsRequest) ProtoMessage() {}

func (x *InfoRefsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRefsRequest.ProtoReflect.Descriptor instead.
func (*InfoRefsRequest) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{0}
}

func (x *InfoRefsRequest) GetBase() *ReadRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *InfoRefsRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *InfoRefsRequest) GetGitConfigOptions() []string {
	if x != nil {
		return x.GitConfigOptions
	}
	return nil
}

func (x *InfoRefsRequest) GetGitProtocol() string {
	if x != nil {
		return x.GitProtocol
	}
	return ""
}

type InfoRefsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *InfoRefsResponse) Reset() {
	*x = InfoRefsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRefsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRefsResponse) ProtoMessage() {}

func (x *InfoRefsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRefsResponse.ProtoReflect.Descriptor instead.
func (*InfoRefsResponse) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{1}
}

func (x *InfoRefsResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ServicePackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base specifies the base parameters.
	// Depending on the service the matching base type has to be passed
	//
	// Types that are assignable to Base:
	//
	//	*ServicePackRequest_ReadBase
	//	*ServicePackRequest_WriteBase
	Base isServicePackRequest_Base `protobuf_oneof:"base"`
	// Service can be: upload-pack or receive-pack
	Service string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	// Raw data to be copied to stdin of 'git upload-pack'
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// Parameters to use with git -c (key=value pairs)
	GitConfigOptions []string `protobuf:"bytes,5,rep,name=git_config_options,json=gitConfigOptions,proto3" json:"git_config_options,omitempty"`
	// Git protocol version
	GitProtocol string `protobuf:"bytes,6,opt,name=git_protocol,json=gitProtocol,proto3" json:"git_protocol,omitempty"`
}

func (x *ServicePackRequest) Reset() {
	*x = ServicePackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePackRequest) ProtoMessage() {}

func (x *ServicePackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePackRequest.ProtoReflect.Descriptor instead.
func (*ServicePackRequest) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{2}
}

func (m *ServicePackRequest) GetBase() isServicePackRequest_Base {
	if m != nil {
		return m.Base
	}
	return nil
}

func (x *ServicePackRequest) GetReadBase() *ReadRequest {
	if x, ok := x.GetBase().(*ServicePackRequest_ReadBase); ok {
		return x.ReadBase
	}
	return nil
}

func (x *ServicePackRequest) GetWriteBase() *WriteRequest {
	if x, ok := x.GetBase().(*ServicePackRequest_WriteBase); ok {
		return x.WriteBase
	}
	return nil
}

func (x *ServicePackRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ServicePackRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ServicePackRequest) GetGitConfigOptions() []string {
	if x != nil {
		return x.GitConfigOptions
	}
	return nil
}

func (x *ServicePackRequest) GetGitProtocol() string {
	if x != nil {
		return x.GitProtocol
	}
	return ""
}

type isServicePackRequest_Base interface {
	isServicePackRequest_Base()
}

type ServicePackRequest_ReadBase struct {
	ReadBase *ReadRequest `protobuf:"bytes,1,opt,name=read_base,json=readBase,proto3,oneof"`
}

type ServicePackRequest_WriteBase struct {
	WriteBase *WriteRequest `protobuf:"bytes,2,opt,name=write_base,json=writeBase,proto3,oneof"`
}

func (*ServicePackRequest_ReadBase) isServicePackRequest_Base() {}

func (*ServicePackRequest_WriteBase) isServicePackRequest_Base() {}

type ServicePackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Raw data from stdout of 'git upload-pack'
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ServicePackResponse) Reset() {
	*x = ServicePackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePackResponse) ProtoMessage() {}

func (x *ServicePackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePackResponse.ProtoReflect.Descriptor instead.
func (*ServicePackResponse) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{3}
}

func (x *ServicePackResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_http_proto protoreflect.FileDescriptor

var file_http_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70,
	0x63, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa2, 0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x67, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x10, 0x67, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x26, 0x0a, 0x10, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x66, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x80, 0x02, 0x0a,
	0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64,
	0x42, 0x61, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x12, 0x67, 0x69, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x10, 0x67, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x69, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22,
	0x29, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x97, 0x01, 0x0a, 0x10, 0x53,
	0x6d, 0x61, 0x72, 0x74, 0x48, 0x54, 0x54, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x08, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x66, 0x73, 0x12, 0x14, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x66, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x6e, 0x65,
	0x73, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_http_proto_rawDescOnce sync.Once
	file_http_proto_rawDescData = file_http_proto_rawDesc
)

func file_http_proto_rawDescGZIP() []byte {
	file_http_proto_rawDescOnce.Do(func() {
		file_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_proto_rawDescData)
	})
	return file_http_proto_rawDescData
}

var file_http_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_http_proto_goTypes = []interface{}{
	(*InfoRefsRequest)(nil),     // 0: rpc.InfoRefsRequest
	(*InfoRefsResponse)(nil),    // 1: rpc.InfoRefsResponse
	(*ServicePackRequest)(nil),  // 2: rpc.ServicePackRequest
	(*ServicePackResponse)(nil), // 3: rpc.ServicePackResponse
	(*ReadRequest)(nil),         // 4: rpc.ReadRequest
	(*WriteRequest)(nil),        // 5: rpc.WriteRequest
}
var file_http_proto_depIdxs = []int32{
	4, // 0: rpc.InfoRefsRequest.base:type_name -> rpc.ReadRequest
	4, // 1: rpc.ServicePackRequest.read_base:type_name -> rpc.ReadRequest
	5, // 2: rpc.ServicePackRequest.write_base:type_name -> rpc.WriteRequest
	0, // 3: rpc.SmartHTTPService.InfoRefs:input_type -> rpc.InfoRefsRequest
	2, // 4: rpc.SmartHTTPService.ServicePack:input_type -> rpc.ServicePackRequest
	1, // 5: rpc.SmartHTTPService.InfoRefs:output_type -> rpc.InfoRefsResponse
	3, // 6: rpc.SmartHTTPService.ServicePack:output_type -> rpc.ServicePackResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_http_proto_init() }
func file_http_proto_init() {
	if File_http_proto != nil {
		return
	}
	file_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRefsRequest); i {
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
		file_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRefsResponse); i {
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
		file_http_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePackRequest); i {
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
		file_http_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePackResponse); i {
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
	file_http_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ServicePackRequest_ReadBase)(nil),
		(*ServicePackRequest_WriteBase)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_http_proto_goTypes,
		DependencyIndexes: file_http_proto_depIdxs,
		MessageInfos:      file_http_proto_msgTypes,
	}.Build()
	File_http_proto = out.File
	file_http_proto_rawDesc = nil
	file_http_proto_goTypes = nil
	file_http_proto_depIdxs = nil
}
