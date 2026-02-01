# gRPC Interceptors

A collection of gRPC interceptors for Go applications.

This package provides the following interceptors:

*   **Kit Interceptor**: Injects the method name into the context.
*   **Logging Interceptor**: Logs request and response data, along with the duration of the call.
*   **Recovery Interceptor**: Gracefully recovers from API panics and logs the error with a stack trace.
*   **Trace ID Reader**: Reads trace IDs from metadata for request tracking.

## Installation

```bash
go get -u github.com/NishanthSpShetty/grpc-interceptors
```

example
```
    go get -u github.com/NishanthSpShetty/grpc-interceptors@v0.4.1
```

## Usage

Register interceptors when setting up gRPC server in application

### Using default interceptors

```
//get Interceptor interface
interceptor := interceptors.NewInterceptor("myservice", logger)

//Calling Get will return ServerOptions
baseServer := grpc.NewServer(interceptor.Get())
```

### Using user defined interceptors
```
func InstrumentationInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		IncRequest(getMethod(info))
		return handler(ctx, req)
	}
}

//get Interceptor interface
interceptor := interceptors.NewInterceptor(
				"myservice",
 				logger,
                  		interceptors.WithInterceptor(InstrumentationInterceptor()))

//Calling Get will return ServerOptions, which is passed to NewServer
baseServer := grpc.NewServer(interceptor.Get())
```

### skip sensitive methods from logging interceptor

You can skip sensitive API request and response from request-response logging interceptor.
```
//get Interceptor interface
var skipMethods = []string{"ReadCredential", "WriteCredential"}

interceptor := interceptors.NewInterceptor(
				"myservice",
 				logger,
				interceptors.WithSkipMethod(skipMethods))

//Calling Get will return ServerOptions, which is passed to NewServer
baseServer := grpc.NewServer(interceptor.Get())
```

## Support
Raise a issue, we will get back to you.

## Roadmap

## Contributing

## Authors and acknowledgment

* Nishanth Shetty <nishanthspshetty@gmail.com>

## License

   GNU GENERAL PUBLIC LICENSE


## Project status
    
    
