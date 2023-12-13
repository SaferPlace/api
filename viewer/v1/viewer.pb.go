// Copyright 2023 SaferPlace

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: viewer/v1/viewer.proto

package viewer

import (
	v1 "api.safer.place/incident/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// ViewInRadiusRequest shows all incidents within specified radius (in meters).
type ViewInRadiusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Center of the search radius
	Center *v1.Coordinates `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty"`
	// Radius from the place that the incident happened in, in meters.
	Radius float64 `protobuf:"fixed64,2,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *ViewInRadiusRequest) Reset() {
	*x = ViewInRadiusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_viewer_v1_viewer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewInRadiusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewInRadiusRequest) ProtoMessage() {}

func (x *ViewInRadiusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_viewer_v1_viewer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewInRadiusRequest.ProtoReflect.Descriptor instead.
func (*ViewInRadiusRequest) Descriptor() ([]byte, []int) {
	return file_viewer_v1_viewer_proto_rawDescGZIP(), []int{0}
}

func (x *ViewInRadiusRequest) GetCenter() *v1.Coordinates {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *ViewInRadiusRequest) GetRadius() float64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

type ViewInRadiusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Incidents []*v1.Incident `protobuf:"bytes,1,rep,name=incidents,proto3" json:"incidents,omitempty"`
}

func (x *ViewInRadiusResponse) Reset() {
	*x = ViewInRadiusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_viewer_v1_viewer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewInRadiusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewInRadiusResponse) ProtoMessage() {}

func (x *ViewInRadiusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_viewer_v1_viewer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewInRadiusResponse.ProtoReflect.Descriptor instead.
func (*ViewInRadiusResponse) Descriptor() ([]byte, []int) {
	return file_viewer_v1_viewer_proto_rawDescGZIP(), []int{1}
}

func (x *ViewInRadiusResponse) GetIncidents() []*v1.Incident {
	if x != nil {
		return x.Incidents
	}
	return nil
}

// ViewInRegionRequest shows all incidents in the provided region.
type ViewInRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region *Region `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// since allows to specify when is the last incident timestamp that we want to use. If empty,
	// the range is determined by the server.
	Since *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=since,proto3" json:"since,omitempty"`
}

func (x *ViewInRegionRequest) Reset() {
	*x = ViewInRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_viewer_v1_viewer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewInRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewInRegionRequest) ProtoMessage() {}

func (x *ViewInRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_viewer_v1_viewer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewInRegionRequest.ProtoReflect.Descriptor instead.
func (*ViewInRegionRequest) Descriptor() ([]byte, []int) {
	return file_viewer_v1_viewer_proto_rawDescGZIP(), []int{2}
}

func (x *ViewInRegionRequest) GetRegion() *Region {
	if x != nil {
		return x.Region
	}
	return nil
}

func (x *ViewInRegionRequest) GetSince() *timestamppb.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

// ViewInRegionResponse returns the list of incidents.
type ViewInRegionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Incidents []*v1.Incident `protobuf:"bytes,1,rep,name=incidents,proto3" json:"incidents,omitempty"`
}

func (x *ViewInRegionResponse) Reset() {
	*x = ViewInRegionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_viewer_v1_viewer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewInRegionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewInRegionResponse) ProtoMessage() {}

func (x *ViewInRegionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_viewer_v1_viewer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewInRegionResponse.ProtoReflect.Descriptor instead.
func (*ViewInRegionResponse) Descriptor() ([]byte, []int) {
	return file_viewer_v1_viewer_proto_rawDescGZIP(), []int{3}
}

func (x *ViewInRegionResponse) GetIncidents() []*v1.Incident {
	if x != nil {
		return x.Incidents
	}
	return nil
}

// ViewIncidentRequest will show the incident based on the given id.
type ViewIncidentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ViewIncidentRequest) Reset() {
	*x = ViewIncidentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_viewer_v1_viewer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewIncidentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewIncidentRequest) ProtoMessage() {}

func (x *ViewIncidentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_viewer_v1_viewer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewIncidentRequest.ProtoReflect.Descriptor instead.
func (*ViewIncidentRequest) Descriptor() ([]byte, []int) {
	return file_viewer_v1_viewer_proto_rawDescGZIP(), []int{4}
}

func (x *ViewIncidentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ViewIncidentResponse returns the incident for review.
type ViewIncidentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Incident *v1.Incident `protobuf:"bytes,1,opt,name=incident,proto3" json:"incident,omitempty"`
}

func (x *ViewIncidentResponse) Reset() {
	*x = ViewIncidentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_viewer_v1_viewer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewIncidentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewIncidentResponse) ProtoMessage() {}

func (x *ViewIncidentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_viewer_v1_viewer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewIncidentResponse.ProtoReflect.Descriptor instead.
func (*ViewIncidentResponse) Descriptor() ([]byte, []int) {
	return file_viewer_v1_viewer_proto_rawDescGZIP(), []int{5}
}

func (x *ViewIncidentResponse) GetIncident() *v1.Incident {
	if x != nil {
		return x.Incident
	}
	return nil
}

// ViewAlertingRequest allows to specify which alerts to view which are alerting.
type ViewAlertingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region *Region `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// since allows to specify when is the last incident timestamp. that we want to use.
	Since *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=since,proto3" json:"since,omitempty"`
}

func (x *ViewAlertingRequest) Reset() {
	*x = ViewAlertingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_viewer_v1_viewer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewAlertingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewAlertingRequest) ProtoMessage() {}

func (x *ViewAlertingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_viewer_v1_viewer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewAlertingRequest.ProtoReflect.Descriptor instead.
func (*ViewAlertingRequest) Descriptor() ([]byte, []int) {
	return file_viewer_v1_viewer_proto_rawDescGZIP(), []int{6}
}

func (x *ViewAlertingRequest) GetRegion() *Region {
	if x != nil {
		return x.Region
	}
	return nil
}

func (x *ViewAlertingRequest) GetSince() *timestamppb.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

// ViewAlertingResponse returns the list of incidents.
type ViewAlertingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Incidents []*v1.Incident `protobuf:"bytes,1,rep,name=incidents,proto3" json:"incidents,omitempty"`
}

func (x *ViewAlertingResponse) Reset() {
	*x = ViewAlertingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_viewer_v1_viewer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewAlertingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewAlertingResponse) ProtoMessage() {}

func (x *ViewAlertingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_viewer_v1_viewer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewAlertingResponse.ProtoReflect.Descriptor instead.
func (*ViewAlertingResponse) Descriptor() ([]byte, []int) {
	return file_viewer_v1_viewer_proto_rawDescGZIP(), []int{7}
}

func (x *ViewAlertingResponse) GetIncidents() []*v1.Incident {
	if x != nil {
		return x.Incidents
	}
	return nil
}

// Region defines a block of coordinates which can be used for getting incidents without revealing
// the precise location of the user, but rather multiple small requests can be made around the user.
// This is formed by creating a box where the area of interest lies between north and south, and
// east and west.
//
//	(north <= $lattitude <= south) && (east <= $longitude <= west)
type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	North float64 `protobuf:"fixed64,1,opt,name=north,proto3" json:"north,omitempty"`
	South float64 `protobuf:"fixed64,2,opt,name=south,proto3" json:"south,omitempty"`
	East  float64 `protobuf:"fixed64,3,opt,name=east,proto3" json:"east,omitempty"`
	West  float64 `protobuf:"fixed64,4,opt,name=west,proto3" json:"west,omitempty"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_viewer_v1_viewer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_viewer_v1_viewer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_viewer_v1_viewer_proto_rawDescGZIP(), []int{8}
}

func (x *Region) GetNorth() float64 {
	if x != nil {
		return x.North
	}
	return 0
}

func (x *Region) GetSouth() float64 {
	if x != nil {
		return x.South
	}
	return 0
}

func (x *Region) GetEast() float64 {
	if x != nil {
		return x.East
	}
	return 0
}

func (x *Region) GetWest() float64 {
	if x != nil {
		return x.West
	}
	return 0
}

var File_viewer_v1_viewer_proto protoreflect.FileDescriptor

var file_viewer_v1_viewer_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x65, 0x77,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5f, 0x0a, 0x13, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64,
	0x69, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75,
	0x73, 0x22, 0x4b, 0x0a, 0x14, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x52, 0x61, 0x64, 0x69, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x69, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x09, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x72,
	0x0a, 0x13, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x30, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x69, 0x6e,
	0x63, 0x65, 0x22, 0x4b, 0x0a, 0x14, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x69, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x25, 0x0a, 0x13, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x14, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x08, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x22, 0x72, 0x0a, 0x13, 0x56, 0x69, 0x65, 0x77, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05,
	0x73, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x4b, 0x0a, 0x14, 0x56, 0x69, 0x65, 0x77, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x09, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x5c, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x72, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6e, 0x6f, 0x72,
	0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x61, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x65, 0x61, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x77, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x77, 0x65, 0x73, 0x74,
	0x32, 0xd3, 0x02, 0x0a, 0x0d, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x52, 0x61, 0x64, 0x69,
	0x75, 0x73, 0x12, 0x1e, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x69, 0x65, 0x77, 0x49, 0x6e, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x69, 0x65, 0x77, 0x49, 0x6e, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x83, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x66,
	0x65, 0x72, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa,
	0x02, 0x09, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x56, 0x69,
	0x65, 0x77, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0a, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_viewer_v1_viewer_proto_rawDescOnce sync.Once
	file_viewer_v1_viewer_proto_rawDescData = file_viewer_v1_viewer_proto_rawDesc
)

func file_viewer_v1_viewer_proto_rawDescGZIP() []byte {
	file_viewer_v1_viewer_proto_rawDescOnce.Do(func() {
		file_viewer_v1_viewer_proto_rawDescData = protoimpl.X.CompressGZIP(file_viewer_v1_viewer_proto_rawDescData)
	})
	return file_viewer_v1_viewer_proto_rawDescData
}

var file_viewer_v1_viewer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_viewer_v1_viewer_proto_goTypes = []interface{}{
	(*ViewInRadiusRequest)(nil),   // 0: viewer.v1.ViewInRadiusRequest
	(*ViewInRadiusResponse)(nil),  // 1: viewer.v1.ViewInRadiusResponse
	(*ViewInRegionRequest)(nil),   // 2: viewer.v1.ViewInRegionRequest
	(*ViewInRegionResponse)(nil),  // 3: viewer.v1.ViewInRegionResponse
	(*ViewIncidentRequest)(nil),   // 4: viewer.v1.ViewIncidentRequest
	(*ViewIncidentResponse)(nil),  // 5: viewer.v1.ViewIncidentResponse
	(*ViewAlertingRequest)(nil),   // 6: viewer.v1.ViewAlertingRequest
	(*ViewAlertingResponse)(nil),  // 7: viewer.v1.ViewAlertingResponse
	(*Region)(nil),                // 8: viewer.v1.Region
	(*v1.Coordinates)(nil),        // 9: incident.v1.Coordinates
	(*v1.Incident)(nil),           // 10: incident.v1.Incident
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_viewer_v1_viewer_proto_depIdxs = []int32{
	9,  // 0: viewer.v1.ViewInRadiusRequest.center:type_name -> incident.v1.Coordinates
	10, // 1: viewer.v1.ViewInRadiusResponse.incidents:type_name -> incident.v1.Incident
	8,  // 2: viewer.v1.ViewInRegionRequest.region:type_name -> viewer.v1.Region
	11, // 3: viewer.v1.ViewInRegionRequest.since:type_name -> google.protobuf.Timestamp
	10, // 4: viewer.v1.ViewInRegionResponse.incidents:type_name -> incident.v1.Incident
	10, // 5: viewer.v1.ViewIncidentResponse.incident:type_name -> incident.v1.Incident
	8,  // 6: viewer.v1.ViewAlertingRequest.region:type_name -> viewer.v1.Region
	11, // 7: viewer.v1.ViewAlertingRequest.since:type_name -> google.protobuf.Timestamp
	10, // 8: viewer.v1.ViewAlertingResponse.incidents:type_name -> incident.v1.Incident
	0,  // 9: viewer.v1.ViewerService.ViewInRadius:input_type -> viewer.v1.ViewInRadiusRequest
	2,  // 10: viewer.v1.ViewerService.ViewInRegion:input_type -> viewer.v1.ViewInRegionRequest
	4,  // 11: viewer.v1.ViewerService.ViewIncident:input_type -> viewer.v1.ViewIncidentRequest
	6,  // 12: viewer.v1.ViewerService.ViewAlerting:input_type -> viewer.v1.ViewAlertingRequest
	1,  // 13: viewer.v1.ViewerService.ViewInRadius:output_type -> viewer.v1.ViewInRadiusResponse
	3,  // 14: viewer.v1.ViewerService.ViewInRegion:output_type -> viewer.v1.ViewInRegionResponse
	5,  // 15: viewer.v1.ViewerService.ViewIncident:output_type -> viewer.v1.ViewIncidentResponse
	7,  // 16: viewer.v1.ViewerService.ViewAlerting:output_type -> viewer.v1.ViewAlertingResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_viewer_v1_viewer_proto_init() }
func file_viewer_v1_viewer_proto_init() {
	if File_viewer_v1_viewer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_viewer_v1_viewer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewInRadiusRequest); i {
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
		file_viewer_v1_viewer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewInRadiusResponse); i {
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
		file_viewer_v1_viewer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewInRegionRequest); i {
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
		file_viewer_v1_viewer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewInRegionResponse); i {
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
		file_viewer_v1_viewer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewIncidentRequest); i {
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
		file_viewer_v1_viewer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewIncidentResponse); i {
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
		file_viewer_v1_viewer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewAlertingRequest); i {
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
		file_viewer_v1_viewer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewAlertingResponse); i {
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
		file_viewer_v1_viewer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
			RawDescriptor: file_viewer_v1_viewer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_viewer_v1_viewer_proto_goTypes,
		DependencyIndexes: file_viewer_v1_viewer_proto_depIdxs,
		MessageInfos:      file_viewer_v1_viewer_proto_msgTypes,
	}.Build()
	File_viewer_v1_viewer_proto = out.File
	file_viewer_v1_viewer_proto_rawDesc = nil
	file_viewer_v1_viewer_proto_goTypes = nil
	file_viewer_v1_viewer_proto_depIdxs = nil
}
