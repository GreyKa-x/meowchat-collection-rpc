// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: collection.proto

package pb

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

type Cat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateAt     int64    `protobuf:"varint,2,opt,name=createAt,proto3" json:"createAt,omitempty"`
	Age          string   `protobuf:"bytes,3,opt,name=age,proto3" json:"age,omitempty"`
	CommunityId  string   `protobuf:"bytes,4,opt,name=communityId,proto3" json:"communityId,omitempty"`
	Color        string   `protobuf:"bytes,5,opt,name=color,proto3" json:"color,omitempty"`
	Details      string   `protobuf:"bytes,6,opt,name=details,proto3" json:"details,omitempty"`
	Name         string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Popularity   int64    `protobuf:"varint,8,opt,name=popularity,proto3" json:"popularity,omitempty"`
	Sex          string   `protobuf:"bytes,9,opt,name=sex,proto3" json:"sex,omitempty"`
	Status       int64    `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	Area         string   `protobuf:"bytes,11,opt,name=area,proto3" json:"area,omitempty"`
	IsSnipped    bool     `protobuf:"varint,12,opt,name=isSnipped,proto3" json:"isSnipped,omitempty"`
	IsSterilized bool     `protobuf:"varint,13,opt,name=isSterilized,proto3" json:"isSterilized,omitempty"`
	Avatars      []string `protobuf:"bytes,14,rep,name=avatars,proto3" json:"avatars,omitempty"`
}

func (x *Cat) Reset() {
	*x = Cat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cat) ProtoMessage() {}

func (x *Cat) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cat.ProtoReflect.Descriptor instead.
func (*Cat) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{0}
}

func (x *Cat) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cat) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Cat) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *Cat) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

func (x *Cat) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Cat) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *Cat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cat) GetPopularity() int64 {
	if x != nil {
		return x.Popularity
	}
	return 0
}

func (x *Cat) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *Cat) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Cat) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *Cat) GetIsSnipped() bool {
	if x != nil {
		return x.IsSnipped
	}
	return false
}

func (x *Cat) GetIsSterilized() bool {
	if x != nil {
		return x.IsSterilized
	}
	return false
}

func (x *Cat) GetAvatars() []string {
	if x != nil {
		return x.Avatars
	}
	return nil
}

type GetManyCatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityId string `protobuf:"bytes,1,opt,name=communityId,proto3" json:"communityId,omitempty"`
	Count       int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"` // 获取的数量
}

func (x *GetManyCatReq) Reset() {
	*x = GetManyCatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManyCatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManyCatReq) ProtoMessage() {}

func (x *GetManyCatReq) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManyCatReq.ProtoReflect.Descriptor instead.
func (*GetManyCatReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{1}
}

func (x *GetManyCatReq) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

func (x *GetManyCatReq) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetManyCatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cats []*Cat `protobuf:"bytes,1,rep,name=cats,proto3" json:"cats,omitempty"`
}

func (x *GetManyCatResp) Reset() {
	*x = GetManyCatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManyCatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManyCatResp) ProtoMessage() {}

func (x *GetManyCatResp) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManyCatResp.ProtoReflect.Descriptor instead.
func (*GetManyCatResp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{2}
}

func (x *GetManyCatResp) GetCats() []*Cat {
	if x != nil {
		return x.Cats
	}
	return nil
}

type GetCatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityId string `protobuf:"bytes,1,opt,name=communityId,proto3" json:"communityId,omitempty"`
}

func (x *GetCatReq) Reset() {
	*x = GetCatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatReq) ProtoMessage() {}

func (x *GetCatReq) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatReq.ProtoReflect.Descriptor instead.
func (*GetCatReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{3}
}

func (x *GetCatReq) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

type GetCatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cat *Cat `protobuf:"bytes,1,opt,name=cat,proto3" json:"cat,omitempty"`
}

func (x *GetCatResp) Reset() {
	*x = GetCatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatResp) ProtoMessage() {}

func (x *GetCatResp) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatResp.ProtoReflect.Descriptor instead.
func (*GetCatResp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{4}
}

func (x *GetCatResp) GetCat() *Cat {
	if x != nil {
		return x.Cat
	}
	return nil
}

type AddCatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cat *Cat `protobuf:"bytes,1,opt,name=cat,proto3" json:"cat,omitempty"`
}

func (x *AddCatReq) Reset() {
	*x = AddCatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCatReq) ProtoMessage() {}

func (x *AddCatReq) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCatReq.ProtoReflect.Descriptor instead.
func (*AddCatReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{5}
}

func (x *AddCatReq) GetCat() *Cat {
	if x != nil {
		return x.Cat
	}
	return nil
}

type AddCatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CatId string `protobuf:"bytes,1,opt,name=catId,proto3" json:"catId,omitempty"`
}

func (x *AddCatResp) Reset() {
	*x = AddCatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCatResp) ProtoMessage() {}

func (x *AddCatResp) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCatResp.ProtoReflect.Descriptor instead.
func (*AddCatResp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{6}
}

func (x *AddCatResp) GetCatId() string {
	if x != nil {
		return x.CatId
	}
	return ""
}

type UpdateCatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cat *Cat `protobuf:"bytes,1,opt,name=cat,proto3" json:"cat,omitempty"`
}

func (x *UpdateCatReq) Reset() {
	*x = UpdateCatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCatReq) ProtoMessage() {}

func (x *UpdateCatReq) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCatReq.ProtoReflect.Descriptor instead.
func (*UpdateCatReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCatReq) GetCat() *Cat {
	if x != nil {
		return x.Cat
	}
	return nil
}

type UpdateCatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCatResp) Reset() {
	*x = UpdateCatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCatResp) ProtoMessage() {}

func (x *UpdateCatResp) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCatResp.ProtoReflect.Descriptor instead.
func (*UpdateCatResp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{8}
}

type DeleteCatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CatId string `protobuf:"bytes,1,opt,name=catId,proto3" json:"catId,omitempty"`
}

func (x *DeleteCatReq) Reset() {
	*x = DeleteCatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCatReq) ProtoMessage() {}

func (x *DeleteCatReq) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCatReq.ProtoReflect.Descriptor instead.
func (*DeleteCatReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCatReq) GetCatId() string {
	if x != nil {
		return x.CatId
	}
	return ""
}

type DeleteCatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCatResp) Reset() {
	*x = DeleteCatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCatResp) ProtoMessage() {}

func (x *DeleteCatResp) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCatResp.ProtoReflect.Descriptor instead.
func (*DeleteCatResp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{10}
}

var File_collection_proto protoreflect.FileDescriptor

var file_collection_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x63, 0x61, 0x74, 0x22, 0xe3, 0x02, 0x0a, 0x03, 0x43, 0x61, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x65,
	0x61, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x73, 0x53, 0x74, 0x65, 0x72, 0x69, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x53, 0x74, 0x65, 0x72, 0x69, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73, 0x22, 0x47, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e,
	0x79, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x04, 0x63, 0x61, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61, 0x74, 0x2e, 0x43, 0x61, 0x74,
	0x52, 0x04, 0x63, 0x61, 0x74, 0x73, 0x22, 0x2d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x03, 0x63, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x63, 0x61, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x03, 0x63, 0x61, 0x74, 0x22,
	0x27, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x03,
	0x63, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61, 0x74, 0x2e,
	0x43, 0x61, 0x74, 0x52, 0x03, 0x63, 0x61, 0x74, 0x22, 0x22, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x03,
	0x63, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61, 0x74, 0x2e,
	0x43, 0x61, 0x74, 0x52, 0x03, 0x63, 0x61, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x24, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x22,
	0x0f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x32, 0xfa, 0x01, 0x0a, 0x03, 0x63, 0x61, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d,
	0x61, 0x6e, 0x79, 0x43, 0x61, 0x74, 0x12, 0x12, 0x2e, 0x63, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x6e, 0x79, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x29, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x12, 0x0e, 0x2e, 0x63, 0x61, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x63, 0x61, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x06, 0x41, 0x64,
	0x64, 0x43, 0x61, 0x74, 0x12, 0x0e, 0x2e, 0x63, 0x61, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x63, 0x61, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x09, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collection_proto_rawDescOnce sync.Once
	file_collection_proto_rawDescData = file_collection_proto_rawDesc
)

func file_collection_proto_rawDescGZIP() []byte {
	file_collection_proto_rawDescOnce.Do(func() {
		file_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_collection_proto_rawDescData)
	})
	return file_collection_proto_rawDescData
}

var file_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_collection_proto_goTypes = []interface{}{
	(*Cat)(nil),            // 0: cat.Cat
	(*GetManyCatReq)(nil),  // 1: cat.GetManyCatReq
	(*GetManyCatResp)(nil), // 2: cat.GetManyCatResp
	(*GetCatReq)(nil),      // 3: cat.GetCatReq
	(*GetCatResp)(nil),     // 4: cat.GetCatResp
	(*AddCatReq)(nil),      // 5: cat.AddCatReq
	(*AddCatResp)(nil),     // 6: cat.AddCatResp
	(*UpdateCatReq)(nil),   // 7: cat.UpdateCatReq
	(*UpdateCatResp)(nil),  // 8: cat.UpdateCatResp
	(*DeleteCatReq)(nil),   // 9: cat.DeleteCatReq
	(*DeleteCatResp)(nil),  // 10: cat.DeleteCatResp
}
var file_collection_proto_depIdxs = []int32{
	0,  // 0: cat.GetManyCatResp.cats:type_name -> cat.Cat
	0,  // 1: cat.GetCatResp.cat:type_name -> cat.Cat
	0,  // 2: cat.AddCatReq.cat:type_name -> cat.Cat
	0,  // 3: cat.UpdateCatReq.cat:type_name -> cat.Cat
	1,  // 4: cat.cat.GetManyCat:input_type -> cat.GetManyCatReq
	3,  // 5: cat.cat.GetCat:input_type -> cat.GetCatReq
	5,  // 6: cat.cat.AddCat:input_type -> cat.AddCatReq
	7,  // 7: cat.cat.UpdateCat:input_type -> cat.UpdateCatReq
	9,  // 8: cat.cat.DeleteCat:input_type -> cat.DeleteCatReq
	2,  // 9: cat.cat.GetManyCat:output_type -> cat.GetManyCatResp
	4,  // 10: cat.cat.GetCat:output_type -> cat.GetCatResp
	6,  // 11: cat.cat.AddCat:output_type -> cat.AddCatResp
	8,  // 12: cat.cat.UpdateCat:output_type -> cat.UpdateCatResp
	10, // 13: cat.cat.DeleteCat:output_type -> cat.DeleteCatResp
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_collection_proto_init() }
func file_collection_proto_init() {
	if File_collection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cat); i {
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
		file_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManyCatReq); i {
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
		file_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManyCatResp); i {
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
		file_collection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatReq); i {
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
		file_collection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatResp); i {
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
		file_collection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCatReq); i {
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
		file_collection_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCatResp); i {
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
		file_collection_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCatReq); i {
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
		file_collection_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCatResp); i {
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
		file_collection_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCatReq); i {
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
		file_collection_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCatResp); i {
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
			RawDescriptor: file_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collection_proto_goTypes,
		DependencyIndexes: file_collection_proto_depIdxs,
		MessageInfos:      file_collection_proto_msgTypes,
	}.Build()
	File_collection_proto = out.File
	file_collection_proto_rawDesc = nil
	file_collection_proto_goTypes = nil
	file_collection_proto_depIdxs = nil
}
