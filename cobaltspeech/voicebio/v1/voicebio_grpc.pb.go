// Copyright (2023--present) Cobalt Speech and Language Inc.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cobaltspeech/voicebio/v1/voicebio.proto

package voicebiov1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	VoiceBioService_Version_FullMethodName           = "/cobaltspeech.voicebio.v1.VoiceBioService/Version"
	VoiceBioService_ListModels_FullMethodName        = "/cobaltspeech.voicebio.v1.VoiceBioService/ListModels"
	VoiceBioService_StreamingEnroll_FullMethodName   = "/cobaltspeech.voicebio.v1.VoiceBioService/StreamingEnroll"
	VoiceBioService_StreamingVerify_FullMethodName   = "/cobaltspeech.voicebio.v1.VoiceBioService/StreamingVerify"
	VoiceBioService_StreamingIdentify_FullMethodName = "/cobaltspeech.voicebio.v1.VoiceBioService/StreamingIdentify"
)

// VoiceBioServiceClient is the client API for VoiceBioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VoiceBioServiceClient interface {
	// Returns version information from the server.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Returns information about the models available on the server.
	ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error)
	// Uses new audio data to perform enrollment of new users, or to update
	// enrollment of existing users. Returns a new or updated voiceprint.
	//
	// Clients should store the returned voiceprint against the ID of the user
	// that provided the audio. This voiceprint can be provided later, with the
	// Verify or Identify requests to match new audio against known speakers.
	//
	// If this call is used to update an existing user's voiceprint, the old
	// voiceprint can be discarded and only the new one can be stored for that
	// user.
	StreamingEnroll(ctx context.Context, opts ...grpc.CallOption) (VoiceBioService_StreamingEnrollClient, error)
	// Compares audio data against the provided voiceprint and verifies whether or
	// not the audio matches against the voiceprint.
	StreamingVerify(ctx context.Context, opts ...grpc.CallOption) (VoiceBioService_StreamingVerifyClient, error)
	// Compares audio data against the provided list of voiceprints and identifies
	// which (or none) of the voiceprints is a match for the given audio.
	StreamingIdentify(ctx context.Context, opts ...grpc.CallOption) (VoiceBioService_StreamingIdentifyClient, error)
}

type voiceBioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVoiceBioServiceClient(cc grpc.ClientConnInterface) VoiceBioServiceClient {
	return &voiceBioServiceClient{cc}
}

func (c *voiceBioServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, VoiceBioService_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voiceBioServiceClient) ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error) {
	out := new(ListModelsResponse)
	err := c.cc.Invoke(ctx, VoiceBioService_ListModels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voiceBioServiceClient) StreamingEnroll(ctx context.Context, opts ...grpc.CallOption) (VoiceBioService_StreamingEnrollClient, error) {
	stream, err := c.cc.NewStream(ctx, &VoiceBioService_ServiceDesc.Streams[0], VoiceBioService_StreamingEnroll_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &voiceBioServiceStreamingEnrollClient{stream}
	return x, nil
}

type VoiceBioService_StreamingEnrollClient interface {
	Send(*StreamingEnrollRequest) error
	CloseAndRecv() (*StreamingEnrollResponse, error)
	grpc.ClientStream
}

type voiceBioServiceStreamingEnrollClient struct {
	grpc.ClientStream
}

func (x *voiceBioServiceStreamingEnrollClient) Send(m *StreamingEnrollRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *voiceBioServiceStreamingEnrollClient) CloseAndRecv() (*StreamingEnrollResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamingEnrollResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *voiceBioServiceClient) StreamingVerify(ctx context.Context, opts ...grpc.CallOption) (VoiceBioService_StreamingVerifyClient, error) {
	stream, err := c.cc.NewStream(ctx, &VoiceBioService_ServiceDesc.Streams[1], VoiceBioService_StreamingVerify_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &voiceBioServiceStreamingVerifyClient{stream}
	return x, nil
}

type VoiceBioService_StreamingVerifyClient interface {
	Send(*StreamingVerifyRequest) error
	CloseAndRecv() (*StreamingVerifyResponse, error)
	grpc.ClientStream
}

type voiceBioServiceStreamingVerifyClient struct {
	grpc.ClientStream
}

func (x *voiceBioServiceStreamingVerifyClient) Send(m *StreamingVerifyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *voiceBioServiceStreamingVerifyClient) CloseAndRecv() (*StreamingVerifyResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamingVerifyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *voiceBioServiceClient) StreamingIdentify(ctx context.Context, opts ...grpc.CallOption) (VoiceBioService_StreamingIdentifyClient, error) {
	stream, err := c.cc.NewStream(ctx, &VoiceBioService_ServiceDesc.Streams[2], VoiceBioService_StreamingIdentify_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &voiceBioServiceStreamingIdentifyClient{stream}
	return x, nil
}

type VoiceBioService_StreamingIdentifyClient interface {
	Send(*StreamingIdentifyRequest) error
	CloseAndRecv() (*StreamingIdentifyResponse, error)
	grpc.ClientStream
}

type voiceBioServiceStreamingIdentifyClient struct {
	grpc.ClientStream
}

func (x *voiceBioServiceStreamingIdentifyClient) Send(m *StreamingIdentifyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *voiceBioServiceStreamingIdentifyClient) CloseAndRecv() (*StreamingIdentifyResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamingIdentifyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VoiceBioServiceServer is the server API for VoiceBioService service.
// All implementations must embed UnimplementedVoiceBioServiceServer
// for forward compatibility
type VoiceBioServiceServer interface {
	// Returns version information from the server.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// Returns information about the models available on the server.
	ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error)
	// Uses new audio data to perform enrollment of new users, or to update
	// enrollment of existing users. Returns a new or updated voiceprint.
	//
	// Clients should store the returned voiceprint against the ID of the user
	// that provided the audio. This voiceprint can be provided later, with the
	// Verify or Identify requests to match new audio against known speakers.
	//
	// If this call is used to update an existing user's voiceprint, the old
	// voiceprint can be discarded and only the new one can be stored for that
	// user.
	StreamingEnroll(VoiceBioService_StreamingEnrollServer) error
	// Compares audio data against the provided voiceprint and verifies whether or
	// not the audio matches against the voiceprint.
	StreamingVerify(VoiceBioService_StreamingVerifyServer) error
	// Compares audio data against the provided list of voiceprints and identifies
	// which (or none) of the voiceprints is a match for the given audio.
	StreamingIdentify(VoiceBioService_StreamingIdentifyServer) error
	mustEmbedUnimplementedVoiceBioServiceServer()
}

// UnimplementedVoiceBioServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVoiceBioServiceServer struct {
}

func (UnimplementedVoiceBioServiceServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedVoiceBioServiceServer) ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModels not implemented")
}
func (UnimplementedVoiceBioServiceServer) StreamingEnroll(VoiceBioService_StreamingEnrollServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingEnroll not implemented")
}
func (UnimplementedVoiceBioServiceServer) StreamingVerify(VoiceBioService_StreamingVerifyServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingVerify not implemented")
}
func (UnimplementedVoiceBioServiceServer) StreamingIdentify(VoiceBioService_StreamingIdentifyServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingIdentify not implemented")
}
func (UnimplementedVoiceBioServiceServer) mustEmbedUnimplementedVoiceBioServiceServer() {}

// UnsafeVoiceBioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VoiceBioServiceServer will
// result in compilation errors.
type UnsafeVoiceBioServiceServer interface {
	mustEmbedUnimplementedVoiceBioServiceServer()
}

func RegisterVoiceBioServiceServer(s grpc.ServiceRegistrar, srv VoiceBioServiceServer) {
	s.RegisterService(&VoiceBioService_ServiceDesc, srv)
}

func _VoiceBioService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoiceBioServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoiceBioService_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoiceBioServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoiceBioService_ListModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoiceBioServiceServer).ListModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoiceBioService_ListModels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoiceBioServiceServer).ListModels(ctx, req.(*ListModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoiceBioService_StreamingEnroll_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VoiceBioServiceServer).StreamingEnroll(&voiceBioServiceStreamingEnrollServer{stream})
}

type VoiceBioService_StreamingEnrollServer interface {
	SendAndClose(*StreamingEnrollResponse) error
	Recv() (*StreamingEnrollRequest, error)
	grpc.ServerStream
}

type voiceBioServiceStreamingEnrollServer struct {
	grpc.ServerStream
}

func (x *voiceBioServiceStreamingEnrollServer) SendAndClose(m *StreamingEnrollResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *voiceBioServiceStreamingEnrollServer) Recv() (*StreamingEnrollRequest, error) {
	m := new(StreamingEnrollRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VoiceBioService_StreamingVerify_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VoiceBioServiceServer).StreamingVerify(&voiceBioServiceStreamingVerifyServer{stream})
}

type VoiceBioService_StreamingVerifyServer interface {
	SendAndClose(*StreamingVerifyResponse) error
	Recv() (*StreamingVerifyRequest, error)
	grpc.ServerStream
}

type voiceBioServiceStreamingVerifyServer struct {
	grpc.ServerStream
}

func (x *voiceBioServiceStreamingVerifyServer) SendAndClose(m *StreamingVerifyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *voiceBioServiceStreamingVerifyServer) Recv() (*StreamingVerifyRequest, error) {
	m := new(StreamingVerifyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VoiceBioService_StreamingIdentify_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VoiceBioServiceServer).StreamingIdentify(&voiceBioServiceStreamingIdentifyServer{stream})
}

type VoiceBioService_StreamingIdentifyServer interface {
	SendAndClose(*StreamingIdentifyResponse) error
	Recv() (*StreamingIdentifyRequest, error)
	grpc.ServerStream
}

type voiceBioServiceStreamingIdentifyServer struct {
	grpc.ServerStream
}

func (x *voiceBioServiceStreamingIdentifyServer) SendAndClose(m *StreamingIdentifyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *voiceBioServiceStreamingIdentifyServer) Recv() (*StreamingIdentifyRequest, error) {
	m := new(StreamingIdentifyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VoiceBioService_ServiceDesc is the grpc.ServiceDesc for VoiceBioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VoiceBioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cobaltspeech.voicebio.v1.VoiceBioService",
	HandlerType: (*VoiceBioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _VoiceBioService_Version_Handler,
		},
		{
			MethodName: "ListModels",
			Handler:    _VoiceBioService_ListModels_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingEnroll",
			Handler:       _VoiceBioService_StreamingEnroll_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamingVerify",
			Handler:       _VoiceBioService_StreamingVerify_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamingIdentify",
			Handler:       _VoiceBioService_StreamingIdentify_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "cobaltspeech/voicebio/v1/voicebio.proto",
}