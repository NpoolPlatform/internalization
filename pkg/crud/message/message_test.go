package message

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
		Message:   fmt.Sprintf("Test Test Test %v", uuid.New().String()),
	}

	resp, err := CreateMessage(context.Background(), &npool.CreateMessageRequest{
		Info: &message,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertMessage(t, resp.Info, &message)
	}

	message1 := npool.Message{
		AppID:     message.AppID,
		LangID:    message.LangID,
		MessageID: uuid.New().String(),
		BatchGet:  true,
		Message:   fmt.Sprintf("Test Test Test Test %v", uuid.New().String()),
	}
	message2 := npool.Message{
		AppID:     message.AppID,
		LangID:    message.LangID,
		MessageID: uuid.New().String(),
		BatchGet:  true,
		Message:   fmt.Sprintf("Test Test Test Test %v", uuid.New().String()),
	}
	resp1, err := CreateMessages(context.Background(), &npool.CreateMessagesRequest{
		Infos: []*npool.Message{&message1, &message2},
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp1.Infos[0].ID, uuid.UUID{}.String())
		assertMessage(t, resp1.Infos[0], &message1)

		assert.NotEqual(t, resp1.Infos[1].ID, uuid.UUID{}.String())
		assertMessage(t, resp1.Infos[1], &message2)
	}

	message1.BatchGet = false
	message1.ID = resp1.Infos[0].ID
	message2.MessageID = uuid.New().String()
	message2.ID = resp1.Infos[1].ID

	resp2, err := UpdateMessages(context.Background(), &npool.UpdateMessagesRequest{
		Infos: []*npool.Message{&message1, &message2},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Infos[0].ID, message1.ID)
		assertMessage(t, resp2.Infos[0], &message1)

		assert.Equal(t, resp2.Infos[1].ID, message2.ID)
		assertMessage(t, resp2.Infos[1], &message2)
	}

	resp3, err := GetMessagesByAppLang(context.Background(), &npool.GetMessagesByAppLangRequest{
		AppID:  message1.AppID,
		LangID: message1.LangID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp3.Infos), 0)
	}

	resp4, err := GetMessageByAppLangMessage(context.Background(), &npool.GetMessageByAppLangMessageRequest{
		AppID:     message1.AppID,
		LangID:    message1.LangID,
		MessageID: message1.MessageID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.Info.ID, message1.ID)
		assertMessage(t, resp4.Info, &message1)
	}
}
