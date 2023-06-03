// Protobuf Version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: modules/products/productsProto/products.proto

package neversuitup_e_commerce_test

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

type FindOneProdcutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOneProdcutReq) Reset() {
	*x = FindOneProdcutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_products_productsProto_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneProdcutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneProdcutReq) ProtoMessage() {}

func (x *FindOneProdcutReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_products_productsProto_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneProdcutReq.ProtoReflect.Descriptor instead.
func (*FindOneProdcutReq) Descriptor() ([]byte, []int) {
	return file_modules_products_productsProto_products_proto_rawDescGZIP(), []int{0}
}

func (x *FindOneProdcutReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_products_productsProto_products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_modules_products_productsProto_products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_modules_products_productsProto_products_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_modules_products_productsProto_products_proto protoreflect.FileDescriptor

var file_modules_products_productsProto_products_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x23, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x63, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0x46, 0x0a,
	0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x32, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x63, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x22, 0x00, 0x28, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x79, 0x61, 0x74, 0x6f, 0x31, 0x35, 0x39, 0x2f, 0x6e, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x75, 0x69, 0x74, 0x75, 0x70, 0x2d, 0x65, 0x2d, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_modules_products_productsProto_products_proto_rawDescOnce sync.Once
	file_modules_products_productsProto_products_proto_rawDescData = file_modules_products_productsProto_products_proto_rawDesc
)

func file_modules_products_productsProto_products_proto_rawDescGZIP() []byte {
	file_modules_products_productsProto_products_proto_rawDescOnce.Do(func() {
		file_modules_products_productsProto_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_products_productsProto_products_proto_rawDescData)
	})
	return file_modules_products_productsProto_products_proto_rawDescData
}

var file_modules_products_productsProto_products_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_modules_products_productsProto_products_proto_goTypes = []interface{}{
	(*FindOneProdcutReq)(nil), // 0: FindOneProdcutReq
	(*Product)(nil),           // 1: Product
}
var file_modules_products_productsProto_products_proto_depIdxs = []int32{
	0, // 0: ProductsServices.FindOneProduct:input_type -> FindOneProdcutReq
	1, // 1: ProductsServices.FindOneProduct:output_type -> Product
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modules_products_productsProto_products_proto_init() }
func file_modules_products_productsProto_products_proto_init() {
	if File_modules_products_productsProto_products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_products_productsProto_products_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneProdcutReq); i {
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
		file_modules_products_productsProto_products_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_modules_products_productsProto_products_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_products_productsProto_products_proto_goTypes,
		DependencyIndexes: file_modules_products_productsProto_products_proto_depIdxs,
		MessageInfos:      file_modules_products_productsProto_products_proto_msgTypes,
	}.Build()
	File_modules_products_productsProto_products_proto = out.File
	file_modules_products_productsProto_products_proto_rawDesc = nil
	file_modules_products_productsProto_products_proto_goTypes = nil
	file_modules_products_productsProto_products_proto_depIdxs = nil
}
