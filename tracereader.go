package interceptors

import (
	"context"

	"github.com/NishanthSpShetty/grpc-interceptors/client"
	"github.com/NishanthSpShetty/log/loggers"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func traceIdReader(in *interceptor, logger zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)

		if ok {
			v := md.Get(client.TRACE_ID)
			if len(v) == 1 {
				ctx = loggers.AddToLogContext(ctx, client.TRACE_ID, v[0])
			}
		}
		return handler(ctx, req)
	}
}
