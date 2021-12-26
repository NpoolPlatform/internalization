// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/internationalization/message/npool"
	crud "github.com/NpoolPlatform/internationalization/pkg/crud/lang"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Add(ctx context.Context, in *npool.AddLangRequest) (*npool.AddLangResponse, error) {
	resp, err := crud.Add(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail add lang: %v", err)
		return &npool.AddLangResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) Update(ctx context.Context, in *npool.UpdateLangRequest) (*npool.UpdateLangResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update lang: %v", err)
		return &npool.UpdateLangResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
