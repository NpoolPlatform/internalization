package message

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

func assertMessage(t *testing.T, actual, expected *npool.Message) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.LangID, expected.LangID)
	assert.Equal(t, actual.MessageID, expected.MessageID)
	assert.Equal(t, actual.BatchGet, expected.BatchGet)
	assert.Equal(t, actual.Message, expected.Message)
}

func TestCRUD(t *testing.T) {
	message := npool.Message{
		AppID:     uuid.New().String(),
		LangID:    uuid.New().String(),
		MessageID: uuid.New().String(),
		BatchGet:  true,
		Message:   "Test Test Test",
	}

	resp, err := CreateMessage(context.Background(), &npool.CreateMessageRequest{
		Info: &message,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertMessage(t, resp.Info, &message)
	}
}
