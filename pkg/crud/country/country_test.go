package country

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/internationalization/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/internationalization"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func assertCountry(t *testing.T, actual, expected *npool.Country) {
	assert.Equal(t, actual.Country, expected.Country)
	assert.Equal(t, actual.Flag, expected.Flag)
	assert.Equal(t, actual.Code, expected.Code)
	assert.Equal(t, actual.Short, expected.Short)
}

func TestCRUD(t *testing.T) {
	id := uuid.New()
	country := npool.Country{
		Country: fmt.Sprintf("zh_CN-%v", id),
		Flag:    fmt.Sprintf("Chinese-%v", id),
		Code:    fmt.Sprintf("ChineseFlag-%v", id),
		Short:   fmt.Sprintf("ChineseShort-%v", id),
	}

	resp, err := Create(context.Background(), &npool.CreateCountryRequest{
		Info: &country,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertCountry(t, resp.Info, &country)
	}

	country.Flag = fmt.Sprintf("ChaineseFlag-%v", uuid.New())
	country.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateCountryRequest{
		Info: &country,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertCountry(t, resp1.Info, &country)
	}

	resp2, err := GetAll(context.Background(), &npool.GetCountriesRequest{})
	if assert.Nil(t, err) {
		found := false
		for _, info := range resp2.Infos {
			if info.ID == country.ID && info.Country == country.Country && info.Short == country.Short && info.Flag == country.Flag {
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}
