// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: shitlist/v1/shitlist.proto

package shitlistv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// GreetRequest is a request to perform a greeting.
type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the person to greet.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shitlist_v1_shitlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shitlist_v1_shitlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_shitlist_v1_shitlist_proto_rawDescGZIP(), []int{0}
}

func (x *GreetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// GreetResponse is a greeting response.
type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// greeting for the user that was requested.
	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shitlist_v1_shitlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shitlist_v1_shitlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_shitlist_v1_shitlist_proto_rawDescGZIP(), []int{1}
}

func (x *GreetResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

// ClickRequest is a request to record a click event.
type ClickRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_id of the user to record a click event for.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ClickRequest) Reset() {
	*x = ClickRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shitlist_v1_shitlist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickRequest) ProtoMessage() {}

func (x *ClickRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shitlist_v1_shitlist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickRequest.ProtoReflect.Descriptor instead.
func (*ClickRequest) Descriptor() ([]byte, []int) {
	return file_shitlist_v1_shitlist_proto_rawDescGZIP(), []int{2}
}

func (x *ClickRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// ClickResponse is a response to a click event.
type ClickResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// clicks recorded for the user.
	Clicks int64 `protobuf:"varint,1,opt,name=clicks,proto3" json:"clicks,omitempty"`
}

func (x *ClickResponse) Reset() {
	*x = ClickResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shitlist_v1_shitlist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickResponse) ProtoMessage() {}

func (x *ClickResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shitlist_v1_shitlist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickResponse.ProtoReflect.Descriptor instead.
func (*ClickResponse) Descriptor() ([]byte, []int) {
	return file_shitlist_v1_shitlist_proto_rawDescGZIP(), []int{3}
}

func (x *ClickResponse) GetClicks() int64 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

// LeadersRequest is a request for the top clickers.
type LeadersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeadersRequest) Reset() {
	*x = LeadersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shitlist_v1_shitlist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeadersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeadersRequest) ProtoMessage() {}

func (x *LeadersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shitlist_v1_shitlist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeadersRequest.ProtoReflect.Descriptor instead.
func (*LeadersRequest) Descriptor() ([]byte, []int) {
	return file_shitlist_v1_shitlist_proto_rawDescGZIP(), []int{4}
}

// Clicker represents a single clicker user.
type Clicker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_id of the user thats clicking.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// clicks is the number of times the user has clicked.
	Clicks int64 `protobuf:"varint,2,opt,name=clicks,proto3" json:"clicks,omitempty"`
}

func (x *Clicker) Reset() {
	*x = Clicker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shitlist_v1_shitlist_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clicker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clicker) ProtoMessage() {}

func (x *Clicker) ProtoReflect() protoreflect.Message {
	mi := &file_shitlist_v1_shitlist_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clicker.ProtoReflect.Descriptor instead.
func (*Clicker) Descriptor() ([]byte, []int) {
	return file_shitlist_v1_shitlist_proto_rawDescGZIP(), []int{5}
}

func (x *Clicker) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Clicker) GetClicks() int64 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

// LeadersResponse is the top clickers.
type LeadersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// top_clickers are the top 10 clicking users.
	TopClickers []*Clicker `protobuf:"bytes,1,rep,name=top_clickers,json=topClickers,proto3" json:"top_clickers,omitempty"`
}

func (x *LeadersResponse) Reset() {
	*x = LeadersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shitlist_v1_shitlist_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeadersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeadersResponse) ProtoMessage() {}

func (x *LeadersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shitlist_v1_shitlist_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeadersResponse.ProtoReflect.Descriptor instead.
func (*LeadersResponse) Descriptor() ([]byte, []int) {
	return file_shitlist_v1_shitlist_proto_rawDescGZIP(), []int{6}
}

func (x *LeadersResponse) GetTopClickers() []*Clicker {
	if x != nil {
		return x.TopClickers
	}
	return nil
}

var File_shitlist_v1_shitlist_proto protoreflect.FileDescriptor

var file_shitlist_v1_shitlist_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x68, 0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68,
	0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x68,
	0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x22, 0x31, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x22,
	0x10, 0x0a, 0x0e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x44, 0x0a, 0x07, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x22, 0x4a, 0x0a, 0x0f, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x74, 0x6f,
	0x70, 0x5f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x68, 0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x0b, 0x74, 0x6f, 0x70, 0x43, 0x6c, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x73, 0x32, 0xdd, 0x01, 0x0a, 0x0f, 0x53, 0x68, 0x69, 0x74, 0x6c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x12, 0x19, 0x2e, 0x73, 0x68, 0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x68,
	0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x05, 0x43, 0x6c, 0x69,
	0x63, 0x6b, 0x12, 0x19, 0x2e, 0x73, 0x68, 0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x68, 0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x07, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x68, 0x69, 0x74, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x68, 0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x69, 0x6b, 0x65, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x68, 0x69, 0x74, 0x6c, 0x69,
	0x73, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x68, 0x69, 0x74, 0x6c, 0x69,
	0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shitlist_v1_shitlist_proto_rawDescOnce sync.Once
	file_shitlist_v1_shitlist_proto_rawDescData = file_shitlist_v1_shitlist_proto_rawDesc
)

func file_shitlist_v1_shitlist_proto_rawDescGZIP() []byte {
	file_shitlist_v1_shitlist_proto_rawDescOnce.Do(func() {
		file_shitlist_v1_shitlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_shitlist_v1_shitlist_proto_rawDescData)
	})
	return file_shitlist_v1_shitlist_proto_rawDescData
}

var file_shitlist_v1_shitlist_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_shitlist_v1_shitlist_proto_goTypes = []interface{}{
	(*GreetRequest)(nil),    // 0: shitlist.v1.GreetRequest
	(*GreetResponse)(nil),   // 1: shitlist.v1.GreetResponse
	(*ClickRequest)(nil),    // 2: shitlist.v1.ClickRequest
	(*ClickResponse)(nil),   // 3: shitlist.v1.ClickResponse
	(*LeadersRequest)(nil),  // 4: shitlist.v1.LeadersRequest
	(*Clicker)(nil),         // 5: shitlist.v1.Clicker
	(*LeadersResponse)(nil), // 6: shitlist.v1.LeadersResponse
}
var file_shitlist_v1_shitlist_proto_depIdxs = []int32{
	5, // 0: shitlist.v1.LeadersResponse.top_clickers:type_name -> shitlist.v1.Clicker
	0, // 1: shitlist.v1.ShitlistService.Greet:input_type -> shitlist.v1.GreetRequest
	2, // 2: shitlist.v1.ShitlistService.Click:input_type -> shitlist.v1.ClickRequest
	4, // 3: shitlist.v1.ShitlistService.Leaders:input_type -> shitlist.v1.LeadersRequest
	1, // 4: shitlist.v1.ShitlistService.Greet:output_type -> shitlist.v1.GreetResponse
	3, // 5: shitlist.v1.ShitlistService.Click:output_type -> shitlist.v1.ClickResponse
	6, // 6: shitlist.v1.ShitlistService.Leaders:output_type -> shitlist.v1.LeadersResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shitlist_v1_shitlist_proto_init() }
func file_shitlist_v1_shitlist_proto_init() {
	if File_shitlist_v1_shitlist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shitlist_v1_shitlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequest); i {
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
		file_shitlist_v1_shitlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponse); i {
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
		file_shitlist_v1_shitlist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickRequest); i {
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
		file_shitlist_v1_shitlist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickResponse); i {
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
		file_shitlist_v1_shitlist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeadersRequest); i {
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
		file_shitlist_v1_shitlist_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clicker); i {
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
		file_shitlist_v1_shitlist_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeadersResponse); i {
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
			RawDescriptor: file_shitlist_v1_shitlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shitlist_v1_shitlist_proto_goTypes,
		DependencyIndexes: file_shitlist_v1_shitlist_proto_depIdxs,
		MessageInfos:      file_shitlist_v1_shitlist_proto_msgTypes,
	}.Build()
	File_shitlist_v1_shitlist_proto = out.File
	file_shitlist_v1_shitlist_proto_rawDesc = nil
	file_shitlist_v1_shitlist_proto_goTypes = nil
	file_shitlist_v1_shitlist_proto_depIdxs = nil
}
