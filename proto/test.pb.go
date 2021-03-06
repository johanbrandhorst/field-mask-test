// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/test.proto

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	proto/test.proto

It has these top-level messages:
	UpdateUserRequest
	User
*/
package test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "google.golang.org/genproto/protobuf/field_mask"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateUserRequest struct {
	// The user resource which replaces the resource on the server.
	User *User `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	// The update mask applies to the resource. For the `FieldMask` definition,
	// see https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask *google_protobuf.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdateUserRequest) Reset()                    { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()               {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateUserRequest) GetUpdateMask() *google_protobuf.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type User struct {
	// Relative resource name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Display name.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// Description.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateUserRequest)(nil), "UpdateUserRequest")
	proto.RegisterType((*User)(nil), "User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestFieldMask service

type TestFieldMaskClient interface {
	// Update an existing user.
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
}

type testFieldMaskClient struct {
	cc *grpc.ClientConn
}

func NewTestFieldMaskClient(cc *grpc.ClientConn) TestFieldMaskClient {
	return &testFieldMaskClient{cc}
}

func (c *testFieldMaskClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/TestFieldMask/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestFieldMask service

type TestFieldMaskServer interface {
	// Update an existing user.
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
}

func RegisterTestFieldMaskServer(s *grpc.Server, srv TestFieldMaskServer) {
	s.RegisterService(&_TestFieldMask_serviceDesc, srv)
}

func _TestFieldMask_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestFieldMaskServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestFieldMask/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestFieldMaskServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestFieldMask_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TestFieldMask",
	HandlerType: (*TestFieldMaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateUser",
			Handler:    _TestFieldMask_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/test.proto",
}

func init() { proto.RegisterFile("proto/test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x4e, 0x84, 0x40,
	0x10, 0x86, 0xc3, 0x1d, 0x9a, 0x38, 0x68, 0xa2, 0x6b, 0x83, 0x68, 0xc1, 0x51, 0x5d, 0xb5, 0x44,
	0xec, 0xb4, 0xb7, 0x30, 0xd1, 0x82, 0x78, 0x95, 0x05, 0xd9, 0x93, 0xb9, 0x0b, 0x81, 0x63, 0x57,
	0x66, 0x31, 0x31, 0xc6, 0xc6, 0x57, 0xf0, 0xd1, 0x7c, 0x05, 0x1f, 0xc4, 0xec, 0x80, 0x5e, 0x71,
	0xdd, 0x30, 0xff, 0x17, 0xe6, 0xff, 0x16, 0x8e, 0x4d, 0xa7, 0xad, 0x4e, 0x2d, 0x92, 0x95, 0x3c,
	0x46, 0xf1, 0x5a, 0xeb, 0x75, 0x83, 0x29, 0x7f, 0x2d, 0xfb, 0x55, 0xba, 0xaa, 0xb0, 0x29, 0x8b,
	0x8d, 0xa2, 0x7a, 0x24, 0x2e, 0x46, 0x42, 0x99, 0x2a, 0x55, 0x6d, 0xab, 0xad, 0xb2, 0x95, 0x6e,
	0x69, 0x48, 0x93, 0x1a, 0x4e, 0x16, 0xa6, 0x54, 0x16, 0x17, 0x84, 0x5d, 0x8e, 0x2f, 0x3d, 0x92,
	0x15, 0x67, 0xe0, 0xf7, 0x84, 0x5d, 0x38, 0x89, 0xbd, 0x79, 0x90, 0xed, 0x49, 0xce, 0x78, 0x25,
	0x6e, 0x20, 0xe8, 0x99, 0xe7, 0x13, 0xe1, 0x94, 0x89, 0x48, 0x0e, 0x37, 0xe4, 0x5f, 0x0b, 0x79,
	0xeb, 0x5a, 0xdc, 0x2b, 0xaa, 0x73, 0x18, 0x70, 0x37, 0x27, 0x05, 0xf8, 0xee, 0x57, 0x42, 0x80,
	0xdf, 0xaa, 0x0d, 0x86, 0x5e, 0xec, 0xcd, 0x0f, 0x72, 0x9e, 0xc5, 0x0c, 0x0e, 0xcb, 0x8a, 0x4c,
	0xa3, 0xde, 0x0a, 0xce, 0xa6, 0x9c, 0x05, 0xe3, 0xee, 0xc1, 0x21, 0x31, 0x04, 0x25, 0xd2, 0x73,
	0x57, 0x19, 0x67, 0x10, 0xfa, 0x23, 0xb1, 0x5d, 0x65, 0x4f, 0x70, 0xf4, 0x88, 0x64, 0xff, 0xaf,
	0x8b, 0x3b, 0x80, 0xad, 0x9e, 0x10, 0x72, 0xc7, 0x35, 0x1a, 0xec, 0x92, 0xd9, 0xe7, 0xf7, 0xcf,
	0xd7, 0xe4, 0x3c, 0x3b, 0xe5, 0x77, 0x7a, 0xbd, 0x4c, 0xdf, 0x9d, 0xae, 0x74, 0x55, 0x3e, 0xae,
	0x59, 0x7d, 0xb9, 0xcf, 0x76, 0x57, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x43, 0x4a, 0xd7, 0xdc,
	0x85, 0x01, 0x00, 0x00,
}
