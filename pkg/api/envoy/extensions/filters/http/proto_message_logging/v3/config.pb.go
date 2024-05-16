// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.4
// source: envoy/extensions/filters/http/proto_message_logging/v3/config.proto

package proto_message_loggingv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v3 "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/config/core/v3"
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

type ProtoMessageLoggingConfig_LogMode int32

const (
	ProtoMessageLoggingConfig_LogMode_UNSPECIFIED ProtoMessageLoggingConfig_LogMode = 0
	// The filter will log the first and the last message for
	// for streaming cases, containing
	// client-side streaming, server-side streaming or bi-directional streaming.
	ProtoMessageLoggingConfig_FIRST_AND_LAST ProtoMessageLoggingConfig_LogMode = 1
)

// Enum value maps for ProtoMessageLoggingConfig_LogMode.
var (
	ProtoMessageLoggingConfig_LogMode_name = map[int32]string{
		0: "LogMode_UNSPECIFIED",
		1: "FIRST_AND_LAST",
	}
	ProtoMessageLoggingConfig_LogMode_value = map[string]int32{
		"LogMode_UNSPECIFIED": 0,
		"FIRST_AND_LAST":      1,
	}
)

func (x ProtoMessageLoggingConfig_LogMode) Enum() *ProtoMessageLoggingConfig_LogMode {
	p := new(ProtoMessageLoggingConfig_LogMode)
	*p = x
	return p
}

func (x ProtoMessageLoggingConfig_LogMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtoMessageLoggingConfig_LogMode) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_enumTypes[0].Descriptor()
}

func (ProtoMessageLoggingConfig_LogMode) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_enumTypes[0]
}

func (x ProtoMessageLoggingConfig_LogMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtoMessageLoggingConfig_LogMode.Descriptor instead.
func (ProtoMessageLoggingConfig_LogMode) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDescGZIP(), []int{0, 0}
}

type MethodLogging_LogDirective int32

const (
	MethodLogging_LogDirective_UNSPECIFIED MethodLogging_LogDirective = 0
	// The value of this field will be logged.
	MethodLogging_LOG MethodLogging_LogDirective = 1
	// It should be only annotated on Message type fields so if the field isn't
	// empty, an empty Struct will be logged.
	MethodLogging_LOG_REDACT MethodLogging_LogDirective = 2
)

// Enum value maps for MethodLogging_LogDirective.
var (
	MethodLogging_LogDirective_name = map[int32]string{
		0: "LogDirective_UNSPECIFIED",
		1: "LOG",
		2: "LOG_REDACT",
	}
	MethodLogging_LogDirective_value = map[string]int32{
		"LogDirective_UNSPECIFIED": 0,
		"LOG":                      1,
		"LOG_REDACT":               2,
	}
)

func (x MethodLogging_LogDirective) Enum() *MethodLogging_LogDirective {
	p := new(MethodLogging_LogDirective)
	*p = x
	return p
}

func (x MethodLogging_LogDirective) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MethodLogging_LogDirective) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_enumTypes[1].Descriptor()
}

func (MethodLogging_LogDirective) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_enumTypes[1]
}

func (x MethodLogging_LogDirective) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MethodLogging_LogDirective.Descriptor instead.
func (MethodLogging_LogDirective) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDescGZIP(), []int{1, 0}
}

type ProtoMessageLoggingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The proto descriptor set binary for the gRPC services.
	//
	// Types that are assignable to DescriptorSet:
	//
	//	*ProtoMessageLoggingConfig_DataSource
	//	*ProtoMessageLoggingConfig_ProtoDescriptorTypedMetadata
	DescriptorSet isProtoMessageLoggingConfig_DescriptorSet `protobuf_oneof:"descriptor_set"`
	Mode          ProtoMessageLoggingConfig_LogMode         `protobuf:"varint,3,opt,name=mode,proto3,enum=envoy.extensions.filters.http.proto_message_logging.v3.ProtoMessageLoggingConfig_LogMode" json:"mode,omitempty"`
	// Specify the message logging info.
	// The key is the fully qualified gRPC method name.
	// “${package}.${Service}.${Method}“, like
	// “endpoints.examples.bookstore.BookStore.GetShelf“
	//
	// The value is the message logging information for individual gRPC methods.
	LoggingByMethod map[string]*MethodLogging `protobuf:"bytes,4,rep,name=logging_by_method,json=loggingByMethod,proto3" json:"logging_by_method,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ProtoMessageLoggingConfig) Reset() {
	*x = ProtoMessageLoggingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoMessageLoggingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoMessageLoggingConfig) ProtoMessage() {}

func (x *ProtoMessageLoggingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoMessageLoggingConfig.ProtoReflect.Descriptor instead.
func (*ProtoMessageLoggingConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDescGZIP(), []int{0}
}

func (m *ProtoMessageLoggingConfig) GetDescriptorSet() isProtoMessageLoggingConfig_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (x *ProtoMessageLoggingConfig) GetDataSource() *v3.DataSource {
	if x, ok := x.GetDescriptorSet().(*ProtoMessageLoggingConfig_DataSource); ok {
		return x.DataSource
	}
	return nil
}

func (x *ProtoMessageLoggingConfig) GetProtoDescriptorTypedMetadata() string {
	if x, ok := x.GetDescriptorSet().(*ProtoMessageLoggingConfig_ProtoDescriptorTypedMetadata); ok {
		return x.ProtoDescriptorTypedMetadata
	}
	return ""
}

func (x *ProtoMessageLoggingConfig) GetMode() ProtoMessageLoggingConfig_LogMode {
	if x != nil {
		return x.Mode
	}
	return ProtoMessageLoggingConfig_LogMode_UNSPECIFIED
}

func (x *ProtoMessageLoggingConfig) GetLoggingByMethod() map[string]*MethodLogging {
	if x != nil {
		return x.LoggingByMethod
	}
	return nil
}

type isProtoMessageLoggingConfig_DescriptorSet interface {
	isProtoMessageLoggingConfig_DescriptorSet()
}

type ProtoMessageLoggingConfig_DataSource struct {
	// It could be passed by a local file through “Datasource.filename“ or
	// embedded in the “Datasource.inline_bytes“.
	DataSource *v3.DataSource `protobuf:"bytes,1,opt,name=data_source,json=dataSource,proto3,oneof"`
}

type ProtoMessageLoggingConfig_ProtoDescriptorTypedMetadata struct {
	// Unimplemented, the key of proto descriptor TypedMetadata.
	// Among filters depending on the proto descriptor, we can have a TypedMetadata
	// for proto descriptors, so that these filters can share one copy of proto
	// descriptor in memory.
	ProtoDescriptorTypedMetadata string `protobuf:"bytes,2,opt,name=proto_descriptor_typed_metadata,json=protoDescriptorTypedMetadata,proto3,oneof"`
}

func (*ProtoMessageLoggingConfig_DataSource) isProtoMessageLoggingConfig_DescriptorSet() {}

func (*ProtoMessageLoggingConfig_ProtoDescriptorTypedMetadata) isProtoMessageLoggingConfig_DescriptorSet() {
}

// This message can be used to support per route config approach later even
// though the Istio doesn't support that so far.
type MethodLogging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mapping of field path to its LogDirective for request messages
	RequestLoggingByField map[string]MethodLogging_LogDirective `protobuf:"bytes,2,rep,name=request_logging_by_field,json=requestLoggingByField,proto3" json:"request_logging_by_field,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging_LogDirective"`
	// The mapping of field path to its LogDirective for response messages
	ResponseLoggingByField map[string]MethodLogging_LogDirective `protobuf:"bytes,3,rep,name=response_logging_by_field,json=responseLoggingByField,proto3" json:"response_logging_by_field,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging_LogDirective"`
}

func (x *MethodLogging) Reset() {
	*x = MethodLogging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodLogging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodLogging) ProtoMessage() {}

func (x *MethodLogging) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodLogging.ProtoReflect.Descriptor instead.
func (*MethodLogging) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDescGZIP(), []int{1}
}

func (x *MethodLogging) GetRequestLoggingByField() map[string]MethodLogging_LogDirective {
	if x != nil {
		return x.RequestLoggingByField
	}
	return nil
}

func (x *MethodLogging) GetResponseLoggingByField() map[string]MethodLogging_LogDirective {
	if x != nil {
		return x.ResponseLoggingByField
	}
	return nil
}

var File_envoy_extensions_filters_http_proto_message_logging_v3_config_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDesc = []byte{
	0x0a, 0x43, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83,
	0x05, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a, 0x0b,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x47, 0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x1c, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x6d, 0x0a, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x59, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x33, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x92, 0x01, 0x0a, 0x11, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x66, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x42, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x89,
	0x01, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x5b, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x33, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x36, 0x0a, 0x07, 0x4c, 0x6f,
	0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4c, 0x41, 0x53, 0x54,
	0x10, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x5f, 0x73, 0x65, 0x74, 0x22, 0xd0, 0x05, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x99, 0x01, 0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x60, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x42,
	0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x9c, 0x01, 0x0a, 0x19, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x61, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x1a, 0x9c, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x68, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x52, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x9d, 0x01, 0x0a, 0x1b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x68, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x52, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x45, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x1c, 0x0a, 0x18, 0x4c, 0x6f, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x4c, 0x4f, 0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x47, 0x5f, 0x52,
	0x45, 0x44, 0x41, 0x43, 0x54, 0x10, 0x02, 0x42, 0xdc, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x0a, 0x44, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x42, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x75, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDescData = file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDesc
)

func file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDescData
}

var file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_goTypes = []interface{}{
	(ProtoMessageLoggingConfig_LogMode)(0), // 0: envoy.extensions.filters.http.proto_message_logging.v3.ProtoMessageLoggingConfig.LogMode
	(MethodLogging_LogDirective)(0),        // 1: envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging.LogDirective
	(*ProtoMessageLoggingConfig)(nil),      // 2: envoy.extensions.filters.http.proto_message_logging.v3.ProtoMessageLoggingConfig
	(*MethodLogging)(nil),                  // 3: envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging
	nil,                                    // 4: envoy.extensions.filters.http.proto_message_logging.v3.ProtoMessageLoggingConfig.LoggingByMethodEntry
	nil,                                    // 5: envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging.RequestLoggingByFieldEntry
	nil,                                    // 6: envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging.ResponseLoggingByFieldEntry
	(*v3.DataSource)(nil),                  // 7: envoy.config.core.v3.DataSource
}
var file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_depIdxs = []int32{
	7, // 0: envoy.extensions.filters.http.proto_message_logging.v3.ProtoMessageLoggingConfig.data_source:type_name -> envoy.config.core.v3.DataSource
	0, // 1: envoy.extensions.filters.http.proto_message_logging.v3.ProtoMessageLoggingConfig.mode:type_name -> envoy.extensions.filters.http.proto_message_logging.v3.ProtoMessageLoggingConfig.LogMode
	4, // 2: envoy.extensions.filters.http.proto_message_logging.v3.ProtoMessageLoggingConfig.logging_by_method:type_name -> envoy.extensions.filters.http.proto_message_logging.v3.ProtoMessageLoggingConfig.LoggingByMethodEntry
	5, // 3: envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging.request_logging_by_field:type_name -> envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging.RequestLoggingByFieldEntry
	6, // 4: envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging.response_logging_by_field:type_name -> envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging.ResponseLoggingByFieldEntry
	3, // 5: envoy.extensions.filters.http.proto_message_logging.v3.ProtoMessageLoggingConfig.LoggingByMethodEntry.value:type_name -> envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging
	1, // 6: envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging.RequestLoggingByFieldEntry.value:type_name -> envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging.LogDirective
	1, // 7: envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging.ResponseLoggingByFieldEntry.value:type_name -> envoy.extensions.filters.http.proto_message_logging.v3.MethodLogging.LogDirective
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_init() }
func file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_init() {
	if File_envoy_extensions_filters_http_proto_message_logging_v3_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoMessageLoggingConfig); i {
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
		file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodLogging); i {
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
	file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ProtoMessageLoggingConfig_DataSource)(nil),
		(*ProtoMessageLoggingConfig_ProtoDescriptorTypedMetadata)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_proto_message_logging_v3_config_proto = out.File
	file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_rawDesc = nil
	file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_goTypes = nil
	file_envoy_extensions_filters_http_proto_message_logging_v3_config_proto_depIdxs = nil
}
