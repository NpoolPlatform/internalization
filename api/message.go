// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/internationalization/message/npool"
	crud "github.com/NpoolPlatform/internationalization/pkg/crud/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateMessage(ctx context.Context, in *npool.CreateMessageRequest) (*npool.CreateMessageResponse, error) {
	resp, err := crud.CreateMessage(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create message: %v", err)
		return &npool.CreateMessageResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func CreateMessages(ctx context.Context, in *npool.CreateMessagesRequest) (*npool.CreateMessagesResponse, error) {
	resp, err := crud.CreateMessages(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create messages: %v", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func UpdateMessage(ctx context.Context, in *npool.UpdateMessageRequest) (*npool.UpdateMessageResponse, error) {
	resp, err := crud.UpdateMessage(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update message: %v", err)
		return &npool.UpdateMessageResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func UpdateMessages(ctx context.Context, in *npool.UpdateMessagesRequest) (*npool.UpdateMessagesResponse, error) {
	resp, err := crud.UpdateMessages(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update messages: %v", err)
		return &npool.UpdateMessagesResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func GetMessages(ctx context.Context, in *npool.GetMessagesRequest) (*npool.GetMessagesResponse, error) {
	resp, err := crud.GetMessages(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get messages: %v", err)
		return &npool.GetMessagesResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func GetMessageByMessageID(ctx context.Context, in *npool.GetMessageByMessageIDRequest) (*npool.GetMessageByMessageIDResponse, error) {
	resp, err := crud.GetMessageByMessageID(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get message by message id: %v", err)
		return &npool.GetMessageByMessageIDResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
