package message

import (
	"context"
	"time"

	"github.com/NpoolPlatform/internationalization/pkg/db"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/message"
	npool "github.com/NpoolPlatform/message/npool/internationalization"

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
		ID:        row.ID.String(),
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

	langID := uuid.UUID{}
	appID := uuid.UUID{}

	bulk := make([]*ent.MessageCreate, len(in.GetInfos()))
	for i, info := range in.GetInfos() {
		if err := validateMessage(info); err != nil {
			return nil, xerrors.Errorf("invalid parameter: %v", err)
		}

		langID = uuid.MustParse(info.GetLangID())
		appID = uuid.MustParse(info.GetAppID())

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

	err = cli.
		Message.
		CreateBulk(bulk...).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create bulk messages: %v", err)
	}

	infos, err := cli.
		Message.
		Query().
		Where(
			message.And(
				message.AppID(appID),
				message.LangID(langID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("Fail query message: %v", err)
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
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	if err := validateMessage(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Message.
		UpdateOneID(id).
		SetMessageID(in.GetInfo().GetMessageID()).
		SetMessage(in.GetInfo().GetMessage()).
		SetBatchGet(in.GetInfo().GetBatchGet()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update message: %v", err)
	}

	return &npool.UpdateMessageResponse{
		Info: dbRowToMessage(info),
	}, nil
}

func UpdateMessages(ctx context.Context, in *npool.UpdateMessagesRequest) (*npool.UpdateMessagesResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	tx, err := cli.Tx(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail get transaction: %v", err)
	}

	msgs := []*npool.Message{}
	for _, info := range in.GetInfos() {
		if err := validateMessage(info); err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return nil, xerrors.Errorf("fail rollback update message: %v, %v", rerr, err)
			}
			return nil, xerrors.Errorf("fail update message: %v", err)
		}

		id, err := uuid.Parse(info.GetID())
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return nil, xerrors.Errorf("fail rollback update message: %v, %v", rerr, err)
			}
			return nil, xerrors.Errorf("invalid id: %v", err)
		}

		msg, err := tx.
			Message.
			UpdateOneID(id).
			SetMessageID(info.GetMessageID()).
			SetMessage(info.GetMessage()).
			SetBatchGet(info.GetBatchGet()).
			Save(ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return nil, xerrors.Errorf("fail rollback update message: %v, %v", rerr, err)
			}
			return nil, xerrors.Errorf("fail update message: %v", err)
		}

		msgs = append(msgs, dbRowToMessage(msg))
	}

	err = tx.Commit()
	if err != nil {
		return nil, xerrors.Errorf("fail commit update message: %v", err)
	}

	return &npool.UpdateMessagesResponse{
		Infos: msgs,
	}, nil
}

func GetMessagesByAppLang(ctx context.Context, in *npool.GetMessagesByAppLangRequest) (*npool.GetMessagesByAppLangResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	langID, err := uuid.Parse(in.GetLangID())
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
		Message.
		Query().
		Where(
			message.And(
				message.AppID(appID),
				message.LangID(langID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query message by lang id: %v", err)
	}

	msgs := []*npool.Message{}
	for _, info := range infos {
		msgs = append(msgs, dbRowToMessage(info))
	}

	return &npool.GetMessagesByAppLangResponse{
		Infos: msgs,
	}, nil
}

func GetMessageByAppLangMessage(ctx context.Context, in *npool.GetMessageByAppLangMessageRequest) (*npool.GetMessageByAppLangMessageResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	langID, err := uuid.Parse(in.GetLangID())
	if err != nil {
		return nil, xerrors.Errorf("invalid lang id: %v", err)
	}

	if in.GetMessageID() == "" {
		return nil, xerrors.Errorf("invalid message id")
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Message.
		Query().
		Where(
			message.And(
				message.AppID(appID),
				message.LangID(langID),
				message.MessageID(in.GetMessageID()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query message by lang id message id: %v", err)
	}

	var msg *npool.Message
	for _, info := range infos {
		msg = dbRowToMessage(info)
		break
	}

	return &npool.GetMessageByAppLangMessageResponse{
		Info: msg,
	}, nil
}
