// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: client.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type InsertHospitalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Phone   string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email   string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *InsertHospitalRequest) Reset() {
	*x = InsertHospitalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertHospitalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertHospitalRequest) ProtoMessage() {}

func (x *InsertHospitalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertHospitalRequest.ProtoReflect.Descriptor instead.
func (*InsertHospitalRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *InsertHospitalRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InsertHospitalRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *InsertHospitalRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InsertHospitalRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *InsertHospitalRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type InsertHospitalReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *InsertHospitalReply) Reset() {
	*x = InsertHospitalReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertHospitalReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertHospitalReply) ProtoMessage() {}

func (x *InsertHospitalReply) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertHospitalReply.ProtoReflect.Descriptor instead.
func (*InsertHospitalReply) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

func (x *InsertHospitalReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Hospital struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City     string  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Address  string  `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Phone    string  `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email    string  `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Distance float32 `protobuf:"fixed32,6,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *Hospital) Reset() {
	*x = Hospital{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hospital) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hospital) ProtoMessage() {}

func (x *Hospital) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hospital.ProtoReflect.Descriptor instead.
func (*Hospital) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{2}
}

func (x *Hospital) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hospital) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Hospital) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Hospital) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Hospital) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Hospital) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type QueryNearestHospitalsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude        float32 `protobuf:"fixed32,1,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude       float32 `protobuf:"fixed32,2,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	NumberHospitals int32   `protobuf:"varint,3,opt,name=numberHospitals,proto3" json:"numberHospitals,omitempty"`
}

func (x *QueryNearestHospitalsRequest) Reset() {
	*x = QueryNearestHospitalsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryNearestHospitalsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryNearestHospitalsRequest) ProtoMessage() {}

func (x *QueryNearestHospitalsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryNearestHospitalsRequest.ProtoReflect.Descriptor instead.
func (*QueryNearestHospitalsRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{3}
}

func (x *QueryNearestHospitalsRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *QueryNearestHospitalsRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *QueryNearestHospitalsRequest) GetNumberHospitals() int32 {
	if x != nil {
		return x.NumberHospitals
	}
	return 0
}

type QueryNearestHospitalsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NearestHospitals []*Hospital `protobuf:"bytes,1,rep,name=nearestHospitals,proto3" json:"nearestHospitals,omitempty"`
	Message          string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *QueryNearestHospitalsReply) Reset() {
	*x = QueryNearestHospitalsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryNearestHospitalsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryNearestHospitalsReply) ProtoMessage() {}

func (x *QueryNearestHospitalsReply) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryNearestHospitalsReply.ProtoReflect.Descriptor instead.
func (*QueryNearestHospitalsReply) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{4}
}

func (x *QueryNearestHospitalsReply) GetNearestHospitals() []*Hospital {
	if x != nil {
		return x.NearestHospitals
	}
	return nil
}

func (x *QueryNearestHospitalsReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a,
	0x15, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2f, 0x0a, 0x13, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x48, 0x6f,
	0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x82, 0x01, 0x0a,
	0x1c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x48, 0x6f, 0x73,
	0x70, 0x69, 0x74, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x4c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c,
	0x73, 0x22, 0x6d, 0x0a, 0x1a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73,
	0x74, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x35, 0x0a, 0x10, 0x6e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x68, 0x6f, 0x73, 0x70,
	0x69, 0x74, 0x61, 0x6c, 0x52, 0x10, 0x6e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x48, 0x6f, 0x73,
	0x70, 0x69, 0x74, 0x61, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xe0, 0x01, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x5a, 0x0a, 0x0e, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x2e,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x48, 0x6f,
	0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x3a, 0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x73,
	0x12, 0x1d, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x48,
	0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x48, 0x6f,
	0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x73,
	0x3a, 0x01, 0x2a, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x44, 0x54, 0x72, 0x65, 0x73, 0x68, 0x79, 0x2f, 0x73, 0x7a, 0x75, 0x6b, 0x61, 0x6a,
	0x2d, 0x73, 0x7a, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_client_proto_goTypes = []interface{}{
	(*InsertHospitalRequest)(nil),        // 0: InsertHospitalRequest
	(*InsertHospitalReply)(nil),          // 1: InsertHospitalReply
	(*Hospital)(nil),                     // 2: hospital
	(*QueryNearestHospitalsRequest)(nil), // 3: QueryNearestHospitalsRequest
	(*QueryNearestHospitalsReply)(nil),   // 4: QueryNearestHospitalsReply
}
var file_client_proto_depIdxs = []int32{
	2, // 0: QueryNearestHospitalsReply.nearestHospitals:type_name -> hospital
	0, // 1: Client.InsertHospital:input_type -> InsertHospitalRequest
	3, // 2: Client.QueryNearestHospitals:input_type -> QueryNearestHospitalsRequest
	1, // 3: Client.InsertHospital:output_type -> InsertHospitalReply
	4, // 4: Client.QueryNearestHospitals:output_type -> QueryNearestHospitalsReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertHospitalRequest); i {
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
		file_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertHospitalReply); i {
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
		file_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hospital); i {
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
		file_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryNearestHospitalsRequest); i {
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
		file_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryNearestHospitalsReply); i {
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
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}