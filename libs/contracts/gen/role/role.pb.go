// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: role/role.proto

package rolev1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RoleType struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OrganizationId int64                  `protobuf:"varint,3,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Permissions    []string               `protobuf:"bytes,4,rep,name=permissions,proto3" json:"permissions,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *RoleType) Reset() {
	*x = RoleType{}
	mi := &file_role_role_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleType) ProtoMessage() {}

func (x *RoleType) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleType.ProtoReflect.Descriptor instead.
func (*RoleType) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{0}
}

func (x *RoleType) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoleType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleType) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *RoleType) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type AddRoleRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Name           string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OrganizationId int64                  `protobuf:"varint,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Permissions    []string               `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *AddRoleRequest) Reset() {
	*x = AddRoleRequest{}
	mi := &file_role_role_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleRequest) ProtoMessage() {}

func (x *AddRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleRequest.ProtoReflect.Descriptor instead.
func (*AddRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{1}
}

func (x *AddRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddRoleRequest) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *AddRoleRequest) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type AddRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRoleResponse) Reset() {
	*x = AddRoleResponse{}
	mi := &file_role_role_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleResponse) ProtoMessage() {}

func (x *AddRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleResponse.ProtoReflect.Descriptor instead.
func (*AddRoleResponse) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{2}
}

func (x *AddRoleResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoleRequest) Reset() {
	*x = GetRoleRequest{}
	mi := &file_role_role_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleRequest) ProtoMessage() {}

func (x *GetRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleRequest.ProtoReflect.Descriptor instead.
func (*GetRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{3}
}

func (x *GetRoleRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          *RoleType              `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoleResponse) Reset() {
	*x = GetRoleResponse{}
	mi := &file_role_role_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleResponse) ProtoMessage() {}

func (x *GetRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleResponse.ProtoReflect.Descriptor instead.
func (*GetRoleResponse) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{4}
}

func (x *GetRoleResponse) GetRole() *RoleType {
	if x != nil {
		return x.Role
	}
	return nil
}

type GetRolesRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	OrganizationId int64                  `protobuf:"varint,1,opt,name=organizationId,proto3" json:"organizationId,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetRolesRequest) Reset() {
	*x = GetRolesRequest{}
	mi := &file_role_role_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesRequest) ProtoMessage() {}

func (x *GetRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesRequest.ProtoReflect.Descriptor instead.
func (*GetRolesRequest) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{5}
}

func (x *GetRolesRequest) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

type GetRolesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Roles         []*RoleType            `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRolesResponse) Reset() {
	*x = GetRolesResponse{}
	mi := &file_role_role_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesResponse) ProtoMessage() {}

func (x *GetRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesResponse.ProtoReflect.Descriptor instead.
func (*GetRolesResponse) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{6}
}

func (x *GetRolesResponse) GetRoles() []*RoleType {
	if x != nil {
		return x.Roles
	}
	return nil
}

type UpdateRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          *RoleType              `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRoleRequest) Reset() {
	*x = UpdateRoleRequest{}
	mi := &file_role_role_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleRequest) ProtoMessage() {}

func (x *UpdateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRoleRequest) GetRole() *RoleType {
	if x != nil {
		return x.Role
	}
	return nil
}

type UpdateRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          *RoleType              `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRoleResponse) Reset() {
	*x = UpdateRoleResponse{}
	mi := &file_role_role_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleResponse) ProtoMessage() {}

func (x *UpdateRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleResponse.ProtoReflect.Descriptor instead.
func (*UpdateRoleResponse) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRoleResponse) GetRole() *RoleType {
	if x != nil {
		return x.Role
	}
	return nil
}

type DeleteRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRoleRequest) Reset() {
	*x = DeleteRoleRequest{}
	mi := &file_role_role_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRequest) ProtoMessage() {}

func (x *DeleteRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRoleRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRoleResponse) Reset() {
	*x = DeleteRoleResponse{}
	mi := &file_role_role_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleResponse) ProtoMessage() {}

func (x *DeleteRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleResponse.ProtoReflect.Descriptor instead.
func (*DeleteRoleResponse) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteRoleResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_role_role_proto protoreflect.FileDescriptor

var file_role_role_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x6f, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x21, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x39, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x22, 0x37, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x38, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x32, 0xbb, 0x03, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x4f, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x12, 0x4f, 0x0a, 0x07, 0x67, 0x65,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x57, 0x0a, 0x08, 0x67,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01,
	0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x2d,
	0x6d, 0x61, 0x6e, 0x79, 0x12, 0x5b, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x17, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a,
	0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x5b, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x17, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x17,
	0x5a, 0x15, 0x6c, 0x65, 0x61, 0x64, 0x49, 0x74, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x3b, 0x72, 0x6f, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_role_proto_rawDescOnce sync.Once
	file_role_role_proto_rawDescData = file_role_role_proto_rawDesc
)

func file_role_role_proto_rawDescGZIP() []byte {
	file_role_role_proto_rawDescOnce.Do(func() {
		file_role_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_role_proto_rawDescData)
	})
	return file_role_role_proto_rawDescData
}

var file_role_role_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_role_role_proto_goTypes = []any{
	(*RoleType)(nil),           // 0: role.RoleType
	(*AddRoleRequest)(nil),     // 1: role.AddRoleRequest
	(*AddRoleResponse)(nil),    // 2: role.AddRoleResponse
	(*GetRoleRequest)(nil),     // 3: role.GetRoleRequest
	(*GetRoleResponse)(nil),    // 4: role.GetRoleResponse
	(*GetRolesRequest)(nil),    // 5: role.GetRolesRequest
	(*GetRolesResponse)(nil),   // 6: role.GetRolesResponse
	(*UpdateRoleRequest)(nil),  // 7: role.UpdateRoleRequest
	(*UpdateRoleResponse)(nil), // 8: role.UpdateRoleResponse
	(*DeleteRoleRequest)(nil),  // 9: role.DeleteRoleRequest
	(*DeleteRoleResponse)(nil), // 10: role.DeleteRoleResponse
}
var file_role_role_proto_depIdxs = []int32{
	0,  // 0: role.GetRoleResponse.role:type_name -> role.RoleType
	0,  // 1: role.GetRolesResponse.roles:type_name -> role.RoleType
	0,  // 2: role.UpdateRoleRequest.role:type_name -> role.RoleType
	0,  // 3: role.UpdateRoleResponse.role:type_name -> role.RoleType
	1,  // 4: role.Role.addRole:input_type -> role.AddRoleRequest
	3,  // 5: role.Role.getRole:input_type -> role.GetRoleRequest
	5,  // 6: role.Role.getRoles:input_type -> role.GetRolesRequest
	7,  // 7: role.Role.updateRole:input_type -> role.UpdateRoleRequest
	9,  // 8: role.Role.deleteRole:input_type -> role.DeleteRoleRequest
	2,  // 9: role.Role.addRole:output_type -> role.AddRoleResponse
	4,  // 10: role.Role.getRole:output_type -> role.GetRoleResponse
	6,  // 11: role.Role.getRoles:output_type -> role.GetRolesResponse
	8,  // 12: role.Role.updateRole:output_type -> role.UpdateRoleResponse
	10, // 13: role.Role.deleteRole:output_type -> role.DeleteRoleResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_role_role_proto_init() }
func file_role_role_proto_init() {
	if File_role_role_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_role_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_role_proto_goTypes,
		DependencyIndexes: file_role_role_proto_depIdxs,
		MessageInfos:      file_role_role_proto_msgTypes,
	}.Build()
	File_role_role_proto = out.File
	file_role_role_proto_rawDesc = nil
	file_role_role_proto_goTypes = nil
	file_role_role_proto_depIdxs = nil
}
