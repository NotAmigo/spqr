// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: protos/coordinator.proto

package proto

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

type ListRoutersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRoutersRequest) Reset() {
	*x = ListRoutersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_coordinate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoutersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutersRequest) ProtoMessage() {}

func (x *ListRoutersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_coordinate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutersRequest.ProtoReflect.Descriptor instead.
func (*ListRoutersRequest) Descriptor() ([]byte, []int) {
	return file_protos_coordinate_proto_rawDescGZIP(), []int{0}
}

type AddRoutersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddRoutersRequest) Reset() {
	*x = AddRoutersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_coordinate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoutersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoutersRequest) ProtoMessage() {}

func (x *AddRoutersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_coordinate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoutersRequest.ProtoReflect.Descriptor instead.
func (*AddRoutersRequest) Descriptor() ([]byte, []int) {
	return file_protos_coordinate_proto_rawDescGZIP(), []int{1}
}

type ListRoutersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRoutersReply) Reset() {
	*x = ListRoutersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_coordinate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoutersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutersReply) ProtoMessage() {}

func (x *ListRoutersReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_coordinate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutersReply.ProtoReflect.Descriptor instead.
func (*ListRoutersReply) Descriptor() ([]byte, []int) {
	return file_protos_coordinate_proto_rawDescGZIP(), []int{2}
}

type AddRoutersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddRoutersReply) Reset() {
	*x = AddRoutersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_coordinate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoutersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoutersReply) ProtoMessage() {}

func (x *AddRoutersReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_coordinate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoutersReply.ProtoReflect.Descriptor instead.
func (*AddRoutersReply) Descriptor() ([]byte, []int) {
	return file_protos_coordinate_proto_rawDescGZIP(), []int{3}
}

var File_protos_coordinate_proto protoreflect.FileDescriptor

var file_protos_coordinate_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11,
	0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xaf, 0x01, 0x0a, 0x0e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0a,
	0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x73, 0x70, 0x71, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_coordinate_proto_rawDescOnce sync.Once
	file_protos_coordinate_proto_rawDescData = file_protos_coordinate_proto_rawDesc
)

func file_protos_coordinate_proto_rawDescGZIP() []byte {
	file_protos_coordinate_proto_rawDescOnce.Do(func() {
		file_protos_coordinate_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_coordinate_proto_rawDescData)
	})
	return file_protos_coordinate_proto_rawDescData
}

var file_protos_coordinate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_coordinate_proto_goTypes = []interface{}{
	(*ListRoutersRequest)(nil), // 0: yandex.spqr.ListRoutersRequest
	(*AddRoutersRequest)(nil),  // 1: yandex.spqr.AddRoutersRequest
	(*ListRoutersReply)(nil),   // 2: yandex.spqr.ListRoutersReply
	(*AddRoutersReply)(nil),    // 3: yandex.spqr.AddRoutersReply
}
var file_protos_coordinate_proto_depIdxs = []int32{
	0, // 0: yandex.spqr.RoutersService.ListRouters:input_type -> yandex.spqr.ListRoutersRequest
	1, // 1: yandex.spqr.RoutersService.AddRouters:input_type -> yandex.spqr.AddRoutersRequest
	2, // 2: yandex.spqr.RoutersService.ListRouters:output_type -> yandex.spqr.ListRoutersReply
	3, // 3: yandex.spqr.RoutersService.AddRouters:output_type -> yandex.spqr.AddRoutersReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_coordinate_proto_init() }
func file_protos_coordinate_proto_init() {
	if File_protos_coordinate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_coordinate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoutersRequest); i {
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
		file_protos_coordinate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoutersRequest); i {
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
		file_protos_coordinate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoutersReply); i {
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
		file_protos_coordinate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoutersReply); i {
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
			RawDescriptor: file_protos_coordinate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_coordinate_proto_goTypes,
		DependencyIndexes: file_protos_coordinate_proto_depIdxs,
		MessageInfos:      file_protos_coordinate_proto_msgTypes,
	}.Build()
	File_protos_coordinate_proto = out.File
	file_protos_coordinate_proto_rawDesc = nil
	file_protos_coordinate_proto_goTypes = nil
	file_protos_coordinate_proto_depIdxs = nil
}
