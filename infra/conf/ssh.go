package conf

import (
	"github.com/imannamdari/xray-core/app/ssh"
)

type SSHConfig struct {
	Enabled     bool   `json:"enabled"`
	IP          string `json:"ip"`
	User        string `json:"user"`
	Password    string `json:"password"`
	RunningPort int32  `json:"runningPort"`
}

func (c *SSHConfig) Build() (*ssh.Config, error) {
	return &ssh.Config{
		Enabled:     c.Enabled,
		Ip:          c.IP,
		User:        c.User,
		Password:    c.Password,
		RunningPort: c.RunningPort,
	}, nil
}
