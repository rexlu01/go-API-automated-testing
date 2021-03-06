// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: proto/requestapi.proto

package protocol

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

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestName   string            `protobuf:"bytes,1,opt,name=requestName,proto3" json:"requestName,omitempty"`
	RequestURL    string            `protobuf:"bytes,2,opt,name=requestURL,proto3" json:"requestURL,omitempty"`
	RequestMethod string            `protobuf:"bytes,3,opt,name=requestMethod,proto3" json:"requestMethod,omitempty"`
	ContentType   string            `protobuf:"bytes,4,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Body          map[string]string `protobuf:"bytes,5,rep,name=body,proto3" json:"body,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Header        map[string]string `protobuf:"bytes,6,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IsPress       bool              `protobuf:"varint,7,opt,name=isPress,proto3" json:"isPress,omitempty"`
	RunTime       int64             `protobuf:"varint,8,opt,name=runTime,proto3" json:"runTime,omitempty"`
	RunTimes      int64             `protobuf:"varint,9,opt,name=runTimes,proto3" json:"runTimes,omitempty"`
	Concurrency   int64             `protobuf:"varint,10,opt,name=Concurrency,proto3" json:"Concurrency,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_requestapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_requestapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_proto_requestapi_proto_rawDescGZIP(), []int{0}
}

func (x *SendRequest) GetRequestName() string {
	if x != nil {
		return x.RequestName
	}
	return ""
}

func (x *SendRequest) GetRequestURL() string {
	if x != nil {
		return x.RequestURL
	}
	return ""
}

func (x *SendRequest) GetRequestMethod() string {
	if x != nil {
		return x.RequestMethod
	}
	return ""
}

func (x *SendRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *SendRequest) GetBody() map[string]string {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *SendRequest) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SendRequest) GetIsPress() bool {
	if x != nil {
		return x.IsPress
	}
	return false
}

func (x *SendRequest) GetRunTime() int64 {
	if x != nil {
		return x.RunTime
	}
	return 0
}

func (x *SendRequest) GetRunTimes() int64 {
	if x != nil {
		return x.RunTimes
	}
	return 0
}

func (x *SendRequest) GetConcurrency() int64 {
	if x != nil {
		return x.Concurrency
	}
	return 0
}

type GetRespons struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponsHeader map[string]string `protobuf:"bytes,1,rep,name=responsHeader,proto3" json:"responsHeader,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ResponsBody   map[string]string `protobuf:"bytes,2,rep,name=responsBody,proto3" json:"responsBody,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StatusCode    int64             `protobuf:"varint,3,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	IsPress       bool              `protobuf:"varint,4,opt,name=isPress,proto3" json:"isPress,omitempty"`
	TpsSet        map[string]string `protobuf:"bytes,5,rep,name=tpsSet,proto3" json:"tpsSet,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExecuteNum    map[string]string `protobuf:"bytes,6,rep,name=executeNum,proto3" json:"executeNum,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ResponseTime  float64           `protobuf:"fixed64,7,opt,name=responseTime,proto3" json:"responseTime,omitempty"`
	ExeuteTime    float64           `protobuf:"fixed64,8,opt,name=exeuteTime,proto3" json:"exeuteTime,omitempty"`
	SuccessRate   float64           `protobuf:"fixed64,9,opt,name=SuccessRate,proto3" json:"SuccessRate,omitempty"`
}

func (x *GetRespons) Reset() {
	*x = GetRespons{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_requestapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRespons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRespons) ProtoMessage() {}

func (x *GetRespons) ProtoReflect() protoreflect.Message {
	mi := &file_proto_requestapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRespons.ProtoReflect.Descriptor instead.
func (*GetRespons) Descriptor() ([]byte, []int) {
	return file_proto_requestapi_proto_rawDescGZIP(), []int{1}
}

func (x *GetRespons) GetResponsHeader() map[string]string {
	if x != nil {
		return x.ResponsHeader
	}
	return nil
}

func (x *GetRespons) GetResponsBody() map[string]string {
	if x != nil {
		return x.ResponsBody
	}
	return nil
}

func (x *GetRespons) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetRespons) GetIsPress() bool {
	if x != nil {
		return x.IsPress
	}
	return false
}

func (x *GetRespons) GetTpsSet() map[string]string {
	if x != nil {
		return x.TpsSet
	}
	return nil
}

func (x *GetRespons) GetExecuteNum() map[string]string {
	if x != nil {
		return x.ExecuteNum
	}
	return nil
}

func (x *GetRespons) GetResponseTime() float64 {
	if x != nil {
		return x.ResponseTime
	}
	return 0
}

func (x *GetRespons) GetExeuteTime() float64 {
	if x != nil {
		return x.ExeuteTime
	}
	return 0
}

func (x *GetRespons) GetSuccessRate() float64 {
	if x != nil {
		return x.SuccessRate
	}
	return 0
}

var File_proto_requestapi_proto protoreflect.FileDescriptor

var file_proto_requestapi_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x61, 0x70,
	0x69, 0x22, 0x8b, 0x04, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x52,
	0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x55, 0x52, 0x4c, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x42, 0x6f, 0x64, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12,
	0x48, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x50,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x50, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x72, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x1a, 0x37, 0x0a, 0x09, 0x42,
	0x6f, 0x64, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xfc, 0x05, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0x5c,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x72, 0x76, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x42, 0x6f,
	0x64, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x50, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x50, 0x72, 0x65, 0x73, 0x73, 0x12, 0x47,
	0x0a, 0x06, 0x74, 0x70, 0x73, 0x53, 0x65, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x70, 0x73, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x74, 0x70, 0x73, 0x53, 0x65, 0x74, 0x12, 0x53, 0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0c,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x65, 0x75, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x65, 0x75, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61,
	0x74, 0x65, 0x1a, 0x40, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x42,
	0x6f, 0x64, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x54, 0x70, 0x73, 0x53, 0x65, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3d, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x64,
	0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x50, 0x49, 0x12, 0x59, 0x0a, 0x0a, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x41, 0x50, 0x49, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x76, 0x69, 0x6c, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x73, 0x6e, 0x69, 0x74, 0x63, 0x68, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f,
	0x75, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_requestapi_proto_rawDescOnce sync.Once
	file_proto_requestapi_proto_rawDescData = file_proto_requestapi_proto_rawDesc
)

func file_proto_requestapi_proto_rawDescGZIP() []byte {
	file_proto_requestapi_proto_rawDescOnce.Do(func() {
		file_proto_requestapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_requestapi_proto_rawDescData)
	})
	return file_proto_requestapi_proto_rawDescData
}

var file_proto_requestapi_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_requestapi_proto_goTypes = []interface{}{
	(*SendRequest)(nil), // 0: go.micro.srv.requestapi.SendRequest
	(*GetRespons)(nil),  // 1: go.micro.srv.requestapi.GetRespons
	nil,                 // 2: go.micro.srv.requestapi.SendRequest.BodyEntry
	nil,                 // 3: go.micro.srv.requestapi.SendRequest.HeaderEntry
	nil,                 // 4: go.micro.srv.requestapi.GetRespons.ResponsHeaderEntry
	nil,                 // 5: go.micro.srv.requestapi.GetRespons.ResponsBodyEntry
	nil,                 // 6: go.micro.srv.requestapi.GetRespons.TpsSetEntry
	nil,                 // 7: go.micro.srv.requestapi.GetRespons.ExecuteNumEntry
}
var file_proto_requestapi_proto_depIdxs = []int32{
	2, // 0: go.micro.srv.requestapi.SendRequest.body:type_name -> go.micro.srv.requestapi.SendRequest.BodyEntry
	3, // 1: go.micro.srv.requestapi.SendRequest.header:type_name -> go.micro.srv.requestapi.SendRequest.HeaderEntry
	4, // 2: go.micro.srv.requestapi.GetRespons.responsHeader:type_name -> go.micro.srv.requestapi.GetRespons.ResponsHeaderEntry
	5, // 3: go.micro.srv.requestapi.GetRespons.responsBody:type_name -> go.micro.srv.requestapi.GetRespons.ResponsBodyEntry
	6, // 4: go.micro.srv.requestapi.GetRespons.tpsSet:type_name -> go.micro.srv.requestapi.GetRespons.TpsSetEntry
	7, // 5: go.micro.srv.requestapi.GetRespons.executeNum:type_name -> go.micro.srv.requestapi.GetRespons.ExecuteNumEntry
	0, // 6: go.micro.srv.requestapi.SendAPI.ProcessAPI:input_type -> go.micro.srv.requestapi.SendRequest
	1, // 7: go.micro.srv.requestapi.SendAPI.ProcessAPI:output_type -> go.micro.srv.requestapi.GetRespons
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_requestapi_proto_init() }
func file_proto_requestapi_proto_init() {
	if File_proto_requestapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_requestapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_proto_requestapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRespons); i {
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
			RawDescriptor: file_proto_requestapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_requestapi_proto_goTypes,
		DependencyIndexes: file_proto_requestapi_proto_depIdxs,
		MessageInfos:      file_proto_requestapi_proto_msgTypes,
	}.Build()
	File_proto_requestapi_proto = out.File
	file_proto_requestapi_proto_rawDesc = nil
	file_proto_requestapi_proto_goTypes = nil
	file_proto_requestapi_proto_depIdxs = nil
}
