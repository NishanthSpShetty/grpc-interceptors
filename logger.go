package interceptors

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func loggingInterceptor(in *interceptor, logger zerolog.Logger) grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		// log request and response data

		begin := time.Now()
		method := getMethod(info)
		logger = logger.With().Str("method", method).Logger()

		skip := in.skipLog(method)

		if !skip {
			request := fmt.Sprintf("%+v", req)
			logger = logger.With().Str("request", request).Logger()
		}

		logger.Info().Send()
		resp, err := handler(ctx, req)

		logger = logger.With().Dur("took", time.Since(begin)).Logger()
		if !skip {
			logger = logger.With().Interface("response", resp).Logger()
		}

		if err != nil {
			logger.Err(err).Send()
		} else {
			logger.Info().Send()
		}

		return resp, err
	}
}
