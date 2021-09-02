// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: rate.proto

package rate

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

type Rate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MarketName string `protobuf:"bytes,2,opt,name=market_name,json=marketName,proto3" json:"market_name,omitempty"`
	High       string `protobuf:"bytes,3,opt,name=high,proto3" json:"high,omitempty"`
	Low        string `protobuf:"bytes,4,opt,name=low,proto3" json:"low,omitempty"`
	Volume     string `protobuf:"bytes,5,opt,name=volume,proto3" json:"volume,omitempty"`
	Timestamp  string `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Rate) Reset() {
	*x = Rate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rate) ProtoMessage() {}

func (x *Rate) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rate.ProtoReflect.Descriptor instead.
func (*Rate) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{0}
}

func (x *Rate) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rate) GetMarketName() string {
	if x != nil {
		return x.MarketName
	}
	return ""
}

func (x *Rate) GetHigh() string {
	if x != nil {
		return x.High
	}
	return ""
}

func (x *Rate) GetLow() string {
	if x != nil {
		return x.Low
	}
	return ""
}

func (x *Rate) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

func (x *Rate) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type GrowthRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FromRateId   int32  `protobuf:"varint,2,opt,name=from_rate_id,json=fromRateId,proto3" json:"from_rate_id,omitempty"`
	ToRateId     int32  `protobuf:"varint,3,opt,name=to_rate_id,json=toRateId,proto3" json:"to_rate_id,omitempty"`
	VolumeGrowth string `protobuf:"bytes,4,opt,name=volume_growth,json=volumeGrowth,proto3" json:"volume_growth,omitempty"`
	HighGrowth   string `protobuf:"bytes,5,opt,name=high_growth,json=highGrowth,proto3" json:"high_growth,omitempty"`
	LowGrowth    string `protobuf:"bytes,6,opt,name=low_growth,json=lowGrowth,proto3" json:"low_growth,omitempty"`
	From         int64  `protobuf:"varint,7,opt,name=from,proto3" json:"from,omitempty"`
	To           int64  `protobuf:"varint,8,opt,name=to,proto3" json:"to,omitempty"`
	FromRate     *Rate  `protobuf:"bytes,9,opt,name=from_rate,json=fromRate,proto3" json:"from_rate,omitempty"`
	ToRate       *Rate  `protobuf:"bytes,10,opt,name=to_rate,json=toRate,proto3" json:"to_rate,omitempty"`
}

func (x *GrowthRecord) Reset() {
	*x = GrowthRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrowthRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrowthRecord) ProtoMessage() {}

func (x *GrowthRecord) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrowthRecord.ProtoReflect.Descriptor instead.
func (*GrowthRecord) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{1}
}

func (x *GrowthRecord) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GrowthRecord) GetFromRateId() int32 {
	if x != nil {
		return x.FromRateId
	}
	return 0
}

func (x *GrowthRecord) GetToRateId() int32 {
	if x != nil {
		return x.ToRateId
	}
	return 0
}

func (x *GrowthRecord) GetVolumeGrowth() string {
	if x != nil {
		return x.VolumeGrowth
	}
	return ""
}

func (x *GrowthRecord) GetHighGrowth() string {
	if x != nil {
		return x.HighGrowth
	}
	return ""
}

func (x *GrowthRecord) GetLowGrowth() string {
	if x != nil {
		return x.LowGrowth
	}
	return ""
}

func (x *GrowthRecord) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *GrowthRecord) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *GrowthRecord) GetFromRate() *Rate {
	if x != nil {
		return x.FromRate
	}
	return nil
}

func (x *GrowthRecord) GetToRate() *Rate {
	if x != nil {
		return x.ToRate
	}
	return nil
}

// Created a blank get request
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromTimestamp int64 `protobuf:"varint,1,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp   int64 `protobuf:"varint,2,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[2]
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
	return file_rate_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetFromTimestamp() int64 {
	if x != nil {
		return x.FromTimestamp
	}
	return 0
}

func (x *GetRequest) GetToTimestamp() int64 {
	if x != nil {
		return x.ToTimestamp
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrowthRecords []*GrowthRecord `protobuf:"bytes,1,rep,name=growth_records,json=growthRecords,proto3" json:"growth_records,omitempty"`
	Data          []*ResponseData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[3]
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
	return file_rate_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetGrowthRecords() []*GrowthRecord {
	if x != nil {
		return x.GrowthRecords
	}
	return nil
}

func (x *Response) GetData() []*ResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From       string                `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To         string                `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	MarketData []*ResponseMarketData `protobuf:"bytes,3,rep,name=market_data,json=marketData,proto3" json:"market_data,omitempty"`
}

func (x *ResponseData) Reset() {
	*x = ResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseData) ProtoMessage() {}

func (x *ResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseData.ProtoReflect.Descriptor instead.
func (*ResponseData) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseData) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *ResponseData) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *ResponseData) GetMarketData() []*ResponseMarketData {
	if x != nil {
		return x.MarketData
	}
	return nil
}

type ResponseMarketData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MarketPair string      `protobuf:"bytes,1,opt,name=market_pair,json=marketPair,proto3" json:"market_pair,omitempty"`
	GrowthData *GrowthData `protobuf:"bytes,2,opt,name=growth_data,json=growthData,proto3" json:"growth_data,omitempty"`
}

func (x *ResponseMarketData) Reset() {
	*x = ResponseMarketData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMarketData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMarketData) ProtoMessage() {}

func (x *ResponseMarketData) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMarketData.ProtoReflect.Descriptor instead.
func (*ResponseMarketData) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseMarketData) GetMarketPair() string {
	if x != nil {
		return x.MarketPair
	}
	return ""
}

func (x *ResponseMarketData) GetGrowthData() *GrowthData {
	if x != nil {
		return x.GrowthData
	}
	return nil
}

type GrowthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VolumeGrowth string `protobuf:"bytes,1,opt,name=volume_growth,json=volumeGrowth,proto3" json:"volume_growth,omitempty"`
	HighGrowth   string `protobuf:"bytes,2,opt,name=high_growth,json=highGrowth,proto3" json:"high_growth,omitempty"`
	LowGrowth    string `protobuf:"bytes,3,opt,name=low_growth,json=lowGrowth,proto3" json:"low_growth,omitempty"`
}

func (x *GrowthData) Reset() {
	*x = GrowthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrowthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrowthData) ProtoMessage() {}

func (x *GrowthData) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrowthData.ProtoReflect.Descriptor instead.
func (*GrowthData) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{6}
}

func (x *GrowthData) GetVolumeGrowth() string {
	if x != nil {
		return x.VolumeGrowth
	}
	return ""
}

func (x *GrowthData) GetHighGrowth() string {
	if x != nil {
		return x.HighGrowth
	}
	return ""
}

func (x *GrowthData) GetLowGrowth() string {
	if x != nil {
		return x.LowGrowth
	}
	return ""
}

var File_rate_proto protoreflect.FileDescriptor

var file_rate_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x61,
	0x74, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x04, 0x52, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x69, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c,
	0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xb5, 0x02, 0x0a, 0x0c, 0x47, 0x72, 0x6f,
	0x77, 0x74, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x52, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x74,
	0x6f, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x74, 0x6f, 0x52, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x47, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x12, 0x1f,
	0x0a, 0x0b, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x47, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x12,
	0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x77, 0x5f, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x27, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x52, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x07, 0x74,
	0x6f, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72,
	0x61, 0x74, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x06, 0x74, 0x6f, 0x52, 0x61, 0x74, 0x65,
	0x22, 0x56, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x6d, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72,
	0x61, 0x74, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x0d, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6d, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x39, 0x0a, 0x0b, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x61, 0x69, 0x72, 0x12, 0x31, 0x0a,
	0x0b, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x77, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x71, 0x0a, 0x0a, 0x47, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23,
	0x0a, 0x0d, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x47, 0x72, 0x6f,
	0x77, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x67, 0x72, 0x6f, 0x77,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x47, 0x72,
	0x6f, 0x77, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x77, 0x5f, 0x67, 0x72, 0x6f, 0x77,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x6f,
	0x77, 0x74, 0x68, 0x32, 0x45, 0x0a, 0x0b, 0x52, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x10, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x62, 0x69, 0x6b, 0x6f, 0x64, 0x65,
	0x2f, 0x70, 0x65, 0x61, 0x71, 0x2d, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2f,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2d, 0x33, 0x2f, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x3b,
	0x72, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rate_proto_rawDescOnce sync.Once
	file_rate_proto_rawDescData = file_rate_proto_rawDesc
)

func file_rate_proto_rawDescGZIP() []byte {
	file_rate_proto_rawDescOnce.Do(func() {
		file_rate_proto_rawDescData = protoimpl.X.CompressGZIP(file_rate_proto_rawDescData)
	})
	return file_rate_proto_rawDescData
}

var file_rate_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rate_proto_goTypes = []interface{}{
	(*Rate)(nil),               // 0: rate.Rate
	(*GrowthRecord)(nil),       // 1: rate.GrowthRecord
	(*GetRequest)(nil),         // 2: rate.GetRequest
	(*Response)(nil),           // 3: rate.Response
	(*ResponseData)(nil),       // 4: rate.ResponseData
	(*ResponseMarketData)(nil), // 5: rate.ResponseMarketData
	(*GrowthData)(nil),         // 6: rate.GrowthData
}
var file_rate_proto_depIdxs = []int32{
	0, // 0: rate.GrowthRecord.from_rate:type_name -> rate.Rate
	0, // 1: rate.GrowthRecord.to_rate:type_name -> rate.Rate
	1, // 2: rate.Response.growth_records:type_name -> rate.GrowthRecord
	4, // 3: rate.Response.data:type_name -> rate.ResponseData
	5, // 4: rate.ResponseData.market_data:type_name -> rate.ResponseMarketData
	6, // 5: rate.ResponseMarketData.growth_data:type_name -> rate.GrowthData
	2, // 6: rate.RateService.GetGrowthRecords:input_type -> rate.GetRequest
	3, // 7: rate.RateService.GetGrowthRecords:output_type -> rate.Response
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_rate_proto_init() }
func file_rate_proto_init() {
	if File_rate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rate); i {
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
		file_rate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrowthRecord); i {
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
		file_rate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_rate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_rate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseData); i {
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
		file_rate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMarketData); i {
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
		file_rate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrowthData); i {
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
			RawDescriptor: file_rate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rate_proto_goTypes,
		DependencyIndexes: file_rate_proto_depIdxs,
		MessageInfos:      file_rate_proto_msgTypes,
	}.Build()
	File_rate_proto = out.File
	file_rate_proto_rawDesc = nil
	file_rate_proto_goTypes = nil
	file_rate_proto_depIdxs = nil
}
