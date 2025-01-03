package server

import (
	"github.com/hoangndst/vision/domain/constant"
	"github.com/hoangndst/vision/server"
	"github.com/hoangndst/vision/server/route"
)

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Port:        DefaultPort,
		Database:    DatabaseOptions{},
		LogFilePath: constant.DefaultLogFilePath,
	}
}

func (o *ServerOptions) Complete(args []string) {}

func (o *ServerOptions) Validate() error {
	return nil
}

func (o *ServerOptions) Config() (*server.Config, error) {
	cfg := server.NewConfig()
	o.Database.ApplyTo(cfg)
	cfg.Port = o.Port
	cfg.LogFilePath = o.LogFilePath
	cfg.GithubToken = o.GithubToken
	return cfg, nil
}

func (o *ServerOptions) Run() error {
	config, err := o.Config()
	if err != nil {
		return err
	}
	if _, err := route.NewRoute(config); err == nil {
		return nil
	}
	return nil
}
