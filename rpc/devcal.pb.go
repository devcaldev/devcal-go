// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: devcal.proto

package rpc

import (
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

type InsertEventParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dtstart string  `protobuf:"bytes,1,opt,name=Dtstart,proto3" json:"Dtstart,omitempty"`
	Dtend   string  `protobuf:"bytes,2,opt,name=Dtend,proto3" json:"Dtend,omitempty"`
	Rrule   *string `protobuf:"bytes,3,opt,name=Rrule,proto3,oneof" json:"Rrule,omitempty"`
	Props   []byte  `protobuf:"bytes,4,opt,name=Props,proto3,oneof" json:"Props,omitempty"`
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

func (x *InsertEventParams) GetDtstart() string {
	if x != nil {
		return x.Dtstart
	}
	return ""
}

func (x *InsertEventParams) GetDtend() string {
	if x != nil {
		return x.Dtend
	}
	return ""
}

func (x *InsertEventParams) GetRrule() string {
	if x != nil && x.Rrule != nil {
		return *x.Rrule
	}
	return ""
}

func (x *InsertEventParams) GetProps() []byte {
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

type ListEventsParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date   string `protobuf:"bytes,1,opt,name=Date,proto3" json:"Date,omitempty"`
	Period string `protobuf:"bytes,2,opt,name=Period,proto3" json:"Period,omitempty"`
	Props  []byte `protobuf:"bytes,3,opt,name=Props,proto3,oneof" json:"Props,omitempty"`
}

func (x *ListEventsParams) Reset() {
	*x = ListEventsParams{}
	mi := &file_devcal_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEventsParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsParams) ProtoMessage() {}

func (x *ListEventsParams) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListEventsParams.ProtoReflect.Descriptor instead.
func (*ListEventsParams) Descriptor() ([]byte, []int) {
	return file_devcal_proto_rawDescGZIP(), []int{2}
}

func (x *ListEventsParams) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ListEventsParams) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *ListEventsParams) GetProps() []byte {
	if x != nil {
		return x.Props
	}
	return nil
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
	Props     []byte                 `protobuf:"bytes,6,opt,name=Props,proto3" json:"Props,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_devcal_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_devcal_proto_rawDescGZIP(), []int{3}
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

func (x *Event) GetProps() []byte {
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
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x44, 0x74, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x44, 0x74, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x74, 0x65, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x44, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x19, 0x0a,
	0x05, 0x52, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x52, 0x72, 0x75, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x70,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x52, 0x72, 0x75, 0x6c, 0x65, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x63, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x50, 0x72, 0x6f,
	0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x70,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x22, 0xc9,
	0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x07, 0x44, 0x74, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x44, 0x74, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x30, 0x0a, 0x05,
	0x44, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x44, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x52, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52,
	0x72, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x32, 0xba, 0x01, 0x0a, 0x0d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0b,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x64, 0x65,
	0x76, 0x63, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0d, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x61, 0x6c, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0d, 0x2e, 0x64, 0x65,
	0x76, 0x63, 0x61, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x64, 0x65, 0x76,
	0x63, 0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0d, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x61, 0x6c, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x64, 0x65, 0x76, 0x63, 0x61,
	0x6c, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_devcal_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_devcal_proto_goTypes = []any{
	(*InsertEventParams)(nil),     // 0: devcal.InsertEventParams
	(*GetEventParams)(nil),        // 1: devcal.GetEventParams
	(*ListEventsParams)(nil),      // 2: devcal.ListEventsParams
	(*Event)(nil),                 // 3: devcal.Event
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_devcal_proto_depIdxs = []int32{
	4, // 0: devcal.Event.Dtstart:type_name -> google.protobuf.Timestamp
	4, // 1: devcal.Event.Dtend:type_name -> google.protobuf.Timestamp
	0, // 2: devcal.EventsService.InsertEvent:input_type -> devcal.InsertEventParams
	1, // 3: devcal.EventsService.GetEvent:input_type -> devcal.GetEventParams
	2, // 4: devcal.EventsService.ListEvents:input_type -> devcal.ListEventsParams
	3, // 5: devcal.EventsService.InsertEvent:output_type -> devcal.Event
	3, // 6: devcal.EventsService.GetEvent:output_type -> devcal.Event
	3, // 7: devcal.EventsService.ListEvents:output_type -> devcal.Event
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_devcal_proto_init() }
func file_devcal_proto_init() {
	if File_devcal_proto != nil {
		return
	}
	file_devcal_proto_msgTypes[0].OneofWrappers = []any{}
	file_devcal_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_devcal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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
