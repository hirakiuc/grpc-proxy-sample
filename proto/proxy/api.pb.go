// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proxy/api.proto

package proxy

import (
	A "github.com/hirakiuc/grpc-proxy-sample/proto/A"
	B "github.com/hirakiuc/grpc-proxy-sample/proto/B"
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

type GreetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GreetingRequest) Reset() {
	*x = GreetingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingRequest) ProtoMessage() {}

func (x *GreetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingRequest.ProtoReflect.Descriptor instead.
func (*GreetingRequest) Descriptor() ([]byte, []int) {
	return file_proxy_api_proto_rawDescGZIP(), []int{0}
}

func (x *GreetingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GreetingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GreetingReply) Reset() {
	*x = GreetingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingReply) ProtoMessage() {}

func (x *GreetingReply) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingReply.ProtoReflect.Descriptor instead.
func (*GreetingReply) Descriptor() ([]byte, []int) {
	return file_proxy_api_proto_rawDescGZIP(), []int{1}
}

func (x *GreetingReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proxy_api_proto protoreflect.FileDescriptor

var file_proxy_api_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x0b, 0x41, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x42, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x3f, 0x0a, 0x03, 0x42, 0x66, 0x66, 0x12, 0x38, 0x0a, 0x08, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x46, 0x0a, 0x08, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x41, 0x12, 0x3a, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x17, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x5f, 0x61, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x40, 0x0a,
	0x08, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x42, 0x12, 0x34, 0x0a, 0x06, 0x53, 0x61, 0x79,
	0x42, 0x79, 0x65, 0x12, 0x15, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x62, 0x2e,
	0x42, 0x79, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x62, 0x2e, 0x42, 0x79, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69,
	0x72, 0x61, 0x6b, 0x69, 0x75, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_api_proto_rawDescOnce sync.Once
	file_proxy_api_proto_rawDescData = file_proxy_api_proto_rawDesc
)

func file_proxy_api_proto_rawDescGZIP() []byte {
	file_proxy_api_proto_rawDescOnce.Do(func() {
		file_proxy_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_api_proto_rawDescData)
	})
	return file_proxy_api_proto_rawDescData
}

var file_proxy_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proxy_api_proto_goTypes = []interface{}{
	(*GreetingRequest)(nil), // 0: proxy.GreetingRequest
	(*GreetingReply)(nil),   // 1: proxy.GreetingReply
	(*A.HelloRequest)(nil),  // 2: backend_a.HelloRequest
	(*B.ByeRequest)(nil),    // 3: backend_b.ByeRequest
	(*A.HelloReply)(nil),    // 4: backend_a.HelloReply
	(*B.ByeReply)(nil),      // 5: backend_b.ByeReply
}
var file_proxy_api_proto_depIdxs = []int32{
	0, // 0: proxy.Bff.Greeting:input_type -> proxy.GreetingRequest
	2, // 1: proxy.BackendA.SayHello:input_type -> backend_a.HelloRequest
	3, // 2: proxy.BackendB.SayBye:input_type -> backend_b.ByeRequest
	1, // 3: proxy.Bff.Greeting:output_type -> proxy.GreetingReply
	4, // 4: proxy.BackendA.SayHello:output_type -> backend_a.HelloReply
	5, // 5: proxy.BackendB.SayBye:output_type -> backend_b.ByeReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proxy_api_proto_init() }
func file_proxy_api_proto_init() {
	if File_proxy_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingRequest); i {
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
		file_proxy_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingReply); i {
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
			RawDescriptor: file_proxy_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_proxy_api_proto_goTypes,
		DependencyIndexes: file_proxy_api_proto_depIdxs,
		MessageInfos:      file_proxy_api_proto_msgTypes,
	}.Build()
	File_proxy_api_proto = out.File
	file_proxy_api_proto_rawDesc = nil
	file_proxy_api_proto_goTypes = nil
	file_proxy_api_proto_depIdxs = nil
}