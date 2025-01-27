// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: github.com/lastbackend/toolkit/examples/service/apis/ptypes/messages.proto

package typespb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type HelloWorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Data map[string]string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HelloWorldRequest) Reset() {
	*x = HelloWorldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldRequest) ProtoMessage() {}

func (x *HelloWorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldRequest.ProtoReflect.Descriptor instead.
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDescGZIP(), []int{0}
}

func (x *HelloWorldRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloWorldRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *HelloWorldRequest) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type HelloWorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type      string            `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Data      map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt int64             `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64             `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *HelloWorldResponse) Reset() {
	*x = HelloWorldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldResponse) ProtoMessage() {}

func (x *HelloWorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldResponse.ProtoReflect.Descriptor instead.
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDescGZIP(), []int{1}
}

func (x *HelloWorldResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HelloWorldResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloWorldResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *HelloWorldResponse) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *HelloWorldResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *HelloWorldResponse) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto protoreflect.FileDescriptor

var file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x73,
	0x74, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6c, 0x61,
	0x73, 0x74, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x72, 0x03, 0x18, 0x80, 0x08, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x4e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6c, 0x61, 0x73, 0x74,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x9a, 0x01, 0x02, 0x08, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x8a, 0x02, 0x0a, 0x12, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x6c, 0x61, 0x73, 0x74, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x61, 0x73, 0x74, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x74, 0x6f, 0x6f, 0x6c,
	0x6b, 0x69, 0x74, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDescOnce sync.Once
	file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDescData = file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDesc
)

func file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDescGZIP() []byte {
	file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDescOnce.Do(func() {
		file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDescData)
	})
	return file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDescData
}

var file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_goTypes = []interface{}{
	(*HelloWorldRequest)(nil),  // 0: lastbackend.example.HelloWorldRequest
	(*HelloWorldResponse)(nil), // 1: lastbackend.example.HelloWorldResponse
	nil,                        // 2: lastbackend.example.HelloWorldRequest.DataEntry
	nil,                        // 3: lastbackend.example.HelloWorldResponse.DataEntry
}
var file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_depIdxs = []int32{
	2, // 0: lastbackend.example.HelloWorldRequest.data:type_name -> lastbackend.example.HelloWorldRequest.DataEntry
	3, // 1: lastbackend.example.HelloWorldResponse.data:type_name -> lastbackend.example.HelloWorldResponse.DataEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_init() }
func file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_init() {
	if File_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldRequest); i {
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
		file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldResponse); i {
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
			RawDescriptor: file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_goTypes,
		DependencyIndexes: file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_depIdxs,
		MessageInfos:      file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_msgTypes,
	}.Build()
	File_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto = out.File
	file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_rawDesc = nil
	file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_goTypes = nil
	file_github_com_lastbackend_toolkit_examples_service_apis_ptypes_messages_proto_depIdxs = nil
}
