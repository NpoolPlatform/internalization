// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	crud "github.com/NpoolPlatform/internationalization/pkg/crud/country"
	npool "github.com/NpoolPlatform/message/npool/internationalization"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCountry(ctx context.Context, in *npool.CreateCountryRequest) (*npool.CreateCountryResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create country: %v", err)
		return &npool.CreateCountryResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateCountries(ctx context.Context, in *npool.CreateCountriesRequest) (*npool.CreateCountriesResponse, error) {
	resp, err := crud.CreateCountries(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create countries: %v", err)
		return &npool.CreateCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateCountry(ctx context.Context, in *npool.UpdateCountryRequest) (*npool.UpdateCountryResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update country: %v", err)
		return &npool.UpdateCountryResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCountry(ctx context.Context, in *npool.GetCountryRequest) (*npool.GetCountryResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get country: %v", err)
		return &npool.GetCountryResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCountries(ctx context.Context, in *npool.GetCountriesRequest) (*npool.GetCountriesResponse, error) {
	resp, err := crud.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get countries: %v", err)
		return &npool.GetCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
