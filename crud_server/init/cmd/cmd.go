package cmd

import (
	"crud-server/config"
	"crud-server/network"
	"crud-server/repository"
	"crud-server/service"
)

type Cmd struct {
	config    *config.Config
	network   *network.Network
	repositoy *repository.Repository
	service   *service.Service
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	c.repositoy = repository.Newrepository()
	c.service = service.NewService(c.repositoy)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
