// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Open Networking Foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: pfcpsim.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// count represents the number of session
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// baseID is used to create incremental IDs for PDRs, FARs, QERs
	BaseID        int32  `protobuf:"varint,2,opt,name=baseID,proto3" json:"baseID,omitempty"`
	NodeBAddress  string `protobuf:"bytes,3,opt,name=nodeBAddress,proto3" json:"nodeBAddress,omitempty"`
	UeAddressPool string `protobuf:"bytes,4,opt,name=ueAddressPool,proto3" json:"ueAddressPool,omitempty"`
	SdfFilter     string `protobuf:"bytes,5,opt,name=sdfFilter,proto3" json:"sdfFilter,omitempty"`
	Qfi           int32  `protobuf:"varint,6,opt,name=qfi,proto3" json:"qfi,omitempty"` // Should be uint8
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pfcpsim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pfcpsim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_pfcpsim_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSessionRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CreateSessionRequest) GetBaseID() int32 {
	if x != nil {
		return x.BaseID
	}
	return 0
}

func (x *CreateSessionRequest) GetNodeBAddress() string {
	if x != nil {
		return x.NodeBAddress
	}
	return ""
}

func (x *CreateSessionRequest) GetUeAddressPool() string {
	if x != nil {
		return x.UeAddressPool
	}
	return ""
}

func (x *CreateSessionRequest) GetSdfFilter() string {
	if x != nil {
		return x.SdfFilter
	}
	return ""
}

func (x *CreateSessionRequest) GetQfi() int32 {
	if x != nil {
		return x.Qfi
	}
	return 0
}

type ModifySessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// count represents the number of session
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// baseID is used to create incremental IDs for PDRs, FARs, QERs
	BaseID        int32  `protobuf:"varint,2,opt,name=baseID,proto3" json:"baseID,omitempty"`
	NodeBAddress  string `protobuf:"bytes,3,opt,name=nodeBAddress,proto3" json:"nodeBAddress,omitempty"`
	UeAddressPool string `protobuf:"bytes,4,opt,name=ueAddressPool,proto3" json:"ueAddressPool,omitempty"`
	BufferFlag    bool   `protobuf:"varint,5,opt,name=bufferFlag,proto3" json:"bufferFlag,omitempty"`
	NotifyCPFlag  bool   `protobuf:"varint,6,opt,name=notifyCPFlag,proto3" json:"notifyCPFlag,omitempty"`
}

func (x *ModifySessionRequest) Reset() {
	*x = ModifySessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pfcpsim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySessionRequest) ProtoMessage() {}

func (x *ModifySessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pfcpsim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySessionRequest.ProtoReflect.Descriptor instead.
func (*ModifySessionRequest) Descriptor() ([]byte, []int) {
	return file_pfcpsim_proto_rawDescGZIP(), []int{1}
}

func (x *ModifySessionRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ModifySessionRequest) GetBaseID() int32 {
	if x != nil {
		return x.BaseID
	}
	return 0
}

func (x *ModifySessionRequest) GetNodeBAddress() string {
	if x != nil {
		return x.NodeBAddress
	}
	return ""
}

func (x *ModifySessionRequest) GetUeAddressPool() string {
	if x != nil {
		return x.UeAddressPool
	}
	return ""
}

func (x *ModifySessionRequest) GetBufferFlag() bool {
	if x != nil {
		return x.BufferFlag
	}
	return false
}

func (x *ModifySessionRequest) GetNotifyCPFlag() bool {
	if x != nil {
		return x.NotifyCPFlag
	}
	return false
}

type ConfigureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the data-plane interface between UPF and gNodeB
	UpfN3Address string `protobuf:"bytes,1,opt,name=upfN3Address,proto3" json:"upfN3Address,omitempty"`
	// the PFCP agent server address
	RemotePeerAddress string `protobuf:"bytes,3,opt,name=remotePeerAddress,proto3" json:"remotePeerAddress,omitempty"`
}

func (x *ConfigureRequest) Reset() {
	*x = ConfigureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pfcpsim_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureRequest) ProtoMessage() {}

func (x *ConfigureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pfcpsim_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureRequest.ProtoReflect.Descriptor instead.
func (*ConfigureRequest) Descriptor() ([]byte, []int) {
	return file_pfcpsim_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigureRequest) GetUpfN3Address() string {
	if x != nil {
		return x.UpfN3Address
	}
	return ""
}

func (x *ConfigureRequest) GetRemotePeerAddress() string {
	if x != nil {
		return x.RemotePeerAddress
	}
	return ""
}

type DeleteSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// baseID is used to decide where to start deleting sessions
	BaseID int32 `protobuf:"varint,2,opt,name=baseID,proto3" json:"baseID,omitempty"`
}

func (x *DeleteSessionRequest) Reset() {
	*x = DeleteSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pfcpsim_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSessionRequest) ProtoMessage() {}

func (x *DeleteSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pfcpsim_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSessionRequest.ProtoReflect.Descriptor instead.
func (*DeleteSessionRequest) Descriptor() ([]byte, []int) {
	return file_pfcpsim_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteSessionRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *DeleteSessionRequest) GetBaseID() int32 {
	if x != nil {
		return x.BaseID
	}
	return 0
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pfcpsim_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pfcpsim_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_pfcpsim_proto_rawDescGZIP(), []int{4}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pfcpsim_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pfcpsim_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pfcpsim_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pfcpsim_proto protoreflect.FileDescriptor

var file_pfcpsim_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x66, 0x63, 0x70, 0x73, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0xbe, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x61, 0x73, 0x65, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x6e,
	0x6f, 0x64, 0x65, 0x42, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x42, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x24, 0x0a, 0x0d, 0x75, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6f, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x64, 0x66, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x64, 0x66, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x66, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x71, 0x66, 0x69, 0x22, 0xd2, 0x01, 0x0a, 0x14, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x61, 0x73, 0x65, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c,
	0x6e, 0x6f, 0x64, 0x65, 0x42, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x42, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x24, 0x0a, 0x0d, 0x75, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6f,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x46, 0x6c, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x43, 0x50, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x43, 0x50, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x64, 0x0a, 0x10, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x75, 0x70, 0x66, 0x4e, 0x33, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x70, 0x66, 0x4e, 0x33, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x44, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x62, 0x61, 0x73, 0x65, 0x49, 0x44, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xda, 0x02,
	0x0a, 0x07, 0x50, 0x46, 0x43, 0x50, 0x53, 0x69, 0x6d, 0x12, 0x33, 0x0a, 0x09, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f,
	0x0a, 0x09, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x32, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x12,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pfcpsim_proto_rawDescOnce sync.Once
	file_pfcpsim_proto_rawDescData = file_pfcpsim_proto_rawDesc
)

func file_pfcpsim_proto_rawDescGZIP() []byte {
	file_pfcpsim_proto_rawDescOnce.Do(func() {
		file_pfcpsim_proto_rawDescData = protoimpl.X.CompressGZIP(file_pfcpsim_proto_rawDescData)
	})
	return file_pfcpsim_proto_rawDescData
}

var file_pfcpsim_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pfcpsim_proto_goTypes = []interface{}{
	(*CreateSessionRequest)(nil), // 0: api.CreateSessionRequest
	(*ModifySessionRequest)(nil), // 1: api.ModifySessionRequest
	(*ConfigureRequest)(nil),     // 2: api.ConfigureRequest
	(*DeleteSessionRequest)(nil), // 3: api.DeleteSessionRequest
	(*EmptyRequest)(nil),         // 4: api.EmptyRequest
	(*Response)(nil),             // 5: api.Response
}
var file_pfcpsim_proto_depIdxs = []int32{
	2, // 0: api.PFCPSim.Configure:input_type -> api.ConfigureRequest
	4, // 1: api.PFCPSim.Associate:input_type -> api.EmptyRequest
	4, // 2: api.PFCPSim.Disassociate:input_type -> api.EmptyRequest
	0, // 3: api.PFCPSim.CreateSession:input_type -> api.CreateSessionRequest
	1, // 4: api.PFCPSim.ModifySession:input_type -> api.ModifySessionRequest
	3, // 5: api.PFCPSim.DeleteSession:input_type -> api.DeleteSessionRequest
	5, // 6: api.PFCPSim.Configure:output_type -> api.Response
	5, // 7: api.PFCPSim.Associate:output_type -> api.Response
	5, // 8: api.PFCPSim.Disassociate:output_type -> api.Response
	5, // 9: api.PFCPSim.CreateSession:output_type -> api.Response
	5, // 10: api.PFCPSim.ModifySession:output_type -> api.Response
	5, // 11: api.PFCPSim.DeleteSession:output_type -> api.Response
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pfcpsim_proto_init() }
func file_pfcpsim_proto_init() {
	if File_pfcpsim_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pfcpsim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequest); i {
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
		file_pfcpsim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifySessionRequest); i {
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
		file_pfcpsim_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureRequest); i {
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
		file_pfcpsim_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSessionRequest); i {
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
		file_pfcpsim_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_pfcpsim_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_pfcpsim_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pfcpsim_proto_goTypes,
		DependencyIndexes: file_pfcpsim_proto_depIdxs,
		MessageInfos:      file_pfcpsim_proto_msgTypes,
	}.Build()
	File_pfcpsim_proto = out.File
	file_pfcpsim_proto_rawDesc = nil
	file_pfcpsim_proto_goTypes = nil
	file_pfcpsim_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PFCPSimClient is the client API for PFCPSim service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PFCPSimClient interface {
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*Response, error)
	// Associate connects PFCPClient to remote peer and starts an association
	Associate(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*Response, error)
	// Disassociate perform teardown of association and disconnects from remote peer.
	Disassociate(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*Response, error)
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Response, error)
	ModifySession(ctx context.Context, in *ModifySessionRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*Response, error)
}

type pFCPSimClient struct {
	cc grpc.ClientConnInterface
}

func NewPFCPSimClient(cc grpc.ClientConnInterface) PFCPSimClient {
	return &pFCPSimClient{cc}
}

func (c *pFCPSimClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.PFCPSim/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pFCPSimClient) Associate(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.PFCPSim/Associate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pFCPSimClient) Disassociate(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.PFCPSim/Disassociate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pFCPSimClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.PFCPSim/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pFCPSimClient) ModifySession(ctx context.Context, in *ModifySessionRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.PFCPSim/ModifySession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pFCPSimClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.PFCPSim/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PFCPSimServer is the server API for PFCPSim service.
type PFCPSimServer interface {
	Configure(context.Context, *ConfigureRequest) (*Response, error)
	// Associate connects PFCPClient to remote peer and starts an association
	Associate(context.Context, *EmptyRequest) (*Response, error)
	// Disassociate perform teardown of association and disconnects from remote peer.
	Disassociate(context.Context, *EmptyRequest) (*Response, error)
	CreateSession(context.Context, *CreateSessionRequest) (*Response, error)
	ModifySession(context.Context, *ModifySessionRequest) (*Response, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*Response, error)
}

// UnimplementedPFCPSimServer can be embedded to have forward compatible implementations.
type UnimplementedPFCPSimServer struct {
}

func (*UnimplementedPFCPSimServer) Configure(context.Context, *ConfigureRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedPFCPSimServer) Associate(context.Context, *EmptyRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Associate not implemented")
}
func (*UnimplementedPFCPSimServer) Disassociate(context.Context, *EmptyRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disassociate not implemented")
}
func (*UnimplementedPFCPSimServer) CreateSession(context.Context, *CreateSessionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (*UnimplementedPFCPSimServer) ModifySession(context.Context, *ModifySessionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySession not implemented")
}
func (*UnimplementedPFCPSimServer) DeleteSession(context.Context, *DeleteSessionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}

func RegisterPFCPSimServer(s *grpc.Server, srv PFCPSimServer) {
	s.RegisterService(&_PFCPSim_serviceDesc, srv)
}

func _PFCPSim_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PFCPSimServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PFCPSim/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PFCPSimServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PFCPSim_Associate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PFCPSimServer).Associate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PFCPSim/Associate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PFCPSimServer).Associate(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PFCPSim_Disassociate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PFCPSimServer).Disassociate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PFCPSim/Disassociate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PFCPSimServer).Disassociate(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PFCPSim_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PFCPSimServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PFCPSim/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PFCPSimServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PFCPSim_ModifySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PFCPSimServer).ModifySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PFCPSim/ModifySession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PFCPSimServer).ModifySession(ctx, req.(*ModifySessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PFCPSim_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PFCPSimServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PFCPSim/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PFCPSimServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PFCPSim_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.PFCPSim",
	HandlerType: (*PFCPSimServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _PFCPSim_Configure_Handler,
		},
		{
			MethodName: "Associate",
			Handler:    _PFCPSim_Associate_Handler,
		},
		{
			MethodName: "Disassociate",
			Handler:    _PFCPSim_Disassociate_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _PFCPSim_CreateSession_Handler,
		},
		{
			MethodName: "ModifySession",
			Handler:    _PFCPSim_ModifySession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _PFCPSim_DeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pfcpsim.proto",
}
