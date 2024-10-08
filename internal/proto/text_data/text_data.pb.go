// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: internal/proto/text_data/text_data.proto

package text_data

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

type PostTextDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Metadata string `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *PostTextDataRequest) Reset() {
	*x = PostTextDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_text_data_text_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTextDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTextDataRequest) ProtoMessage() {}

func (x *PostTextDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_text_data_text_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTextDataRequest.ProtoReflect.Descriptor instead.
func (*PostTextDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_text_data_text_data_proto_rawDescGZIP(), []int{0}
}

func (x *PostTextDataRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *PostTextDataRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type PostTextDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Metadata  string `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *PostTextDataResponse) Reset() {
	*x = PostTextDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_text_data_text_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTextDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTextDataResponse) ProtoMessage() {}

func (x *PostTextDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_text_data_text_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTextDataResponse.ProtoReflect.Descriptor instead.
func (*PostTextDataResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_text_data_text_data_proto_rawDescGZIP(), []int{1}
}

func (x *PostTextDataResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostTextDataResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *PostTextDataResponse) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *PostTextDataResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PostTextDataResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetTextDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Metadata string `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *GetTextDataRequest) Reset() {
	*x = GetTextDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_text_data_text_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextDataRequest) ProtoMessage() {}

func (x *GetTextDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_text_data_text_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextDataRequest.ProtoReflect.Descriptor instead.
func (*GetTextDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_text_data_text_data_proto_rawDescGZIP(), []int{2}
}

func (x *GetTextDataRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetTextDataRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type TextData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId   string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Text      string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Metadata  string `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *TextData) Reset() {
	*x = TextData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_text_data_text_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextData) ProtoMessage() {}

func (x *TextData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_text_data_text_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextData.ProtoReflect.Descriptor instead.
func (*TextData) Descriptor() ([]byte, []int) {
	return file_internal_proto_text_data_text_data_proto_rawDescGZIP(), []int{3}
}

func (x *TextData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TextData) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *TextData) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TextData) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *TextData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *TextData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetTextDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text []*TextData `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
}

func (x *GetTextDataResponse) Reset() {
	*x = GetTextDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_text_data_text_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextDataResponse) ProtoMessage() {}

func (x *GetTextDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_text_data_text_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextDataResponse.ProtoReflect.Descriptor instead.
func (*GetTextDataResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_text_data_text_data_proto_rawDescGZIP(), []int{4}
}

func (x *GetTextDataResponse) GetText() []*TextData {
	if x != nil {
		return x.Text
	}
	return nil
}

var File_internal_proto_text_data_text_data_proto protoreflect.FileDescriptor

var file_internal_proto_text_data_text_data_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x45, 0x0a, 0x13, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x50, 0x6f, 0x73,
	0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x44, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa3, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3a, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0xa8, 0x01, 0x0a, 0x0f, 0x54, 0x65, 0x78, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x10, 0x50,
	0x6f, 0x73, 0x74, 0x53, 0x61, 0x76, 0x65, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x61, 0x64, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x73, 0x6d, 0x6b, 0x64, 0x65, 0x6e, 0x69, 0x73, 0x2f, 0x79, 0x61, 0x70, 0x2d, 0x67,
	0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_text_data_text_data_proto_rawDescOnce sync.Once
	file_internal_proto_text_data_text_data_proto_rawDescData = file_internal_proto_text_data_text_data_proto_rawDesc
)

func file_internal_proto_text_data_text_data_proto_rawDescGZIP() []byte {
	file_internal_proto_text_data_text_data_proto_rawDescOnce.Do(func() {
		file_internal_proto_text_data_text_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_text_data_text_data_proto_rawDescData)
	})
	return file_internal_proto_text_data_text_data_proto_rawDescData
}

var file_internal_proto_text_data_text_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_proto_text_data_text_data_proto_goTypes = []interface{}{
	(*PostTextDataRequest)(nil),  // 0: proto.PostTextDataRequest
	(*PostTextDataResponse)(nil), // 1: proto.PostTextDataResponse
	(*GetTextDataRequest)(nil),   // 2: proto.GetTextDataRequest
	(*TextData)(nil),             // 3: proto.TextData
	(*GetTextDataResponse)(nil),  // 4: proto.GetTextDataResponse
}
var file_internal_proto_text_data_text_data_proto_depIdxs = []int32{
	3, // 0: proto.GetTextDataResponse.text:type_name -> proto.TextData
	0, // 1: proto.TextDataService.PostSaveTextData:input_type -> proto.PostTextDataRequest
	2, // 2: proto.TextDataService.GetLoadTextData:input_type -> proto.GetTextDataRequest
	1, // 3: proto.TextDataService.PostSaveTextData:output_type -> proto.PostTextDataResponse
	4, // 4: proto.TextDataService.GetLoadTextData:output_type -> proto.GetTextDataResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_proto_text_data_text_data_proto_init() }
func file_internal_proto_text_data_text_data_proto_init() {
	if File_internal_proto_text_data_text_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_text_data_text_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTextDataRequest); i {
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
		file_internal_proto_text_data_text_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTextDataResponse); i {
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
		file_internal_proto_text_data_text_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextDataRequest); i {
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
		file_internal_proto_text_data_text_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextData); i {
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
		file_internal_proto_text_data_text_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextDataResponse); i {
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
			RawDescriptor: file_internal_proto_text_data_text_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_text_data_text_data_proto_goTypes,
		DependencyIndexes: file_internal_proto_text_data_text_data_proto_depIdxs,
		MessageInfos:      file_internal_proto_text_data_text_data_proto_msgTypes,
	}.Build()
	File_internal_proto_text_data_text_data_proto = out.File
	file_internal_proto_text_data_text_data_proto_rawDesc = nil
	file_internal_proto_text_data_text_data_proto_goTypes = nil
	file_internal_proto_text_data_text_data_proto_depIdxs = nil
}
