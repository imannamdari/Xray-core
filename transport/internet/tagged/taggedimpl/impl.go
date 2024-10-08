package taggedimpl

import (
	"context"

	"github.com/imannamdari/xray-core/common/errors"
	"github.com/imannamdari/xray-core/common/net"
	"github.com/imannamdari/xray-core/common/net/cnc"
	"github.com/imannamdari/xray-core/common/session"
	"github.com/imannamdari/xray-core/core"
	"github.com/imannamdari/xray-core/features/routing"
	"github.com/imannamdari/xray-core/transport/internet/tagged"
)

func DialTaggedOutbound(ctx context.Context, dest net.Destination, tag string) (net.Conn, error) {
	var dispatcher routing.Dispatcher
	if core.FromContext(ctx) == nil {
		return nil, errors.New("Instance context variable is not in context, dial denied. ")
	}
	if err := core.RequireFeatures(ctx, func(dispatcherInstance routing.Dispatcher) {
		dispatcher = dispatcherInstance
	}); err != nil {
		return nil, errors.New("Required Feature dispatcher not resolved").Base(err)
	}

	content := new(session.Content)
	content.SkipDNSResolve = true

	ctx = session.ContextWithContent(ctx, content)
	ctx = session.SetForcedOutboundTagToContext(ctx, tag)

	r, err := dispatcher.Dispatch(ctx, dest)
	if err != nil {
		return nil, err
	}
	var readerOpt cnc.ConnectionOption
	if dest.Network == net.Network_TCP {
		readerOpt = cnc.ConnectionOutputMulti(r.Reader)
	} else {
		readerOpt = cnc.ConnectionOutputMultiUDP(r.Reader)
	}
	return cnc.NewConnection(cnc.ConnectionInputMulti(r.Writer), readerOpt), nil
}

func init() {
	tagged.Dialer = DialTaggedOutbound
}
