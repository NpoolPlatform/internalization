// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	crud "github.com/NpoolPlatform/internationalization/pkg/crud/message"
	npool "github.com/NpoolPlatform/message/npool/internationalization"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateMessage(ctx context.Context, in *npool.CreateMessageRequest) (*npool.CreateMessageResponse, error) {
	resp, err := crud.CreateMessage(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create message: %v", err)
		return &npool.CreateMessageResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) CreateMessages(ctx context.Context, in *npool.CreateMessagesRequest) (*npool.CreateMessagesResponse, error) {
	resp, err := crud.CreateMessages(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create messages: %v", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateMessage(ctx context.Context, in *npool.UpdateMessageRequest) (*npool.UpdateMessageResponse, error) {
	resp, err := crud.UpdateMessage(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update message: %v", err)
		return &npool.UpdateMessageResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateMessages(ctx context.Context, in *npool.UpdateMessagesRequest) (*npool.UpdateMessagesResponse, error) {
	resp, err := crud.UpdateMessages(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update messages: %v", err)
		return &npool.UpdateMessagesResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetMessagesByLangID(ctx context.Context, in *npool.GetMessagesByLangIDRequest) (*npool.GetMessagesByLangIDResponse, error) {
	resp, err := crud.GetMessagesByLangID(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get messages by lang id: %v", err)
		return &npool.GetMessagesByLangIDResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetMessageByLangIDMessageID(ctx context.Context, in *npool.GetMessageByLangIDMessageIDRequest) (*npool.GetMessageByLangIDMessageIDResponse, error) {
	resp, err := crud.GetMessageByLangIDMessageID(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get message by lang id message id: %v", err)
		return &npool.GetMessageByLangIDMessageIDResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
