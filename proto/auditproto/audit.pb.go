// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: proto/audit.proto

package auditproto

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	BoxId     int32  `protobuf:"varint,2,opt,name=boxId,proto3" json:"boxId,omitempty"`
	ErrorMsg  string `json:"errMsgha" bson:"errMsgho" protobuf:"bytes,12,opt,name=errorMsg,proto3"` // #json: "errMsgha" #bson: "errMsgho"
	ErrorCode int32  `protobuf:"varint,13,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	Timestamp int64  `protobuf:"varint,14,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_audit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_audit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_audit_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Message) GetBoxId() int32 {
	if x != nil {
		return x.BoxId
	}
	return 0
}

func (x *Message) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *Message) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_proto_audit_proto protoreflect.FileDescriptor

var file_proto_audit_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8f, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_audit_proto_rawDescOnce sync.Once
	file_proto_audit_proto_rawDescData = file_proto_audit_proto_rawDesc
)

func file_proto_audit_proto_rawDescGZIP() []byte {
	file_proto_audit_proto_rawDescOnce.Do(func() {
		file_proto_audit_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_audit_proto_rawDescData)
	})
	return file_proto_audit_proto_rawDescData
}

var file_proto_audit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_audit_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: commonproto.Message
}
var file_proto_audit_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_audit_proto_init() }
func file_proto_audit_proto_init() {
	if File_proto_audit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_audit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_proto_audit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_audit_proto_goTypes,
		DependencyIndexes: file_proto_audit_proto_depIdxs,
		MessageInfos:      file_proto_audit_proto_msgTypes,
	}.Build()
	File_proto_audit_proto = out.File
	file_proto_audit_proto_rawDesc = nil
	file_proto_audit_proto_goTypes = nil
	file_proto_audit_proto_depIdxs = nil
}
