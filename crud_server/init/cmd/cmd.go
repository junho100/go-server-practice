package cmd

import (
	"crud-server/config"
	"fmt"
)

type Cmd struct {
	config *config.Config
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	fmt.Println(c.config.Server.Port)

	return c
}
