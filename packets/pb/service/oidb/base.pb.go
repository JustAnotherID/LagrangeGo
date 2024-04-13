// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/base.proto

package oidb

import (
	anypb "google.golang.org/protobuf/types/known/anypb"
)

type OidbSvcTrpcTcpBase struct {
	Command    uint32          `protobuf:"varint,1,opt"`
	SubCommand uint32          `protobuf:"varint,2,opt"`
	ErrorCode  uint32          `protobuf:"varint,3,opt"`
	Body       *anypb.Any      `protobuf:"bytes,4,opt"`
	ErrorMsg   string          `protobuf:"bytes,5,opt"`
	Lafter     *OidbLafter     `protobuf:"bytes,7,opt"`
	Properties []*OidbProperty `protobuf:"bytes,11,rep"`
	Reserved   int32           `protobuf:"varint,12,opt"`
}