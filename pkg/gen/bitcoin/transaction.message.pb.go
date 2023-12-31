// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: pkg/proto/bitcoin/transaction.message.proto

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

type GetRawTransactionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId string `protobuf:"bytes,1,opt,name=TxId,json=tx_id,proto3" json:"TxId,omitempty"`
}

func (x *GetRawTransactionInput) Reset() {
	*x = GetRawTransactionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_bitcoin_transaction_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRawTransactionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRawTransactionInput) ProtoMessage() {}

func (x *GetRawTransactionInput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_bitcoin_transaction_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRawTransactionInput.ProtoReflect.Descriptor instead.
func (*GetRawTransactionInput) Descriptor() ([]byte, []int) {
	return file_pkg_proto_bitcoin_transaction_message_proto_rawDescGZIP(), []int{0}
}

func (x *GetRawTransactionInput) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

type GetRawTransactionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hex           string  `protobuf:"bytes,1,opt,name=Hex,json=hex,proto3" json:"Hex,omitempty"`
	TxId          string  `protobuf:"bytes,2,opt,name=TxId,json=tx_id,proto3" json:"TxId,omitempty"`
	Hash          string  `protobuf:"bytes,3,opt,name=Hash,json=hash,proto3" json:"Hash,omitempty"`
	Size          int32   `protobuf:"varint,4,opt,name=Size,json=size,proto3" json:"Size,omitempty"`
	Vsize         int32   `protobuf:"varint,5,opt,name=Vsize,json=vsize,proto3" json:"Vsize,omitempty"`
	Weight        int32   `protobuf:"varint,6,opt,name=Weight,json=weight,proto3" json:"Weight,omitempty"`
	Version       uint32  `protobuf:"varint,7,opt,name=Version,json=version,proto3" json:"Version,omitempty"`
	LockTime      uint32  `protobuf:"varint,8,opt,name=LockTime,json=lockTime,proto3" json:"LockTime,omitempty"`
	Vin           []*Vin  `protobuf:"bytes,9,rep,name=Vin,json=vin,proto3" json:"Vin,omitempty"`
	Vout          []*Vout `protobuf:"bytes,10,rep,name=Vout,json=vout,proto3" json:"Vout,omitempty"`
	BlockHash     string  `protobuf:"bytes,11,opt,name=BlockHash,json=blockHash,proto3" json:"BlockHash,omitempty"`
	Confirmations uint64  `protobuf:"varint,12,opt,name=Confirmations,json=confirmations,proto3" json:"Confirmations,omitempty"`
	Time          int64   `protobuf:"varint,13,opt,name=Time,json=time,proto3" json:"Time,omitempty"`
	BlockTime     int64   `protobuf:"varint,14,opt,name=BlockTime,json=block_time,proto3" json:"BlockTime,omitempty"`
}

func (x *GetRawTransactionOutput) Reset() {
	*x = GetRawTransactionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_bitcoin_transaction_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRawTransactionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRawTransactionOutput) ProtoMessage() {}

func (x *GetRawTransactionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_bitcoin_transaction_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRawTransactionOutput.ProtoReflect.Descriptor instead.
func (*GetRawTransactionOutput) Descriptor() ([]byte, []int) {
	return file_pkg_proto_bitcoin_transaction_message_proto_rawDescGZIP(), []int{1}
}

func (x *GetRawTransactionOutput) GetHex() string {
	if x != nil {
		return x.Hex
	}
	return ""
}

func (x *GetRawTransactionOutput) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *GetRawTransactionOutput) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *GetRawTransactionOutput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetRawTransactionOutput) GetVsize() int32 {
	if x != nil {
		return x.Vsize
	}
	return 0
}

func (x *GetRawTransactionOutput) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *GetRawTransactionOutput) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *GetRawTransactionOutput) GetLockTime() uint32 {
	if x != nil {
		return x.LockTime
	}
	return 0
}

func (x *GetRawTransactionOutput) GetVin() []*Vin {
	if x != nil {
		return x.Vin
	}
	return nil
}

func (x *GetRawTransactionOutput) GetVout() []*Vout {
	if x != nil {
		return x.Vout
	}
	return nil
}

func (x *GetRawTransactionOutput) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *GetRawTransactionOutput) GetConfirmations() uint64 {
	if x != nil {
		return x.Confirmations
	}
	return 0
}

func (x *GetRawTransactionOutput) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *GetRawTransactionOutput) GetBlockTime() int64 {
	if x != nil {
		return x.BlockTime
	}
	return 0
}

type Vout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value               float64  `protobuf:"fixed64,1,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
	N                   uint32   `protobuf:"varint,2,opt,name=N,json=n,proto3" json:"N,omitempty"`
	ScriptPubKeyAms     string   `protobuf:"bytes,3,opt,name=ScriptPubKeyAms,json=script_pub_key_ams,proto3" json:"ScriptPubKeyAms,omitempty"`
	ScriptPubKeyHex     string   `protobuf:"bytes,4,opt,name=ScriptPubKeyHex,json=script_pub_key_hex,proto3" json:"ScriptPubKeyHex,omitempty"`
	ScriptPubKeyReqSigs int32    `protobuf:"varint,5,opt,name=ScriptPubKeyReqSigs,json=script_pub_key_req_sigs,proto3" json:"ScriptPubKeyReqSigs,omitempty"`
	ScriptPubKeyType    string   `protobuf:"bytes,6,opt,name=ScriptPubKeyType,json=script_pub_key_type,proto3" json:"ScriptPubKeyType,omitempty"`
	Address             []string `protobuf:"bytes,7,rep,name=Address,json=address,proto3" json:"Address,omitempty"`
}

func (x *Vout) Reset() {
	*x = Vout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_bitcoin_transaction_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vout) ProtoMessage() {}

func (x *Vout) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_bitcoin_transaction_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vout.ProtoReflect.Descriptor instead.
func (*Vout) Descriptor() ([]byte, []int) {
	return file_pkg_proto_bitcoin_transaction_message_proto_rawDescGZIP(), []int{2}
}

func (x *Vout) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Vout) GetN() uint32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *Vout) GetScriptPubKeyAms() string {
	if x != nil {
		return x.ScriptPubKeyAms
	}
	return ""
}

func (x *Vout) GetScriptPubKeyHex() string {
	if x != nil {
		return x.ScriptPubKeyHex
	}
	return ""
}

func (x *Vout) GetScriptPubKeyReqSigs() int32 {
	if x != nil {
		return x.ScriptPubKeyReqSigs
	}
	return 0
}

func (x *Vout) GetScriptPubKeyType() string {
	if x != nil {
		return x.ScriptPubKeyType
	}
	return ""
}

func (x *Vout) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

type Vin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coinbase     string   `protobuf:"bytes,1,opt,name=Coinbase,json=coinbase,proto3" json:"Coinbase,omitempty"`
	TxId         string   `protobuf:"bytes,2,opt,name=TxId,json=tx_id,proto3" json:"TxId,omitempty"`
	Vout         uint32   `protobuf:"varint,3,opt,name=Vout,json=vout,proto3" json:"Vout,omitempty"`
	ScriptSigAms string   `protobuf:"bytes,4,opt,name=ScriptSigAms,json=script_sig_ams,proto3" json:"ScriptSigAms,omitempty"`
	ScriptSigHex string   `protobuf:"bytes,5,opt,name=ScriptSigHex,json=script_sig_hex,proto3" json:"ScriptSigHex,omitempty"`
	Sequence     uint32   `protobuf:"varint,6,opt,name=Sequence,json=sequence,proto3" json:"Sequence,omitempty"`
	Witness      []string `protobuf:"bytes,7,rep,name=Witness,json=witness,proto3" json:"Witness,omitempty"`
}

func (x *Vin) Reset() {
	*x = Vin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_bitcoin_transaction_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vin) ProtoMessage() {}

func (x *Vin) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_bitcoin_transaction_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vin.ProtoReflect.Descriptor instead.
func (*Vin) Descriptor() ([]byte, []int) {
	return file_pkg_proto_bitcoin_transaction_message_proto_rawDescGZIP(), []int{3}
}

func (x *Vin) GetCoinbase() string {
	if x != nil {
		return x.Coinbase
	}
	return ""
}

func (x *Vin) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *Vin) GetVout() uint32 {
	if x != nil {
		return x.Vout
	}
	return 0
}

func (x *Vin) GetScriptSigAms() string {
	if x != nil {
		return x.ScriptSigAms
	}
	return ""
}

func (x *Vin) GetScriptSigHex() string {
	if x != nil {
		return x.ScriptSigHex
	}
	return ""
}

func (x *Vin) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Vin) GetWitness() []string {
	if x != nil {
		return x.Witness
	}
	return nil
}

var File_pkg_proto_bitcoin_transaction_message_proto protoreflect.FileDescriptor

var file_pkg_proto_bitcoin_transaction_message_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x74, 0x63,
	0x6f, 0x69, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62,
	0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x22, 0x2d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x13, 0x0a, 0x04, 0x54, 0x78, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x78, 0x5f, 0x69, 0x64, 0x22, 0x86, 0x03, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x48, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x68, 0x65, 0x78, 0x12, 0x13, 0x0a, 0x04, 0x54, 0x78, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x56, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x6b,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x03, 0x56, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x56, 0x69, 0x6e, 0x52,
	0x03, 0x76, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x04, 0x56, 0x6f, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x56, 0x6f, 0x75,
	0x74, 0x52, 0x04, 0x76, 0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x83,
	0x02, 0x0a, 0x04, 0x56, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0c, 0x0a,
	0x01, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x6e, 0x12, 0x2b, 0x0a, 0x0f, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x41, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x70, 0x75, 0x62,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x0f, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x48, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x68, 0x65, 0x78, 0x12, 0x34, 0x0a, 0x13, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x53, 0x69, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x17, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x73, 0x69, 0x67, 0x73, 0x12, 0x2d, 0x0a, 0x10, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x70, 0x75,
	0x62, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x03, 0x56, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x13, 0x0a, 0x04, 0x54, 0x78, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x56, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x76, 0x6f, 0x75,
	0x74, 0x12, 0x24, 0x0a, 0x0c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x67, 0x41, 0x6d,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f,
	0x73, 0x69, 0x67, 0x5f, 0x61, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0c, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x53, 0x69, 0x67, 0x48, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x5f, 0x68, 0x65, 0x78, 0x12, 0x1a, 0x0a,
	0x08, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x77, 0x69, 0x74, 0x6e,
	0x65, 0x73, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_bitcoin_transaction_message_proto_rawDescOnce sync.Once
	file_pkg_proto_bitcoin_transaction_message_proto_rawDescData = file_pkg_proto_bitcoin_transaction_message_proto_rawDesc
)

func file_pkg_proto_bitcoin_transaction_message_proto_rawDescGZIP() []byte {
	file_pkg_proto_bitcoin_transaction_message_proto_rawDescOnce.Do(func() {
		file_pkg_proto_bitcoin_transaction_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_bitcoin_transaction_message_proto_rawDescData)
	})
	return file_pkg_proto_bitcoin_transaction_message_proto_rawDescData
}

var file_pkg_proto_bitcoin_transaction_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_proto_bitcoin_transaction_message_proto_goTypes = []interface{}{
	(*GetRawTransactionInput)(nil),  // 0: bitcoin.GetRawTransactionInput
	(*GetRawTransactionOutput)(nil), // 1: bitcoin.GetRawTransactionOutput
	(*Vout)(nil),                    // 2: bitcoin.Vout
	(*Vin)(nil),                     // 3: bitcoin.Vin
}
var file_pkg_proto_bitcoin_transaction_message_proto_depIdxs = []int32{
	3, // 0: bitcoin.GetRawTransactionOutput.Vin:type_name -> bitcoin.Vin
	2, // 1: bitcoin.GetRawTransactionOutput.Vout:type_name -> bitcoin.Vout
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_proto_bitcoin_transaction_message_proto_init() }
func file_pkg_proto_bitcoin_transaction_message_proto_init() {
	if File_pkg_proto_bitcoin_transaction_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_bitcoin_transaction_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRawTransactionInput); i {
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
		file_pkg_proto_bitcoin_transaction_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRawTransactionOutput); i {
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
		file_pkg_proto_bitcoin_transaction_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vout); i {
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
		file_pkg_proto_bitcoin_transaction_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vin); i {
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
			RawDescriptor: file_pkg_proto_bitcoin_transaction_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_bitcoin_transaction_message_proto_goTypes,
		DependencyIndexes: file_pkg_proto_bitcoin_transaction_message_proto_depIdxs,
		MessageInfos:      file_pkg_proto_bitcoin_transaction_message_proto_msgTypes,
	}.Build()
	File_pkg_proto_bitcoin_transaction_message_proto = out.File
	file_pkg_proto_bitcoin_transaction_message_proto_rawDesc = nil
	file_pkg_proto_bitcoin_transaction_message_proto_goTypes = nil
	file_pkg_proto_bitcoin_transaction_message_proto_depIdxs = nil
}
