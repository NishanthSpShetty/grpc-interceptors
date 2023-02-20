package interceptors

import (
	"context"
	"runtime/debug"

	"github.com/gogo/status"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func recoveryInterceptor(in *interceptor, logger zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				// log error details and stack trace
				methodName := getMethod(info)
				logger.Err(err).Str("method", methodName).Str("stacktrace", string(debug.Stack())).Msg("failed to handle the request [PANIC]")
				err = status.Errorf(codes.Internal, "%v in call to method '%s'", r, methodName)
			}
		}()

		resp, err := handler(ctx, req)
		panicked = false
		return resp, err
	}
}
