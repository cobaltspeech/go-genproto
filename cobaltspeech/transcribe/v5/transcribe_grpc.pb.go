// Copyright (2022--present) Cobalt Speech and Language Inc.

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

//*
//Cobalt ASR is a state-of-the-art speech recognition system, which uses deep-learning models for fast, accurate speech recognition.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cobaltspeech/transcribe/v5/transcribe.proto

package transcribev5

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
	TranscribeService_Version_FullMethodName            = "/cobaltspeech.transcribe.v5.TranscribeService/Version"
	TranscribeService_ListModels_FullMethodName         = "/cobaltspeech.transcribe.v5.TranscribeService/ListModels"
	TranscribeService_StreamingRecognize_FullMethodName = "/cobaltspeech.transcribe.v5.TranscribeService/StreamingRecognize"
	TranscribeService_CompileContext_FullMethodName     = "/cobaltspeech.transcribe.v5.TranscribeService/CompileContext"
)

// TranscribeServiceClient is the client API for TranscribeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranscribeServiceClient interface {
	// Queries the version of the server.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Retrieves a list of available speech recognition models.
	ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error)
	// Performs bidirectional streaming speech recognition. Receive results while
	// sending audio. This method is only available via GRPC and not via
	// HTTP+JSON. However, a web browser may use websockets to use this service.
	StreamingRecognize(ctx context.Context, opts ...grpc.CallOption) (TranscribeService_StreamingRecognizeClient, error)
	// Compiles recognition context information, such as a specialized list of
	// words or phrases, into a compact, efficient form to send with subsequent
	// `StreamingRecognize` requests to customize speech recognition. For example,
	// a list of contact names may be compiled in a mobile app and sent with each
	// recognition request so that the app user's contact names are more likely to
	// be recognized than arbitrary names. This pre-compilation ensures that there
	// is no added latency for the recognition request. It is important to note
	// that in order to compile context for a model, that model has to support
	// context in the first place, which can be verified by checking its
	// `ModelAttributes.ContextInfo` obtained via the `ListModels` method. Also,
	// the compiled data will be model specific; that is, the data compiled for
	// one model will generally not be usable with a different model.
	CompileContext(ctx context.Context, in *CompileContextRequest, opts ...grpc.CallOption) (*CompileContextResponse, error)
}

type transcribeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTranscribeServiceClient(cc grpc.ClientConnInterface) TranscribeServiceClient {
	return &transcribeServiceClient{cc}
}

func (c *transcribeServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, TranscribeService_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transcribeServiceClient) ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error) {
	out := new(ListModelsResponse)
	err := c.cc.Invoke(ctx, TranscribeService_ListModels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transcribeServiceClient) StreamingRecognize(ctx context.Context, opts ...grpc.CallOption) (TranscribeService_StreamingRecognizeClient, error) {
	stream, err := c.cc.NewStream(ctx, &TranscribeService_ServiceDesc.Streams[0], TranscribeService_StreamingRecognize_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transcribeServiceStreamingRecognizeClient{stream}
	return x, nil
}

type TranscribeService_StreamingRecognizeClient interface {
	Send(*StreamingRecognizeRequest) error
	Recv() (*StreamingRecognizeResponse, error)
	grpc.ClientStream
}

type transcribeServiceStreamingRecognizeClient struct {
	grpc.ClientStream
}

func (x *transcribeServiceStreamingRecognizeClient) Send(m *StreamingRecognizeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transcribeServiceStreamingRecognizeClient) Recv() (*StreamingRecognizeResponse, error) {
	m := new(StreamingRecognizeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transcribeServiceClient) CompileContext(ctx context.Context, in *CompileContextRequest, opts ...grpc.CallOption) (*CompileContextResponse, error) {
	out := new(CompileContextResponse)
	err := c.cc.Invoke(ctx, TranscribeService_CompileContext_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranscribeServiceServer is the server API for TranscribeService service.
// All implementations must embed UnimplementedTranscribeServiceServer
// for forward compatibility
type TranscribeServiceServer interface {
	// Queries the version of the server.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// Retrieves a list of available speech recognition models.
	ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error)
	// Performs bidirectional streaming speech recognition. Receive results while
	// sending audio. This method is only available via GRPC and not via
	// HTTP+JSON. However, a web browser may use websockets to use this service.
	StreamingRecognize(TranscribeService_StreamingRecognizeServer) error
	// Compiles recognition context information, such as a specialized list of
	// words or phrases, into a compact, efficient form to send with subsequent
	// `StreamingRecognize` requests to customize speech recognition. For example,
	// a list of contact names may be compiled in a mobile app and sent with each
	// recognition request so that the app user's contact names are more likely to
	// be recognized than arbitrary names. This pre-compilation ensures that there
	// is no added latency for the recognition request. It is important to note
	// that in order to compile context for a model, that model has to support
	// context in the first place, which can be verified by checking its
	// `ModelAttributes.ContextInfo` obtained via the `ListModels` method. Also,
	// the compiled data will be model specific; that is, the data compiled for
	// one model will generally not be usable with a different model.
	CompileContext(context.Context, *CompileContextRequest) (*CompileContextResponse, error)
	mustEmbedUnimplementedTranscribeServiceServer()
}

// UnimplementedTranscribeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTranscribeServiceServer struct {
}

func (UnimplementedTranscribeServiceServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedTranscribeServiceServer) ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModels not implemented")
}
func (UnimplementedTranscribeServiceServer) StreamingRecognize(TranscribeService_StreamingRecognizeServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingRecognize not implemented")
}
func (UnimplementedTranscribeServiceServer) CompileContext(context.Context, *CompileContextRequest) (*CompileContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompileContext not implemented")
}
func (UnimplementedTranscribeServiceServer) mustEmbedUnimplementedTranscribeServiceServer() {}

// UnsafeTranscribeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranscribeServiceServer will
// result in compilation errors.
type UnsafeTranscribeServiceServer interface {
	mustEmbedUnimplementedTranscribeServiceServer()
}

func RegisterTranscribeServiceServer(s grpc.ServiceRegistrar, srv TranscribeServiceServer) {
	s.RegisterService(&TranscribeService_ServiceDesc, srv)
}

func _TranscribeService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranscribeServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranscribeService_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranscribeServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranscribeService_ListModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranscribeServiceServer).ListModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranscribeService_ListModels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranscribeServiceServer).ListModels(ctx, req.(*ListModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranscribeService_StreamingRecognize_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TranscribeServiceServer).StreamingRecognize(&transcribeServiceStreamingRecognizeServer{stream})
}

type TranscribeService_StreamingRecognizeServer interface {
	Send(*StreamingRecognizeResponse) error
	Recv() (*StreamingRecognizeRequest, error)
	grpc.ServerStream
}

type transcribeServiceStreamingRecognizeServer struct {
	grpc.ServerStream
}

func (x *transcribeServiceStreamingRecognizeServer) Send(m *StreamingRecognizeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transcribeServiceStreamingRecognizeServer) Recv() (*StreamingRecognizeRequest, error) {
	m := new(StreamingRecognizeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TranscribeService_CompileContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompileContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranscribeServiceServer).CompileContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranscribeService_CompileContext_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranscribeServiceServer).CompileContext(ctx, req.(*CompileContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TranscribeService_ServiceDesc is the grpc.ServiceDesc for TranscribeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TranscribeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cobaltspeech.transcribe.v5.TranscribeService",
	HandlerType: (*TranscribeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _TranscribeService_Version_Handler,
		},
		{
			MethodName: "ListModels",
			Handler:    _TranscribeService_ListModels_Handler,
		},
		{
			MethodName: "CompileContext",
			Handler:    _TranscribeService_CompileContext_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingRecognize",
			Handler:       _TranscribeService_StreamingRecognize_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cobaltspeech/transcribe/v5/transcribe.proto",
}