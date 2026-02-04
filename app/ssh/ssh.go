package ssh

import (
	"context"
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"

	"github.com/armon/go-socks5"
	"github.com/sirupsen/logrus"

	"github.com/imannamdari/xray-core/common"
)

type SSH struct {
	enabled     bool
	ip          string
	user        string
	password    string
	runningPort int32
}

func New(ctx context.Context, config *Config) (*SSH, error) {
	sshS := &SSH{
		enabled:     config.GetEnabled(),
		ip:          config.GetIp(),
		user:        config.GetUser(),
		password:    config.GetPassword(),
		runningPort: config.GetRunningPort(),
	}
	if sshS.enabled {
		if err := sshS.startTunnel(); err != nil {
			logrus.WithError(err).Error("failed to start ssh tunnel")
			return nil, fmt.Errorf("failed to start tunnel: %w", err)
		}
	}
	return sshS, nil
}

func (s *SSH) startTunnel() error {
	cfg := &ssh.ClientConfig{
		User:            s.user,
		Auth:            []ssh.AuthMethod{ssh.Password(s.password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	c, err := ssh.Dial("tcp", s.ip+":22", cfg)
	if err != nil {
		return fmt.Errorf("failed to dial ssh tcp: %w", err)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:"+fmt.Sprintf("%d", s.runningPort))
	if err != nil {
		return fmt.Errorf("failed to listen on %d: %w", s.runningPort, err)
	}
	go func() {
		for {
			localConn, _ := listener.Accept()
			go func() {
				if err := s.handleSocks(c, localConn); err != nil {
					logrus.WithError(err).Error("failed to handle socks")
				}
			}()
		}
	}()
	return nil
}

func (s *SSH) handleSocks(sshC *ssh.Client, localConn net.Conn) error {
	conf := &socks5.Config{
		Dial: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return sshC.Dial(network, addr)
		},
	}
	server, err := socks5.New(conf)
	if err != nil {
		return fmt.Errorf("failed to create new socks: %w", err)
	}
	if err := server.ServeConn(localConn); err != nil {
		return fmt.Errorf("failed to server connection: %w", err)
	}
	return nil
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return New(ctx, config.(*Config))
	}))
}
