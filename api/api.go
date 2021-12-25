package api

import (
	"context"

	"github.com/NpoolPlatform/internalization/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedInternalizationServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterInternalizationServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterInternalizationHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
