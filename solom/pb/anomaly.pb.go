// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: anomaly.proto

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

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SolBalance  *SubscribeRequestFilterSolBalance        `protobuf:"bytes,1,opt,name=sol_balance,json=solBalance,proto3" json:"sol_balance,omitempty"`
	Pumpfun     *SubscribeRequestFilterPumpFun           `protobuf:"bytes,2,opt,name=pumpfun,proto3" json:"pumpfun,omitempty"`
	Transaction *SubscribeRequestFilterTransactionFilter `protobuf:"bytes,3,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Source      *string                                  `protobuf:"bytes,4,opt,name=source,proto3,oneof" json:"source,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeRequest) GetSolBalance() *SubscribeRequestFilterSolBalance {
	if x != nil {
		return x.SolBalance
	}
	return nil
}

func (x *SubscribeRequest) GetPumpfun() *SubscribeRequestFilterPumpFun {
	if x != nil {
		return x.Pumpfun
	}
	return nil
}

func (x *SubscribeRequest) GetTransaction() *SubscribeRequestFilterTransactionFilter {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *SubscribeRequest) GetSource() string {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return ""
}

type SubscribeRequestFilterSolBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min *uint64 `protobuf:"varint,1,opt,name=min,proto3,oneof" json:"min,omitempty"`
	Max *uint64 `protobuf:"varint,2,opt,name=max,proto3,oneof" json:"max,omitempty"`
}

func (x *SubscribeRequestFilterSolBalance) Reset() {
	*x = SubscribeRequestFilterSolBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequestFilterSolBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequestFilterSolBalance) ProtoMessage() {}

func (x *SubscribeRequestFilterSolBalance) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequestFilterSolBalance.ProtoReflect.Descriptor instead.
func (*SubscribeRequestFilterSolBalance) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{1}
}

func (x *SubscribeRequestFilterSolBalance) GetMin() uint64 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *SubscribeRequestFilterSolBalance) GetMax() uint64 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

type SubscribeRequestFilterPumpFun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsPumpfun *bool `protobuf:"varint,1,opt,name=is_pumpfun,json=isPumpfun,proto3,oneof" json:"is_pumpfun,omitempty"`
}

func (x *SubscribeRequestFilterPumpFun) Reset() {
	*x = SubscribeRequestFilterPumpFun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequestFilterPumpFun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequestFilterPumpFun) ProtoMessage() {}

func (x *SubscribeRequestFilterPumpFun) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequestFilterPumpFun.ProtoReflect.Descriptor instead.
func (*SubscribeRequestFilterPumpFun) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeRequestFilterPumpFun) GetIsPumpfun() bool {
	if x != nil && x.IsPumpfun != nil {
		return *x.IsPumpfun
	}
	return false
}

type SubscribeRequestFilterTransactionFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mint   []string `protobuf:"bytes,1,rep,name=mint,proto3" json:"mint,omitempty"`
	Signer []string `protobuf:"bytes,2,rep,name=signer,proto3" json:"signer,omitempty"`
	Min    *uint64  `protobuf:"varint,3,opt,name=min,proto3,oneof" json:"min,omitempty"`
	Max    *uint64  `protobuf:"varint,4,opt,name=max,proto3,oneof" json:"max,omitempty"`
}

func (x *SubscribeRequestFilterTransactionFilter) Reset() {
	*x = SubscribeRequestFilterTransactionFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequestFilterTransactionFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequestFilterTransactionFilter) ProtoMessage() {}

func (x *SubscribeRequestFilterTransactionFilter) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequestFilterTransactionFilter.ProtoReflect.Descriptor instead.
func (*SubscribeRequestFilterTransactionFilter) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeRequestFilterTransactionFilter) GetMint() []string {
	if x != nil {
		return x.Mint
	}
	return nil
}

func (x *SubscribeRequestFilterTransactionFilter) GetSigner() []string {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *SubscribeRequestFilterTransactionFilter) GetMin() uint64 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *SubscribeRequestFilterTransactionFilter) GetMax() uint64 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

type SubscribeRequestFilterOHLC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmmId        string `protobuf:"bytes,1,opt,name=amm_id,json=ammId,proto3" json:"amm_id,omitempty"`
	TimeInterval int64  `protobuf:"varint,2,opt,name=time_interval,json=timeInterval,proto3" json:"time_interval,omitempty"`
}

func (x *SubscribeRequestFilterOHLC) Reset() {
	*x = SubscribeRequestFilterOHLC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequestFilterOHLC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequestFilterOHLC) ProtoMessage() {}

func (x *SubscribeRequestFilterOHLC) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequestFilterOHLC.ProtoReflect.Descriptor instead.
func (*SubscribeRequestFilterOHLC) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{4}
}

func (x *SubscribeRequestFilterOHLC) GetAmmId() string {
	if x != nil {
		return x.AmmId
	}
	return ""
}

func (x *SubscribeRequestFilterOHLC) GetTimeInterval() int64 {
	if x != nil {
		return x.TimeInterval
	}
	return 0
}

type SubscribeUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dex       string `protobuf:"bytes,1,opt,name=dex,proto3" json:"dex,omitempty"`
	AmmId     string `protobuf:"bytes,2,opt,name=amm_id,json=ammId,proto3" json:"amm_id,omitempty"`
	Mint      string `protobuf:"bytes,3,opt,name=mint,proto3" json:"mint,omitempty"`
	Signature string `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SubscribeUpdate) Reset() {
	*x = SubscribeUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeUpdate) ProtoMessage() {}

func (x *SubscribeUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeUpdate.ProtoReflect.Descriptor instead.
func (*SubscribeUpdate) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{5}
}

func (x *SubscribeUpdate) GetDex() string {
	if x != nil {
		return x.Dex
	}
	return ""
}

func (x *SubscribeUpdate) GetAmmId() string {
	if x != nil {
		return x.AmmId
	}
	return ""
}

func (x *SubscribeUpdate) GetMint() string {
	if x != nil {
		return x.Mint
	}
	return ""
}

func (x *SubscribeUpdate) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type AmmId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmmId string `protobuf:"bytes,1,opt,name=amm_id,json=ammId,proto3" json:"amm_id,omitempty"`
}

func (x *AmmId) Reset() {
	*x = AmmId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmmId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmmId) ProtoMessage() {}

func (x *AmmId) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmmId.ProtoReflect.Descriptor instead.
func (*AmmId) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{6}
}

func (x *AmmId) GetAmmId() string {
	if x != nil {
		return x.AmmId
	}
	return ""
}

type PriceAllWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price map[string]float64 `protobuf:"bytes,1,rep,name=price,proto3" json:"price,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *PriceAllWindow) Reset() {
	*x = PriceAllWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceAllWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceAllWindow) ProtoMessage() {}

func (x *PriceAllWindow) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceAllWindow.ProtoReflect.Descriptor instead.
func (*PriceAllWindow) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{7}
}

func (x *PriceAllWindow) GetPrice() map[string]float64 {
	if x != nil {
		return x.Price
	}
	return nil
}

type ActionMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume map[string]int64 `protobuf:"bytes,1,rep,name=volume,proto3" json:"volume,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ActionMap) Reset() {
	*x = ActionMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionMap) ProtoMessage() {}

func (x *ActionMap) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionMap.ProtoReflect.Descriptor instead.
func (*ActionMap) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{8}
}

func (x *ActionMap) GetVolume() map[string]int64 {
	if x != nil {
		return x.Volume
	}
	return nil
}

type VolumeAllWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume map[string]*ActionMap `protobuf:"bytes,1,rep,name=volume,proto3" json:"volume,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VolumeAllWindow) Reset() {
	*x = VolumeAllWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeAllWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeAllWindow) ProtoMessage() {}

func (x *VolumeAllWindow) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeAllWindow.ProtoReflect.Descriptor instead.
func (*VolumeAllWindow) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{9}
}

func (x *VolumeAllWindow) GetVolume() map[string]*ActionMap {
	if x != nil {
		return x.Volume
	}
	return nil
}

var File_anomaly_proto protoreflect.FileDescriptor

var file_anomaly_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x22, 0x96, 0x02, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0b, 0x73,
	0x6f, 0x6c, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53,
	0x6f, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x73, 0x6f, 0x6c, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x70, 0x75, 0x6d, 0x70, 0x66, 0x75, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x75, 0x6d, 0x70, 0x46, 0x75, 0x6e, 0x52, 0x07, 0x70, 0x75,
	0x6d, 0x70, 0x66, 0x75, 0x6e, 0x12, 0x50, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0x60, 0x0a, 0x20, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x6f, 0x6c, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x48, 0x00, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x61,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x88, 0x01,
	0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x69, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x61,
	0x78, 0x22, 0x52, 0x0a, 0x1d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x75, 0x6d, 0x70, 0x46,
	0x75, 0x6e, 0x12, 0x22, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x6d, 0x70, 0x66, 0x75, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x69, 0x73, 0x50, 0x75, 0x6d, 0x70,
	0x66, 0x75, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x73, 0x5f, 0x70, 0x75,
	0x6d, 0x70, 0x66, 0x75, 0x6e, 0x22, 0x93, 0x01, 0x0a, 0x27, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x15, 0x0a,
	0x03, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x69,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x48, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f,
	0x6d, 0x69, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x61, 0x78, 0x22, 0x58, 0x0a, 0x1a, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x48, 0x4c, 0x43, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x6d, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6d, 0x6d, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x6c, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x65, 0x78, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x6d,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6d, 0x6d, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x1e, 0x0a, 0x05, 0x41, 0x6d, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x61, 0x6d, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6d,
	0x6d, 0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x41, 0x6c, 0x6c,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x36, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x41, 0x6c, 0x6c, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x38,
	0x0a, 0x0a, 0x50, 0x72, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7c, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x12, 0x34, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9a, 0x01, 0x0a, 0x0f, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x41, 0x6c, 0x6c, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x3a, 0x0a, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x6d, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41, 0x6c, 0x6c, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x1a, 0x4b, 0x0a, 0x0b, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0xc7, 0x01, 0x0a, 0x07, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x12,
	0x42, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x17, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x41,
	0x6c, 0x6c, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x0c, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d,
	0x2e, 0x41, 0x6d, 0x6d, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x41, 0x6c, 0x6c, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41, 0x6c, 0x6c, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x0c, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x41, 0x6d,
	0x6d, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x41, 0x6c, 0x6c, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x22, 0x00, 0x42, 0x0b, 0x5a,
	0x09, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_anomaly_proto_rawDescOnce sync.Once
	file_anomaly_proto_rawDescData = file_anomaly_proto_rawDesc
)

func file_anomaly_proto_rawDescGZIP() []byte {
	file_anomaly_proto_rawDescOnce.Do(func() {
		file_anomaly_proto_rawDescData = protoimpl.X.CompressGZIP(file_anomaly_proto_rawDescData)
	})
	return file_anomaly_proto_rawDescData
}

var file_anomaly_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_anomaly_proto_goTypes = []interface{}{
	(*SubscribeRequest)(nil),                        // 0: solom.SubscribeRequest
	(*SubscribeRequestFilterSolBalance)(nil),        // 1: solom.SubscribeRequestFilterSolBalance
	(*SubscribeRequestFilterPumpFun)(nil),           // 2: solom.SubscribeRequestFilterPumpFun
	(*SubscribeRequestFilterTransactionFilter)(nil), // 3: solom.SubscribeRequestFilterTransactionFilter
	(*SubscribeRequestFilterOHLC)(nil),              // 4: solom.SubscribeRequestFilterOHLC
	(*SubscribeUpdate)(nil),                         // 5: solom.SubscribeUpdate
	(*AmmId)(nil),                                   // 6: solom.AmmId
	(*PriceAllWindow)(nil),                          // 7: solom.PriceAllWindow
	(*ActionMap)(nil),                               // 8: solom.ActionMap
	(*VolumeAllWindow)(nil),                         // 9: solom.VolumeAllWindow
	nil,                                             // 10: solom.PriceAllWindow.PriceEntry
	nil,                                             // 11: solom.ActionMap.VolumeEntry
	nil,                                             // 12: solom.VolumeAllWindow.VolumeEntry
}
var file_anomaly_proto_depIdxs = []int32{
	1,  // 0: solom.SubscribeRequest.sol_balance:type_name -> solom.SubscribeRequestFilterSolBalance
	2,  // 1: solom.SubscribeRequest.pumpfun:type_name -> solom.SubscribeRequestFilterPumpFun
	3,  // 2: solom.SubscribeRequest.transaction:type_name -> solom.SubscribeRequestFilterTransactionFilter
	10, // 3: solom.PriceAllWindow.price:type_name -> solom.PriceAllWindow.PriceEntry
	11, // 4: solom.ActionMap.volume:type_name -> solom.ActionMap.VolumeEntry
	12, // 5: solom.VolumeAllWindow.volume:type_name -> solom.VolumeAllWindow.VolumeEntry
	8,  // 6: solom.VolumeAllWindow.VolumeEntry.value:type_name -> solom.ActionMap
	0,  // 7: solom.Anomaly.Subscribe:input_type -> solom.SubscribeRequest
	6,  // 8: solom.Anomaly.GetPriceAllWindow:input_type -> solom.AmmId
	6,  // 9: solom.Anomaly.GetVolumeAllWindow:input_type -> solom.AmmId
	5,  // 10: solom.Anomaly.Subscribe:output_type -> solom.SubscribeUpdate
	7,  // 11: solom.Anomaly.GetPriceAllWindow:output_type -> solom.PriceAllWindow
	9,  // 12: solom.Anomaly.GetVolumeAllWindow:output_type -> solom.VolumeAllWindow
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_anomaly_proto_init() }
func file_anomaly_proto_init() {
	if File_anomaly_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_anomaly_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_anomaly_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequestFilterSolBalance); i {
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
		file_anomaly_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequestFilterPumpFun); i {
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
		file_anomaly_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequestFilterTransactionFilter); i {
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
		file_anomaly_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequestFilterOHLC); i {
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
		file_anomaly_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeUpdate); i {
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
		file_anomaly_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmmId); i {
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
		file_anomaly_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceAllWindow); i {
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
		file_anomaly_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionMap); i {
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
		file_anomaly_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeAllWindow); i {
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
	file_anomaly_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_anomaly_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_anomaly_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_anomaly_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anomaly_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_anomaly_proto_goTypes,
		DependencyIndexes: file_anomaly_proto_depIdxs,
		MessageInfos:      file_anomaly_proto_msgTypes,
	}.Build()
	File_anomaly_proto = out.File
	file_anomaly_proto_rawDesc = nil
	file_anomaly_proto_goTypes = nil
	file_anomaly_proto_depIdxs = nil
}
