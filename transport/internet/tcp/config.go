package tcp

import (
	"github.com/imannamdari/xray-core/common"
	"github.com/imannamdari/xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
