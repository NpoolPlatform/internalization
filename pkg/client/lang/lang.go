//nolint:nolintlint,dupl
package lang

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	npool "github.com/NpoolPlatform/message/npool/internationalization"

	constant "github.com/NpoolPlatform/internationalization/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.InternationalizationClient) (*npool.Lang, error)

func do(ctx context.Context, handler handler) (*npool.Lang, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewInternationalizationClient(conn)

	return handler(_ctx, cli)
}

func ExistLang(ctx context.Context, id string) (bool, error) {
	lang, err := do(ctx, func(_ctx context.Context, cli npool.InternationalizationClient) (*npool.Lang, error) {
		lang, err := cli.GetLang(ctx, &npool.GetLangRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return lang.Info, nil
	})
	if err != nil {
		return false, err
	}
	if lang != nil {
		return true, nil
	}
	return false, nil
}
