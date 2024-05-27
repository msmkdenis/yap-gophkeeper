// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: internal/proto/credit_card/credit_card.proto

package credit_card

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

type PostCreditCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number    string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	OwnerName string `protobuf:"bytes,2,opt,name=owner_name,json=ownerName,proto3" json:"owner_name,omitempty"`
	ExpiresAt string `protobuf:"bytes,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	CvvCode   string `protobuf:"bytes,4,opt,name=cvv_code,json=cvvCode,proto3" json:"cvv_code,omitempty"`
	PinCode   string `protobuf:"bytes,5,opt,name=pin_code,json=pinCode,proto3" json:"pin_code,omitempty"`
	Metadata  string `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *PostCreditCardRequest) Reset() {
	*x = PostCreditCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_credit_card_credit_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCreditCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCreditCardRequest) ProtoMessage() {}

func (x *PostCreditCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_credit_card_credit_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCreditCardRequest.ProtoReflect.Descriptor instead.
func (*PostCreditCardRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_credit_card_credit_card_proto_rawDescGZIP(), []int{0}
}

func (x *PostCreditCardRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *PostCreditCardRequest) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

func (x *PostCreditCardRequest) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

func (x *PostCreditCardRequest) GetCvvCode() string {
	if x != nil {
		return x.CvvCode
	}
	return ""
}

func (x *PostCreditCardRequest) GetPinCode() string {
	if x != nil {
		return x.PinCode
	}
	return ""
}

func (x *PostCreditCardRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type PostCreditCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId   string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Number    string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	OwnerName string `protobuf:"bytes,4,opt,name=owner_name,json=ownerName,proto3" json:"owner_name,omitempty"`
	ExpiresAt string `protobuf:"bytes,5,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	CvvCode   string `protobuf:"bytes,6,opt,name=cvv_code,json=cvvCode,proto3" json:"cvv_code,omitempty"`
	PinCode   string `protobuf:"bytes,7,opt,name=pin_code,json=pinCode,proto3" json:"pin_code,omitempty"`
	Metadata  string `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	CreatedAt string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *PostCreditCardResponse) Reset() {
	*x = PostCreditCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_credit_card_credit_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCreditCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCreditCardResponse) ProtoMessage() {}

func (x *PostCreditCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_credit_card_credit_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCreditCardResponse.ProtoReflect.Descriptor instead.
func (*PostCreditCardResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_credit_card_credit_card_proto_rawDescGZIP(), []int{1}
}

func (x *PostCreditCardResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostCreditCardResponse) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *PostCreditCardResponse) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *PostCreditCardResponse) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

func (x *PostCreditCardResponse) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

func (x *PostCreditCardResponse) GetCvvCode() string {
	if x != nil {
		return x.CvvCode
	}
	return ""
}

func (x *PostCreditCardResponse) GetPinCode() string {
	if x != nil {
		return x.PinCode
	}
	return ""
}

func (x *PostCreditCardResponse) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *PostCreditCardResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PostCreditCardResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetCreditCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCreditCardRequest) Reset() {
	*x = GetCreditCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_credit_card_credit_card_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCreditCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCreditCardRequest) ProtoMessage() {}

func (x *GetCreditCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_credit_card_credit_card_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCreditCardRequest.ProtoReflect.Descriptor instead.
func (*GetCreditCardRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_credit_card_credit_card_proto_rawDescGZIP(), []int{2}
}

type CreditCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId   string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Number    string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	OwnerName string `protobuf:"bytes,4,opt,name=owner_name,json=ownerName,proto3" json:"owner_name,omitempty"`
	ExpiresAt string `protobuf:"bytes,5,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	CvvCode   string `protobuf:"bytes,6,opt,name=cvv_code,json=cvvCode,proto3" json:"cvv_code,omitempty"`
	PinCode   string `protobuf:"bytes,7,opt,name=pin_code,json=pinCode,proto3" json:"pin_code,omitempty"`
	Metadata  string `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	CreatedAt string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CreditCard) Reset() {
	*x = CreditCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_credit_card_credit_card_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditCard) ProtoMessage() {}

func (x *CreditCard) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_credit_card_credit_card_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditCard.ProtoReflect.Descriptor instead.
func (*CreditCard) Descriptor() ([]byte, []int) {
	return file_internal_proto_credit_card_credit_card_proto_rawDescGZIP(), []int{3}
}

func (x *CreditCard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreditCard) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *CreditCard) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *CreditCard) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

func (x *CreditCard) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

func (x *CreditCard) GetCvvCode() string {
	if x != nil {
		return x.CvvCode
	}
	return ""
}

func (x *CreditCard) GetPinCode() string {
	if x != nil {
		return x.PinCode
	}
	return ""
}

func (x *CreditCard) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *CreditCard) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CreditCard) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetCreditCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cards []*CreditCard `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *GetCreditCardResponse) Reset() {
	*x = GetCreditCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_credit_card_credit_card_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCreditCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCreditCardResponse) ProtoMessage() {}

func (x *GetCreditCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_credit_card_credit_card_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCreditCardResponse.ProtoReflect.Descriptor instead.
func (*GetCreditCardResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_credit_card_credit_card_proto_rawDescGZIP(), []int{4}
}

func (x *GetCreditCardResponse) GetCards() []*CreditCard {
	if x != nil {
		return x.Cards
	}
	return nil
}

var File_internal_proto_credit_card_credit_card_proto protoreflect.FileDescriptor

var file_internal_proto_credit_card_credit_card_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x76, 0x76, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x76, 0x76, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa9, 0x02, 0x0a, 0x16, 0x50, 0x6f, 0x73, 0x74,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x76, 0x76, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x76, 0x76, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9d, 0x02, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x76, 0x76, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x76, 0x76, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x69, 0x6e, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x40, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x32, 0xb6, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x61, 0x76, 0x65, 0x43,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x61,
	0x64, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x73, 0x6d, 0x6b, 0x64, 0x65, 0x6e, 0x69, 0x73, 0x2f, 0x79,
	0x61, 0x70, 0x2d, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_credit_card_credit_card_proto_rawDescOnce sync.Once
	file_internal_proto_credit_card_credit_card_proto_rawDescData = file_internal_proto_credit_card_credit_card_proto_rawDesc
)

func file_internal_proto_credit_card_credit_card_proto_rawDescGZIP() []byte {
	file_internal_proto_credit_card_credit_card_proto_rawDescOnce.Do(func() {
		file_internal_proto_credit_card_credit_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_credit_card_credit_card_proto_rawDescData)
	})
	return file_internal_proto_credit_card_credit_card_proto_rawDescData
}

var file_internal_proto_credit_card_credit_card_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_proto_credit_card_credit_card_proto_goTypes = []interface{}{
	(*PostCreditCardRequest)(nil),  // 0: proto.PostCreditCardRequest
	(*PostCreditCardResponse)(nil), // 1: proto.PostCreditCardResponse
	(*GetCreditCardRequest)(nil),   // 2: proto.GetCreditCardRequest
	(*CreditCard)(nil),             // 3: proto.CreditCard
	(*GetCreditCardResponse)(nil),  // 4: proto.GetCreditCardResponse
}
var file_internal_proto_credit_card_credit_card_proto_depIdxs = []int32{
	3, // 0: proto.GetCreditCardResponse.cards:type_name -> proto.CreditCard
	0, // 1: proto.CreditCardService.PostSaveCreditCard:input_type -> proto.PostCreditCardRequest
	2, // 2: proto.CreditCardService.GetLoadCreditCard:input_type -> proto.GetCreditCardRequest
	1, // 3: proto.CreditCardService.PostSaveCreditCard:output_type -> proto.PostCreditCardResponse
	4, // 4: proto.CreditCardService.GetLoadCreditCard:output_type -> proto.GetCreditCardResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_proto_credit_card_credit_card_proto_init() }
func file_internal_proto_credit_card_credit_card_proto_init() {
	if File_internal_proto_credit_card_credit_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_credit_card_credit_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCreditCardRequest); i {
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
		file_internal_proto_credit_card_credit_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCreditCardResponse); i {
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
		file_internal_proto_credit_card_credit_card_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCreditCardRequest); i {
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
		file_internal_proto_credit_card_credit_card_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditCard); i {
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
		file_internal_proto_credit_card_credit_card_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCreditCardResponse); i {
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
			RawDescriptor: file_internal_proto_credit_card_credit_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_credit_card_credit_card_proto_goTypes,
		DependencyIndexes: file_internal_proto_credit_card_credit_card_proto_depIdxs,
		MessageInfos:      file_internal_proto_credit_card_credit_card_proto_msgTypes,
	}.Build()
	File_internal_proto_credit_card_credit_card_proto = out.File
	file_internal_proto_credit_card_credit_card_proto_rawDesc = nil
	file_internal_proto_credit_card_credit_card_proto_goTypes = nil
	file_internal_proto_credit_card_credit_card_proto_depIdxs = nil
}
