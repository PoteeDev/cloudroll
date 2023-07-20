package gateway

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	pb "github.com/PoteeDev/cloudroll/proto"
	"github.com/PoteeDev/cloudroll/src/interceptors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

var (
	callbackPath = "/auth/callback"
	key          = []byte("test1234test1234")
	clientID     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	keyPath      = os.Getenv("KEY_PATH")
	issuer       = os.Getenv("ISSUER")
	port         = os.Getenv("PORT")
	scopes       = strings.Split(os.Getenv("SCOPES"), " ")
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

	withCors := cors.New(cors.Options{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(mux)

	s := &http.Server{
		Addr:    addr,
		Handler: withCors,
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
