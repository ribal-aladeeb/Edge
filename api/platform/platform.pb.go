// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform.proto

package platform

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Class struct {
	Id                   int32      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Topic                string     `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Instructor           *Teacher   `protobuf:"bytes,3,opt,name=instructor" json:"instructor,omitempty"`
	Students             []*Student `protobuf:"bytes,4,rep,name=students" json:"students,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Class) Reset()         { *m = Class{} }
func (m *Class) String() string { return proto.CompactTextString(m) }
func (*Class) ProtoMessage()    {}
func (*Class) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_8c1e7ce88f2b2cd8, []int{0}
}
func (m *Class) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Class.Unmarshal(m, b)
}
func (m *Class) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Class.Marshal(b, m, deterministic)
}
func (dst *Class) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Class.Merge(dst, src)
}
func (m *Class) XXX_Size() int {
	return xxx_messageInfo_Class.Size(m)
}
func (m *Class) XXX_DiscardUnknown() {
	xxx_messageInfo_Class.DiscardUnknown(m)
}

var xxx_messageInfo_Class proto.InternalMessageInfo

func (m *Class) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Class) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Class) GetInstructor() *Teacher {
	if m != nil {
		return m.Instructor
	}
	return nil
}

func (m *Class) GetStudents() []*Student {
	if m != nil {
		return m.Students
	}
	return nil
}

type Parent struct {
	Id                   int32      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FirstName            string     `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName             string     `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Email                string     `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Children             []*Student `protobuf:"bytes,5,rep,name=children" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Parent) Reset()         { *m = Parent{} }
func (m *Parent) String() string { return proto.CompactTextString(m) }
func (*Parent) ProtoMessage()    {}
func (*Parent) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_8c1e7ce88f2b2cd8, []int{1}
}
func (m *Parent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parent.Unmarshal(m, b)
}
func (m *Parent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parent.Marshal(b, m, deterministic)
}
func (dst *Parent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parent.Merge(dst, src)
}
func (m *Parent) XXX_Size() int {
	return xxx_messageInfo_Parent.Size(m)
}
func (m *Parent) XXX_DiscardUnknown() {
	xxx_messageInfo_Parent.DiscardUnknown(m)
}

var xxx_messageInfo_Parent proto.InternalMessageInfo

func (m *Parent) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Parent) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Parent) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Parent) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Parent) GetChildren() []*Student {
	if m != nil {
		return m.Children
	}
	return nil
}

type Student struct {
	Id                   int32      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FirstName            string     `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName             string     `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Mother               *Parent    `protobuf:"bytes,5,opt,name=mother" json:"mother,omitempty"`
	Father               *Parent    `protobuf:"bytes,6,opt,name=father" json:"father,omitempty"`
	Teachers             []*Teacher `protobuf:"bytes,7,rep,name=teachers" json:"teachers,omitempty"`
	Subjects             []*Class   `protobuf:"bytes,8,rep,name=subjects" json:"subjects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Student) Reset()         { *m = Student{} }
func (m *Student) String() string { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()    {}
func (*Student) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_8c1e7ce88f2b2cd8, []int{2}
}
func (m *Student) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Student.Unmarshal(m, b)
}
func (m *Student) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Student.Marshal(b, m, deterministic)
}
func (dst *Student) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Student.Merge(dst, src)
}
func (m *Student) XXX_Size() int {
	return xxx_messageInfo_Student.Size(m)
}
func (m *Student) XXX_DiscardUnknown() {
	xxx_messageInfo_Student.DiscardUnknown(m)
}

var xxx_messageInfo_Student proto.InternalMessageInfo

func (m *Student) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Student) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Student) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Student) GetMother() *Parent {
	if m != nil {
		return m.Mother
	}
	return nil
}

func (m *Student) GetFather() *Parent {
	if m != nil {
		return m.Father
	}
	return nil
}

func (m *Student) GetTeachers() []*Teacher {
	if m != nil {
		return m.Teachers
	}
	return nil
}

func (m *Student) GetSubjects() []*Class {
	if m != nil {
		return m.Subjects
	}
	return nil
}

type Status struct {
	Success              bool     `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_8c1e7ce88f2b2cd8, []int{3}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Teacher struct {
	Id                   int32      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FirstName            string     `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName             string     `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Email                string     `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Students             []*Student `protobuf:"bytes,5,rep,name=students" json:"students,omitempty"`
	Subjects             []*Class   `protobuf:"bytes,6,rep,name=subjects" json:"subjects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Teacher) Reset()         { *m = Teacher{} }
func (m *Teacher) String() string { return proto.CompactTextString(m) }
func (*Teacher) ProtoMessage()    {}
func (*Teacher) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_8c1e7ce88f2b2cd8, []int{4}
}
func (m *Teacher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Teacher.Unmarshal(m, b)
}
func (m *Teacher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Teacher.Marshal(b, m, deterministic)
}
func (dst *Teacher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Teacher.Merge(dst, src)
}
func (m *Teacher) XXX_Size() int {
	return xxx_messageInfo_Teacher.Size(m)
}
func (m *Teacher) XXX_DiscardUnknown() {
	xxx_messageInfo_Teacher.DiscardUnknown(m)
}

var xxx_messageInfo_Teacher proto.InternalMessageInfo

func (m *Teacher) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Teacher) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Teacher) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Teacher) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Teacher) GetStudents() []*Student {
	if m != nil {
		return m.Students
	}
	return nil
}

func (m *Teacher) GetSubjects() []*Class {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func init() {
	proto.RegisterType((*Class)(nil), "platform.Class")
	proto.RegisterType((*Parent)(nil), "platform.Parent")
	proto.RegisterType((*Student)(nil), "platform.Student")
	proto.RegisterType((*Status)(nil), "platform.Status")
	proto.RegisterType((*Teacher)(nil), "platform.Teacher")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Platform service

type PlatformClient interface {
	AddClass(ctx context.Context, in *Class, opts ...grpc.CallOption) (*Status, error)
	AddStudents(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Status, error)
	RegisterParent(ctx context.Context, in *Parent, opts ...grpc.CallOption) (*Status, error)
	RegisterTeacher(ctx context.Context, in *Teacher, opts ...grpc.CallOption) (*Status, error)
}

type platformClient struct {
	cc *grpc.ClientConn
}

func NewPlatformClient(cc *grpc.ClientConn) PlatformClient {
	return &platformClient{cc}
}

func (c *platformClient) AddClass(ctx context.Context, in *Class, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/platform.Platform/AddClass", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) AddStudents(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/platform.Platform/AddStudents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) RegisterParent(ctx context.Context, in *Parent, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/platform.Platform/registerParent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) RegisterTeacher(ctx context.Context, in *Teacher, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/platform.Platform/registerTeacher", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Platform service

type PlatformServer interface {
	AddClass(context.Context, *Class) (*Status, error)
	AddStudents(context.Context, *Student) (*Status, error)
	RegisterParent(context.Context, *Parent) (*Status, error)
	RegisterTeacher(context.Context, *Teacher) (*Status, error)
}

func RegisterPlatformServer(s *grpc.Server, srv PlatformServer) {
	s.RegisterService(&_Platform_serviceDesc, srv)
}

func _Platform_AddClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Class)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).AddClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.Platform/AddClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).AddClass(ctx, req.(*Class))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_AddStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Student)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).AddStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.Platform/AddStudents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).AddStudents(ctx, req.(*Student))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_RegisterParent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Parent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).RegisterParent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.Platform/RegisterParent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).RegisterParent(ctx, req.(*Parent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_RegisterTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Teacher)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).RegisterTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.Platform/RegisterTeacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).RegisterTeacher(ctx, req.(*Teacher))
	}
	return interceptor(ctx, in, info, handler)
}

var _Platform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "platform.Platform",
	HandlerType: (*PlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddClass",
			Handler:    _Platform_AddClass_Handler,
		},
		{
			MethodName: "AddStudents",
			Handler:    _Platform_AddStudents_Handler,
		},
		{
			MethodName: "registerParent",
			Handler:    _Platform_RegisterParent_Handler,
		},
		{
			MethodName: "registerTeacher",
			Handler:    _Platform_RegisterTeacher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform.proto",
}

func init() { proto.RegisterFile("platform.proto", fileDescriptor_platform_8c1e7ce88f2b2cd8) }

var fileDescriptor_platform_8c1e7ce88f2b2cd8 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x95, 0xd3, 0xc9, 0xa3, 0x77, 0xa4, 0xce, 0x60, 0xcd, 0xc2, 0x02, 0x21, 0x45, 0x59, 0x45,
	0x42, 0x8c, 0x44, 0x41, 0xac, 0xd8, 0x54, 0xec, 0x51, 0x95, 0xb2, 0x47, 0x6e, 0xec, 0xb6, 0x46,
	0x79, 0x54, 0xbe, 0xce, 0x27, 0xb0, 0x66, 0xcd, 0x17, 0xf1, 0x11, 0xfc, 0x0c, 0x8a, 0x13, 0xa7,
	0x8f, 0xb4, 0xea, 0x06, 0xb1, 0xbc, 0x3e, 0xe7, 0xc4, 0xe7, 0x9e, 0x7b, 0x1d, 0x98, 0xed, 0x0b,
	0x6e, 0x36, 0xb5, 0x2e, 0x9f, 0xf7, 0xba, 0x36, 0x35, 0x8d, 0x5c, 0x9d, 0xfc, 0x24, 0xe0, 0x7f,
	0x2e, 0x38, 0x22, 0x9d, 0x81, 0xa7, 0x04, 0x23, 0x31, 0x49, 0xfd, 0xcc, 0x53, 0x82, 0x3e, 0x81,
	0x6f, 0xea, 0xbd, 0xca, 0x99, 0x17, 0x93, 0x74, 0x9a, 0x75, 0x05, 0x7d, 0x07, 0xa0, 0x2a, 0x34,
	0xba, 0xc9, 0x4d, 0xad, 0xd9, 0x24, 0x26, 0xe9, 0xfd, 0xfc, 0xc5, 0xf3, 0xf0, 0xf9, 0xaf, 0x92,
	0xe7, 0x3b, 0xa9, 0xb3, 0x23, 0x12, 0x7d, 0x0b, 0x11, 0x9a, 0x46, 0xc8, 0xca, 0x20, 0xbb, 0x8b,
	0x27, 0xa7, 0x82, 0x55, 0x87, 0x64, 0x03, 0x25, 0xf9, 0x45, 0x20, 0x58, 0x72, 0x2d, 0x2b, 0x33,
	0xb2, 0xf4, 0x1a, 0x60, 0xa3, 0x34, 0x9a, 0x6f, 0x15, 0x2f, 0x65, 0xef, 0x6b, 0x6a, 0x4f, 0xbe,
	0xf0, 0x52, 0xd2, 0x57, 0x30, 0x2d, 0xb8, 0x43, 0x27, 0x16, 0x8d, 0xda, 0x03, 0x0b, 0x3e, 0x81,
	0x2f, 0x4b, 0xae, 0x0a, 0x76, 0xd7, 0xb5, 0x63, 0x8b, 0xd6, 0x5b, 0xbe, 0x53, 0x85, 0xd0, 0xb2,
	0x62, 0xfe, 0x55, 0x6f, 0x8e, 0x92, 0xfc, 0xf0, 0x20, 0xec, 0x4f, 0xff, 0xa9, 0xb9, 0x14, 0x82,
	0xb2, 0x36, 0x3b, 0xa9, 0x99, 0x6f, 0x13, 0x7d, 0x3c, 0x98, 0xe8, 0xa2, 0xc8, 0x7a, 0xbc, 0x65,
	0x6e, 0xb8, 0x65, 0x06, 0xd7, 0x98, 0x1d, 0xde, 0xb6, 0x66, 0xba, 0x69, 0x20, 0x0b, 0xcf, 0x5b,
	0x73, 0x73, 0x1a, 0x28, 0xf4, 0x0d, 0x44, 0xd8, 0xac, 0xbf, 0xcb, 0xdc, 0x20, 0x8b, 0x2c, 0xfd,
	0xe1, 0x40, 0xb7, 0x1b, 0x92, 0x0d, 0x84, 0xe4, 0x13, 0x04, 0x2b, 0xc3, 0x4d, 0x83, 0x94, 0x41,
	0x88, 0x4d, 0x9e, 0x4b, 0x44, 0x1b, 0x45, 0x94, 0xb9, 0xb2, 0x45, 0x4a, 0x89, 0xc8, 0xb7, 0x2e,
	0x0c, 0x57, 0x26, 0xbf, 0x09, 0x84, 0xbd, 0x81, 0xff, 0x33, 0xe2, 0x61, 0xfd, 0xfc, 0x9b, 0xeb,
	0x77, 0x92, 0x43, 0x70, 0x23, 0x87, 0xf9, 0x1f, 0x02, 0xd1, 0xb2, 0x07, 0xdb, 0x8b, 0x16, 0x42,
	0x74, 0x8f, 0xe9, 0x5c, 0xf3, 0xf2, 0xf1, 0xf8, 0x4e, 0x9b, 0xdc, 0x1c, 0xee, 0x17, 0x42, 0xac,
	0xdc, 0xbd, 0x63, 0x53, 0x17, 0x34, 0x1f, 0x60, 0xa6, 0xe5, 0x56, 0xa1, 0x91, 0xba, 0x7f, 0x22,
	0xa3, 0xf9, 0x5f, 0x50, 0x7d, 0x84, 0x07, 0xa7, 0x72, 0xb1, 0x8f, 0x57, 0x61, 0xac, 0x5b, 0x07,
	0xf6, 0x67, 0xf1, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0xe5, 0x3c, 0x08, 0x3e, 0x04,
	0x00, 0x00,
}