// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: einride/bigquery/public/v1/san_fransisco_transit_stop_time.proto

package publicv1

import (
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
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

// Protobuf schema for the BigQuery public table:
//
//  bigquery-public-data.san_francisco_transit_muni.stop_times
type SanFransiscoTransitStopTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StopId         int64                `protobuf:"varint,1,opt,name=stop_id,json=stopId,proto3" json:"stop_id,omitempty"`                           // INTEGER NULLABLE
	TripId         int64                `protobuf:"varint,2,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`                           // INTEGER NULLABLE
	StopSequence   int64                `protobuf:"varint,3,opt,name=stop_sequence,json=stopSequence,proto3" json:"stop_sequence,omitempty"`         // INTEGER NULLABLE
	ArrivalTime    *timeofday.TimeOfDay `protobuf:"bytes,4,opt,name=arrival_time,json=arrivalTime,proto3" json:"arrival_time,omitempty"`             // TIME NULLABLE
	ArrivesNextDay bool                 `protobuf:"varint,5,opt,name=arrives_next_day,json=arrivesNextDay,proto3" json:"arrives_next_day,omitempty"` // BOOLEAN NULLABLE
	DepartureTime  *timeofday.TimeOfDay `protobuf:"bytes,6,opt,name=departure_time,json=departureTime,proto3" json:"departure_time,omitempty"`       // TIME NULLABLE
	DepartsNextDay bool                 `protobuf:"varint,7,opt,name=departs_next_day,json=departsNextDay,proto3" json:"departs_next_day,omitempty"` // BOOLEAN NULLABLE
	DropoffType    string               `protobuf:"bytes,8,opt,name=dropoff_type,json=dropoffType,proto3" json:"dropoff_type,omitempty"`             // STRING NULLABLE
	ExactTimepoint bool                 `protobuf:"varint,9,opt,name=exact_timepoint,json=exactTimepoint,proto3" json:"exact_timepoint,omitempty"`   // BOOLEAN NULLABLE
}

func (x *SanFransiscoTransitStopTime) Reset() {
	*x = SanFransiscoTransitStopTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SanFransiscoTransitStopTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SanFransiscoTransitStopTime) ProtoMessage() {}

func (x *SanFransiscoTransitStopTime) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SanFransiscoTransitStopTime.ProtoReflect.Descriptor instead.
func (*SanFransiscoTransitStopTime) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescGZIP(), []int{0}
}

func (x *SanFransiscoTransitStopTime) GetStopId() int64 {
	if x != nil {
		return x.StopId
	}
	return 0
}

func (x *SanFransiscoTransitStopTime) GetTripId() int64 {
	if x != nil {
		return x.TripId
	}
	return 0
}

func (x *SanFransiscoTransitStopTime) GetStopSequence() int64 {
	if x != nil {
		return x.StopSequence
	}
	return 0
}

func (x *SanFransiscoTransitStopTime) GetArrivalTime() *timeofday.TimeOfDay {
	if x != nil {
		return x.ArrivalTime
	}
	return nil
}

func (x *SanFransiscoTransitStopTime) GetArrivesNextDay() bool {
	if x != nil {
		return x.ArrivesNextDay
	}
	return false
}

func (x *SanFransiscoTransitStopTime) GetDepartureTime() *timeofday.TimeOfDay {
	if x != nil {
		return x.DepartureTime
	}
	return nil
}

func (x *SanFransiscoTransitStopTime) GetDepartsNextDay() bool {
	if x != nil {
		return x.DepartsNextDay
	}
	return false
}

func (x *SanFransiscoTransitStopTime) GetDropoffType() string {
	if x != nil {
		return x.DropoffType
	}
	return ""
}

func (x *SanFransiscoTransitStopTime) GetExactTimepoint() bool {
	if x != nil {
		return x.ExactTimepoint
	}
	return false
}

var File_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto protoreflect.FileDescriptor

var file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6e,
	0x5f, 0x66, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x66, 0x64, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x03, 0x0a, 0x1b,
	0x53, 0x61, 0x6e, 0x46, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x73, 0x63, 0x6f, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x69, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x74, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x6f, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79,
	0x52, 0x0b, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65, 0x73, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65, 0x73,
	0x4e, 0x65, 0x78, 0x74, 0x44, 0x61, 0x79, 0x12, 0x3d, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x73, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x44, 0x61, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x66, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x66, 0x66, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x78,
	0x61, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0xb2, 0x02, 0x0a,
	0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x42,
	0x20, 0x53, 0x61, 0x6e, 0x46, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x73, 0x63, 0x6f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x61, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x74, 0x65, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x42, 0x50, 0xaa, 0x02, 0x1a, 0x45,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1b, 0x45, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x27, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x5f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1d, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x42, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x3a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescOnce sync.Once
	file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescData = file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDesc
)

func file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescGZIP() []byte {
	file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescData)
	})
	return file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescData
}

var file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_goTypes = []interface{}{
	(*SanFransiscoTransitStopTime)(nil), // 0: einride.bigquery.public.v1.SanFransiscoTransitStopTime
	(*timeofday.TimeOfDay)(nil),         // 1: google.type.TimeOfDay
}
var file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_depIdxs = []int32{
	1, // 0: einride.bigquery.public.v1.SanFransiscoTransitStopTime.arrival_time:type_name -> google.type.TimeOfDay
	1, // 1: einride.bigquery.public.v1.SanFransiscoTransitStopTime.departure_time:type_name -> google.type.TimeOfDay
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_init() }
func file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_init() {
	if File_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SanFransiscoTransitStopTime); i {
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
			RawDescriptor: file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_msgTypes,
	}.Build()
	File_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto = out.File
	file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDesc = nil
	file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_goTypes = nil
	file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_depIdxs = nil
}
