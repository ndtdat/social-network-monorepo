// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: purchase/model/transaction_type.proto

package model

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

type TransactionType int32

const (
	TransactionType_TT_NONE     TransactionType = 0
	TransactionType_TT_BUY      TransactionType = 1
	TransactionType_VS_DISCOUNT TransactionType = 2
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "TT_NONE",
		1: "TT_BUY",
		2: "VS_DISCOUNT",
	}
	TransactionType_value = map[string]int32{
		"TT_NONE":     0,
		"TT_BUY":      1,
		"VS_DISCOUNT": 2,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_purchase_model_transaction_type_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_purchase_model_transaction_type_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_purchase_model_transaction_type_proto_rawDescGZIP(), []int{0}
}

var File_purchase_model_transaction_type_proto protoreflect.FileDescriptor

var file_purchase_model_transaction_type_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2a, 0x3b,
	0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x54, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x54, 0x54, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x53,
	0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x02, 0x42, 0x56, 0x48, 0x01, 0x5a,
	0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x64, 0x74, 0x64,
	0x61, 0x74, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3b, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_purchase_model_transaction_type_proto_rawDescOnce sync.Once
	file_purchase_model_transaction_type_proto_rawDescData = file_purchase_model_transaction_type_proto_rawDesc
)

func file_purchase_model_transaction_type_proto_rawDescGZIP() []byte {
	file_purchase_model_transaction_type_proto_rawDescOnce.Do(func() {
		file_purchase_model_transaction_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_purchase_model_transaction_type_proto_rawDescData)
	})
	return file_purchase_model_transaction_type_proto_rawDescData
}

var file_purchase_model_transaction_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_purchase_model_transaction_type_proto_goTypes = []interface{}{
	(TransactionType)(0), // 0: model.TransactionType
}
var file_purchase_model_transaction_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_purchase_model_transaction_type_proto_init() }
func file_purchase_model_transaction_type_proto_init() {
	if File_purchase_model_transaction_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_purchase_model_transaction_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_purchase_model_transaction_type_proto_goTypes,
		DependencyIndexes: file_purchase_model_transaction_type_proto_depIdxs,
		EnumInfos:         file_purchase_model_transaction_type_proto_enumTypes,
	}.Build()
	File_purchase_model_transaction_type_proto = out.File
	file_purchase_model_transaction_type_proto_rawDesc = nil
	file_purchase_model_transaction_type_proto_goTypes = nil
	file_purchase_model_transaction_type_proto_depIdxs = nil
}
