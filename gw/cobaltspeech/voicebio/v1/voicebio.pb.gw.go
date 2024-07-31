// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: cobaltspeech/voicebio/v1/voicebio.proto

/*
Package voicebiov1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package voicebiov1

import (
	"context"
	"io"
	"net/http"

	extVoicebiov1 "github.com/cobaltspeech/go-genproto/cobaltspeech/voicebio/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_VoiceBioService_Version_0(ctx context.Context, marshaler runtime.Marshaler, client extVoicebiov1.VoiceBioServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extVoicebiov1.VersionRequest
	var metadata runtime.ServerMetadata

	msg, err := client.Version(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_VoiceBioService_Version_0(ctx context.Context, marshaler runtime.Marshaler, server extVoicebiov1.VoiceBioServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extVoicebiov1.VersionRequest
	var metadata runtime.ServerMetadata

	msg, err := server.Version(ctx, &protoReq)
	return msg, metadata, err

}

func request_VoiceBioService_ListModels_0(ctx context.Context, marshaler runtime.Marshaler, client extVoicebiov1.VoiceBioServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extVoicebiov1.ListModelsRequest
	var metadata runtime.ServerMetadata

	msg, err := client.ListModels(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_VoiceBioService_ListModels_0(ctx context.Context, marshaler runtime.Marshaler, server extVoicebiov1.VoiceBioServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extVoicebiov1.ListModelsRequest
	var metadata runtime.ServerMetadata

	msg, err := server.ListModels(ctx, &protoReq)
	return msg, metadata, err

}

func request_VoiceBioService_StreamingEnroll_0(ctx context.Context, marshaler runtime.Marshaler, client extVoicebiov1.VoiceBioServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.StreamingEnroll(ctx)
	if err != nil {
		grpclog.Errorf("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	for {
		var protoReq extVoicebiov1.StreamingEnrollRequest
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Errorf("Failed to decode request: %v", err)
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		if err = stream.Send(&protoReq); err != nil {
			if err == io.EOF {
				break
			}
			grpclog.Errorf("Failed to send request: %v", err)
			return nil, metadata, err
		}
	}

	if err := stream.CloseSend(); err != nil {
		grpclog.Errorf("Failed to terminate client stream: %v", err)
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		grpclog.Errorf("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header

	msg, err := stream.CloseAndRecv()
	metadata.TrailerMD = stream.Trailer()
	return msg, metadata, err

}

func request_VoiceBioService_StreamingVerify_0(ctx context.Context, marshaler runtime.Marshaler, client extVoicebiov1.VoiceBioServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.StreamingVerify(ctx)
	if err != nil {
		grpclog.Errorf("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	for {
		var protoReq extVoicebiov1.StreamingVerifyRequest
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Errorf("Failed to decode request: %v", err)
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		if err = stream.Send(&protoReq); err != nil {
			if err == io.EOF {
				break
			}
			grpclog.Errorf("Failed to send request: %v", err)
			return nil, metadata, err
		}
	}

	if err := stream.CloseSend(); err != nil {
		grpclog.Errorf("Failed to terminate client stream: %v", err)
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		grpclog.Errorf("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header

	msg, err := stream.CloseAndRecv()
	metadata.TrailerMD = stream.Trailer()
	return msg, metadata, err

}

func request_VoiceBioService_StreamingIdentify_0(ctx context.Context, marshaler runtime.Marshaler, client extVoicebiov1.VoiceBioServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.StreamingIdentify(ctx)
	if err != nil {
		grpclog.Errorf("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	for {
		var protoReq extVoicebiov1.StreamingIdentifyRequest
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Errorf("Failed to decode request: %v", err)
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		if err = stream.Send(&protoReq); err != nil {
			if err == io.EOF {
				break
			}
			grpclog.Errorf("Failed to send request: %v", err)
			return nil, metadata, err
		}
	}

	if err := stream.CloseSend(); err != nil {
		grpclog.Errorf("Failed to terminate client stream: %v", err)
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		grpclog.Errorf("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header

	msg, err := stream.CloseAndRecv()
	metadata.TrailerMD = stream.Trailer()
	return msg, metadata, err

}

// RegisterVoiceBioServiceHandlerServer registers the http handlers for service VoiceBioService to "mux".
// UnaryRPC     :call VoiceBioServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterVoiceBioServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterVoiceBioServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server extVoicebiov1.VoiceBioServiceServer) error {

	mux.Handle("GET", pattern_VoiceBioService_Version_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/cobaltspeech.voicebio.v1.VoiceBioService/Version", runtime.WithHTTPPathPattern("/api/voicebio/v1/version"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_VoiceBioService_Version_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VoiceBioService_Version_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_VoiceBioService_ListModels_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/cobaltspeech.voicebio.v1.VoiceBioService/ListModels", runtime.WithHTTPPathPattern("/api/voicebio/v1/list-models"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_VoiceBioService_ListModels_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VoiceBioService_ListModels_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_VoiceBioService_StreamingEnroll_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("GET", pattern_VoiceBioService_StreamingVerify_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("GET", pattern_VoiceBioService_StreamingIdentify_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterVoiceBioServiceHandlerFromEndpoint is same as RegisterVoiceBioServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterVoiceBioServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterVoiceBioServiceHandler(ctx, mux, conn)
}

// RegisterVoiceBioServiceHandler registers the http handlers for service VoiceBioService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterVoiceBioServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterVoiceBioServiceHandlerClient(ctx, mux, extVoicebiov1.NewVoiceBioServiceClient(conn))
}

// RegisterVoiceBioServiceHandlerClient registers the http handlers for service VoiceBioService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "extVoicebiov1.VoiceBioServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "extVoicebiov1.VoiceBioServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "extVoicebiov1.VoiceBioServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterVoiceBioServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client extVoicebiov1.VoiceBioServiceClient) error {

	mux.Handle("GET", pattern_VoiceBioService_Version_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/cobaltspeech.voicebio.v1.VoiceBioService/Version", runtime.WithHTTPPathPattern("/api/voicebio/v1/version"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_VoiceBioService_Version_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VoiceBioService_Version_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_VoiceBioService_ListModels_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/cobaltspeech.voicebio.v1.VoiceBioService/ListModels", runtime.WithHTTPPathPattern("/api/voicebio/v1/list-models"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_VoiceBioService_ListModels_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VoiceBioService_ListModels_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_VoiceBioService_StreamingEnroll_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/cobaltspeech.voicebio.v1.VoiceBioService/StreamingEnroll", runtime.WithHTTPPathPattern("/api/voicebio/v1/streaming-enroll"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_VoiceBioService_StreamingEnroll_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VoiceBioService_StreamingEnroll_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_VoiceBioService_StreamingVerify_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/cobaltspeech.voicebio.v1.VoiceBioService/StreamingVerify", runtime.WithHTTPPathPattern("/api/voicebio/v1/streaming-verify"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_VoiceBioService_StreamingVerify_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VoiceBioService_StreamingVerify_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_VoiceBioService_StreamingIdentify_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/cobaltspeech.voicebio.v1.VoiceBioService/StreamingIdentify", runtime.WithHTTPPathPattern("/api/voicebio/v1/streaming-identify"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_VoiceBioService_StreamingIdentify_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VoiceBioService_StreamingIdentify_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_VoiceBioService_Version_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "voicebio", "v1", "version"}, ""))

	pattern_VoiceBioService_ListModels_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "voicebio", "v1", "list-models"}, ""))

	pattern_VoiceBioService_StreamingEnroll_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "voicebio", "v1", "streaming-enroll"}, ""))

	pattern_VoiceBioService_StreamingVerify_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "voicebio", "v1", "streaming-verify"}, ""))

	pattern_VoiceBioService_StreamingIdentify_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "voicebio", "v1", "streaming-identify"}, ""))
)

var (
	forward_VoiceBioService_Version_0 = runtime.ForwardResponseMessage

	forward_VoiceBioService_ListModels_0 = runtime.ForwardResponseMessage

	forward_VoiceBioService_StreamingEnroll_0 = runtime.ForwardResponseMessage

	forward_VoiceBioService_StreamingVerify_0 = runtime.ForwardResponseMessage

	forward_VoiceBioService_StreamingIdentify_0 = runtime.ForwardResponseMessage
)
