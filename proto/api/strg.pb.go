// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: strg.proto

package strg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image []byte `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *SendImageRequest) Reset() {
	*x = SendImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendImageRequest) ProtoMessage() {}

func (x *SendImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendImageRequest.ProtoReflect.Descriptor instead.
func (*SendImageRequest) Descriptor() ([]byte, []int) {
	return file_strg_proto_rawDescGZIP(), []int{0}
}

func (x *SendImageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SendImageRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type GetImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetImageRequest) Reset() {
	*x = GetImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageRequest) ProtoMessage() {}

func (x *GetImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageRequest.ProtoReflect.Descriptor instead.
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return file_strg_proto_rawDescGZIP(), []int{1}
}

func (x *GetImageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetImageResponse) Reset() {
	*x = GetImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageResponse) ProtoMessage() {}

func (x *GetImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageResponse.ProtoReflect.Descriptor instead.
func (*GetImageResponse) Descriptor() ([]byte, []int) {
	return file_strg_proto_rawDescGZIP(), []int{2}
}

func (x *GetImageResponse) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type GetImagesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*Images `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *GetImagesListResponse) Reset() {
	*x = GetImagesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImagesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesListResponse) ProtoMessage() {}

func (x *GetImagesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesListResponse.ProtoReflect.Descriptor instead.
func (*GetImagesListResponse) Descriptor() ([]byte, []int) {
	return file_strg_proto_rawDescGZIP(), []int{3}
}

func (x *GetImagesListResponse) GetImages() []*Images {
	if x != nil {
		return x.Images
	}
	return nil
}

type Images struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreationTime     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	ModificationTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=modificationTime,proto3" json:"modificationTime,omitempty"`
}

func (x *Images) Reset() {
	*x = Images{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Images) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Images) ProtoMessage() {}

func (x *Images) ProtoReflect() protoreflect.Message {
	mi := &file_strg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Images.ProtoReflect.Descriptor instead.
func (*Images) Descriptor() ([]byte, []int) {
	return file_strg_proto_rawDescGZIP(), []int{4}
}

func (x *Images) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Images) GetCreationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *Images) GetModificationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ModificationTime
	}
	return nil
}

var File_strg_proto protoreflect.FileDescriptor

var file_strg_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x74, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x74,
	0x72, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3c, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x25,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0x3d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x74, 0x72, 0x67, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0xa4,
	0x01, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a,
	0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x46, 0x0a,
	0x10, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x10, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xd2, 0x01, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x72, 0x67, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x15, 0x2e, 0x73, 0x74, 0x72, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x74, 0x72, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x73, 0x74,
	0x72, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b,
	0x73, 0x74, 0x72, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strg_proto_rawDescOnce sync.Once
	file_strg_proto_rawDescData = file_strg_proto_rawDesc
)

func file_strg_proto_rawDescGZIP() []byte {
	file_strg_proto_rawDescOnce.Do(func() {
		file_strg_proto_rawDescData = protoimpl.X.CompressGZIP(file_strg_proto_rawDescData)
	})
	return file_strg_proto_rawDescData
}

var file_strg_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_strg_proto_goTypes = []interface{}{
	(*SendImageRequest)(nil),      // 0: strg.SendImageRequest
	(*GetImageRequest)(nil),       // 1: strg.GetImageRequest
	(*GetImageResponse)(nil),      // 2: strg.GetImageResponse
	(*GetImagesListResponse)(nil), // 3: strg.GetImagesListResponse
	(*Images)(nil),                // 4: strg.Images
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_strg_proto_depIdxs = []int32{
	4, // 0: strg.GetImagesListResponse.images:type_name -> strg.Images
	5, // 1: strg.Images.creationTime:type_name -> google.protobuf.Timestamp
	5, // 2: strg.Images.modificationTime:type_name -> google.protobuf.Timestamp
	0, // 3: strg.ImageStorage.SendImage:input_type -> strg.SendImageRequest
	1, // 4: strg.ImageStorage.GetImage:input_type -> strg.GetImageRequest
	6, // 5: strg.ImageStorage.GetImagesList:input_type -> google.protobuf.Empty
	6, // 6: strg.ImageStorage.SendImage:output_type -> google.protobuf.Empty
	2, // 7: strg.ImageStorage.GetImage:output_type -> strg.GetImageResponse
	3, // 8: strg.ImageStorage.GetImagesList:output_type -> strg.GetImagesListResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_strg_proto_init() }
func file_strg_proto_init() {
	if File_strg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendImageRequest); i {
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
		file_strg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageRequest); i {
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
		file_strg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageResponse); i {
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
		file_strg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImagesListResponse); i {
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
		file_strg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Images); i {
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
			RawDescriptor: file_strg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strg_proto_goTypes,
		DependencyIndexes: file_strg_proto_depIdxs,
		MessageInfos:      file_strg_proto_msgTypes,
	}.Build()
	File_strg_proto = out.File
	file_strg_proto_rawDesc = nil
	file_strg_proto_goTypes = nil
	file_strg_proto_depIdxs = nil
}
