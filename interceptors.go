package interceptors

import (
	"strings"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func getMethod(info *grpc.UnaryServerInfo) string {
	splits := strings.Split(info.FullMethod, "/")
	return splits[len(splits)-1]
}

type Interceptor interface {
	Get() grpc.ServerOption
}

type interceptor struct {
	options     []grpc.UnaryServerInterceptor
	skipMethods map[string]struct{}
}

func (i *interceptor) skipLog(method string) bool {
	// quick good lookup function
	_, ok := i.skipMethods[method]
	return ok
}

type InterceptorOption func(in *interceptor)

// WithInterceptor appends a user defined interceptor to the chain
func WithInterceptor(userInterceptor grpc.UnaryServerInterceptor) InterceptorOption {
	return func(in *interceptor) {
		in.options = append(in.options, userInterceptor)
	}
}

// Get return the unary interceptors composing of all default and user options
func (in *interceptor) Get() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		in.options...,
	)
}

func WithSkipMethod(methods []string) InterceptorOption {
	return func(in *interceptor) {
		in.skipMethods = make(map[string]struct{}, len(methods))
		for _, m := range methods {
			in.skipMethods[m] = struct{}{}
		}
	}
}

func NewInterceptor(service string, logger zerolog.Logger, options ...InterceptorOption) Interceptor {
	in := &interceptor{}

	// apply default interceptors
	WithInterceptor(kitgrpc.Interceptor)(in)
	WithInterceptor(traceIdReader(in, logger))(in)
	WithInterceptor(loggingInterceptor(in, logger))(in)
	WithInterceptor(recoveryInterceptor(in, logger))(in)

	for _, option := range options {
		option(in)
	}
	return in
}
