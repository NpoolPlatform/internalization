package lang

import (
	"context"
	"time"

	"github.com/NpoolPlatform/internationalization/message/npool"
	"github.com/NpoolPlatform/internationalization/pkg/db"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent"

	"golang.org/x/xerrors"
)

const (
	dbTimeout = 30 * time.Second
)

func validateLang(info *npool.Lang) error {
	if info.Lang == "" {
		return xerrors.Errorf("invalid lang")
	}
	if info.Name == "" {
		return xerrors.Errorf("invalid lang name")
	}
	if info.Logo == "" {
		return xerrors.Errorf("invalid lang logo")
	}
	return nil
}

func dbRowToLang(row *ent.Lang) *npool.Lang {
	return &npool.Lang{
		ID:   row.ID.String(),
		Lang: row.Lang,
		Name: row.Name,
		Logo: row.Logo,
	}
}

func Add(ctx context.Context, in *npool.AddLangRequest) (*npool.AddLangResponse, error) {
	if err := validateLang(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Lang.
		Create().
		SetLang(in.GetInfo().GetLang()).
		SetName(in.GetInfo().GetName()).
		SetLogo(in.GetInfo().GetLogo()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create lang: %v", err)
	}

	return &npool.AddLangResponse{
		Info: dbRowToLang(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateLangRequest) (*npool.UpdateLangResponse, error) {
	return nil, nil
}

func GetAll(ctx context.Context, in *npool.GetLangsRequest) (*npool.GetLangsResponse, error) {
	return nil, nil
}
