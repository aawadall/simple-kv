// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kv_service.proto

package kv_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type KeyValueRecord struct {
	Key                  string            `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte            `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *KeyValueRecord) Reset()         { *m = KeyValueRecord{} }
func (m *KeyValueRecord) String() string { return proto.CompactTextString(m) }
func (*KeyValueRecord) ProtoMessage()    {}
func (*KeyValueRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{0}
}

func (m *KeyValueRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValueRecord.Unmarshal(m, b)
}
func (m *KeyValueRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValueRecord.Marshal(b, m, deterministic)
}
func (m *KeyValueRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValueRecord.Merge(m, src)
}
func (m *KeyValueRecord) XXX_Size() int {
	return xxx_messageInfo_KeyValueRecord.Size(m)
}
func (m *KeyValueRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValueRecord.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValueRecord proto.InternalMessageInfo

func (m *KeyValueRecord) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValueRecord) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyValueRecord) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type SetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{2}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SetResponse struct {
	Response             *UniversalResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{3}
}

func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (m *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(m, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

func (m *SetResponse) GetResponse() *UniversalResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type DeleteRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{4}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DeleteResponse struct {
	Response             *UniversalResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{5}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetResponse() *UniversalResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type SetMetadataRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	MetadataKey          string   `protobuf:"bytes,2,opt,name=metadata_key,json=metadataKey,proto3" json:"metadata_key,omitempty"`
	MetadataValue        string   `protobuf:"bytes,3,opt,name=metadata_value,json=metadataValue,proto3" json:"metadata_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetMetadataRequest) Reset()         { *m = SetMetadataRequest{} }
func (m *SetMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*SetMetadataRequest) ProtoMessage()    {}
func (*SetMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{6}
}

func (m *SetMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetMetadataRequest.Unmarshal(m, b)
}
func (m *SetMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetMetadataRequest.Marshal(b, m, deterministic)
}
func (m *SetMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMetadataRequest.Merge(m, src)
}
func (m *SetMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_SetMetadataRequest.Size(m)
}
func (m *SetMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetMetadataRequest proto.InternalMessageInfo

func (m *SetMetadataRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetMetadataRequest) GetMetadataKey() string {
	if m != nil {
		return m.MetadataKey
	}
	return ""
}

func (m *SetMetadataRequest) GetMetadataValue() string {
	if m != nil {
		return m.MetadataValue
	}
	return ""
}

type SetMetadataResponse struct {
	Response             *UniversalResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SetMetadataResponse) Reset()         { *m = SetMetadataResponse{} }
func (m *SetMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*SetMetadataResponse) ProtoMessage()    {}
func (*SetMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{7}
}

func (m *SetMetadataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetMetadataResponse.Unmarshal(m, b)
}
func (m *SetMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetMetadataResponse.Marshal(b, m, deterministic)
}
func (m *SetMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMetadataResponse.Merge(m, src)
}
func (m *SetMetadataResponse) XXX_Size() int {
	return xxx_messageInfo_SetMetadataResponse.Size(m)
}
func (m *SetMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetMetadataResponse proto.InternalMessageInfo

func (m *SetMetadataResponse) GetResponse() *UniversalResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type DeleteMetadataRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	MetadataKey          string   `protobuf:"bytes,2,opt,name=metadata_key,json=metadataKey,proto3" json:"metadata_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMetadataRequest) Reset()         { *m = DeleteMetadataRequest{} }
func (m *DeleteMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMetadataRequest) ProtoMessage()    {}
func (*DeleteMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{8}
}

func (m *DeleteMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMetadataRequest.Unmarshal(m, b)
}
func (m *DeleteMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMetadataRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMetadataRequest.Merge(m, src)
}
func (m *DeleteMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMetadataRequest.Size(m)
}
func (m *DeleteMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMetadataRequest proto.InternalMessageInfo

func (m *DeleteMetadataRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DeleteMetadataRequest) GetMetadataKey() string {
	if m != nil {
		return m.MetadataKey
	}
	return ""
}

type DeleteMetadataResponse struct {
	Response             *UniversalResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DeleteMetadataResponse) Reset()         { *m = DeleteMetadataResponse{} }
func (m *DeleteMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteMetadataResponse) ProtoMessage()    {}
func (*DeleteMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{9}
}

func (m *DeleteMetadataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMetadataResponse.Unmarshal(m, b)
}
func (m *DeleteMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMetadataResponse.Marshal(b, m, deterministic)
}
func (m *DeleteMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMetadataResponse.Merge(m, src)
}
func (m *DeleteMetadataResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteMetadataResponse.Size(m)
}
func (m *DeleteMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMetadataResponse proto.InternalMessageInfo

func (m *DeleteMetadataResponse) GetResponse() *UniversalResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type GetAllMetadataRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllMetadataRequest) Reset()         { *m = GetAllMetadataRequest{} }
func (m *GetAllMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllMetadataRequest) ProtoMessage()    {}
func (*GetAllMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{10}
}

func (m *GetAllMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllMetadataRequest.Unmarshal(m, b)
}
func (m *GetAllMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllMetadataRequest.Marshal(b, m, deterministic)
}
func (m *GetAllMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllMetadataRequest.Merge(m, src)
}
func (m *GetAllMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllMetadataRequest.Size(m)
}
func (m *GetAllMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllMetadataRequest proto.InternalMessageInfo

func (m *GetAllMetadataRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetAllMetadataResponse struct {
	Response             *UniversalResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Metadata             map[string]string  `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetAllMetadataResponse) Reset()         { *m = GetAllMetadataResponse{} }
func (m *GetAllMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllMetadataResponse) ProtoMessage()    {}
func (*GetAllMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{11}
}

func (m *GetAllMetadataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllMetadataResponse.Unmarshal(m, b)
}
func (m *GetAllMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllMetadataResponse.Marshal(b, m, deterministic)
}
func (m *GetAllMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllMetadataResponse.Merge(m, src)
}
func (m *GetAllMetadataResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllMetadataResponse.Size(m)
}
func (m *GetAllMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllMetadataResponse proto.InternalMessageInfo

func (m *GetAllMetadataResponse) GetResponse() *UniversalResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *GetAllMetadataResponse) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type FindRequest struct {
	PartialKey           string   `protobuf:"bytes,1,opt,name=partial_key,json=partialKey,proto3" json:"partial_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{12}
}

func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (m *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(m, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetPartialKey() string {
	if m != nil {
		return m.PartialKey
	}
	return ""
}

type FindResponse struct {
	Response             *UniversalResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Records              []string           `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FindResponse) Reset()         { *m = FindResponse{} }
func (m *FindResponse) String() string { return proto.CompactTextString(m) }
func (*FindResponse) ProtoMessage()    {}
func (*FindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{13}
}

func (m *FindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResponse.Unmarshal(m, b)
}
func (m *FindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResponse.Marshal(b, m, deterministic)
}
func (m *FindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResponse.Merge(m, src)
}
func (m *FindResponse) XXX_Size() int {
	return xxx_messageInfo_FindResponse.Size(m)
}
func (m *FindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindResponse proto.InternalMessageInfo

func (m *FindResponse) GetResponse() *UniversalResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *FindResponse) GetRecords() []string {
	if m != nil {
		return m.Records
	}
	return nil
}

type FindByMetadataRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindByMetadataRequest) Reset()         { *m = FindByMetadataRequest{} }
func (m *FindByMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*FindByMetadataRequest) ProtoMessage()    {}
func (*FindByMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{14}
}

func (m *FindByMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByMetadataRequest.Unmarshal(m, b)
}
func (m *FindByMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByMetadataRequest.Marshal(b, m, deterministic)
}
func (m *FindByMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByMetadataRequest.Merge(m, src)
}
func (m *FindByMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_FindByMetadataRequest.Size(m)
}
func (m *FindByMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindByMetadataRequest proto.InternalMessageInfo

func (m *FindByMetadataRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type FindByMetadataResponse struct {
	Response             *UniversalResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Records              []string           `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FindByMetadataResponse) Reset()         { *m = FindByMetadataResponse{} }
func (m *FindByMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*FindByMetadataResponse) ProtoMessage()    {}
func (*FindByMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{15}
}

func (m *FindByMetadataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByMetadataResponse.Unmarshal(m, b)
}
func (m *FindByMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByMetadataResponse.Marshal(b, m, deterministic)
}
func (m *FindByMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByMetadataResponse.Merge(m, src)
}
func (m *FindByMetadataResponse) XXX_Size() int {
	return xxx_messageInfo_FindByMetadataResponse.Size(m)
}
func (m *FindByMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindByMetadataResponse proto.InternalMessageInfo

func (m *FindByMetadataResponse) GetResponse() *UniversalResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *FindByMetadataResponse) GetRecords() []string {
	if m != nil {
		return m.Records
	}
	return nil
}

type UniversalResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UniversalResponse) Reset()         { *m = UniversalResponse{} }
func (m *UniversalResponse) String() string { return proto.CompactTextString(m) }
func (*UniversalResponse) ProtoMessage()    {}
func (*UniversalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2489677d3d3be1b1, []int{16}
}

func (m *UniversalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UniversalResponse.Unmarshal(m, b)
}
func (m *UniversalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UniversalResponse.Marshal(b, m, deterministic)
}
func (m *UniversalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniversalResponse.Merge(m, src)
}
func (m *UniversalResponse) XXX_Size() int {
	return xxx_messageInfo_UniversalResponse.Size(m)
}
func (m *UniversalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UniversalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UniversalResponse proto.InternalMessageInfo

func (m *UniversalResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UniversalResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*KeyValueRecord)(nil), "aawadall.simple_kv.proto_api.KeyValueRecord")
	proto.RegisterMapType((map[string]string)(nil), "aawadall.simple_kv.proto_api.KeyValueRecord.MetadataEntry")
	proto.RegisterType((*GetRequest)(nil), "aawadall.simple_kv.proto_api.GetRequest")
	proto.RegisterType((*SetRequest)(nil), "aawadall.simple_kv.proto_api.SetRequest")
	proto.RegisterType((*SetResponse)(nil), "aawadall.simple_kv.proto_api.SetResponse")
	proto.RegisterType((*DeleteRequest)(nil), "aawadall.simple_kv.proto_api.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "aawadall.simple_kv.proto_api.DeleteResponse")
	proto.RegisterType((*SetMetadataRequest)(nil), "aawadall.simple_kv.proto_api.SetMetadataRequest")
	proto.RegisterType((*SetMetadataResponse)(nil), "aawadall.simple_kv.proto_api.SetMetadataResponse")
	proto.RegisterType((*DeleteMetadataRequest)(nil), "aawadall.simple_kv.proto_api.DeleteMetadataRequest")
	proto.RegisterType((*DeleteMetadataResponse)(nil), "aawadall.simple_kv.proto_api.DeleteMetadataResponse")
	proto.RegisterType((*GetAllMetadataRequest)(nil), "aawadall.simple_kv.proto_api.GetAllMetadataRequest")
	proto.RegisterType((*GetAllMetadataResponse)(nil), "aawadall.simple_kv.proto_api.GetAllMetadataResponse")
	proto.RegisterMapType((map[string]string)(nil), "aawadall.simple_kv.proto_api.GetAllMetadataResponse.MetadataEntry")
	proto.RegisterType((*FindRequest)(nil), "aawadall.simple_kv.proto_api.FindRequest")
	proto.RegisterType((*FindResponse)(nil), "aawadall.simple_kv.proto_api.FindResponse")
	proto.RegisterType((*FindByMetadataRequest)(nil), "aawadall.simple_kv.proto_api.FindByMetadataRequest")
	proto.RegisterType((*FindByMetadataResponse)(nil), "aawadall.simple_kv.proto_api.FindByMetadataResponse")
	proto.RegisterType((*UniversalResponse)(nil), "aawadall.simple_kv.proto_api.UniversalResponse")
}

func init() { proto.RegisterFile("kv_service.proto", fileDescriptor_2489677d3d3be1b1) }

var fileDescriptor_2489677d3d3be1b1 = []byte{
	// 642 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xce, 0x47, 0x3f, 0xc7, 0x6d, 0xdf, 0xbe, 0x4b, 0x13, 0x45, 0x11, 0x82, 0x74, 0x25, 0xa4,
	0x14, 0xa8, 0x53, 0xd2, 0x1e, 0x50, 0x39, 0x11, 0x3e, 0x72, 0x08, 0x5c, 0x1c, 0xd1, 0x43, 0x05,
	0x58, 0x9b, 0x64, 0x04, 0x56, 0x9c, 0xd8, 0x5d, 0xaf, 0x8d, 0x72, 0x40, 0x48, 0x5c, 0xf9, 0x6d,
	0xfc, 0x27, 0x64, 0x3b, 0x9b, 0xc4, 0xae, 0x1b, 0x27, 0x12, 0xe6, 0x96, 0x59, 0xcf, 0xcc, 0xf3,
	0xcc, 0xec, 0xcc, 0xb3, 0x81, 0xc3, 0xa1, 0xa7, 0x3b, 0xc8, 0x3d, 0xa3, 0x8f, 0xaa, 0xcd, 0x2d,
	0x61, 0x91, 0xfb, 0x8c, 0x7d, 0x63, 0x03, 0x66, 0x9a, 0xaa, 0x63, 0x8c, 0x6c, 0x13, 0xf5, 0xa1,
	0x17, 0x7e, 0xd1, 0x99, 0x6d, 0xd0, 0xdf, 0x79, 0x38, 0xe8, 0xe0, 0xe4, 0x8a, 0x99, 0x2e, 0x6a,
	0xd8, 0xb7, 0xf8, 0x80, 0x1c, 0x42, 0x71, 0x88, 0x93, 0x4a, 0xa1, 0x96, 0xaf, 0xef, 0x6a, 0xfe,
	0x4f, 0x72, 0x04, 0x9b, 0x9e, 0xef, 0x50, 0x29, 0xd6, 0xf2, 0xf5, 0x3d, 0x2d, 0x34, 0xc8, 0x15,
	0xec, 0x8c, 0x50, 0xb0, 0x01, 0x13, 0xac, 0xb2, 0x51, 0x2b, 0xd6, 0x95, 0xe6, 0xa5, 0xba, 0x0c,
	0x4b, 0x8d, 0xe2, 0xa8, 0xef, 0xa7, 0xc1, 0x6f, 0xc6, 0x82, 0x4f, 0xb4, 0x59, 0xae, 0xea, 0x0b,
	0xd8, 0x8f, 0x7c, 0x92, 0x84, 0xf2, 0x09, 0x84, 0x42, 0x92, 0xa1, 0x71, 0x59, 0x78, 0x9e, 0xa7,
	0x0f, 0x00, 0xda, 0x28, 0x34, 0xbc, 0x71, 0xd1, 0x11, 0xb7, 0x23, 0xe9, 0x05, 0x40, 0x77, 0xc9,
	0xf7, 0x68, 0x66, 0x59, 0x2a, 0xbd, 0x06, 0x25, 0x88, 0x72, 0x6c, 0x6b, 0xec, 0x20, 0xe9, 0xc0,
	0x0e, 0x9f, 0xfe, 0x0e, 0x62, 0x95, 0x66, 0x63, 0x79, 0xe5, 0x1f, 0xc6, 0x86, 0x87, 0xdc, 0x61,
	0xa6, 0x4c, 0xa1, 0xcd, 0x12, 0xd0, 0x63, 0xd8, 0x7f, 0x8d, 0x26, 0x0a, 0xbc, 0x9b, 0xf4, 0x27,
	0x38, 0x90, 0x2e, 0x59, 0x30, 0xb0, 0x81, 0x74, 0x51, 0xc8, 0x9e, 0xdf, 0xdd, 0x9b, 0x63, 0xd8,
	0x93, 0x97, 0xa4, 0xcf, 0x27, 0x44, 0x91, 0x67, 0x1d, 0x9c, 0x90, 0x47, 0x70, 0x30, 0x73, 0x99,
	0x8f, 0xcc, 0xae, 0xb6, 0x2f, 0x4f, 0x83, 0x01, 0xa0, 0x3d, 0xb8, 0x17, 0x41, 0xcc, 0xa2, 0xaa,
	0x77, 0x50, 0x0a, 0x9b, 0xf6, 0x37, 0x0a, 0xa3, 0x08, 0xe5, 0x78, 0xb6, 0x2c, 0x48, 0x9f, 0x40,
	0xa9, 0x8d, 0xe2, 0xa5, 0x69, 0xa6, 0x92, 0xa6, 0x3f, 0x0b, 0x50, 0x8e, 0xfb, 0x66, 0x40, 0x89,
	0x7c, 0x5e, 0x58, 0xf3, 0x42, 0xb0, 0xe6, 0xad, 0xe5, 0xc9, 0x92, 0x49, 0x65, 0xb3, 0xee, 0x2a,
	0x28, 0x6f, 0x8d, 0xf1, 0x40, 0x76, 0xe9, 0x21, 0x28, 0x36, 0xe3, 0xc2, 0x60, 0xa6, 0x3e, 0x4f,
	0x01, 0xd3, 0x23, 0xff, 0x1a, 0x5d, 0xd8, 0x0b, 0xfd, 0xb3, 0xe8, 0x54, 0x05, 0xb6, 0x79, 0x20,
	0x6d, 0x4e, 0xd0, 0xa8, 0x5d, 0x4d, 0x9a, 0xf4, 0x14, 0x4a, 0x3e, 0x6c, 0x6b, 0x12, 0xbf, 0xd6,
	0x23, 0xd8, 0xbc, 0x71, 0x91, 0x4b, 0xaa, 0xa1, 0x41, 0x7f, 0x40, 0x39, 0xee, 0xfe, 0x6f, 0xf9,
	0xbe, 0x82, 0xff, 0x6f, 0x05, 0xfa, 0xee, 0x8e, 0xdb, 0xef, 0xa3, 0xe3, 0x04, 0xd0, 0x3b, 0x9a,
	0x34, 0xfd, 0x2a, 0x90, 0x73, 0x8b, 0xcb, 0xfb, 0x09, 0x8c, 0xe6, 0xaf, 0x6d, 0xf8, 0x4f, 0x4a,
	0x7e, 0x37, 0x7c, 0x92, 0x88, 0x0e, 0xc5, 0x36, 0x0a, 0x52, 0x4f, 0x9d, 0xa0, 0x69, 0x83, 0xaa,
	0x4f, 0xd7, 0x79, 0x52, 0x68, 0x8e, 0x7c, 0x84, 0x62, 0x37, 0x1d, 0x60, 0xfe, 0x04, 0x54, 0x4f,
	0x56, 0xf0, 0x9c, 0x2e, 0x67, 0x8e, 0x20, 0x6c, 0x85, 0x2a, 0x40, 0x9e, 0x2c, 0x0f, 0x8b, 0x28,
	0x7a, 0x5a, 0x11, 0x51, 0x6d, 0xa7, 0x39, 0x22, 0x82, 0xe7, 0x46, 0x5e, 0x3e, 0x39, 0x4b, 0xa5,
	0x18, 0x1b, 0xab, 0xea, 0xb3, 0x35, 0x22, 0x66, 0xa8, 0xdf, 0xe5, 0x2b, 0x33, 0x03, 0x3e, 0x5f,
	0x85, 0x77, 0x1c, 0xfb, 0x62, 0xbd, 0xa0, 0x45, 0xf8, 0xa8, 0x72, 0xa4, 0xc1, 0x27, 0x0a, 0x65,
	0x1a, 0x7c, 0xb2, 0x38, 0xd1, 0x1c, 0xd1, 0x61, 0xc3, 0xdf, 0x39, 0x92, 0x32, 0x0f, 0x0b, 0x6a,
	0x53, 0x7d, 0xbc, 0x8a, 0xeb, 0x62, 0x7d, 0xd1, 0xa5, 0x4e, 0xab, 0x2f, 0x51, 0x31, 0xd2, 0xea,
	0x4b, 0xd6, 0x0d, 0x9a, 0x6b, 0x9d, 0x5d, 0xab, 0x5f, 0x0c, 0xf1, 0xd5, 0xed, 0xa9, 0x7d, 0x6b,
	0xd4, 0x90, 0x39, 0x1a, 0x61, 0x8e, 0xd3, 0xa1, 0xd7, 0x08, 0x72, 0x34, 0xe6, 0x7f, 0x1f, 0x7b,
	0x5b, 0xc1, 0xc9, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x81, 0x60, 0x89, 0x53, 0x0a,
	0x00, 0x00,
}
