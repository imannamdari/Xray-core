package splithttp

import (
	"context"

	"github.com/imannamdari/xray-core/common"
	"github.com/imannamdari/xray-core/common/errors"
)

//go:generate go run github.com/imannamdari/xray-core/common/errors/errorgen

const protocolName = "splithttp"

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return nil, errors.New("splithttp is a transport protocol.")
	}))
}
