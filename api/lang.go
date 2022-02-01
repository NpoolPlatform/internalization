// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	crud "github.com/NpoolPlatform/internationalization/pkg/crud/lang"
	npool "github.com/NpoolPlatform/message/npool/internationalization"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddLang(ctx context.Context, in *npool.AddLangRequest) (*npool.AddLangResponse, error) {
	resp, err := crud.Add(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail add lang: %v", err)
		return &npool.AddLangResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetLang(ctx context.Context, in *npool.GetLangRequest) (*npool.GetLangResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get lang: %v", err)
		return &npool.GetLangResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateLang(ctx context.Context, in *npool.UpdateLangRequest) (*npool.UpdateLangResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update lang: %v", err)
		return &npool.UpdateLangResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetLangs(ctx context.Context, in *npool.GetLangsRequest) (*npool.GetLangsResponse, error) {
	resp, err := crud.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get langs: %v", err)
		return &npool.GetLangsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
