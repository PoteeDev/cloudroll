package interceptors

import (
	"context"
	"log"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/zitadel/oidc/v2/pkg/client/rs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(provider rs.ResourceServer) grpc.UnaryServerInterceptor {
	authFn := func(ctx context.Context) (context.Context, error) {
		token, err := auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}
		resp, err := rs.Introspect(ctx, provider, token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid auth token")
		}
		log.Println(resp)

		// if !ok || value == "" || value != requestedValue {
		// 	return nil, status.Error(codes.Unauthenticated, "invalid auth token")
		// }
		// NOTE: You can also pass the token in the context for further interceptors or gRPC service code.
		return ctx, nil
	}

	// Setup auth matcher.
	allButHealthZ := func(ctx context.Context, callMeta interceptors.CallMeta) bool {
		return healthpb.Health_ServiceDesc.ServiceName != callMeta.Service
	}
	return selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(authFn), selector.MatchFunc(allButHealthZ))
}
