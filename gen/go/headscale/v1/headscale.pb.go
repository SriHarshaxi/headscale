// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: headscale/v1/headscale.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_headscale_v1_headscale_proto protoreflect.FileDescriptor

var file_headscale_v1_headscale_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68,
	0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69,
	0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb1, 0x16, 0x0a, 0x10, 0x48, 0x65,
	0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21,
	0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x7c, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22,
	0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x96, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x22, 0x2e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2f, 0x7b, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x72, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x2f, 0x7b, 0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x80,
	0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x24, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x12, 0x76, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x25,
	0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x65, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x87, 0x01, 0x0a,
	0x10, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x12, 0x25, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x50, 0x72,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x65, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x2f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x24, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x61, 0x75, 0x74, 0x68, 0x6b,
	0x65, 0x79, 0x12, 0x89, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x62, 0x75, 0x67, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x27, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x62, 0x75, 0x67, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x75,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x1f, 0x2e, 0x68,
	0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x7b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x74, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x61, 0x67, 0x73,
	0x12, 0x1c, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x22, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x7b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x80, 0x01, 0x0a, 0x0f,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12,
	0x24, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x7e,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12,
	0x22, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x2a, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x2f, 0x7b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x85,
	0x01, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x12, 0x22, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x25, 0x22, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x2f, 0x7b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x90, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x22, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x68,
	0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x22, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x7b, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x7b,
	0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x6e, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x68,
	0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x0b, 0x4d, 0x6f,
	0x76, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x20, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x65,
	0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x7b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x8b,
	0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x24, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x7b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x97, 0x01, 0x0a,
	0x13, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x25, 0x22, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x2f, 0x7b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x70, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x69, 0x6b, 0x65, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x77, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x68, 0x65,
	0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x6a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73,
	0x12, 0x20, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x61, 0x6e,
	0x66, 0x6f, 0x6e, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_headscale_v1_headscale_proto_goTypes = []interface{}{
	(*GetNamespaceRequest)(nil),         // 0: headscale.v1.GetNamespaceRequest
	(*CreateNamespaceRequest)(nil),      // 1: headscale.v1.CreateNamespaceRequest
	(*RenameNamespaceRequest)(nil),      // 2: headscale.v1.RenameNamespaceRequest
	(*DeleteNamespaceRequest)(nil),      // 3: headscale.v1.DeleteNamespaceRequest
	(*ListNamespacesRequest)(nil),       // 4: headscale.v1.ListNamespacesRequest
	(*CreatePreAuthKeyRequest)(nil),     // 5: headscale.v1.CreatePreAuthKeyRequest
	(*ExpirePreAuthKeyRequest)(nil),     // 6: headscale.v1.ExpirePreAuthKeyRequest
	(*ListPreAuthKeysRequest)(nil),      // 7: headscale.v1.ListPreAuthKeysRequest
	(*DebugCreateMachineRequest)(nil),   // 8: headscale.v1.DebugCreateMachineRequest
	(*GetMachineRequest)(nil),           // 9: headscale.v1.GetMachineRequest
	(*SetTagsRequest)(nil),              // 10: headscale.v1.SetTagsRequest
	(*RegisterMachineRequest)(nil),      // 11: headscale.v1.RegisterMachineRequest
	(*DeleteMachineRequest)(nil),        // 12: headscale.v1.DeleteMachineRequest
	(*ExpireMachineRequest)(nil),        // 13: headscale.v1.ExpireMachineRequest
	(*RenameMachineRequest)(nil),        // 14: headscale.v1.RenameMachineRequest
	(*ListMachinesRequest)(nil),         // 15: headscale.v1.ListMachinesRequest
	(*MoveMachineRequest)(nil),          // 16: headscale.v1.MoveMachineRequest
	(*GetMachineRouteRequest)(nil),      // 17: headscale.v1.GetMachineRouteRequest
	(*EnableMachineRoutesRequest)(nil),  // 18: headscale.v1.EnableMachineRoutesRequest
	(*CreateApiKeyRequest)(nil),         // 19: headscale.v1.CreateApiKeyRequest
	(*ExpireApiKeyRequest)(nil),         // 20: headscale.v1.ExpireApiKeyRequest
	(*ListApiKeysRequest)(nil),          // 21: headscale.v1.ListApiKeysRequest
	(*GetNamespaceResponse)(nil),        // 22: headscale.v1.GetNamespaceResponse
	(*CreateNamespaceResponse)(nil),     // 23: headscale.v1.CreateNamespaceResponse
	(*RenameNamespaceResponse)(nil),     // 24: headscale.v1.RenameNamespaceResponse
	(*DeleteNamespaceResponse)(nil),     // 25: headscale.v1.DeleteNamespaceResponse
	(*ListNamespacesResponse)(nil),      // 26: headscale.v1.ListNamespacesResponse
	(*CreatePreAuthKeyResponse)(nil),    // 27: headscale.v1.CreatePreAuthKeyResponse
	(*ExpirePreAuthKeyResponse)(nil),    // 28: headscale.v1.ExpirePreAuthKeyResponse
	(*ListPreAuthKeysResponse)(nil),     // 29: headscale.v1.ListPreAuthKeysResponse
	(*DebugCreateMachineResponse)(nil),  // 30: headscale.v1.DebugCreateMachineResponse
	(*GetMachineResponse)(nil),          // 31: headscale.v1.GetMachineResponse
	(*SetTagsResponse)(nil),             // 32: headscale.v1.SetTagsResponse
	(*RegisterMachineResponse)(nil),     // 33: headscale.v1.RegisterMachineResponse
	(*DeleteMachineResponse)(nil),       // 34: headscale.v1.DeleteMachineResponse
	(*ExpireMachineResponse)(nil),       // 35: headscale.v1.ExpireMachineResponse
	(*RenameMachineResponse)(nil),       // 36: headscale.v1.RenameMachineResponse
	(*ListMachinesResponse)(nil),        // 37: headscale.v1.ListMachinesResponse
	(*MoveMachineResponse)(nil),         // 38: headscale.v1.MoveMachineResponse
	(*GetMachineRouteResponse)(nil),     // 39: headscale.v1.GetMachineRouteResponse
	(*EnableMachineRoutesResponse)(nil), // 40: headscale.v1.EnableMachineRoutesResponse
	(*CreateApiKeyResponse)(nil),        // 41: headscale.v1.CreateApiKeyResponse
	(*ExpireApiKeyResponse)(nil),        // 42: headscale.v1.ExpireApiKeyResponse
	(*ListApiKeysResponse)(nil),         // 43: headscale.v1.ListApiKeysResponse
}
var file_headscale_v1_headscale_proto_depIdxs = []int32{
	0,  // 0: headscale.v1.HeadscaleService.GetNamespace:input_type -> headscale.v1.GetNamespaceRequest
	1,  // 1: headscale.v1.HeadscaleService.CreateNamespace:input_type -> headscale.v1.CreateNamespaceRequest
	2,  // 2: headscale.v1.HeadscaleService.RenameNamespace:input_type -> headscale.v1.RenameNamespaceRequest
	3,  // 3: headscale.v1.HeadscaleService.DeleteNamespace:input_type -> headscale.v1.DeleteNamespaceRequest
	4,  // 4: headscale.v1.HeadscaleService.ListNamespaces:input_type -> headscale.v1.ListNamespacesRequest
	5,  // 5: headscale.v1.HeadscaleService.CreatePreAuthKey:input_type -> headscale.v1.CreatePreAuthKeyRequest
	6,  // 6: headscale.v1.HeadscaleService.ExpirePreAuthKey:input_type -> headscale.v1.ExpirePreAuthKeyRequest
	7,  // 7: headscale.v1.HeadscaleService.ListPreAuthKeys:input_type -> headscale.v1.ListPreAuthKeysRequest
	8,  // 8: headscale.v1.HeadscaleService.DebugCreateMachine:input_type -> headscale.v1.DebugCreateMachineRequest
	9,  // 9: headscale.v1.HeadscaleService.GetMachine:input_type -> headscale.v1.GetMachineRequest
	10, // 10: headscale.v1.HeadscaleService.SetTags:input_type -> headscale.v1.SetTagsRequest
	11, // 11: headscale.v1.HeadscaleService.RegisterMachine:input_type -> headscale.v1.RegisterMachineRequest
	12, // 12: headscale.v1.HeadscaleService.DeleteMachine:input_type -> headscale.v1.DeleteMachineRequest
	13, // 13: headscale.v1.HeadscaleService.ExpireMachine:input_type -> headscale.v1.ExpireMachineRequest
	14, // 14: headscale.v1.HeadscaleService.RenameMachine:input_type -> headscale.v1.RenameMachineRequest
	15, // 15: headscale.v1.HeadscaleService.ListMachines:input_type -> headscale.v1.ListMachinesRequest
	16, // 16: headscale.v1.HeadscaleService.MoveMachine:input_type -> headscale.v1.MoveMachineRequest
	17, // 17: headscale.v1.HeadscaleService.GetMachineRoute:input_type -> headscale.v1.GetMachineRouteRequest
	18, // 18: headscale.v1.HeadscaleService.EnableMachineRoutes:input_type -> headscale.v1.EnableMachineRoutesRequest
	19, // 19: headscale.v1.HeadscaleService.CreateApiKey:input_type -> headscale.v1.CreateApiKeyRequest
	20, // 20: headscale.v1.HeadscaleService.ExpireApiKey:input_type -> headscale.v1.ExpireApiKeyRequest
	21, // 21: headscale.v1.HeadscaleService.ListApiKeys:input_type -> headscale.v1.ListApiKeysRequest
	22, // 22: headscale.v1.HeadscaleService.GetNamespace:output_type -> headscale.v1.GetNamespaceResponse
	23, // 23: headscale.v1.HeadscaleService.CreateNamespace:output_type -> headscale.v1.CreateNamespaceResponse
	24, // 24: headscale.v1.HeadscaleService.RenameNamespace:output_type -> headscale.v1.RenameNamespaceResponse
	25, // 25: headscale.v1.HeadscaleService.DeleteNamespace:output_type -> headscale.v1.DeleteNamespaceResponse
	26, // 26: headscale.v1.HeadscaleService.ListNamespaces:output_type -> headscale.v1.ListNamespacesResponse
	27, // 27: headscale.v1.HeadscaleService.CreatePreAuthKey:output_type -> headscale.v1.CreatePreAuthKeyResponse
	28, // 28: headscale.v1.HeadscaleService.ExpirePreAuthKey:output_type -> headscale.v1.ExpirePreAuthKeyResponse
	29, // 29: headscale.v1.HeadscaleService.ListPreAuthKeys:output_type -> headscale.v1.ListPreAuthKeysResponse
	30, // 30: headscale.v1.HeadscaleService.DebugCreateMachine:output_type -> headscale.v1.DebugCreateMachineResponse
	31, // 31: headscale.v1.HeadscaleService.GetMachine:output_type -> headscale.v1.GetMachineResponse
	32, // 32: headscale.v1.HeadscaleService.SetTags:output_type -> headscale.v1.SetTagsResponse
	33, // 33: headscale.v1.HeadscaleService.RegisterMachine:output_type -> headscale.v1.RegisterMachineResponse
	34, // 34: headscale.v1.HeadscaleService.DeleteMachine:output_type -> headscale.v1.DeleteMachineResponse
	35, // 35: headscale.v1.HeadscaleService.ExpireMachine:output_type -> headscale.v1.ExpireMachineResponse
	36, // 36: headscale.v1.HeadscaleService.RenameMachine:output_type -> headscale.v1.RenameMachineResponse
	37, // 37: headscale.v1.HeadscaleService.ListMachines:output_type -> headscale.v1.ListMachinesResponse
	38, // 38: headscale.v1.HeadscaleService.MoveMachine:output_type -> headscale.v1.MoveMachineResponse
	39, // 39: headscale.v1.HeadscaleService.GetMachineRoute:output_type -> headscale.v1.GetMachineRouteResponse
	40, // 40: headscale.v1.HeadscaleService.EnableMachineRoutes:output_type -> headscale.v1.EnableMachineRoutesResponse
	41, // 41: headscale.v1.HeadscaleService.CreateApiKey:output_type -> headscale.v1.CreateApiKeyResponse
	42, // 42: headscale.v1.HeadscaleService.ExpireApiKey:output_type -> headscale.v1.ExpireApiKeyResponse
	43, // 43: headscale.v1.HeadscaleService.ListApiKeys:output_type -> headscale.v1.ListApiKeysResponse
	22, // [22:44] is the sub-list for method output_type
	0,  // [0:22] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_headscale_v1_headscale_proto_init() }
func file_headscale_v1_headscale_proto_init() {
	if File_headscale_v1_headscale_proto != nil {
		return
	}
	file_headscale_v1_namespace_proto_init()
	file_headscale_v1_preauthkey_proto_init()
	file_headscale_v1_machine_proto_init()
	file_headscale_v1_routes_proto_init()
	file_headscale_v1_apikey_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_headscale_v1_headscale_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_headscale_v1_headscale_proto_goTypes,
		DependencyIndexes: file_headscale_v1_headscale_proto_depIdxs,
	}.Build()
	File_headscale_v1_headscale_proto = out.File
	file_headscale_v1_headscale_proto_rawDesc = nil
	file_headscale_v1_headscale_proto_goTypes = nil
	file_headscale_v1_headscale_proto_depIdxs = nil
}
