// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/v1/idp_service.proto

package apiv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IdentityProvider_Type int32

const (
	IdentityProvider_TYPE_UNSPECIFIED IdentityProvider_Type = 0
	IdentityProvider_OAUTH2           IdentityProvider_Type = 1
)

// Enum value maps for IdentityProvider_Type.
var (
	IdentityProvider_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "OAUTH2",
	}
	IdentityProvider_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"OAUTH2":           1,
	}
)

func (x IdentityProvider_Type) Enum() *IdentityProvider_Type {
	p := new(IdentityProvider_Type)
	*p = x
	return p
}

func (x IdentityProvider_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentityProvider_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_idp_service_proto_enumTypes[0].Descriptor()
}

func (IdentityProvider_Type) Type() protoreflect.EnumType {
	return &file_api_v1_idp_service_proto_enumTypes[0]
}

func (x IdentityProvider_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentityProvider_Type.Descriptor instead.
func (IdentityProvider_Type) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_idp_service_proto_rawDescGZIP(), []int{0, 0}
}

type IdentityProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the identityProvider.
	// Format: identityProviders/{id}
	Name             string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type             IdentityProvider_Type   `protobuf:"varint,2,opt,name=type,proto3,enum=memos.api.v1.IdentityProvider_Type" json:"type,omitempty"`
	Title            string                  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	IdentifierFilter string                  `protobuf:"bytes,4,opt,name=identifier_filter,json=identifierFilter,proto3" json:"identifier_filter,omitempty"`
	Config           *IdentityProviderConfig `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *IdentityProvider) Reset() {
	*x = IdentityProvider{}
	mi := &file_api_v1_idp_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdentityProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProvider) ProtoMessage() {}

func (x *IdentityProvider) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_idp_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProvider.ProtoReflect.Descriptor instead.
func (*IdentityProvider) Descriptor() ([]byte, []int) {
	return file_api_v1_idp_service_proto_rawDescGZIP(), []int{0}
}

func (x *IdentityProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IdentityProvider) GetType() IdentityProvider_Type {
	if x != nil {
		return x.Type
	}
	return IdentityProvider_TYPE_UNSPECIFIED
}

func (x *IdentityProvider) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *IdentityProvider) GetIdentifierFilter() string {
	if x != nil {
		return x.IdentifierFilter
	}
	return ""
}

func (x *IdentityProvider) GetConfig() *IdentityProviderConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type IdentityProviderConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Config:
	//
	//	*IdentityProviderConfig_Oauth2Config
	Config isIdentityProviderConfig_Config `protobuf_oneof:"config"`
}

func (x *IdentityProviderConfig) Reset() {
	*x = IdentityProviderConfig{}
	mi := &file_api_v1_idp_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdentityProviderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProviderConfig) ProtoMessage() {}

func (x *IdentityProviderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_idp_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProviderConfig.ProtoReflect.Descriptor instead.
func (*IdentityProviderConfig) Descriptor() ([]byte, []int) {
	return file_api_v1_idp_service_proto_rawDescGZIP(), []int{1}
}

func (m *IdentityProviderConfig) GetConfig() isIdentityProviderConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *IdentityProviderConfig) GetOauth2Config() *OAuth2Config {
	if x, ok := x.GetConfig().(*IdentityProviderConfig_Oauth2Config); ok {
		return x.Oauth2Config
	}
	return nil
}

type isIdentityProviderConfig_Config interface {
	isIdentityProviderConfig_Config()
}

type IdentityProviderConfig_Oauth2Config struct {
	Oauth2Config *OAuth2Config `protobuf:"bytes,1,opt,name=oauth2_config,json=oauth2Config,proto3,oneof"`
}

func (*IdentityProviderConfig_Oauth2Config) isIdentityProviderConfig_Config() {}

type FieldMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier  string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FieldMapping) Reset() {
	*x = FieldMapping{}
	mi := &file_api_v1_idp_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMapping) ProtoMessage() {}

func (x *FieldMapping) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_idp_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMapping.ProtoReflect.Descriptor instead.
func (*FieldMapping) Descriptor() ([]byte, []int) {
	return file_api_v1_idp_service_proto_rawDescGZIP(), []int{2}
}

func (x *FieldMapping) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *FieldMapping) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *FieldMapping) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type OAuth2Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string        `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string        `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	AuthUrl      string        `protobuf:"bytes,3,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	TokenUrl     string        `protobuf:"bytes,4,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	UserInfoUrl  string        `protobuf:"bytes,5,opt,name=user_info_url,json=userInfoUrl,proto3" json:"user_info_url,omitempty"`
	Scopes       []string      `protobuf:"bytes,6,rep,name=scopes,proto3" json:"scopes,omitempty"`
	FieldMapping *FieldMapping `protobuf:"bytes,7,opt,name=field_mapping,json=fieldMapping,proto3" json:"field_mapping,omitempty"`
}

func (x *OAuth2Config) Reset() {
	*x = OAuth2Config{}
	mi := &file_api_v1_idp_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OAuth2Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2Config) ProtoMessage() {}

func (x *OAuth2Config) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_idp_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2Config.ProtoReflect.Descriptor instead.
func (*OAuth2Config) Descriptor() ([]byte, []int) {
	return file_api_v1_idp_service_proto_rawDescGZIP(), []int{3}
}

func (x *OAuth2Config) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *OAuth2Config) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *OAuth2Config) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

func (x *OAuth2Config) GetTokenUrl() string {
	if x != nil {
		return x.TokenUrl
	}
	return ""
}

func (x *OAuth2Config) GetUserInfoUrl() string {
	if x != nil {
		return x.UserInfoUrl
	}
	return ""
}

func (x *OAuth2Config) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *OAuth2Config) GetFieldMapping() *FieldMapping {
	if x != nil {
		return x.FieldMapping
	}
	return nil
}

type ListIdentityProvidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListIdentityProvidersRequest) Reset() {
	*x = ListIdentityProvidersRequest{}
	mi := &file_api_v1_idp_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListIdentityProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIdentityProvidersRequest) ProtoMessage() {}

func (x *ListIdentityProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_idp_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIdentityProvidersRequest.ProtoReflect.Descriptor instead.
func (*ListIdentityProvidersRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_idp_service_proto_rawDescGZIP(), []int{4}
}

type ListIdentityProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentityProviders []*IdentityProvider `protobuf:"bytes,1,rep,name=identity_providers,json=identityProviders,proto3" json:"identity_providers,omitempty"`
}

func (x *ListIdentityProvidersResponse) Reset() {
	*x = ListIdentityProvidersResponse{}
	mi := &file_api_v1_idp_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListIdentityProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIdentityProvidersResponse) ProtoMessage() {}

func (x *ListIdentityProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_idp_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIdentityProvidersResponse.ProtoReflect.Descriptor instead.
func (*ListIdentityProvidersResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_idp_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListIdentityProvidersResponse) GetIdentityProviders() []*IdentityProvider {
	if x != nil {
		return x.IdentityProviders
	}
	return nil
}

type GetIdentityProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the identityProvider to get.
	// Format: identityProviders/{id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetIdentityProviderRequest) Reset() {
	*x = GetIdentityProviderRequest{}
	mi := &file_api_v1_idp_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIdentityProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdentityProviderRequest) ProtoMessage() {}

func (x *GetIdentityProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_idp_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdentityProviderRequest.ProtoReflect.Descriptor instead.
func (*GetIdentityProviderRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_idp_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetIdentityProviderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateIdentityProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identityProvider to create.
	IdentityProvider *IdentityProvider `protobuf:"bytes,1,opt,name=identity_provider,json=identityProvider,proto3" json:"identity_provider,omitempty"`
}

func (x *CreateIdentityProviderRequest) Reset() {
	*x = CreateIdentityProviderRequest{}
	mi := &file_api_v1_idp_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateIdentityProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIdentityProviderRequest) ProtoMessage() {}

func (x *CreateIdentityProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_idp_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIdentityProviderRequest.ProtoReflect.Descriptor instead.
func (*CreateIdentityProviderRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_idp_service_proto_rawDescGZIP(), []int{7}
}

func (x *CreateIdentityProviderRequest) GetIdentityProvider() *IdentityProvider {
	if x != nil {
		return x.IdentityProvider
	}
	return nil
}

type UpdateIdentityProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identityProvider to update.
	IdentityProvider *IdentityProvider `protobuf:"bytes,1,opt,name=identity_provider,json=identityProvider,proto3" json:"identity_provider,omitempty"`
	// The update mask applies to the resource. Only the top level fields of
	// IdentityProvider are supported.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateIdentityProviderRequest) Reset() {
	*x = UpdateIdentityProviderRequest{}
	mi := &file_api_v1_idp_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateIdentityProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIdentityProviderRequest) ProtoMessage() {}

func (x *UpdateIdentityProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_idp_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIdentityProviderRequest.ProtoReflect.Descriptor instead.
func (*UpdateIdentityProviderRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_idp_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateIdentityProviderRequest) GetIdentityProvider() *IdentityProvider {
	if x != nil {
		return x.IdentityProvider
	}
	return nil
}

func (x *UpdateIdentityProviderRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteIdentityProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the identityProvider to delete.
	// Format: identityProviders/{id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteIdentityProviderRequest) Reset() {
	*x = DeleteIdentityProviderRequest{}
	mi := &file_api_v1_idp_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteIdentityProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIdentityProviderRequest) ProtoMessage() {}

func (x *DeleteIdentityProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_idp_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIdentityProviderRequest.ProtoReflect.Descriptor instead.
func (*DeleteIdentityProviderRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_idp_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteIdentityProviderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_v1_idp_service_proto protoreflect.FileDescriptor

var file_api_v1_idp_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x70, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x6d, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a,
	0x02, 0x0a, 0x10, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x28, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x32, 0x10, 0x01, 0x22, 0x65, 0x0a, 0x16, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x0d, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d,
	0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x32, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x32, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x67, 0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x85, 0x02, 0x0a, 0x0c,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x6d, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x22, 0x1e, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x6e, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x11, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x22, 0x30, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6c, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x11, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x10, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x22, 0xa9, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x11, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x10, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22,
	0x33, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0xce, 0x06, 0x0a, 0x17, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x93, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x2e, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x92, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x28,
	0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x31, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x9b, 0x01, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x11, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x19,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0xd6, 0x01, 0x0a, 0x16, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x22, 0x6f, 0xda, 0x41, 0x1d, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x49, 0x3a, 0x11, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x32, 0x34, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f,
	0x2a, 0x7d, 0x12, 0x90, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x2e,
	0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x31, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x24, 0x2a, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0xa7, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65,
	0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x49, 0x64, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x4d, 0x41, 0x58, 0xaa, 0x02, 0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x70,
	0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0e, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_idp_service_proto_rawDescOnce sync.Once
	file_api_v1_idp_service_proto_rawDescData = file_api_v1_idp_service_proto_rawDesc
)

func file_api_v1_idp_service_proto_rawDescGZIP() []byte {
	file_api_v1_idp_service_proto_rawDescOnce.Do(func() {
		file_api_v1_idp_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_idp_service_proto_rawDescData)
	})
	return file_api_v1_idp_service_proto_rawDescData
}

var file_api_v1_idp_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_idp_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_v1_idp_service_proto_goTypes = []any{
	(IdentityProvider_Type)(0),            // 0: memos.api.v1.IdentityProvider.Type
	(*IdentityProvider)(nil),              // 1: memos.api.v1.IdentityProvider
	(*IdentityProviderConfig)(nil),        // 2: memos.api.v1.IdentityProviderConfig
	(*FieldMapping)(nil),                  // 3: memos.api.v1.FieldMapping
	(*OAuth2Config)(nil),                  // 4: memos.api.v1.OAuth2Config
	(*ListIdentityProvidersRequest)(nil),  // 5: memos.api.v1.ListIdentityProvidersRequest
	(*ListIdentityProvidersResponse)(nil), // 6: memos.api.v1.ListIdentityProvidersResponse
	(*GetIdentityProviderRequest)(nil),    // 7: memos.api.v1.GetIdentityProviderRequest
	(*CreateIdentityProviderRequest)(nil), // 8: memos.api.v1.CreateIdentityProviderRequest
	(*UpdateIdentityProviderRequest)(nil), // 9: memos.api.v1.UpdateIdentityProviderRequest
	(*DeleteIdentityProviderRequest)(nil), // 10: memos.api.v1.DeleteIdentityProviderRequest
	(*fieldmaskpb.FieldMask)(nil),         // 11: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),                 // 12: google.protobuf.Empty
}
var file_api_v1_idp_service_proto_depIdxs = []int32{
	0,  // 0: memos.api.v1.IdentityProvider.type:type_name -> memos.api.v1.IdentityProvider.Type
	2,  // 1: memos.api.v1.IdentityProvider.config:type_name -> memos.api.v1.IdentityProviderConfig
	4,  // 2: memos.api.v1.IdentityProviderConfig.oauth2_config:type_name -> memos.api.v1.OAuth2Config
	3,  // 3: memos.api.v1.OAuth2Config.field_mapping:type_name -> memos.api.v1.FieldMapping
	1,  // 4: memos.api.v1.ListIdentityProvidersResponse.identity_providers:type_name -> memos.api.v1.IdentityProvider
	1,  // 5: memos.api.v1.CreateIdentityProviderRequest.identity_provider:type_name -> memos.api.v1.IdentityProvider
	1,  // 6: memos.api.v1.UpdateIdentityProviderRequest.identity_provider:type_name -> memos.api.v1.IdentityProvider
	11, // 7: memos.api.v1.UpdateIdentityProviderRequest.update_mask:type_name -> google.protobuf.FieldMask
	5,  // 8: memos.api.v1.IdentityProviderService.ListIdentityProviders:input_type -> memos.api.v1.ListIdentityProvidersRequest
	7,  // 9: memos.api.v1.IdentityProviderService.GetIdentityProvider:input_type -> memos.api.v1.GetIdentityProviderRequest
	8,  // 10: memos.api.v1.IdentityProviderService.CreateIdentityProvider:input_type -> memos.api.v1.CreateIdentityProviderRequest
	9,  // 11: memos.api.v1.IdentityProviderService.UpdateIdentityProvider:input_type -> memos.api.v1.UpdateIdentityProviderRequest
	10, // 12: memos.api.v1.IdentityProviderService.DeleteIdentityProvider:input_type -> memos.api.v1.DeleteIdentityProviderRequest
	6,  // 13: memos.api.v1.IdentityProviderService.ListIdentityProviders:output_type -> memos.api.v1.ListIdentityProvidersResponse
	1,  // 14: memos.api.v1.IdentityProviderService.GetIdentityProvider:output_type -> memos.api.v1.IdentityProvider
	1,  // 15: memos.api.v1.IdentityProviderService.CreateIdentityProvider:output_type -> memos.api.v1.IdentityProvider
	1,  // 16: memos.api.v1.IdentityProviderService.UpdateIdentityProvider:output_type -> memos.api.v1.IdentityProvider
	12, // 17: memos.api.v1.IdentityProviderService.DeleteIdentityProvider:output_type -> google.protobuf.Empty
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_v1_idp_service_proto_init() }
func file_api_v1_idp_service_proto_init() {
	if File_api_v1_idp_service_proto != nil {
		return
	}
	file_api_v1_idp_service_proto_msgTypes[1].OneofWrappers = []any{
		(*IdentityProviderConfig_Oauth2Config)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_idp_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_idp_service_proto_goTypes,
		DependencyIndexes: file_api_v1_idp_service_proto_depIdxs,
		EnumInfos:         file_api_v1_idp_service_proto_enumTypes,
		MessageInfos:      file_api_v1_idp_service_proto_msgTypes,
	}.Build()
	File_api_v1_idp_service_proto = out.File
	file_api_v1_idp_service_proto_rawDesc = nil
	file_api_v1_idp_service_proto_goTypes = nil
	file_api_v1_idp_service_proto_depIdxs = nil
}
