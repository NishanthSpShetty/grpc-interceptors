package interceptors

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func loggingInterceptor(in *interceptor, logger zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		// log request and response data

		begin := time.Now()
		method := getMethod(info)
		log := logger.With().Str("method", method).Logger()

		skip := in.skipLog(method)

		if !skip {
			log = log.With().Interface("request", req).Logger()
		}

		log.Debug().Send()
		resp, err := handler(ctx, req)

		log = log.With().Dur("took", time.Since(begin)).Logger()
		if !skip {
			log = log.With().Interface("response", resp).Logger()
		}

		if err != nil {
			log.Err(err).Send()
		} else {
			log.Info().Send()
		}

		return resp, err
	}
}
