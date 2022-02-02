package applang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	crud "github.com/NpoolPlatform/internationalization/pkg/crud/applang"
	langcrud "github.com/NpoolPlatform/internationalization/pkg/crud/lang"
	npool "github.com/NpoolPlatform/message/npool/internationalization"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func GetByApp(ctx context.Context, in *npool.GetAppLangInfosByAppRequest) (*npool.GetAppLangInfosByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	resp, err := crud.GetByApp(ctx, &npool.GetAppLangsByAppRequest{
		AppID: appID.String(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get langs by app: %v", err)
	}

	infos := []*npool.AppLangInfo{}
	for _, info := range resp.Infos {
		resp1, err := langcrud.Get(ctx, &npool.GetLangRequest{
			ID: info.LangID,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail get lang: %v", err)
		}
		if resp1.Info == nil {
			logger.Sugar().Warnf("invalid lang")
			continue
		}
		infos = append(infos, &npool.AppLangInfo{
			AppLang: info,
			Lang:    resp1.Info,
		})
	}

	return &npool.GetAppLangInfosByAppResponse{
		Infos: infos,
	}, nil
}
