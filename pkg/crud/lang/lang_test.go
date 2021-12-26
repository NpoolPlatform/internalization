package lang

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/internationalization/message/npool"
	"github.com/NpoolPlatform/internationalization/pkg/test-init" //nolint

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

	inTesting = true
}

func assertLang(t *testing.T, actual, expected *npool.Lang) {
	assert.Equal(t, actual.Lang, expected.Lang)
	assert.Equal(t, actual.Name, expected.Name)
	assert.Equal(t, actual.Logo, expected.Logo)
}

func TestCRUD(t *testing.T) {
	id := uuid.New()
	lang := npool.Lang{
		Lang: fmt.Sprintf("zh_CN-%v", id),
		Name: fmt.Sprintf("Chinese-%v", id),
		Logo: fmt.Sprintf("ChineseLogo-%v", id),
	}

	resp, err := Add(context.Background(), &npool.AddLangRequest{
		Info: &lang,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertLang(t, resp.Info, &lang)
	}

	lang.Logo = fmt.Sprintf("ChaineseLogo-%v", uuid.New())
	lang.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateLangRequest{
		Info: &lang,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertLang(t, resp1.Info, &lang)
	}

	resp2, err := GetAll(context.Background(), &npool.GetLangsRequest{})
	if assert.Nil(t, err) {
		found := false
		for _, info := range resp2.Infos {
			if info.ID == lang.ID && info.Lang == lang.Lang && info.Name == lang.Name && info.Logo == lang.Logo {
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}
