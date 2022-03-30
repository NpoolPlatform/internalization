package country

import (
	"context"
	"time"

	"github.com/NpoolPlatform/internationalization/pkg/db"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/country"
	npool "github.com/NpoolPlatform/message/npool/internationalization"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

const (
	dbTimeout = 30 * time.Second
)

func validateCountry(info *npool.Country) error {
	if info.Country == "" {
		return xerrors.Errorf("invalid country")
	}
	if info.Flag == "" {
		return xerrors.Errorf("invalid country flag")
	}
	if info.Short == "" {
		return xerrors.Errorf("invalid country short")
	}
	if info.Code == "" {
		return xerrors.Errorf("invalid country code")
	}
	return nil
}

func dbRowToCountry(row *ent.Country) *npool.Country {
	return &npool.Country{
		ID:      row.ID.String(),
		Country: row.Country,
		Short:   row.Short,
		Flag:    row.Flag,
		Code:    row.Code,
	}
}

func Create(ctx context.Context, in *npool.CreateCountryRequest) (*npool.CreateCountryResponse, error) {
	if err := validateCountry(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Country.
		Create().
		SetCountry(in.GetInfo().GetCountry()).
		SetFlag(in.GetInfo().GetFlag()).
		SetCode(in.GetInfo().GetCode()).
		SetShort(in.GetInfo().GetShort()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create country: %v", err)
	}

	return &npool.CreateCountryResponse{
		Info: dbRowToCountry(info),
	}, nil
}

func CreateCountries(ctx context.Context, in *npool.CreateCountriesRequest) (*npool.CreateCountriesResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	bulk := make([]*ent.CountryCreate, len(in.GetInfos()))
	for i, info := range in.GetInfos() {
		if err := validateCountry(info); err != nil {
			return nil, xerrors.Errorf("invalid parameter: %v", err)
		}

		bulk[i] = cli.
			Country.
			Create().
			SetCountry(info.GetCountry()).
			SetFlag(info.GetFlag()).
			SetCode(info.GetCode()).
			SetShort(info.GetShort())
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Country.
		CreateBulk(bulk...).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create bulk countries: %v", err)
	}

	countries := []*npool.Country{}
	for _, info := range infos {
		countries = append(countries, dbRowToCountry(info))
	}

	return &npool.CreateCountriesResponse{
		Infos: countries,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateCountryRequest) (*npool.UpdateCountryResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid country id: %v", err)
	}

	if err := validateCountry(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Country.
		UpdateOneID(id).
		SetCountry(in.GetInfo().GetCountry()).
		SetFlag(in.GetInfo().GetFlag()).
		SetCode(in.GetInfo().GetCode()).
		SetShort(in.GetInfo().GetShort()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update country: %v", err)
	}

	return &npool.UpdateCountryResponse{
		Info: dbRowToCountry(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetCountryRequest) (*npool.GetCountryResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid country id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Country.
		Query().
		Where(
			country.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query country: %v", err)
	}

	var _country *npool.Country
	for _, info := range infos {
		_country = dbRowToCountry(info)
		break
	}

	return &npool.GetCountryResponse{
		Info: _country,
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetCountriesRequest) (*npool.GetCountriesResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Country.
		Query().
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query countries: %v", err)
	}

	countrys := []*npool.Country{}
	for _, info := range infos {
		countrys = append(countrys, dbRowToCountry(info))
	}

	return &npool.GetCountriesResponse{
		Infos: countrys,
	}, nil
}
