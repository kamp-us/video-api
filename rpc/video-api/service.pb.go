// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: rpc/video-api/service.proto

package video_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetVideoRequest) Reset() {
	*x = GetVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_video_api_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoRequest) ProtoMessage() {}

func (x *GetVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_video_api_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoRequest.ProtoReflect.Descriptor instead.
func (*GetVideoRequest) Descriptor() ([]byte, []int) {
	return file_rpc_video_api_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetVideoRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type CreateVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	URI    string `protobuf:"bytes,3,opt,name=URI,proto3" json:"URI,omitempty"`
}

func (x *CreateVideoRequest) Reset() {
	*x = CreateVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_video_api_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVideoRequest) ProtoMessage() {}

func (x *CreateVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_video_api_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVideoRequest.ProtoReflect.Descriptor instead.
func (*CreateVideoRequest) Descriptor() ([]byte, []int) {
	return file_rpc_video_api_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVideoRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateVideoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateVideoRequest) GetURI() string {
	if x != nil {
		return x.URI
	}
	return ""
}

type UpdateVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string                  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	URI  *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=URI,proto3" json:"URI,omitempty"`
}

func (x *UpdateVideoRequest) Reset() {
	*x = UpdateVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_video_api_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVideoRequest) ProtoMessage() {}

func (x *UpdateVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_video_api_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVideoRequest.ProtoReflect.Descriptor instead.
func (*UpdateVideoRequest) Descriptor() ([]byte, []int) {
	return file_rpc_video_api_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateVideoRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateVideoRequest) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *UpdateVideoRequest) GetURI() *wrapperspb.StringValue {
	if x != nil {
		return x.URI
	}
	return nil
}

type DeleteVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteVideoRequest) Reset() {
	*x = DeleteVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_video_api_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVideoRequest) ProtoMessage() {}

func (x *DeleteVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_video_api_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVideoRequest.ProtoReflect.Descriptor instead.
func (*DeleteVideoRequest) Descriptor() ([]byte, []int) {
	return file_rpc_video_api_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteVideoRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetBatchVideosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetBatchVideosRequest) Reset() {
	*x = GetBatchVideosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_video_api_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchVideosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchVideosRequest) ProtoMessage() {}

func (x *GetBatchVideosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_video_api_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchVideosRequest.ProtoReflect.Descriptor instead.
func (*GetBatchVideosRequest) Descriptor() ([]byte, []int) {
	return file_rpc_video_api_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetBatchVideosRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetBatchVideosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos []*Video `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *GetBatchVideosResponse) Reset() {
	*x = GetBatchVideosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_video_api_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchVideosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchVideosResponse) ProtoMessage() {}

func (x *GetBatchVideosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_video_api_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchVideosResponse.ProtoReflect.Descriptor instead.
func (*GetBatchVideosResponse) Descriptor() ([]byte, []int) {
	return file_rpc_video_api_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetBatchVideosResponse) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	URI    string `protobuf:"bytes,4,opt,name=URI,proto3" json:"URI,omitempty"`
	Slug   string `protobuf:"bytes,5,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_video_api_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_video_api_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_rpc_video_api_service_proto_rawDescGZIP(), []int{6}
}

func (x *Video) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Video) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Video) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Video) GetURI() string {
	if x != nil {
		return x.URI
	}
	return ""
}

func (x *Video) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

var File_rpc_video_api_service_proto protoreflect.FileDescriptor

var file_rpc_video_api_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6b,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x61, 0x70, 0x69, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x53,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x49, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x55, 0x52, 0x49, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x03,
	0x55, 0x52, 0x49, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x55, 0x52, 0x49, 0x22, 0x24, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x29, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x48, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x22, 0x6a, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x52, 0x49, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x49, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x32, 0x97, 0x03, 0x0a, 0x08, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x50, 0x49,
	0x12, 0x44, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x20, 0x2e, 0x6b,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x61, 0x70, 0x69,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x4a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x23, 0x2e, 0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6b, 0x61, 0x6d,
	0x70, 0x75, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x12, 0x4a, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x23, 0x2e, 0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x23, 0x2e,
	0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x61, 0x70, 0x69, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x61, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x26, 0x2e, 0x6b,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a,
	0x0e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_video_api_service_proto_rawDescOnce sync.Once
	file_rpc_video_api_service_proto_rawDescData = file_rpc_video_api_service_proto_rawDesc
)

func file_rpc_video_api_service_proto_rawDescGZIP() []byte {
	file_rpc_video_api_service_proto_rawDescOnce.Do(func() {
		file_rpc_video_api_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_video_api_service_proto_rawDescData)
	})
	return file_rpc_video_api_service_proto_rawDescData
}

var file_rpc_video_api_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_video_api_service_proto_goTypes = []interface{}{
	(*GetVideoRequest)(nil),        // 0: kampus.videoapi.GetVideoRequest
	(*CreateVideoRequest)(nil),     // 1: kampus.videoapi.CreateVideoRequest
	(*UpdateVideoRequest)(nil),     // 2: kampus.videoapi.UpdateVideoRequest
	(*DeleteVideoRequest)(nil),     // 3: kampus.videoapi.DeleteVideoRequest
	(*GetBatchVideosRequest)(nil),  // 4: kampus.videoapi.GetBatchVideosRequest
	(*GetBatchVideosResponse)(nil), // 5: kampus.videoapi.GetBatchVideosResponse
	(*Video)(nil),                  // 6: kampus.videoapi.Video
	(*wrapperspb.StringValue)(nil), // 7: google.protobuf.StringValue
	(*emptypb.Empty)(nil),          // 8: google.protobuf.Empty
}
var file_rpc_video_api_service_proto_depIdxs = []int32{
	7, // 0: kampus.videoapi.UpdateVideoRequest.name:type_name -> google.protobuf.StringValue
	7, // 1: kampus.videoapi.UpdateVideoRequest.URI:type_name -> google.protobuf.StringValue
	6, // 2: kampus.videoapi.GetBatchVideosResponse.videos:type_name -> kampus.videoapi.Video
	0, // 3: kampus.videoapi.VideoAPI.GetVideo:input_type -> kampus.videoapi.GetVideoRequest
	1, // 4: kampus.videoapi.VideoAPI.CreateVideo:input_type -> kampus.videoapi.CreateVideoRequest
	2, // 5: kampus.videoapi.VideoAPI.UpdateVideo:input_type -> kampus.videoapi.UpdateVideoRequest
	3, // 6: kampus.videoapi.VideoAPI.DeleteVideo:input_type -> kampus.videoapi.DeleteVideoRequest
	4, // 7: kampus.videoapi.VideoAPI.GetBatchVideos:input_type -> kampus.videoapi.GetBatchVideosRequest
	6, // 8: kampus.videoapi.VideoAPI.GetVideo:output_type -> kampus.videoapi.Video
	6, // 9: kampus.videoapi.VideoAPI.CreateVideo:output_type -> kampus.videoapi.Video
	8, // 10: kampus.videoapi.VideoAPI.UpdateVideo:output_type -> google.protobuf.Empty
	8, // 11: kampus.videoapi.VideoAPI.DeleteVideo:output_type -> google.protobuf.Empty
	5, // 12: kampus.videoapi.VideoAPI.GetBatchVideos:output_type -> kampus.videoapi.GetBatchVideosResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_video_api_service_proto_init() }
func file_rpc_video_api_service_proto_init() {
	if File_rpc_video_api_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_video_api_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoRequest); i {
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
		file_rpc_video_api_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVideoRequest); i {
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
		file_rpc_video_api_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVideoRequest); i {
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
		file_rpc_video_api_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVideoRequest); i {
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
		file_rpc_video_api_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchVideosRequest); i {
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
		file_rpc_video_api_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchVideosResponse); i {
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
		file_rpc_video_api_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
			RawDescriptor: file_rpc_video_api_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_video_api_service_proto_goTypes,
		DependencyIndexes: file_rpc_video_api_service_proto_depIdxs,
		MessageInfos:      file_rpc_video_api_service_proto_msgTypes,
	}.Build()
	File_rpc_video_api_service_proto = out.File
	file_rpc_video_api_service_proto_rawDesc = nil
	file_rpc_video_api_service_proto_goTypes = nil
	file_rpc_video_api_service_proto_depIdxs = nil
}
