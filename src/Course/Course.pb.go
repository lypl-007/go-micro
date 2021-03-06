// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v4.0.0
// source: Course.proto

package Course

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CourseModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId   int32  `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	CourseName string `protobuf:"bytes,2,opt,name=course_name,json=courseName,proto3" json:"course_name,omitempty"`
}

func (x *CourseModel) Reset() {
	*x = CourseModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseModel) ProtoMessage() {}

func (x *CourseModel) ProtoReflect() protoreflect.Message {
	mi := &file_Course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseModel.ProtoReflect.Descriptor instead.
func (*CourseModel) Descriptor() ([]byte, []int) {
	return file_Course_proto_rawDescGZIP(), []int{0}
}

func (x *CourseModel) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *CourseModel) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_Course_proto_rawDescGZIP(), []int{1}
}

func (x *ListRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*CourseModel `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Course_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Course_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_Course_proto_rawDescGZIP(), []int{2}
}

func (x *ListResponse) GetResult() []*CourseModel {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_Course_proto protoreflect.FileDescriptor

var file_Course_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x4b, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0x4a, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x54,
	0x6f, 0x70, 0x12, 0x13, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Course_proto_rawDescOnce sync.Once
	file_Course_proto_rawDescData = file_Course_proto_rawDesc
)

func file_Course_proto_rawDescGZIP() []byte {
	file_Course_proto_rawDescOnce.Do(func() {
		file_Course_proto_rawDescData = protoimpl.X.CompressGZIP(file_Course_proto_rawDescData)
	})
	return file_Course_proto_rawDescData
}

var file_Course_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_Course_proto_goTypes = []interface{}{
	(*CourseModel)(nil),  // 0: Course.CourseModel
	(*ListRequest)(nil),  // 1: Course.ListRequest
	(*ListResponse)(nil), // 2: Course.ListResponse
}
var file_Course_proto_depIdxs = []int32{
	0, // 0: Course.ListResponse.result:type_name -> Course.CourseModel
	1, // 1: Course.CourseService.ListForTop:input_type -> Course.ListRequest
	2, // 2: Course.CourseService.ListForTop:output_type -> Course.ListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Course_proto_init() }
func file_Course_proto_init() {
	if File_Course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseModel); i {
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
		file_Course_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_Course_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_Course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Course_proto_goTypes,
		DependencyIndexes: file_Course_proto_depIdxs,
		MessageInfos:      file_Course_proto_msgTypes,
	}.Build()
	File_Course_proto = out.File
	file_Course_proto_rawDesc = nil
	file_Course_proto_goTypes = nil
	file_Course_proto_depIdxs = nil
}
