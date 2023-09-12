// Copyright (2021) Cobalt Speech and Language Inc.

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
// source: cobaltspeech/interact/v3/interact.proto

package interactv3

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
	InteractService_Version_FullMethodName               = "/cobaltspeech.interact.v3.InteractService/Version"
	InteractService_ListModels_FullMethodName            = "/cobaltspeech.interact.v3.InteractService/ListModels"
	InteractService_CreateSession_FullMethodName         = "/cobaltspeech.interact.v3.InteractService/CreateSession"
	InteractService_DeleteSession_FullMethodName         = "/cobaltspeech.interact.v3.InteractService/DeleteSession"
	InteractService_UpdateSession_FullMethodName         = "/cobaltspeech.interact.v3.InteractService/UpdateSession"
	InteractService_StreamASR_FullMethodName             = "/cobaltspeech.interact.v3.InteractService/StreamASR"
	InteractService_StreamTTS_FullMethodName             = "/cobaltspeech.interact.v3.InteractService/StreamTTS"
	InteractService_Transcribe_FullMethodName            = "/cobaltspeech.interact.v3.InteractService/Transcribe"
	InteractService_StreamASRWithPartials_FullMethodName = "/cobaltspeech.interact.v3.InteractService/StreamASRWithPartials"
)

// InteractServiceClient is the client API for InteractService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractServiceClient interface {
	// Returns version information from the server.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// ListModels returns information about the Cobalt Interact models
	// the server can access.
	ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error)
	// Create a new Cobalt Interact session. Also returns a list of
	// actions to take next.
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	// Delete the session. Behavior is undefined if the given
	// TokenData in the request is used again after this function is called.
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error)
	// Process input for a session and get an updated session with
	// a list of actions to take next. This is the only method
	// that modifies the Cobalt Interact session state.
	UpdateSession(ctx context.Context, in *UpdateSessionRequest, opts ...grpc.CallOption) (*UpdateSessionResponse, error)
	// Create an ASR stream. A result is returned when the
	// stream is closed by the client (which forces the ASR to
	// endpoint), or when a transcript becomes available on its
	// own, in which case the stream is closed by the server.
	// The ASR result may be used in the UpdateSession method.
	// <br/><br/>
	// If the session has a wakeword enabled, and the client
	// application is using Cobalt Interact and Cobalt Transcribe
	// to handle the wakeword processing, this method will not
	// return a result until the wakeword condition has been
	// satisfied. Utterances without the required wakeword will
	// be discarded and no transcription will be returned.
	StreamASR(ctx context.Context, opts ...grpc.CallOption) (InteractService_StreamASRClient, error)
	// Create a TTS stream to receive audio for the given reply.
	// The stream will close when TTS is finished. The client
	// may also close the stream early to cancel the speech
	// synthesis.
	StreamTTS(ctx context.Context, in *StreamTTSRequest, opts ...grpc.CallOption) (InteractService_StreamTTSClient, error)
	// Create an ASR stream for transcription. Unlike StreamASR,
	// Transcribe does not listen for a wakeword. This method
	// returns a bi-directional stream, and its intended use is
	// for situations where a user may say anything at all, whether
	// it is short or long, and the application wants to save the
	// transcript (e.g., take a note, send a message).
	// <br/><br/>
	// The first message sent to the server must be the TranscribeAction,
	// with remaining messages sending audio data.
	// Messages received from the server will include the current
	// best partial transcription until the full transcription is
	// ready. The stream ends when either the client application
	// closes it, a predefined duration of silence (non-speech)
	// occurs, or the end-transcription intent is recognized.
	Transcribe(ctx context.Context, opts ...grpc.CallOption) (InteractService_TranscribeClient, error)
	// Performs bidirectional streaming speech recognition. Receive results while
	// sending audio. Each result will either be a partial ASR result, or a final
	// result. Partial results will be sent as soon as they are ready, and all
	// results will be sent, regardless of any wakeword configuration in the
	// session. A final result will be sent exactly once, and the stream will be
	// closed then. If a session has a wakeword enabled, the final result will
	// only be emitted if the required wakeword is present. The ASRResult in the
	// final message maybe used in the UpdateSession method for further dialog
	// processing.
	StreamASRWithPartials(ctx context.Context, opts ...grpc.CallOption) (InteractService_StreamASRWithPartialsClient, error)
}

type interactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractServiceClient(cc grpc.ClientConnInterface) InteractServiceClient {
	return &interactServiceClient{cc}
}

func (c *interactServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, InteractService_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactServiceClient) ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error) {
	out := new(ListModelsResponse)
	err := c.cc.Invoke(ctx, InteractService_ListModels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, InteractService_CreateSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactServiceClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error) {
	out := new(DeleteSessionResponse)
	err := c.cc.Invoke(ctx, InteractService_DeleteSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactServiceClient) UpdateSession(ctx context.Context, in *UpdateSessionRequest, opts ...grpc.CallOption) (*UpdateSessionResponse, error) {
	out := new(UpdateSessionResponse)
	err := c.cc.Invoke(ctx, InteractService_UpdateSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactServiceClient) StreamASR(ctx context.Context, opts ...grpc.CallOption) (InteractService_StreamASRClient, error) {
	stream, err := c.cc.NewStream(ctx, &InteractService_ServiceDesc.Streams[0], InteractService_StreamASR_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &interactServiceStreamASRClient{stream}
	return x, nil
}

type InteractService_StreamASRClient interface {
	Send(*StreamASRRequest) error
	CloseAndRecv() (*StreamASRResponse, error)
	grpc.ClientStream
}

type interactServiceStreamASRClient struct {
	grpc.ClientStream
}

func (x *interactServiceStreamASRClient) Send(m *StreamASRRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *interactServiceStreamASRClient) CloseAndRecv() (*StreamASRResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamASRResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *interactServiceClient) StreamTTS(ctx context.Context, in *StreamTTSRequest, opts ...grpc.CallOption) (InteractService_StreamTTSClient, error) {
	stream, err := c.cc.NewStream(ctx, &InteractService_ServiceDesc.Streams[1], InteractService_StreamTTS_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &interactServiceStreamTTSClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InteractService_StreamTTSClient interface {
	Recv() (*StreamTTSResponse, error)
	grpc.ClientStream
}

type interactServiceStreamTTSClient struct {
	grpc.ClientStream
}

func (x *interactServiceStreamTTSClient) Recv() (*StreamTTSResponse, error) {
	m := new(StreamTTSResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *interactServiceClient) Transcribe(ctx context.Context, opts ...grpc.CallOption) (InteractService_TranscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &InteractService_ServiceDesc.Streams[2], InteractService_Transcribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &interactServiceTranscribeClient{stream}
	return x, nil
}

type InteractService_TranscribeClient interface {
	Send(*TranscribeRequest) error
	Recv() (*TranscribeResponse, error)
	grpc.ClientStream
}

type interactServiceTranscribeClient struct {
	grpc.ClientStream
}

func (x *interactServiceTranscribeClient) Send(m *TranscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *interactServiceTranscribeClient) Recv() (*TranscribeResponse, error) {
	m := new(TranscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *interactServiceClient) StreamASRWithPartials(ctx context.Context, opts ...grpc.CallOption) (InteractService_StreamASRWithPartialsClient, error) {
	stream, err := c.cc.NewStream(ctx, &InteractService_ServiceDesc.Streams[3], InteractService_StreamASRWithPartials_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &interactServiceStreamASRWithPartialsClient{stream}
	return x, nil
}

type InteractService_StreamASRWithPartialsClient interface {
	Send(*StreamASRWithPartialsRequest) error
	Recv() (*StreamASRWithPartialsResponse, error)
	grpc.ClientStream
}

type interactServiceStreamASRWithPartialsClient struct {
	grpc.ClientStream
}

func (x *interactServiceStreamASRWithPartialsClient) Send(m *StreamASRWithPartialsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *interactServiceStreamASRWithPartialsClient) Recv() (*StreamASRWithPartialsResponse, error) {
	m := new(StreamASRWithPartialsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InteractServiceServer is the server API for InteractService service.
// All implementations must embed UnimplementedInteractServiceServer
// for forward compatibility
type InteractServiceServer interface {
	// Returns version information from the server.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// ListModels returns information about the Cobalt Interact models
	// the server can access.
	ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error)
	// Create a new Cobalt Interact session. Also returns a list of
	// actions to take next.
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	// Delete the session. Behavior is undefined if the given
	// TokenData in the request is used again after this function is called.
	DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error)
	// Process input for a session and get an updated session with
	// a list of actions to take next. This is the only method
	// that modifies the Cobalt Interact session state.
	UpdateSession(context.Context, *UpdateSessionRequest) (*UpdateSessionResponse, error)
	// Create an ASR stream. A result is returned when the
	// stream is closed by the client (which forces the ASR to
	// endpoint), or when a transcript becomes available on its
	// own, in which case the stream is closed by the server.
	// The ASR result may be used in the UpdateSession method.
	// <br/><br/>
	// If the session has a wakeword enabled, and the client
	// application is using Cobalt Interact and Cobalt Transcribe
	// to handle the wakeword processing, this method will not
	// return a result until the wakeword condition has been
	// satisfied. Utterances without the required wakeword will
	// be discarded and no transcription will be returned.
	StreamASR(InteractService_StreamASRServer) error
	// Create a TTS stream to receive audio for the given reply.
	// The stream will close when TTS is finished. The client
	// may also close the stream early to cancel the speech
	// synthesis.
	StreamTTS(*StreamTTSRequest, InteractService_StreamTTSServer) error
	// Create an ASR stream for transcription. Unlike StreamASR,
	// Transcribe does not listen for a wakeword. This method
	// returns a bi-directional stream, and its intended use is
	// for situations where a user may say anything at all, whether
	// it is short or long, and the application wants to save the
	// transcript (e.g., take a note, send a message).
	// <br/><br/>
	// The first message sent to the server must be the TranscribeAction,
	// with remaining messages sending audio data.
	// Messages received from the server will include the current
	// best partial transcription until the full transcription is
	// ready. The stream ends when either the client application
	// closes it, a predefined duration of silence (non-speech)
	// occurs, or the end-transcription intent is recognized.
	Transcribe(InteractService_TranscribeServer) error
	// Performs bidirectional streaming speech recognition. Receive results while
	// sending audio. Each result will either be a partial ASR result, or a final
	// result. Partial results will be sent as soon as they are ready, and all
	// results will be sent, regardless of any wakeword configuration in the
	// session. A final result will be sent exactly once, and the stream will be
	// closed then. If a session has a wakeword enabled, the final result will
	// only be emitted if the required wakeword is present. The ASRResult in the
	// final message maybe used in the UpdateSession method for further dialog
	// processing.
	StreamASRWithPartials(InteractService_StreamASRWithPartialsServer) error
	mustEmbedUnimplementedInteractServiceServer()
}

// UnimplementedInteractServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInteractServiceServer struct {
}

func (UnimplementedInteractServiceServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedInteractServiceServer) ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModels not implemented")
}
func (UnimplementedInteractServiceServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedInteractServiceServer) DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedInteractServiceServer) UpdateSession(context.Context, *UpdateSessionRequest) (*UpdateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSession not implemented")
}
func (UnimplementedInteractServiceServer) StreamASR(InteractService_StreamASRServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamASR not implemented")
}
func (UnimplementedInteractServiceServer) StreamTTS(*StreamTTSRequest, InteractService_StreamTTSServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTTS not implemented")
}
func (UnimplementedInteractServiceServer) Transcribe(InteractService_TranscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Transcribe not implemented")
}
func (UnimplementedInteractServiceServer) StreamASRWithPartials(InteractService_StreamASRWithPartialsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamASRWithPartials not implemented")
}
func (UnimplementedInteractServiceServer) mustEmbedUnimplementedInteractServiceServer() {}

// UnsafeInteractServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractServiceServer will
// result in compilation errors.
type UnsafeInteractServiceServer interface {
	mustEmbedUnimplementedInteractServiceServer()
}

func RegisterInteractServiceServer(s grpc.ServiceRegistrar, srv InteractServiceServer) {
	s.RegisterService(&InteractService_ServiceDesc, srv)
}

func _InteractService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractService_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractService_ListModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServiceServer).ListModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractService_ListModels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServiceServer).ListModels(ctx, req.(*ListModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractService_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractService_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServiceServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractService_DeleteSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServiceServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractService_UpdateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServiceServer).UpdateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractService_UpdateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServiceServer).UpdateSession(ctx, req.(*UpdateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractService_StreamASR_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InteractServiceServer).StreamASR(&interactServiceStreamASRServer{stream})
}

type InteractService_StreamASRServer interface {
	SendAndClose(*StreamASRResponse) error
	Recv() (*StreamASRRequest, error)
	grpc.ServerStream
}

type interactServiceStreamASRServer struct {
	grpc.ServerStream
}

func (x *interactServiceStreamASRServer) SendAndClose(m *StreamASRResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *interactServiceStreamASRServer) Recv() (*StreamASRRequest, error) {
	m := new(StreamASRRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _InteractService_StreamTTS_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamTTSRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InteractServiceServer).StreamTTS(m, &interactServiceStreamTTSServer{stream})
}

type InteractService_StreamTTSServer interface {
	Send(*StreamTTSResponse) error
	grpc.ServerStream
}

type interactServiceStreamTTSServer struct {
	grpc.ServerStream
}

func (x *interactServiceStreamTTSServer) Send(m *StreamTTSResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InteractService_Transcribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InteractServiceServer).Transcribe(&interactServiceTranscribeServer{stream})
}

type InteractService_TranscribeServer interface {
	Send(*TranscribeResponse) error
	Recv() (*TranscribeRequest, error)
	grpc.ServerStream
}

type interactServiceTranscribeServer struct {
	grpc.ServerStream
}

func (x *interactServiceTranscribeServer) Send(m *TranscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *interactServiceTranscribeServer) Recv() (*TranscribeRequest, error) {
	m := new(TranscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _InteractService_StreamASRWithPartials_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InteractServiceServer).StreamASRWithPartials(&interactServiceStreamASRWithPartialsServer{stream})
}

type InteractService_StreamASRWithPartialsServer interface {
	Send(*StreamASRWithPartialsResponse) error
	Recv() (*StreamASRWithPartialsRequest, error)
	grpc.ServerStream
}

type interactServiceStreamASRWithPartialsServer struct {
	grpc.ServerStream
}

func (x *interactServiceStreamASRWithPartialsServer) Send(m *StreamASRWithPartialsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *interactServiceStreamASRWithPartialsServer) Recv() (*StreamASRWithPartialsRequest, error) {
	m := new(StreamASRWithPartialsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InteractService_ServiceDesc is the grpc.ServiceDesc for InteractService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cobaltspeech.interact.v3.InteractService",
	HandlerType: (*InteractServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _InteractService_Version_Handler,
		},
		{
			MethodName: "ListModels",
			Handler:    _InteractService_ListModels_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _InteractService_CreateSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _InteractService_DeleteSession_Handler,
		},
		{
			MethodName: "UpdateSession",
			Handler:    _InteractService_UpdateSession_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamASR",
			Handler:       _InteractService_StreamASR_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamTTS",
			Handler:       _InteractService_StreamTTS_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Transcribe",
			Handler:       _InteractService_Transcribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamASRWithPartials",
			Handler:       _InteractService_StreamASRWithPartials_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cobaltspeech/interact/v3/interact.proto",
}
