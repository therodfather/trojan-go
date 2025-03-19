// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api.proto

package service

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

type SetUsersRequest_Operation int32

const (
	SetUsersRequest_Add    SetUsersRequest_Operation = 0
	SetUsersRequest_Delete SetUsersRequest_Operation = 1
	SetUsersRequest_Modify SetUsersRequest_Operation = 2
)

// Enum value maps for SetUsersRequest_Operation.
var (
	SetUsersRequest_Operation_name = map[int32]string{
		0: "Add",
		1: "Delete",
		2: "Modify",
	}
	SetUsersRequest_Operation_value = map[string]int32{
		"Add":    0,
		"Delete": 1,
		"Modify": 2,
	}
)

func (x SetUsersRequest_Operation) Enum() *SetUsersRequest_Operation {
	p := new(SetUsersRequest_Operation)
	*p = x
	return p
}

func (x SetUsersRequest_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetUsersRequest_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (SetUsersRequest_Operation) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x SetUsersRequest_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetUsersRequest_Operation.Descriptor instead.
func (SetUsersRequest_Operation) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{10, 0}
}

type Traffic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadTraffic   uint64 `protobuf:"varint,1,opt,name=upload_traffic,json=uploadTraffic,proto3" json:"upload_traffic,omitempty"`
	DownloadTraffic uint64 `protobuf:"varint,2,opt,name=download_traffic,json=downloadTraffic,proto3" json:"download_traffic,omitempty"`
}

func (x *Traffic) Reset() {
	*x = Traffic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Traffic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Traffic) ProtoMessage() {}

func (x *Traffic) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Traffic.ProtoReflect.Descriptor instead.
func (*Traffic) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *Traffic) GetUploadTraffic() uint64 {
	if x != nil {
		return x.UploadTraffic
	}
	return 0
}

func (x *Traffic) GetDownloadTraffic() uint64 {
	if x != nil {
		return x.DownloadTraffic
	}
	return 0
}

type Speed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadSpeed   uint64 `protobuf:"varint,1,opt,name=upload_speed,json=uploadSpeed,proto3" json:"upload_speed,omitempty"`
	DownloadSpeed uint64 `protobuf:"varint,2,opt,name=download_speed,json=downloadSpeed,proto3" json:"download_speed,omitempty"`
}

func (x *Speed) Reset() {
	*x = Speed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Speed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Speed) ProtoMessage() {}

func (x *Speed) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Speed.ProtoReflect.Descriptor instead.
func (*Speed) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *Speed) GetUploadSpeed() uint64 {
	if x != nil {
		return x.UploadSpeed
	}
	return 0
}

func (x *Speed) GetDownloadSpeed() uint64 {
	if x != nil {
		return x.DownloadSpeed
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Hash     string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type UserStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	TrafficTotal *Traffic `protobuf:"bytes,2,opt,name=traffic_total,json=trafficTotal,proto3" json:"traffic_total,omitempty"`
	SpeedCurrent *Speed   `protobuf:"bytes,3,opt,name=speed_current,json=speedCurrent,proto3" json:"speed_current,omitempty"`
	SpeedLimit   *Speed   `protobuf:"bytes,4,opt,name=speed_limit,json=speedLimit,proto3" json:"speed_limit,omitempty"`
	IpCurrent    int32    `protobuf:"varint,5,opt,name=ip_current,json=ipCurrent,proto3" json:"ip_current,omitempty"`
	IpLimit      int32    `protobuf:"varint,6,opt,name=ip_limit,json=ipLimit,proto3" json:"ip_limit,omitempty"`
}

func (x *UserStatus) Reset() {
	*x = UserStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStatus) ProtoMessage() {}

func (x *UserStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStatus.ProtoReflect.Descriptor instead.
func (*UserStatus) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *UserStatus) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserStatus) GetTrafficTotal() *Traffic {
	if x != nil {
		return x.TrafficTotal
	}
	return nil
}

func (x *UserStatus) GetSpeedCurrent() *Speed {
	if x != nil {
		return x.SpeedCurrent
	}
	return nil
}

func (x *UserStatus) GetSpeedLimit() *Speed {
	if x != nil {
		return x.SpeedLimit
	}
	return nil
}

func (x *UserStatus) GetIpCurrent() int32 {
	if x != nil {
		return x.IpCurrent
	}
	return 0
}

func (x *UserStatus) GetIpLimit() int32 {
	if x != nil {
		return x.IpLimit
	}
	return 0
}

type GetTrafficRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetTrafficRequest) Reset() {
	*x = GetTrafficRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrafficRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrafficRequest) ProtoMessage() {}

func (x *GetTrafficRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrafficRequest.ProtoReflect.Descriptor instead.
func (*GetTrafficRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetTrafficRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetTrafficResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Info         string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	TrafficTotal *Traffic `protobuf:"bytes,3,opt,name=traffic_total,json=trafficTotal,proto3" json:"traffic_total,omitempty"`
	SpeedCurrent *Speed   `protobuf:"bytes,4,opt,name=speed_current,json=speedCurrent,proto3" json:"speed_current,omitempty"`
}

func (x *GetTrafficResponse) Reset() {
	*x = GetTrafficResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrafficResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrafficResponse) ProtoMessage() {}

func (x *GetTrafficResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrafficResponse.ProtoReflect.Descriptor instead.
func (*GetTrafficResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetTrafficResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetTrafficResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *GetTrafficResponse) GetTrafficTotal() *Traffic {
	if x != nil {
		return x.TrafficTotal
	}
	return nil
}

func (x *GetTrafficResponse) GetSpeedCurrent() *Speed {
	if x != nil {
		return x.SpeedCurrent
	}
	return nil
}

type ListUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

type ListUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *UserStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *ListUsersResponse) GetStatus() *UserStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type GetUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *GetUsersRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool        `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Info    string      `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Status  *UserStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

func (x *GetUsersResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetUsersResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *GetUsersResponse) GetStatus() *UserStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type SetUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    *UserStatus               `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Operation SetUsersRequest_Operation `protobuf:"varint,2,opt,name=operation,proto3,enum=trojan.api.SetUsersRequest_Operation" json:"operation,omitempty"`
}

func (x *SetUsersRequest) Reset() {
	*x = SetUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUsersRequest) ProtoMessage() {}

func (x *SetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUsersRequest.ProtoReflect.Descriptor instead.
func (*SetUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{10}
}

func (x *SetUsersRequest) GetStatus() *UserStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *SetUsersRequest) GetOperation() SetUsersRequest_Operation {
	if x != nil {
		return x.Operation
	}
	return SetUsersRequest_Add
}

type SetUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Info    string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SetUsersResponse) Reset() {
	*x = SetUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUsersResponse) ProtoMessage() {}

func (x *SetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUsersResponse.ProtoReflect.Descriptor instead.
func (*SetUsersResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{11}
}

func (x *SetUsersResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SetUsersResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x72, 0x6f,
	0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x5b, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x22, 0x51, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x22, 0x36, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22,
	0x92, 0x02, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74,
	0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x72,
	0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x36,
	0x0a, 0x0d, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x0c, 0x73, 0x70, 0x65, 0x65, 0x64, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72,
	0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x0a,
	0x73, 0x70, 0x65, 0x65, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x69, 0x70, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x70, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x70, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x39, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0xb4, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x72,
	0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x36,
	0x0a, 0x0d, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x0c, 0x73, 0x70, 0x65, 0x65, 0x64, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x70, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x6f,
	0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x0f, 0x53,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x43,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x25, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x10,
	0x02, 0x22, 0x40, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x32, 0x64, 0x0a, 0x13, 0x54, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x1d, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xfd, 0x01, 0x0a, 0x13, 0x54, 0x72,
	0x6f, 0x6a, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4c, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1c,
	0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74,
	0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x4b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x2e, 0x74, 0x72,
	0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4b, 0x0a, 0x08,
	0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x34, 0x67, 0x65, 0x66, 0x61, 0x75, 0x31,
	0x74, 0x2f, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_proto_goTypes = []interface{}{
	(SetUsersRequest_Operation)(0), // 0: trojan.api.SetUsersRequest.Operation
	(*Traffic)(nil),                // 1: trojan.api.Traffic
	(*Speed)(nil),                  // 2: trojan.api.Speed
	(*User)(nil),                   // 3: trojan.api.User
	(*UserStatus)(nil),             // 4: trojan.api.UserStatus
	(*GetTrafficRequest)(nil),      // 5: trojan.api.GetTrafficRequest
	(*GetTrafficResponse)(nil),     // 6: trojan.api.GetTrafficResponse
	(*ListUsersRequest)(nil),       // 7: trojan.api.ListUsersRequest
	(*ListUsersResponse)(nil),      // 8: trojan.api.ListUsersResponse
	(*GetUsersRequest)(nil),        // 9: trojan.api.GetUsersRequest
	(*GetUsersResponse)(nil),       // 10: trojan.api.GetUsersResponse
	(*SetUsersRequest)(nil),        // 11: trojan.api.SetUsersRequest
	(*SetUsersResponse)(nil),       // 12: trojan.api.SetUsersResponse
}
var file_api_proto_depIdxs = []int32{
	3,  // 0: trojan.api.UserStatus.user:type_name -> trojan.api.User
	1,  // 1: trojan.api.UserStatus.traffic_total:type_name -> trojan.api.Traffic
	2,  // 2: trojan.api.UserStatus.speed_current:type_name -> trojan.api.Speed
	2,  // 3: trojan.api.UserStatus.speed_limit:type_name -> trojan.api.Speed
	3,  // 4: trojan.api.GetTrafficRequest.user:type_name -> trojan.api.User
	1,  // 5: trojan.api.GetTrafficResponse.traffic_total:type_name -> trojan.api.Traffic
	2,  // 6: trojan.api.GetTrafficResponse.speed_current:type_name -> trojan.api.Speed
	4,  // 7: trojan.api.ListUsersResponse.status:type_name -> trojan.api.UserStatus
	3,  // 8: trojan.api.GetUsersRequest.user:type_name -> trojan.api.User
	4,  // 9: trojan.api.GetUsersResponse.status:type_name -> trojan.api.UserStatus
	4,  // 10: trojan.api.SetUsersRequest.status:type_name -> trojan.api.UserStatus
	0,  // 11: trojan.api.SetUsersRequest.operation:type_name -> trojan.api.SetUsersRequest.Operation
	5,  // 12: trojan.api.TrojanClientService.GetTraffic:input_type -> trojan.api.GetTrafficRequest
	7,  // 13: trojan.api.TrojanServerService.ListUsers:input_type -> trojan.api.ListUsersRequest
	9,  // 14: trojan.api.TrojanServerService.GetUsers:input_type -> trojan.api.GetUsersRequest
	11, // 15: trojan.api.TrojanServerService.SetUsers:input_type -> trojan.api.SetUsersRequest
	6,  // 16: trojan.api.TrojanClientService.GetTraffic:output_type -> trojan.api.GetTrafficResponse
	8,  // 17: trojan.api.TrojanServerService.ListUsers:output_type -> trojan.api.ListUsersResponse
	10, // 18: trojan.api.TrojanServerService.GetUsers:output_type -> trojan.api.GetUsersResponse
	12, // 19: trojan.api.TrojanServerService.SetUsers:output_type -> trojan.api.SetUsersResponse
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Traffic); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Speed); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStatus); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrafficRequest); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrafficResponse); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersRequest); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersResponse); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRequest); i {
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
		file_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersResponse); i {
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
		file_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUsersRequest); i {
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
		file_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUsersResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
