package network

import "github.com/gin-gonic/gin"

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(port)
}

func (n *Network) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.GET(path, handler...)
}

func (n *Network) registerPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.POST(path, handler...)
}
func (n *Network) registerUPDATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.PUT(path, handler...)
}

func (n *Network) registerDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.DELETE(path, handler...)
}
