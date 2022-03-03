package lang

import (
	"context"
	"time"

	constant "github.com/NpoolPlatform/internationalization/pkg/const"
	"github.com/NpoolPlatform/internationalization/pkg/db"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/lang"
	npool "github.com/NpoolPlatform/message/npool/internationalization"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

const (
	dbTimeout = 30 * time.Second
)

var inTesting = false

func validateLang(info *npool.Lang) error {
	if info.Lang == "" {
		return xerrors.Errorf("invalid lang")
	}

	if !inTesting {
		if _, ok := constant.Langs[info.Lang]; !ok {
			return xerrors.Errorf("lang not supported")
		}
	}

	if info.Logo == "" {
		return xerrors.Errorf("invalid lang logo")
	}
	return nil
}

func dbRowToLang(row *ent.Lang) *npool.Lang {
	return &npool.Lang{
		ID:    row.ID.String(),
		Lang:  row.Lang,
		Name:  row.Name,
		Logo:  row.Logo,
		Short: row.Short,
	}
}

func Add(ctx context.Context, in *npool.AddLangRequest) (*npool.AddLangResponse, error) {
	if err := validateLang(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	if !inTesting {
		in.GetInfo().Name = constant.Langs[in.GetInfo().GetLang()].Name
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
		SetShort(in.GetInfo().GetShort()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create lang: %v", err)
	}

	return &npool.AddLangResponse{
		Info: dbRowToLang(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateLangRequest) (*npool.UpdateLangResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid lang id: %v", err)
	}

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
		UpdateOneID(id).
		SetLogo(in.GetInfo().GetLogo()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update lang: %v", err)
	}

	return &npool.UpdateLangResponse{
		Info: dbRowToLang(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetLangRequest) (*npool.GetLangResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid lang id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Lang.
		Query().
		Where(
			lang.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query langs: %v", err)
	}

	var _lang *npool.Lang
	for _, info := range infos {
		_lang = dbRowToLang(info)
		break
	}

	return &npool.GetLangResponse{
		Info: _lang,
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetLangsRequest) (*npool.GetLangsResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Lang.
		Query().
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query langs: %v", err)
	}

	langs := []*npool.Lang{}
	for _, info := range infos {
		langs = append(langs, dbRowToLang(info))
	}

	return &npool.GetLangsResponse{
		Infos: langs,
	}, nil
}
