// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: taucorder/v1/tns.proto

package taucorderv1

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

// Data Structures
type TNSListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node  *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Depth int32 `protobuf:"varint,2,opt,name=depth,proto3" json:"depth,omitempty"`
}

func (x *TNSListRequest) Reset() {
	*x = TNSListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taucorder_v1_tns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TNSListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TNSListRequest) ProtoMessage() {}

func (x *TNSListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taucorder_v1_tns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TNSListRequest.ProtoReflect.Descriptor instead.
func (*TNSListRequest) Descriptor() ([]byte, []int) {
	return file_taucorder_v1_tns_proto_rawDescGZIP(), []int{0}
}

func (x *TNSListRequest) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *TNSListRequest) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

type TNSFetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Path *TNSPath `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *TNSFetchRequest) Reset() {
	*x = TNSFetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taucorder_v1_tns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TNSFetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TNSFetchRequest) ProtoMessage() {}

func (x *TNSFetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taucorder_v1_tns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TNSFetchRequest.ProtoReflect.Descriptor instead.
func (*TNSFetchRequest) Descriptor() ([]byte, []int) {
	return file_taucorder_v1_tns_proto_rawDescGZIP(), []int{1}
}

func (x *TNSFetchRequest) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *TNSFetchRequest) GetPath() *TNSPath {
	if x != nil {
		return x.Path
	}
	return nil
}

type TNSLookupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Types that are assignable to Match:
	//
	//	*TNSLookupRequest_Prefix
	//	*TNSLookupRequest_Regex
	Match isTNSLookupRequest_Match `protobuf_oneof:"match"`
}

func (x *TNSLookupRequest) Reset() {
	*x = TNSLookupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taucorder_v1_tns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TNSLookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TNSLookupRequest) ProtoMessage() {}

func (x *TNSLookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taucorder_v1_tns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TNSLookupRequest.ProtoReflect.Descriptor instead.
func (*TNSLookupRequest) Descriptor() ([]byte, []int) {
	return file_taucorder_v1_tns_proto_rawDescGZIP(), []int{2}
}

func (x *TNSLookupRequest) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (m *TNSLookupRequest) GetMatch() isTNSLookupRequest_Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (x *TNSLookupRequest) GetPrefix() *TNSPath {
	if x, ok := x.GetMatch().(*TNSLookupRequest_Prefix); ok {
		return x.Prefix
	}
	return nil
}

func (x *TNSLookupRequest) GetRegex() *TNSPath {
	if x, ok := x.GetMatch().(*TNSLookupRequest_Regex); ok {
		return x.Regex
	}
	return nil
}

type isTNSLookupRequest_Match interface {
	isTNSLookupRequest_Match()
}

type TNSLookupRequest_Prefix struct {
	Prefix *TNSPath `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type TNSLookupRequest_Regex struct {
	Regex *TNSPath `protobuf:"bytes,3,opt,name=regex,proto3,oneof"`
}

func (*TNSLookupRequest_Prefix) isTNSLookupRequest_Match() {}

func (*TNSLookupRequest_Regex) isTNSLookupRequest_Match() {}

type TNSPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Leafs []string `protobuf:"bytes,1,rep,name=leafs,proto3" json:"leafs,omitempty"`
}

func (x *TNSPath) Reset() {
	*x = TNSPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taucorder_v1_tns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TNSPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TNSPath) ProtoMessage() {}

func (x *TNSPath) ProtoReflect() protoreflect.Message {
	mi := &file_taucorder_v1_tns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TNSPath.ProtoReflect.Descriptor instead.
func (*TNSPath) Descriptor() ([]byte, []int) {
	return file_taucorder_v1_tns_proto_rawDescGZIP(), []int{3}
}

func (x *TNSPath) GetLeafs() []string {
	if x != nil {
		return x.Leafs
	}
	return nil
}

type TNSObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path *TNSPath `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Json string   `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *TNSObject) Reset() {
	*x = TNSObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taucorder_v1_tns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TNSObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TNSObject) ProtoMessage() {}

func (x *TNSObject) ProtoReflect() protoreflect.Message {
	mi := &file_taucorder_v1_tns_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TNSObject.ProtoReflect.Descriptor instead.
func (*TNSObject) Descriptor() ([]byte, []int) {
	return file_taucorder_v1_tns_proto_rawDescGZIP(), []int{4}
}

func (x *TNSObject) GetPath() *TNSPath {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *TNSObject) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

type TNSPaths struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paths []*TNSPath `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (x *TNSPaths) Reset() {
	*x = TNSPaths{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taucorder_v1_tns_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TNSPaths) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TNSPaths) ProtoMessage() {}

func (x *TNSPaths) ProtoReflect() protoreflect.Message {
	mi := &file_taucorder_v1_tns_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TNSPaths.ProtoReflect.Descriptor instead.
func (*TNSPaths) Descriptor() ([]byte, []int) {
	return file_taucorder_v1_tns_proto_rawDescGZIP(), []int{5}
}

func (x *TNSPaths) GetPaths() []*TNSPath {
	if x != nil {
		return x.Paths
	}
	return nil
}

var File_taucorder_v1_tns_proto protoreflect.FileDescriptor

var file_taucorder_v1_tns_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x61, 0x75, 0x63, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x74, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4e, 0x0a, 0x0e, 0x54, 0x4e, 0x53, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x65, 0x70, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74,
	0x68, 0x22, 0x64, 0x0a, 0x0f, 0x54, 0x4e, 0x53, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x75,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4e, 0x53, 0x50, 0x61, 0x74,
	0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xa3, 0x01, 0x0a, 0x10, 0x54, 0x4e, 0x53, 0x4c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04,
	0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x75,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x6e, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4e, 0x53, 0x50, 0x61, 0x74, 0x68, 0x48, 0x00, 0x52, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x2d, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4e, 0x53, 0x50, 0x61, 0x74, 0x68, 0x48, 0x00, 0x52, 0x05, 0x72,
	0x65, 0x67, 0x65, 0x78, 0x42, 0x07, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x22, 0x1f, 0x0a,
	0x07, 0x54, 0x4e, 0x53, 0x50, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x61, 0x66,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x61, 0x66, 0x73, 0x22, 0x4a,
	0x0a, 0x09, 0x54, 0x4e, 0x53, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x75, 0x63,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4e, 0x53, 0x50, 0x61, 0x74, 0x68,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x08, 0x54, 0x4e,
	0x53, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4e, 0x53, 0x50, 0x61, 0x74, 0x68, 0x52, 0x05, 0x70, 0x61,
	0x74, 0x68, 0x73, 0x32, 0xd8, 0x02, 0x0a, 0x0a, 0x54, 0x4e, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x74, 0x61, 0x75,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4e, 0x53, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x61, 0x75, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4e, 0x53, 0x50, 0x61, 0x74, 0x68, 0x30,
	0x01, 0x12, 0x3f, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x74, 0x61, 0x75,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4e, 0x53, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x61, 0x75, 0x63,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4e, 0x53, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x1e, 0x2e, 0x74,
	0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4e, 0x53, 0x4c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74,
	0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4e, 0x53, 0x50,
	0x61, 0x74, 0x68, 0x73, 0x12, 0x4a, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e,
	0x74, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x3c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x74, 0x61, 0x75,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x1c,
	0x2e, 0x74, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x30, 0x01, 0x42, 0xb6,
	0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x42, 0x08, 0x54, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x75, 0x62,
	0x79, 0x74, 0x65, 0x2f, 0x74, 0x61, 0x75, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x75, 0x63,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x74, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x75,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02,
	0x0c, 0x54, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c,
	0x54, 0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x54,
	0x61, 0x75, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x54, 0x61, 0x75, 0x63, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_taucorder_v1_tns_proto_rawDescOnce sync.Once
	file_taucorder_v1_tns_proto_rawDescData = file_taucorder_v1_tns_proto_rawDesc
)

func file_taucorder_v1_tns_proto_rawDescGZIP() []byte {
	file_taucorder_v1_tns_proto_rawDescOnce.Do(func() {
		file_taucorder_v1_tns_proto_rawDescData = protoimpl.X.CompressGZIP(file_taucorder_v1_tns_proto_rawDescData)
	})
	return file_taucorder_v1_tns_proto_rawDescData
}

var file_taucorder_v1_tns_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_taucorder_v1_tns_proto_goTypes = []any{
	(*TNSListRequest)(nil),        // 0: taucorder.v1.TNSListRequest
	(*TNSFetchRequest)(nil),       // 1: taucorder.v1.TNSFetchRequest
	(*TNSLookupRequest)(nil),      // 2: taucorder.v1.TNSLookupRequest
	(*TNSPath)(nil),               // 3: taucorder.v1.TNSPath
	(*TNSObject)(nil),             // 4: taucorder.v1.TNSObject
	(*TNSPaths)(nil),              // 5: taucorder.v1.TNSPaths
	(*Node)(nil),                  // 6: taucorder.v1.Node
	(*ConsensusStateRequest)(nil), // 7: taucorder.v1.ConsensusStateRequest
	(*ConsensusState)(nil),        // 8: taucorder.v1.ConsensusState
}
var file_taucorder_v1_tns_proto_depIdxs = []int32{
	6,  // 0: taucorder.v1.TNSListRequest.node:type_name -> taucorder.v1.Node
	6,  // 1: taucorder.v1.TNSFetchRequest.node:type_name -> taucorder.v1.Node
	3,  // 2: taucorder.v1.TNSFetchRequest.path:type_name -> taucorder.v1.TNSPath
	6,  // 3: taucorder.v1.TNSLookupRequest.node:type_name -> taucorder.v1.Node
	3,  // 4: taucorder.v1.TNSLookupRequest.prefix:type_name -> taucorder.v1.TNSPath
	3,  // 5: taucorder.v1.TNSLookupRequest.regex:type_name -> taucorder.v1.TNSPath
	3,  // 6: taucorder.v1.TNSObject.path:type_name -> taucorder.v1.TNSPath
	3,  // 7: taucorder.v1.TNSPaths.paths:type_name -> taucorder.v1.TNSPath
	0,  // 8: taucorder.v1.TNSService.List:input_type -> taucorder.v1.TNSListRequest
	1,  // 9: taucorder.v1.TNSService.Fetch:input_type -> taucorder.v1.TNSFetchRequest
	2,  // 10: taucorder.v1.TNSService.Lookup:input_type -> taucorder.v1.TNSLookupRequest
	7,  // 11: taucorder.v1.TNSService.State:input_type -> taucorder.v1.ConsensusStateRequest
	6,  // 12: taucorder.v1.TNSService.States:input_type -> taucorder.v1.Node
	3,  // 13: taucorder.v1.TNSService.List:output_type -> taucorder.v1.TNSPath
	4,  // 14: taucorder.v1.TNSService.Fetch:output_type -> taucorder.v1.TNSObject
	5,  // 15: taucorder.v1.TNSService.Lookup:output_type -> taucorder.v1.TNSPaths
	8,  // 16: taucorder.v1.TNSService.State:output_type -> taucorder.v1.ConsensusState
	8,  // 17: taucorder.v1.TNSService.States:output_type -> taucorder.v1.ConsensusState
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_taucorder_v1_tns_proto_init() }
func file_taucorder_v1_tns_proto_init() {
	if File_taucorder_v1_tns_proto != nil {
		return
	}
	file_taucorder_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_taucorder_v1_tns_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TNSListRequest); i {
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
		file_taucorder_v1_tns_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TNSFetchRequest); i {
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
		file_taucorder_v1_tns_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TNSLookupRequest); i {
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
		file_taucorder_v1_tns_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TNSPath); i {
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
		file_taucorder_v1_tns_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TNSObject); i {
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
		file_taucorder_v1_tns_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TNSPaths); i {
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
	file_taucorder_v1_tns_proto_msgTypes[2].OneofWrappers = []any{
		(*TNSLookupRequest_Prefix)(nil),
		(*TNSLookupRequest_Regex)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_taucorder_v1_tns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_taucorder_v1_tns_proto_goTypes,
		DependencyIndexes: file_taucorder_v1_tns_proto_depIdxs,
		MessageInfos:      file_taucorder_v1_tns_proto_msgTypes,
	}.Build()
	File_taucorder_v1_tns_proto = out.File
	file_taucorder_v1_tns_proto_rawDesc = nil
	file_taucorder_v1_tns_proto_goTypes = nil
	file_taucorder_v1_tns_proto_depIdxs = nil
}
