package udp

import (
	"github.com/imannamdari/xray-core/common"
	"github.com/imannamdari/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
