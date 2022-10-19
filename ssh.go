package webssh

import (
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	defaultUser = "intel"
)

type (
	Option func(*SSHOption)
)

type SSHOption struct {
	HostAddr string
	User     string
	Password string
	KeyValue string
	Timeout  time.Duration
}

func WithHostAddr(host string) Option {
	return func(s *SSHOption) {
		if host != "" {
			s.HostAddr = host
		}
	}
}

func WithKeyValue(val string) Option {
	return func(s *SSHOption) {
		if val != "" {
			s.KeyValue = val
		}
	}
}

func WithUser(user string) Option {
	return func(s *SSHOption) {
		if user != "" {
			s.User = user
		}
	}
}

func WithTimeOut(t time.Duration) Option {
	return func(s *SSHOption) {
		s.Timeout = t
	}
}

func newSSHClient(opt ...Option) (*ssh.Client, error) {
	conf := &SSHOption{
		Timeout: time.Second * 5,
		User:    defaultUser,
	}
	for _, o := range opt {
		o(conf)
	}
	config := &ssh.ClientConfig{
		Timeout:         conf.Timeout,
		User:            conf.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //忽略know_hosts检查
	}
	if conf.KeyValue != "" {
		signer, err := ssh.ParsePrivateKey([]byte(conf.KeyValue))
		if err != nil {
			return nil, err
		}
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	} else {
		config.Auth = []ssh.AuthMethod{ssh.Password(conf.Password)}
	}

	c, err := ssh.Dial("tcp", conf.HostAddr, config)
	if err != nil {
		return nil, err
	}
	return c, nil
}
