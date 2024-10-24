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

	Source string                        `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Volume *SubscribeRequestFilterVolume `protobuf:"bytes,2,opt,name=volume,proto3" json:"volume,omitempty"`
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

func (x *SubscribeRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *SubscribeRequest) GetVolume() *SubscribeRequestFilterVolume {
	if x != nil {
		return x.Volume
	}
	return nil
}

type SubscribeRequestFilterVolume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Condition:
	//
	//	*SubscribeRequestFilterVolume_Above
	//	*SubscribeRequestFilterVolume_Below
	//	*SubscribeRequestFilterVolume_Between
	Condition isSubscribeRequestFilterVolume_Condition `protobuf_oneof:"condition"`
}

func (x *SubscribeRequestFilterVolume) Reset() {
	*x = SubscribeRequestFilterVolume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequestFilterVolume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequestFilterVolume) ProtoMessage() {}

func (x *SubscribeRequestFilterVolume) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SubscribeRequestFilterVolume.ProtoReflect.Descriptor instead.
func (*SubscribeRequestFilterVolume) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{1}
}

func (m *SubscribeRequestFilterVolume) GetCondition() isSubscribeRequestFilterVolume_Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (x *SubscribeRequestFilterVolume) GetAbove() *VolumeAbove {
	if x, ok := x.GetCondition().(*SubscribeRequestFilterVolume_Above); ok {
		return x.Above
	}
	return nil
}

func (x *SubscribeRequestFilterVolume) GetBelow() *VolumeBelow {
	if x, ok := x.GetCondition().(*SubscribeRequestFilterVolume_Below); ok {
		return x.Below
	}
	return nil
}

func (x *SubscribeRequestFilterVolume) GetBetween() *VolumeBetween {
	if x, ok := x.GetCondition().(*SubscribeRequestFilterVolume_Between); ok {
		return x.Between
	}
	return nil
}

type isSubscribeRequestFilterVolume_Condition interface {
	isSubscribeRequestFilterVolume_Condition()
}

type SubscribeRequestFilterVolume_Above struct {
	Above *VolumeAbove `protobuf:"bytes,1,opt,name=above,proto3,oneof"`
}

type SubscribeRequestFilterVolume_Below struct {
	Below *VolumeBelow `protobuf:"bytes,2,opt,name=below,proto3,oneof"`
}

type SubscribeRequestFilterVolume_Between struct {
	Between *VolumeBetween `protobuf:"bytes,3,opt,name=between,proto3,oneof"`
}

func (*SubscribeRequestFilterVolume_Above) isSubscribeRequestFilterVolume_Condition() {}

func (*SubscribeRequestFilterVolume_Below) isSubscribeRequestFilterVolume_Condition() {}

func (*SubscribeRequestFilterVolume_Between) isSubscribeRequestFilterVolume_Condition() {}

type VolumeAbove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Threshold float32 `protobuf:"fixed32,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
}

func (x *VolumeAbove) Reset() {
	*x = VolumeAbove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeAbove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeAbove) ProtoMessage() {}

func (x *VolumeAbove) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use VolumeAbove.ProtoReflect.Descriptor instead.
func (*VolumeAbove) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{2}
}

func (x *VolumeAbove) GetThreshold() float32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

type VolumeBelow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Threshold float32 `protobuf:"fixed32,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
}

func (x *VolumeBelow) Reset() {
	*x = VolumeBelow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeBelow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeBelow) ProtoMessage() {}

func (x *VolumeBelow) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use VolumeBelow.ProtoReflect.Descriptor instead.
func (*VolumeBelow) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{3}
}

func (x *VolumeBelow) GetThreshold() float32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

type VolumeBetween struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min float32 `protobuf:"fixed32,1,opt,name=min,proto3" json:"min,omitempty"`
	Max float32 `protobuf:"fixed32,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *VolumeBetween) Reset() {
	*x = VolumeBetween{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeBetween) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeBetween) ProtoMessage() {}

func (x *VolumeBetween) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use VolumeBetween.ProtoReflect.Descriptor instead.
func (*VolumeBetween) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{4}
}

func (x *VolumeBetween) GetMin() float32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *VolumeBetween) GetMax() float32 {
	if x != nil {
		return x.Max
	}
	return 0
}

type SubscribeUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UpdateOneof:
	//
	//	*SubscribeUpdate_Volume
	UpdateOneof isSubscribeUpdate_UpdateOneof `protobuf_oneof:"update_oneof"`
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

func (m *SubscribeUpdate) GetUpdateOneof() isSubscribeUpdate_UpdateOneof {
	if m != nil {
		return m.UpdateOneof
	}
	return nil
}

func (x *SubscribeUpdate) GetVolume() *SubscribeUpdateVolume {
	if x, ok := x.GetUpdateOneof().(*SubscribeUpdate_Volume); ok {
		return x.Volume
	}
	return nil
}

type isSubscribeUpdate_UpdateOneof interface {
	isSubscribeUpdate_UpdateOneof()
}

type SubscribeUpdate_Volume struct {
	Volume *SubscribeUpdateVolume `protobuf:"bytes,2,opt,name=volume,proto3,oneof"`
}

func (*SubscribeUpdate_Volume) isSubscribeUpdate_UpdateOneof() {}

type SubscribeUpdateVolume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmmId        string  `protobuf:"bytes,1,opt,name=amm_id,json=ammId,proto3" json:"amm_id,omitempty"`
	Mint         string  `protobuf:"bytes,2,opt,name=mint,proto3" json:"mint,omitempty"`
	Name         string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Ticker       string  `protobuf:"bytes,4,opt,name=ticker,proto3" json:"ticker,omitempty"`
	BaseVault    float64 `protobuf:"fixed64,5,opt,name=base_vault,json=baseVault,proto3" json:"base_vault,omitempty"`
	TokenPrice   float64 `protobuf:"fixed64,6,opt,name=token_price,json=tokenPrice,proto3" json:"token_price,omitempty"`
	Volume_1M    float64 `protobuf:"fixed64,7,opt,name=volume_1m,json=volume1m,proto3" json:"volume_1m,omitempty"`
	IsPumpFun    bool    `protobuf:"varint,8,opt,name=is_pump_fun,json=isPumpFun,proto3" json:"is_pump_fun,omitempty"`
	TriggerValue int64   `protobuf:"varint,9,opt,name=trigger_value,json=triggerValue,proto3" json:"trigger_value,omitempty"`
}

func (x *SubscribeUpdateVolume) Reset() {
	*x = SubscribeUpdateVolume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeUpdateVolume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeUpdateVolume) ProtoMessage() {}

func (x *SubscribeUpdateVolume) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SubscribeUpdateVolume.ProtoReflect.Descriptor instead.
func (*SubscribeUpdateVolume) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{6}
}

func (x *SubscribeUpdateVolume) GetAmmId() string {
	if x != nil {
		return x.AmmId
	}
	return ""
}

func (x *SubscribeUpdateVolume) GetMint() string {
	if x != nil {
		return x.Mint
	}
	return ""
}

func (x *SubscribeUpdateVolume) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubscribeUpdateVolume) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *SubscribeUpdateVolume) GetBaseVault() float64 {
	if x != nil {
		return x.BaseVault
	}
	return 0
}

func (x *SubscribeUpdateVolume) GetTokenPrice() float64 {
	if x != nil {
		return x.TokenPrice
	}
	return 0
}

func (x *SubscribeUpdateVolume) GetVolume_1M() float64 {
	if x != nil {
		return x.Volume_1M
	}
	return 0
}

func (x *SubscribeUpdateVolume) GetIsPumpFun() bool {
	if x != nil {
		return x.IsPumpFun
	}
	return false
}

func (x *SubscribeUpdateVolume) GetTriggerValue() int64 {
	if x != nil {
		return x.TriggerValue
	}
	return 0
}

type AnomalyTxn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmmId        string `protobuf:"bytes,1,opt,name=amm_id,json=ammId,proto3" json:"amm_id,omitempty"`
	Mint         string `protobuf:"bytes,2,opt,name=mint,proto3" json:"mint,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Ticker       string `protobuf:"bytes,4,opt,name=ticker,proto3" json:"ticker,omitempty"`
	BaseVault    string `protobuf:"bytes,5,opt,name=base_vault,json=baseVault,proto3" json:"base_vault,omitempty"`
	TokenPrice   string `protobuf:"bytes,6,opt,name=token_price,json=tokenPrice,proto3" json:"token_price,omitempty"`
	Volume_1M    string `protobuf:"bytes,7,opt,name=volume_1m,json=volume1m,proto3" json:"volume_1m,omitempty"`
	IsPumpFun    string `protobuf:"bytes,8,opt,name=is_pump_fun,json=isPumpFun,proto3" json:"is_pump_fun,omitempty"`
	TriggerValue string `protobuf:"bytes,9,opt,name=trigger_value,json=triggerValue,proto3" json:"trigger_value,omitempty"`
	Signature    string `protobuf:"bytes,10,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *AnomalyTxn) Reset() {
	*x = AnomalyTxn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnomalyTxn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnomalyTxn) ProtoMessage() {}

func (x *AnomalyTxn) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AnomalyTxn.ProtoReflect.Descriptor instead.
func (*AnomalyTxn) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{7}
}

func (x *AnomalyTxn) GetAmmId() string {
	if x != nil {
		return x.AmmId
	}
	return ""
}

func (x *AnomalyTxn) GetMint() string {
	if x != nil {
		return x.Mint
	}
	return ""
}

func (x *AnomalyTxn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AnomalyTxn) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *AnomalyTxn) GetBaseVault() string {
	if x != nil {
		return x.BaseVault
	}
	return ""
}

func (x *AnomalyTxn) GetTokenPrice() string {
	if x != nil {
		return x.TokenPrice
	}
	return ""
}

func (x *AnomalyTxn) GetVolume_1M() string {
	if x != nil {
		return x.Volume_1M
	}
	return ""
}

func (x *AnomalyTxn) GetIsPumpFun() string {
	if x != nil {
		return x.IsPumpFun
	}
	return ""
}

func (x *AnomalyTxn) GetTriggerValue() string {
	if x != nil {
		return x.TriggerValue
	}
	return ""
}

func (x *AnomalyTxn) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_anomaly_proto protoreflect.FileDescriptor

var file_anomaly_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x22, 0x67, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22,
	0xb5, 0x01, 0x0a, 0x1c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x05, 0x61, 0x62, 0x6f, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41, 0x62,
	0x6f, 0x76, 0x65, 0x48, 0x00, 0x52, 0x05, 0x61, 0x62, 0x6f, 0x76, 0x65, 0x12, 0x2a, 0x0a, 0x05,
	0x62, 0x65, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x6d, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x65, 0x6c, 0x6f, 0x77, 0x48,
	0x00, 0x52, 0x05, 0x62, 0x65, 0x6c, 0x6f, 0x77, 0x12, 0x30, 0x0a, 0x07, 0x62, 0x65, 0x74, 0x77,
	0x65, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x6d, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x48,
	0x00, 0x52, 0x07, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x0b, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x41, 0x62, 0x6f, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x22, 0x2b, 0x0a, 0x0b, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x65,
	0x6c, 0x6f, 0x77, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x22, 0x33, 0x0a, 0x0d, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x65, 0x74, 0x77, 0x65,
	0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x59, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x6d, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x22, 0x90, 0x02, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61,
	0x6d, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6d, 0x6d,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x56, 0x61, 0x75, 0x6c,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x31, 0x6d, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x31, 0x6d, 0x12,
	0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x6d, 0x70, 0x5f, 0x66, 0x75, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x75, 0x6d, 0x70, 0x46, 0x75, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xa3, 0x02, 0x0a, 0x0a, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79,
	0x54, 0x78, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x6d, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6d, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x61, 0x73, 0x65, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x31, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x31, 0x6d, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x70, 0x75,
	0x6d, 0x70, 0x5f, 0x66, 0x75, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73,
	0x50, 0x75, 0x6d, 0x70, 0x46, 0x75, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0x4d, 0x0a, 0x07, 0x41, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x12, 0x42, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x6d, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_anomaly_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_anomaly_proto_goTypes = []interface{}{
	(*SubscribeRequest)(nil),             // 0: solom.SubscribeRequest
	(*SubscribeRequestFilterVolume)(nil), // 1: solom.SubscribeRequestFilterVolume
	(*VolumeAbove)(nil),                  // 2: solom.VolumeAbove
	(*VolumeBelow)(nil),                  // 3: solom.VolumeBelow
	(*VolumeBetween)(nil),                // 4: solom.VolumeBetween
	(*SubscribeUpdate)(nil),              // 5: solom.SubscribeUpdate
	(*SubscribeUpdateVolume)(nil),        // 6: solom.SubscribeUpdateVolume
	(*AnomalyTxn)(nil),                   // 7: solom.AnomalyTxn
}
var file_anomaly_proto_depIdxs = []int32{
	1, // 0: solom.SubscribeRequest.volume:type_name -> solom.SubscribeRequestFilterVolume
	2, // 1: solom.SubscribeRequestFilterVolume.above:type_name -> solom.VolumeAbove
	3, // 2: solom.SubscribeRequestFilterVolume.below:type_name -> solom.VolumeBelow
	4, // 3: solom.SubscribeRequestFilterVolume.between:type_name -> solom.VolumeBetween
	6, // 4: solom.SubscribeUpdate.volume:type_name -> solom.SubscribeUpdateVolume
	0, // 5: solom.Anomaly.Subscribe:input_type -> solom.SubscribeRequest
	5, // 6: solom.Anomaly.Subscribe:output_type -> solom.SubscribeUpdate
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
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
			switch v := v.(*SubscribeRequestFilterVolume); i {
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
			switch v := v.(*VolumeAbove); i {
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
			switch v := v.(*VolumeBelow); i {
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
			switch v := v.(*VolumeBetween); i {
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
			switch v := v.(*SubscribeUpdateVolume); i {
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
			switch v := v.(*AnomalyTxn); i {
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
	file_anomaly_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SubscribeRequestFilterVolume_Above)(nil),
		(*SubscribeRequestFilterVolume_Below)(nil),
		(*SubscribeRequestFilterVolume_Between)(nil),
	}
	file_anomaly_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*SubscribeUpdate_Volume)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anomaly_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
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
