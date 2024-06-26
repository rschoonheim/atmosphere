// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: api/proto/tls.proto

package environment

import (
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

type TlsConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertFile string `protobuf:"bytes,1,opt,name=cert_file,json=certFile,proto3" json:"cert_file,omitempty"`
	KeyFile  string `protobuf:"bytes,2,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty"`
}

func (x *TlsConfiguration) Reset() {
	*x = TlsConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_tls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TlsConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TlsConfiguration) ProtoMessage() {}

func (x *TlsConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_tls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TlsConfiguration.ProtoReflect.Descriptor instead.
func (*TlsConfiguration) Descriptor() ([]byte, []int) {
	return file_api_proto_tls_proto_rawDescGZIP(), []int{0}
}

func (x *TlsConfiguration) GetCertFile() string {
	if x != nil {
		return x.CertFile
	}
	return ""
}

func (x *TlsConfiguration) GetKeyFile() string {
	if x != nil {
		return x.KeyFile
	}
	return ""
}

var File_api_proto_tls_proto protoreflect.FileDescriptor

var file_api_proto_tls_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x61, 0x74, 0x6d, 0x6f, 0x73, 0x70, 0x68, 0x65, 0x72,
	0x65, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x22, 0x4a, 0x0a, 0x10, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x2a, 0x5a, 0x28,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x3b, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_tls_proto_rawDescOnce sync.Once
	file_api_proto_tls_proto_rawDescData = file_api_proto_tls_proto_rawDesc
)

func file_api_proto_tls_proto_rawDescGZIP() []byte {
	file_api_proto_tls_proto_rawDescOnce.Do(func() {
		file_api_proto_tls_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_tls_proto_rawDescData)
	})
	return file_api_proto_tls_proto_rawDescData
}

var file_api_proto_tls_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_proto_tls_proto_goTypes = []interface{}{
	(*TlsConfiguration)(nil), // 0: atmosphere.environment.v1.TlsConfiguration
}
var file_api_proto_tls_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_tls_proto_init() }
func file_api_proto_tls_proto_init() {
	if File_api_proto_tls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_tls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TlsConfiguration); i {
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
			RawDescriptor: file_api_proto_tls_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_tls_proto_goTypes,
		DependencyIndexes: file_api_proto_tls_proto_depIdxs,
		MessageInfos:      file_api_proto_tls_proto_msgTypes,
	}.Build()
	File_api_proto_tls_proto = out.File
	file_api_proto_tls_proto_rawDesc = nil
	file_api_proto_tls_proto_goTypes = nil
	file_api_proto_tls_proto_depIdxs = nil
}
