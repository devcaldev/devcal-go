// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: devcal.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type InsertEventParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dtstart *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=Dtstart,proto3" json:"Dtstart,omitempty"`
	Dtend   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=Dtend,proto3" json:"Dtend,omitempty"`
	Rrule   *string                `protobuf:"bytes,3,opt,name=Rrule,proto3,oneof" json:"Rrule,omitempty"`
	Props   *structpb.Struct       `protobuf:"bytes,4,opt,name=Props,proto3,oneof" json:"Props,omitempty"`
}

func (x *InsertEventParams) Reset() {
	*x = InsertEventParams{}
	mi := &file_devcal_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertEventParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertEventParams) ProtoMessage() {}

func (x *InsertEventParams) ProtoReflect() protoreflect.Message {
	mi := &file_devcal_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertEventParams.ProtoReflect.Descriptor instead.
func (*InsertEventParams) Descriptor() ([]byte, []int) {
	return file_devcal_proto_rawDescGZIP(), []int{0}
}

func (x *InsertEventParams) GetDtstart() *timestamppb.Timestamp {
	if x != nil {
		return x.Dtstart
	}
	return nil
}

func (x *InsertEventParams) GetDtend() *timestamppb.Timestamp {
	if x != nil {
		return x.Dtend
	}
	return nil
}

func (x *InsertEventParams) GetRrule() string {
	if x != nil && x.Rrule != nil {
		return *x.Rrule
	}
	return ""
}

func (x *InsertEventParams) GetProps() *structpb.Struct {
	if x != nil {
		return x.Props
	}
	return nil
}

type GetEventParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetEventParams) Reset() {
	*x = GetEventParams{}
	mi := &file_devcal_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventParams) ProtoMessage() {}

func (x *GetEventParams) ProtoReflect() protoreflect.Message {
	mi := &file_devcal_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventParams.ProtoReflect.Descriptor instead.
func (*GetEventParams) Descriptor() ([]byte, []int) {
	return file_devcal_proto_rawDescGZIP(), []int{1}
}

func (x *GetEventParams) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type ListEventsRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=Date,proto3" json:"Date,omitempty"`
	Period string                 `protobuf:"bytes,2,opt,name=Period,proto3" json:"Period,omitempty"`
}

func (x *ListEventsRange) Reset() {
	*x = ListEventsRange{}
	mi := &file_devcal_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEventsRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsRange) ProtoMessage() {}

func (x *ListEventsRange) ProtoReflect() protoreflect.Message {
	mi := &file_devcal_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsRange.ProtoReflect.Descriptor instead.
func (*ListEventsRange) Descriptor() ([]byte, []int) {
	return file_devcal_proto_rawDescGZIP(), []int{2}
}

func (x *ListEventsRange) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *ListEventsRange) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

type ListEventsParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range *ListEventsRange `protobuf:"bytes,1,opt,name=Range,proto3,oneof" json:"Range,omitempty"`
	Props *structpb.Struct `protobuf:"bytes,3,opt,name=Props,proto3,oneof" json:"Props,omitempty"`
}

func (x *ListEventsParams) Reset() {
	*x = ListEventsParams{}
	mi := &file_devcal_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEventsParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsParams) ProtoMessage() {}

func (x *ListEventsParams) ProtoReflect() protoreflect.Message {
	mi := &file_devcal_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsParams.ProtoReflect.Descriptor instead.
func (*ListEventsParams) Descriptor() ([]byte, []int) {
	return file_devcal_proto_rawDescGZIP(), []int{3}
}

func (x *ListEventsParams) GetRange() *ListEventsRange {
	if x != nil {
		return x.Range
	}
	return nil
}

func (x *ListEventsParams) GetProps() *structpb.Struct {
	if x != nil {
		return x.Props
	}
	return nil
}

type UpdateEventParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Dtstart *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=Dtstart,proto3,oneof" json:"Dtstart,omitempty"`
	Dtend   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=Dtend,proto3,oneof" json:"Dtend,omitempty"`
	Rrule   *string                `protobuf:"bytes,4,opt,name=Rrule,proto3,oneof" json:"Rrule,omitempty"`
	Props   *structpb.Struct       `protobuf:"bytes,5,opt,name=Props,proto3,oneof" json:"Props,omitempty"`
}

func (x *UpdateEventParams) Reset() {
	*x = UpdateEventParams{}
	mi := &file_devcal_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEventParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventParams) ProtoMessage() {}

func (x *UpdateEventParams) ProtoReflect() protoreflect.Message {
	mi := &file_devcal_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventParams.ProtoReflect.Descriptor instead.
func (*UpdateEventParams) Descriptor() ([]byte, []int) {
	return file_devcal_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEventParams) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateEventParams) GetDtstart() *timestamppb.Timestamp {
	if x != nil {
		return x.Dtstart
	}
	return nil
}

func (x *UpdateEventParams) GetDtend() *timestamppb.Timestamp {
	if x != nil {
		return x.Dtend
	}
	return nil
}

func (x *UpdateEventParams) GetRrule() string {
	if x != nil && x.Rrule != nil {
		return *x.Rrule
	}
	return ""
}

func (x *UpdateEventParams) GetProps() *structpb.Struct {
	if x != nil {
		return x.Props
	}
	return nil
}

type DeleteEventParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteEventParams) Reset() {
	*x = DeleteEventParams{}
	mi := &file_devcal_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEventParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventParams) ProtoMessage() {}

func (x *DeleteEventParams) ProtoReflect() protoreflect.Message {
	mi := &file_devcal_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventParams.ProtoReflect.Descriptor instead.
func (*DeleteEventParams) Descriptor() ([]byte, []int) {
	return file_devcal_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteEventParams) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	AccountID string                 `protobuf:"bytes,2,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	Dtstart   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=Dtstart,proto3" json:"Dtstart,omitempty"`
	Dtend     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=Dtend,proto3" json:"Dtend,omitempty"`
	Rrule     string                 `protobuf:"bytes,5,opt,name=Rrule,proto3" json:"Rrule,omitempty"`
	Props     *structpb.Struct       `protobuf:"bytes,6,opt,name=Props,proto3" json:"Props,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_devcal_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_devcal_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_devcal_proto_rawDescGZIP(), []int{6}
}

func (x *Event) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Event) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *Event) GetDtstart() *timestamppb.Timestamp {
	if x != nil {
		return x.Dtstart
	}
	return nil
}

func (x *Event) GetDtend() *timestamppb.Timestamp {
	if x != nil {
		return x.Dtend
	}
	return nil
}

func (x *Event) GetRrule() string {
	if x != nil {
		return x.Rrule
	}
	return ""
}

func (x *Event) GetProps() *structpb.Struct {
	if x != nil {
		return x.Props
	}
	return nil
}

var File_devcal_proto protoreflect.FileDescriptor

var file_devcal_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x65, 0x76, 0x63, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x65, 0x76, 0x63, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x44, 0x74, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x44, 0x74, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x30,
	0x0a, 0x05, 0x44, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x44, 0x74, 0x65, 0x6e, 0x64,
	0x12, 0x19, 0x0a, 0x05, 0x52, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x05, 0x52, 0x72, 0x75, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x50,
	0x72, 0x6f, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x48, 0x01, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x52, 0x72, 0x75, 0x6c, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x50, 0x72,
	0x6f, 0x70, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x59, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x22, 0x8e, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x61, 0x6c, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x50, 0x72, 0x6f,
	0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x48, 0x01, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x50, 0x72, 0x6f, 0x70,
	0x73, 0x22, 0x8e, 0x02, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x07, 0x44, 0x74, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x07, 0x44, 0x74, 0x73, 0x74, 0x61, 0x72, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x35, 0x0a, 0x05, 0x44, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52,
	0x05, 0x44, 0x74, 0x65, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x52, 0x72, 0x75,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x52, 0x72, 0x75, 0x6c,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x03, 0x52, 0x05,
	0x50, 0x72, 0x6f, 0x70, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x44, 0x74, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x44, 0x74, 0x65, 0x6e, 0x64, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x52, 0x72, 0x75, 0x6c, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x50, 0x72, 0x6f,
	0x70, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0xe2, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x34, 0x0a, 0x07, 0x44, 0x74, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x44, 0x74,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x44, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x05, 0x44, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x72, 0x75, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x2d, 0x0a,
	0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x32, 0xc2, 0x02, 0x0a,
	0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39,
	0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e,
	0x64, 0x65, 0x76, 0x63, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0d, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x61,
	0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x61, 0x6c, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0d, 0x2e,
	0x64, 0x65, 0x76, 0x63, 0x61, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x64,
	0x65, 0x76, 0x63, 0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0d, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x61, 0x6c, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x42, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x61,
	0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x64,
	0x65, 0x76, 0x63, 0x61, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x64, 0x65, 0x76, 0x63, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_devcal_proto_rawDescOnce sync.Once
	file_devcal_proto_rawDescData = file_devcal_proto_rawDesc
)

func file_devcal_proto_rawDescGZIP() []byte {
	file_devcal_proto_rawDescOnce.Do(func() {
		file_devcal_proto_rawDescData = protoimpl.X.CompressGZIP(file_devcal_proto_rawDescData)
	})
	return file_devcal_proto_rawDescData
}

var file_devcal_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_devcal_proto_goTypes = []any{
	(*InsertEventParams)(nil),     // 0: devcal.InsertEventParams
	(*GetEventParams)(nil),        // 1: devcal.GetEventParams
	(*ListEventsRange)(nil),       // 2: devcal.ListEventsRange
	(*ListEventsParams)(nil),      // 3: devcal.ListEventsParams
	(*UpdateEventParams)(nil),     // 4: devcal.UpdateEventParams
	(*DeleteEventParams)(nil),     // 5: devcal.DeleteEventParams
	(*Event)(nil),                 // 6: devcal.Event
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 8: google.protobuf.Struct
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_devcal_proto_depIdxs = []int32{
	7,  // 0: devcal.InsertEventParams.Dtstart:type_name -> google.protobuf.Timestamp
	7,  // 1: devcal.InsertEventParams.Dtend:type_name -> google.protobuf.Timestamp
	8,  // 2: devcal.InsertEventParams.Props:type_name -> google.protobuf.Struct
	7,  // 3: devcal.ListEventsRange.Date:type_name -> google.protobuf.Timestamp
	2,  // 4: devcal.ListEventsParams.Range:type_name -> devcal.ListEventsRange
	8,  // 5: devcal.ListEventsParams.Props:type_name -> google.protobuf.Struct
	7,  // 6: devcal.UpdateEventParams.Dtstart:type_name -> google.protobuf.Timestamp
	7,  // 7: devcal.UpdateEventParams.Dtend:type_name -> google.protobuf.Timestamp
	8,  // 8: devcal.UpdateEventParams.Props:type_name -> google.protobuf.Struct
	7,  // 9: devcal.Event.Dtstart:type_name -> google.protobuf.Timestamp
	7,  // 10: devcal.Event.Dtend:type_name -> google.protobuf.Timestamp
	8,  // 11: devcal.Event.Props:type_name -> google.protobuf.Struct
	0,  // 12: devcal.EventsService.InsertEvent:input_type -> devcal.InsertEventParams
	1,  // 13: devcal.EventsService.GetEvent:input_type -> devcal.GetEventParams
	3,  // 14: devcal.EventsService.ListEvents:input_type -> devcal.ListEventsParams
	4,  // 15: devcal.EventsService.UpdateEvent:input_type -> devcal.UpdateEventParams
	5,  // 16: devcal.EventsService.DeleteEvent:input_type -> devcal.DeleteEventParams
	6,  // 17: devcal.EventsService.InsertEvent:output_type -> devcal.Event
	6,  // 18: devcal.EventsService.GetEvent:output_type -> devcal.Event
	6,  // 19: devcal.EventsService.ListEvents:output_type -> devcal.Event
	9,  // 20: devcal.EventsService.UpdateEvent:output_type -> google.protobuf.Empty
	9,  // 21: devcal.EventsService.DeleteEvent:output_type -> google.protobuf.Empty
	17, // [17:22] is the sub-list for method output_type
	12, // [12:17] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_devcal_proto_init() }
func file_devcal_proto_init() {
	if File_devcal_proto != nil {
		return
	}
	file_devcal_proto_msgTypes[0].OneofWrappers = []any{}
	file_devcal_proto_msgTypes[3].OneofWrappers = []any{}
	file_devcal_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_devcal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_devcal_proto_goTypes,
		DependencyIndexes: file_devcal_proto_depIdxs,
		MessageInfos:      file_devcal_proto_msgTypes,
	}.Build()
	File_devcal_proto = out.File
	file_devcal_proto_rawDesc = nil
	file_devcal_proto_goTypes = nil
	file_devcal_proto_depIdxs = nil
}
