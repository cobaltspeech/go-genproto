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
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cobaltspeech/privacyscreen/v1/privacyscreen.proto

package privacyscreenv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PrivacyScreenService_Version_FullMethodName                         = "/cobaltspeech.privacyscreen.v1.PrivacyScreenService/Version"
	PrivacyScreenService_ListModels_FullMethodName                      = "/cobaltspeech.privacyscreen.v1.PrivacyScreenService/ListModels"
	PrivacyScreenService_RedactText_FullMethodName                      = "/cobaltspeech.privacyscreen.v1.PrivacyScreenService/RedactText"
	PrivacyScreenService_RedactTranscript_FullMethodName                = "/cobaltspeech.privacyscreen.v1.PrivacyScreenService/RedactTranscript"
	PrivacyScreenService_StreamingRedactTranscribedAudio_FullMethodName = "/cobaltspeech.privacyscreen.v1.PrivacyScreenService/StreamingRedactTranscribedAudio"
	PrivacyScreenService_StreamingTranscribeAndRedact_FullMethodName    = "/cobaltspeech.privacyscreen.v1.PrivacyScreenService/StreamingTranscribeAndRedact"
)

// PrivacyScreenServiceClient is the client API for PrivacyScreenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service that implements the Cobalt Privacy Screen API.
type PrivacyScreenServiceClient interface {
	// Returns version information from the server.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// ListModels returns information about the models the server can access.
	ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error)
	// Redact text using a redaction engine that is configured with the provided
	// redaction configuration.
	RedactText(ctx context.Context, in *RedactTextRequest, opts ...grpc.CallOption) (*RedactTextResponse, error)
	// redacts transcript using a redaction engine that is configured with the
	// provided redaction configuration.
	RedactTranscript(ctx context.Context, in *RedactTranscriptRequest, opts ...grpc.CallOption) (*RedactTranscriptResponse, error)
	// Performs bidirectional streaming redaction on transcribed audio. Receive
	// redacted audio while sending audio. The transcription of audio data must be
	// ready before sending the audio.
	StreamingRedactTranscribedAudio(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamingRedactTranscribedAudioRequest, StreamingRedactTranscribedAudioResponse], error)
	// Performs bidirectional streaming speech recognition and redaction. Receive
	// redacted audio and transcriptions while sending audio.
	StreamingTranscribeAndRedact(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamingTranscribeAndRedactRequest, StreamingTranscribeAndRedactResponse], error)
}

type privacyScreenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivacyScreenServiceClient(cc grpc.ClientConnInterface) PrivacyScreenServiceClient {
	return &privacyScreenServiceClient{cc}
}

func (c *privacyScreenServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, PrivacyScreenService_Version_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyScreenServiceClient) ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListModelsResponse)
	err := c.cc.Invoke(ctx, PrivacyScreenService_ListModels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyScreenServiceClient) RedactText(ctx context.Context, in *RedactTextRequest, opts ...grpc.CallOption) (*RedactTextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RedactTextResponse)
	err := c.cc.Invoke(ctx, PrivacyScreenService_RedactText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyScreenServiceClient) RedactTranscript(ctx context.Context, in *RedactTranscriptRequest, opts ...grpc.CallOption) (*RedactTranscriptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RedactTranscriptResponse)
	err := c.cc.Invoke(ctx, PrivacyScreenService_RedactTranscript_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyScreenServiceClient) StreamingRedactTranscribedAudio(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamingRedactTranscribedAudioRequest, StreamingRedactTranscribedAudioResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PrivacyScreenService_ServiceDesc.Streams[0], PrivacyScreenService_StreamingRedactTranscribedAudio_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamingRedactTranscribedAudioRequest, StreamingRedactTranscribedAudioResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PrivacyScreenService_StreamingRedactTranscribedAudioClient = grpc.BidiStreamingClient[StreamingRedactTranscribedAudioRequest, StreamingRedactTranscribedAudioResponse]

func (c *privacyScreenServiceClient) StreamingTranscribeAndRedact(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamingTranscribeAndRedactRequest, StreamingTranscribeAndRedactResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PrivacyScreenService_ServiceDesc.Streams[1], PrivacyScreenService_StreamingTranscribeAndRedact_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamingTranscribeAndRedactRequest, StreamingTranscribeAndRedactResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PrivacyScreenService_StreamingTranscribeAndRedactClient = grpc.BidiStreamingClient[StreamingTranscribeAndRedactRequest, StreamingTranscribeAndRedactResponse]

// PrivacyScreenServiceServer is the server API for PrivacyScreenService service.
// All implementations must embed UnimplementedPrivacyScreenServiceServer
// for forward compatibility.
//
// Service that implements the Cobalt Privacy Screen API.
type PrivacyScreenServiceServer interface {
	// Returns version information from the server.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// ListModels returns information about the models the server can access.
	ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error)
	// Redact text using a redaction engine that is configured with the provided
	// redaction configuration.
	RedactText(context.Context, *RedactTextRequest) (*RedactTextResponse, error)
	// redacts transcript using a redaction engine that is configured with the
	// provided redaction configuration.
	RedactTranscript(context.Context, *RedactTranscriptRequest) (*RedactTranscriptResponse, error)
	// Performs bidirectional streaming redaction on transcribed audio. Receive
	// redacted audio while sending audio. The transcription of audio data must be
	// ready before sending the audio.
	StreamingRedactTranscribedAudio(grpc.BidiStreamingServer[StreamingRedactTranscribedAudioRequest, StreamingRedactTranscribedAudioResponse]) error
	// Performs bidirectional streaming speech recognition and redaction. Receive
	// redacted audio and transcriptions while sending audio.
	StreamingTranscribeAndRedact(grpc.BidiStreamingServer[StreamingTranscribeAndRedactRequest, StreamingTranscribeAndRedactResponse]) error
	mustEmbedUnimplementedPrivacyScreenServiceServer()
}

// UnimplementedPrivacyScreenServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPrivacyScreenServiceServer struct{}

func (UnimplementedPrivacyScreenServiceServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedPrivacyScreenServiceServer) ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModels not implemented")
}
func (UnimplementedPrivacyScreenServiceServer) RedactText(context.Context, *RedactTextRequest) (*RedactTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedactText not implemented")
}
func (UnimplementedPrivacyScreenServiceServer) RedactTranscript(context.Context, *RedactTranscriptRequest) (*RedactTranscriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedactTranscript not implemented")
}
func (UnimplementedPrivacyScreenServiceServer) StreamingRedactTranscribedAudio(grpc.BidiStreamingServer[StreamingRedactTranscribedAudioRequest, StreamingRedactTranscribedAudioResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamingRedactTranscribedAudio not implemented")
}
func (UnimplementedPrivacyScreenServiceServer) StreamingTranscribeAndRedact(grpc.BidiStreamingServer[StreamingTranscribeAndRedactRequest, StreamingTranscribeAndRedactResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamingTranscribeAndRedact not implemented")
}
func (UnimplementedPrivacyScreenServiceServer) mustEmbedUnimplementedPrivacyScreenServiceServer() {}
func (UnimplementedPrivacyScreenServiceServer) testEmbeddedByValue()                              {}

// UnsafePrivacyScreenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivacyScreenServiceServer will
// result in compilation errors.
type UnsafePrivacyScreenServiceServer interface {
	mustEmbedUnimplementedPrivacyScreenServiceServer()
}

func RegisterPrivacyScreenServiceServer(s grpc.ServiceRegistrar, srv PrivacyScreenServiceServer) {
	// If the following call pancis, it indicates UnimplementedPrivacyScreenServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PrivacyScreenService_ServiceDesc, srv)
}

func _PrivacyScreenService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyScreenServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivacyScreenService_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyScreenServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyScreenService_ListModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyScreenServiceServer).ListModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivacyScreenService_ListModels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyScreenServiceServer).ListModels(ctx, req.(*ListModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyScreenService_RedactText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedactTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyScreenServiceServer).RedactText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivacyScreenService_RedactText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyScreenServiceServer).RedactText(ctx, req.(*RedactTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyScreenService_RedactTranscript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedactTranscriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyScreenServiceServer).RedactTranscript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivacyScreenService_RedactTranscript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyScreenServiceServer).RedactTranscript(ctx, req.(*RedactTranscriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyScreenService_StreamingRedactTranscribedAudio_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PrivacyScreenServiceServer).StreamingRedactTranscribedAudio(&grpc.GenericServerStream[StreamingRedactTranscribedAudioRequest, StreamingRedactTranscribedAudioResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PrivacyScreenService_StreamingRedactTranscribedAudioServer = grpc.BidiStreamingServer[StreamingRedactTranscribedAudioRequest, StreamingRedactTranscribedAudioResponse]

func _PrivacyScreenService_StreamingTranscribeAndRedact_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PrivacyScreenServiceServer).StreamingTranscribeAndRedact(&grpc.GenericServerStream[StreamingTranscribeAndRedactRequest, StreamingTranscribeAndRedactResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PrivacyScreenService_StreamingTranscribeAndRedactServer = grpc.BidiStreamingServer[StreamingTranscribeAndRedactRequest, StreamingTranscribeAndRedactResponse]

// PrivacyScreenService_ServiceDesc is the grpc.ServiceDesc for PrivacyScreenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrivacyScreenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cobaltspeech.privacyscreen.v1.PrivacyScreenService",
	HandlerType: (*PrivacyScreenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _PrivacyScreenService_Version_Handler,
		},
		{
			MethodName: "ListModels",
			Handler:    _PrivacyScreenService_ListModels_Handler,
		},
		{
			MethodName: "RedactText",
			Handler:    _PrivacyScreenService_RedactText_Handler,
		},
		{
			MethodName: "RedactTranscript",
			Handler:    _PrivacyScreenService_RedactTranscript_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingRedactTranscribedAudio",
			Handler:       _PrivacyScreenService_StreamingRedactTranscribedAudio_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamingTranscribeAndRedact",
			Handler:       _PrivacyScreenService_StreamingTranscribeAndRedact_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cobaltspeech/privacyscreen/v1/privacyscreen.proto",
}
