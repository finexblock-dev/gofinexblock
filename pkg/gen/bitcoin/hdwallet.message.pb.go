// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: pkg/proto/bitcoin/hdwallet.message.proto

package bitcoin

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

type UnspentOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash        string  `protobuf:"bytes,1,opt,name=TxHash,json=tx_hash,proto3" json:"TxHash,omitempty"`
	Address       string  `protobuf:"bytes,2,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
	Account       string  `protobuf:"bytes,3,opt,name=Account,json=account,proto3" json:"Account,omitempty"`
	Amount        float64 `protobuf:"fixed64,4,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
	Confirmations int64   `protobuf:"varint,5,opt,name=Confirmations,json=confirmations,proto3" json:"Confirmations,omitempty"`
	Spendable     bool    `protobuf:"varint,6,opt,name=Spendable,json=spendable,proto3" json:"Spendable,omitempty"`
}

func (x *UnspentOutput) Reset() {
	*x = UnspentOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnspentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnspentOutput) ProtoMessage() {}

func (x *UnspentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnspentOutput.ProtoReflect.Descriptor instead.
func (*UnspentOutput) Descriptor() ([]byte, []int) {
	return file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescGZIP(), []int{0}
}

func (x *UnspentOutput) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *UnspentOutput) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UnspentOutput) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UnspentOutput) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UnspentOutput) GetConfirmations() int64 {
	if x != nil {
		return x.Confirmations
	}
	return 0
}

func (x *UnspentOutput) GetSpendable() bool {
	if x != nil {
		return x.Spendable
	}
	return false
}

type GetNewAddressInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNewAddressInput) Reset() {
	*x = GetNewAddressInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewAddressInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewAddressInput) ProtoMessage() {}

func (x *GetNewAddressInput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewAddressInput.ProtoReflect.Descriptor instead.
func (*GetNewAddressInput) Descriptor() ([]byte, []int) {
	return file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescGZIP(), []int{1}
}

type GetNewAddressOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
}

func (x *GetNewAddressOutput) Reset() {
	*x = GetNewAddressOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewAddressOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewAddressOutput) ProtoMessage() {}

func (x *GetNewAddressOutput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewAddressOutput.ProtoReflect.Descriptor instead.
func (*GetNewAddressOutput) Descriptor() ([]byte, []int) {
	return file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescGZIP(), []int{2}
}

func (x *GetNewAddressOutput) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ListUnspentInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinConf uint64 `protobuf:"varint,1,opt,name=MinConf,json=minconf,proto3" json:"MinConf,omitempty"`
	MaxConf uint64 `protobuf:"varint,2,opt,name=MaxConf,json=maxconf,proto3" json:"MaxConf,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=Address,json=addresses,proto3" json:"Address,omitempty"`
}

func (x *ListUnspentInput) Reset() {
	*x = ListUnspentInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUnspentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUnspentInput) ProtoMessage() {}

func (x *ListUnspentInput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUnspentInput.ProtoReflect.Descriptor instead.
func (*ListUnspentInput) Descriptor() ([]byte, []int) {
	return file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescGZIP(), []int{3}
}

func (x *ListUnspentInput) GetMinConf() uint64 {
	if x != nil {
		return x.MinConf
	}
	return 0
}

func (x *ListUnspentInput) GetMaxConf() uint64 {
	if x != nil {
		return x.MaxConf
	}
	return 0
}

func (x *ListUnspentInput) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ListUnspentOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnspentOutputs []*UnspentOutput `protobuf:"bytes,1,rep,name=UnspentOutputs,json=unspent_outputs,proto3" json:"UnspentOutputs,omitempty"`
}

func (x *ListUnspentOutput) Reset() {
	*x = ListUnspentOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUnspentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUnspentOutput) ProtoMessage() {}

func (x *ListUnspentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUnspentOutput.ProtoReflect.Descriptor instead.
func (*ListUnspentOutput) Descriptor() ([]byte, []int) {
	return file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescGZIP(), []int{4}
}

func (x *ListUnspentOutput) GetUnspentOutputs() []*UnspentOutput {
	if x != nil {
		return x.UnspentOutputs
	}
	return nil
}

type SendToAddressInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToAddress string  `protobuf:"bytes,1,opt,name=ToAddress,json=to_address,proto3" json:"ToAddress,omitempty"`
	Amount    float64 `protobuf:"fixed64,2,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
}

func (x *SendToAddressInput) Reset() {
	*x = SendToAddressInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendToAddressInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendToAddressInput) ProtoMessage() {}

func (x *SendToAddressInput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendToAddressInput.ProtoReflect.Descriptor instead.
func (*SendToAddressInput) Descriptor() ([]byte, []int) {
	return file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescGZIP(), []int{5}
}

func (x *SendToAddressInput) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *SendToAddressInput) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type SendToAddressOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash string `protobuf:"bytes,1,opt,name=TxHash,json=tx_hash,proto3" json:"TxHash,omitempty"`
}

func (x *SendToAddressOutput) Reset() {
	*x = SendToAddressOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendToAddressOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendToAddressOutput) ProtoMessage() {}

func (x *SendToAddressOutput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendToAddressOutput.ProtoReflect.Descriptor instead.
func (*SendToAddressOutput) Descriptor() ([]byte, []int) {
	return file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescGZIP(), []int{6}
}

func (x *SendToAddressOutput) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

var File_pkg_proto_bitcoin_hdwallet_message_proto protoreflect.FileDescriptor

var file_pkg_proto_bitcoin_hdwallet_message_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x74, 0x63,
	0x6f, 0x69, 0x6e, 0x2f, 0x68, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x69, 0x74, 0x63,
	0x6f, 0x69, 0x6e, 0x22, 0xb8, 0x01, 0x0a, 0x0d, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x12, 0x18,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x14,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x22, 0x2f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x62, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x6e, 0x73,
	0x70, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x63,
	0x6f, 0x6e, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x63, 0x6f, 0x6e, 0x66, 0x12, 0x1a, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x3f,
	0x0a, 0x0e, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0f,
	0x75, 0x6e, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22,
	0x4b, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x13,
	0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x42, 0x0a, 0x5a, 0x08,
	0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescOnce sync.Once
	file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescData = file_pkg_proto_bitcoin_hdwallet_message_proto_rawDesc
)

func file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescGZIP() []byte {
	file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescOnce.Do(func() {
		file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescData)
	})
	return file_pkg_proto_bitcoin_hdwallet_message_proto_rawDescData
}

var file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_proto_bitcoin_hdwallet_message_proto_goTypes = []interface{}{
	(*UnspentOutput)(nil),       // 0: bitcoin.UnspentOutput
	(*GetNewAddressInput)(nil),  // 1: bitcoin.GetNewAddressInput
	(*GetNewAddressOutput)(nil), // 2: bitcoin.GetNewAddressOutput
	(*ListUnspentInput)(nil),    // 3: bitcoin.ListUnspentInput
	(*ListUnspentOutput)(nil),   // 4: bitcoin.ListUnspentOutput
	(*SendToAddressInput)(nil),  // 5: bitcoin.SendToAddressInput
	(*SendToAddressOutput)(nil), // 6: bitcoin.SendToAddressOutput
}
var file_pkg_proto_bitcoin_hdwallet_message_proto_depIdxs = []int32{
	0, // 0: bitcoin.ListUnspentOutput.UnspentOutputs:type_name -> bitcoin.UnspentOutput
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_bitcoin_hdwallet_message_proto_init() }
func file_pkg_proto_bitcoin_hdwallet_message_proto_init() {
	if File_pkg_proto_bitcoin_hdwallet_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnspentOutput); i {
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
		file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewAddressInput); i {
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
		file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewAddressOutput); i {
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
		file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUnspentInput); i {
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
		file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUnspentOutput); i {
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
		file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendToAddressInput); i {
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
		file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendToAddressOutput); i {
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
			RawDescriptor: file_pkg_proto_bitcoin_hdwallet_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_bitcoin_hdwallet_message_proto_goTypes,
		DependencyIndexes: file_pkg_proto_bitcoin_hdwallet_message_proto_depIdxs,
		MessageInfos:      file_pkg_proto_bitcoin_hdwallet_message_proto_msgTypes,
	}.Build()
	File_pkg_proto_bitcoin_hdwallet_message_proto = out.File
	file_pkg_proto_bitcoin_hdwallet_message_proto_rawDesc = nil
	file_pkg_proto_bitcoin_hdwallet_message_proto_goTypes = nil
	file_pkg_proto_bitcoin_hdwallet_message_proto_depIdxs = nil
}
