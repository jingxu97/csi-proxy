// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: github.com/kubernetes-csi/csi-proxy/integrationtests/csiapigen/new_group/api/v1alpha1/api.proto

package v1alpha1

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

type FooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubMessage *FooRequestSubMessage `protobuf:"bytes,1,opt,name=subMessage,proto3" json:"subMessage,omitempty"`
	Blob       []byte                `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (x *FooRequest) Reset() {
	*x = FooRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooRequest) ProtoMessage() {}

func (x *FooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooRequest.ProtoReflect.Descriptor instead.
func (*FooRequest) Descriptor() ([]byte, []int) {
	return file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDescGZIP(), []int{0}
}

func (x *FooRequest) GetSubMessage() *FooRequestSubMessage {
	if x != nil {
		return x.SubMessage
	}
	return nil
}

func (x *FooRequest) GetBlob() []byte {
	if x != nil {
		return x.Blob
	}
	return nil
}

type FooRequestSubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blah          int32                      `protobuf:"varint,1,opt,name=blah,proto3" json:"blah,omitempty"`
	SubSubMessage []*FooRequestSubSubMessage `protobuf:"bytes,2,rep,name=subSubMessage,proto3" json:"subSubMessage,omitempty"`
}

func (x *FooRequestSubMessage) Reset() {
	*x = FooRequestSubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooRequestSubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooRequestSubMessage) ProtoMessage() {}

func (x *FooRequestSubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooRequestSubMessage.ProtoReflect.Descriptor instead.
func (*FooRequestSubMessage) Descriptor() ([]byte, []int) {
	return file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDescGZIP(), []int{1}
}

func (x *FooRequestSubMessage) GetBlah() int32 {
	if x != nil {
		return x.Blah
	}
	return 0
}

func (x *FooRequestSubMessage) GetSubSubMessage() []*FooRequestSubSubMessage {
	if x != nil {
		return x.SubSubMessage
	}
	return nil
}

type FooRequestSubSubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bools []bool `protobuf:"varint,1,rep,packed,name=bools,proto3" json:"bools,omitempty"`
}

func (x *FooRequestSubSubMessage) Reset() {
	*x = FooRequestSubSubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooRequestSubSubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooRequestSubSubMessage) ProtoMessage() {}

func (x *FooRequestSubSubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooRequestSubSubMessage.ProtoReflect.Descriptor instead.
func (*FooRequestSubSubMessage) Descriptor() ([]byte, []int) {
	return file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDescGZIP(), []int{2}
}

func (x *FooRequestSubSubMessage) GetBools() []bool {
	if x != nil {
		return x.Bools
	}
	return nil
}

type FooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response32 int32                 `protobuf:"varint,1,opt,name=response32,proto3" json:"response32,omitempty"`
	SubMessage *FooRequestSubMessage `protobuf:"bytes,2,opt,name=subMessage,proto3" json:"subMessage,omitempty"`
}

func (x *FooResponse) Reset() {
	*x = FooResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooResponse) ProtoMessage() {}

func (x *FooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooResponse.ProtoReflect.Descriptor instead.
func (*FooResponse) Descriptor() ([]byte, []int) {
	return file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDescGZIP(), []int{3}
}

func (x *FooResponse) GetResponse32() int32 {
	if x != nil {
		return x.Response32
	}
	return 0
}

func (x *FooResponse) GetSubMessage() *FooRequestSubMessage {
	if x != nil {
		return x.SubMessage
	}
	return nil
}

var File_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto protoreflect.FileDescriptor

var file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDesc = []byte{
	0x0a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x63, 0x73, 0x69, 0x2f, 0x63, 0x73, 0x69, 0x2d,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x63, 0x73, 0x69, 0x61, 0x70, 0x69, 0x67, 0x65, 0x6e,
	0x2f, 0x6e, 0x65, 0x77, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x60, 0x0a, 0x0a, 0x46,
	0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x75, 0x62,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x73,
	0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6c, 0x6f,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x22, 0x73, 0x0a,
	0x14, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x75, 0x62, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6c, 0x61, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x6c, 0x61, 0x68, 0x12, 0x47, 0x0a, 0x0d, 0x73, 0x75, 0x62,
	0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x75, 0x62, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x2f, 0x0a, 0x17, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x75, 0x62, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x08, 0x52, 0x05, 0x62, 0x6f,
	0x6f, 0x6c, 0x73, 0x22, 0x6d, 0x0a, 0x0b, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x33, 0x32,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x33, 0x32, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x75, 0x62, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x40, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x34,
	0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x14, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x63, 0x73,
	0x69, 0x2f, 0x63, 0x73, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x63, 0x73, 0x69,
	0x61, 0x70, 0x69, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x65, 0x77, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDescOnce sync.Once
	file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDescData = file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDesc
)

func file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDescGZIP() []byte {
	file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDescOnce.Do(func() {
		file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDescData)
	})
	return file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDescData
}

var file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_goTypes = []interface{}{
	(*FooRequest)(nil),              // 0: v1alpha1.FooRequest
	(*FooRequestSubMessage)(nil),    // 1: v1alpha1.FooRequestSubMessage
	(*FooRequestSubSubMessage)(nil), // 2: v1alpha1.FooRequestSubSubMessage
	(*FooResponse)(nil),             // 3: v1alpha1.FooResponse
}
var file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_depIdxs = []int32{
	1, // 0: v1alpha1.FooRequest.subMessage:type_name -> v1alpha1.FooRequestSubMessage
	2, // 1: v1alpha1.FooRequestSubMessage.subSubMessage:type_name -> v1alpha1.FooRequestSubSubMessage
	1, // 2: v1alpha1.FooResponse.subMessage:type_name -> v1alpha1.FooRequestSubMessage
	0, // 3: v1alpha1.NewGroup.Foo:input_type -> v1alpha1.FooRequest
	3, // 4: v1alpha1.NewGroup.Foo:output_type -> v1alpha1.FooResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_init()
}
func file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_init() {
	if File_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FooRequest); i {
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
		file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FooRequestSubMessage); i {
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
		file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FooRequestSubSubMessage); i {
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
		file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FooResponse); i {
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
			RawDescriptor: file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_goTypes,
		DependencyIndexes: file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_depIdxs,
		MessageInfos:      file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_msgTypes,
	}.Build()
	File_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto = out.File
	file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_rawDesc = nil
	file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_goTypes = nil
	file_github_com_kubernetes_csi_csi_proxy_integrationtests_csiapigen_new_group_api_v1alpha1_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NewGroupClient is the client API for NewGroup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NewGroupClient interface {
	Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error)
}

type newGroupClient struct {
	cc grpc.ClientConnInterface
}

func NewNewGroupClient(cc grpc.ClientConnInterface) NewGroupClient {
	return &newGroupClient{cc}
}

func (c *newGroupClient) Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error) {
	out := new(FooResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NewGroup/Foo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewGroupServer is the server API for NewGroup service.
type NewGroupServer interface {
	Foo(context.Context, *FooRequest) (*FooResponse, error)
}

// UnimplementedNewGroupServer can be embedded to have forward compatible implementations.
type UnimplementedNewGroupServer struct {
}

func (*UnimplementedNewGroupServer) Foo(context.Context, *FooRequest) (*FooResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Foo not implemented")
}

func RegisterNewGroupServer(s *grpc.Server, srv NewGroupServer) {
	s.RegisterService(&_NewGroup_serviceDesc, srv)
}

func _NewGroup_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewGroupServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NewGroup/Foo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewGroupServer).Foo(ctx, req.(*FooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NewGroup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.NewGroup",
	HandlerType: (*NewGroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Foo",
			Handler:    _NewGroup_Foo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/kubernetes-csi/csi-proxy/integrationtests/csiapigen/new_group/api/v1alpha1/api.proto",
}
