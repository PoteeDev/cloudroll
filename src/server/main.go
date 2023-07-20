package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	proto "github.com/PoteeDev/cloudroll/proto"
	"github.com/PoteeDev/cloudroll/src/interceptors"
	grpcmw "github.com/zitadel/zitadel-go/v2/pkg/api/middleware/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	keyPath = os.Getenv("KEY")
	issuer  = os.Getenv("ISSUER")
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
	// shutdown opentelemetry exporter
	defer func() { _ = interceptors.Exporter.Shutdown(context.Background()) }()

	authIntrospector, err := grpcmw.NewIntrospectionInterceptor(issuer, keyPath)
	if err != nil {
		log.Fatalln("oauth error:", err.Error())
	}
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.OpenTelemetryInterceptor(),
			interceptors.PrometheusIntercepotor(),
			authIntrospector.Unary(),
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
