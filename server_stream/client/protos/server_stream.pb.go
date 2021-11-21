// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: server_stream.proto

package client

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start  int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Finish int64 `protobuf:"varint,2,opt,name=finish,proto3" json:"finish,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_server_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_server_stream_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Request) GetFinish() int64 {
	if x != nil {
		return x.Finish
	}
	return 0
}

var File_server_stream_proto protoreflect.FileDescriptor

var file_server_stream_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x37, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x32, 0x56, 0x0a, 0x0c, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x46, 0x0a, 0x0d, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x14, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x61, 0x62, 0x6f, 0x75, 0x72, 0x6f, 0x75, 0x6d, 0x69,
	0x6e, 0x65, 0x2f, 0x73, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x3b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_server_stream_proto_rawDescOnce sync.Once
	file_server_stream_proto_rawDescData = file_server_stream_proto_rawDesc
)

func file_server_stream_proto_rawDescGZIP() []byte {
	file_server_stream_proto_rawDescOnce.Do(func() {
		file_server_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_stream_proto_rawDescData)
	})
	return file_server_stream_proto_rawDescData
}

var file_server_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_server_stream_proto_goTypes = []interface{}{
	(*Request)(nil),               // 0: grpc_stream.Request
	(*wrapperspb.Int64Value)(nil), // 1: google.protobuf.Int64Value
}
var file_server_stream_proto_depIdxs = []int32{
	0, // 0: grpc_stream.ServerStream.startToFinish:input_type -> grpc_stream.Request
	1, // 1: grpc_stream.ServerStream.startToFinish:output_type -> google.protobuf.Int64Value
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_stream_proto_init() }
func file_server_stream_proto_init() {
	if File_server_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
			RawDescriptor: file_server_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_stream_proto_goTypes,
		DependencyIndexes: file_server_stream_proto_depIdxs,
		MessageInfos:      file_server_stream_proto_msgTypes,
	}.Build()
	File_server_stream_proto = out.File
	file_server_stream_proto_rawDesc = nil
	file_server_stream_proto_goTypes = nil
	file_server_stream_proto_depIdxs = nil
}
