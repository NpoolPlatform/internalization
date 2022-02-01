package applang

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

func assertAppLang(t *testing.T, actual, expected *npool.AppLang) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.LangID, expected.LangID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	lang := npool.AppLang{
		AppID:  uuid.New().String(),
		LangID: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateAppLangRequest{
		Info: &lang,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertAppLang(t, resp.Info, &lang)
	}

	resp1, err := Get(context.Background(), &npool.GetAppLangRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Info.ID, resp1.Info.ID)
		assertAppLang(t, resp1.Info, &lang)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetAppLangsByAppRequest{
		AppID: lang.AppID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp2.Infos), 0)
	}
}
