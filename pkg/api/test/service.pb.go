// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: api/test/service.proto

package test

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestGetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestGetRequest) Reset() {
	*x = TestGetRequest{}
	mi := &file_api_test_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestGetRequest) ProtoMessage() {}

func (x *TestGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestGetRequest.ProtoReflect.Descriptor instead.
func (*TestGetRequest) Descriptor() ([]byte, []int) {
	return file_api_test_service_proto_rawDescGZIP(), []int{0}
}

type TestGetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ans           string                 `protobuf:"bytes,1,opt,name=ans,proto3" json:"ans,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestGetResponse) Reset() {
	*x = TestGetResponse{}
	mi := &file_api_test_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestGetResponse) ProtoMessage() {}

func (x *TestGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestGetResponse.ProtoReflect.Descriptor instead.
func (*TestGetResponse) Descriptor() ([]byte, []int) {
	return file_api_test_service_proto_rawDescGZIP(), []int{1}
}

func (x *TestGetResponse) GetAns() string {
	if x != nil {
		return x.Ans
	}
	return ""
}

var File_api_test_service_proto protoreflect.FileDescriptor

var file_api_test_service_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e,
	0x54, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x23,
	0x0a, 0x0f, 0x54, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x6e, 0x73, 0x32, 0x54, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x45, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x65, 0x72, 0x65, 0x76, 0x65, 0x72, 0x7a,
	0x65, 0x76, 0x49, 0x76, 0x61, 0x6e, 0x2f, 0x72, 0x61, 0x7a, 0x72, 0x61, 0x62, 0x6f, 0x74, 0x6b,
	0x61, 0x2d, 0x62, 0x69, 0x7a, 0x6e, 0x65, 0x73, 0x2d, 0x70, 0x72, 0x69, 0x6c, 0x6f, 0x6a, 0x65,
	0x6e, 0x69, 0x69, 0x2d, 0x6b, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x72, 0x73, 0x6b, 0x61, 0x79,
	0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_api_test_service_proto_rawDescOnce sync.Once
	file_api_test_service_proto_rawDescData []byte
)

func file_api_test_service_proto_rawDescGZIP() []byte {
	file_api_test_service_proto_rawDescOnce.Do(func() {
		file_api_test_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_test_service_proto_rawDesc), len(file_api_test_service_proto_rawDesc)))
	})
	return file_api_test_service_proto_rawDescData
}

var file_api_test_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_test_service_proto_goTypes = []any{
	(*TestGetRequest)(nil),  // 0: test.TestGetRequest
	(*TestGetResponse)(nil), // 1: test.TestGetResponse
}
var file_api_test_service_proto_depIdxs = []int32{
	0, // 0: test.TestService.Get:input_type -> test.TestGetRequest
	1, // 1: test.TestService.Get:output_type -> test.TestGetResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_test_service_proto_init() }
func file_api_test_service_proto_init() {
	if File_api_test_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_test_service_proto_rawDesc), len(file_api_test_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_test_service_proto_goTypes,
		DependencyIndexes: file_api_test_service_proto_depIdxs,
		MessageInfos:      file_api_test_service_proto_msgTypes,
	}.Build()
	File_api_test_service_proto = out.File
	file_api_test_service_proto_goTypes = nil
	file_api_test_service_proto_depIdxs = nil
}
