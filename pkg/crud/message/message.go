package message

import (
	"context"
	"time"

	"github.com/NpoolPlatform/internationalization/message/npool"
	"github.com/NpoolPlatform/internationalization/pkg/db"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

const (
	dbTimeout = 30 * time.Second
)

func validateMessage(info *npool.Message) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetLangID()); err != nil {
		return xerrors.Errorf("invalid lang id: %v", err)
	}
	if info.GetMessageID() == "" {
		return xerrors.Errorf("invalid message id")
	}
	if info.GetMessage() == "" {
		return xerrors.Errorf("empty message")
	}
	return nil
}

func dbRowToMessage(row *ent.Message) *npool.Message {
	return &npool.Message{
		AppID:     row.AppID.String(),
		LangID:    row.LangID.String(),
		MessageID: row.MessageID,
		BatchGet:  row.BatchGet,
		Message:   row.Message,
	}
}

func CreateMessage(ctx context.Context, in *npool.CreateMessageRequest) (*npool.CreateMessageResponse, error) {
	if err := validateMessage(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameters: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Message.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetLangID(uuid.MustParse(in.GetInfo().GetLangID())).
		SetMessageID(in.GetInfo().GetMessageID()).
		SetBatchGet(in.GetInfo().GetBatchGet()).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create message: %v", err)
	}

	return &npool.CreateMessageResponse{
		Info: dbRowToMessage(info),
	}, nil
}

func CreateMessages(ctx context.Context, in *npool.CreateMessagesRequest) (*npool.CreateMessagesResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	bulk := make([]*ent.MessageCreate, len(in.GetInfos()))
	for i, info := range in.GetInfos() {
		if err := validateMessage(info); err != nil {
			return nil, xerrors.Errorf("invalid parameter: %v", err)
		}

		bulk[i] = cli.
			Message.
			Create().
			SetAppID(uuid.MustParse(info.GetAppID())).
			SetLangID(uuid.MustParse(info.GetLangID())).
			SetMessageID(info.GetMessageID()).
			SetBatchGet(info.GetBatchGet()).
			SetMessage(info.GetMessage())
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Message.
		CreateBulk(bulk...).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create bulk messages: %v", err)
	}

	msgs := []*npool.Message{}
	for _, info := range infos {
		msgs = append(msgs, dbRowToMessage(info))
	}

	return &npool.CreateMessagesResponse{
		Infos: msgs,
	}, nil
}

func UpdateMessage(ctx context.Context, in *npool.UpdateMessageRequest) (*npool.UpdateMessageResponse, error) {
	return nil, nil
}

func UpdateMessages(ctx context.Context, in *npool.UpdateMessagesRequest) (*npool.UpdateMessagesResponse, error) {
	return nil, nil
}

func GetMessages(ctx context.Context, in *npool.GetMessagesRequest) (*npool.GetMessagesResponse, error) {
	return nil, nil
}

func GetMessageByMessageID(ctx context.Context, in *npool.GetMessageByMessageIDRequest) (*npool.GetMessageByMessageIDResponse, error) {
	return nil, nil
}
