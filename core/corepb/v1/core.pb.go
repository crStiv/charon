// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: core/corepb/v1/core.proto

package v1

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

type Duty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot uint64 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"` // uint64
	Type int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"` // core.DutyType
}

func (x *Duty) Reset() {
	*x = Duty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_corepb_v1_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Duty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Duty) ProtoMessage() {}

func (x *Duty) ProtoReflect() protoreflect.Message {
	mi := &file_core_corepb_v1_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Duty.ProtoReflect.Descriptor instead.
func (*Duty) Descriptor() ([]byte, []int) {
	return file_core_corepb_v1_core_proto_rawDescGZIP(), []int{0}
}

func (x *Duty) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *Duty) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type UnsignedDataSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set map[string][]byte `protobuf:"bytes,1,rep,name=set,proto3" json:"set,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[core.PubKey]core.UnsignedData
}

func (x *UnsignedDataSet) Reset() {
	*x = UnsignedDataSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_corepb_v1_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsignedDataSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsignedDataSet) ProtoMessage() {}

func (x *UnsignedDataSet) ProtoReflect() protoreflect.Message {
	mi := &file_core_corepb_v1_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsignedDataSet.ProtoReflect.Descriptor instead.
func (*UnsignedDataSet) Descriptor() ([]byte, []int) {
	return file_core_corepb_v1_core_proto_rawDescGZIP(), []int{1}
}

func (x *UnsignedDataSet) GetSet() map[string][]byte {
	if x != nil {
		return x.Set
	}
	return nil
}

type ParSignedDataSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set map[string]*ParSignedData `protobuf:"bytes,1,rep,name=set,proto3" json:"set,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[core.PubKey]core.ParSignedData
}

func (x *ParSignedDataSet) Reset() {
	*x = ParSignedDataSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_corepb_v1_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParSignedDataSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParSignedDataSet) ProtoMessage() {}

func (x *ParSignedDataSet) ProtoReflect() protoreflect.Message {
	mi := &file_core_corepb_v1_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParSignedDataSet.ProtoReflect.Descriptor instead.
func (*ParSignedDataSet) Descriptor() ([]byte, []int) {
	return file_core_corepb_v1_core_proto_rawDescGZIP(), []int{2}
}

func (x *ParSignedDataSet) GetSet() map[string]*ParSignedData {
	if x != nil {
		return x.Set
	}
	return nil
}

type ParSignedData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`                          // []byte
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`                // core.Signature
	ShareIdx  int32  `protobuf:"varint,3,opt,name=share_idx,json=shareIdx,proto3" json:"share_idx,omitempty"` // int
}

func (x *ParSignedData) Reset() {
	*x = ParSignedData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_corepb_v1_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParSignedData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParSignedData) ProtoMessage() {}

func (x *ParSignedData) ProtoReflect() protoreflect.Message {
	mi := &file_core_corepb_v1_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParSignedData.ProtoReflect.Descriptor instead.
func (*ParSignedData) Descriptor() ([]byte, []int) {
	return file_core_corepb_v1_core_proto_rawDescGZIP(), []int{3}
}

func (x *ParSignedData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ParSignedData) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *ParSignedData) GetShareIdx() int32 {
	if x != nil {
		return x.ShareIdx
	}
	return 0
}

var File_core_corepb_v1_core_proto protoreflect.FileDescriptor

var file_core_corepb_v1_core_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x22, 0x2e, 0x0a, 0x04, 0x44,
	0x75, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x0f,
	0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x12,
	0x3a, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x2e, 0x53, 0x65,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x73, 0x65, 0x74, 0x1a, 0x36, 0x0a, 0x08, 0x53,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xa6, 0x01, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x12, 0x3b, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x73, 0x65, 0x74, 0x1a, 0x55, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5e, 0x0a, 0x0d,
	0x50, 0x61, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x73, 0x68, 0x61, 0x72, 0x65, 0x49, 0x64, 0x78, 0x42, 0x2e, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x62, 0x6f, 0x6c, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_corepb_v1_core_proto_rawDescOnce sync.Once
	file_core_corepb_v1_core_proto_rawDescData = file_core_corepb_v1_core_proto_rawDesc
)

func file_core_corepb_v1_core_proto_rawDescGZIP() []byte {
	file_core_corepb_v1_core_proto_rawDescOnce.Do(func() {
		file_core_corepb_v1_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_corepb_v1_core_proto_rawDescData)
	})
	return file_core_corepb_v1_core_proto_rawDescData
}

var file_core_corepb_v1_core_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_core_corepb_v1_core_proto_goTypes = []interface{}{
	(*Duty)(nil),             // 0: core.corepb.v1.Duty
	(*UnsignedDataSet)(nil),  // 1: core.corepb.v1.UnsignedDataSet
	(*ParSignedDataSet)(nil), // 2: core.corepb.v1.ParSignedDataSet
	(*ParSignedData)(nil),    // 3: core.corepb.v1.ParSignedData
	nil,                      // 4: core.corepb.v1.UnsignedDataSet.SetEntry
	nil,                      // 5: core.corepb.v1.ParSignedDataSet.SetEntry
}
var file_core_corepb_v1_core_proto_depIdxs = []int32{
	4, // 0: core.corepb.v1.UnsignedDataSet.set:type_name -> core.corepb.v1.UnsignedDataSet.SetEntry
	5, // 1: core.corepb.v1.ParSignedDataSet.set:type_name -> core.corepb.v1.ParSignedDataSet.SetEntry
	3, // 2: core.corepb.v1.ParSignedDataSet.SetEntry.value:type_name -> core.corepb.v1.ParSignedData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_core_corepb_v1_core_proto_init() }
func file_core_corepb_v1_core_proto_init() {
	if File_core_corepb_v1_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_corepb_v1_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Duty); i {
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
		file_core_corepb_v1_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsignedDataSet); i {
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
		file_core_corepb_v1_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParSignedDataSet); i {
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
		file_core_corepb_v1_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParSignedData); i {
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
			RawDescriptor: file_core_corepb_v1_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_corepb_v1_core_proto_goTypes,
		DependencyIndexes: file_core_corepb_v1_core_proto_depIdxs,
		MessageInfos:      file_core_corepb_v1_core_proto_msgTypes,
	}.Build()
	File_core_corepb_v1_core_proto = out.File
	file_core_corepb_v1_core_proto_rawDesc = nil
	file_core_corepb_v1_core_proto_goTypes = nil
	file_core_corepb_v1_core_proto_depIdxs = nil
}
