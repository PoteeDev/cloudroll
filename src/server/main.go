package server

import (
	"context"
	"fmt"
	"net"

	proto "github.com/PoteeDev/cloudroll/proto"
	"github.com/PoteeDev/cloudroll/src/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func Run(ctx context.Context, network, address string) error {

	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			fmt.Printf("Failed to close %s %s: %v\n", network, address, err)
		}
	}()

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.PrometheusIntercepotor(),
			interceptors.AuthInterceptor(),
		),
	)
	interceptors.SrvMetrics.InitializeMetrics(s)
	// proto.RegisterEchoServiceServer(s, newEchoServer())
	proto.RegisterCloudrollServiceServer(s, newCloudrollServer())
	reflection.Register(s)

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}
