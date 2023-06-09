// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: KeychainService.proto

package lightwalletrpc

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

type KeyLocator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// The family of key being identified.
	KeyFamily uint32 `protobuf:"varint,1,opt,name=keyFamily,proto3" json:"keyFamily,omitempty"`
	/// The precise index of the key being identified.
	KeyIndex uint32 `protobuf:"varint,2,opt,name=keyIndex,proto3" json:"keyIndex,omitempty"`
}

func (x *KeyLocator) Reset() {
	*x = KeyLocator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KeychainService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyLocator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyLocator) ProtoMessage() {}

func (x *KeyLocator) ProtoReflect() protoreflect.Message {
	mi := &file_KeychainService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyLocator.ProtoReflect.Descriptor instead.
func (*KeyLocator) Descriptor() ([]byte, []int) {
	return file_KeychainService_proto_rawDescGZIP(), []int{0}
}

func (x *KeyLocator) GetKeyFamily() uint32 {
	if x != nil {
		return x.KeyFamily
	}
	return 0
}

func (x *KeyLocator) GetKeyIndex() uint32 {
	if x != nil {
		return x.KeyIndex
	}
	return 0
}

type KeyDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	//The raw bytes of the key being identified. Encoded as Hex. Either this or the KeyLocator
	//must be specified.
	PubKey string `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	//*
	//The key locator that identifies which key to use for signing. Either this
	//or the raw bytes of the target key must be specified.
	Locator *KeyLocator `protobuf:"bytes,2,opt,name=locator,proto3" json:"locator,omitempty"`
}

func (x *KeyDescriptor) Reset() {
	*x = KeyDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KeychainService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyDescriptor) ProtoMessage() {}

func (x *KeyDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_KeychainService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyDescriptor.ProtoReflect.Descriptor instead.
func (*KeyDescriptor) Descriptor() ([]byte, []int) {
	return file_KeychainService_proto_rawDescGZIP(), []int{1}
}

func (x *KeyDescriptor) GetPubKey() string {
	if x != nil {
		return x.PubKey
	}
	return ""
}

func (x *KeyDescriptor) GetLocator() *KeyLocator {
	if x != nil {
		return x.Locator
	}
	return nil
}

type DeriveNextKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyFamily uint32 `protobuf:"varint,1,opt,name=keyFamily,proto3" json:"keyFamily,omitempty"`
}

func (x *DeriveNextKeyReq) Reset() {
	*x = DeriveNextKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KeychainService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeriveNextKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeriveNextKeyReq) ProtoMessage() {}

func (x *DeriveNextKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_KeychainService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeriveNextKeyReq.ProtoReflect.Descriptor instead.
func (*DeriveNextKeyReq) Descriptor() ([]byte, []int) {
	return file_KeychainService_proto_rawDescGZIP(), []int{2}
}

func (x *DeriveNextKeyReq) GetKeyFamily() uint32 {
	if x != nil {
		return x.KeyFamily
	}
	return 0
}

type IsOurAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOur bool `protobuf:"varint,1,opt,name=isOur,proto3" json:"isOur,omitempty"`
}

func (x *IsOurAddressResponse) Reset() {
	*x = IsOurAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KeychainService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsOurAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsOurAddressResponse) ProtoMessage() {}

func (x *IsOurAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_KeychainService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsOurAddressResponse.ProtoReflect.Descriptor instead.
func (*IsOurAddressResponse) Descriptor() ([]byte, []int) {
	return file_KeychainService_proto_rawDescGZIP(), []int{3}
}

func (x *IsOurAddressResponse) GetIsOur() bool {
	if x != nil {
		return x.IsOur
	}
	return false
}

var File_KeychainService_proto protoreflect.FileDescriptor

var file_KeychainService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x4b, 0x65, 0x79, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x72, 0x70, 0x63, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x5d, 0x0a,
	0x0d, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4b, 0x65, 0x79, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x30, 0x0a, 0x10,
	0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x4e, 0x65, 0x78, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22, 0x2c,
	0x0a, 0x14, 0x49, 0x73, 0x4f, 0x75, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x4f, 0x75, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4f, 0x75, 0x72, 0x32, 0xc9, 0x02, 0x0a,
	0x0f, 0x4b, 0x65, 0x79, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x50, 0x0a, 0x0d, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x4e, 0x65, 0x78, 0x74, 0x4b, 0x65,
	0x79, 0x12, 0x20, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x72,
	0x70, 0x63, 0x2e, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x4e, 0x65, 0x78, 0x74, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x72, 0x70, 0x63, 0x2e, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x12, 0x46, 0x0a, 0x09, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x1a, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x72, 0x70, 0x63,
	0x2e, 0x4b, 0x65, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x1d, 0x2e, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4b, 0x65, 0x79,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x4a, 0x0a, 0x0d, 0x44, 0x65,
	0x72, 0x69, 0x76, 0x65, 0x50, 0x72, 0x69, 0x76, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x2e, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4b, 0x65, 0x79,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x78, 0x45,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x12, 0x50, 0x0a, 0x0c, 0x49, 0x73, 0x4f, 0x75, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x78, 0x45, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x64, 0x1a, 0x24, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x49, 0x73, 0x4f, 0x75, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KeychainService_proto_rawDescOnce sync.Once
	file_KeychainService_proto_rawDescData = file_KeychainService_proto_rawDesc
)

func file_KeychainService_proto_rawDescGZIP() []byte {
	file_KeychainService_proto_rawDescOnce.Do(func() {
		file_KeychainService_proto_rawDescData = protoimpl.X.CompressGZIP(file_KeychainService_proto_rawDescData)
	})
	return file_KeychainService_proto_rawDescData
}

var file_KeychainService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_KeychainService_proto_goTypes = []interface{}{
	(*KeyLocator)(nil),           // 0: lightwalletrpc.KeyLocator
	(*KeyDescriptor)(nil),        // 1: lightwalletrpc.KeyDescriptor
	(*DeriveNextKeyReq)(nil),     // 2: lightwalletrpc.DeriveNextKeyReq
	(*IsOurAddressResponse)(nil), // 3: lightwalletrpc.IsOurAddressResponse
	(*HexEncoded)(nil),           // 4: lightwalletrpc.HexEncoded
}
var file_KeychainService_proto_depIdxs = []int32{
	0, // 0: lightwalletrpc.KeyDescriptor.locator:type_name -> lightwalletrpc.KeyLocator
	2, // 1: lightwalletrpc.KeychainService.DeriveNextKey:input_type -> lightwalletrpc.DeriveNextKeyReq
	0, // 2: lightwalletrpc.KeychainService.DeriveKey:input_type -> lightwalletrpc.KeyLocator
	1, // 3: lightwalletrpc.KeychainService.DerivePrivKey:input_type -> lightwalletrpc.KeyDescriptor
	4, // 4: lightwalletrpc.KeychainService.IsOurAddress:input_type -> lightwalletrpc.HexEncoded
	1, // 5: lightwalletrpc.KeychainService.DeriveNextKey:output_type -> lightwalletrpc.KeyDescriptor
	1, // 6: lightwalletrpc.KeychainService.DeriveKey:output_type -> lightwalletrpc.KeyDescriptor
	4, // 7: lightwalletrpc.KeychainService.DerivePrivKey:output_type -> lightwalletrpc.HexEncoded
	3, // 8: lightwalletrpc.KeychainService.IsOurAddress:output_type -> lightwalletrpc.IsOurAddressResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_KeychainService_proto_init() }
func file_KeychainService_proto_init() {
	if File_KeychainService_proto != nil {
		return
	}
	file_Common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_KeychainService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyLocator); i {
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
		file_KeychainService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyDescriptor); i {
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
		file_KeychainService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeriveNextKeyReq); i {
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
		file_KeychainService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsOurAddressResponse); i {
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
			RawDescriptor: file_KeychainService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_KeychainService_proto_goTypes,
		DependencyIndexes: file_KeychainService_proto_depIdxs,
		MessageInfos:      file_KeychainService_proto_msgTypes,
	}.Build()
	File_KeychainService_proto = out.File
	file_KeychainService_proto_rawDesc = nil
	file_KeychainService_proto_goTypes = nil
	file_KeychainService_proto_depIdxs = nil
}
