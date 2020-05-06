// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/group_service.proto

package monitoring

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The `ListGroup` request.
type ListGroupsRequest struct {
	// Required. The project whose groups are to be listed. The format is
	// `"projects/{project_id_or_number}"`.
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	// An optional filter consisting of a single group name.  The filters limit
	// the groups returned based on their parent-child relationship with the
	// specified group. If no filter is specified, all groups are returned.
	//
	// Types that are valid to be assigned to Filter:
	//	*ListGroupsRequest_ChildrenOfGroup
	//	*ListGroupsRequest_AncestorsOfGroup
	//	*ListGroupsRequest_DescendantsOfGroup
	Filter isListGroupsRequest_Filter `protobuf_oneof:"filter"`
	// A positive number that is the maximum number of results to return.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return additional results from the previous method call.
	PageToken            string   `protobuf:"bytes,6,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGroupsRequest) Reset()         { *m = ListGroupsRequest{} }
func (m *ListGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*ListGroupsRequest) ProtoMessage()    {}
func (*ListGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ad21d0ed55c55a, []int{0}
}

func (m *ListGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupsRequest.Unmarshal(m, b)
}
func (m *ListGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupsRequest.Marshal(b, m, deterministic)
}
func (m *ListGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupsRequest.Merge(m, src)
}
func (m *ListGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_ListGroupsRequest.Size(m)
}
func (m *ListGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupsRequest proto.InternalMessageInfo

func (m *ListGroupsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isListGroupsRequest_Filter interface {
	isListGroupsRequest_Filter()
}

type ListGroupsRequest_ChildrenOfGroup struct {
	ChildrenOfGroup string `protobuf:"bytes,2,opt,name=children_of_group,json=childrenOfGroup,proto3,oneof"`
}

type ListGroupsRequest_AncestorsOfGroup struct {
	AncestorsOfGroup string `protobuf:"bytes,3,opt,name=ancestors_of_group,json=ancestorsOfGroup,proto3,oneof"`
}

type ListGroupsRequest_DescendantsOfGroup struct {
	DescendantsOfGroup string `protobuf:"bytes,4,opt,name=descendants_of_group,json=descendantsOfGroup,proto3,oneof"`
}

func (*ListGroupsRequest_ChildrenOfGroup) isListGroupsRequest_Filter() {}

func (*ListGroupsRequest_AncestorsOfGroup) isListGroupsRequest_Filter() {}

func (*ListGroupsRequest_DescendantsOfGroup) isListGroupsRequest_Filter() {}

func (m *ListGroupsRequest) GetFilter() isListGroupsRequest_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListGroupsRequest) GetChildrenOfGroup() string {
	if x, ok := m.GetFilter().(*ListGroupsRequest_ChildrenOfGroup); ok {
		return x.ChildrenOfGroup
	}
	return ""
}

func (m *ListGroupsRequest) GetAncestorsOfGroup() string {
	if x, ok := m.GetFilter().(*ListGroupsRequest_AncestorsOfGroup); ok {
		return x.AncestorsOfGroup
	}
	return ""
}

func (m *ListGroupsRequest) GetDescendantsOfGroup() string {
	if x, ok := m.GetFilter().(*ListGroupsRequest_DescendantsOfGroup); ok {
		return x.DescendantsOfGroup
	}
	return ""
}

func (m *ListGroupsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListGroupsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListGroupsRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListGroupsRequest_ChildrenOfGroup)(nil),
		(*ListGroupsRequest_AncestorsOfGroup)(nil),
		(*ListGroupsRequest_DescendantsOfGroup)(nil),
	}
}

// The `ListGroups` response.
type ListGroupsResponse struct {
	// The groups that match the specified filters.
	Group []*Group `protobuf:"bytes,1,rep,name=group,proto3" json:"group,omitempty"`
	// If there are more results than have been returned, then this field is set
	// to a non-empty value.  To see the additional results,
	// use that value as `pageToken` in the next call to this method.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGroupsResponse) Reset()         { *m = ListGroupsResponse{} }
func (m *ListGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*ListGroupsResponse) ProtoMessage()    {}
func (*ListGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ad21d0ed55c55a, []int{1}
}

func (m *ListGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupsResponse.Unmarshal(m, b)
}
func (m *ListGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupsResponse.Marshal(b, m, deterministic)
}
func (m *ListGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupsResponse.Merge(m, src)
}
func (m *ListGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_ListGroupsResponse.Size(m)
}
func (m *ListGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupsResponse proto.InternalMessageInfo

func (m *ListGroupsResponse) GetGroup() []*Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ListGroupsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The `GetGroup` request.
type GetGroupRequest struct {
	// Required. The group to retrieve. The format is
	// `"projects/{project_id_or_number}/groups/{group_id}"`.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupRequest) Reset()         { *m = GetGroupRequest{} }
func (m *GetGroupRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupRequest) ProtoMessage()    {}
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ad21d0ed55c55a, []int{2}
}

func (m *GetGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupRequest.Unmarshal(m, b)
}
func (m *GetGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupRequest.Marshal(b, m, deterministic)
}
func (m *GetGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupRequest.Merge(m, src)
}
func (m *GetGroupRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupRequest.Size(m)
}
func (m *GetGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupRequest proto.InternalMessageInfo

func (m *GetGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The `CreateGroup` request.
type CreateGroupRequest struct {
	// Required. The project in which to create the group. The format is
	// `"projects/{project_id_or_number}"`.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Required. A group definition. It is an error to define the `name` field because
	// the system assigns the name.
	Group *Group `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	// If true, validate this request but do not create the group.
	ValidateOnly         bool     `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGroupRequest) Reset()         { *m = CreateGroupRequest{} }
func (m *CreateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGroupRequest) ProtoMessage()    {}
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ad21d0ed55c55a, []int{3}
}

func (m *CreateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupRequest.Unmarshal(m, b)
}
func (m *CreateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupRequest.Marshal(b, m, deterministic)
}
func (m *CreateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupRequest.Merge(m, src)
}
func (m *CreateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGroupRequest.Size(m)
}
func (m *CreateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupRequest proto.InternalMessageInfo

func (m *CreateGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateGroupRequest) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *CreateGroupRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// The `UpdateGroup` request.
type UpdateGroupRequest struct {
	// Required. The new definition of the group.  All fields of the existing group,
	// excepting `name`, are replaced with the corresponding fields of this group.
	Group *Group `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	// If true, validate this request but do not update the existing group.
	ValidateOnly         bool     `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateGroupRequest) Reset()         { *m = UpdateGroupRequest{} }
func (m *UpdateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateGroupRequest) ProtoMessage()    {}
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ad21d0ed55c55a, []int{4}
}

func (m *UpdateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGroupRequest.Unmarshal(m, b)
}
func (m *UpdateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGroupRequest.Marshal(b, m, deterministic)
}
func (m *UpdateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGroupRequest.Merge(m, src)
}
func (m *UpdateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateGroupRequest.Size(m)
}
func (m *UpdateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGroupRequest proto.InternalMessageInfo

func (m *UpdateGroupRequest) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *UpdateGroupRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// The `DeleteGroup` request. The default behavior is to be able to delete a
// single group without any descendants.
type DeleteGroupRequest struct {
	// Required. The group to delete. The format is
	// `"projects/{project_id_or_number}/groups/{group_id}"`.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// If this field is true, then the request means to delete a group with all
	// its descendants. Otherwise, the request means to delete a group only when
	// it has no descendants. The default value is false.
	Recursive            bool     `protobuf:"varint,4,opt,name=recursive,proto3" json:"recursive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteGroupRequest) Reset()         { *m = DeleteGroupRequest{} }
func (m *DeleteGroupRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteGroupRequest) ProtoMessage()    {}
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ad21d0ed55c55a, []int{5}
}

func (m *DeleteGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGroupRequest.Unmarshal(m, b)
}
func (m *DeleteGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGroupRequest.Marshal(b, m, deterministic)
}
func (m *DeleteGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGroupRequest.Merge(m, src)
}
func (m *DeleteGroupRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteGroupRequest.Size(m)
}
func (m *DeleteGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGroupRequest proto.InternalMessageInfo

func (m *DeleteGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteGroupRequest) GetRecursive() bool {
	if m != nil {
		return m.Recursive
	}
	return false
}

// The `ListGroupMembers` request.
type ListGroupMembersRequest struct {
	// Required. The group whose members are listed. The format is
	// `"projects/{project_id_or_number}/groups/{group_id}"`.
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	// A positive number that is the maximum number of results to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return additional results from the previous method call.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// An optional [list filter](/monitoring/api/learn_more#filtering) describing
	// the members to be returned.  The filter may reference the type, labels, and
	// metadata of monitored resources that comprise the group.
	// For example, to return only resources representing Compute Engine VM
	// instances, use this filter:
	//
	//     resource.type = "gce_instance"
	Filter string `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
	// An optional time interval for which results should be returned. Only
	// members that were part of the group during the specified interval are
	// included in the response.  If no interval is provided then the group
	// membership over the last minute is returned.
	Interval             *TimeInterval `protobuf:"bytes,6,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListGroupMembersRequest) Reset()         { *m = ListGroupMembersRequest{} }
func (m *ListGroupMembersRequest) String() string { return proto.CompactTextString(m) }
func (*ListGroupMembersRequest) ProtoMessage()    {}
func (*ListGroupMembersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ad21d0ed55c55a, []int{6}
}

func (m *ListGroupMembersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupMembersRequest.Unmarshal(m, b)
}
func (m *ListGroupMembersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupMembersRequest.Marshal(b, m, deterministic)
}
func (m *ListGroupMembersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupMembersRequest.Merge(m, src)
}
func (m *ListGroupMembersRequest) XXX_Size() int {
	return xxx_messageInfo_ListGroupMembersRequest.Size(m)
}
func (m *ListGroupMembersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupMembersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupMembersRequest proto.InternalMessageInfo

func (m *ListGroupMembersRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListGroupMembersRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListGroupMembersRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListGroupMembersRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListGroupMembersRequest) GetInterval() *TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

// The `ListGroupMembers` response.
type ListGroupMembersResponse struct {
	// A set of monitored resources in the group.
	Members []*monitoredres.MonitoredResource `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	// If there are more results than have been returned, then this field is
	// set to a non-empty value.  To see the additional results, use that value as
	// `pageToken` in the next call to this method.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The total number of elements matching this request.
	TotalSize            int32    `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGroupMembersResponse) Reset()         { *m = ListGroupMembersResponse{} }
func (m *ListGroupMembersResponse) String() string { return proto.CompactTextString(m) }
func (*ListGroupMembersResponse) ProtoMessage()    {}
func (*ListGroupMembersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ad21d0ed55c55a, []int{7}
}

func (m *ListGroupMembersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupMembersResponse.Unmarshal(m, b)
}
func (m *ListGroupMembersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupMembersResponse.Marshal(b, m, deterministic)
}
func (m *ListGroupMembersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupMembersResponse.Merge(m, src)
}
func (m *ListGroupMembersResponse) XXX_Size() int {
	return xxx_messageInfo_ListGroupMembersResponse.Size(m)
}
func (m *ListGroupMembersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupMembersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupMembersResponse proto.InternalMessageInfo

func (m *ListGroupMembersResponse) GetMembers() []*monitoredres.MonitoredResource {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *ListGroupMembersResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListGroupMembersResponse) GetTotalSize() int32 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func init() {
	proto.RegisterType((*ListGroupsRequest)(nil), "google.monitoring.v3.ListGroupsRequest")
	proto.RegisterType((*ListGroupsResponse)(nil), "google.monitoring.v3.ListGroupsResponse")
	proto.RegisterType((*GetGroupRequest)(nil), "google.monitoring.v3.GetGroupRequest")
	proto.RegisterType((*CreateGroupRequest)(nil), "google.monitoring.v3.CreateGroupRequest")
	proto.RegisterType((*UpdateGroupRequest)(nil), "google.monitoring.v3.UpdateGroupRequest")
	proto.RegisterType((*DeleteGroupRequest)(nil), "google.monitoring.v3.DeleteGroupRequest")
	proto.RegisterType((*ListGroupMembersRequest)(nil), "google.monitoring.v3.ListGroupMembersRequest")
	proto.RegisterType((*ListGroupMembersResponse)(nil), "google.monitoring.v3.ListGroupMembersResponse")
}

func init() {
	proto.RegisterFile("google/monitoring/v3/group_service.proto", fileDescriptor_21ad21d0ed55c55a)
}

var fileDescriptor_21ad21d0ed55c55a = []byte{
	// 999 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0x67, 0xed, 0x24, 0xb5, 0xc7, 0xad, 0xd2, 0x8c, 0xaa, 0xd6, 0xdd, 0xa4, 0xad, 0xb3, 0xfd,
	0x67, 0x85, 0x66, 0x57, 0xc4, 0xe2, 0x8f, 0x1a, 0x51, 0xc9, 0x2e, 0x28, 0x20, 0x51, 0x1a, 0x6d,
	0x43, 0x85, 0x50, 0x24, 0x6b, 0xb2, 0x7e, 0x76, 0x06, 0x76, 0x67, 0xb6, 0xb3, 0x63, 0x87, 0x14,
	0x85, 0x03, 0x37, 0x24, 0x84, 0x84, 0xb8, 0xa0, 0x7e, 0x02, 0xca, 0x81, 0x0f, 0xd2, 0x23, 0xdc,
	0x72, 0xea, 0x81, 0x13, 0x1f, 0xa1, 0x27, 0xb4, 0xb3, 0x3b, 0xf1, 0xc6, 0x7f, 0x93, 0x48, 0x1c,
	0xf7, 0xbd, 0xdf, 0xbc, 0xdf, 0x7b, 0x6f, 0xde, 0xfb, 0xcd, 0xa2, 0x6a, 0x87, 0xf3, 0x8e, 0x0f,
	0x4e, 0xc0, 0x19, 0x95, 0x5c, 0x50, 0xd6, 0x71, 0x7a, 0x35, 0xa7, 0x23, 0x78, 0x37, 0x6c, 0x46,
	0x20, 0x7a, 0xd4, 0x03, 0x3b, 0x14, 0x5c, 0x72, 0x7c, 0x29, 0x41, 0xda, 0x7d, 0xa4, 0xdd, 0xab,
	0x99, 0x4b, 0xe9, 0x79, 0x12, 0x52, 0x87, 0x30, 0xc6, 0x25, 0x91, 0x94, 0xb3, 0x28, 0x39, 0x63,
	0x5e, 0xc9, 0x78, 0x3d, 0x9f, 0x02, 0x93, 0xa9, 0xe3, 0x46, 0xc6, 0xd1, 0xa6, 0xe0, 0xb7, 0x9a,
	0x3b, 0xb0, 0x4b, 0x7a, 0x94, 0x8b, 0x14, 0x70, 0x33, 0x03, 0x48, 0x19, 0xa1, 0xd5, 0x14, 0x10,
	0xf1, 0xae, 0xd0, 0x29, 0x99, 0x57, 0x33, 0xa0, 0x01, 0xd7, 0xf2, 0xc8, 0xba, 0x3c, 0x1e, 0x04,
	0x9c, 0xa5, 0x90, 0xca, 0xf8, 0xd2, 0x53, 0xc4, 0x62, 0x8a, 0x50, 0x5f, 0x3b, 0xdd, 0xb6, 0x03,
	0x41, 0x28, 0xf7, 0x13, 0xa7, 0xf5, 0x22, 0x8f, 0x16, 0x3e, 0xa3, 0x91, 0xdc, 0x88, 0x0f, 0x44,
	0x2e, 0x3c, 0xeb, 0x42, 0x24, 0xf1, 0x3a, 0x9a, 0x61, 0x24, 0x80, 0xf2, 0xb9, 0x8a, 0x51, 0x2d,
	0x36, 0xee, 0xbe, 0xae, 0xe7, 0xde, 0xd4, 0x97, 0xf1, 0x8d, 0x4c, 0xd3, 0x92, 0x98, 0x24, 0xa4,
	0x91, 0xed, 0xf1, 0xc0, 0x51, 0xc7, 0x5d, 0x75, 0x08, 0xbb, 0x68, 0xc1, 0xdb, 0xa5, 0x7e, 0x4b,
	0x00, 0x6b, 0xf2, 0x76, 0x53, 0xa5, 0x52, 0xce, 0xa9, 0x48, 0xb7, 0xde, 0xd4, 0x97, 0xd1, 0xb4,
	0x30, 0x9f, 0xbc, 0xe5, 0xce, 0xeb, 0x00, 0x8f, 0xdb, 0xca, 0x84, 0xb7, 0x10, 0x26, 0xcc, 0x83,
	0x48, 0x72, 0x11, 0xf5, 0x83, 0xe6, 0x4f, 0x15, 0xf4, 0xe2, 0x51, 0x04, 0x1d, 0xf5, 0x4b, 0x74,
	0xa9, 0x05, 0x91, 0x07, 0xac, 0x45, 0x98, 0xcc, 0xc4, 0x9d, 0x39, 0x55, 0x5c, 0x9c, 0x89, 0xa1,
	0x23, 0x2f, 0xa2, 0x62, 0x48, 0x3a, 0xd0, 0x8c, 0xe8, 0x73, 0x28, 0xcf, 0x56, 0x8c, 0xea, 0xac,
	0x5b, 0x88, 0x0d, 0x4f, 0xe8, 0x73, 0xc0, 0xd7, 0x10, 0x52, 0x4e, 0xc9, 0xbf, 0x01, 0x56, 0x9e,
	0x8b, 0xc9, 0x5c, 0x05, 0xdf, 0x8a, 0x0d, 0x8d, 0x02, 0x9a, 0x6b, 0x53, 0x5f, 0x82, 0xb0, 0x38,
	0xc2, 0xd9, 0xbb, 0x89, 0x42, 0xce, 0x22, 0xc0, 0xef, 0xa0, 0xd9, 0x24, 0x4d, 0xa3, 0x92, 0xaf,
	0x96, 0xd6, 0x16, 0xed, 0x51, 0x23, 0x6d, 0x27, 0x37, 0x92, 0x20, 0xf1, 0x1d, 0x34, 0xcf, 0xe0,
	0x5b, 0xd9, 0xcc, 0xd0, 0xaa, 0x0b, 0x71, 0x2f, 0xc4, 0xe6, 0x4d, 0x4d, 0x6d, 0x7d, 0x8e, 0xe6,
	0x37, 0x20, 0xe1, 0x1b, 0x1c, 0x85, 0x7c, 0x76, 0x14, 0xd0, 0xc9, 0x46, 0xc1, 0xfa, 0xd3, 0x40,
	0xf8, 0xa1, 0x00, 0x22, 0x61, 0x64, 0xcc, 0x99, 0xb3, 0x8c, 0xd7, 0x7b, 0xba, 0xfc, 0xb8, 0x82,
	0xc9, 0xe5, 0x37, 0xf2, 0xaf, 0xeb, 0x39, 0xdd, 0x83, 0x9b, 0xe8, 0x42, 0x8f, 0xf8, 0xb4, 0x45,
	0x24, 0x34, 0x39, 0xf3, 0xf7, 0x55, 0x45, 0x05, 0xf7, 0xbc, 0x36, 0x3e, 0x66, 0xfe, 0xbe, 0xf5,
	0x0c, 0xe1, 0x2f, 0xc2, 0xd6, 0x60, 0xbe, 0xff, 0x2b, 0x25, 0x47, 0xf8, 0x23, 0xf0, 0x61, 0x4c,
	0x8b, 0xce, 0xd2, 0x76, 0xbc, 0x84, 0x8a, 0x02, 0xbc, 0xae, 0x88, 0x68, 0x2f, 0x69, 0x72, 0xc1,
	0xed, 0x1b, 0xac, 0x7f, 0x0d, 0x74, 0xe5, 0x68, 0xac, 0x1e, 0x41, 0xb0, 0x03, 0x62, 0xf2, 0xe2,
	0x9f, 0x94, 0xf6, 0xd8, 0xd0, 0xe7, 0x27, 0x0e, 0xfd, 0xcc, 0xc0, 0xd0, 0xe3, 0xcb, 0x7a, 0xe8,
	0xd5, 0xb6, 0x14, 0xdd, 0xf4, 0x0b, 0x3f, 0x40, 0x05, 0xca, 0x24, 0x88, 0x1e, 0xf1, 0xd5, 0xa6,
	0x94, 0xd6, 0xac, 0xd1, 0xdd, 0xdf, 0xa2, 0x01, 0x7c, 0x9a, 0x22, 0xdd, 0xa3, 0x33, 0xd6, 0x0b,
	0x03, 0x95, 0x87, 0x8b, 0x4d, 0x37, 0xe9, 0x7d, 0x74, 0x2e, 0x48, 0x4c, 0xe9, 0x2e, 0x5d, 0xd3,
	0xb1, 0x49, 0x48, 0xed, 0x47, 0x5a, 0xb0, 0xdd, 0x54, 0x94, 0x5d, 0x8d, 0x3e, 0xe9, 0x3e, 0xc5,
	0x45, 0x4b, 0x2e, 0x89, 0x9f, 0x6d, 0x49, 0x51, 0x59, 0xe2, 0x9e, 0xac, 0xfd, 0x56, 0x40, 0xe7,
	0x55, 0x62, 0x4f, 0x92, 0x37, 0x0a, 0xff, 0x64, 0x20, 0xd4, 0xdf, 0x78, 0x7c, 0x77, 0x74, 0xa9,
	0x43, 0x7a, 0x6d, 0x56, 0xa7, 0x03, 0x93, 0x92, 0xad, 0xb7, 0x0f, 0xeb, 0xea, 0xb2, 0x7e, 0xf8,
	0xfb, 0x9f, 0x5f, 0x73, 0xd7, 0xf1, 0x52, 0xfc, 0x58, 0x7c, 0x17, 0x1b, 0x3e, 0x0c, 0x05, 0xff,
	0x1a, 0x3c, 0x19, 0x39, 0x2b, 0x07, 0xc9, 0xf3, 0x11, 0xe1, 0x03, 0x54, 0xd0, 0x72, 0x80, 0x6f,
	0x8f, 0x19, 0xfa, 0xe3, 0x72, 0x61, 0x4e, 0xda, 0x0d, 0x6b, 0x35, 0x4b, 0x5e, 0xc1, 0xd7, 0x47,
	0x91, 0xa7, 0xdc, 0xce, 0xca, 0x01, 0xfe, 0xd9, 0x40, 0xa5, 0x8c, 0x7a, 0xe0, 0x31, 0x55, 0x0e,
	0x0b, 0xcc, 0xe4, 0x2c, 0x3e, 0x38, 0xac, 0xa3, 0x98, 0xf4, 0x9e, 0x62, 0x52, 0xb9, 0xdc, 0xb6,
	0x26, 0x36, 0xe2, 0x7e, 0xba, 0xcf, 0xbf, 0x18, 0xa8, 0x94, 0x91, 0x87, 0x71, 0x09, 0x0d, 0x2b,
	0xc8, 0xe4, 0x84, 0xd6, 0x0f, 0xeb, 0xb3, 0xfd, 0x5c, 0x56, 0xcd, 0x5b, 0x2a, 0x97, 0xe4, 0x09,
	0x1f, 0xdb, 0x1d, 0x9d, 0xd3, 0xf7, 0xa8, 0x94, 0x91, 0x8f, 0x71, 0x29, 0x0d, 0x2b, 0x8c, 0x79,
	0x59, 0x23, 0xf5, 0x7f, 0x81, 0xfd, 0x71, 0xfc, 0x5f, 0x30, 0x70, 0x49, 0x2b, 0xd3, 0x2e, 0xe9,
	0x77, 0x03, 0x5d, 0x1c, 0x5c, 0x30, 0xbc, 0x3a, 0x65, 0x1e, 0x8f, 0xab, 0x8e, 0x69, 0x9f, 0x14,
	0x9e, 0x0e, 0xf1, 0xbb, 0xd9, 0x14, 0xab, 0xf8, 0xce, 0xe4, 0x14, 0x9d, 0x74, 0x6b, 0xcd, 0x97,
	0xc6, 0xab, 0xfa, 0xd5, 0xb1, 0x52, 0xf6, 0x57, 0xfd, 0x47, 0x63, 0x57, 0xca, 0x30, 0xba, 0xef,
	0x38, 0x7b, 0x7b, 0x7b, 0x83, 0x42, 0x47, 0xba, 0x72, 0xd7, 0xf1, 0x7c, 0xde, 0x6d, 0xad, 0x86,
	0x3e, 0x91, 0x6d, 0x2e, 0x82, 0x7b, 0xd3, 0xe0, 0x7d, 0xae, 0x53, 0x40, 0x6d, 0x01, 0xa4, 0xd5,
	0x78, 0x69, 0xa0, 0xb2, 0xc7, 0x83, 0x91, 0x8d, 0x69, 0x2c, 0x64, 0x45, 0x63, 0x33, 0xbe, 0xbe,
	0x4d, 0xe3, 0xab, 0x07, 0x29, 0xb4, 0xc3, 0x7d, 0xc2, 0x3a, 0x36, 0x17, 0x1d, 0xa7, 0x03, 0x4c,
	0x5d, 0xae, 0xd3, 0x67, 0x3c, 0xfe, 0x9f, 0xb8, 0xde, 0xff, 0xfa, 0x23, 0x67, 0x6e, 0x24, 0x01,
	0x1e, 0xc6, 0x45, 0x6a, 0xf5, 0x8b, 0x19, 0x9f, 0xd6, 0x5e, 0x69, 0xe7, 0xb6, 0x72, 0x6e, 0xf7,
	0x9d, 0xdb, 0x4f, 0x6b, 0x3b, 0x73, 0x8a, 0xa4, 0xf6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c,
	0x8b, 0xf0, 0x28, 0x86, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupServiceClient interface {
	// Lists the existing groups.
	ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error)
	// Gets a single group.
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// Creates a new group.
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// Updates an existing group.
	// You can change any group attributes except `name`.
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// Deletes an existing group.
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists the monitored resources that are members of a group.
	ListGroupMembers(ctx context.Context, in *ListGroupMembersRequest, opts ...grpc.CallOption) (*ListGroupMembersResponse, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error) {
	out := new(ListGroupsResponse)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.GroupService/ListGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.GroupService/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.GroupService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.GroupService/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.GroupService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) ListGroupMembers(ctx context.Context, in *ListGroupMembersRequest, opts ...grpc.CallOption) (*ListGroupMembersResponse, error) {
	out := new(ListGroupMembersResponse)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.GroupService/ListGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
type GroupServiceServer interface {
	// Lists the existing groups.
	ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error)
	// Gets a single group.
	GetGroup(context.Context, *GetGroupRequest) (*Group, error)
	// Creates a new group.
	CreateGroup(context.Context, *CreateGroupRequest) (*Group, error)
	// Updates an existing group.
	// You can change any group attributes except `name`.
	UpdateGroup(context.Context, *UpdateGroupRequest) (*Group, error)
	// Deletes an existing group.
	DeleteGroup(context.Context, *DeleteGroupRequest) (*empty.Empty, error)
	// Lists the monitored resources that are members of a group.
	ListGroupMembers(context.Context, *ListGroupMembersRequest) (*ListGroupMembersResponse, error)
}

// UnimplementedGroupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGroupServiceServer struct {
}

func (*UnimplementedGroupServiceServer) ListGroups(ctx context.Context, req *ListGroupsRequest) (*ListGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (*UnimplementedGroupServiceServer) GetGroup(ctx context.Context, req *GetGroupRequest) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (*UnimplementedGroupServiceServer) CreateGroup(ctx context.Context, req *CreateGroupRequest) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (*UnimplementedGroupServiceServer) UpdateGroup(ctx context.Context, req *UpdateGroupRequest) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (*UnimplementedGroupServiceServer) DeleteGroup(ctx context.Context, req *DeleteGroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (*UnimplementedGroupServiceServer) ListGroupMembers(ctx context.Context, req *ListGroupMembersRequest) (*ListGroupMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroupMembers not implemented")
}

func RegisterGroupServiceServer(s *grpc.Server, srv GroupServiceServer) {
	s.RegisterService(&_GroupService_serviceDesc, srv)
}

func _GroupService_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/ListGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListGroups(ctx, req.(*ListGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_ListGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/ListGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListGroupMembers(ctx, req.(*ListGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.monitoring.v3.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGroups",
			Handler:    _GroupService_ListGroups_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _GroupService_GetGroup_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _GroupService_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _GroupService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _GroupService_DeleteGroup_Handler,
		},
		{
			MethodName: "ListGroupMembers",
			Handler:    _GroupService_ListGroupMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/monitoring/v3/group_service.proto",
}