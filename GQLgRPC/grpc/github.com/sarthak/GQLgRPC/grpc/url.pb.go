// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: proto/url.proto

package grpc

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

type Urlpath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *Urlpath) Reset() {
	*x = Urlpath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Urlpath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Urlpath) ProtoMessage() {}

func (x *Urlpath) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Urlpath.ProtoReflect.Descriptor instead.
func (*Urlpath) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{0}
}

func (x *Urlpath) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type Disecturlpath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme            string `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Topleveldomain    string `protobuf:"bytes,2,opt,name=topleveldomain,proto3" json:"topleveldomain,omitempty"`
	Secondleveldomain string `protobuf:"bytes,3,opt,name=secondleveldomain,proto3" json:"secondleveldomain,omitempty"`
}

func (x *Disecturlpath) Reset() {
	*x = Disecturlpath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disecturlpath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disecturlpath) ProtoMessage() {}

func (x *Disecturlpath) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disecturlpath.ProtoReflect.Descriptor instead.
func (*Disecturlpath) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{1}
}

func (x *Disecturlpath) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Disecturlpath) GetTopleveldomain() string {
	if x != nil {
		return x.Topleveldomain
	}
	return ""
}

func (x *Disecturlpath) GetSecondleveldomain() string {
	if x != nil {
		return x.Secondleveldomain
	}
	return ""
}

var File_proto_url_proto protoreflect.FileDescriptor

var file_proto_url_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x1b, 0x0a, 0x07, 0x75, 0x72, 0x6c, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x52, 0x4c, 0x22, 0x7d, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x6c, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x74, 0x6f, 0x70, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x6f, 0x70, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x32, 0x38, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x65, 0x63, 0x74, 0x75, 0x72, 0x6c,
	0x12, 0x2b, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x75,
	0x72, 0x6c, 0x70, 0x61, 0x74, 0x68, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x69,
	0x73, 0x65, 0x63, 0x74, 0x75, 0x72, 0x6c, 0x70, 0x61, 0x74, 0x68, 0x22, 0x00, 0x42, 0x21, 0x5a,
	0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x72, 0x74,
	0x68, 0x61, 0x6b, 0x2f, 0x47, 0x51, 0x4c, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_url_proto_rawDescOnce sync.Once
	file_proto_url_proto_rawDescData = file_proto_url_proto_rawDesc
)

func file_proto_url_proto_rawDescGZIP() []byte {
	file_proto_url_proto_rawDescOnce.Do(func() {
		file_proto_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_url_proto_rawDescData)
	})
	return file_proto_url_proto_rawDescData
}

var file_proto_url_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_url_proto_goTypes = []interface{}{
	(*Urlpath)(nil),       // 0: grpc.urlpath
	(*Disecturlpath)(nil), // 1: grpc.disecturlpath
}
var file_proto_url_proto_depIdxs = []int32{
	0, // 0: grpc.disecturl.url:input_type -> grpc.urlpath
	1, // 1: grpc.disecturl.url:output_type -> grpc.disecturlpath
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_url_proto_init() }
func file_proto_url_proto_init() {
	if File_proto_url_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Urlpath); i {
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
		file_proto_url_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Disecturlpath); i {
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
			RawDescriptor: file_proto_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_url_proto_goTypes,
		DependencyIndexes: file_proto_url_proto_depIdxs,
		MessageInfos:      file_proto_url_proto_msgTypes,
	}.Build()
	File_proto_url_proto = out.File
	file_proto_url_proto_rawDesc = nil
	file_proto_url_proto_goTypes = nil
	file_proto_url_proto_depIdxs = nil
}
