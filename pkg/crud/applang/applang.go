package applang

import (
	"context"
	"time"

	"github.com/NpoolPlatform/internationalization/pkg/db"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/applang"
	npool "github.com/NpoolPlatform/message/npool/internationalization"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

const (
	dbTimeout = 30 * time.Second
)

func validateAppLang(info *npool.AppLang) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetLangID()); err != nil {
		return xerrors.Errorf("invalid lang id: %v", err)
	}
	return nil
}

func dbRowToLang(row *ent.AppLang) *npool.AppLang {
	return &npool.AppLang{
		ID:       row.ID.String(),
		AppID:    row.AppID.String(),
		LangID:   row.LangID.String(),
		MainLang: row.MainLang,
	}
}

func Create(ctx context.Context, in *npool.CreateAppLangRequest) (*npool.CreateAppLangResponse, error) {
	if err := validateAppLang(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		AppLang.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetLangID(uuid.MustParse(in.GetInfo().GetLangID())).
		SetMainLang(in.GetInfo().GetMainLang()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create lang: %v", err)
	}

	return &npool.CreateAppLangResponse{
		Info: dbRowToLang(info),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppLangsByAppRequest) (*npool.GetAppLangsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		AppLang.
		Query().
		Where(
			applang.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query langs: %v", err)
	}

	langs := []*npool.AppLang{}
	for _, info := range infos {
		langs = append(langs, dbRowToLang(info))
	}

	return &npool.GetAppLangsByAppResponse{
		Infos: langs,
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppLangRequest) (*npool.GetAppLangResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app lang id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		AppLang.
		Query().
		Where(
			applang.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query langs: %v", err)
	}

	var lang *npool.AppLang
	for _, info := range infos {
		lang = dbRowToLang(info)
		break
	}

	return &npool.GetAppLangResponse{
		Info: lang,
	}, nil
}
