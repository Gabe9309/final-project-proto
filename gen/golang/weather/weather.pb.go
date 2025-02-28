// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: weather/weather.proto

package weather_service

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

type WeatherData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location    string  `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Temperature float64 `protobuf:"fixed64,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *WeatherData) Reset() {
	*x = WeatherData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherData) ProtoMessage() {}

func (x *WeatherData) ProtoReflect() protoreflect.Message {
	mi := &file_weather_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherData.ProtoReflect.Descriptor instead.
func (*WeatherData) Descriptor() ([]byte, []int) {
	return file_weather_weather_proto_rawDescGZIP(), []int{0}
}

func (x *WeatherData) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *WeatherData) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *WeatherData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_weather_weather_proto protoreflect.FileDescriptor

var file_weather_weather_proto_rawDesc = []byte{
	0x0a, 0x15, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x6d, 0x0a, 0x0b, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x5c, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1c, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x00, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x62, 0x65, 0x39, 0x33, 0x30, 0x39, 0x2f, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3b, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weather_weather_proto_rawDescOnce sync.Once
	file_weather_weather_proto_rawDescData = file_weather_weather_proto_rawDesc
)

func file_weather_weather_proto_rawDescGZIP() []byte {
	file_weather_weather_proto_rawDescOnce.Do(func() {
		file_weather_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_weather_proto_rawDescData)
	})
	return file_weather_weather_proto_rawDescData
}

var file_weather_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_weather_weather_proto_goTypes = []interface{}{
	(*WeatherData)(nil), // 0: weather_service.WeatherData
}
var file_weather_weather_proto_depIdxs = []int32{
	0, // 0: weather_service.WeatherService.GetWeather:input_type -> weather_service.WeatherData
	0, // 1: weather_service.WeatherService.GetWeather:output_type -> weather_service.WeatherData
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_weather_weather_proto_init() }
func file_weather_weather_proto_init() {
	if File_weather_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherData); i {
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
			RawDescriptor: file_weather_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_weather_proto_goTypes,
		DependencyIndexes: file_weather_weather_proto_depIdxs,
		MessageInfos:      file_weather_weather_proto_msgTypes,
	}.Build()
	File_weather_weather_proto = out.File
	file_weather_weather_proto_rawDesc = nil
	file_weather_weather_proto_goTypes = nil
	file_weather_weather_proto_depIdxs = nil
}
