package ssh

import (
	"github.com/xtls/xray-core/features"
)

type SSHEngine interface {
	features.Feature
}