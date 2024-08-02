// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: cobaltspeech/cubic/v5/cubic.proto

/*
Package cubicv5 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package cubicv5

import (
	"context"
	"io"
	"net/http"

	extCubicv5 "github.com/cobaltspeech/go-genproto/cobaltspeech/cubic/v5"
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

func request_CubicService_Version_0(ctx context.Context, marshaler runtime.Marshaler, client extCubicv5.CubicServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extCubicv5.VersionRequest
	var metadata runtime.ServerMetadata

	msg, err := client.Version(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CubicService_Version_0(ctx context.Context, marshaler runtime.Marshaler, server extCubicv5.CubicServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extCubicv5.VersionRequest
	var metadata runtime.ServerMetadata

	msg, err := server.Version(ctx, &protoReq)
	return msg, metadata, err

}

func request_CubicService_ListModels_0(ctx context.Context, marshaler runtime.Marshaler, client extCubicv5.CubicServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extCubicv5.ListModelsRequest
	var metadata runtime.ServerMetadata

	msg, err := client.ListModels(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CubicService_ListModels_0(ctx context.Context, marshaler runtime.Marshaler, server extCubicv5.CubicServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extCubicv5.ListModelsRequest
	var metadata runtime.ServerMetadata

	msg, err := server.ListModels(ctx, &protoReq)
	return msg, metadata, err

}

func request_CubicService_StreamingRecognize_0(ctx context.Context, marshaler runtime.Marshaler, client extCubicv5.CubicServiceClient, req *http.Request, pathParams map[string]string) (extCubicv5.CubicService_StreamingRecognizeClient, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.StreamingRecognize(ctx)
	if err != nil {
		grpclog.Errorf("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq extCubicv5.StreamingRecognizeRequest
		err := dec.Decode(&protoReq)
		if err == io.EOF {
			return err
		}
		if err != nil {
			grpclog.Errorf("Failed to decode request: %v", err)
			return err
		}
		if err := stream.Send(&protoReq); err != nil {
			grpclog.Errorf("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	go func() {
		for {
			if err := handleSend(); err != nil {
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Errorf("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Errorf("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}

func request_CubicService_CompileContext_0(ctx context.Context, marshaler runtime.Marshaler, client extCubicv5.CubicServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extCubicv5.CompileContextRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CompileContext(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CubicService_CompileContext_0(ctx context.Context, marshaler runtime.Marshaler, server extCubicv5.CubicServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extCubicv5.CompileContextRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CompileContext(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterCubicServiceHandlerServer registers the http handlers for service CubicService to "mux".
// UnaryRPC     :call CubicServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterCubicServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterCubicServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server extCubicv5.CubicServiceServer) error {

	mux.Handle("GET", pattern_CubicService_Version_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/cobaltspeech.cubic.v5.CubicService/Version", runtime.WithHTTPPathPattern("/api/cubic/v5/version"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CubicService_Version_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CubicService_Version_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_CubicService_ListModels_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/cobaltspeech.cubic.v5.CubicService/ListModels", runtime.WithHTTPPathPattern("/api/cubic/v5/list-models"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CubicService_ListModels_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CubicService_ListModels_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CubicService_StreamingRecognize_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("POST", pattern_CubicService_CompileContext_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/cobaltspeech.cubic.v5.CubicService/CompileContext", runtime.WithHTTPPathPattern("/api/cubic/v5/compile-context"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CubicService_CompileContext_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CubicService_CompileContext_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterCubicServiceHandlerFromEndpoint is same as RegisterCubicServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCubicServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterCubicServiceHandler(ctx, mux, conn)
}

// RegisterCubicServiceHandler registers the http handlers for service CubicService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCubicServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCubicServiceHandlerClient(ctx, mux, extCubicv5.NewCubicServiceClient(conn))
}

// RegisterCubicServiceHandlerClient registers the http handlers for service CubicService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "extCubicv5.CubicServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "extCubicv5.CubicServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "extCubicv5.CubicServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterCubicServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client extCubicv5.CubicServiceClient) error {

	mux.Handle("GET", pattern_CubicService_Version_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/cobaltspeech.cubic.v5.CubicService/Version", runtime.WithHTTPPathPattern("/api/cubic/v5/version"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CubicService_Version_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CubicService_Version_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_CubicService_ListModels_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/cobaltspeech.cubic.v5.CubicService/ListModels", runtime.WithHTTPPathPattern("/api/cubic/v5/list-models"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CubicService_ListModels_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CubicService_ListModels_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CubicService_StreamingRecognize_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/cobaltspeech.cubic.v5.CubicService/StreamingRecognize", runtime.WithHTTPPathPattern("/api/cubic/v5/streaming-recognize"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CubicService_StreamingRecognize_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CubicService_StreamingRecognize_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CubicService_CompileContext_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/cobaltspeech.cubic.v5.CubicService/CompileContext", runtime.WithHTTPPathPattern("/api/cubic/v5/compile-context"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CubicService_CompileContext_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CubicService_CompileContext_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CubicService_Version_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "cubic", "v5", "version"}, ""))

	pattern_CubicService_ListModels_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "cubic", "v5", "list-models"}, ""))

	pattern_CubicService_StreamingRecognize_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "cubic", "v5", "streaming-recognize"}, ""))

	pattern_CubicService_CompileContext_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "cubic", "v5", "compile-context"}, ""))
)

var (
	forward_CubicService_Version_0 = runtime.ForwardResponseMessage

	forward_CubicService_ListModels_0 = runtime.ForwardResponseMessage

	forward_CubicService_StreamingRecognize_0 = runtime.ForwardResponseStream

	forward_CubicService_CompileContext_0 = runtime.ForwardResponseMessage
)
