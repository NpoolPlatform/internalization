// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	crud "github.com/NpoolPlatform/internationalization/pkg/crud/message"
	npool "github.com/NpoolPlatform/message/npool/internationalization"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) CreateMessage(ctx context.Context, in *npool.CreateMessageRequest) (*npool.CreateMessageResponse, error) {
	info := in.GetInfo()
	if _, err := uuid.Parse(in.GetTargetLangID()); err == nil {
		info.LangID = in.GetTargetLangID()
	}

	resp, err := crud.CreateMessage(ctx, &npool.CreateMessageRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("fail create message: %v", err)
		return &npool.CreateMessageResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateMessageForOtherApp(ctx context.Context, in *npool.CreateMessageForOtherAppRequest) (*npool.CreateMessageForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := crud.CreateMessage(ctx, &npool.CreateMessageRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("fail create message: %v", err)
		return &npool.CreateMessageForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateMessageForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) CreateMessages(ctx context.Context, in *npool.CreateMessagesRequest) (*npool.CreateMessagesResponse, error) {
	resp, err := crud.CreateMessages(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create messages: %v", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateMessage(ctx context.Context, in *npool.UpdateMessageRequest) (*npool.UpdateMessageResponse, error) {
	resp, err := crud.UpdateMessage(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update message: %v", err)
		return &npool.UpdateMessageResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateMessages(ctx context.Context, in *npool.UpdateMessagesRequest) (*npool.UpdateMessagesResponse, error) {
	resp, err := crud.UpdateMessages(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update messages: %v", err)
		return &npool.UpdateMessagesResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetMessagesByAppLang(ctx context.Context, in *npool.GetMessagesByAppLangRequest) (*npool.GetMessagesByAppLangResponse, error) {
	resp, err := crud.GetMessagesByAppLang(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get messages by lang id: %v", err)
		return &npool.GetMessagesByAppLangResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetMessagesByOtherAppLang(ctx context.Context, in *npool.GetMessagesByOtherAppLangRequest) (*npool.GetMessagesByOtherAppLangResponse, error) {
	resp, err := crud.GetMessagesByAppLang(ctx, &npool.GetMessagesByAppLangRequest{
		AppID:  in.GetTargetAppID(),
		LangID: in.GetLangID(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail get messages by lang id: %v", err)
		return &npool.GetMessagesByOtherAppLangResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetMessagesByOtherAppLangResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetMessageByAppLangMessage(ctx context.Context, in *npool.GetMessageByAppLangMessageRequest) (*npool.GetMessageByAppLangMessageResponse, error) {
	resp, err := crud.GetMessageByAppLangMessage(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get message by lang id message id: %v", err)
		return &npool.GetMessageByAppLangMessageResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
