package client

import (
	"context"
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

func DialContext(ctx context.Context, target string) (*grpc.ClientConn, error) {
	streamInterceptors := []grpc.StreamClientInterceptor{
		grpc_opentracing.StreamClientInterceptor(),
		grpc_prometheus.StreamClientInterceptor,
		grpc_retry.StreamClientInterceptor(),
	}

	unaryInterceptors := []grpc.UnaryClientInterceptor{
		grpc_opentracing.UnaryClientInterceptor(),
		grpc_prometheus.UnaryClientInterceptor,
		grpc_retry.UnaryClientInterceptor(),
	}

	dialOptions := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(streamInterceptors...)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(unaryInterceptors...)),
		grpc.WithBlock(),
	}
	conn, err := grpc.DialContext(ctx, target, dialOptions...)
	if err != nil {
		return nil, err
	}
	go func() {
		<-ctx.Done()
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	return conn, nil
}
