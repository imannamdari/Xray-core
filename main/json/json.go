package json

import (
	"context"
	"io"

	"github.com/imannamdari/xray-core/common"
	"github.com/imannamdari/xray-core/common/cmdarg"
	"github.com/imannamdari/xray-core/common/errors"
	"github.com/imannamdari/xray-core/core"
	"github.com/imannamdari/xray-core/infra/conf"
	"github.com/imannamdari/xray-core/infra/conf/serial"
	"github.com/imannamdari/xray-core/main/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "JSON",
		Extension: []string{"json"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				cf := &conf.Config{}
				for i, arg := range v {
					errors.LogInfo(context.Background(), "Reading config: ", arg)
					r, err := confloader.LoadConfig(arg)
					if err != nil {
						return nil, errors.New("failed to read config: ", arg).Base(err)
					}
					c, err := serial.DecodeJSONConfig(r)
					if err != nil {
						return nil, errors.New("failed to decode config: ", arg).Base(err)
					}
					if i == 0 {
						// This ensure even if the muti-json parser do not support a setting,
						// It is still respected automatically for the first configure file
						*cf = *c
						continue
					}
					cf.Override(c, arg)
				}
				return cf.Build()
			case io.Reader:
				return serial.LoadJSONConfig(v)
			default:
				return nil, errors.New("unknown type")
			}
		},
	}))
}
