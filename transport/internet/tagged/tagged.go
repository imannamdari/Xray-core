package tagged

import (
	"context"

	"github.com/imannamdari/xray-core/common/net"
	"github.com/imannamdari/xray-core/features/routing"
)

type DialFunc func(ctx context.Context, dispatcher routing.Dispatcher, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
