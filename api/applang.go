package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	crud "github.com/NpoolPlatform/internationalization/pkg/crud/applang"
	mw "github.com/NpoolPlatform/internationalization/pkg/middleware/applang"
	npool "github.com/NpoolPlatform/message/npool/internationalization"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppLang(ctx context.Context, in *npool.CreateAppLangRequest) (*npool.CreateAppLangResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create lang: %v", err)
		return &npool.CreateAppLangResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppLangForOtherApp(ctx context.Context, in *npool.CreateAppLangForOtherAppRequest) (*npool.CreateAppLangForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := crud.Create(ctx, &npool.CreateAppLangRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorw("fail create lang: %v", err)
		return &npool.CreateAppLangForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateAppLangForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UpdateAppLang(ctx context.Context, in *npool.UpdateAppLangRequest) (*npool.UpdateAppLangResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update lang: %v", err)
		return &npool.UpdateAppLangResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppLang(ctx context.Context, in *npool.GetAppLangRequest) (*npool.GetAppLangResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get lang: %v", err)
		return &npool.GetAppLangResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppLangsByApp(ctx context.Context, in *npool.GetAppLangsByAppRequest) (*npool.GetAppLangsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app langs: %v", err)
		return &npool.GetAppLangsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppLangInfosByApp(ctx context.Context, in *npool.GetAppLangInfosByAppRequest) (*npool.GetAppLangInfosByAppResponse, error) {
	resp, err := mw.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app lang infos: %v", err)
		return &npool.GetAppLangInfosByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppLangInfosByOtherApp(ctx context.Context, in *npool.GetAppLangInfosByOtherAppRequest) (*npool.GetAppLangInfosByOtherAppResponse, error) {
	resp, err := mw.GetByApp(ctx, &npool.GetAppLangInfosByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorw("fail get other app lang infos: %v", err)
		return &npool.GetAppLangInfosByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppLangInfosByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}
