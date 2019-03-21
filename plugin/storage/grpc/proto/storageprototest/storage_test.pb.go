// Code generated by protoc-gen-go. DO NOT EDIT.

// Copyright (c) 2019 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// source: storage_test.proto

package storageprototest

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

type GetTraceRequest struct {
	TraceId              []byte   `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTraceRequest) Reset()         { *m = GetTraceRequest{} }
func (m *GetTraceRequest) String() string { return proto.CompactTextString(m) }
func (*GetTraceRequest) ProtoMessage()    {}
func (*GetTraceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_test_9c088335406789b5, []int{0}
}
func (m *GetTraceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTraceRequest.Unmarshal(m, b)
}
func (m *GetTraceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTraceRequest.Marshal(b, m, deterministic)
}
func (dst *GetTraceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTraceRequest.Merge(dst, src)
}
func (m *GetTraceRequest) XXX_Size() int {
	return xxx_messageInfo_GetTraceRequest.Size(m)
}
func (m *GetTraceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTraceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTraceRequest proto.InternalMessageInfo

func (m *GetTraceRequest) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTraceRequest)(nil), "storageprototest.GetTraceRequest")
}

func init() { proto.RegisterFile("storage_test.proto", fileDescriptor_storage_test_9c088335406789b5) }

var fileDescriptor_storage_test_9c088335406789b5 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0x8d, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x80, 0x8a, 0x81, 0x79, 0x20, 0x71, 0x25, 0x1d, 0x2e, 0x7e, 0xf7, 0xd4, 0x92, 0x90, 0xa2, 0xc4,
	0xe4, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x49, 0x2e, 0x8e, 0x12, 0x10, 0x3f,
	0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x27, 0x88, 0x1d, 0xcc, 0xf7, 0x4c, 0x49, 0x62,
	0x03, 0x6b, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x56, 0x76, 0x68, 0x5c, 0x00, 0x00,
	0x00,
}
