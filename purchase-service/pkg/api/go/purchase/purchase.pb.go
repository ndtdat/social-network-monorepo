// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: purchase/purchase.proto

package purchase

import (
	rpc "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/rpc"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_purchase_purchase_proto protoreflect.FileDescriptor

var file_purchase_purchase_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34,
	0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x5f, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x62, 0x79, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x6c, 0x0a, 0x08, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x12, 0x60, 0x0a, 0x1c, 0x49, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x75,
	0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x49, 0x44,
	0x12, 0x28, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65,
	0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x5b, 0x48, 0x01, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x64, 0x74, 0x64, 0x61, 0x74, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70,
	0x6f, 0x2f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x3b, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_purchase_purchase_proto_goTypes = []interface{}{
	(*rpc.IAllocateVoucherByCampaignIDRequest)(nil), // 0: rpc.IAllocateVoucherByCampaignIDRequest
	(*emptypb.Empty)(nil),                           // 1: google.protobuf.Empty
}
var file_purchase_purchase_proto_depIdxs = []int32{
	0, // 0: purchase.Purchase.IAllocateVoucherByCampaignID:input_type -> rpc.IAllocateVoucherByCampaignIDRequest
	1, // 1: purchase.Purchase.IAllocateVoucherByCampaignID:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_purchase_purchase_proto_init() }
func file_purchase_purchase_proto_init() {
	if File_purchase_purchase_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_purchase_purchase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_purchase_purchase_proto_goTypes,
		DependencyIndexes: file_purchase_purchase_proto_depIdxs,
	}.Build()
	File_purchase_purchase_proto = out.File
	file_purchase_purchase_proto_rawDesc = nil
	file_purchase_purchase_proto_goTypes = nil
	file_purchase_purchase_proto_depIdxs = nil
}
