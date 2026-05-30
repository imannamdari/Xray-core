package outbound

import (
	"github.com/imannamdari/xray-core/common/net"
	"github.com/imannamdari/xray-core/common/protocol"
)

// As a stub command consumer.
func (h *Handler) handleCommand(dest net.Destination, cmd protocol.ResponseCommand) {
	switch cmd.(type) {
	default:
	}
}
