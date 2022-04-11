// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: objectstore.proto

package pluginextensionv2

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

type EntityReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"` // dashboard | alert | datasource | panel (library) | ...
	Grn  string `protobuf:"bytes,2,opt,name=grn,proto3" json:"grn,omitempty"`   // includes folder pattern -- may include git branch info
}

func (x *EntityReference) Reset() {
	*x = EntityReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityReference) ProtoMessage() {}

func (x *EntityReference) ProtoReflect() protoreflect.Message {
	mi := &file_objectstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityReference.ProtoReflect.Descriptor instead.
func (*EntityReference) Descriptor() ([]byte, []int) {
	return file_objectstore_proto_rawDescGZIP(), []int{0}
}

func (x *EntityReference) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *EntityReference) GetGrn() string {
	if x != nil {
		return x.Grn
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref *EntityReference `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_objectstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_objectstore_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetRef() *EntityReference {
	if x != nil {
		return x.Ref
	}
	return nil
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref *EntityReference `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// lineage
	Created     int64  `protobuf:"varint,9,opt,name=created,proto3" json:"created,omitempty"`
	Updated     int64  `protobuf:"varint,10,opt,name=updated,proto3" json:"updated,omitempty"`        //  repeated string author = 11; // everyone who touched it?
	ContentType string `protobuf:"bytes,12,opt,name=contentType,proto3" json:"contentType,omitempty"` // json | csv | image/png | ...
	Body        []byte `protobuf:"bytes,16,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_objectstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_objectstore_proto_rawDescGZIP(), []int{2}
}

func (x *Object) GetRef() *EntityReference {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *Object) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Object) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Object) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Object) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Object *Object `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_objectstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_objectstore_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetResponse) GetObject() *Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  PluginContext pluginContext = 1;
	Ref         *EntityReference `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	ContentType string           `protobuf:"bytes,7,opt,name=contentType,proto3" json:"contentType,omitempty"` // json | csv | image/png | ...
	Body        []byte           `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_objectstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_objectstore_proto_rawDescGZIP(), []int{4}
}

func (x *WriteRequest) GetRef() *EntityReference {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *WriteRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *WriteRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type WriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_objectstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_objectstore_proto_rawDescGZIP(), []int{5}
}

func (x *WriteResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  PluginContext pluginContext = 1;
	Ref *EntityReference `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectstore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_objectstore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_objectstore_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequest) GetRef() *EntityReference {
	if x != nil {
		return x.Ref
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectstore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_objectstore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_objectstore_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type WatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  PluginContext pluginContext = 1;
	KindFilter string `protobuf:"bytes,2,opt,name=kind_filter,json=kindFilter,proto3" json:"kind_filter,omitempty"` // filter to a specific kind of entity (ie, only dashboards).
	GrnFilter  string `protobuf:"bytes,3,opt,name=grn_filter,json=grnFilter,proto3" json:"grn_filter,omitempty"`    // the entity ID or folder prefix
	Recursive  bool   `protobuf:"varint,4,opt,name=recursive,proto3" json:"recursive,omitempty"`                    // when true, all changes to this + any substrings will be sent to the caller
	Details    bool   `protobuf:"varint,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectstore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_objectstore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_objectstore_proto_rawDescGZIP(), []int{8}
}

func (x *WatchRequest) GetKindFilter() string {
	if x != nil {
		return x.KindFilter
	}
	return ""
}

func (x *WatchRequest) GetGrnFilter() string {
	if x != nil {
		return x.GrnFilter
	}
	return ""
}

func (x *WatchRequest) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

func (x *WatchRequest) GetDetails() bool {
	if x != nil {
		return x.Details
	}
	return false
}

type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64            `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Ref       *EntityReference `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	Action    string           `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"` // CREATE + UPDATE + DELETE, TODO: enum.
	// when details = true, send the whole object.
	Object *Object `protobuf:"bytes,4,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectstore_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_objectstore_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_objectstore_proto_rawDescGZIP(), []int{9}
}

func (x *WatchResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *WatchResponse) GetRef() *EntityReference {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *WatchResponse) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *WatchResponse) GetObject() *Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KindFilter string `protobuf:"bytes,2,opt,name=kind_filter,json=kindFilter,proto3" json:"kind_filter,omitempty"` // filter to a specific kind of entity (ie, only dashboards).
	GrnFilter  string `protobuf:"bytes,3,opt,name=grn_filter,json=grnFilter,proto3" json:"grn_filter,omitempty"`    // the entity ID or folder prefix
	Recursive  bool   `protobuf:"varint,4,opt,name=recursive,proto3" json:"recursive,omitempty"`                    // when true, all changes to this + any substrings will be sent to the caller
	Details    bool   `protobuf:"varint,5,opt,name=details,proto3" json:"details,omitempty"`
	Limit      uint32 `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	AfterGrn   string `protobuf:"bytes,7,opt,name=after_grn,json=afterGrn,proto3" json:"after_grn,omitempty"`
	BeforeGrn  string `protobuf:"bytes,8,opt,name=before_grn,json=beforeGrn,proto3" json:"before_grn,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectstore_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_objectstore_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_objectstore_proto_rawDescGZIP(), []int{10}
}

func (x *ListRequest) GetKindFilter() string {
	if x != nil {
		return x.KindFilter
	}
	return ""
}

func (x *ListRequest) GetGrnFilter() string {
	if x != nil {
		return x.GrnFilter
	}
	return ""
}

func (x *ListRequest) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

func (x *ListRequest) GetDetails() bool {
	if x != nil {
		return x.Details
	}
	return false
}

func (x *ListRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRequest) GetAfterGrn() string {
	if x != nil {
		return x.AfterGrn
	}
	return ""
}

func (x *ListRequest) GetBeforeGrn() string {
	if x != nil {
		return x.BeforeGrn
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*Object `protobuf:"bytes,2,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectstore_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_objectstore_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_objectstore_proto_rawDescGZIP(), []int{11}
}

func (x *ListResponse) GetObjects() []*Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

var File_objectstore_proto protoreflect.FileDescriptor

var file_objectstore_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x22, 0x37, 0x0a, 0x0f, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x67, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x72, 0x6e, 0x22,
	0x42, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a,
	0x03, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x03,
	0x72, 0x65, 0x66, 0x22, 0xa8, 0x01, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x34,
	0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x03, 0x72, 0x65, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x54,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0x7a, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x23, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x45, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0x24, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x69, 0x6e, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x69, 0x6e, 0x64, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x0d,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x34, 0x0a, 0x03, 0x72,
	0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x03, 0x72, 0x65,
	0x66, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0xd7, 0x01, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x6b, 0x69, 0x6e, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6b, 0x69, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x67, 0x72, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x67, 0x72, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x47, 0x72, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x5f, 0x67, 0x72, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x47, 0x72, 0x6e, 0x22, 0x43, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x32, 0x9b, 0x03, 0x0a, 0x05,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x1d, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x50, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x76, 0x32, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x76, 0x32, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x53, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x20, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f, 0x3b,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x76,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_objectstore_proto_rawDescOnce sync.Once
	file_objectstore_proto_rawDescData = file_objectstore_proto_rawDesc
)

func file_objectstore_proto_rawDescGZIP() []byte {
	file_objectstore_proto_rawDescOnce.Do(func() {
		file_objectstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_objectstore_proto_rawDescData)
	})
	return file_objectstore_proto_rawDescData
}

var file_objectstore_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_objectstore_proto_goTypes = []interface{}{
	(*EntityReference)(nil), // 0: pluginextensionv2.EntityReference
	(*GetRequest)(nil),      // 1: pluginextensionv2.GetRequest
	(*Object)(nil),          // 2: pluginextensionv2.Object
	(*GetResponse)(nil),     // 3: pluginextensionv2.GetResponse
	(*WriteRequest)(nil),    // 4: pluginextensionv2.WriteRequest
	(*WriteResponse)(nil),   // 5: pluginextensionv2.WriteResponse
	(*DeleteRequest)(nil),   // 6: pluginextensionv2.DeleteRequest
	(*DeleteResponse)(nil),  // 7: pluginextensionv2.DeleteResponse
	(*WatchRequest)(nil),    // 8: pluginextensionv2.WatchRequest
	(*WatchResponse)(nil),   // 9: pluginextensionv2.WatchResponse
	(*ListRequest)(nil),     // 10: pluginextensionv2.ListRequest
	(*ListResponse)(nil),    // 11: pluginextensionv2.ListResponse
}
var file_objectstore_proto_depIdxs = []int32{
	0,  // 0: pluginextensionv2.GetRequest.ref:type_name -> pluginextensionv2.EntityReference
	0,  // 1: pluginextensionv2.Object.ref:type_name -> pluginextensionv2.EntityReference
	2,  // 2: pluginextensionv2.GetResponse.object:type_name -> pluginextensionv2.Object
	0,  // 3: pluginextensionv2.WriteRequest.ref:type_name -> pluginextensionv2.EntityReference
	0,  // 4: pluginextensionv2.DeleteRequest.ref:type_name -> pluginextensionv2.EntityReference
	0,  // 5: pluginextensionv2.WatchResponse.ref:type_name -> pluginextensionv2.EntityReference
	2,  // 6: pluginextensionv2.WatchResponse.object:type_name -> pluginextensionv2.Object
	2,  // 7: pluginextensionv2.ListResponse.objects:type_name -> pluginextensionv2.Object
	10, // 8: pluginextensionv2.Store.ListStore:input_type -> pluginextensionv2.ListRequest
	1,  // 9: pluginextensionv2.Store.GetEntity:input_type -> pluginextensionv2.GetRequest
	4,  // 10: pluginextensionv2.Store.WriteEntity:input_type -> pluginextensionv2.WriteRequest
	6,  // 11: pluginextensionv2.Store.DeleteEntity:input_type -> pluginextensionv2.DeleteRequest
	8,  // 12: pluginextensionv2.Store.WatchStore:input_type -> pluginextensionv2.WatchRequest
	11, // 13: pluginextensionv2.Store.ListStore:output_type -> pluginextensionv2.ListResponse
	3,  // 14: pluginextensionv2.Store.GetEntity:output_type -> pluginextensionv2.GetResponse
	5,  // 15: pluginextensionv2.Store.WriteEntity:output_type -> pluginextensionv2.WriteResponse
	7,  // 16: pluginextensionv2.Store.DeleteEntity:output_type -> pluginextensionv2.DeleteResponse
	9,  // 17: pluginextensionv2.Store.WatchStore:output_type -> pluginextensionv2.WatchResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_objectstore_proto_init() }
func file_objectstore_proto_init() {
	if File_objectstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_objectstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityReference); i {
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
		file_objectstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_objectstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_objectstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_objectstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_objectstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResponse); i {
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
		file_objectstore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_objectstore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_objectstore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRequest); i {
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
		file_objectstore_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResponse); i {
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
		file_objectstore_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_objectstore_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_objectstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_objectstore_proto_goTypes,
		DependencyIndexes: file_objectstore_proto_depIdxs,
		MessageInfos:      file_objectstore_proto_msgTypes,
	}.Build()
	File_objectstore_proto = out.File
	file_objectstore_proto_rawDesc = nil
	file_objectstore_proto_goTypes = nil
	file_objectstore_proto_depIdxs = nil
}