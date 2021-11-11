// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.5.1
// source: sdp.proto

package sdppb

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

type SDPSessionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionName string        `protobuf:"bytes,1,opt,name=actionName,proto3" json:"actionName,omitempty"`
	SdpReq     []*Dictionary `protobuf:"bytes,2,rep,name=sdpReq,proto3" json:"sdpReq,omitempty"`
	Version    string        `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *SDPSessionReq) Reset() {
	*x = SDPSessionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SDPSessionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SDPSessionReq) ProtoMessage() {}

func (x *SDPSessionReq) ProtoReflect() protoreflect.Message {
	mi := &file_sdp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SDPSessionReq.ProtoReflect.Descriptor instead.
func (*SDPSessionReq) Descriptor() ([]byte, []int) {
	return file_sdp_proto_rawDescGZIP(), []int{0}
}

func (x *SDPSessionReq) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *SDPSessionReq) GetSdpReq() []*Dictionary {
	if x != nil {
		return x.SdpReq
	}
	return nil
}

func (x *SDPSessionReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type SDPSessionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status       bool         `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Set          string       `protobuf:"bytes,3,opt,name=set,proto3" json:"set,omitempty"`
	LogMessage   string       `protobuf:"bytes,4,opt,name=logMessage,proto3" json:"logMessage,omitempty"`
	ResponseData *ReqResponse `protobuf:"bytes,5,opt,name=responseData,proto3" json:"responseData,omitempty"`
}

func (x *SDPSessionResp) Reset() {
	*x = SDPSessionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SDPSessionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SDPSessionResp) ProtoMessage() {}

func (x *SDPSessionResp) ProtoReflect() protoreflect.Message {
	mi := &file_sdp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SDPSessionResp.ProtoReflect.Descriptor instead.
func (*SDPSessionResp) Descriptor() ([]byte, []int) {
	return file_sdp_proto_rawDescGZIP(), []int{1}
}

func (x *SDPSessionResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SDPSessionResp) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *SDPSessionResp) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

func (x *SDPSessionResp) GetLogMessage() string {
	if x != nil {
		return x.LogMessage
	}
	return ""
}

func (x *SDPSessionResp) GetResponseData() *ReqResponse {
	if x != nil {
		return x.ResponseData
	}
	return nil
}

type ReqResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ReqResponse) Reset() {
	*x = ReqResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqResponse) ProtoMessage() {}

func (x *ReqResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqResponse.ProtoReflect.Descriptor instead.
func (*ReqResponse) Descriptor() ([]byte, []int) {
	return file_sdp_proto_rawDescGZIP(), []int{2}
}

func (x *ReqResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *ReqResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Dictionary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key             string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value           string            `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ByteValue       []byte            `protobuf:"bytes,3,opt,name=byteValue,proto3" json:"byteValue,omitempty"`
	ListValue       []string          `protobuf:"bytes,4,rep,name=listValue,proto3" json:"listValue,omitempty"`
	DictionaryValue []*ListDictionary `protobuf:"bytes,5,rep,name=dictionaryValue,proto3" json:"dictionaryValue,omitempty"`
}

func (x *Dictionary) Reset() {
	*x = Dictionary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dictionary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dictionary) ProtoMessage() {}

func (x *Dictionary) ProtoReflect() protoreflect.Message {
	mi := &file_sdp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dictionary.ProtoReflect.Descriptor instead.
func (*Dictionary) Descriptor() ([]byte, []int) {
	return file_sdp_proto_rawDescGZIP(), []int{3}
}

func (x *Dictionary) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Dictionary) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Dictionary) GetByteValue() []byte {
	if x != nil {
		return x.ByteValue
	}
	return nil
}

func (x *Dictionary) GetListValue() []string {
	if x != nil {
		return x.ListValue
	}
	return nil
}

func (x *Dictionary) GetDictionaryValue() []*ListDictionary {
	if x != nil {
		return x.DictionaryValue
	}
	return nil
}

type ListDictionary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ListDictionary) Reset() {
	*x = ListDictionary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDictionary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDictionary) ProtoMessage() {}

func (x *ListDictionary) ProtoReflect() protoreflect.Message {
	mi := &file_sdp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDictionary.ProtoReflect.Descriptor instead.
func (*ListDictionary) Descriptor() ([]byte, []int) {
	return file_sdp_proto_rawDescGZIP(), []int{4}
}

func (x *ListDictionary) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ListDictionary) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_sdp_proto protoreflect.FileDescriptor

var file_sdp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x64, 0x70,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22,
	0x81, 0x01, 0x0a, 0x0d, 0x53, 0x44, 0x50, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x64, 0x70, 0x52, 0x65, 0x71, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x64, 0x70, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x06, 0x73, 0x64, 0x70, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0xb9, 0x01, 0x0a, 0x0e, 0x53, 0x44, 0x50, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f,
	0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x73, 0x64, 0x70, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x3d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xbe,
	0x01, 0x0a, 0x0a, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x64, 0x70,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x0f,
	0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x38, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x67, 0x0a, 0x0a, 0x53, 0x44, 0x50,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x53, 0x44, 0x50, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x73, 0x64, 0x70, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x44, 0x50, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x73, 0x64, 0x70, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x44,
	0x50, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x73, 0x64, 0x70, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdp_proto_rawDescOnce sync.Once
	file_sdp_proto_rawDescData = file_sdp_proto_rawDesc
)

func file_sdp_proto_rawDescGZIP() []byte {
	file_sdp_proto_rawDescOnce.Do(func() {
		file_sdp_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdp_proto_rawDescData)
	})
	return file_sdp_proto_rawDescData
}

var file_sdp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sdp_proto_goTypes = []interface{}{
	(*SDPSessionReq)(nil),  // 0: sdpgeneratedclient.SDPSessionReq
	(*SDPSessionResp)(nil), // 1: sdpgeneratedclient.SDPSessionResp
	(*ReqResponse)(nil),    // 2: sdpgeneratedclient.ReqResponse
	(*Dictionary)(nil),     // 3: sdpgeneratedclient.Dictionary
	(*ListDictionary)(nil), // 4: sdpgeneratedclient.ListDictionary
}
var file_sdp_proto_depIdxs = []int32{
	3, // 0: sdpgeneratedclient.SDPSessionReq.sdpReq:type_name -> sdpgeneratedclient.Dictionary
	2, // 1: sdpgeneratedclient.SDPSessionResp.responseData:type_name -> sdpgeneratedclient.ReqResponse
	4, // 2: sdpgeneratedclient.Dictionary.dictionaryValue:type_name -> sdpgeneratedclient.ListDictionary
	0, // 3: sdpgeneratedclient.SDPService.SDPSession:input_type -> sdpgeneratedclient.SDPSessionReq
	1, // 4: sdpgeneratedclient.SDPService.SDPSession:output_type -> sdpgeneratedclient.SDPSessionResp
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sdp_proto_init() }
func file_sdp_proto_init() {
	if File_sdp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SDPSessionReq); i {
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
		file_sdp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SDPSessionResp); i {
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
		file_sdp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqResponse); i {
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
		file_sdp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dictionary); i {
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
		file_sdp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDictionary); i {
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
			RawDescriptor: file_sdp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sdp_proto_goTypes,
		DependencyIndexes: file_sdp_proto_depIdxs,
		MessageInfos:      file_sdp_proto_msgTypes,
	}.Build()
	File_sdp_proto = out.File
	file_sdp_proto_rawDesc = nil
	file_sdp_proto_goTypes = nil
	file_sdp_proto_depIdxs = nil
}
