// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: proto/bitcoin/blockchain.proto

package bitcoin

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

var File_proto_bitcoin_blockchain_proto protoreflect.FileDescriptor

var file_proto_bitcoin_blockchain_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xea, 0x01, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1d, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1c, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x62, 0x69, 0x74, 0x63,
	0x6f, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_proto_bitcoin_blockchain_proto_goTypes = []interface{}{
	(*GetBlockCountRequest)(nil),  // 0: bitcoin.GetBlockCountRequest
	(*GetBlockHashRequest)(nil),   // 1: bitcoin.GetBlockHashRequest
	(*GetBlockRequest)(nil),       // 2: bitcoin.GetBlockRequest
	(*GetBlockCountResponse)(nil), // 3: bitcoin.GetBlockCountResponse
	(*GetBlockHashResponse)(nil),  // 4: bitcoin.GetBlockHashResponse
	(*GetBlockResponse)(nil),      // 5: bitcoin.GetBlockResponse
}
var file_proto_bitcoin_blockchain_proto_depIdxs = []int32{
	0, // 0: bitcoin.Blockchain.GetBlockCount:input_type -> bitcoin.GetBlockCountRequest
	1, // 1: bitcoin.Blockchain.GetBlockHash:input_type -> bitcoin.GetBlockHashRequest
	2, // 2: bitcoin.Blockchain.GetBlock:input_type -> bitcoin.GetBlockRequest
	3, // 3: bitcoin.Blockchain.GetBlockCount:output_type -> bitcoin.GetBlockCountResponse
	4, // 4: bitcoin.Blockchain.GetBlockHash:output_type -> bitcoin.GetBlockHashResponse
	5, // 5: bitcoin.Blockchain.GetBlock:output_type -> bitcoin.GetBlockResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_bitcoin_blockchain_proto_init() }
func file_proto_bitcoin_blockchain_proto_init() {
	if File_proto_bitcoin_blockchain_proto != nil {
		return
	}
	file_proto_bitcoin_blockchain_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_bitcoin_blockchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_bitcoin_blockchain_proto_goTypes,
		DependencyIndexes: file_proto_bitcoin_blockchain_proto_depIdxs,
	}.Build()
	File_proto_bitcoin_blockchain_proto = out.File
	file_proto_bitcoin_blockchain_proto_rawDesc = nil
	file_proto_bitcoin_blockchain_proto_goTypes = nil
	file_proto_bitcoin_blockchain_proto_depIdxs = nil
}