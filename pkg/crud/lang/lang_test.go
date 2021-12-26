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
}
