// Copyright 2022 SaferPlace

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: incident/v1/incident.proto

package incident

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

// Resolution defines how the incident was resolved, or not yet.
type Resolution int32

const (
	Resolution_RESOLUTION_UNSPECIFIED Resolution = 0
	Resolution_RESOLUTION_REJECTED    Resolution = 1
	Resolution_RESOLUTION_ACCEPTED    Resolution = 2
	Resolution_RESOLUTION_ALERTED     Resolution = 3
)

// Enum value maps for Resolution.
var (
	Resolution_name = map[int32]string{
		0: "RESOLUTION_UNSPECIFIED",
		1: "RESOLUTION_REJECTED",
		2: "RESOLUTION_ACCEPTED",
		3: "RESOLUTION_ALERTED",
	}
	Resolution_value = map[string]int32{
		"RESOLUTION_UNSPECIFIED": 0,
		"RESOLUTION_REJECTED":    1,
		"RESOLUTION_ACCEPTED":    2,
		"RESOLUTION_ALERTED":     3,
	}
)

func (x Resolution) Enum() *Resolution {
	p := new(Resolution)
	*p = x
	return p
}

func (x Resolution) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Resolution) Descriptor() protoreflect.EnumDescriptor {
	return file_incident_v1_incident_proto_enumTypes[0].Descriptor()
}

func (Resolution) Type() protoreflect.EnumType {
	return &file_incident_v1_incident_proto_enumTypes[0]
}

func (x Resolution) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Resolution.Descriptor instead.
func (Resolution) EnumDescriptor() ([]byte, []int) {
	return file_incident_v1_incident_proto_rawDescGZIP(), []int{0}
}

// Incident defines an individual incident as it happened.
type Incident struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the report. This would be generated on ingestion, and therefore
	// does not have to be specified by the client as it will be overriden.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Timestamp (unix) at which the incident occured, in seconds.
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Lat (lattitude) of the incident.
	Lat float32 `protobuf:"fixed32,3,opt,name=lat,proto3" json:"lat,omitempty"`
	// Lon (Longitude) of the incident.
	Lon float32 `protobuf:"fixed32,4,opt,name=lon,proto3" json:"lon,omitempty"`
	// Description provided by the user when sending the incident.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resolution of the incident.
	Resolution Resolution `protobuf:"varint,6,opt,name=resolution,proto3,enum=incident.v1.Resolution" json:"resolution,omitempty"`
	// Comments provided by the reviewers.
	ReviewerComments []*Comment `protobuf:"bytes,7,rep,name=reviewer_comments,json=reviewerComments,proto3" json:"reviewer_comments,omitempty"`
	// Tags added by the reviewer to categorize the incident. This might be
	// updated in the future to maybe not be free-form but rather just specific
	// categories.
	Tags []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Incident) Reset() {
	*x = Incident{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incident_v1_incident_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Incident) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Incident) ProtoMessage() {}

func (x *Incident) ProtoReflect() protoreflect.Message {
	mi := &file_incident_v1_incident_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Incident.ProtoReflect.Descriptor instead.
func (*Incident) Descriptor() ([]byte, []int) {
	return file_incident_v1_incident_proto_rawDescGZIP(), []int{0}
}

func (x *Incident) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Incident) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Incident) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Incident) GetLon() float32 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *Incident) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Incident) GetResolution() Resolution {
	if x != nil {
		return x.Resolution
	}
	return Resolution_RESOLUTION_UNSPECIFIED
}

func (x *Incident) GetReviewerComments() []*Comment {
	if x != nil {
		return x.ReviewerComments
	}
	return nil
}

func (x *Incident) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// Comment left by the reviewer.
type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp (unix) of the comment, in seconds.
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// AuthorID is the author of the comment
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	// Message left in the comment
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incident_v1_incident_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_incident_v1_incident_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_incident_v1_incident_proto_rawDescGZIP(), []int{1}
}

func (x *Comment) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Comment) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Comment) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_incident_v1_incident_proto protoreflect.FileDescriptor

var file_incident_v1_incident_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x8e, 0x02, 0x0a, 0x08, 0x49, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x11, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x10, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x5e, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x72, 0x0a, 0x0a, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x53, 0x4f,
	0x4c, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x55, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x43, 0x45,
	0x50, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x55,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x93,
	0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x42, 0x0d, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x24, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x66, 0x65, 0x72, 0x2e, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0xa2, 0x02, 0x03, 0x49, 0x58, 0x58, 0xaa,
	0x02, 0x0b, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b,
	0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x49, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_incident_v1_incident_proto_rawDescOnce sync.Once
	file_incident_v1_incident_proto_rawDescData = file_incident_v1_incident_proto_rawDesc
)

func file_incident_v1_incident_proto_rawDescGZIP() []byte {
	file_incident_v1_incident_proto_rawDescOnce.Do(func() {
		file_incident_v1_incident_proto_rawDescData = protoimpl.X.CompressGZIP(file_incident_v1_incident_proto_rawDescData)
	})
	return file_incident_v1_incident_proto_rawDescData
}

var file_incident_v1_incident_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_incident_v1_incident_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_incident_v1_incident_proto_goTypes = []interface{}{
	(Resolution)(0),  // 0: incident.v1.Resolution
	(*Incident)(nil), // 1: incident.v1.Incident
	(*Comment)(nil),  // 2: incident.v1.Comment
}
var file_incident_v1_incident_proto_depIdxs = []int32{
	0, // 0: incident.v1.Incident.resolution:type_name -> incident.v1.Resolution
	2, // 1: incident.v1.Incident.reviewer_comments:type_name -> incident.v1.Comment
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_incident_v1_incident_proto_init() }
func file_incident_v1_incident_proto_init() {
	if File_incident_v1_incident_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_incident_v1_incident_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Incident); i {
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
		file_incident_v1_incident_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
			RawDescriptor: file_incident_v1_incident_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_incident_v1_incident_proto_goTypes,
		DependencyIndexes: file_incident_v1_incident_proto_depIdxs,
		EnumInfos:         file_incident_v1_incident_proto_enumTypes,
		MessageInfos:      file_incident_v1_incident_proto_msgTypes,
	}.Build()
	File_incident_v1_incident_proto = out.File
	file_incident_v1_incident_proto_rawDesc = nil
	file_incident_v1_incident_proto_goTypes = nil
	file_incident_v1_incident_proto_depIdxs = nil
}
