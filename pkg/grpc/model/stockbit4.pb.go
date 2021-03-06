// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: stockbit4.proto

package __

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

type GetMovieDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=Keyword,proto3" json:"Keyword,omitempty"`
	Page    int32  `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
}

func (x *GetMovieDataRequest) Reset() {
	*x = GetMovieDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stockbit4_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieDataRequest) ProtoMessage() {}

func (x *GetMovieDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stockbit4_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieDataRequest.ProtoReflect.Descriptor instead.
func (*GetMovieDataRequest) Descriptor() ([]byte, []int) {
	return file_stockbit4_proto_rawDescGZIP(), []int{0}
}

func (x *GetMovieDataRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *GetMovieDataRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetGmStatShopTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopID uint64 `protobuf:"varint,1,opt,name=ShopID,proto3" json:"ShopID,omitempty"`
}

func (x *GetGmStatShopTransactionRequest) Reset() {
	*x = GetGmStatShopTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stockbit4_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGmStatShopTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGmStatShopTransactionRequest) ProtoMessage() {}

func (x *GetGmStatShopTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stockbit4_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGmStatShopTransactionRequest.ProtoReflect.Descriptor instead.
func (*GetGmStatShopTransactionRequest) Descriptor() ([]byte, []int) {
	return file_stockbit4_proto_rawDescGZIP(), []int{1}
}

func (x *GetGmStatShopTransactionRequest) GetShopID() uint64 {
	if x != nil {
		return x.ShopID
	}
	return 0
}

type MovieData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=Poster,proto3" json:"Poster,omitempty"`
}

func (x *MovieData) Reset() {
	*x = MovieData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stockbit4_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieData) ProtoMessage() {}

func (x *MovieData) ProtoReflect() protoreflect.Message {
	mi := &file_stockbit4_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieData.ProtoReflect.Descriptor instead.
func (*MovieData) Descriptor() ([]byte, []int) {
	return file_stockbit4_proto_rawDescGZIP(), []int{2}
}

func (x *MovieData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieData) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieData) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *MovieData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieData) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type GetMovieDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MoviesData []*MovieData `protobuf:"bytes,1,rep,name=MoviesData,proto3" json:"MoviesData,omitempty"`
}

func (x *GetMovieDataResponse) Reset() {
	*x = GetMovieDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stockbit4_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieDataResponse) ProtoMessage() {}

func (x *GetMovieDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stockbit4_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieDataResponse.ProtoReflect.Descriptor instead.
func (*GetMovieDataResponse) Descriptor() ([]byte, []int) {
	return file_stockbit4_proto_rawDescGZIP(), []int{3}
}

func (x *GetMovieDataResponse) GetMoviesData() []*MovieData {
	if x != nil {
		return x.MoviesData
	}
	return nil
}

var File_stockbit4_proto protoreflect.FileDescriptor

var file_stockbit4_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74, 0x34, 0x22, 0x43, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x61, 0x67,
	0x65, 0x22, 0x39, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x47, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x53, 0x68,
	0x6f, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x44, 0x22, 0x79, 0x0a, 0x09,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59,
	0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0x4c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x0a, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74, 0x34, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x32, 0x5e, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69,
	0x74, 0x34, 0x12, 0x51, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74, 0x34, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74, 0x34, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_stockbit4_proto_rawDescOnce sync.Once
	file_stockbit4_proto_rawDescData = file_stockbit4_proto_rawDesc
)

func file_stockbit4_proto_rawDescGZIP() []byte {
	file_stockbit4_proto_rawDescOnce.Do(func() {
		file_stockbit4_proto_rawDescData = protoimpl.X.CompressGZIP(file_stockbit4_proto_rawDescData)
	})
	return file_stockbit4_proto_rawDescData
}

var file_stockbit4_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stockbit4_proto_goTypes = []interface{}{
	(*GetMovieDataRequest)(nil),             // 0: stockbit4.GetMovieDataRequest
	(*GetGmStatShopTransactionRequest)(nil), // 1: stockbit4.GetGmStatShopTransactionRequest
	(*MovieData)(nil),                       // 2: stockbit4.MovieData
	(*GetMovieDataResponse)(nil),            // 3: stockbit4.GetMovieDataResponse
}
var file_stockbit4_proto_depIdxs = []int32{
	2, // 0: stockbit4.GetMovieDataResponse.MoviesData:type_name -> stockbit4.MovieData
	0, // 1: stockbit4.Stockbit4.GetMovieData:input_type -> stockbit4.GetMovieDataRequest
	3, // 2: stockbit4.Stockbit4.GetMovieData:output_type -> stockbit4.GetMovieDataResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stockbit4_proto_init() }
func file_stockbit4_proto_init() {
	if File_stockbit4_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stockbit4_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieDataRequest); i {
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
		file_stockbit4_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGmStatShopTransactionRequest); i {
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
		file_stockbit4_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieData); i {
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
		file_stockbit4_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieDataResponse); i {
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
			RawDescriptor: file_stockbit4_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stockbit4_proto_goTypes,
		DependencyIndexes: file_stockbit4_proto_depIdxs,
		MessageInfos:      file_stockbit4_proto_msgTypes,
	}.Build()
	File_stockbit4_proto = out.File
	file_stockbit4_proto_rawDesc = nil
	file_stockbit4_proto_goTypes = nil
	file_stockbit4_proto_depIdxs = nil
}
