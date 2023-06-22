package gateway

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/PoteeDev/cloudroll/proto"
	"github.com/PoteeDev/cloudroll/src/interceptors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

// Run starts the example http gateway.
// "endpoint" is passed to register the gRPC server.
// "addr" is passed to net.Listen.
func Run(ctx context.Context, addr string, endpoint string, opts ...runtime.ServeMuxOption) error {
	mux := runtime.NewServeMux(opts...)
	reg := prometheus.NewRegistry()
	reg.MustRegister(interceptors.SrvMetrics)
	mux.HandlePath("GET", "/metrics", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		promhttp.HandlerFor(
			reg,
			promhttp.HandlerOpts{
				// Opt into OpenMetrics e.g. to support exemplars.
				EnableOpenMetrics: true,
			},
		).ServeHTTP(w, r)
	})
	grpcOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterCloudrollServiceHandlerFromEndpoint(ctx, mux, endpoint, grpcOpts)
	if err != nil {
		fmt.Printf("Failed to register endpoint server: %v", err)
		return err
	}

	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			fmt.Printf("Failed to shutdown http gateway server: %v", err)
		}
	}()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Failed to listen and serve: %v", err)
		return err
	}

	return nil
}
