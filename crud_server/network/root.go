package network

import "github.com/gin-gonic/gin"

type Network struct {
	engin *gin.Engine
}

func NewNetwork() *Network {
	r := &Network{
		engin: gin.New(),
	}

	newUserRouter(r)

	return r
}
