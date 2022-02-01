package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/internationalization"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedInternationalizationServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterInternationalizationServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterInternationalizationHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
