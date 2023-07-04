// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: proto/ethereum/hdwallet.proto

package ethereum

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_ethereum_hdwallet_proto protoreflect.FileDescriptor

var file_proto_ethereum_hdwallet_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2f, 0x68, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x1a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x68, 0x64, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x9a, 0x01, 0x0a, 0x08, 0x48, 0x44, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x49, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x1b, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_proto_ethereum_hdwallet_proto_goTypes = []interface{}{
	(*CreateWalletInput)(nil),  // 0: ethereum.CreateWalletInput
	(*GetBalanceInput)(nil),    // 1: ethereum.GetBalanceInput
	(*CreateWalletOutput)(nil), // 2: ethereum.CreateWalletOutput
	(*GetBalanceOutput)(nil),   // 3: ethereum.GetBalanceOutput
}
var file_proto_ethereum_hdwallet_proto_depIdxs = []int32{
	0, // 0: ethereum.HDWallet.CreateWallet:input_type -> ethereum.CreateWalletInput
	1, // 1: ethereum.HDWallet.GetBalance:input_type -> ethereum.GetBalanceInput
	2, // 2: ethereum.HDWallet.CreateWallet:output_type -> ethereum.CreateWalletOutput
	3, // 3: ethereum.HDWallet.GetBalance:output_type -> ethereum.GetBalanceOutput
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ethereum_hdwallet_proto_init() }
func file_proto_ethereum_hdwallet_proto_init() {
	if File_proto_ethereum_hdwallet_proto != nil {
		return
	}
	file_proto_ethereum_hdwallet_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ethereum_hdwallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ethereum_hdwallet_proto_goTypes,
		DependencyIndexes: file_proto_ethereum_hdwallet_proto_depIdxs,
	}.Build()
	File_proto_ethereum_hdwallet_proto = out.File
	file_proto_ethereum_hdwallet_proto_rawDesc = nil
	file_proto_ethereum_hdwallet_proto_goTypes = nil
	file_proto_ethereum_hdwallet_proto_depIdxs = nil
}