// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connect.proto

/*
Package connectproto is a generated protocol buffer package.

It is generated from these files:
	connect.proto

It has these top-level messages:
	Event
	Response
*/
package connectproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	Id       string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Time     int64     `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Response *Response `protobuf:"bytes,3,opt,name=response" json:"response,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Event) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type Response struct {
	Udid        string `protobuf:"bytes,1,opt,name=udid" json:"udid,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Status      string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	RequestType string `protobuf:"bytes,4,opt,name=request_type,json=requestType" json:"request_type,omitempty"`
	CommandUuid string `protobuf:"bytes,5,opt,name=command_uuid,json=commandUuid" json:"command_uuid,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetUdid() string {
	if m != nil {
		return m.Udid
	}
	return ""
}

func (m *Response) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Response) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *Response) GetCommandUuid() string {
	if m != nil {
		return m.CommandUuid
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "connectproto.Event")
	proto.RegisterType((*Response)(nil), "connectproto.Response")
}

func init() { proto.RegisterFile("connect.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xb1, 0x52, 0xc3, 0x30,
	0x10, 0x44, 0x47, 0x4e, 0x62, 0xe2, 0x4b, 0xa0, 0xb8, 0x22, 0xa8, 0x34, 0xa9, 0x5c, 0xb9, 0x08,
	0xdf, 0x40, 0x41, 0xab, 0x81, 0xda, 0x13, 0x7c, 0x57, 0xa8, 0xb0, 0x64, 0xa4, 0x13, 0x33, 0xf9,
	0x10, 0xfe, 0x97, 0xb1, 0x10, 0x1e, 0xba, 0x7b, 0xbb, 0x3b, 0xbb, 0x07, 0xf7, 0xa3, 0x77, 0x8e,
	0x47, 0xe9, 0xe7, 0xe0, 0xc5, 0xe3, 0xb1, 0x60, 0xa6, 0xf3, 0x00, 0xbb, 0x97, 0x2f, 0x76, 0x82,
	0x0f, 0x50, 0x59, 0xd2, 0xaa, 0x55, 0x5d, 0x63, 0x2a, 0x4b, 0x88, 0xb0, 0x15, 0x3b, 0xb1, 0xae,
	0x5a, 0xd5, 0x6d, 0x4c, 0xbe, 0xf1, 0x02, 0xfb, 0xc0, 0x71, 0xf6, 0x2e, 0xb2, 0xde, 0xb4, 0xaa,
	0x3b, 0x5c, 0x4e, 0xfd, 0xff, 0xb6, 0xde, 0x14, 0xd7, 0xac, 0xb9, 0xf3, 0xb7, 0x82, 0xfd, 0x9f,
	0xbc, 0x94, 0x26, 0x5a, 0x67, 0xf2, 0x8d, 0x8f, 0x70, 0x97, 0x22, 0x87, 0xc1, 0x52, 0xde, 0x6a,
	0x4c, 0xbd, 0xe0, 0x2b, 0xe1, 0x09, 0xea, 0x28, 0x57, 0x49, 0x31, 0x6f, 0x35, 0xa6, 0x10, 0x3e,
	0xc1, 0x31, 0xf0, 0x67, 0xe2, 0x28, 0x83, 0xdc, 0x66, 0xd6, 0xdb, 0xec, 0x1e, 0x8a, 0xf6, 0x76,
	0x9b, 0x79, 0x89, 0x8c, 0x7e, 0x9a, 0xae, 0x8e, 0x86, 0x94, 0x2c, 0xe9, 0xdd, 0x6f, 0xa4, 0x68,
	0xef, 0xc9, 0xd2, 0x47, 0x9d, 0x3f, 0x7e, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x70, 0xa9, 0xb3,
	0xa5, 0x1e, 0x01, 0x00, 0x00,
}
