package interceptors

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func recoveryInterceptor(in *interceptor, logger zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				// log error details and stack trace
				methodName := getMethod(info)
				err = fmt.Errorf("%v in call to method '%s'", r, methodName)
				logger.Err(err).Str("method", methodName).Str("stacktrace", string(debug.Stack())).Msg("failed to handle the request [PANIC]")
			}
		}()

		resp, err := handler(ctx, req)
		panicked = false
		return resp, err
	}
}
