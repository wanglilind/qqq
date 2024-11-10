// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: api/proto/identity.proto

package identity

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

type RegisterIdentityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                      // 用户ID
	BiometricData []byte `protobuf:"bytes,2,opt,name=biometric_data,json=biometricData,proto3" json:"biometric_data,omitempty"` // 生物特征数据
	NationalId    string `protobuf:"bytes,3,opt,name=national_id,json=nationalId,proto3" json:"national_id,omitempty"`          // 国家身份证号
	CountryCode   string `protobuf:"bytes,4,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`       // 国家代码
	BirthDate     string `protobuf:"bytes,5,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`             // 出生日期
}

func (x *RegisterIdentityRequest) Reset() {
	*x = RegisterIdentityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterIdentityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterIdentityRequest) ProtoMessage() {}

func (x *RegisterIdentityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterIdentityRequest.ProtoReflect.Descriptor instead.
func (*RegisterIdentityRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_identity_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterIdentityRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RegisterIdentityRequest) GetBiometricData() []byte {
	if x != nil {
		return x.BiometricData
	}
	return nil
}

func (x *RegisterIdentityRequest) GetNationalId() string {
	if x != nil {
		return x.NationalId
	}
	return ""
}

func (x *RegisterIdentityRequest) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *RegisterIdentityRequest) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

type RegisterIdentityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentityId string `protobuf:"bytes,1,opt,name=identity_id,json=identityId,proto3" json:"identity_id,omitempty"` // 身份ID
	Status     string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`                           // 状态
}

func (x *RegisterIdentityResponse) Reset() {
	*x = RegisterIdentityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterIdentityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterIdentityResponse) ProtoMessage() {}

func (x *RegisterIdentityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterIdentityResponse.ProtoReflect.Descriptor instead.
func (*RegisterIdentityResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_identity_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterIdentityResponse) GetIdentityId() string {
	if x != nil {
		return x.IdentityId
	}
	return ""
}

func (x *RegisterIdentityResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type VerifyIdentityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentityId    string `protobuf:"bytes,1,opt,name=identity_id,json=identityId,proto3" json:"identity_id,omitempty"`          // 身份ID
	BiometricData []byte `protobuf:"bytes,2,opt,name=biometric_data,json=biometricData,proto3" json:"biometric_data,omitempty"` // 生物特征数据
}

func (x *VerifyIdentityRequest) Reset() {
	*x = VerifyIdentityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyIdentityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyIdentityRequest) ProtoMessage() {}

func (x *VerifyIdentityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyIdentityRequest.ProtoReflect.Descriptor instead.
func (*VerifyIdentityRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_identity_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyIdentityRequest) GetIdentityId() string {
	if x != nil {
		return x.IdentityId
	}
	return ""
}

func (x *VerifyIdentityRequest) GetBiometricData() []byte {
	if x != nil {
		return x.BiometricData
	}
	return nil
}

type VerifyIdentityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid      bool   `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`                             // 验证结果
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                          // 验证信息
	VerifyTime int64  `protobuf:"varint,3,opt,name=verify_time,json=verifyTime,proto3" json:"verify_time,omitempty"` // 验证时间
}

func (x *VerifyIdentityResponse) Reset() {
	*x = VerifyIdentityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_identity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyIdentityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyIdentityResponse) ProtoMessage() {}

func (x *VerifyIdentityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_identity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyIdentityResponse.ProtoReflect.Descriptor instead.
func (*VerifyIdentityResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_identity_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyIdentityResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *VerifyIdentityResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *VerifyIdentityResponse) GetVerifyTime() int64 {
	if x != nil {
		return x.VerifyTime
	}
	return 0
}

type GetIdentityStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentityId string `protobuf:"bytes,1,opt,name=identity_id,json=identityId,proto3" json:"identity_id,omitempty"` // 身份ID
}

func (x *GetIdentityStatusRequest) Reset() {
	*x = GetIdentityStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_identity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdentityStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdentityStatusRequest) ProtoMessage() {}

func (x *GetIdentityStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_identity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdentityStatusRequest.ProtoReflect.Descriptor instead.
func (*GetIdentityStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_identity_proto_rawDescGZIP(), []int{4}
}

func (x *GetIdentityStatusRequest) GetIdentityId() string {
	if x != nil {
		return x.IdentityId
	}
	return ""
}

type GetIdentityStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`                                  // 身份状态
	LastActivity int64  `protobuf:"varint,2,opt,name=last_activity,json=lastActivity,proto3" json:"last_activity,omitempty"` // 最后活动时间
	IsActive     bool   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`             // 是否活跃
}

func (x *GetIdentityStatusResponse) Reset() {
	*x = GetIdentityStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_identity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdentityStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdentityStatusResponse) ProtoMessage() {}

func (x *GetIdentityStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_identity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdentityStatusResponse.ProtoReflect.Descriptor instead.
func (*GetIdentityStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_identity_proto_rawDescGZIP(), []int{5}
}

func (x *GetIdentityStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetIdentityStatusResponse) GetLastActivity() int64 {
	if x != nil {
		return x.LastActivity
	}
	return 0
}

func (x *GetIdentityStatusResponse) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type UpdateBiometricDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentityId       string `protobuf:"bytes,1,opt,name=identity_id,json=identityId,proto3" json:"identity_id,omitempty"`                     // 身份ID
	NewBiometricData []byte `protobuf:"bytes,2,opt,name=new_biometric_data,json=newBiometricData,proto3" json:"new_biometric_data,omitempty"` // 新的生物特征数据
}

func (x *UpdateBiometricDataRequest) Reset() {
	*x = UpdateBiometricDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_identity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBiometricDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBiometricDataRequest) ProtoMessage() {}

func (x *UpdateBiometricDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_identity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBiometricDataRequest.ProtoReflect.Descriptor instead.
func (*UpdateBiometricDataRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_identity_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBiometricDataRequest) GetIdentityId() string {
	if x != nil {
		return x.IdentityId
	}
	return ""
}

func (x *UpdateBiometricDataRequest) GetNewBiometricData() []byte {
	if x != nil {
		return x.NewBiometricData
	}
	return nil
}

type UpdateBiometricDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                         // 更新结果
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                          // 更新信息
	UpdateTime int64  `protobuf:"varint,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"` // 更新时间
}

func (x *UpdateBiometricDataResponse) Reset() {
	*x = UpdateBiometricDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_identity_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBiometricDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBiometricDataResponse) ProtoMessage() {}

func (x *UpdateBiometricDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_identity_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBiometricDataResponse.ProtoReflect.Descriptor instead.
func (*UpdateBiometricDataResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_identity_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateBiometricDataResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateBiometricDataResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateBiometricDataResponse) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

var File_api_proto_identity_proto protoreflect.FileDescriptor

var file_api_proto_identity_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0xbc, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x69, 0x6f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x62, 0x69, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x65, 0x22, 0x53, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5f, 0x0a, 0x15, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x69, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x62, 0x69, 0x6f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x22, 0x69, 0x0a, 0x16, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x22, 0x75, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x6b, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x69, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x65, 0x77, 0x5f, 0x62,
	0x69, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x10, 0x6e, 0x65, 0x77, 0x42, 0x69, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x44, 0x61, 0x74, 0x61, 0x22, 0x72, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x69, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x83, 0x03, 0x0a, 0x0f, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a,
	0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x21, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x22, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x24, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61,
	0x6e, 0x67, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x64, 0x2f, 0x71, 0x71, 0x71, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_identity_proto_rawDescOnce sync.Once
	file_api_proto_identity_proto_rawDescData = file_api_proto_identity_proto_rawDesc
)

func file_api_proto_identity_proto_rawDescGZIP() []byte {
	file_api_proto_identity_proto_rawDescOnce.Do(func() {
		file_api_proto_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_identity_proto_rawDescData)
	})
	return file_api_proto_identity_proto_rawDescData
}

var file_api_proto_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_identity_proto_goTypes = []interface{}{
	(*RegisterIdentityRequest)(nil),     // 0: identity.RegisterIdentityRequest
	(*RegisterIdentityResponse)(nil),    // 1: identity.RegisterIdentityResponse
	(*VerifyIdentityRequest)(nil),       // 2: identity.VerifyIdentityRequest
	(*VerifyIdentityResponse)(nil),      // 3: identity.VerifyIdentityResponse
	(*GetIdentityStatusRequest)(nil),    // 4: identity.GetIdentityStatusRequest
	(*GetIdentityStatusResponse)(nil),   // 5: identity.GetIdentityStatusResponse
	(*UpdateBiometricDataRequest)(nil),  // 6: identity.UpdateBiometricDataRequest
	(*UpdateBiometricDataResponse)(nil), // 7: identity.UpdateBiometricDataResponse
}
var file_api_proto_identity_proto_depIdxs = []int32{
	0, // 0: identity.IdentityService.RegisterIdentity:input_type -> identity.RegisterIdentityRequest
	2, // 1: identity.IdentityService.VerifyIdentity:input_type -> identity.VerifyIdentityRequest
	4, // 2: identity.IdentityService.GetIdentityStatus:input_type -> identity.GetIdentityStatusRequest
	6, // 3: identity.IdentityService.UpdateBiometricData:input_type -> identity.UpdateBiometricDataRequest
	1, // 4: identity.IdentityService.RegisterIdentity:output_type -> identity.RegisterIdentityResponse
	3, // 5: identity.IdentityService.VerifyIdentity:output_type -> identity.VerifyIdentityResponse
	5, // 6: identity.IdentityService.GetIdentityStatus:output_type -> identity.GetIdentityStatusResponse
	7, // 7: identity.IdentityService.UpdateBiometricData:output_type -> identity.UpdateBiometricDataResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_identity_proto_init() }
func file_api_proto_identity_proto_init() {
	if File_api_proto_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterIdentityRequest); i {
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
		file_api_proto_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterIdentityResponse); i {
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
		file_api_proto_identity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyIdentityRequest); i {
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
		file_api_proto_identity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyIdentityResponse); i {
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
		file_api_proto_identity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdentityStatusRequest); i {
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
		file_api_proto_identity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdentityStatusResponse); i {
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
		file_api_proto_identity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBiometricDataRequest); i {
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
		file_api_proto_identity_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBiometricDataResponse); i {
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
			RawDescriptor: file_api_proto_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_identity_proto_goTypes,
		DependencyIndexes: file_api_proto_identity_proto_depIdxs,
		MessageInfos:      file_api_proto_identity_proto_msgTypes,
	}.Build()
	File_api_proto_identity_proto = out.File
	file_api_proto_identity_proto_rawDesc = nil
	file_api_proto_identity_proto_goTypes = nil
	file_api_proto_identity_proto_depIdxs = nil
}
