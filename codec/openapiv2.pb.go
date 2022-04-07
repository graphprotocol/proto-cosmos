// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.0
// source: openapiv2.proto

package options

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

type JSONSchema_JSONSchemaSimpleTypes int32

const (
	JSONSchema_UNKNOWN JSONSchema_JSONSchemaSimpleTypes = 0
	JSONSchema_ARRAY   JSONSchema_JSONSchemaSimpleTypes = 1
	JSONSchema_BOOLEAN JSONSchema_JSONSchemaSimpleTypes = 2
	JSONSchema_INTEGER JSONSchema_JSONSchemaSimpleTypes = 3
	JSONSchema_NULL    JSONSchema_JSONSchemaSimpleTypes = 4
	JSONSchema_NUMBER  JSONSchema_JSONSchemaSimpleTypes = 5
	JSONSchema_OBJECT  JSONSchema_JSONSchemaSimpleTypes = 6
	JSONSchema_STRING  JSONSchema_JSONSchemaSimpleTypes = 7
)

// Enum value maps for JSONSchema_JSONSchemaSimpleTypes.
var (
	JSONSchema_JSONSchemaSimpleTypes_name = map[int32]string{
		0: "UNKNOWN",
		1: "ARRAY",
		2: "BOOLEAN",
		3: "INTEGER",
		4: "NULL",
		5: "NUMBER",
		6: "OBJECT",
		7: "STRING",
	}
	JSONSchema_JSONSchemaSimpleTypes_value = map[string]int32{
		"UNKNOWN": 0,
		"ARRAY":   1,
		"BOOLEAN": 2,
		"INTEGER": 3,
		"NULL":    4,
		"NUMBER":  5,
		"OBJECT":  6,
		"STRING":  7,
	}
)

func (x JSONSchema_JSONSchemaSimpleTypes) Enum() *JSONSchema_JSONSchemaSimpleTypes {
	p := new(JSONSchema_JSONSchemaSimpleTypes)
	*p = x
	return p
}

func (x JSONSchema_JSONSchemaSimpleTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JSONSchema_JSONSchemaSimpleTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_openapiv2_proto_enumTypes[0].Descriptor()
}

func (JSONSchema_JSONSchemaSimpleTypes) Type() protoreflect.EnumType {
	return &file_openapiv2_proto_enumTypes[0]
}

func (x JSONSchema_JSONSchemaSimpleTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JSONSchema_JSONSchemaSimpleTypes.Descriptor instead.
func (JSONSchema_JSONSchemaSimpleTypes) EnumDescriptor() ([]byte, []int) {
	return file_openapiv2_proto_rawDescGZIP(), []int{0, 0}
}

type JSONSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref                string                             `protobuf:"bytes,3,opt,name=ref,proto3" json:"ref,omitempty"`
	Title              string                             `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description        string                             `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Default            string                             `protobuf:"bytes,7,opt,name=default,proto3" json:"default,omitempty"`
	ReadOnly           bool                               `protobuf:"varint,8,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	Example            string                             `protobuf:"bytes,9,opt,name=example,proto3" json:"example,omitempty"`
	MultipleOf         float64                            `protobuf:"fixed64,10,opt,name=multiple_of,json=multipleOf,proto3" json:"multiple_of,omitempty"`
	Maximum            float64                            `protobuf:"fixed64,11,opt,name=maximum,proto3" json:"maximum,omitempty"`
	ExclusiveMaximum   bool                               `protobuf:"varint,12,opt,name=exclusive_maximum,json=exclusiveMaximum,proto3" json:"exclusive_maximum,omitempty"`
	Minimum            float64                            `protobuf:"fixed64,13,opt,name=minimum,proto3" json:"minimum,omitempty"`
	ExclusiveMinimum   bool                               `protobuf:"varint,14,opt,name=exclusive_minimum,json=exclusiveMinimum,proto3" json:"exclusive_minimum,omitempty"`
	MaxLength          uint64                             `protobuf:"varint,15,opt,name=max_length,json=maxLength,proto3" json:"max_length,omitempty"`
	MinLength          uint64                             `protobuf:"varint,16,opt,name=min_length,json=minLength,proto3" json:"min_length,omitempty"`
	Pattern            string                             `protobuf:"bytes,17,opt,name=pattern,proto3" json:"pattern,omitempty"`
	MaxItems           uint64                             `protobuf:"varint,20,opt,name=max_items,json=maxItems,proto3" json:"max_items,omitempty"`
	MinItems           uint64                             `protobuf:"varint,21,opt,name=min_items,json=minItems,proto3" json:"min_items,omitempty"`
	UniqueItems        bool                               `protobuf:"varint,22,opt,name=unique_items,json=uniqueItems,proto3" json:"unique_items,omitempty"`
	MaxProperties      uint64                             `protobuf:"varint,24,opt,name=max_properties,json=maxProperties,proto3" json:"max_properties,omitempty"`
	MinProperties      uint64                             `protobuf:"varint,25,opt,name=min_properties,json=minProperties,proto3" json:"min_properties,omitempty"`
	Required           []string                           `protobuf:"bytes,26,rep,name=required,proto3" json:"required,omitempty"`
	Array              []string                           `protobuf:"bytes,34,rep,name=array,proto3" json:"array,omitempty"`
	Type               []JSONSchema_JSONSchemaSimpleTypes `protobuf:"varint,35,rep,packed,name=type,proto3,enum=grpc.gateway.protoc_gen_openapiv2.options.JSONSchema_JSONSchemaSimpleTypes" json:"type,omitempty"`
	Format             string                             `protobuf:"bytes,36,opt,name=format,proto3" json:"format,omitempty"`
	Enum               []string                           `protobuf:"bytes,46,rep,name=enum,proto3" json:"enum,omitempty"`
	FieldConfiguration *JSONSchema_FieldConfiguration     `protobuf:"bytes,1001,opt,name=field_configuration,json=fieldConfiguration,proto3" json:"field_configuration,omitempty"`
}

func (x *JSONSchema) Reset() {
	*x = JSONSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapiv2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JSONSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSONSchema) ProtoMessage() {}

func (x *JSONSchema) ProtoReflect() protoreflect.Message {
	mi := &file_openapiv2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSONSchema.ProtoReflect.Descriptor instead.
func (*JSONSchema) Descriptor() ([]byte, []int) {
	return file_openapiv2_proto_rawDescGZIP(), []int{0}
}

func (x *JSONSchema) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *JSONSchema) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *JSONSchema) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *JSONSchema) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *JSONSchema) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

func (x *JSONSchema) GetExample() string {
	if x != nil {
		return x.Example
	}
	return ""
}

func (x *JSONSchema) GetMultipleOf() float64 {
	if x != nil {
		return x.MultipleOf
	}
	return 0
}

func (x *JSONSchema) GetMaximum() float64 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

func (x *JSONSchema) GetExclusiveMaximum() bool {
	if x != nil {
		return x.ExclusiveMaximum
	}
	return false
}

func (x *JSONSchema) GetMinimum() float64 {
	if x != nil {
		return x.Minimum
	}
	return 0
}

func (x *JSONSchema) GetExclusiveMinimum() bool {
	if x != nil {
		return x.ExclusiveMinimum
	}
	return false
}

func (x *JSONSchema) GetMaxLength() uint64 {
	if x != nil {
		return x.MaxLength
	}
	return 0
}

func (x *JSONSchema) GetMinLength() uint64 {
	if x != nil {
		return x.MinLength
	}
	return 0
}

func (x *JSONSchema) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *JSONSchema) GetMaxItems() uint64 {
	if x != nil {
		return x.MaxItems
	}
	return 0
}

func (x *JSONSchema) GetMinItems() uint64 {
	if x != nil {
		return x.MinItems
	}
	return 0
}

func (x *JSONSchema) GetUniqueItems() bool {
	if x != nil {
		return x.UniqueItems
	}
	return false
}

func (x *JSONSchema) GetMaxProperties() uint64 {
	if x != nil {
		return x.MaxProperties
	}
	return 0
}

func (x *JSONSchema) GetMinProperties() uint64 {
	if x != nil {
		return x.MinProperties
	}
	return 0
}

func (x *JSONSchema) GetRequired() []string {
	if x != nil {
		return x.Required
	}
	return nil
}

func (x *JSONSchema) GetArray() []string {
	if x != nil {
		return x.Array
	}
	return nil
}

func (x *JSONSchema) GetType() []JSONSchema_JSONSchemaSimpleTypes {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *JSONSchema) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *JSONSchema) GetEnum() []string {
	if x != nil {
		return x.Enum
	}
	return nil
}

func (x *JSONSchema) GetFieldConfiguration() *JSONSchema_FieldConfiguration {
	if x != nil {
		return x.FieldConfiguration
	}
	return nil
}

type JSONSchema_FieldConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PathParamName string `protobuf:"bytes,47,opt,name=path_param_name,json=pathParamName,proto3" json:"path_param_name,omitempty"`
}

func (x *JSONSchema_FieldConfiguration) Reset() {
	*x = JSONSchema_FieldConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapiv2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JSONSchema_FieldConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSONSchema_FieldConfiguration) ProtoMessage() {}

func (x *JSONSchema_FieldConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_openapiv2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSONSchema_FieldConfiguration.ProtoReflect.Descriptor instead.
func (*JSONSchema_FieldConfiguration) Descriptor() ([]byte, []int) {
	return file_openapiv2_proto_rawDescGZIP(), []int{0, 0}
}

func (x *JSONSchema_FieldConfiguration) GetPathParamName() string {
	if x != nil {
		return x.PathParamName
	}
	return ""
}

var File_openapiv2_proto protoreflect.FileDescriptor

var file_openapiv2_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x29, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x99, 0x09, 0x0a,
	0x0a, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x5f, 0x6f, 0x66, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x65, 0x4f, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75,
	0x6d, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x6d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x78, 0x63, 0x6c,
	0x75, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x10, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x4d, 0x69,
	0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x18, 0x10, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x6d, 0x61, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69,
	0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d,
	0x69, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61,
	0x78, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6d, 0x69, 0x6e, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x22, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x12, 0x5f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x23, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67,
	0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x2e, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x7a, 0x0a, 0x13, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xe9,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x12, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x3c, 0x0a, 0x12, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x74,
	0x68, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x2f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x77, 0x0a, 0x15, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x52, 0x52, 0x41, 0x59,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52,
	0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x06, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02,
	0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x12,
	0x10, 0x13, 0x4a, 0x04, 0x08, 0x13, 0x10, 0x14, 0x4a, 0x04, 0x08, 0x17, 0x10, 0x18, 0x4a, 0x04,
	0x08, 0x1b, 0x10, 0x1c, 0x4a, 0x04, 0x08, 0x1c, 0x10, 0x1d, 0x4a, 0x04, 0x08, 0x1d, 0x10, 0x1e,
	0x4a, 0x04, 0x08, 0x1e, 0x10, 0x22, 0x4a, 0x04, 0x08, 0x25, 0x10, 0x2a, 0x4a, 0x04, 0x08, 0x2a,
	0x10, 0x2b, 0x4a, 0x04, 0x08, 0x2b, 0x10, 0x2e, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_openapiv2_proto_rawDescOnce sync.Once
	file_openapiv2_proto_rawDescData = file_openapiv2_proto_rawDesc
)

func file_openapiv2_proto_rawDescGZIP() []byte {
	file_openapiv2_proto_rawDescOnce.Do(func() {
		file_openapiv2_proto_rawDescData = protoimpl.X.CompressGZIP(file_openapiv2_proto_rawDescData)
	})
	return file_openapiv2_proto_rawDescData
}

var file_openapiv2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_openapiv2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_openapiv2_proto_goTypes = []interface{}{
	(JSONSchema_JSONSchemaSimpleTypes)(0), // 0: grpc.gateway.protoc_gen_openapiv2.options.JSONSchema.JSONSchemaSimpleTypes
	(*JSONSchema)(nil),                    // 1: grpc.gateway.protoc_gen_openapiv2.options.JSONSchema
	(*JSONSchema_FieldConfiguration)(nil), // 2: grpc.gateway.protoc_gen_openapiv2.options.JSONSchema.FieldConfiguration
}
var file_openapiv2_proto_depIdxs = []int32{
	0, // 0: grpc.gateway.protoc_gen_openapiv2.options.JSONSchema.type:type_name -> grpc.gateway.protoc_gen_openapiv2.options.JSONSchema.JSONSchemaSimpleTypes
	2, // 1: grpc.gateway.protoc_gen_openapiv2.options.JSONSchema.field_configuration:type_name -> grpc.gateway.protoc_gen_openapiv2.options.JSONSchema.FieldConfiguration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_openapiv2_proto_init() }
func file_openapiv2_proto_init() {
	if File_openapiv2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_openapiv2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JSONSchema); i {
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
		file_openapiv2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JSONSchema_FieldConfiguration); i {
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
			RawDescriptor: file_openapiv2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_openapiv2_proto_goTypes,
		DependencyIndexes: file_openapiv2_proto_depIdxs,
		EnumInfos:         file_openapiv2_proto_enumTypes,
		MessageInfos:      file_openapiv2_proto_msgTypes,
	}.Build()
	File_openapiv2_proto = out.File
	file_openapiv2_proto_rawDesc = nil
	file_openapiv2_proto_goTypes = nil
	file_openapiv2_proto_depIdxs = nil
}
